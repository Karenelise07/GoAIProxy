/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Text output from the Code Interpreter tool call as part of a run step.
type RunStepDetailsToolCallsCodeOutputLogsObject struct {
	// Always `logs`.
	Type_ string `json:"type"`
	// The text output from the Code Interpreter tool call.
	Logs string `json:"logs"`
}
