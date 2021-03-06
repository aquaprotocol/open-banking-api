/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Lista atrybutów uprawnienia do zainicjowania paczki przelewu / The list of attributes of the initiate bundle of payments request privilege
type PrivilegeBundleTransfers struct {

	// Lista struktur danych przelewów zwykłych, które zostaną zlecone. Wymagane gdy typeOfTransfers = domestic. / The list of data structures of domestic transfers to be initiated. Required when typeOfTransfers = domestic.
	DomesticTransfers []PaymentDomesticRequestBundled `json:"domesticTransfers,omitempty"`

	// Łączna kwota wszystkich przelewów z paczki / The total amount of all transfers from bundle
	TransfersTotalAmount string `json:"transfersTotalAmount,omitempty"`

	// Typ przelewów, które zostaną zlecone w ramach paczki / The type of transfers that will be initiated through the bundle
	TypeOfTransfers string `json:"typeOfTransfers"`

	// Lista struktur danych przelewów zagranicznych innych niż SEPA, które zostaną zlecone. Wymagane gdy typeOfTransfers = nonEEA. / The list of data structures of non SEPA foreign transfers to be initiated. Required when typeOfTransfers = nonEEA.
	NonEEATransfers []PaymentNonEeaRequestBundled `json:"nonEEATransfers,omitempty"`

	// Lista struktur danych przelewów do urzędu podatkowego, które zostaną zlecone. Wymagane gdy typeOfTransfers = tax. / The list of data structures of tax transfers to be initiated. Required when typeOfTransfers = tax.
	TaxTransfers []PaymentTaxRequestBundled `json:"taxTransfers,omitempty"`

	// Rodzaj limitu zgody / Type of limit of usages
	ScopeUsageLimit string `json:"scopeUsageLimit,omitempty"`

	// Lista struktur danych przelewów zagranicznych SEPA, które zostaną zlecone. Wymagane gdy typeOfTransfers = EEA. / The list of data structures of SEPA foreign transfers to be initiated. Required when typeOfTransfers = EEA.
	EEATransfers []PaymentEeaRequestBundled `json:"EEATransfers,omitempty"`
}
