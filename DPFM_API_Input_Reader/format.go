package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-exconf-customer-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *CustomerSDC) ConvertToBusinessPartnerCustomer() *requests.BusinessPartnerCustomer {
	data := sdc.BusinessPartnerCustomer
	return &requests.BusinessPartnerCustomer{
		BusinessPartner: data.BusinessPartner,
		Customer:        data.Customer,
	}
}

func (sdc *PartnerFunctionContactSDC) ConvertToBusinessPartnerCustomerPartnerFunctionContact() *requests.BusinessPartnerCustomerPartnerFunctionContact {
	data := sdc.BusinessPartnerCustomerPartnerFunctionContact
	return &requests.BusinessPartnerCustomerPartnerFunctionContact{
		BusinessPartner: data.BusinessPartner,
		Customer:        data.Customer,
		PartnerCounter:  data.PartnerCounter,
		ContactID:       data.ContactID,
	}
}

func (sdc *PartnerPlantSDC) ConvertToBusinessPartnerCustomerPartnerPlant() *requests.BusinessPartnerCustomerPartnerPlant {
	data := sdc.BusinessPartnerCustomerPartnerPlant
	return &requests.BusinessPartnerCustomerPartnerPlant{
		BusinessPartner:                data.BusinessPartner,
		Customer:                       data.Customer,
		PartnerCounter:                 data.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		PlantCounter:                   data.PlantCounter,
	}
}

func (sdc *FinInstSDC) ConvertToBusinessPartnerCustomerFinInst() *requests.BusinessPartnerCustomerFinInst {
	data := sdc.BusinessPartnerCustomerFinInst
	return &requests.BusinessPartnerCustomerFinInst{
		BusinessPartner:       data.BusinessPartner,
		Customer:              data.Customer,
		FinInstIdentification: data.FinInstIdentification,
		ValidityEndDate:       data.ValidityEndDate,
		ValidityStartDate:     data.ValidityStartDate,
	}
}
