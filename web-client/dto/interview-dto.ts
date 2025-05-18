export enum InterviewStatus {
  Pending = "pending",
  Active = "active",
  Abandoned = "abandoned",
  Completed = "completed",
}

export interface InterviewerDto {
  id: string;
  name: string;
  avatar_url: string;
  entry_message: string;
}

export interface InterviewDto {
  id: string;
  status: InterviewStatus;
  interviewer: InterviewerDto | null;
}
