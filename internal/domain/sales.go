package domain

type Sales struct {
	Id        int     `csv:"id"`
	IdInvoice int     `csv:"id_invoice"`
	IdProduct int     `csv:"id_product"`
	Quantity  float32 `csv:"quantity"`
}
