"use client";

import { InterviewerDto } from "@/dto/interviewer-dto";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { CheckIcon, ChevronDownIcon } from "lucide-react";
import { useState } from "react";
import { RANDOM_INTERVIEWER } from "@/data/interviewer-constants";

interface DashboardInterviewSetupInterviewerSelectorProps {
  options: InterviewerDto[];
  selected: InterviewerDto | null;
  onSelect: (interviewer: InterviewerDto | null) => void;
}

export function DashboardInterviewSetupInterviewerSelector({
  options,
  selected,
  onSelect,
}: DashboardInterviewSetupInterviewerSelectorProps) {
  const [isOpen, setIsOpen] = useState(false);

  const normalizedInterviewer = selected ?? RANDOM_INTERVIEWER;

  const createSelectHandler = (interviewer: InterviewerDto | null) => () => {
    onSelect(interviewer);
    setIsOpen(false);
  };

  return (
    <div className="space-y-4">
      <Card className="border-2 border-[#E12D39]/20 rounded-xl overflow-hidden">
        <CardContent className="p-4">
          <div className="flex items-center gap-4">
            <Avatar className="h-16 w-16 border-2 border-[#E12D39]/20">
              <AvatarImage
                src={normalizedInterviewer.avatar_url}
                alt={normalizedInterviewer.name}
              />
              <AvatarFallback className="bg-[#E12D39] text-white text-lg">
                {normalizedInterviewer.name
                  .split(" ")
                  .map((n) => n[0])
                  .join("")}
              </AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <h4 className="font-bold text-[#0C3B5F] text-lg">{normalizedInterviewer.name}</h4>
              <p className="text-gray-600 text-sm mb-2">{normalizedInterviewer.description}</p>
            </div>
          </div>
        </CardContent>
      </Card>

      <DropdownMenu open={isOpen} onOpenChange={setIsOpen}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            className="w-full justify-between rounded-xl border-2 border-gray-200 hover:border-[#E12D39]/30 h-12"
          >
            <span className="text-gray-700">Choose a different interviewer</span>
            <ChevronDownIcon
              className={`h-4 w-4 transition-transform ${isOpen ? "rotate-180" : ""}`}
            />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="w-[400px] rounded-xl p-2" align="start">
          <DropdownMenuItem
            key={RANDOM_INTERVIEWER.id}
            className="p-0 rounded-lg"
            onClick={createSelectHandler(null)}
          >
            <div
              className={`w-full p-3 rounded-lg transition-colors ${
                selected === null
                  ? "bg-[#E12D39]/10 border border-[#E12D39]/20"
                  : "hover:bg-gray-50"
              }`}
            >
              <div className="flex items-center gap-3">
                <Avatar className="h-12 w-12 border border-gray-200">
                  <AvatarImage
                    src={RANDOM_INTERVIEWER.avatar_url || "/placeholder.svg"}
                    alt={RANDOM_INTERVIEWER.name}
                  />
                  <AvatarFallback className="bg-[#E12D39] text-white">
                    {RANDOM_INTERVIEWER.name
                      .split(" ")
                      .map((n) => n[0])
                      .join("")}
                  </AvatarFallback>
                </Avatar>
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2">
                    <h5 className="font-medium text-[#0C3B5F] truncate">
                      {RANDOM_INTERVIEWER.name}
                    </h5>
                    {selected === null && (
                      <CheckIcon className="h-4 w-4 text-[#E12D39] flex-shrink-0" />
                    )}
                  </div>
                  <p className="text-gray-600 text-sm line-clamp-2">
                    {RANDOM_INTERVIEWER.description}
                  </p>
                </div>
              </div>
            </div>
          </DropdownMenuItem>

          {options.map((interviewer) => (
            <DropdownMenuItem
              key={interviewer.id}
              className="p-0 rounded-lg"
              onClick={createSelectHandler(interviewer)}
            >
              <div
                className={`w-full p-3 rounded-lg transition-colors ${
                  normalizedInterviewer.id === interviewer.id
                    ? "bg-[#E12D39]/10 border border-[#E12D39]/20"
                    : "hover:bg-gray-50"
                }`}
              >
                <div className="flex items-center gap-3">
                  <Avatar className="h-12 w-12 border border-gray-200">
                    <AvatarImage
                      src={interviewer.avatar_url || "/placeholder.svg"}
                      alt={interviewer.name}
                    />
                    <AvatarFallback className="bg-[#E12D39] text-white">
                      {interviewer.name
                        .split(" ")
                        .map((n) => n[0])
                        .join("")}
                    </AvatarFallback>
                  </Avatar>
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-2">
                      <h5 className="font-medium text-[#0C3B5F] truncate">{interviewer.name}</h5>
                      {normalizedInterviewer.id === interviewer.id && (
                        <CheckIcon className="h-4 w-4 text-[#E12D39] flex-shrink-0" />
                      )}
                    </div>
                    <p className="text-gray-600 text-sm line-clamp-2">{interviewer.description}</p>
                  </div>
                </div>
              </div>
            </DropdownMenuItem>
          ))}
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
}
