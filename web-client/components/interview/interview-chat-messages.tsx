import { AlertTriangleIcon, Loader2Icon } from "lucide-react";
import { useEffect, useRef } from "react";
import { Skeleton } from "@/components/ui/skeleton";
import { useInterviewStore } from "@/store/interview-store";
import { InterviewMessageDto, InterviewMessageRole } from "@/dto/interview-message-dto";

export function InterviewChatMessages() {
  const messages = useInterviewStore((state) => state.messages);
  const isPending = useInterviewStore((state) => state.isPending);
  const isCompletionPending = useInterviewStore((state) => state.isCompletionPending);

  const previousMessagesLength = useRef(messages.length);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottomHandler = () => {
    if (!messagesEndRef.current) {
      return;
    }

    messagesEndRef.current.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    if (messages.length <= previousMessagesLength.current) {
      return;
    }
    previousMessagesLength.current = messages.length;
    scrollToBottomHandler();
  }, [messages]);

  useEffect(() => {
    if (isPending) {
      scrollToBottomHandler();
    }
  }, [isPending]);

  return (
    <div className="flex-1 overflow-y-auto p-6 space-y-6">
      {messages.map((message) => {
        if (message.role === InterviewMessageRole.User) {
          return <UserMessage key={message.id} message={message} />;
        } else if (message.role === InterviewMessageRole.Interviewer) {
          return <InterviewerMessage key={message.id} message={message} />;
        }

        return null;
      })}

      {isPending && !isCompletionPending && <PendingMessage />}

      {isCompletionPending && <CompletionPendingMessage />}

      <div ref={messagesEndRef} />
    </div>
  );
}

interface MessageComponentProps {
  message: InterviewMessageDto;
}

function InterviewerMessage({ message }: MessageComponentProps) {
  return (
    <div className="flex justify-start">
      <div className="max-w-[80%] bg-white border border-gray-100 shadow-sm rounded-2xl rounded-tl-sm p-4">
        <p className="text-gray-800">{message.content_text}</p>
        {message.tips_text && (
          <div className="mt-3 pt-3 border-t border-gray-200 text-sm flex items-start gap-2">
            <AlertTriangleIcon className="h-4 w-4 text-amber-500 flex-shrink-0 mt-0.5" />
            <span className="text-amber-700">{message.tips_text}</span>
          </div>
        )}
      </div>
    </div>
  );
}

function UserMessage({ message }: MessageComponentProps) {
  return (
    <div className="flex justify-end">
      <div className="max-w-[80%] bg-[#E12D39] text-white rounded-2xl rounded-tr-sm p-4">
        <p className="text-white">{message.content_text}</p>
      </div>
    </div>
  );
}

function PendingMessage() {
  return (
    <div className="flex justify-start">
      <div className="w-[40%] bg-white border border-gray-100 shadow-sm rounded-2xl rounded-tl-sm p-4">
        <Skeleton className="w-full h-4" />

        <div className="mt-3 pt-3 border-t border-gray-200 text-sm flex items-start gap-2">
          <Skeleton className="w-[70%] h-4" />
        </div>
      </div>
    </div>
  );
}

function CompletionPendingMessage() {
  return (
    <div className="flex justify-center">
      <div className="flex items-center gap-2 bg-white border border-gray-100 shadow-sm rounded-2xl p-4 w-[60%]">
        <div className="flex items-center justify-center">
          <Loader2Icon className="size-10 text-[#E12D39] animate-spin" />
        </div>
        <div>
          <p className="text-lg font-bold text-gray-800">Generating Results</p>
          <p className="text-sm text-gray-500">
            Please wait while we analyze your response. It can take up to 30 seconds.
          </p>
        </div>
      </div>
    </div>
  );
}
