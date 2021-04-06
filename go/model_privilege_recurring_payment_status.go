/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Lista atrybutów uprawnienia do żądania statusu płatności cyklicznej / The list of attributes of recurring payment status request privilege
type PrivilegeRecurringPaymentStatus struct {

	RecurringPaymentId *RecurringPaymentId `json:"recurringPaymentId,omitempty"`

	// Rodzaj limitu zgody / Type of limit of usages
	ScopeUsageLimit string `json:"scopeUsageLimit,omitempty"`

	// Identyfikator płatności cyklicznej nadany przez TPP. Jeżeli to pole jest puste należy uzupełnić pole recurringPaymentId./ Recurring payment identifier set by TPP. If this field is empty then recurringPaymentId must be filled.
	TppRecurringPaymentId string `json:"tppRecurringPaymentId,omitempty"`
}
