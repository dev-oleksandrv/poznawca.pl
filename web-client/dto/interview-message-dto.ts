export enum InterviewMessageRole {
  User = "user",
  Interviewer = "interviewer",
  System = "system",
}

export enum InterviewMessageType {
  Default = "default",
  Error = "error",
}

export interface InterviewMessageDto {
  id: string;
  content_text: string;
  content_translation_text: string;
  tips_text: string;
  role: InterviewMessageRole;
  type: InterviewMessageType;
  is_last_message: boolean;
}
