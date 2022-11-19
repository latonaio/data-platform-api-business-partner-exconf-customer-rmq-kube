package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-exconf-customer-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBusinessPartnerCustomer() *requests.BusinessPartnerCustomer {
	data := sdc.BusinessPartnerCustomer
	return &requests.BusinessPartnerCustomer{
		BusinessPartner: data.BusinessPartner,
		Customer:        data.Customer,
	}
}
