package domain

type Farmer struct {
	MsgToken       string `firestore:"msgToken"`
	DeviceID       string `firestore:"deviceId"`
	PhoneNumber    int    `firestore:"phoneNumber"`
	Company        string `firestore:"company"`
	XDay           int    `firestore:"xDay"`
	SubscriptionID int    `firestore:"subscriptionId"`
	Status         int    `firestore:"status"`
	IsOnline       bool   `firestore:"isOnline"`
}

type FarmerCloudMessage struct {
	ExecCodes    []string `json:"execCodes"`
	Company      string   `json:"company"`
	IDRecharge   string   `json:"idRecharge"`
	Mount        int      `json:"mount"`
	FarmerNumber int      `json:"farmerNumber"`
}
