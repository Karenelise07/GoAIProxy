/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Usage statistics related to the run step. This value will be `null` while the run step's status is `in_progress`.
type RunStepCompletionUsage struct {
	// Number of completion tokens used over the course of the run step.
	CompletionTokens int32 `json:"completion_tokens"`
	// Number of prompt tokens used over the course of the run step.
	PromptTokens int32 `json:"prompt_tokens"`
	// Total number of tokens used (prompt + completion).
	TotalTokens int32 `json:"total_tokens"`
}
