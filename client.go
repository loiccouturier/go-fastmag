package gofastmag

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	CreateCustomer()
	UpdateCustomer()
	UpdateCustomerEmail()
	UpdateCustomerPassword()
	CreateShippingAddress()
	UpdateLoyaltyCard()
	SellLoyaltyCard()
	RenewLoyaltyCard()
	CreateSale()
	GetPMEAmount()
	CreateReservation()
	CreateOrder()
	CreateMovement()
	GetLastBl()
	GetDetailBl()
	Transfer()
	CreateAdvancePayment()
	ConvertReservationToSale()
	Cancel()
	CancelSale()
	CancelAdvancePayment()
	CancelOrder()
	CancelReservation()
	SetOrigin()
	GetStock()
	AddInformation()
	HandleCustomerInAccount()
	GetCreditDetail()
	ExecuteQuery(data string, out interface{}) error
}

type client struct {
	host   string
	folder string
	store  string
	user   string
	pwd    string
}

func NewClient(host, folder, store, user, pwd string) Client {
	return &client{
		host:   host,
		folder: folder,
		store:  store,
		user:   user,
		pwd:    pwd,
	}
}

func (c *client) CreateCustomer() {
}

func (c *client) UpdateCustomer() {
}

func (c *client) UpdateCustomerEmail() {
}

func (c *client) UpdateCustomerPassword() {
}

func (c *client) CreateShippingAddress() {
}

func (c *client) UpdateLoyaltyCard() {
}

func (c *client) SellLoyaltyCard() {
}

func (c *client) RenewLoyaltyCard() {
}

func (c *client) CreateSale() {
}

func (c *client) GetPMEAmount() {
}

func (c *client) CreateReservation() {
}

func (c *client) CreateOrder() {
}

func (c *client) CreateMovement() {
}

func (c *client) GetLastBl() {
}

func (c *client) GetDetailBl() {
}

func (c *client) Transfer() {
}

func (c *client) CreateAdvancePayment() {
}

func (c *client) ConvertReservationToSale() {
}

func (c *client) Cancel() {
}

func (c *client) CancelSale() {
}

func (c *client) CancelAdvancePayment() {
}

func (c *client) CancelOrder() {
}

func (c *client) CancelReservation() {
}

func (c *client) SetOrigin() {
}

func (c *client) GetStock() {
}

func (c *client) AddInformation() {
}

func (c *client) HandleCustomerInAccount() {
}

func (c *client) GetCreditDetail() {
}

func (c *client) ExecuteQuery(data string, out interface{}) error {
	uri := "/EDIQUERY.IPS"

	response, err := c.call(uri, data)
	if err != nil {
		return err
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})

	err = gocsv.UnmarshalBytes(response, out)
	if err != nil {
		return err
	}

	return nil
}

// PDA
// VSHOPSTATUS
// RECEPTIONAR
// RECEPTIONTRACKINGINTERNE
// COMPLEMENT
// REMISEBQE
// QUERY

func (c *client) call(uri string, data string) ([]byte, error) {
	hc := http.Client{}

	form := url.Values{}
	form.Add("Enseigne", c.folder)
	form.Add("Magasin", c.store)
	form.Add("Compte", c.user)
	form.Add("MotPasse", c.pwd)
	form.Add("data", data)

	req, err := http.NewRequest("POST", c.host+uri, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseData, err
}
