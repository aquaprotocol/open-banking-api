/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa opisująca informacje wymagane do pobrania informacji o statusie zainicjowanej płatności w ramach zapytania zbiorczego / Class defining a pair of information required to retreive the status of initiated payment
type PaymentTokenEntry struct {

	// Access token użyty do zainicjowania płatności / Access token used to initiate payment
	AccessToken string `json:"accessToken"`

	PaymentId *PaymentId `json:"paymentId,omitempty"`

	// Identyfikator transakcji nadany przez TPP. Wymagany warunkowo - jeśli TPP nie otrzymał identyfikatora paymentId od ASPSP. / Identifier of transaction established by TPP. Conditionally required - in case TPP has not received paymentId from ASPSP.
	TppTransactionId string `json:"tppTransactionId,omitempty"`
}
