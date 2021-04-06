/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Lista atrybutów uprawnienia usługi AIS dla których została wyrażona zgoda przez PSU / The list of attributes of the privilege of the AIS service that the PSU's consent has been confirmed
type PrivilegeAisAspspOutSimple struct {

	// Rodzaj limitu zgody / Type of limit of usages
	ScopeUsageLimit string `json:"scopeUsageLimit,omitempty"`
}
