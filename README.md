# open-banking-api

Chainlink Virtual Hackathon Spring 2021

The European Union's PSD2 directive allows third parties to access bank functions. The Open Banking API is a response to these problems. It describes Payment Initiation Service (PIS), Account Information Service (AIS) and Confirmation of the Availability of Funds (CAF). In the AquaProtocol now only domestic payment is implemented in fiat currency deposition and withdrawal. This component is not only API it is also mock bank account functions which are invoked during change owner procedure. Thanks to this approach, the testing process is not dependent on Bank Access.

Base of the API was created from the swagger file obtained from https://developers.pkobp.pl/documentation

For purpose of this project [ProvideDomesticTransfer](https://github.com/aquaprotocol/open-banking-api/blob/61148fcb1a1d88e3d847cfa98627d9a455f8a863/go/api_pis.go#L43) function is mocking.

### ProvideDomesticTransfer
1. Read data from database
2. Modification of the Sender account balance
3. Modification of the Recipient account balance
4. Save data to database/ update amount state
5. Create AddPaymentResponse


Request:
```go
type PaymentDomesticRequest struct {

    Sender *SenderPisDomestic `json:"sender"`

    // The way the transfer should be settled
    System string `json:"system"`

    // Payment execution mode. The superior information deciding which mode is to be used to execute payment.
    ExecutionMode string `json:"executionMode,omitempty"`

    RequestHeader *RequestHeaderCallback `json:"requestHeader"`

    // Tryb pilności / Urgency mode
    DeliveryMode string `json:"deliveryMode"`

    Recipient *RecipientPis `json:"recipient"`

    TransferData *TransferDataCurrencyRequired `json:"transferData"`

    // Indicates if payment should be holded
    Hold bool `json:"hold,omitempty"`

    // Transaction ID (TPP)
    TppTransactionId string `json:"tppTransactionId"`
}
```

Response:
```go
type AddPaymentResponse struct {

    // Detailed payment status
    DetailedStatus string `json:"detailedStatus"`

    GeneralStatus *PaymentStatus `json:"generalStatus"`

    PaymentId *PaymentId `json:"paymentId"`

    ResponseHeader *ResponseHeader `json:"responseHeader"`
}
```

### SQL MOCK DB

```sql
CREATE DATABASE bank_mock;

use bank_mock;

CREATE TABLE account_info (
account_number VARCHAR(30),
booking_balance FLOAT(20,2),
currency VARCHAR(20),
available_balance FLOAT(20,2),
account_type_name VARCHAR(255)
);

show tables;
#drop table account_info;
select * from account_info;

INSERT INTO account_info (account_number, booking_balance, currency, available_balance, account_type_name)
VALUES ('58150017805507845039441513', 100.00, 'PLN', '100.00', 'Junior');

INSERT INTO account_info (account_number, booking_balance, currency, available_balance, account_type_name)
VALUES ('45858110278832612545618695', 100.00, 'PLN', '100.00', 'Firma');
```
### Run service
`go run main.go`




