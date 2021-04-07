package swagger

import (
	"fmt"
	"github.com/rs/xid"
	"log"
	"strconv"
)

func ProvideDomesticTransfer(pDR PaymentDomesticRequest) AddPaymentResponse  {

	db, err := ProvideDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pDR.Sender.AccountNumber.Value)

	results := db.QueryRow("SELECT currency, available_balance, booking_balance FROM account_info where account_number =?", 	pDR.Sender.AccountNumber.Value)
	var account AccountInfoDb

	err = results.Scan(&account.Currency, &account.AvailableBalance, &account.BookingBalance)

	fmt.Println(account.AvailableBalance)

	availableBalance, _ := strconv.ParseFloat(account.AvailableBalance,64)

	fmt.Println(availableBalance)

	amount,_ := strconv.ParseFloat(pDR.TransferData.Amount,64)

	result := availableBalance - amount


 	db.QueryRow("UPDATE account_info set available_balance=? where account_number =?", result,pDR.Sender.AccountNumber.Value)


	results3 := db.QueryRow("SELECT currency, available_balance, booking_balance FROM account_info where account_number =?", 	pDR.Recipient.AccountNumber.Value)
	var account1 AccountInfoDb

	err = results3.Scan(&account1.Currency, &account1.AvailableBalance, &account1.BookingBalance)


	availableBalance1, _ := strconv.ParseFloat(account1.AvailableBalance,64)

	result1 := availableBalance1 + amount

	db.QueryRow("UPDATE account_info set available_balance=? where account_number =?", result1, pDR.Recipient.AccountNumber.Value)


	fmt.Println(result1)

	guid := xid.New()

	status:= SUBMITTED

	return AddPaymentResponse{
		PaymentId: &PaymentId{ guid.String()},
		GeneralStatus: &status,
	}

}

