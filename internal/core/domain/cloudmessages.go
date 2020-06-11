package domain

type RechargeMessage struct {
	ExecCodes     []string `json:"execCodes"`
	TargetCompany string   `json:"targetCompany"`
	IDRecharge    string   `json:"idRecharge"`
	Mount         int      `json:"mount"`
	FarmerNumber  int      `json:"farmerNumber"`
}

type AdminMessage struct {
	ExecCodes    []string `json:"execCodes"`
	IDMessage    string   `json:"idMessage"`
	FarmerNumber int      `json:"farmerNumber"`
}
