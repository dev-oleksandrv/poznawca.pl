import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { InfoIcon, Loader2Icon, XIcon } from "lucide-react";
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
import { Button } from "@/components/ui/button";
import { InterviewerDto } from "@/dto/interviewer-dto";
import { TextWithPopover } from "@/components/ui/text-with-popover";

interface InterviewChatSidebarProps {
  interviewer: InterviewerDto;

  isPending?: boolean;
  isCompletionAvailable?: boolean;
  isCompletionPending?: boolean;
  isCompleted?: boolean;

  onCompleteInterview?: () => void;
}

export function InterviewChatSidebar({
  interviewer,
  isCompletionPending = false,
  isCompletionAvailable = false,
  isCompleted = false,
  isPending = false,
  onCompleteInterview,
}: InterviewChatSidebarProps) {
  return (
    <div className="w-80 bg-white border-r border-gray-100 flex flex-col">
      <div className="px-6 py-4 border-b border-gray-100 flex items-center justify-between">
        <span className="text-xl font-bold text-[#E12D39]">poznawca.pl</span>
      </div>

      <div className="p-6 flex flex-col items-center">
        <div className="relative">
          <Avatar className="h-24 w-24 mb-4 border-4 border-white shadow-md">
            <AvatarImage src={interviewer.avatar_url} alt={interviewer.name} />
            <AvatarFallback className="bg-[#E12D39] text-white text-xl">
              {interviewer.name.slice(0, 2).toUpperCase()}
            </AvatarFallback>
          </Avatar>
        </div>
        <h2 className="font-bold text-xl text-[#0C3B5F]">{interviewer.name}</h2>
        {interviewer.description && (
          <TextWithPopover
            contentClassName="text-sm text-gray-500 text-center mt-1"
            content={interviewer.description}
          />
        )}
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

      {!isCompleted && (
        <div className="p-6 border-t border-gray-100">
          {!isCompletionAvailable && (
            <p className="text-xs text-gray-500 mb-2">
              You must send at least 5 messages before you can end the interview.
            </p>
          )}

          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button
                disabled={isCompletionPending || !isCompletionAvailable || isPending}
                variant="outline"
                className="w-full border-[#E12D39] text-[#E12D39] hover:bg-[#E12D39] hover:text-white"
              >
                {isCompletionPending ? (
                  <>
                    <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
                    Ending Interview...
                  </>
                ) : (
                  <>
                    <XIcon className="mr-2 h-4 w-4" />
                    End Interview
                  </>
                )}
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
                  onClick={onCompleteInterview}
                  className="bg-[#E12D39] hover:bg-[#c0252f] rounded-lg"
                >
                  End Interview
                </AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </div>
      )}
    </div>
  );
}
