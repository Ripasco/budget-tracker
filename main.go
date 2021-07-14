package main

import m "budget/models"

func main() {
	// var User m.Transaction
	// var Budget m.Transaction
	user := m.NewUserCreate("Aslan", "Rancho", "test@mail.com")
	budget := m.NewBudget("Famaly Budget", 13400.0, user)
	user.Pay(200.0, "Payed internet", budget.Id)
	budget.Info()
}
