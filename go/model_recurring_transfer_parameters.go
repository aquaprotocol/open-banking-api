/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa definiująca parametry płatności cyklicznej / A class defining recurring payment parameters
type RecurringTransferParameters struct {

	Frequency *RecurringTransferFrequency `json:"frequency"`

	// Ostatnia możliwa data wykonania płatności w definiowanym cyklu. Format: YYYY-MM-DD / The last possible payment execution date (of the defined cycle).
	EndDate string `json:"endDate,omitempty"`

	// Data wykonania pierwszej płatności w definiowanym cyklu. Format: YYYY-MM-DD / First payment execution date (of the defined cycle).
	StartDate string `json:"startDate"`

	// Rodzaj przesunięcia, który należy zastosować do wykonania przelewu w przypadku, gdy dzień wolny jest planowaną datą przelewu / Type of offset that should be used for transfer execution in case of day off being the planned date of transfer
	DayOffOffsetType string `json:"dayOffOffsetType"`
}
