package services

import (
	"fmt"

	"github.com/hossainabid/go-ims/domain"
	"github.com/hossainabid/go-ims/logger"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/worker"
)

type Mail struct {
	userRepo    domain.UserRepository
	productRepo domain.ProductRepository
	mailRepo    domain.MailRepository
	workerPool  *worker.Pool
}

func NewMailService(userRepo domain.UserRepository, productRepo domain.ProductRepository, mailRepo domain.MailRepository, workerPool *worker.Pool) *Mail {
	return &Mail{
		userRepo:    userRepo,
		productRepo: productRepo,
		mailRepo:    mailRepo,
		workerPool:  workerPool,
	}
}

func (m *Mail) SendEmail(reqData types.EmailPayload) error {
	err := m.mailRepo.SendEmail(&reqData)
	if err != nil {
		logger.Error(fmt.Sprintf("err: []%v occurred while sending email to: %s", err, reqData.MailTo))
		return err
	}
	return nil
}

func (m *Mail) SendLowStockEmail(productId int) error {
	product, err := m.productRepo.ReadProductByID(productId)
	if err != nil {
		return err
	}

	if product.ThresholdQty > product.WarehouseQty {
		user, err := m.userRepo.ReadUser(product.CreatedBy)
		if err != nil {
			return err
		}

		emailPayload := types.EmailPayload{
			MailTo:  user.Email,
			Subject: "Low stock notification for product: " + product.Name,
			Body: map[string]interface{}{
				"product": product,
			},
		}

		// Add the email sending task to the worker pool
		task := worker.NewTask(func() error {
			return m.SendEmail(emailPayload)
		}, func(err error) {
			logger.Error("Failed to send email: ", err, " to user: ", user.Email)
		}, 3)

		m.workerPool.AddTask(task)
	}

	return nil
}
