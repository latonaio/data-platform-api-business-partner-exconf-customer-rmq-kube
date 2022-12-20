package requests

type BusinessPartnerCustomerPartnerFunctionContact struct {
	BusinessPartner *int `json:"BusinessPartner"`
	Customer        *int `json:"Customer"`
	PartnerCounter  *int `json:"PartnerCounter"`
	ContactID       *int `json:"ContactID"`
}
