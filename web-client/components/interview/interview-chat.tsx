"use client";

import { useInterviewStore } from "@/store/interview-store";
import { useEffect, useMemo, useRef } from "react";
import { InterviewWsManager } from "@/ws/interview-ws-manager";
import { InterviewChatSidebar } from "@/components/interview/interview-chat-sidebar";
import {
  InterviewerMessageSentEvent,
  InterviewEventType,
  ResultsSentEvent,
  UserMessageSentEvent,
} from "@/ws/interview-events";
import { InterviewChatMessages } from "@/components/interview/interview-chat-messages";
import { InterviewChatInput } from "@/components/interview/interview-chat-input";
import { InterviewMessageRole } from "@/dto/interview-message-dto";
import { INTERVIEW_MIN_MESSAGES_BEFORE_COMPLETION } from "@/data/interviewer-constants";
import { InterviewChatResult } from "@/components/interview/interview-chat-result";
import {
  InterviewChatSoundController,
  InterviewChatSoundControllerRef,
} from "@/components/interview/interview-chat-sound-controller";

export function InterviewChat() {
  const interview = useInterviewStore((state) => state.interview);
  const isPending = useInterviewStore((state) => state.isPending);
  const setIsPending = useInterviewStore((state) => state.setIsPending);
  const setResult = useInterviewStore((state) => state.setResult);
  const messages = useInterviewStore((state) => state.messages);
  const addMessage = useInterviewStore((state) => state.addMessage);
  const userInput = useInterviewStore((state) => state.userInput);
  const setUserInput = useInterviewStore((state) => state.setUserInput);
  const isCompletionPending = useInterviewStore((state) => state.isCompletionPending);
  const interviewResult = useInterviewStore((state) => state.result);
  const setIsCompletionPending = useInterviewStore((state) => state.setIsCompletionPending);

  const wsManagerRef = useRef<InterviewWsManager>(null);
  const soundControllerRef = useRef<InterviewChatSoundControllerRef>(null);

  const sendUserMessageHandler = () => {
    if (!interview || !userInput.trim() || !wsManagerRef.current) {
      return;
    }

    addMessage({
      id: Date.now().toString(),
      role: InterviewMessageRole.User,
      content_text: userInput.trim(),
      created_at: new Date().toISOString(),
    });

    wsManagerRef.current.sendEvent({
      type: InterviewEventType.UserMessageSent,
      content: userInput.trim(),
    } as UserMessageSentEvent);
  };

  const sendUserCompleteInterviewHandler = () => {
    if (!interview || !wsManagerRef.current) {
      return;
    }

    wsManagerRef.current.sendEvent({
      type: InterviewEventType.UserCompleteInterview,
    });
    setIsPending(true);
  };

  useEffect(() => {
    if (!interview || wsManagerRef.current !== null) {
      return;
    }

    const interviewWsManager = new InterviewWsManager(interview.id);

    interviewWsManager
      .subscribe(InterviewEventType.InterviewerMessagePending, () => setIsPending(true))
      .subscribe(InterviewEventType.ResultsPending, () => setIsCompletionPending(true))
      .subscribe(
        InterviewEventType.InterviewerMessageSent,
        (event: InterviewerMessageSentEvent) => {
          addMessage(event.data);
          setUserInput("");
          if (!event.data.is_last_message) {
            setIsPending(false);
          }
          if (soundControllerRef.current) {
            soundControllerRef.current.playNewMessageSound();
          }
        },
      )
      .subscribe(InterviewEventType.ResultsSent, (event: ResultsSentEvent) => {
        setResult(event.data);
        setIsCompletionPending(false);
        if (soundControllerRef.current) {
          soundControllerRef.current.playInterviewCompletedSound();
        }
      });

    wsManagerRef.current = interviewWsManager;
  }, [interview]);

  const isCompletionAvailable = useMemo(() => {
    const userMessages = messages.filter((msg) => msg.role === InterviewMessageRole.User);

    return userMessages.length > INTERVIEW_MIN_MESSAGES_BEFORE_COMPLETION;
  }, [messages]);

  if (!interview) {
    return <div className="flex items-center justify-center min-h-screen">Loading...</div>;
  }

  return (
    <div className="flex h-screen bg-[#f8f9fa] relative">
      <InterviewChatSidebar
        interviewer={interview.interviewer}
        isPending={isPending}
        isCompletionPending={isCompletionPending}
        isCompletionAvailable={isCompletionAvailable}
        isCompleted={!!interviewResult}
        onCompleteInterview={sendUserCompleteInterviewHandler}
      />

      <div className="flex-1 flex flex-col">
        <InterviewChatMessages />

        {!isCompletionPending && !interviewResult && (
          <InterviewChatInput onSubmit={sendUserMessageHandler} />
        )}

        {!!interviewResult && (
          <div className="absolute top-0 left-0 w-full h-full bg-opacity-60 bg-black flex items-center justify-center z-10">
            <InterviewChatResult />
          </div>
        )}
      </div>

      <InterviewChatSoundController ref={soundControllerRef} />
    </div>
  );
}
