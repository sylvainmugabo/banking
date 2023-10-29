package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func (d CustomerRepositoryStub) Find(id string) (*Customer, error) {

	for _, row := range d.customers {
		if row.Id == id {
			return &row, nil
		}

	}
	return nil, nil
}

func NewCustomerSub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "123", Name: "Mugabo", DateofBirth: "10/10/2011", City: "New Market", ZipCode: "21774"},
		{Id: "113", Name: "Liam", DateofBirth: "10/10/2011", City: "New Market", ZipCode: "21774"},
	}

	return CustomerRepositoryStub{customers: customers}
}
