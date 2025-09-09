package models

type Document struct {
	ID       int    `reindex:"id"`
	Name     string `reindex:"name"`
	CreateAt string `reindex:"createAt"`
	UpdateAt string `reindex:"updateAt"`
}

type Person struct {
	ID         int        `reindex:"id,,pk"`
	FirstName  string     `reindex:"firstName" json:"firstName,omitempty"`
	SecondName string     `reindex:"secondName" json:"secondName,omitempty"`
	MiddleName string     `reindex:"middleName" json:"middleName,omitempty"`
	CreateAt   string     `reindex:"createAt" json:"createAt,omitempty"`
	UpdateAt   string     `reindex:"updateAt" json:"updateAt,omitempty"`
	Documents  []Document `reindex:"documents,json"`
}
