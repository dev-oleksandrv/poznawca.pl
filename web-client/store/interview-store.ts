import { create } from "zustand";
import { InterviewDto } from "@/dto/interview-dto";
import { InterviewMessageDto } from "@/dto/interview-message-dto";
import { InterviewResultDto } from "@/dto/interview-result-dto";

export interface InterviewStore {
  interview: InterviewDto | null;
  messages: InterviewMessageDto[];
  result: InterviewResultDto | null;
  userInput: string;
  isPending: boolean;
  isCompleted: boolean;

  addMessage: (message: any) => void;
  setInterview: (interview: any | null) => void;
  setResult: (result: any | null) => void;
  setUserInput: (userInput: string) => void;
  setIsPending: (isPending: boolean) => void;
  setIsCompleted: (isCompleted: boolean) => void;
}

export const useInterviewStore = create<InterviewStore>((set) => ({
  interview: null,
  messages: [],
  result: null,
  userInput: "",
  isPending: false,
  isCompleted: false,

  addMessage: (message) => set((state) => ({ messages: [...state.messages, message] })),
  setInterview: (interview) => set({ interview }),
  setResult: (result) => set({ result }),
  setUserInput: (userInput) => set({ userInput }),
  setIsPending: (isPending) => set({ isPending }),
  setIsCompleted: (isCompleted) => set({ isCompleted }),
}));
