"use client";

import { InterviewDto } from "@/dto/interview-dto";
import { useInterviewStore } from "@/store/interview-store";
import { ReactNode, useEffect } from "react";

interface InterviewChatStoreWrapperProps {
  children: ReactNode;
  interview: InterviewDto;
}

export function InterviewChatStoreWrapper({ children, interview }: InterviewChatStoreWrapperProps) {
  useEffect(() => {
    if (!interview || !!useInterviewStore.getState().interview) {
      return;
    }

    useInterviewStore.setState({
      interview,
    });

    return () => {
      // useInterviewStore.setState(useInterviewStore.getInitialState());
    };
  }, [interview]);

  return children;
}
