import { KyInstance } from "ky";
import { APIResponseWrapper } from "@/dto/common-dto";
import { InterviewerDto } from "@/dto/interviewer-dto";
import { apiClient } from "@/api/client";

export class InterviewerService {
  private readonly URL_PREFIX = "interviewer";

  constructor(private readonly apiClient: KyInstance) {}

  public getList() {
    return this.apiClient
      .get(`${this.URL_PREFIX}/list`)
      .json<APIResponseWrapper<InterviewerDto[]>>();
  }
}

export const interviewerServiceSingleton = new InterviewerService(apiClient);
