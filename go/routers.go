/*
 * OpenAI API
 *
 * The OpenAI REST API. Please see https://platform.openai.com/docs/api-reference for more details.
 *
 * API version: 2.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	// 为静态文件服务添加路由
	// 请确保将"path/to/your/static/files"替换为您的静态文件实际所在的目录路径
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./html"))))

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"CancelRun",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}/runs/{run_id}/cancel",
		CancelRun,
	},

	Route{
		"CreateAssistant",
		strings.ToUpper("Post"),
		"/v1/assistants",
		CreateAssistant,
	},

	Route{
		"CreateAssistantFile",
		strings.ToUpper("Post"),
		"/v1/assistants/{assistant_id}/files",
		CreateAssistantFile,
	},

	Route{
		"CreateMessage",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}/messages",
		CreateMessage,
	},

	Route{
		"CreateRun",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}/runs",
		CreateRun,
	},

	Route{
		"CreateThread",
		strings.ToUpper("Post"),
		"/v1/threads",
		CreateThread,
	},

	Route{
		"CreateThreadAndRun",
		strings.ToUpper("Post"),
		"/v1/threads/runs",
		CreateThreadAndRun,
	},

	Route{
		"DeleteAssistant",
		strings.ToUpper("Delete"),
		"/v1/assistants/{assistant_id}",
		DeleteAssistant,
	},

	Route{
		"DeleteAssistantFile",
		strings.ToUpper("Delete"),
		"/v1/assistants/{assistant_id}/files/{file_id}",
		DeleteAssistantFile,
	},

	Route{
		"DeleteThread",
		strings.ToUpper("Delete"),
		"/v1/threads/{thread_id}",
		DeleteThread,
	},

	Route{
		"GetAssistant",
		strings.ToUpper("Get"),
		"/v1/assistants/{assistant_id}",
		GetAssistant,
	},

	Route{
		"GetAssistantFile",
		strings.ToUpper("Get"),
		"/v1/assistants/{assistant_id}/files/{file_id}",
		GetAssistantFile,
	},

	Route{
		"GetMessage",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/messages/{message_id}",
		GetMessage,
	},

	Route{
		"GetMessageFile",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/messages/{message_id}/files/{file_id}",
		GetMessageFile,
	},

	Route{
		"GetRun",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/runs/{run_id}",
		GetRun,
	},

	Route{
		"GetRunStep",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/runs/{run_id}/steps/{step_id}",
		GetRunStep,
	},

	Route{
		"GetThread",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}",
		GetThread,
	},

	Route{
		"ListAssistantFiles",
		strings.ToUpper("Get"),
		"/v1/assistants/{assistant_id}/files",
		ListAssistantFiles,
	},

	Route{
		"ListAssistants",
		strings.ToUpper("Get"),
		"/v1/assistants",
		ListAssistants,
	},

	Route{
		"ListMessageFiles",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/messages/{message_id}/files",
		ListMessageFiles,
	},

	Route{
		"ListMessages",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/messages",
		ListMessages,
	},

	Route{
		"ListRunSteps",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/runs/{run_id}/steps",
		ListRunSteps,
	},

	Route{
		"ListRuns",
		strings.ToUpper("Get"),
		"/v1/threads/{thread_id}/runs",
		ListRuns,
	},

	Route{
		"ModifyAssistant",
		strings.ToUpper("Post"),
		"/v1/assistants/{assistant_id}",
		ModifyAssistant,
	},

	Route{
		"ModifyMessage",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}/messages/{message_id}",
		ModifyMessage,
	},

	Route{
		"ModifyRun",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}/runs/{run_id}",
		ModifyRun,
	},

	Route{
		"ModifyThread",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}",
		ModifyThread,
	},

	Route{
		"SubmitToolOuputsToRun",
		strings.ToUpper("Post"),
		"/v1/threads/{thread_id}/runs/{run_id}/submit_tool_outputs",
		SubmitToolOuputsToRun,
	},

	Route{
		"CreateSpeech",
		strings.ToUpper("Post"),
		"/v1/audio/speech",
		CreateSpeech,
	},

	Route{
		"CreateTranscription",
		strings.ToUpper("Post"),
		"/v1/audio/transcriptions",
		CreateTranscription,
	},

	Route{
		"CreateTranslation",
		strings.ToUpper("Post"),
		"/v1/audio/translations",
		CreateTranslation,
	},

	Route{
		"CreateChatCompletion",
		strings.ToUpper("Post"),
		"/v1/chat/completions",
		CreateChatCompletion,
	},

	Route{
		"CreateCompletion",
		strings.ToUpper("Post"),
		"/v1/completions",
		CreateCompletion,
	},

	Route{
		"CreateEmbedding",
		strings.ToUpper("Post"),
		"/v1/embeddings",
		CreateEmbedding,
	},

	Route{
		"CreateFile",
		strings.ToUpper("Post"),
		"/v1/files",
		CreateFile,
	},

	Route{
		"DeleteFile",
		strings.ToUpper("Delete"),
		"/v1/files/{file_id}",
		DeleteFile,
	},

	Route{
		"DownloadFile",
		strings.ToUpper("Get"),
		"/v1/files/{file_id}/content",
		DownloadFile,
	},

	Route{
		"ListFiles",
		strings.ToUpper("Get"),
		"/v1/files",
		ListFiles,
	},

	Route{
		"RetrieveFile",
		strings.ToUpper("Get"),
		"/v1/files/{file_id}",
		RetrieveFile,
	},

	Route{
		"CancelFineTuningJob",
		strings.ToUpper("Post"),
		"/v1/fine_tuning/jobs/{fine_tuning_job_id}/cancel",
		CancelFineTuningJob,
	},

	Route{
		"CreateFineTuningJob",
		strings.ToUpper("Post"),
		"/v1/fine_tuning/jobs",
		CreateFineTuningJob,
	},

	Route{
		"ListFineTuningEvents",
		strings.ToUpper("Get"),
		"/v1/fine_tuning/jobs/{fine_tuning_job_id}/events",
		ListFineTuningEvents,
	},

	Route{
		"ListPaginatedFineTuningJobs",
		strings.ToUpper("Get"),
		"/v1/fine_tuning/jobs",
		ListPaginatedFineTuningJobs,
	},

	Route{
		"RetrieveFineTuningJob",
		strings.ToUpper("Get"),
		"/v1/fine_tuning/jobs/{fine_tuning_job_id}",
		RetrieveFineTuningJob,
	},

	Route{
		"CreateImage",
		strings.ToUpper("Post"),
		"/v1/images/generations",
		CreateImage,
	},

	Route{
		"CreateImageEdit",
		strings.ToUpper("Post"),
		"/v1/images/edits",
		CreateImageEdit,
	},

	Route{
		"CreateImageVariation",
		strings.ToUpper("Post"),
		"/v1/images/variations",
		CreateImageVariation,
	},

	Route{
		"DeleteModel",
		strings.ToUpper("Delete"),
		"/v1/models/{model}",
		DeleteModel,
	},

	Route{
		"ListModels",
		strings.ToUpper("Get"),
		"/v1/models",
		ListModels,
	},

	Route{
		"RetrieveModel",
		strings.ToUpper("Get"),
		"/v1/models/{model}",
		RetrieveModel,
	},

	Route{
		"CreateModeration",
		strings.ToUpper("Post"),
		"/v1/moderations",
		CreateModeration,
	},
}
