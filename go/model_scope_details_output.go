/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Zakres, limity, parametry i czas ważności zgód, które potwierdza ASPSP / Scope, limits, parameters and expiration time of consents that are confimed by ASPSP
type ScopeDetailsOutput struct {

	// The list of structures describing details of consent obtained by TPP.
	PrivilegeList []ScopeDetailsOutputPrivilegeList `json:"privilegeList,omitempty"`

	// Data ważności zgody / consent valid until, YYYY-MM-DDThh:mm:ss[.mmm]
	ScopeTimeLimit time.Time `json:"scopeTimeLimit"`

	// Id of consent
	ConsentId string `json:"consentId"`

	// Throttling policy
	ThrottlingPolicy string `json:"throttlingPolicy"`
}