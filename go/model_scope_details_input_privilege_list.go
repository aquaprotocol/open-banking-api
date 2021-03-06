/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Pełna lista uprawnień będących przedmiotem zapytania o zgodę PSU / The list of all privileges that are the subject of the request for PSU's consent
type ScopeDetailsInputPrivilegeList struct {

	PiscancelRecurringPayment *PrivilegeCancelRecurringPayment `json:"pis:cancelRecurringPayment,omitempty"`

	Pisrecurring *PrivilegeRecurringPayment `json:"pis:recurring,omitempty"`

	Pisbundle *PrivilegeBundleTransfers `json:"pis:bundle,omitempty"`

	AisgetTransactionsRejected *PrivilegeAisAspspIn `json:"ais:getTransactionsRejected,omitempty"`

	PisnonEEA *PrivilegeForeignTransferNonEea `json:"pis:nonEEA,omitempty"`

	AisgetTransactionsDone *PrivilegeAisAspspIn `json:"ais:getTransactionsDone,omitempty"`

	AisgetTransactionsCancelled *PrivilegeAisAspspIn `json:"ais:getTransactionsCancelled,omitempty"`

	PisgetRecurringPayment *PrivilegeRecurringPaymentStatus `json:"pis:getRecurringPayment,omitempty"`

	Pisdomestic *PrivilegeDomesticTransfer `json:"pis:domestic,omitempty"`

	AisgetAccount *PrivilegeAisAspspInSimple `json:"ais:getAccount,omitempty"`

	PisgetPayment *PrivilegePayment `json:"pis:getPayment,omitempty"`

	PiscancelPayment *PrivilegeCancelPayment `json:"pis:cancelPayment,omitempty"`

	AisgetTransactionDetail *PrivilegeAisAspspInSimple `json:"ais:getTransactionDetail,omitempty"`

	AisgetHolds *PrivilegeAisAspspIn `json:"ais:getHolds,omitempty"`

	PisgetBundle *PrivilegeBundle `json:"pis:getBundle,omitempty"`

	AccountNumber *AccountNumber `json:"accountNumber,omitempty"`

	AisgetTransactionsPending *PrivilegeAisAspspIn `json:"ais:getTransactionsPending,omitempty"`

	AisAccountsgetAccounts *PrivilegeAisAspspInSimple `json:"ais-accounts:getAccounts,omitempty"`

	PisEEA *PrivilegeForeignTransferEea `json:"pis:EEA,omitempty"`

	AisgetTransactionsScheduled *PrivilegeAisAspspIn `json:"ais:getTransactionsScheduled,omitempty"`

	Pistax *PrivilegeTaxTransfer `json:"pis:tax,omitempty"`
}
