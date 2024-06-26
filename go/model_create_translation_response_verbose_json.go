/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateTranslationResponseVerboseJson struct {
	// The language of the output translation (always `english`).
	Language string `json:"language"`
	// The duration of the input audio.
	Duration string `json:"duration"`
	// The translated text.
	Text string `json:"text"`
	// Segments of the translated text and their corresponding details.
	Segments []TranscriptionSegment `json:"segments,omitempty"`
}
