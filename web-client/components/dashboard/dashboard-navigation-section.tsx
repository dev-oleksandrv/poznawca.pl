import { Card } from "@/components/ui/card";
import { ChevronRightIcon, FileTextIcon, MessageSquareIcon, ZapIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export function DashboardNavigationSection() {
  return (
    <section className="mb-8">
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Card className="overflow-hidden rounded-2xl border-0 shadow-md hover:shadow-lg transition-all bg-[#E12D39]/10 p-6 flex flex-col">
          <div className="flex-1">
            <div className="bg-[#E12D39] w-12 h-12 rounded-xl flex items-center justify-center mb-4">
              <MessageSquareIcon className="h-6 w-6 text-white" />
            </div>
            <h3 className="text-xl font-bold text-[#E12D39] mb-2">Mock Interview</h3>
            <p className="text-gray-600 mb-4">
              Practice your interview skills with a simulated Karta Stalego Pobytu interview.
            </p>
          </div>

          <div className="flex justify-between items-center">
            <div className="flex items-center text-sm text-gray-500">
              <ZapIcon className="h-4 w-4 mr-1 text-[#E12D39]" />
              <span>Costs 5 energy points</span>
            </div>

            <Button className="bg-[#E12D39] hover:bg-[#c0252f]" asChild>
              <Link href="/portal/interview">
                <span>Start</span>
                <ChevronRightIcon className="h-4 w-4 ml-1" />
              </Link>
            </Button>
          </div>
        </Card>

        <Card className="overflow-hidden rounded-2xl border-0 shadow-md hover:shadow-lg transition-all bg-[#0C3B5F]/10 p-6 flex flex-col">
          <div className="flex-1">
            <div className="bg-[#0C3B5F] w-12 h-12 rounded-xl flex items-center justify-center mb-4">
              <FileTextIcon className="h-6 w-6 text-white" />
            </div>
            <h3 className="text-xl font-bold text-[#0C3B5F] mb-2">Knowledge Tests</h3>
            <p className="text-gray-600 mb-4">
              Test your knowledge about Polish history, culture, and society with interactive
              quizzes.
            </p>
          </div>

          <div className="flex justify-between items-center">
            <div className="flex items-center text-sm text-gray-500">
              <ZapIcon className="h-4 w-4 mr-1 text-[#0C3B5F]" />
              <span>Costs 1 energy points</span>
            </div>

            <Button className="bg-[#0C3B5F] hover:bg-[#0a3050]" asChild>
              <Link href="/portal/tests">
                <span>Start</span>
                <ChevronRightIcon className="h-4 w-4 ml-1" />
              </Link>
            </Button>
          </div>
        </Card>
      </div>
    </section>
  );
}
