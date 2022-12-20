package requests

type BusinessPartnerCustomerPartnerPlant struct {
	BusinessPartner                *int    `json:"BusinessPartner"`
	Customer                       *int    `json:"Customer"`
	PartnerCounter                 *int    `json:"PartnerCounter"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	PlantCounter                   *int    `json:"PlantCounter"`
}
