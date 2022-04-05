package entity

// representasi table di database
type DetailsAddress struct {
	City     string
	Province string
	Country  string
}

type Costumer struct {
	Id       string
	Name     string
	Username string
	Address  DetailsAddress
}
