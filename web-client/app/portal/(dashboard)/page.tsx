import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import {
  AwardIcon,
  BarChart3Icon,
  ChevronLeftIcon,
  ChevronRightIcon,
  CrownIcon,
  FileTextIcon,
  MessageSquareIcon,
  ZapIcon,
} from "lucide-react";
import Link from "next/link";

export default function PortalDashboardPage() {
  return (
    <>
      <section className="mb-8">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <Card className="overflow-hidden rounded-2xl border-0 shadow-md hover:shadow-lg transition-all bg-[#E12D39]/10 p-6 flex flex-col">
            <div className="flex-1">
              <div className="bg-[#E12D39] w-12 h-12 rounded-xl flex items-center justify-center mb-4">
                <MessageSquareIcon className="h-6 w-6 text-white" />
              </div>
              <h3 className="text-xl font-bold text-[#0C3B5F] mb-2">Mock Interview</h3>
              <p className="text-gray-600 mb-4">
                Practice your interview skills with a simulated Karta Stalego Pobytu interview.
              </p>
              <div className="flex items-center text-sm text-gray-500">
                <ZapIcon className="h-4 w-4 mr-1 text-[#E12D39]" />
                <span>Costs 5 energy points</span>
              </div>
            </div>

            <div className="flex justify-end">
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
              <div className="flex items-center text-sm text-gray-500">
                <ZapIcon className="h-4 w-4 mr-1 text-[#0C3B5F]" />
                <span>Costs 1 energy points</span>
              </div>
            </div>

            <div className="flex justify-end">
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

      <section className="mb-8">
        <Card className="bg-gradient-to-r from-[#0C3B5F] to-[#1a5a8a] border-0 overflow-hidden rounded-2xl shadow-lg p-0">
          <CardContent className="p-0">
            <div className="flex flex-col md:flex-row items-center">
              <div className="p-6 md:p-8 flex-1">
                <div className="flex items-center gap-3 mb-4">
                  <div className="bg-[#FFD700] p-2 rounded-full">
                    <CrownIcon className="h-6 w-6 text-[#0C3B5F]" />
                  </div>
                  <h3 className="font-bold text-xl text-white">Unlimited Access with Premium</h3>
                </div>
                <p className="text-gray-200 mb-6">
                  Get unlimited energy, detailed feedback, and exclusive practice materials.
                </p>
                <div className="flex items-center gap-4">
                  <div className="text-white">
                    <p className="text-sm opacity-75">Only</p>
                    <p className="font-bold text-2xl">$2.99/month</p>
                  </div>
                  <Button className="bg-[#FFD700] text-[#0C3B5F] hover:bg-[#e6c200] rounded-xl">
                    Get Premium
                  </Button>
                </div>
              </div>
              <div className="hidden md:block w-1/3 h-full p-8">
                <div className="bg-white/10 p-4 rounded-xl h-full flex flex-col justify-center">
                  <h4 className="font-bold text-white mb-4">Premium Benefits:</h4>
                  <ul className="space-y-3">
                    <li className="flex items-center gap-2 text-white">
                      <ZapIcon className="h-4 w-4 text-[#FFD700]" />
                      <span>Unlimited energy</span>
                    </li>
                    <li className="flex items-center gap-2 text-white">
                      <AwardIcon className="h-4 w-4 text-[#FFD700]" />
                      <span>Detailed feedback</span>
                    </li>
                    <li className="flex items-center gap-2 text-white">
                      <BarChart3Icon className="h-4 w-4 text-[#FFD700]" />
                      <span>Progress tracking</span>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>
      </section>

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
    </>
  );
}
