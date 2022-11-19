package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-exconf-customer-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-exconf-customer-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-business-partner-exconf-customer-rmq-kube/database"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
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

func (e *ExistenceConf) Conf(input *dpfm_api_input_reader.SDC) *dpfm_api_output_formatter.BusinessPartnerCustomer {
	businessPartner := *input.BusinessPartnerCustomer.BusinessPartner
	customer := *input.BusinessPartnerCustomer.Customer
	notKeyExistence := make([]dpfm_api_output_formatter.BusinessPartnerCustomer, 0, 1)
	KeyExistence := make([]dpfm_api_output_formatter.BusinessPartnerCustomer, 0, 1)

	existData := &dpfm_api_output_formatter.BusinessPartnerCustomer{
		BusinessPartner: businessPartner,
		Customer:        customer,
		ExistenceConf:   false,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if !e.confBusinessPartnerCustomer(businessPartner, customer) {
			notKeyExistence = append(
				notKeyExistence,
				dpfm_api_output_formatter.BusinessPartnerCustomer{businessPartner, customer, false},
			)
			return
		}
		KeyExistence = append(KeyExistence, dpfm_api_output_formatter.BusinessPartnerCustomer{businessPartner, customer, true})
	}()

	wg.Wait()

	if len(KeyExistence) == 0 {
		return existData
	}
	if len(notKeyExistence) > 0 {
		return existData
	}

	existData.ExistenceConf = true
	return existData
}

func (e *ExistenceConf) confBusinessPartnerCustomer(businessPartner int, customer int) bool {
	rows, err := e.db.Query(
		`SELECT Customer 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_customer_data 
		WHERE (BusinessPartner, Customer) = (?, ?);`, businessPartner, customer,
	)
	if err != nil {
		e.l.Error(err)
		return false
	}
	if err != nil {
		e.l.Error(err)
		return false
	}

	for rows.Next() {
		var businessPartner int
		var customer int
		err := rows.Scan(&customer)
		if err != nil {
			e.l.Error(err)
			continue
		}
		if businessPartner == businessPartner {
			return true
		}
	}
	return false
}
