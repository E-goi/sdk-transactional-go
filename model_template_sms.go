/*
 * Transactional API
 *
 * # Introduction    This API is a service provided by [E-goi](www.e-goi.com) to send transactional messages.    Transactional Messaging is a 1-to-1 communication channel, usually from an organization directed to a specific consumer.   They can be triggered by:    * __Actions__ - The consumer interacts with the organization (ie.: online shopping);  * __Time__ - The consumer's actions are time-bounded by the organization (ie.: period of inactivity).    Because of the nature of these messages,   it is expected that the consumer is interested in the content of these messages.   Therefore, they have a different treatment from marketing messages, and have a higher acceptance and opening rate.    >DISCLAIMER  >  >Please notice that this platform is more delicate in regards to the nature and processing of its messages.  >  >It should NOT be used as a mean of mass marketing, scams, phishing or overall unruly behaviour.   >  >Failure to comply may lead to limitation of use and even termination of account.    ***
 *
 * API version: V2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TemplateSms struct {
	To []string `json:"to"`
	// The id of the sender.
	From string `json:"from"`
	// The id of the template.
	TemplateId string `json:"templateId"`
	// A list of the merge tags to replace information gaps in the email. The format of each tag should be ``tagName:tagValue``
	MergeTags []string `json:"mergeTags,omitempty"`
	// The priority of the message.
	Priority string `json:"priority,omitempty"`
	// Extra Data related to the campaign to be re-sent via WebHook.
	CustomData string `json:"customData,omitempty"`
	// option to register the message. Its an exclusive feature for senders that require it.
	Registered bool `json:"registered,omitempty"`
	// The group of the message.
	Group string `json:"group,omitempty"`
}
