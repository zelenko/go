package main

// https://godoc.org/html/template
// https://golang.org/pkg/html/template/
import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// Account has first and last name
type Account struct {
	FirstName string
	LastName  string
}

// Purchase is the purchase order placed by specific Accont
type Purchase struct {
	Date          time.Time
	Description   string
	AmountInCents int
}

// Statement can have many Purchase Orders
type Statement struct {
	FromDate  time.Time
	ToDate    time.Time
	Account   Account
	Purchases []Purchase
}

func main() {
	fmap := template.FuncMap{
		"formatAsDollars": formatAsDollars,
		"formatAsDate":    formatAsDate,
		"urgentNote":      urgentNote,
	}
	t := template.Must(template.New("htmlTemplate").Funcs(fmap).Parse(htmlTemplate))
	err := t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		panic(err)
	}
}

func formatAsDollars(valueInCents int) (string, error) {
	dollars := valueInCents / 100
	cents := valueInCents % 100
	return fmt.Sprintf("$%d.%2d", dollars, cents), nil
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%d/%d", day, month, year)
}

func urgentNote(acc Account) string {
	return fmt.Sprintf("You have earned 100 VIP points that can be used for purchases")
}

func createMockStatement() Statement {
	return Statement{
		FromDate: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
		Account: Account{
			FirstName: "John",
			LastName:  "Dow",
		},
		Purchases: []Purchase{
			Purchase{
				Date:          time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
				Description:   "Shovel",
				AmountInCents: 2326,
			},
			Purchase{
				Date:          time.Date(2016, 1, 8, 0, 0, 0, 0, time.UTC),
				Description:   "Staple remover",
				AmountInCents: 5432,
			},
		},
	}
}

var htmlTemplate = `
{{with .Account -}}
Dear {{.FirstName}} {{.LastName}},
{{- end}}

Below are your account statement details for period from {{.FromDate | formatAsDate}} to {{.ToDate | formatAsDate}}.

{{if .Purchases -}}
    Your purchases:
    {{- range .Purchases }}
        {{ .Date | formatAsDate}} {{ printf "%-20s" .Description }} {{.AmountInCents | formatAsDollars -}}
    {{- end}}
{{- else}}
You didn't make any purchases during the period.
{{- end}}

{{$note := urgentNote .Account -}}
{{if $note -}}
Note: {{$note}}
{{- end}}

Best Wishes,
Customer Service
`
