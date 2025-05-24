import { InterviewMessageDto } from "@/dto/interview-message-dto";
import { InterviewResultDto } from "@/dto/interview-result-dto";

export enum InterviewEventType {
  ErrorMessageSent = "error_message_sent",
  InterviewerMessagePending = "interviewer_message_pending",
  InterviewerMessageSent = "interviewer_message_sent",
  ResultsPending = "results_pending",
  ResultsSent = "results_sent",
  UserMessageSent = "user_message_sent",
  UserCompleteInterview = "user_complete_interview",
  InterviewCompleted = "interview_completed",
}

export interface BaseInterviewEvent {
  type: InterviewEventType;
}

export interface ErrorMessageSentEvent extends BaseInterviewEvent {
  type: InterviewEventType.ErrorMessageSent;
  error_key: string;
}

export interface InterviewerMessagePendingEvent extends BaseInterviewEvent {
  type: InterviewEventType.InterviewerMessagePending;
}

export interface InterviewerMessageSentEvent extends BaseInterviewEvent {
  type: InterviewEventType.InterviewerMessageSent;
  data: InterviewMessageDto;
}

export interface ResultsPendingEvent extends BaseInterviewEvent {
  type: InterviewEventType.ResultsPending;
}

export interface ResultsSentEvent extends BaseInterviewEvent {
  type: InterviewEventType.ResultsSent;
  data: InterviewResultDto;
}

export interface UserMessageSentEvent extends BaseInterviewEvent {
  type: InterviewEventType.UserMessageSent;
  content: string;
}

export interface UserCompleteInterviewEvent extends BaseInterviewEvent {
  type: InterviewEventType.UserCompleteInterview;
}

export interface InterviewCompletedEvent extends BaseInterviewEvent {
  type: InterviewEventType.InterviewCompleted;
}
