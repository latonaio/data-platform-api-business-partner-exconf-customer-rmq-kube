package dpfm_api_output_formatter

type MetaData struct {
	ConnectionKey                                 string                                         `json:"connection_key"`
	Result                                        bool                                           `json:"result"`
	RedisKey                                      string                                         `json:"redis_key"`
	Filepath                                      string                                         `json:"filepath"`
	APIStatusCode                                 int                                            `json:"api_status_code"`
	RuntimeSessionID                              string                                         `json:"runtime_session_id"`
	BusinessPartnerID                             *int                                           `json:"business_partner"`
	ServiceLabel                                  string                                         `json:"service_label"`
	BusinessPartnerCustomer                       *BusinessPartnerCustomer                       `json:"BusinessPartnerCustomer,omitempty"`
	BusinessPartnerCustomerPartnerFunctionContact *BusinessPartnerCustomerPartnerFunctionContact `json:"BusinessPartnerCustomerPartnerFunctionContact,omitempty"`
	BusinessPartnerCustomerPartnerPlant           *BusinessPartnerCustomerPartnerPlant           `json:"BusinessPartnerCustomerPartnerPlant,omitempty"`
	BusinessPartnerCustomerFinInst                *BusinessPartnerCustomerFinInst                `json:"BusinessPartnerCustomerFinInst,omitempty"`
	APISchema                                     string                                         `json:"api_schema"`
	Accepter                                      []string                                       `json:"accepter"`
	Deleted                                       bool                                           `json:"deleted"`
}

type BusinessPartnerCustomer struct {
	BusinessPartner int  `json:"BusinessPartner"`
	Customer        int  `json:"Customer"`
	ExistenceConf   bool `json:"ExistenceConf"`
}

type BusinessPartnerCustomerPartnerFunctionContact struct {
	BusinessPartner int  `json:"BusinessPartner"`
	Customer        int  `json:"Customer"`
	PartnerCounter  int  `json:"PartnerCounter"`
	ContactID       int  `json:"ContactID"`
	ExistenceConf   bool `json:"ExistenceConf"`
}

type BusinessPartnerCustomerPartnerPlant struct {
	BusinessPartner                int    `json:"BusinessPartner"`
	Customer                       int    `json:"Customer"`
	PartnerCounter                 int    `json:"PartnerCounter"`
	PartnerFunction                string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner int    `json:"PartnerFunctionBusinessPartner"`
	PlantCounter                   int    `json:"PlantCounter"`
	ExistenceConf                  bool   `json:"ExistenceConf"`
}

type BusinessPartnerCustomerFinInst struct {
	BusinessPartner       int    `json:"BusinessPartner"`
	Customer              int    `json:"Customer"`
	FinInstIdentification int    `json:"FinInstIdentification"`
	ValidityEndDate       string `json:"ValidityEndDate"`
	ValidityStartDate     string `json:"ValidityStartDate"`
	ExistenceConf         bool   `json:"ExistenceConf"`
}
