/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// References an image [File](/docs/api-reference/files) in the content of a message.
type MessageContentImageFileObject struct {
	// Always `image_file`.
	Type_ string `json:"type"`

	ImageFile *MessageContentImageFileObjectImageFile `json:"image_file"`
}
