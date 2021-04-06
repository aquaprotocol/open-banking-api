/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa odpowiedzi na żądanie zdefiniowania nowej płatności cyklicznej. / Data returned for the request of creation of new recurring payment
type RecurringPaymentResponse struct {

	RecurringPaymentId *RecurringPaymentId `json:"recurringPaymentId,omitempty"`

	// Status płatności cyklicznej / Status of recurring payment
	RecurringPaymentStatus string `json:"recurringPaymentStatus,omitempty"`

	Recurrence *RecurringTransferParameters `json:"recurrence,omitempty"`

	// Szczegółowy status płatności cyklicznej / Recurring payment detailed status
	RecurringPaymentDetailedStatus string `json:"recurringPaymentDetailedStatus,omitempty"`
}