package message

// BulkMessage usually contains a single recipient per message
// and is generally used to send the same content to many recipients,
// optionally customizing the message via the use of MergeData.
type BulkMessage struct {

	//The message subject.
	Subject string

	// Plain text portion of the message body.
	// (Optional) Either PlainTextBody, HtmlBody, both or ApiTemplate must be used to set the body.
	PlainTextBody string

	// HTML portion of the message body.
	// (Optional) Either PlainTextBody, HtmlBody, both or ApiTemplate must be used to set the body.
	HtmlBody string

	// Api Template for the message body.
	// (Optional) Either PlainTextBody, HtmlBody, both or ApiTemplate must be used to set the body.
	ApiTemplate string

	//Custom MailingId for the message.
	//See https://www.socketlabs.com/blog/best-practices-for-using-custom-mailingids-and-messageids/ for more information.
	//(Optional)
	MailingId string

	//Custom MessageId for the message.
	//(Optional)
	MessageId string

	//From Address for the message.
	From EmailAddress

	//ReplyTo Address for the message.
	//(Optional)
	ReplyTo EmailAddress

	//The optional character set for the message.
	//(Optional)
	CharSet string

	// Optional collection of message attachments.
	// (Optional)
	Attachments []Attachment

	// Optional collection of custom headers for the message.
	// (Optional)
	CustomHeaders []CustomHeader

	//Collection of To Recipients for the message.
	To []BulkRecipient

	// Optional Merge data that is global across the whole message.
	// (Optional)
	Global map[string]string
}

// AddToBulkRecipient add an email address to the To recipients collection
func (bulk *BulkMessage) AddToBulkRecipient(email string) BulkRecipient {
	to := NewBulkRecipient(email)
	bulk.To = append(bulk.To, to)
	return to
}

// AddToFriendlyBulkRecipient adds an email address paired with a friendly name to the To Recipients collection
func (bulk *BulkMessage) AddToFriendlyBulkRecipient(email string, friendly string) BulkRecipient {
	to := NewFriendlyBulkRecipient(email, friendly)
	bulk.To = append(bulk.To, to)
	return to
}

// AddCustomHeader adds a custom header to the message
func (bulk *BulkMessage) AddCustomHeader(name string, value string) {
	customHeader := NewCustomHeader(name, value)
	bulk.CustomHeaders = append(bulk.CustomHeaders, customHeader)
}

// AddGlobalMergeData adds global merge data
func (bulk *BulkMessage) AddGlobalMergeData(key string, value string) {
	if bulk.Global == nil || len(bulk.Global) == 0 {
		bulk.Global = make(map[string]string)
	}
	bulk.Global[key] = value
}
