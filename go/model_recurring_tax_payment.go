/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Struktura danych przelewu do urzędu podatkowego, użyta do zdefinowania nowej płatności cyklicznej. / Data structure of tax payment that is to be used to define new recurring payment.
type RecurringTaxPayment struct {

	Sender *SenderPisDomestic `json:"sender"`

	// Droga jaką przelew ma być rozliczony / The way the transfer should be settled
	System string `json:"system"`

	// Tryb pilności / Urgency mode
	DeliveryMode string `json:"deliveryMode"`

	UsInfo *TransactionInfoTax `json:"usInfo,omitempty"`

	TransferData *TransferDataCurrencyRequiredTax `json:"transferData"`

	Recipient *RecipientPisTax `json:"recipient"`

	// Czy założyć blokadę (w przypadku np. zlecenia przelewu w dniu wolnym) / Indicates if payment should be holded
	Hold bool `json:"hold,omitempty"`
}
