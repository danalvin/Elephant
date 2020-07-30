package services

import (
	"elephant/data"
	"elephant/log"
	"elephant/models"
	"elephant/payments"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/uniplaces/carbon"
)

// Service -
type Service struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

// NewService -
func NewService() *Service {
	return &Service{
		Logger: log.GetLogger(),
		DB:     data.GetDB(),
	}
}

// Disburse -
// TODO: - use context to propagate cancellation sigs
func (s *Service) Disburse() {

	// 1. Init payment gateway
	payway, err := payments.New()

	if err != nil {
		s.Logger.Errorf("cannot init payment gateway : %v", err)
		return
	}

	// 2. Fetch all accounts
	var employees []models.Employee

	gor := s.DB.Raw(`SELECT * FROM employees 
				WHERE 
					deleted_at IS NULL AND 
					active = true`,
	).Scan(&employees)

	if err := gor.Error; err != nil {
		s.Logger.Errorf("could not fetch employeed records : %v", err)
	}

	// 3. save payments
	if len(employees) > 0 {

		if err := s.savePayment(employees); err != nil {
			s.Logger.Errorf("cannot save payment records to payments table : %v", err)
			return
		}
	}

	// 4. prepare payload and disburse payments
	var payload []*payments.Payload

	for _, k := range employees {
		payload = append(payload, &payments.Payload{AccountNumber: k.PhoneNumber, Amount: fmt.Sprint(k.Wage)})
	}

	res, err := payway.PayBulk(payload)

	if err != nil {
		s.Logger.Errorf("cannot disburse payments : %v", err)
	}

	// 5. update table records
	if err := s.updatePayment(res); err != nil {
		s.Logger.Errorf("cannot update payment records : %v", err)
	}
}

// savePayment - before actual payment is done, save payment record
func (s *Service) savePayment(emp []models.Employee) error {

	var (
		err error
		
		// begin transaction - 
		tx  = s.DB.Begin()
	)

	// initial query
	query := fmt.Sprintf(`INSERT INTO payments (amount, employee_id, date_paid, confirmed)
			VALUES `)

	var values []string

	// wait for it -
	for _, val := range emp {

		values = append(values, fmt.Sprintf(`(%v, %v, %v, %v),`, val.Wage, val.ID, carbon.Now(), false))
	}

	// save the records - limit could be upto 1k records at least mysql can support that
	for _, chunk := range splitChunks(values, 20) {

		// expand the values now
		var v = ``

		for _, c := range chunk {
			v += c
		}

		q := query + v

		if err = tx.Exec(q).Error; err != nil {
			tx.Rollback()
			return err
		}

	}

	// commit transaction  and return
	return tx.Commit().Error
}

// splitChunks - returns chunks
func splitChunks(load []string, limit int) [][]string {

	var chunks = [][]string{}
	var start, stop = 0, limit

	for i := 0; i <= len(load)/limit; i++ {

		var batch []string

		if len(load) < limit {
			batch = load
		} else {
			batch = load[start:stop]
		}

		chunks = append(chunks, batch)
		batch = nil

		if stop+limit >= len(load) {
			start, stop = stop, len(load)
		} else {
			start, stop = stop, stop+limit
		}

	}

	return chunks
}

func (s *Service) updatePayment([]*payments.PaywayResponse) error {

	
	return nil
}
