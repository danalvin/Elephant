package mpesa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"elephant/config"
	"elephant/payments"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/spf13/viper"
)

var (
	conf    = config.GetConfig()
	baseURL = "https://sandbox.safaricom.co.ke"
	env     = conf.GetString("app.environment")
)

const (
	b2cAPIEndpoint = "%v/mpesa/b2c/v1/paymentrequest"
	accessTokenURL = "%v/oauth/v1/generate?grant_type=client_credentials"
)

// Mpesa - B2C service
type Mpesa struct {
	conf *viper.Viper
}

// B2C -
type B2C struct {
	InitiatorName      string `json:"InitiatorName"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	Amount             string `json:"Amount"`
	PartyA             string `json:"PartyA"`
	PartyB             string `json:"PartyB"`
	Remarks            string `json:"Remarks"`
	QueueTimeoutURL    string `json:"QueueTimeoutURL"`
	ResultsURL         string `json:"ResultsURL"`
	Occasion           string `json:"Occasion"` // optional field
}

// init registration
func init() {

	if env == "production" {

		baseURL = conf.GetString("gateways.payments.mpesa.production_url")

		if strings.TrimSpace(baseURL) == "" {
			panic("mpesa: production url not provided")
		}
	}

	payments.Register("mpesa", new(Mpesa))
}

// New - mpesa instance
func (m *Mpesa) New() error {

	return nil
}

// PayOne -
func (m *Mpesa) PayOne(account payments.AccountNumber, amount string, ref string) (*payments.PaywayResponse, error) {

	var payer = "testapi115"

	// prepare payload
	payload := &B2C{
		InitiatorName:      payer,
		SecurityCredential: securityCredential("Safaricom007@"),
		CommandID:          "SalaryPayment",
		Amount:             amount,
		PartyA:             "174379",
		PartyB:             string(account),
		Remarks:            "test transaction",
		QueueTimeoutURL:    "https://test.com",
		ResultsURL:         "https://test.com",
		// Occasion:           "string",
	}

	// fmt.Printf("nothign ---------- %v\n", securityCredential())

	data, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(b2cAPIEndpoint, baseURL), bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+getAccessToken())

	if err != nil {
		return nil, err
	}

	client := new(http.Client)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	jdata, err := ioutil.ReadAll(res.Body)

	// print the data -
	fmt.Printf("data : %s\n", jdata)

	// prepare response
	resData := &payments.PaywayResponse{}

	return resData, nil
}

// PayBulk -
func (m *Mpesa) PayBulk([]*payments.Payload) ([]*payments.PaywayResponse, error) {

	//
	return nil, nil
}

// PayMany - pay many wallets same amount
func (m *Mpesa) PayMany(accounts []payments.AccountNumber, amount string) (*payments.PaywayResponse, error) {
	return nil, nil
}

// Balance - money in the account
func (m *Mpesa) Balance() (*payments.PaywayResponse, error) {
	return nil, nil
}

// setAccessToken - does what needs to be done
func getAccessToken() string {

	// 1. request access token
	req, err := http.NewRequest("GET", fmt.Sprintf(accessTokenURL, baseURL), nil)
	req.Header.Set("Authorization", "Basic "+getPassword())
	res, err := new(http.Client).Do(req)
	if err != nil {
		log.Fatalf("mpesa: cannot make request for acquiring accesss token : %v", err)
		return ""
	}
	// 2. read response
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("mpesa: cannot parse response from access token req : %v", err)
		return ""
	}
	aT, _ := jsonparser.GetString(data, "access_token")
	return aT
}

func getPassword() string {
	secret := conf.GetString("gateways.mpesa.consumer_secret")
	key := conf.GetString("gateways.mpesa.consumer_key")
	return base64.StdEncoding.EncodeToString([]byte(key + ":" + secret))
}

// securityCred -
func securityCredential(initName string) string {

	// 1. download cert
	certPath := "https://developer.safaricom.co.ke/sites/default/files/cert/cert_sandbox/cert.cer"
	if env == "production" {
		certPath = "https://developer.safaricom.co.ke/sites/default/files/cert/cert_prod/cert.cer"
	}
	res, err := http.Get(certPath)
	if err != nil {
		log.Fatalf("cannot fetch cert : %v", err)
	}

	data, err := ioutil.ReadAll(res.Body)

	// 2. do encryption
	pem, _ := pem.Decode(data)
	cert, err := x509.ParseCertificate(pem.Bytes)
	key := cert.PublicKey.(*rsa.PublicKey)
	ciph, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(initName))
	if err != nil {
		log.Fatalf("cannot encrypt message : %v", err)
	}

	return base64.StdEncoding.EncodeToString(ciph)
}
