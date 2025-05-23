import { KyInstance } from "ky";
import { APIResponseWrapper } from "@/dto/common-dto";
import { apiClient } from "@/api/client";
import { CreateInterviewRequestDto, CreateInterviewResponseDto } from "@/dto/interview-dto";

export class InterviewService {
  private readonly URL_PREFIX = "interview";

  constructor(private readonly apiClient: KyInstance) {}

  public createInterview(input: CreateInterviewRequestDto) {
    return this.apiClient
      .post(this.URL_PREFIX, {
        json: input,
      })
      .json<APIResponseWrapper<CreateInterviewResponseDto>>();
  }
}

export const interviewServiceSingleton = new InterviewService(apiClient);
