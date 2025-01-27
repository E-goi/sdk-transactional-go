/*
 * Transactional API
 *
 * # Introduction    This API is a service provided by [E-goi](www.e-goi.com) to send transactional messages.    Transactional Messaging is a 1-to-1 communication channel, usually from an organization directed to a specific consumer.   They can be triggered by:    * __Actions__ - The consumer interacts with the organization (ie.: online shopping);  * __Time__ - The consumer's actions are time-bounded by the organization (ie.: period of inactivity).    Because of the nature of these messages,   it is expected that the consumer is interested in the content of these messages.   Therefore, they have a different treatment from marketing messages, and have a higher acceptance and opening rate.    >DISCLAIMER  >  >Please notice that this platform is more delicate in regards to the nature and processing of its messages.  >  >It should NOT be used as a mean of mass marketing, scams, phishing or overall unruly behaviour.   >  >Failure to comply may lead to limitation of use and even termination of account.    ***
 *
 * API version: V2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type CompleteEmailMessageResponse struct {
	Bcc []string `json:"bcc,omitempty"`
	GroupId int32 `json:"groupId,omitempty"`
	TextBody string `json:"textBody,omitempty"`
	HtmlBody string `json:"htmlBody,omitempty"`
	TemplateId string `json:"templateId,omitempty"`
	SenderId int32 `json:"senderId,omitempty"`
	Id int32 `json:"id,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	OpenTracking bool `json:"openTracking,omitempty"`
	Cc []string `json:"cc,omitempty"`
	Subject string `json:"subject,omitempty"`
	Status string `json:"status,omitempty"`
	Unsubscribe bool `json:"unsubscribe,omitempty"`
	SubscriberId int32 `json:"subscriberId,omitempty"`
	ClickTracking bool `json:"clickTracking,omitempty"`
	DomainId int32 `json:"domainId,omitempty"`
	ReplyToId int32 `json:"replyToId,omitempty"`
	Domain string `json:"domain,omitempty"`
	Name string `json:"name,omitempty"`
	Register bool `json:"register,omitempty"`
}
