/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ChatCompletionRequestFunctionMessage struct {
	// The role of the messages author, in this case `function`.
	Role string `json:"role"`
	// The contents of the function message.
	Content string `json:"content"`
	// The name of the function to call.
	Name string `json:"name"`
}
