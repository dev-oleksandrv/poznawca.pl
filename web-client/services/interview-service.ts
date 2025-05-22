import { KyInstance } from "ky";
import { APIResponseWrapper } from "@/dto/common-dto";
// import { InterviewDto } from "@/dto/interview-dto";
import { apiClient } from "@/api/client";

export class InterviewService {
  private readonly URL_PREFIX = "interview";

  constructor(private readonly apiClient: KyInstance) {}

  // public getByID(id: string) {
  //   return this.apiClient.get(`${this.URL_PREFIX}/${id}`).json<APIResponseWrapper<InterviewDto>>();
  // }
  //
  // public createInterview() {
  //   return this.apiClient.post(this.URL_PREFIX).json<APIResponseWrapper<InterviewDto>>();
  // }
}

export const interviewServiceSingleton = new InterviewService(apiClient);
