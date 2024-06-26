/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AssistantToolsFunction struct {
	// The type of tool being defined: `function`
	Type_ string `json:"type"`

	Function *FunctionObject `json:"function"`
}
