"use client";

import { InterviewChatSidebar } from "@/components/interview/interview-chat-sidebar";
import { InterviewDto } from "@/dto/interview-dto";
import { InterviewChatMessages } from "@/components/interview/interview-chat-messages";
import { Card, CardContent, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { Home, RotateCcw } from "lucide-react";

export interface InterviewChatAbandonedProps {
  interview: InterviewDto;
}

export function InterviewChatAbandoned({ interview }: InterviewChatAbandonedProps) {
  return (
    <div className="flex h-screen bg-[#f8f9fa] relative">
      <InterviewChatSidebar isCompleted isCompletionAvailable interviewer={interview.interviewer} />

      <div className="flex-1 flex flex-col">
        <div className="p-6 pb-0">
          <Card className="p-4 mb-6">
            <CardContent className="p-0">
              <div className="text-lg font-bold text-red-600">
                This interview has been abandoned.
              </div>
              <div className="text-gray-500 mt-1">
                You can no longer interact with this interview.
              </div>

              <div className="flex flex-wrap gap-2 items-center mt-4">
                <Button className="rounded-xl bg-[#0C3B5F] hover:bg-[#0a3050]" asChild>
                  <Link href="/portal">
                    <Home className="mr-2 h-4 w-4" />
                    Back to Main
                  </Link>
                </Button>
                <Button variant="outline" className="rounded-xl" asChild>
                  <Link href="/portal/interview">
                    <RotateCcw className="mr-2 h-4 w-4" />
                    Start new Interview
                  </Link>
                </Button>
              </div>
            </CardContent>
          </Card>

          <h3 className="text-lg font-bold text-gray-800">Conversation history:</h3>
        </div>

        <InterviewChatMessages messages={interview.messages} />
      </div>
    </div>
  );
}
