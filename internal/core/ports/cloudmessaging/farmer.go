package messagingport

import "github.com/reynaldoqs/urLix_resolver/internal/core/domain"

type FarmerMessenger interface {
	Notify(farmer *domain.Farmer, fcmsg *domain.FarmerCloudMessage) error
}
