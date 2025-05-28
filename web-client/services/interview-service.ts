import { KyInstance } from "ky";
import { APIResponseWrapper } from "@/dto/common-dto";
import { apiClient } from "@/api/client";
import {
  CreateInterviewRequestDto,
  CreateInterviewResponseDto,
  InterviewDto,
} from "@/dto/interview-dto";

export class InterviewService {
  private readonly URL_PREFIX = "interview";

  constructor(private readonly apiClient: KyInstance) {}

  public getByID(id: string) {
    return this.apiClient.get(`${this.URL_PREFIX}/${id}`).json<APIResponseWrapper<InterviewDto>>();
  }

  public getList() {
    return this.apiClient.get(`${this.URL_PREFIX}/list`).json<APIResponseWrapper<InterviewDto[]>>();
  }

  public createInterview(input: CreateInterviewRequestDto) {
    return this.apiClient
      .post(this.URL_PREFIX, {
        json: input,
      })
      .json<APIResponseWrapper<CreateInterviewResponseDto>>();
  }
}

export const interviewServiceSingleton = new InterviewService(apiClient);
