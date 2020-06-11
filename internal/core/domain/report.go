package domain

type RechargeReport struct {
	IDRecharge    string         `json:"idRecharge"`
	FarmerNumber  int            `json:"farmerNumber"`
	Successful    bool           `json:"successful"`
	CodeResponses []codeResponse `json:"codeRespones"`
}

type codeResponse struct {
	Code     string `json:"code"`
	Response string `json:"response"`
}
