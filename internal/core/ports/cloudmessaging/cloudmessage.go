package messagingport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type CloudMessenger interface {
	RechargeNotify(farmer *domain.Farmer, fcmsg *domain.RechargeMessage) error
	AdminNotify(farmer *domain.Farmer, fcmsg *domain.AdminMessage) error
}
