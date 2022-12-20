package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-exconf-customer-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-exconf-customer-rmq-kube/DPFM_API_Output_Formatter"
	"encoding/json"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewExistenceConf(ctx context.Context, db *database.Mysql, l *logger.Logger) *ExistenceConf {
	return &ExistenceConf{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (e *ExistenceConf) Conf(msg rabbitmq.RabbitmqMessage) interface{} {
	var ret interface{}
	ret = map[string]interface{}{
		"ExistenceConf": false,
	}
	input := make(map[string]interface{})
	err := json.Unmarshal(msg.Raw(), &input)
	if err != nil {
		return ret
	}

	_, ok := input["BusinessPartnerCustomer"]
	if ok {
		input := &dpfm_api_input_reader.CustomerSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.confBusinessPartnerCustomer(input)
		goto endProcess
	}
	_, ok = input["BusinessPartnerCustomerPartnerFunctionContact"]
	if ok {
		input := &dpfm_api_input_reader.PartnerFunctionContactSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfBusinessPartnerCustomerPartnerFunctionContact(input)
		goto endProcess
	}
	_, ok = input["BusinessPartnerCustomerPartnerPlant"]
	if ok {
		input := &dpfm_api_input_reader.PartnerPlantSDC{}
		err = json.Unmarshal(msg.Raw(), input)
		ret = e.ConfBusinessPartnerCustomerPartnerPlant(input)
		goto endProcess
	}

	err = xerrors.Errorf("can not get exconf check target")
endProcess:
	if err != nil {
		e.l.Error(err)
	}
	return ret
}

func (e *ExistenceConf) confBusinessPartnerCustomer(input *dpfm_api_input_reader.CustomerSDC) *dpfm_api_output_formatter.BusinessPartnerCustomer {
	exconf := dpfm_api_output_formatter.BusinessPartnerCustomer{
		ExistenceConf: false,
	}
	if input.BusinessPartnerCustomer.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomer.Customer == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerCustomer{
		BusinessPartner: *input.BusinessPartnerCustomer.BusinessPartner,
		Customer:        *input.BusinessPartnerCustomer.Customer,
		ExistenceConf:   false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerCustomer
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data 
		WHERE (businessPartner, customer) = (?, ?);`, exconf.BusinessPartner, exconf.Customer,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfBusinessPartnerCustomerPartnerFunctionContact(input *dpfm_api_input_reader.PartnerFunctionContactSDC) *dpfm_api_output_formatter.BusinessPartnerCustomerPartnerFunctionContact {
	exconf := dpfm_api_output_formatter.BusinessPartnerCustomerPartnerFunctionContact{
		ExistenceConf: false,
	}
	if input.BusinessPartnerCustomerPartnerFunctionContact.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerFunctionContact.Customer == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerFunctionContact.PartnerCounter == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerFunctionContact.ContactID == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerCustomerPartnerFunctionContact{
		BusinessPartner: *input.BusinessPartnerCustomerPartnerFunctionContact.BusinessPartner,
		Customer:        *input.BusinessPartnerCustomerPartnerFunctionContact.Customer,
		PartnerCounter:  *input.BusinessPartnerCustomerPartnerFunctionContact.PartnerCounter,
		ContactID:       *input.BusinessPartnerCustomerPartnerFunctionContact.ContactID,
		ExistenceConf:   false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerCustomerPartnerFunctionContact
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_function_contact_data 
		WHERE (businessPartner, customer, partnerCounter, contactID) = (?, ?, ?, ?);`, exconf.BusinessPartner, exconf.Customer, exconf.PartnerCounter, exconf.ContactID,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfBusinessPartnerCustomerPartnerPlant(input *dpfm_api_input_reader.PartnerPlantSDC) *dpfm_api_output_formatter.BusinessPartnerCustomerPartnerPlant {
	exconf := dpfm_api_output_formatter.BusinessPartnerCustomerPartnerPlant{
		ExistenceConf: false,
	}
	if input.BusinessPartnerCustomerPartnerPlant.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerPlant.Customer == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerPlant.PartnerCounter == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerPlant.PartnerFunction == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerPlant.PartnerFunctionBusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerPartnerPlant.PlantCounter == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerCustomerPartnerPlant{
		BusinessPartner:                *input.BusinessPartnerCustomerPartnerPlant.BusinessPartner,
		Customer:                       *input.BusinessPartnerCustomerPartnerPlant.Customer,
		PartnerCounter:                 *input.BusinessPartnerCustomerPartnerPlant.PartnerCounter,
		PartnerFunction:                *input.BusinessPartnerCustomerPartnerPlant.PartnerFunction,
		PartnerFunctionBusinessPartner: *input.BusinessPartnerCustomerPartnerPlant.PartnerFunctionBusinessPartner,
		PlantCounter:                   *input.BusinessPartnerCustomerPartnerPlant.PlantCounter,
		ExistenceConf:                  false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerCustomerPartnerPlant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_partner_plant_data 
		WHERE (businessPartner, customer, partnerCounter, partnerFunction, partnerFunctionBusinessPartner, plantCounter) = (?, ?, ?, ?, ?, ?);`, exconf.BusinessPartner, exconf.Customer, exconf.PartnerCounter, exconf.PartnerFunction, exconf.PartnerFunctionBusinessPartner, exconf.PlantCounter,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}

func (e *ExistenceConf) ConfBusinessPartnerCustomerFinInst(input *dpfm_api_input_reader.FinInstSDC) *dpfm_api_output_formatter.BusinessPartnerCustomerFinInst {
	exconf := dpfm_api_output_formatter.BusinessPartnerCustomerFinInst{
		ExistenceConf: false,
	}
	if input.BusinessPartnerCustomerFinInst.BusinessPartner == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerFinInst.Customer == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerFinInst.FinInstIdentification == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerFinInst.ValidityEndDate == nil {
		return &exconf
	}
	if input.BusinessPartnerCustomerFinInst.ValidityStartDate == nil {
		return &exconf
	}
	exconf = dpfm_api_output_formatter.BusinessPartnerCustomerFinInst{
		BusinessPartner:       *input.BusinessPartnerCustomerFinInst.BusinessPartner,
		Customer:              *input.BusinessPartnerCustomerFinInst.Customer,
		FinInstIdentification: *input.BusinessPartnerCustomerFinInst.FinInstIdentification,
		ValidityEndDate:       *input.BusinessPartnerCustomerFinInst.ValidityEndDate,
		ValidityStartDate:     *input.BusinessPartnerCustomerFinInst.ValidityStartDate,
		ExistenceConf:         false,
	}

	rows, err := e.db.Query(
		`SELECT BusinessPartnerCustomerFinInst
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_fin_inst_data 
		WHERE (businessPartner, customer, finInstIdentification, validityEndDate, validityStartDate) = (?, ?, ?, ?, ?);`, exconf.BusinessPartner, exconf.Customer, exconf.FinInstIdentification, exconf.ValidityEndDate, exconf.ValidityStartDate,
	)
	if err != nil {
		e.l.Error(err)
		return &exconf
	}

	exconf.ExistenceConf = rows.Next()
	return &exconf
}
