import { useState } from "react";
import { useRouter } from "next/navigation";
import { InterviewerDto } from "@/dto/interviewer-dto";
import { CreateInterviewRequestDto } from "@/dto/interview-dto";
import { interviewServiceSingleton } from "@/services/interview-service";

export function useSetupInterviewHandler() {
  const router = useRouter();

  const [selectedInterviewer, setSelectedInterviewer] = useState<InterviewerDto | null>(null);
  const [error, setError] = useState<string | null>("");
  const [isPending, setIsPending] = useState<boolean>(false);

  const setupInterviewHandler = async () => {
    setIsPending(true);

    const input: CreateInterviewRequestDto = {};
    if (selectedInterviewer !== null) {
      input.interviewer_id = selectedInterviewer.id;
    }

    try {
      const response = await interviewServiceSingleton.createInterview(input);

      if (response.data) {
        router.push(`/interview/${response.data.id}`);
      }
    } catch (error) {
      setError((error as Error).message);
      setIsPending(false);
    }
  };

  return {
    selectedInterviewer,
    error,
    isPending,
    setupInterviewHandler,
    selectInterviewerHandler: setSelectedInterviewer,
  };
}
