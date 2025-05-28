import { InterviewDto, InterviewStatus } from "@/dto/interview-dto";
import {
  AlertTriangle,
  CalendarDaysIcon,
  CheckCircle,
  MessageSquareIcon,
  XCircle,
} from "lucide-react";
import { Badge } from "@/components/ui/badge";
import { useMemo } from "react";
import { cn } from "@/lib/utils";
import { useRouter } from "next/navigation";

interface DashboardRecentResultsItemProps {
  interview: InterviewDto;
}

export function DashboardRecentResultsItem({ interview }: DashboardRecentResultsItemProps) {
  const router = useRouter();

  const getScoreColor = (score: number) => {
    if (score >= 80) return "text-green-600";
    if (score >= 60) return "text-amber-600";
    return "text-red-600";
  };

  const getScoreBg = (score: number) => {
    if (score >= 80) return "bg-green-100";
    if (score >= 60) return "bg-amber-100";
    return "bg-red-100";
  };

  const getScoreIcon = (score: number) => {
    if (score >= 80) return <CheckCircle className={`size-3 ${getScoreColor(score)}`} />;
    if (score >= 60) return <AlertTriangle className={`size-3 ${getScoreColor(score)}`} />;
    return <XCircle className={`size-3 ${getScoreColor(score)}`} />;
  };

  const statusBadge = useMemo(() => {
    if (interview.status === InterviewStatus.Abandoned) {
      return (
        <Badge className="bg-red-100 text-red-800 hover:bg-red-100 flex items-center gap-1 rounded-full">
          <XCircle className="h-3 w-3" />
          <span>Abandoned</span>
        </Badge>
      );
    }

    if (interview.status === InterviewStatus.Completed && interview.result) {
      return (
        <Badge
          className={`flex items-center gap-1 ${getScoreBg(interview.result.total_score)} ${getScoreColor(interview.result.total_score)} rounded-full`}
        >
          {getScoreIcon(interview.result.total_score)}
          <span>{interview.result.total_score}%</span>
        </Badge>
      );
    }
  }, [interview]);

  const clickHandler = () => {
    if (interview.status === InterviewStatus.Completed && interview.result) {
      router.push(`/portal/interview/${interview.id}`);
    }
  };

  return (
    <div
      onClick={clickHandler}
      className={cn("flex gap-2 p-4 gap-4", {
        "cursor-pointer": interview.status === InterviewStatus.Completed && interview.result,
      })}
    >
      <div className="size-12 rounded-xl flex items-center justify-center bg-[#E12D39]">
        <MessageSquareIcon className="size-6 text-white" />
      </div>

      <div className="flex-1 flex flex-col">
        <h4 className="font-medium text-[#0C3B5F]">Interview with {interview.interviewer.name}</h4>
        <div className="flex items-center text-xs text-gray-500">
          <CalendarDaysIcon className="h-3 w-3 mr-1" />
          <span>{interview.updated_at}</span>
        </div>
      </div>

      <div className="flex flex-col justify-between gap-2">{statusBadge}</div>
    </div>
  );
}
