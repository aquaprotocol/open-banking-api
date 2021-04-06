package swagger

type AccountInfoDb struct{
	BookingBalance string `json:"booking_balance"`
	Currency string `json:"currency"`
	AvailableBalance string `json:"available_balance"`
	AccountNumber string `json:"account_number"`
	AccountTypeName string `json:"account_type_name"`

/*	return AccountInfo{
	BookingBalance: "125.45", Currency: "PLN",
	AvailableBalance: "125.45",
	AccountNumber: "123456789977", AccountTypeName: "Junior"}*/
}
