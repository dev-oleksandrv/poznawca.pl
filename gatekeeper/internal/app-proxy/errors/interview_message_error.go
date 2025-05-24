package errors

type InterviewMessageErrorKey string

const (
	InterviewMessageErrorKeyFailedToInitializeInterview       InterviewMessageErrorKey = "failed_to_initialize_interview"
	InterviewMessageErrorKeyFailedToSendMessage               InterviewMessageErrorKey = "failed_to_send_message"
	InterviewMessageErrorKeyFailedToProcessMessage            InterviewMessageErrorKey = "failed_to_process_message"
	InterviewMessageErrorKeyFailedToGetResults                InterviewMessageErrorKey = "failed_to_get_results"
	InterviewMessageErrorKeyFailedToCompleteInterview         InterviewMessageErrorKey = "failed_to_complete_interview"
	InterviewMessageErrorKeyFailedToCheckCompleteAvailability InterviewMessageErrorKey = "failed_to_check_complete_availability"
)
