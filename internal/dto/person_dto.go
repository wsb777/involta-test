package dto

type DocumentCreate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CreateAt string `json:"createAt"`
}

type DocumentGet struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CreateAt string `json:"createAt"`
}

type PersonCreate struct {
	FirstName  string           `json:"firstName"`
	SecondName string           `json:"secondName"`
	MiddleName string           `json:"middleName"`
	Sort       string           `json:"sort"`
	Documents  []DocumentCreate `json:"documents"`
}

type PersonUpdate struct {
	ID         int              `json:"id"`
	FirstName  string           `json:"firstName"`
	SecondName string           `json:"secondName"`
	MiddleName string           `json:"middleName"`
	Sort       string           `json:"sort"`
	Documents  []DocumentCreate `json:"documents"`
}

type PersonGet struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	MiddleName string `json:"middleName"`
	Sort       int
	CreateAt   string        `json:"createAt"`
	UpdateAt   string        `json:"updateAt"`
	Documents  []DocumentGet `json:"documents"`
}

type PersonDelete struct {
	ID int `json:"id"`
}

type SearchParams struct {
	Limit  int
	LastID int
	Text   string
}
