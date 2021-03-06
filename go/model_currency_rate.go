/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa zawierająca kursy przewalutowań / Currency rate
type CurrencyRate struct {

	// Waluta rachunku, kod ISO / to Currency, ISO code
	ToCurrency string `json:"toCurrency,omitempty"`

	// Waluta rachunku, kod ISO / from Currency, ISO code
	FromCurrency string `json:"fromCurrency,omitempty"`

	// Kursy przewalutowania / Currency exchange rate
	Rate float64 `json:"rate,omitempty"`
}
