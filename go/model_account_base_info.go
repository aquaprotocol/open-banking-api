/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa informacji o koncie / Account Information Class
type AccountBaseInfo struct {

	// Numer rachunku (częściowo zamaskowany) / Account number (partly masked)
	AccountNumber string `json:"accountNumber"`

	AccountType *DictionaryItem `json:"accountType"`

	// Nazwa typu rachunku (definiowana przez ASPSP) / Account's type name (defined by ASPSP)
	AccountTypeName string `json:"accountTypeName"`
}
