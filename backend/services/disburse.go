package service

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Service -
type Service struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}


// Disburse - 
func (s *Service)Disburse() {

	// 1. Fetch all accounts 

	// 2. format payment payloads

	// 3. send money

	// 4. update records
}