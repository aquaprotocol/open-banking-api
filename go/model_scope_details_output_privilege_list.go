/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Pełna lista uprawnień, dla których została wyrażona zgoda przez PSU / The list of all privileges that the PSU's consent has been confirmed
type ScopeDetailsOutputPrivilegeList struct {

	PiscancelRecurringPayment *PrivilegeCancelRecurringPayment `json:"pis:cancelRecurringPayment,omitempty"`

	Pisrecurring *PrivilegeRecurringPayment `json:"pis:recurring,omitempty"`

	Pisbundle *PrivilegeBundleTransfers `json:"pis:bundle,omitempty"`

	AisgetTransactionsRejected *PrivilegeAisAspspOut `json:"ais:getTransactionsRejected,omitempty"`

	PisnonEEA *PrivilegeForeignTransferNonEea `json:"pis:nonEEA,omitempty"`

	AisgetTransactionsDone *PrivilegeAisAspspOut `json:"ais:getTransactionsDone,omitempty"`

	AisgetTransactionsCancelled *PrivilegeAisAspspOut `json:"ais:getTransactionsCancelled,omitempty"`

	PisgetRecurringPayment *PrivilegeRecurringPaymentStatus `json:"pis:getRecurringPayment,omitempty"`

	Pisdomestic *PrivilegeDomesticTransfer `json:"pis:domestic,omitempty"`

	AisgetAccount *PrivilegeAisAspspOutSimple `json:"ais:getAccount,omitempty"`

	PisgetPayment *PrivilegePayment `json:"pis:getPayment,omitempty"`

	PiscancelPayment *PrivilegeCancelPayment `json:"pis:cancelPayment,omitempty"`

	AisgetTransactionDetail *PrivilegeAisAspspOutSimple `json:"ais:getTransactionDetail,omitempty"`

	AisgetHolds *PrivilegeAisAspspOut `json:"ais:getHolds,omitempty"`

	PisgetBundle *PrivilegeBundle `json:"pis:getBundle,omitempty"`

	AccountNumber *AccountNumber `json:"accountNumber,omitempty"`

	AisgetTransactionsPending *PrivilegeAisAspspOut `json:"ais:getTransactionsPending,omitempty"`

	AisAccountsgetAccounts *PrivilegeAisAspspOutSimple `json:"ais-accounts:getAccounts,omitempty"`

	PisEEA *PrivilegeForeignTransferEea `json:"pis:EEA,omitempty"`

	AisgetTransactionsScheduled *PrivilegeAisAspspOut `json:"ais:getTransactionsScheduled,omitempty"`

	Pistax *PrivilegeTaxTransfer `json:"pis:tax,omitempty"`
}
