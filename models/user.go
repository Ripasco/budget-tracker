package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Password string
	Email    string
	Budgets  []*Budget
}

func NewUserCreate(name string, password string, Email string) *User {
	newId, _ := uuid.New()
	var user = &User{
		Id:       newId,
		Name:     name,
		Password: password,
		Email:    Email,
		Budgets:  []*Budget{},
	}
	return user
}

func (u User) Info() {
	fmt.Println(u)
}
func (u User) Pay(sum float64, name string, id uuid.UUID) error {
	for _, val := range u.Budgets {
		if val.Id == id {
			err := val.Pay(sum, name, u.Id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Pay success!!!")
			fmt.Printf("\n\nName: %s \n\nSum: %f \n\n", name, sum)
		}
	}
	return nil
}
