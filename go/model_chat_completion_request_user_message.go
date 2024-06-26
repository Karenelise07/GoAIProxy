/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ChatCompletionRequestUserMessage struct {
	// The contents of the user message.
	Content string `json:"content"`
	// The role of the messages author, in this case `user`.
	Role string `json:"role"`
	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name string `json:"name,omitempty"`
}
