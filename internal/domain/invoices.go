package domain

type Invoices struct {
	Id         int     `csv:"id"`
	Datetime   string  `csv:"datetime"`
	IdCustomer string  `csv:"id_customer"`
	Total      float64 `csv:"total"`
}
