/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A chat completion delta generated by streamed model responses.
type ChatCompletionStreamResponseDelta struct {
	// The contents of the chunk message.
	Content string `json:"content,omitempty"`

	FunctionCall *ChatCompletionStreamResponseDeltaFunctionCall `json:"function_call,omitempty"`

	ToolCalls []ChatCompletionMessageToolCallChunk `json:"tool_calls,omitempty"`
	// The role of the author of this message.
	Role string `json:"role,omitempty"`
}
