package requests

type BusinessPartnerCustomerFinInst struct {
	BusinessPartner       *int    `json:"BusinessPartner"`
	Customer              *int    `json:"Customer"`
	FinInstIdentification *int    `json:"FinInstIdentification"`
	ValidityEndDate       *string `json:"ValidityEndDate"`
	ValidityStartDate     *string `json:"ValidityStartDate"`
}
