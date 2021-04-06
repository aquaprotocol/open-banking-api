package swagger

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func ProvideAccountInfo(number *string) *AccountInfo {
	fmt.Println("ProvideAccountInfo")


	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bank_mock")

	if err != nil {
		fmt.Println("Account info error in sql open")
	}

	results := db.QueryRow("SELECT account_number, currency, available_balance, booking_balance, account_type_name FROM account_info where account_number =?", number)
	var account AccountInfoDb

	err = results.Scan(&account.AccountNumber, &account.Currency, &account.AvailableBalance, &account.BookingBalance, &account.AccountTypeName)

	defer db.Close()

	return &AccountInfo{
		BookingBalance: account.BookingBalance, Currency: account.Currency,
		AvailableBalance: account.AvailableBalance,
		AccountHolderType: "",
		AccountNumber:    account.AccountNumber, AccountTypeName: account.AccountTypeName}
}
