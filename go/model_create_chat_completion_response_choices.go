/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateChatCompletionResponseChoices struct {
	// The reason the model stopped generating tokens. This will be `stop` if the model hit a natural stop point or a provided stop sequence, `length` if the maximum number of tokens specified in the request was reached, `content_filter` if content was omitted due to a flag from our content filters, `tool_calls` if the model called a tool, or `function_call` (deprecated) if the model called a function. 
	FinishReason string `json:"finish_reason"`
	// The index of the choice in the list of choices.
	Index int32 `json:"index"`

	Message *ChatCompletionResponseMessage `json:"message"`

	Logprobs *CreateChatCompletionResponseLogprobs `json:"logprobs"`
}
