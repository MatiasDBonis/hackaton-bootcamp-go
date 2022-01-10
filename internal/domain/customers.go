package domain

type Customers struct {
	Id        int    `csv:"id"`
	LastName  string `csv:"last_name"`
	FirstName string `csv:"first_name"`
	Condition string `csv:"condition"`
}
