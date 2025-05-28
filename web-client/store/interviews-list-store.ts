import { InterviewDto } from "@/dto/interview-dto";
import { create } from "zustand";

export interface InterviewsListState {
  isPending: boolean;
  error: string | null;
  interviews: InterviewDto[];
  setIsPending: (isPending: boolean) => void;
  setError: (error: string | null) => void;
  setInterviews: (interviews: InterviewDto[]) => void;
}

export const useInterviewsListStore = create<InterviewsListState>((set) => ({
  isPending: false,
  error: null,
  interviews: [],

  setIsPending: (isPending: boolean) => set({ isPending }),
  setError: (error: string | null) => set({ error }),
  setInterviews: (interviews: InterviewDto[]) => set({ interviews }),
}));
