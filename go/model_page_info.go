/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa zawierająca dane pozwalajace na korzystanie z mechanizmu stronicowania / Paging Information Class
type PageInfo struct {

	// Użyte w celu stronicowania danych:  Identyfikator następnej strony rezultatów / Used for paging the results. Identifier of the next page.
	NextPage string `json:"nextPage,omitempty"`

	// Użyte w celu stronicowania danych: Identyfikator poprzedniej strony rezultatów / Used for paging the results. Identifier of the previous page.
	PreviousPage string `json:"previousPage,omitempty"`
}