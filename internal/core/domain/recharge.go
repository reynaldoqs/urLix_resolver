package domain

import "time"

// Recharge is a global representation of an recharge
type Recharge struct {
	ID          string    `json:"id"`
	PhoneNumber uint32    `json:"phoneNumber" validate:"required,gte=9999999,lte=100000000"`
	Company     string    `json:"company" validate:"required,oneof=entel viva tigo"`
	CardNumber  uint64    `json:"cardNumber" validate:"required,min=99999999999999,max=10000000000000000"`
	Status      uint8     `json:"status"`
	IDResolver  string    `json:"idResolver"`
	Mount       int       `json:"mount" validate:"gte=10"`
	CreatedAt   time.Time `json:"createdAt"`
	ResolvedAt  time.Time `json:"resolvedAt"`
}
