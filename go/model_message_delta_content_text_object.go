/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The text content that is part of a message.
type MessageDeltaContentTextObject struct {
	// The index of the content part in the message.
	Index int32 `json:"index"`
	// Always `text`.
	Type_ string `json:"type"`

	Text *MessageDeltaContentTextObjectText `json:"text,omitempty"`
}
