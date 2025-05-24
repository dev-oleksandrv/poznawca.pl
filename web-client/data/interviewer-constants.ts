import { InterviewerDto } from "@/dto/interviewer-dto";

export const RANDOM_INTERVIEWER: InterviewerDto = {
  id: "random-interviewer",
  name: "Random Interviewer",
  avatar_url: "",
  description: "Let us choose an interviewer for you!",
  description_translation_key: "interviewer_random",
};

export const INTERVIEW_MIN_MESSAGES_BEFORE_COMPLETION = 5;
