import { InterviewDto, InterviewStatus } from "@/dto/interview-dto";
import { interviewServiceSingleton } from "@/services/interview-service";
import { redirect } from "next/navigation";
import { DashboardInterviewResult } from "@/components/dashboard/dashboard-interview-result";

interface InterviewResultPageProps {
  params: Promise<{ id: string }>;
}

export default async function InterviewResultPage({ params }: InterviewResultPageProps) {
  const { id } = await params;

  let interview: InterviewDto;
  try {
    const response = await interviewServiceSingleton.getByID(id);
    interview = response.data;
  } catch (error) {
    console.error("Failed to fetch interview:", error);
    return <div>Error loading interview</div>;
  }

  if (interview.status !== InterviewStatus.Completed) {
    return redirect("/portal");
  }

  return <DashboardInterviewResult interview={interview} />;
}
