"use client";

import { InterviewDto } from "@/dto/interview-dto";
import { Card } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { TextWithPopover } from "@/components/ui/text-with-popover";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { InterviewChatResult } from "@/components/interview/interview-chat-result";
import { InterviewChatMessages } from "@/components/interview/interview-chat-messages";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { Home, RotateCcw } from "lucide-react";

interface DashboardInterviewResultProps {
  interview: InterviewDto;
}

export function DashboardInterviewResult({ interview }: DashboardInterviewResultProps) {
  return (
    <div className="container mx-auto flex-1 flex flex-col md:flex-row gap-4">
      <div className="flex-1">
        <Card className="p-4 flex flex-col gap-2 items-center m-0">
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
          </div>
          <h2 className="font-bold text-xl text-[#0C3B5F]">{interview.interviewer.name}</h2>
          {interview.interviewer.description && (
            <TextWithPopover
              contentClassName="text-sm text-gray-500 text-center mt-1"
              content={interview.interviewer.description}
            />
          )}
          <div className="flex flex-col gap-2 w-full mt-2">
            <Button variant="outline" className="rounded-xl w-full" asChild>
              <Link href="/portal/interview">
                <RotateCcw className="mr-2 h-4 w-4" />
                Try Again
              </Link>
            </Button>
            <Link href="/portal">
              <Button className="w-full rounded-xl bg-[#0C3B5F] hover:bg-[#0a3050]">
                <Home className="mr-2 h-4 w-4" />
                Back to Main
              </Button>
            </Link>
          </div>
        </Card>
      </div>
      <Tabs defaultValue="result" className="w-full flex-[3] flex flex-col gap-4">
        <TabsList className="grid grid-cols-2 gap-1 bg-gray-200 rounded-lg">
          <TabsTrigger value="result" className="rounded-lg">
            Results
          </TabsTrigger>
          <TabsTrigger value="history" className="rounded-lg">
            Message History
          </TabsTrigger>
        </TabsList>

        <TabsContent value="result" className="flex-1 max-h-full overflow-hidden mt-0">
          <InterviewChatResult
            result={interview.result!}
            rootClassName="max-w-full max-h-full flex-1 overflow-hidden shadow-none max-h-full"
          />
        </TabsContent>

        <TabsContent value="history" asChild>
          <Card className="m-0">
            <InterviewChatMessages messages={interview.messages} />
          </Card>
        </TabsContent>
      </Tabs>
    </div>
  );
}
