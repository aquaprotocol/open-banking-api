/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Lista atrybutów uprawnienia do żądania statusu płatności / The list of attributes of the payment status request privilege
type PrivilegePayment struct {

	// Rodzaj limitu zgody / Type of limit of usages
	ScopeUsageLimit string `json:"scopeUsageLimit,omitempty"`

	PaymentId *PaymentId `json:"paymentId,omitempty"`

	// Identyfikator transakcji nadany przez TPP. Jeżeli to pole jest puste należy uzupełnić pole paymentId. / Identifier of transaction established by TPP. If this field is empty then paymentId must be filled.
	TppTransactionId string `json:"tppTransactionId,omitempty"`
}
