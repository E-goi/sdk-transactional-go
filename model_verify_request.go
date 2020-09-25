/*
 * Transactional API
 *
 * # Introduction    This API is a service provided by [E-goi](www.e-goi.com) to send transactional messages.    Transactional Messaging is a 1-to-1 communication channel, usually from an organization directed to a specific consumer.   They can be triggered by:    * __Actions__ - The consumer interacts with the organization (ie.: online shopping);  * __Time__ - The consumer's actions are time-bounded by the organization (ie.: period of inactivity).    Because of the nature of these messages,   it is expected that the consumer is interested in the content of these messages.   Therefore, they have a different treatment from marketing messages, and have a higher acceptance and opening rate.    >DISCLAIMER  >  >Please notice that this platform is more delicate in regards to the nature and processing of its messages.  >  >It should NOT be used as a mean of mass marketing, scams, phishing or overall unruly behaviour.   >  >Failure to comply may lead to limitation of use and even termination of account.    ***
 *
 * API version: V2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerifyRequest struct {
	// The id of the sender.
	SenderId string `json:"senderId"`
	// The phone number of the recipient of the code. Format should be `countryCode`-`phoneNumber`
	To string `json:"to"`
	// the message that the user is going to receive. location of the code should be given by using the tag {{ code }}. All instances of {{ code }} will be replaced with a code that is in accordance with the settings. If the tag is not present, the code will be added at the end of the message.
	Message string `json:"message,omitempty"`
	// the length of the code in characters. Maximum is 20 characters. Default is 6 characters.
	CodeLength int32 `json:"codeLength,omitempty"`
	// the period of time the code is valid after the message has been sent, in seconds. Defaults to 600 seconds (10 min). Maximum validity is 3600 seconds (1 hour).
	CodeValidity int32 `json:"codeValidity,omitempty"`
	// the format of the code. can be:  Format | Description | RegEx --- | --- | --- numeric | numbers | [0-9] alpha | letters | [A-Z] alphanumeric | numbers and letters | [0-9A-Z] hexa | numbers and letters up to F | [0-9A-F]
	CodeType string `json:"codeType,omitempty"`
}
