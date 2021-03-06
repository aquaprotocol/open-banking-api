/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Klasa zapytania o rejestrację aplikacji klienckiej / Request class for the client application registration
type RegistrationRequest struct {

	SoftwareVersion string `json:"software_version,omitempty"`

	Scope string `json:"scope,omitempty"`

	Address string `json:"address"`

	LogoUri string `json:"logo_uri,omitempty"`

	ClientName string `json:"client_name,omitempty"`

	Contacts []string `json:"contacts,omitempty"`

	TokenEndpointAuthMethod *TokenEndpointAuthMethod `json:"token_endpoint_auth_method,omitempty"`

	GrantTypes []GrantType `json:"grant_types,omitempty"`

	Phone string `json:"phone"`

	ClientUri string `json:"client_uri,omitempty"`

	PolicyUri string `json:"policy_uri,omitempty"`

	Email string `json:"email"`

	RedirectUris []string `json:"redirect_uris,omitempty"`

	SoftwareId string `json:"software_id,omitempty"`

	SoftwareStatement string `json:"software_statement"`

	TosUri string `json:"tos_uri,omitempty"`

	Jwks *Jwks `json:"jwks,omitempty"`

	ResponseTypes []ResponseType `json:"response_types,omitempty"`

	ExampleExtensionParameter string `json:"example_extension_parameter,omitempty"`

	JwksUri string `json:"jwks_uri,omitempty"`

	DeliveryAddress string `json:"delivery_address"`
}
