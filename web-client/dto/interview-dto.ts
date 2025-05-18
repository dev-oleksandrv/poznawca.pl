export enum InterviewStatus {
  Pending = "pending",
  Active = "active",
  Abandoned = "abandoned",
  Completed = "completed",
}

export enum InterviewMessageRole {
  User = "user",
  Interviewer = "interviewer",
  System = "system",
}

export interface InterviewerDto {
  id: string;
  name: string;
  avatar_url: string;
  entry_message: string;
}

export interface InterviewMessageDto {
  id: string;
  content_text: string;
  content_translation_text: string;
  tips_text: string;
  role: InterviewMessageRole;
}

export interface InterviewDto {
  id: string;
  status: InterviewStatus;
  interviewer: InterviewerDto;
}
