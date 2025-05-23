import { DashboardInterviewSetupError } from "@/components/dashboard/dashboard-interview-setup-error";
import { InterviewerDto } from "@/dto/interviewer-dto";
import { interviewerServiceSingleton } from "@/services/interviewer-service";
import { DashboardInterviewSetup } from "@/components/dashboard/dashboard-interview-setup";

export default async function DashboardInterviewSetupPage() {
  let interviewers: InterviewerDto[] = [];
  try {
    const response = await interviewerServiceSingleton.getList();
    interviewers = response.data;
  } catch (error) {
    console.log(error);
  }

  if (!interviewers.length) {
    return <DashboardInterviewSetupError />;
  }

  return <DashboardInterviewSetup interviewers={interviewers} />;
}
