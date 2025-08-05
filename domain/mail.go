package domain

import (
	"github.com/hossainabid/go-ims/types"
)

type (
	MailService interface {
		SendEmail(reqData types.EmailPayload) error
		SendLowStockEmail(productId int) error
	}

	MailRepository interface {
		SendEmail(reqData *types.EmailPayload) error
	}
)
