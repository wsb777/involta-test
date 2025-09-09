package dto

type PersonCreate struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	MiddleName string `json:"middleName"`
}

type PersonID struct {
	ID int `json:"id"`
}

type PersonUpdate struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	MiddleName string `json:"middleName"`
}

type PersonGet struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	MiddleName string `json:"middleName"`
}
