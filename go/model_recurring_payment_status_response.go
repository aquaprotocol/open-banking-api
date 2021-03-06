/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa odpowiedzi na pytanie o status płatności cyklicznej / Response data with status of recurring payment
type RecurringPaymentStatusResponse struct {

	RecurringPaymentId *RecurringPaymentId `json:"recurringPaymentId"`

	// Identyfikator płatności cyklicznej nadany przez TPP. / Recurring payment identifier set by TPP.
	TppRecurringPaymentId string `json:"tppRecurringPaymentId"`

	// Szczegółowy status płatności cyklicznej / Recurring payment detailed status
	RecurringPaymentDetailedStatus string `json:"recurringPaymentDetailedStatus,omitempty"`

	// Status płatności cyklicznej / Recurring payment status
	RecurringPaymentStatus string `json:"recurringPaymentStatus,omitempty"`

	ResponseHeader *ResponseHeader `json:"responseHeader"`
}
