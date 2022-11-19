package requests

type BusinessPartnerCustomer struct {
	BusinessPartner *int `json:"BusinessPartner"`
	Customer        *int `json:"Customer"`
}
