package service

import (
	"context"
	"errors"
	"fmt"
	balanceHistoryDomain "simple-wallet/internal/module/balance_history/domain"
	balanceHistoryRepository "simple-wallet/internal/module/balance_history/repository"
	"simple-wallet/internal/module/transaction/domain"
	"simple-wallet/internal/module/transaction/repository"
	userDomain "simple-wallet/internal/module/user/domain"
	userRepository "simple-wallet/internal/module/user/repository"
	walletRepository "simple-wallet/internal/module/wallet/repository"
	"simple-wallet/pkg/db"
	"time"

	log "github.com/sirupsen/logrus"
)

var logFormat = "[TransactionService][%v][UserID:%v]"

type TransactionService struct {
	repo      repository.TransactionRepositoryInterface
	userRepo  userRepository.UserRepositoryInterface
	dbWrapper *db.GormDBWrapper
}

func NewTransactionService(
	repo repository.TransactionRepositoryInterface,
	userRepo userRepository.UserRepositoryInterface,
	dbWrapper *db.GormDBWrapper,
) TransactionServiceInterface {
	return &TransactionService{
		dbWrapper: dbWrapper,
		repo:      repo,
		userRepo:  userRepo,
	}
}

func (s *TransactionService) GetByReferenceID(ctx context.Context, referenceID string) *domain.TransactionEntity {
	return s.repo.GetByReferenceID(ctx, referenceID)
}

func (s *TransactionService) DeductBalance(ctx context.Context, request domain.DeductBalanceRequest) (data *domain.DeductBalanceResponse, err error) {
	logf := fmt.Sprintf(logFormat, "DeductBalance", request.UserID)
	message := ""

	if request.User == nil {
		request.User = s.userRepo.GetByID(ctx, request.UserID)
	}

	if request.User == nil {
		return nil, errors.New("user not found")
	}

	if request.User.Status != userDomain.StatusActive {
		return nil, errors.New("user not active")
	}

	tx := s.dbWrapper.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	walletRepo := walletRepository.NewWalletRepository(tx)
	wallet := walletRepo.GetByUserIDForLocking(ctx, request.WalletID)
	if wallet == nil {
		return nil, errors.New("record not found")
	}

	if wallet.Balance < request.Amount {
		return nil, errors.New("insufficient balance")
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
		return nil, errors.New(message)
	}

	trx := domain.TransactionEntity{
		WalletID:              request.WalletID,
		Amount:                request.Amount,
		ReceiverBank:          request.ReceiverBank,
		ReceiverAccountNumber: request.ReceiverAccountNumber,
		Status:                domain.StatusPending,
		ReferenceID:           request.ReferenceID,
		CreatedAt:             uint(time.Now().Unix()),
	}

	trxID, err := repository.NewTransactionRepository(tx).Create(ctx, trx)
	if err != nil {
		tx.Rollback()
		message = "error create deduct transaction"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return nil, errors.New(message)
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
		return nil, errors.New(message)
	}

	err = tx.Commit().Error
	if err != nil {
		message = "error commit transaction"
		log.Errorf("%v %v: %v", logf, message, err.Error())
		return nil, errors.New(message)
	}

	message = "Success deduct balance"
	log.Infof("%v %v: %v", logf, message, "")

	data = &domain.DeductBalanceResponse{
		WalletID:   request.WalletID,
		NewBalance: wallet.Balance,
		Status:     trx.Status,
		CreatedAt:  trx.CreatedAt,
	}

	return data, nil
}
