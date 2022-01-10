package domain

type Products struct {
	Id          int     `csv:"id"`
	Description string  `csv:"description"`
	Price       float32 `csv:"price"`
}
