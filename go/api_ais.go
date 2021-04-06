/*
 * PKO BP API 2.1.1
 *
 * Specyfikacja interfejsu dla usług świadczonych przez ASPSP w oparciu o dostęp do rachunków płatniczych. Opracowane przez Związek Banków Polskich i podmioty współpracujące / Interface specification for services provided by ASPSP based on access to payment accounts. Prepared by the Polish Bank Association and its affiliates
 *
 * API version: 1.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */


/*{"invokeDeposition":{"depType":"C","accountNumberFrom":"0xa7578120e2550984701fcCC7dbfa44Ae5A89bFD2","accountNumberTo":"","amount":"1","ethAccount":"0x4e52b55514FD23f2F86926230eFC8a5c0a694418"}}*/
package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DeleteConsent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

type name struct {
	AccountNumber string `json:"accountNumber"`
}

//curl -X POST localhost:8080/v2_1_1.1/accounts/v2_1_1.1/getAccount -d '{"accountNumber":"58150017805507845039441513"}'
func GetAccount(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("GetAccount request body error")
	}

	var aIR AccountInfoRequest
	err = json.Unmarshal(b, &aIR)
	fmt.Println("GetAccount account number from request " + aIR.AccountNumber)

	pAI := ProvideAccountInfo(&aIR.AccountNumber)

	js,_ :=json.Marshal(pAI)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetHolds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTransactionsCancelled(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTransactionsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTransactionsPending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTransactionsRejected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetTransactionsScheduled(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}