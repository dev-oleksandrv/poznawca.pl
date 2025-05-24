import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Loader2Icon, SendIcon } from "lucide-react";
import { ChangeEvent, FormEvent, KeyboardEvent } from "react";
import { useInterviewStore } from "@/store/interview-store";

interface InterviewChatInputProps {
  onSubmit: () => void;
}

export function InterviewChatInput({ onSubmit }: InterviewChatInputProps) {
  const userInput = useInterviewStore((state) => state.userInput);
  const isPending = useInterviewStore((state) => state.isPending);
  const setUserInput = useInterviewStore((state) => state.setUserInput);

  const isReadyForSubmission = !isPending && userInput.trim().length > 0;

  const submitHandler = () => {
    if (!isReadyForSubmission) {
      return;
    }
    onSubmit();
  };

  const formSubmitHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    submitHandler();
  };

  const inputChangeHandler = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setUserInput(e.target.value);
  };

  const keyDownHandler = (e: KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      submitHandler();
    }
  };

  return (
    <div className="border-t border-gray-100 bg-white p-4">
      <form onSubmit={formSubmitHandler} className="flex items-end gap-2">
        <Textarea
          value={userInput}
          onChange={inputChangeHandler}
          onKeyDown={keyDownHandler}
          placeholder="Type your answer here..."
          className="flex-1 min-h-[80px] resize-none rounded-xl border-gray-200 focus:border-[#E12D39] focus:ring-[#E12D39]"
          disabled={isPending}
        />
        <Button
          type="submit"
          className={`bg-[#E12D39] hover:bg-[#c0252f] h-12 w-12 rounded-xl flex items-center justify-center ${
            !isReadyForSubmission ? "opacity-50 cursor-not-allowed" : ""
          }`}
          disabled={!isReadyForSubmission}
        >
          {isPending ? (
            <Loader2Icon className="h-5 w-5 animate-spin text-white" />
          ) : (
            <SendIcon className="h-5 w-5" />
          )}
        </Button>
      </form>
    </div>
  );
}
