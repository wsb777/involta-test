package models

type Document struct {
	ID       string `reindex:"id"`
	Name     string `reindex:"name"`
	CreateAt string `reindex:"createAt"`
	UpdateAt string `reindex:"updateAt"`
}

type Person struct {
	ID         string     `reindex:"id,,pk"`
	FirstName  string     `reindex:"firstName"`
	SecondName string     `reindex:"secondName"`
	MiddleName string     `reindex:"middleName"`
	CreateAt   string     `reindex:"createAt"`
	UpdateAt   string     `reindex:"updateAt"`
	Documents  []Document `reindex:"documents,json"`
}
