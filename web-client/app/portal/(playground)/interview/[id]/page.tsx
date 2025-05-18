import { interviewServiceSingleton } from "@/services/interview-service";
import { InterviewStatus } from "@/dto/interview-dto";
import { InterviewChat } from "@/components/interview/interview-chat";

interface InternalInterviewPageProps {
  params: Promise<{
    id: string;
  }>;
}

export default async function InternalInterviewPage({ params }: InternalInterviewPageProps) {
  const { id } = await params;

  let interview = (await interviewServiceSingleton.getByID(id)).data;

  if (interview.status !== InterviewStatus.Pending) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <p className="text-red-500">Interview is not in a pending state.</p>
      </div>
    );
  }

  return <InterviewChat interview={interview} />;
}
