/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TranscriptionWord struct {
	// The text content of the word.
	Word string `json:"word"`
	// Start time of the word in seconds.
	Start float32 `json:"start"`
	// End time of the word in seconds.
	End float32 `json:"end"`
}
