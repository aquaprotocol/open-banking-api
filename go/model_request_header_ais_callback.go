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

// Klasa zawierająca informacje o PSU and callback URL / PSU and callback URL Information Class
type RequestHeaderAisCallback struct {

	// Oryginalna data wysłania, format: YYYY-MM-DDThh:mm:ss[.mmm] / Send date
	SendDate time.Time `json:"sendDate,omitempty"`

	// Adres IP końcowego urządzenia PSU. Wymagany dla isDirectPsu=true. / IP address of PSU's terminal device. Required when isDirectPsu=true.
	IpAddress string `json:"ipAddress,omitempty"`

	// Browser agent dla PSU / PSU browser agent
	UserAgent string `json:"userAgent,omitempty"`

	// Identyfikator TPP / TPP ID
	TppId string `json:"tppId"`

	// Identyfikator żądania w formacie UUID (Wariant 1, Wersja 1) zgodnym ze standardem RFC 4122 / Request ID using UUID format (Variant 1, Version 1) described in RFC 4122 standard
	RequestId string `json:"requestId"`

	// API key dla wywołania funkcji zwrotnej / callback API key
	ApiKey string `json:"apiKey,omitempty"`

	// (true/false) Znacznik informujący czy request jest przesłany bezpośrednio przez PSU. Domyślna wartość to false. / Is request sent by PSU directly. false is default value.
	IsDirectPsu bool `json:"isDirectPsu,omitempty"`

	// adres funkcji zwrotnej / callback URL
	CallbackURL string `json:"callbackURL,omitempty"`

	// Token autoryzacyjny / Authorization token
	Token string `json:"token"`
}
