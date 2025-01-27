/*
 * Transactional API
 *
 * # Introduction    This API is a service provided by [E-goi](www.e-goi.com) to send transactional messages.    Transactional Messaging is a 1-to-1 communication channel, usually from an organization directed to a specific consumer.   They can be triggered by:    * __Actions__ - The consumer interacts with the organization (ie.: online shopping);  * __Time__ - The consumer's actions are time-bounded by the organization (ie.: period of inactivity).    Because of the nature of these messages,   it is expected that the consumer is interested in the content of these messages.   Therefore, they have a different treatment from marketing messages, and have a higher acceptance and opening rate.    >DISCLAIMER  >  >Please notice that this platform is more delicate in regards to the nature and processing of its messages.  >  >It should NOT be used as a mean of mass marketing, scams, phishing or overall unruly behaviour.   >  >Failure to comply may lead to limitation of use and even termination of account.    ***
 *
 * API version: V2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SmsLengthCalculatorResponse struct {
	// GSM0338 message or not
	Gsm0338 bool `json:"gsm0338,omitempty"`
	// Number of characters in the SMS text
	CharCount int32 `json:"charCount,omitempty"`
	// number of parts that compose the message
	Parts int32 `json:"parts,omitempty"`
	// Number of characters per message
	CharPerMessage int32 `json:"charPerMessage,omitempty"`
	// SMS Text Message
	Message string `json:"message,omitempty"`
}
