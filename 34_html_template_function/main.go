package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// Account is a type
type Account struct {
	FirstName string
	LastName  string
}

// Purchase is a type
type Purchase struct {
	Date          time.Time
	Description   string
	AmountInCents int
}

// Output is a type
type Output struct {
	Message string
}

// Statement is a type
type Statement struct {
	FromDate  time.Time
	ToDate    time.Time
	Account   Account
	Purchases []Purchase
}

// The results are in CLI to demonstrate how templates work.
func main() {

	functionMap := template.FuncMap{
		"formatAsDollars": formatAsDollars,
		"formatAsDate":    formatAsDate,
		"urgentNote":      urgentNote,
		"signature":       signature,
		"list":            list,
	}

	t := template.Must(template.New("email.html").Funcs(functionMap).ParseFiles("email.html"))
	err := t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		panic(err)
	}
}

// format as dollars
func formatAsDollars(valueInCents int) (string, error) {
	dollars := valueInCents / 100
	cents := valueInCents % 100
	return fmt.Sprintf("$%d.%2d", dollars, cents), nil
}

// formatting for date
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", day, month, year)
}

// Demonstration of passing arguments back from the template.
func urgentNote(acc Account) string {
	return fmt.Sprintf(acc.FirstName + ", you have earned 100 VIP points that can be used for purchases")
}

// Dates
func createMockStatement() Statement {
	return Statement{
		FromDate: time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC),
		Account: Account{
			FirstName: "John",
			LastName:  "Dow",
		},
		Purchases: []Purchase{
			{
				Date:          time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC),
				Description:   "Shovel",
				AmountInCents: 2326,
			},
			{
				Date:          time.Date(2018, 2, 21, 0, 0, 0, 0, time.UTC),
				Description:   "Staple remover",
				AmountInCents: 5432,
			},
		},
	}
}

// Passing single string to template
func signature() string {
	return fmt.Sprintf("www.example.com")
}

// Passing a slice of structures to template as a list
func list() (r []Output) {
	//r := []Output{}
	r = append(r, Output{Message: fmt.Sprint("one")})
	r = append(r, Output{Message: fmt.Sprint("two")})
	r = append(r, Output{Message: fmt.Sprint("three")})
	r = append(r, Output{Message: fmt.Sprint("four")})

	return
}
