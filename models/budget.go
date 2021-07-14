package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type Transaction interface {
	Pay(float64, string, uuid.UUID) error
	Info()
}

type Budget struct {
	Id          uuid.UUID
	Name        string
	Balance     float64
	Founders    []uuid.UUID
	Author      uuid.UUID
	Childs      []Budget
	CreatedTime time.Time
	UpdateTime  time.Time
	Action      []Turnover
}

type Turnover struct {
	Time   time.Time
	Author uuid.UUID
	Sum    float64
	Name   string
}

func NewBudget(name string, balance float64, user *User) *Budget {
	defer fmt.Println("Budget Created by: ", user.Name)
	newID, _ := uuid.New()
	b := &Budget{
		Id:          newID,
		Name:        name,
		Balance:     balance,
		Founders:    []uuid.UUID{},
		Author:      user.Id,
		Childs:      []Budget{},
		CreatedTime: time.Now(),
		Action:      []Turnover{},
	}

	user.Budgets = append(user.Budgets, b)
	return b
}

func (b Budget) Info() {
	fmt.Printf("\nBudgetInfo:\nId: %d\nName: %s\nBalance: %f\nFounders: %d\nAuthor: %d\nCreatedTime: %s\n", b.Id, b.Name, b.Balance, b.Founders, b.Author, b.CreatedTime)
}
func (b *Budget) Pay(sum float64, name string, id uuid.UUID) error {
	b.Balance -= sum
	b.UpdateTime = time.Now()
	b.Action = append(b.Action, Turnover{
		Time:   time.Now(),
		Author: id,
		Sum:    sum,
		Name:   name,
	})
	return nil
}
