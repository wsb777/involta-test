package models

type Document struct {
	ID       int    `reindex:"id"`
	Name     string `reindex:"name"`
	CreateAt string `reindex:"createAt"`
	UpdateAt string `reindex:"updateAt"`
}

type Person struct {
	ID         int        `reindex:"id,,pk"`
	FirstName  string     `reindex:"firstName,text" json:"firstName,omitempty"`
	SecondName string     `reindex:"secondName,text" json:"secondName,omitempty"`
	MiddleName string     `reindex:"middleName,text" json:"middleName,omitempty"`
	CreateAt   string     `reindex:"createAt" json:"createAt,omitempty"`
	UpdateAt   string     `reindex:"updateAt" json:"updateAt,omitempty"`
	Documents  []Document `reindex:"documents,json"`
	_          struct{}   `reindex:"firstName+secondName+middleName=fullName,text,composite"`
}

type SearchParams struct {
	Limit  int
	LastID int
	Text   string
}
