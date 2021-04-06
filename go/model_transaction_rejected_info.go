/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Klasa opisująca odrzuconą transakcję płatniczą
type TransactionRejectedInfo struct {

	// Typ transakcji / Transaction type
	TransactionType string `json:"transactionType,omitempty"`

	// Data operacji, YYYY-MM-DDThh:mm:ss[.mmm] / Date of the operation
	TradeDate time.Time `json:"tradeDate,omitempty"`

	// Kod ISO waluty transakcji / Currency (ISO)
	Currency string `json:"currency,omitempty"`

	// Tytuł transakcji / Description of the transaction
	Description string `json:"description,omitempty"`

	AuxData *ModelMap `json:"auxData,omitempty"`

	// Kwota transakcji / Amount of the transaction
	Amount string `json:"amount"`

	// ID elementu (transakcji lub blokadzie) nadany przez ASPSP / Element (transaction or hold) ID (ASPSP)
	ItemId string `json:"itemId"`

	// Kod dla każdej transakcji/operacji wykonanej przy użyciu karty / A code of each transaction performed by card
	Mcc string `json:"mcc,omitempty"`

	Sender *SenderRecipient `json:"sender,omitempty"`

	// Powod odrzucenia
	RejectionReason string `json:"rejectionReason,omitempty"`

	Recipient *SenderRecipient `json:"recipient,omitempty"`

	Initiator *NameAddress `json:"initiator,omitempty"`

	// Kategoria transakcji obciążenie / Transaction category (debit)
	TransactionCategory string `json:"transactionCategory"`

	// Data odrzucenia, YYYY-MM-DDThh:mm:ss[.mmm]
	RejectionDate time.Time `json:"rejectionDate,omitempty"`
}