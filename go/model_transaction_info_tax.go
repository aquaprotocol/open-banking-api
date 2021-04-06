/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa informacji danych dla przelewu podatkowego do Urzędu Skarbowego lub Izby Celnej / Tax Transfer
type TransactionInfoTax struct {

	PayerInfo *Payor `json:"payerInfo"`

	// Identyfikator zobowiązania, z którego wynika należność podatku np. decyzja, tytuł wykonawczy, postanowienie / Obligation ID
	ObligationId string `json:"obligationId,omitempty"`

	// Typ okresu. Wymagany warunkowo - w zależności od wartości parametru formCode. / Period type. Required conditionally - depending on formCode parameter value.
	PeriodType string `json:"periodType"`

	// Numer okresu. Wymagany warunkowo - w zależności od wartości parametru formCode. / Period number. Required conditionally - depending on formCode parameter value.
	PeriodId string `json:"periodId"`

	// Symbol formularza Urzędu Skarbowego lub Izby Celnej / Tax authority form symbol
	FormCode string `json:"formCode"`

	// Rok okresu. Wymagany warunkowo - w zależności od wartości parametru formCode. / Period year. Required conditionally - depending on formCode parameter value.
	Year int32 `json:"year"`
}