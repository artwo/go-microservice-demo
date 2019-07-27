package model

import "errors"

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type PersonNoID struct {
	FirstName string   `json:"firstName" binding:"required"`
	LastName  string   `json:"lastName" binding:"required"`
	Address   *Address `json:"address" binding:"required"`
}

type Address struct {
	City  string `json:"city,omitempty" binding:"required"`
	State string `json:"state,omitempty" binding:"required"`
}

func (p *Person) Validate() []error {
	var errs []error
	if p.FirstName == "" {
		errs = append(errs, errors.New("firstName field of Person is empty"))
	}
	if p.LastName == "" {
		errs = append(errs, errors.New("lastName field of Person is empty"))
	}

	errs = append(errs, p.Address.Validate()...)
	return errs
}

func (p *PersonNoID) Validate() []error {
	var errs []error
	if p.FirstName == "" {
		errs = append(errs, errors.New("firstName field of Person is empty"))
	}
	if p.LastName == "" {
		errs = append(errs, errors.New("lastName field of Person is empty"))
	}

	errs = append(errs, p.Address.Validate()...)
	return errs
}

func (a *Address) Validate() []error {
	var errs []error
	if (a == nil || a == &Address{}) {
		errs = append(errs, errors.New("address object of Person missing or empty"))
		return errs
	}
	if a.City == "" {
		errs = append(errs, errors.New("city field of Address is empty"))
	}
	if a.State == "" {
		errs = append(errs, errors.New("state field of Address is empty"))
	}
	return errs
}
