/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CreateThreadAndRunRequest struct {
	// The ID of the [assistant](/docs/api-reference/assistants) to use to execute this run.
	AssistantId string `json:"assistant_id"`

	Thread *CreateThreadRequest `json:"thread,omitempty"`
	// The ID of the [Model](/docs/api-reference/models) to be used to execute this run. If a value is provided here, it will override the model associated with the assistant. If not, the model associated with the assistant will be used.
	Model string `json:"model,omitempty"`
	// Override the default system message of the assistant. This is useful for modifying the behavior on a per-run basis.
	Instructions string `json:"instructions,omitempty"`
	// Override the tools the assistant can use for this run. This is useful for modifying the behavior on a per-run basis.
	Tools []OneOfCreateThreadAndRunRequestToolsItems `json:"tools,omitempty"`
	// Set of 16 key-value pairs that can be attached to an object. This can be useful for storing additional information about the object in a structured format. Keys can be a maximum of 64 characters long and values can be a maxium of 512 characters long. 
	Metadata *interface{} `json:"metadata,omitempty"`
	// If `true`, returns a stream of events that happen during the Run as server-sent events, terminating when the Run enters a terminal state with a `data: [DONE]` message. 
	Stream bool `json:"stream,omitempty"`
}
