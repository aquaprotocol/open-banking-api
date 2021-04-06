/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TokenEndpointAuthMethod string

// List of TokenEndpointAuthMethod
const (
	CLIENT_SECRET_POST TokenEndpointAuthMethod = "client_secret_post"
	CLIENT_SECRET_BASIC TokenEndpointAuthMethod = "client_secret_basic"
	NONES TokenEndpointAuthMethod = "none"
)
