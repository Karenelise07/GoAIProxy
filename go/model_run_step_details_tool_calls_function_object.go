/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RunStepDetailsToolCallsFunctionObject struct {
	// The ID of the tool call object.
	Id string `json:"id"`
	// The type of tool call. This is always going to be `function` for this type of tool call.
	Type_ string `json:"type"`

	Function *RunStepDetailsToolCallsFunctionObjectFunction `json:"function"`
}
