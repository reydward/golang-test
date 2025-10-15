package main

import (
	"fmt"
	gomail "gopkg.in/gomail.v2"
	"strings"
)

type NumberOfTransactions struct {
	Month            string `db:"month"`
	TransactionCount int    `db:"transaction_count"`
}

func main() {
	abc := gomail.NewMessage()

	userName := "Eduard Reyes"
	totalBalance := fmt.Sprintf("%f", 1000.20)
	averageDebitAmount := fmt.Sprintf("%f", -15.38)
	averageCreditAmount := fmt.Sprintf("%f", 35.25)
	numberOfTransactions := []NumberOfTransactions{
		{
			Month:            "July",
			TransactionCount: 2,
		},
		{
			Month:            "August",
			TransactionCount: 3,
		},
	}

	var numberOfTransactionsHTML strings.Builder
	for _, transaction := range numberOfTransactions {
		numberOfTransactionsHTML.WriteString(fmt.Sprintf("<p>%s: %d transactions</p>\n", transaction.Month, transaction.TransactionCount))
	}

	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<style>
				.container { width: 800px; margin: 0 auto; font-family: AppleGothic, sans-serif; font-size: small}
				.header { text-align: center; margin-bottom: 20px; }
				.summary { text-align: center; margin-bottom: 20px; }
				table { width: 100%; border-collapse: collapse; }
				th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
				th { background-color: #f2f2f2; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<img style="display:block" height="35" src="https://ci3.googleusercontent.com/meips/ADKq_NY01inpn_TMEeLOn8dwDeAXcXIMXc6P7Rh9vFT49_YXSyC1w4XmLu1fxxhQhXP4iErjgabTGxrIgw2SqJUV_gy2eZ-Q1r7NKbpMLhLjgy8OU-XAQeXj6m-3i_K1Pef3Vdgln83142_fLqapyBxTARNVm-DqjEnW2qE8GrPyvvHsD2Afv0APKYlicU63u_3f_w=s0-d-e1-ft#https://d1.awsstatic.com/logos/aws-logo/full-color/AWS-Logo_Full-Color_100x60.5e9a396f44e8a57089a6c62488cd51517ae74f4b.png" align="right" hspace="6" vspace="6" alt="Description: awslogo" title="Description: awslogo" class="CToWUd" data-bit="iit">				</div>
				<div class="header">
					<h1>Account summary</h1>
				</div>
				<div class="summary">
					Dear ` + userName + `, here is a summary of your account.
				</div>
				<table>
					<tr>
						<th>Total balance:</th>
						<td>` + totalBalance + `</td>
					</tr>
					<tr>
						<th>Transactions per month:</th>
						<td>` + numberOfTransactionsHTML.String() + `</td>
					</tr>
					<tr>
						<th>Average debit amount:</th>
						<td>` + averageDebitAmount + `</td>
					</tr>
					<tr>
						<th>Average credit amount:</th>
						<td>` + averageCreditAmount + `</td>
					</tr>
				</table>
			</div>
		</body>
		</html>
		`
	abc.SetHeader("From", "reydward@gmail.com")
	abc.SetHeader("To", "info@brainsmartsolutions.com")
	abc.SetHeader("Subject", "Test Email Subject abc")
	abc.SetBody("text/html", body)

	a := gomail.NewDialer("smtp.gmail.com", 587, "reydward@gmail.com", "juti gdfm qauq kymt")

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
