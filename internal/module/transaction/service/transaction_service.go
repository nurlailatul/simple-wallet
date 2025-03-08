package service

import (
	"context"
	"errors"
	"fmt"
	"simple-wallet/internal/infrastructure/database"
	balanceHistoryDomain "simple-wallet/internal/module/balance_history/domain"
	balanceHistoryRepository "simple-wallet/internal/module/balance_history/repository"
	"simple-wallet/internal/module/transaction/domain"
	"simple-wallet/internal/module/transaction/repository"
	walletRepository "simple-wallet/internal/module/wallet/repository"
	"time"

	log "github.com/sirupsen/logrus"
)

var logFormat = "[TransactionService][%v][UserID:%v]"

type TransactionService struct {
	balanceHistoryRepo balanceHistoryRepository.BalanceHistoryRepositoryInterface
	repo               repository.TransactionRepositoryInterface
	walletRepo         walletRepository.WalletRepositoryInterface
	connection         database.MySQLDBInterface
}

func NewTransactionService(
	balanceHistoryRepo balanceHistoryRepository.BalanceHistoryRepositoryInterface,
	repo repository.TransactionRepositoryInterface,
	walletRepo walletRepository.WalletRepositoryInterface,
	connection database.MySQLDBInterface,
) TransactionServiceInterface {
	return &TransactionService{
		balanceHistoryRepo: balanceHistoryRepo,
		repo:               repo,
		walletRepo:         walletRepo,
		connection:         connection,
	}
}

func (s *TransactionService) GetByReferenceID(ctx context.Context, referenceID string) *domain.TransactionEntity {
	return s.repo.GetByReferenceID(ctx, referenceID)
}

func (s *TransactionService) DeductBalance(ctx context.Context, request domain.DeductBalanceRequest) error {
	// log.SetFormatter(&log.JSONFormatter{})
	logf := fmt.Sprintf(logFormat, "DeductBalance", request.UserID)
	message := ""

	// Start a transaction
	tx := s.connection.BeginTx(ctx, s.connection.GetConnection().DB)
	defer func() {
		if r := recover(); r != nil {
			s.connection.RollbackTx(tx)
		}
	}()

	wallet := s.walletRepo.GetByUserIDForLocking(tx, request.WalletID)
	if wallet == nil {
		return errors.New("record not found")
	}

	if wallet.Balance < request.Amount {
		return errors.New("insufficient balance")
	}

	// Deduct balance and update wallet
	originAmount := wallet.Balance
	wallet.Balance -= request.Amount
	finalAmount := wallet.Balance

	err := s.walletRepo.Update(tx, wallet)
	if err != nil {
		s.connection.RollbackTx(tx)

		message = "error deduct wallet balance"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return errors.New(message)
	}

	trx := domain.TransactionEntity{
		WalletID:              request.WalletID,
		Amount:                request.Amount,
		ReceiverBank:          request.ReceiverBank,
		ReceiverAccountNumber: request.ReceiverAccountNumber,
		Status:                domain.StatusPending,
		ReferenceID:           request.ReferenceID,
		CreatedAt:             time.Now(),
	}

	trxID, err := s.repo.Create(tx, trx)
	if err != nil {
		s.connection.RollbackTx(tx)

		message = "error create deduct transaction"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return errors.New(message)
	}

	history := balanceHistoryDomain.BalanceHistoryEntity{
		WalletID:        request.WalletID,
		TransactionID:   trxID,
		TransactionType: 1,
		OriginAmount:    originAmount,
		Amount:          request.Amount,
		OperationType:   1,
		FinalAmount:     finalAmount,
	}

	s.connection.RollbackTx(tx)

	if err = s.balanceHistoryRepo.Create(tx, history); err != nil {
		s.connection.RollbackTx(tx)

		message = "error create balance history"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return errors.New(message)
	}

	s.connection.CommitTx(tx)

	message = "Success deduct balance"
	log.Infof("%v %v: %v", logf, message, "")

	return nil
}
