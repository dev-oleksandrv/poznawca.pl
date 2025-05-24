"use client";

import { InterviewerDto } from "@/dto/interviewer-dto";
import { create } from "zustand/react";
import { createJSONStorage, persist } from "zustand/middleware";

export interface InterviewSetupStore {
  interviewer: InterviewerDto | null;
  setInterviewer: (interviewer: InterviewerDto | null) => void;
}

export const useInterviewSetupStore = create<InterviewSetupStore>()(
  persist(
    (set) => ({
      interviewer: null,
      setInterviewer: (interviewer: InterviewerDto | null) => set({ interviewer }),
    }),
    {
      name: "interview-setup-store",
      partialize: (state) => ({ interviewer: state.interviewer }),
      storage: createJSONStorage(() => localStorage),
    },
  ),
);
