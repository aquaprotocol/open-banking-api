/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa odpowiedzi na zapytanie TPP o autoryzację PSU do wykonania usługi interfejsu XS2A
type AuthorizeResponse struct {

	// Adres po stronie ASPSP, na który TPP powinien dokonać przekierowania przeglądarki PSU w celu jego uwierzytelnienia
	AspspRedirectUri string `json:"aspspRedirectUri"`

	ResponseHeader *ResponseHeader `json:"responseHeader"`
}
