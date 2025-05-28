"use client";

import { Button } from "@/components/ui/button";
import { ChevronLeftIcon, ChevronRightIcon, Loader2Icon, XCircleIcon } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import { useInterviewsListStore } from "@/store/interviews-list-store";
import { Card } from "@/components/ui/card";
import { interviewServiceSingleton } from "@/services/interview-service";
import { DashboardRecentResultsItem } from "@/components/dashboard/dashboard-recent-results-item";

export function DashboardRecentResultsSection() {
  const [currentPage, setCurrentPage] = useState(1);

  const interviewsListStore = useInterviewsListStore();

  const totalPages = Math.ceil(interviewsListStore.interviews.length / 5);

  const paginatedResults = useMemo(() => {
    const startIndex = (currentPage - 1) * 5;
    return interviewsListStore.interviews.slice(startIndex, startIndex + 5);
  }, [interviewsListStore.interviews, currentPage]);

  const nextPageHandler = () => {
    if (currentPage < totalPages) {
      setCurrentPage((prev) => prev + 1);
    }
  };

  const prevPageHandler = () => {
    if (currentPage > 1) {
      setCurrentPage((prev) => prev - 1);
    }
  };

  useEffect(() => {
    interviewsListStore.setIsPending(true);

    interviewServiceSingleton
      .getList()
      .then((response) => interviewsListStore.setInterviews(response.data))
      .catch((error) => interviewsListStore.setError(error?.message))
      .finally(() => interviewsListStore.setIsPending(false));
  }, []);

  return (
    <section>
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-bold text-[#0C3B5F]">Recent Results</h2>
        {interviewsListStore.interviews.length > 0 && (
          <div className="flex items-center gap-2">
            <Button
              variant="outline"
              size="icon"
              className="rounded-full h-8 w-8"
              disabled={currentPage === 1}
              onClick={prevPageHandler}
            >
              <ChevronLeftIcon className="h-4 w-4" />
            </Button>
            <Button
              variant="outline"
              size="icon"
              className="rounded-full h-8 w-8"
              disabled={currentPage >= totalPages}
              onClick={nextPageHandler}
            >
              <ChevronRightIcon className="h-4 w-4" />
            </Button>
          </div>
        )}
      </div>

      {interviewsListStore.interviews.length > 0 ? (
        <Card>
          {paginatedResults.map((interview) => (
            <DashboardRecentResultsItem key={interview.id} interview={interview} />
          ))}
        </Card>
      ) : (
        <>
          {interviewsListStore.isPending && (
            <Card className="p-4 flex flex-col items-center">
              <Loader2Icon className="size-8 text-[#E12D39] animate-spin" />

              <p className="text-sm text-gray-600 mt-2">Loading recent results...</p>
            </Card>
          )}

          {interviewsListStore.error && (
            <Card className="p-4 flex flex-col items-center">
              <XCircleIcon className="size-8 text-[#E12D39]" />

              <p className="text-sm text-gray-600 mt-2">{interviewsListStore.error}</p>
            </Card>
          )}
        </>
      )}
    </section>
  );
}
