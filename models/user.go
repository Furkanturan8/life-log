package models

type User struct {
	ID      int    `bun:",pk,autoincrement"`
	Name    string `bun:"name" json:"name"`
	Surname string `bun:"surname" json:"surname"`
	Email   string `bun:"email" json:"email"`
}

func (User) TableName() string {
	return "users"
}
