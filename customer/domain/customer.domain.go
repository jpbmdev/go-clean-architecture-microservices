package domain

type CustomerEntity struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
}

type CreateCustomerDto struct {
	Name string `json:"name"`
}
