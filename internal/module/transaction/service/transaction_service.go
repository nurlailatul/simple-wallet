package service

import (
	"context"
	"errors"
	"fmt"
	balanceHistoryDomain "simple-wallet/internal/module/balance_history/domain"
	balanceHistoryRepository "simple-wallet/internal/module/balance_history/repository"
	"simple-wallet/internal/module/transaction/domain"
	"simple-wallet/internal/module/transaction/repository"
	walletRepository "simple-wallet/internal/module/wallet/repository"
	"simple-wallet/pkg/db"
	"time"

	log "github.com/sirupsen/logrus"
)

var logFormat = "[TransactionService][%v][UserID:%v]"

type TransactionService struct {
	repo      repository.TransactionRepositoryInterface
	dbWrapper *db.GormDBWrapper
}

func NewTransactionService(
	repo repository.TransactionRepositoryInterface,
	dbWrapper *db.GormDBWrapper,
) TransactionServiceInterface {
	return &TransactionService{
		dbWrapper: dbWrapper,
		repo:      repo,
	}
}

func (s *TransactionService) GetByReferenceID(ctx context.Context, referenceID string) *domain.TransactionEntity {
	return s.repo.GetByReferenceID(ctx, referenceID)
}

func (s *TransactionService) DeductBalance(ctx context.Context, request domain.DeductBalanceRequest) error {
	logf := fmt.Sprintf(logFormat, "DeductBalance", request.UserID)
	var (
		err     error
		message string
	)

	tx := s.dbWrapper.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	walletRepo := walletRepository.NewWalletRepository(tx)
	wallet := walletRepo.GetByUserIDForLocking(ctx, request.WalletID)
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

	err = walletRepo.Update(ctx, wallet)
	if err != nil {
		tx.Rollback()
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

	trxID, err := repository.NewTransactionRepository(tx).Create(ctx, trx)
	if err != nil {
		tx.Rollback()
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

	if err = balanceHistoryRepository.NewBalanceHistoryRepository(tx).Create(ctx, history); err != nil {
		tx.Rollback()

		message = "error create balance history"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return errors.New(message)
	}

	err = tx.Commit().Error
	if err != nil {
		message = "error commit transaction"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return errors.New(message)
	}

	message = "Success deduct balance"
	log.Infof("%v %v: %v", logf, message, "")

	return nil
}
