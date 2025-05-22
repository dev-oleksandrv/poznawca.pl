import { Button } from "@/components/ui/button";
import { ChevronLeftIcon, ChevronRightIcon } from "lucide-react";

export function DashboardRecentResultsSection() {
  return (
    <section>
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl font-bold text-[#0C3B5F]">Recent Results</h2>
        <div className="flex items-center gap-2">
          <Button variant="outline" size="icon" className="rounded-full h-8 w-8">
            <ChevronLeftIcon className="h-4 w-4" />
          </Button>
          <Button variant="outline" size="icon" className="rounded-full h-8 w-8">
            <ChevronRightIcon className="h-4 w-4" />
          </Button>
        </div>
      </div>
      {/*<RecentResults />*/}
    </section>
  );
}
