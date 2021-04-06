/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa informacji danych dla przelewu do ZUS / ZUS Transfer Information Class
type TransactionInfoZus struct {

	// Identyfikator typu płatności / Payment type ID
	PaymentTypeId string `json:"paymentTypeId,omitempty"`

	PayerInfo *SocialSecurityPayor `json:"payerInfo,omitempty"`

	// Numer tytułu wykonawczego / Obligation identifier or number
	ObligationId string `json:"obligationId,omitempty"`

	// Numer deklaracji / Contribution identifier
	ContributionId string `json:"contributionId,omitempty"`

	// Okres deklaracji / Contribution period, Format: MMYYYY
	ContributionPeriod string `json:"contributionPeriod,omitempty"`

	// Typ wpłaty / Contribution type
	ContributionType string `json:"contributionType,omitempty"`
}
