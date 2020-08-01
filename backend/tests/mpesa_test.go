package tests

import (
	"elephant/payments/gateways/mpesa"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pesa = new(mpesa.Mpesa)

func TestNewFunc(t *testing.T) {

	err := pesa.New()

	// payments.New()

	assert.NoError(t, err)
}

func TestPayOne(t *testing.T) {

	_, err := pesa.PayOne("254708374149", "100", "001")

	assert.NoError(t, err)
}
