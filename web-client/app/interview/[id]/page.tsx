import { InterviewDto, InterviewStatus } from "@/dto/interview-dto";
import { redirect } from "next/navigation";
import { interviewServiceSingleton } from "@/services/interview-service";
import { InterviewChat } from "@/components/interview/interview-chat";
import { InterviewChatStoreWrapper } from "@/components/interview/interview-chat-store-wrapper";
import { InterviewChatAbandoned } from "@/components/interview/interview-chat-abandoned";

interface InterviewPageProps {
  params: Promise<{ id: string }>;
}

export default async function InterviewPage({ params }: InterviewPageProps) {
  const { id } = await params;

  let interview: InterviewDto;
  try {
    const response = await interviewServiceSingleton.getByID(id);
    interview = response.data;
  } catch (error) {
    console.error("Failed to fetch interview:", error);
    return <div>Error loading interview</div>;
  }

  if (interview.status !== InterviewStatus.Pending) {
    if (interview.status === InterviewStatus.Completed) {
      return redirect(`/portal/interview/${id}`);
    }

    return <InterviewChatAbandoned interview={interview} />;
  }

  return (
    <InterviewChatStoreWrapper interview={interview}>
      <InterviewChat />
    </InterviewChatStoreWrapper>
  );
}
