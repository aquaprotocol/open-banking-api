/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa zapytania o anulowanie płatności cyklicznej. / Recurring payment cancelation request class.
type CancelRecurringPaymentRequest struct {

	RecurringPaymentId *RecurringPaymentId `json:"recurringPaymentId"`

	RequestHeader *RequestHeader `json:"requestHeader"`
}