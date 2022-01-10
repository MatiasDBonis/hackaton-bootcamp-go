package domain

type Invoices struct {
	Id         int     `csv:"id"`
	Datetime   string  `csv:"datetime"`
	IdCustomer int     `csv:"id_customer"`
	Total      float64 `csv:"total"`
}
