/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa zapytania o status płatności cyklicznej / Recurring payment status Request Class
type RecurringPaymentStatusRequest struct {

	RecurringPaymentId *RecurringPaymentId `json:"recurringPaymentId,omitempty"`

	// Identyfikator płatności cyklicznej nadany przez TPP. Wymagany warunkowo - jeśli TPP nie otrzymał identyfikatora recurringPaymentId od ASPSP. / Recurring payment identifier set by TPP. Conditionally required - in case TPP has not received recurringPaymentId from ASPSP.
	TppRecurringPaymentId string `json:"tppRecurringPaymentId,omitempty"`

	RequestHeader *RequestHeader `json:"requestHeader"`
}
