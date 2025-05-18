import { InterviewDto, InterviewMessageDto, InterviewMessageRole } from "@/dto/interview-dto";
import { useEffect, useRef, useState } from "react";
import { WSManager, WSManagerEventType } from "@/ws/manager";
import { SystemMessageSentEvent } from "@/ws/events";

function createUserMessage(content: string) {
  return {
    id: Date.now().toString(),
    content_text: content,
    content_translation_text: "",
    tips_text: "",
    role: InterviewMessageRole.User,
  };
}

export function useInterviewChat(interview: InterviewDto) {
  const [isPending, setIsPending] = useState(false);
  const [messages, setMessages] = useState<InterviewMessageDto[]>([]);

  const socketRef = useRef<WSManager>(null);

  const sendMessageHandler = (inputValue: string) => {
    const trimmedInputValue = inputValue.trim();

    if (!socketRef.current || isPending || !trimmedInputValue) {
      return;
    }

    socketRef.current.sendMessage(inputValue.trim());

    setIsPending(true);
    setMessages((prev) => [...prev, createUserMessage(trimmedInputValue)]);
  };

  useEffect(() => {
    if (socketRef.current !== null || !interview) {
      return;
    }

    socketRef.current = new WSManager(interview);

    socketRef.current.on(WSManagerEventType.SystemMessageSent, (event: SystemMessageSentEvent) => {
      setMessages((prev) => [...prev, event.details]);
      setIsPending(false);
    });

    socketRef.current.on(
      WSManagerEventType.SystemMessagePending,
      (event: SystemMessageSentEvent) => {
        setIsPending(true);
      },
    );

    // return () => {
    //   if (socketRef.current) {
    //     socketRef.current.close();
    //   }
    // };
  }, [interview]);

  return {
    messages,
    isPending,
    send: sendMessageHandler,
  };
}
