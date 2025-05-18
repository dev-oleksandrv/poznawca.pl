"use client";

import { InterviewDto } from "@/dto/interview-dto";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { AlertTriangleIcon, InfoIcon, Loader2Icon, SendIcon, XIcon } from "lucide-react";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Textarea } from "@/components/ui/textarea";
import { FormEvent, KeyboardEvent, useEffect, useRef, useState } from "react";
import { WSManager, WSManagerEventType } from "@/ws/manager";
import { SystemMessageSentEvent } from "@/ws/events";

interface InterviewChatProps {
  interview: InterviewDto;
}

export function InterviewChat({ interview }: InterviewChatProps) {
  const [messages, setMessages] = useState<any[]>([]);
  const [inputValue, setInputValue] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const socketRef = useRef<WSManager>(null);

  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottomHandler = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    scrollToBottomHandler();
  }, [messages]);

  const sendMessageHandler = () => {
    if (!socketRef.current) {
      return;
    }

    setIsSubmitting(true);

    socketRef.current.sendMessage(inputValue.trim());

    setMessages((prev) => [
      ...prev,
      {
        id: Date.now().toString(),
        sender: "user",
        content: inputValue.trim(),
      },
    ]);

    setIsSubmitting(false);
    setInputValue("");
  };

  const formSubmitHandler = async (e: FormEvent) => {
    e.preventDefault();

    sendMessageHandler();
  };

  const keyDownHandler = (e: KeyboardEvent) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();

      sendMessageHandler();
    }
  };

  useEffect(() => {
    if (socketRef.current !== null) {
      return;
    }

    socketRef.current = new WSManager(interview);

    socketRef.current.on(WSManagerEventType.SystemMessageSent, (event: SystemMessageSentEvent) => {
      setMessages((prev) => [
        ...prev,
        {
          id: Date.now().toString(),
          sender: "system",
          content: event.content_text,
          feedback: event.tips_text,
        },
      ]);
    });

    return () => {
      if (socketRef.current) {
        socketRef.current.close();
      }
    };
  }, [interview]);

  if (!interview.interviewer) {
    return null;
  }

  return (
    <div className="flex h-screen bg-[#f8f9fa]">
      <div className="w-80 bg-white border-r border-gray-100 flex flex-col">
        <div className="px-6 py-4 border-b border-gray-100 flex items-center justify-between">
          <span className="text-xl font-bold text-[#E12D39]">poznawca.pl</span>
        </div>

        <div className="p-6 flex flex-col items-center">
          <div className="relative">
            <Avatar className="h-24 w-24 mb-4 border-4 border-white shadow-md">
              <AvatarImage
                src={interview.interviewer.avatar_url}
                alt={interview.interviewer.name}
              />
              <AvatarFallback className="bg-[#E12D39] text-white text-xl">
                {interview.interviewer.name.slice(0, 2).toUpperCase()}
              </AvatarFallback>
            </Avatar>
            <div className="absolute bottom-3 right-0 w-5 h-5 bg-green-500 rounded-full border-2 border-white"></div>
          </div>
          <h2 className="font-bold text-xl text-[#0C3B5F]">{interview.interviewer.name}</h2>
          <p className="text-sm text-gray-500 text-center mt-1">
            {/*{interview.interviewer.description}*/}
            Lorem ipsum dolor sit amet, consectetur adipisicing elit.
          </p>
        </div>

        <div className="flex-1 p-6 overflow-y-auto">
          <div className="bg-[#f8f9fa] rounded-xl p-4 mb-4">
            <h3 className="font-medium text-[#0C3B5F] flex items-center gap-2 mb-2">
              <InfoIcon className="h-4 w-4 text-[#E12D39]" />
              Interview Tips
            </h3>
            <ul className="text-sm text-gray-600 space-y-2">
              <li>• Answer questions clearly and concisely</li>
              <li>• Maintain a formal tone throughout</li>
              <li>• Provide specific details when asked</li>
              <li>• Be honest in your responses</li>
            </ul>
          </div>
        </div>

        <div className="p-6 border-t border-gray-100">
          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button
                variant="outline"
                className="w-full border-[#E12D39] text-[#E12D39] hover:bg-[#E12D39] hover:text-white"
              >
                <XIcon className="mr-2 h-4 w-4" />
                End Interview
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent className="rounded-xl">
              <AlertDialogHeader>
                <AlertDialogTitle>End Interview?</AlertDialogTitle>
                <AlertDialogDescription>
                  Are you sure you want to end this interview? Your progress will be saved, but you
                  will use 5 energy points.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel className="rounded-lg">Cancel</AlertDialogCancel>
                <AlertDialogAction
                  onClick={() => alert("not implemented yet")}
                  className="bg-[#E12D39] hover:bg-[#c0252f] rounded-lg"
                >
                  End Interview
                </AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </div>
      </div>

      <div className="flex-1 flex flex-col">
        <div className="flex-1 overflow-y-auto p-6 space-y-6">
          {messages.map((message) => (
            <div
              key={message.id}
              className={`flex ${message.sender === "user" ? "justify-end" : "justify-start"}`}
            >
              <div
                className={`max-w-[80%] ${
                  message.sender === "user"
                    ? "bg-[#E12D39] text-white rounded-2xl rounded-tr-sm"
                    : "bg-white border border-gray-100 shadow-sm rounded-2xl rounded-tl-sm"
                } p-4`}
              >
                <p className={message.sender === "user" ? "text-white" : "text-gray-800"}>
                  {message.content}
                </p>
                {message.feedback && (
                  <div className="mt-3 pt-3 border-t border-gray-200 text-sm flex items-start gap-2">
                    <AlertTriangleIcon className="h-4 w-4 text-amber-500 flex-shrink-0 mt-0.5" />
                    <span className="text-amber-700">{message.feedback}</span>
                  </div>
                )}
              </div>
            </div>
          ))}
          <div ref={messagesEndRef} />
        </div>

        {/* Input area */}
        <div className="border-t border-gray-100 bg-white p-4">
          <form onSubmit={formSubmitHandler} className="flex items-end gap-2">
            <Textarea
              value={inputValue}
              onChange={(e) => setInputValue(e.target.value)}
              onKeyDown={keyDownHandler}
              placeholder="Type your answer here..."
              className="flex-1 min-h-[80px] resize-none rounded-xl border-gray-200 focus:border-[#E12D39] focus:ring-[#E12D39]"
              disabled={isSubmitting}
            />
            <Button
              type="submit"
              className={`bg-[#E12D39] hover:bg-[#c0252f] h-12 w-12 rounded-xl flex items-center justify-center ${
                isSubmitting || !inputValue.trim() ? "opacity-50 cursor-not-allowed" : ""
              }`}
              disabled={isSubmitting || !inputValue.trim()}
            >
              {isSubmitting ? (
                <Loader2Icon className="h-5 w-5 animate-spin text-white" />
              ) : (
                <SendIcon className="h-5 w-5" />
              )}
            </Button>
          </form>
        </div>
      </div>
    </div>
  );
}
