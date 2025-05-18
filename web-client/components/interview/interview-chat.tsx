"use client";

import { InterviewDto } from "@/dto/interview-dto";
import { useState } from "react";
import { InterviewChatSidebar } from "@/components/interview/interview-chat-sidebar";
import { InterviewChatInput } from "@/components/interview/interview-chat-input";
import { InterviewChatMessages } from "@/components/interview/interview-chat-messages";
import { useInterviewChat } from "@/hooks/use-interview-chat";

interface InterviewChatProps {
  interview: InterviewDto;
}

export function InterviewChat({ interview }: InterviewChatProps) {
  const { send, messages, isPending } = useInterviewChat(interview);

  const [inputValue, setInputValue] = useState("");

  const sendMessageHandler = () => {
    send(inputValue);

    setInputValue("");
  };

  return (
    <div className="flex h-screen bg-[#f8f9fa]">
      <InterviewChatSidebar interviewer={interview.interviewer} />

      <div className="flex-1 flex flex-col">
        <InterviewChatMessages messages={messages} isPending={isPending} />

        <InterviewChatInput
          value={inputValue}
          isSubmitting={isPending}
          onInputChange={setInputValue}
          onSubmit={sendMessageHandler}
        />
      </div>
    </div>
  );
}
