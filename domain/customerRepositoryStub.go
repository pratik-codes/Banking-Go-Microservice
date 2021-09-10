package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Pratik", City: "San Francisco", Zipcode: "3123123", DateofBirth: "12/12/1998", Status: "1"},
		{Id: "2", Name: "Neil", City: "Pune", Zipcode: "3223123", DateofBirth: "12/12/1998", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}