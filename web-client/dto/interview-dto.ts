import { InterviewMessageDto } from "@/dto/interview-message-dto";
import { InterviewerDto } from "@/dto/interviewer-dto";
import { InterviewResultDto } from "@/dto/interview-result-dto";

export enum InterviewStatus {
  Pending = "pending",
  Active = "active",
  Completed = "completed",
  Abandoned = "abandoned",
}

export interface CreateInterviewRequestDto {
  interviewer_id?: string;
}

export interface CreateInterviewResponseDto {
  id: string;
  status: string;
}

export interface InterviewDto {
  id: string;
  status: InterviewStatus;
  interviewer: InterviewerDto;
  messages: InterviewMessageDto[];
  result: InterviewResultDto | null;
  updated_at: string;
}
