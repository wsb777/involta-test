package dto

type Document struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PersonCreate struct {
	FirstName  string     `json:"firstName"`
	SecondName string     `json:"secondName"`
	MiddleName string     `json:"middleName"`
	Documents  []Document `json:"documents"`
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

type PersonDelete struct {
	ID int `json:"id"`
}
