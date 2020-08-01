package payments

import (
	"elephant/config"
	"elephant/log"
	"fmt"
	"sync"
)

var (
	conf     = config.GetConfig()
	l        = log.GetLogger()
	mu       sync.RWMutex
	gateways = make(map[string]Gateway) // cache
)

// Payway - payment gateway
type Payway struct {
	PayGateway Gateway
}

// Gateway - all payment gateways should implement these methods
type Gateway interface {

	// Initialize new gateway
	New() error

	// PayOne - pay to one wallet
	PayOne(account AccountNumber, amount string, ref string) (*PaywayResponse, error)

	// PayBulk -
	PayBulk([]*Payload) ([]*PaywayResponse, error)

	// PayMany - pay many wallets same amount
	PayMany(accounts []AccountNumber, amount string) (*PaywayResponse, error)

	// Balance - money in the account
	Balance() (*PaywayResponse, error)
}

// AccountNumber - could be phone_number, bank acc no or email or bitcoin address
type AccountNumber string

// Payload -
type Payload struct {
	AccountNumber string `json:"account_number"`
	Amount        string `json:"amount"`
	PaymentID     string `json:"payment_id"` // id of the transaction
}

// PaywayResponse -
type PaywayResponse struct {
	TransactionID string        `json:"transaction_id"`
	AccountID     AccountNumber `json:"account_id"`
	HTTPStatus    string        `json:"http_status"`
	Status        string        `json:"status"`
	IsSuccessful  bool          `json:"is_successful"`
	StatusCode    string        `json:"status_code"`
	AccountBal    string        `json:"account_balance"`
}

// Register -
func Register(name string, pway Gateway) {

	// acquire mutex lock
	mu.Lock()

	defer mu.Unlock()

	if pway == nil {
		l.Fatalf("cannot register payment gateway %v ", name)
	}

	if _, ok := gateways[name]; ok {
		l.Fatalf("Payment Gateway already registered : %v", name)
	}

	// register the gateway
	gateways[name] = pway

}

// New - initialize new Payment Gateway
func New() (*Payway, error) {

	// get default gateway from configuration file
	provider := conf.GetString("gateways.payments.default")

	if provider == "" {
		return nil, fmt.Errorf("No default payment gateway provided")
	}

	// acquare read mutex
	mu.RLock()

	defer mu.RUnlock()

	// fetch gateway from cache
	gateway, ok := gateways[provider]

	if !ok {

		return nil, fmt.Errorf("Gateway not registered")
	}

	// Init gateway
	if err := gateway.New(); err != nil {

		return nil, err
	}

	return &Payway{PayGateway: gateway}, nil
}

// Name -
func (p *Payway) Name() string {

	return p.Name()
}

// PayOne - credits money to one account number
func (p *Payway) PayOne(account string, amount float64) (*PaywayResponse, error) {

	return p.PayOne(account, amount)
}

// PayMany -
func (p *Payway) PayMany(accounts []AccountNumber, amount float64) (*PaywayResponse, error) {

	return p.PayMany(accounts, amount)
}

// PayBulk -
func (p *Payway) PayBulk(load []*Payload) ([]*PaywayResponse, error) {

	return p.PayBulk(load)
}

// Balance -
func (p *Payway) Balance() (*PaywayResponse, error) {

	return p.Balance()
}
