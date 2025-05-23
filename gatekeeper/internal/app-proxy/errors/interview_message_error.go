package errors

type InterviewMessageErrorKey string

const (
	InterviewMessageErrorKeyFailedToInitializeInterview InterviewMessageErrorKey = "failed_to_initialize_interview"
	InterviewMessageErrorKeyFailedToSendMessage         InterviewMessageErrorKey = "failed_to_send_message"
)
