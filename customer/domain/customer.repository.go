package domain

type CustomerRepository interface {
	GetCustomers() ([]CustomerEntity, error)
	CreateCustomer(CreateCustomerDto) (CustomerEntity, error)
}
