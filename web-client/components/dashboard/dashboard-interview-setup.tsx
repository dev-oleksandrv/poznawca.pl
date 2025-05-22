"use client";

import { Alert, AlertTitle, AlertDescription } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { ArrowLeftIcon, ArrowRightIcon, Loader2Icon, MessageSquareIcon } from "lucide-react";
import Link from "next/link";
import { useSetupInterviewHandler } from "@/hooks/use-setup-interview-handler";

export function DashboardInterviewSetup() {
  const { error, isPending, setupInterviewHandler } = useSetupInterviewHandler();

  return (
    <Card className="flex flex-col flex-1 w-full bg-white rounded-2xl shadow-xl overflow-hidden border-0 p-0">
      <div className="flex flex-col md:flex-row h-full flex-1">
        <div className="w-full md:w-5/12 bg-gradient-to-br from-[#E12D39] to-[#c0252f] p-8 flex flex-col justify-between text-white">
          <div>
            <div className="flex-none container mx-auto pb-2 flex items-center mb-4">
              <Link href="/portal" className="flex items-center gap-2 text-sm ">
                <ArrowLeftIcon className="size-4" />
                <span>Back</span>
              </Link>
            </div>

            <div className="bg-white/20 w-12 h-12 rounded-xl flex items-center justify-center mb-6">
              <MessageSquareIcon className="h-6 w-6 text-white" />
            </div>
            <h1 className="text-3xl font-bold mb-4">Mock Interview</h1>
            <p className="opacity-90 mb-6">
              Practice your Karta Stalego Pobytu interview in a realistic simulation.
            </p>
          </div>

          <div className="space-y-4">
            <div className="bg-white/10 rounded-xl p-4">
              <h3 className="font-medium mb-2">What to expect:</h3>
              <ul className="space-y-2 text-sm">
                <li className="flex items-start gap-2">
                  <span className="bg-white text-[#E12D39] rounded-full w-5 h-5 flex items-center justify-center text-xs flex-shrink-0 mt-0.5">
                    1
                  </span>
                  <span>
                    Answer questions about your personal information and reasons for staying in
                    Poland
                  </span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="bg-white text-[#E12D39] rounded-full w-5 h-5 flex items-center justify-center text-xs flex-shrink-0 mt-0.5">
                    2
                  </span>
                  <span>Receive real-time feedback on your answers</span>
                </li>
                <li className="flex items-start gap-2">
                  <span className="bg-white text-[#E12D39] rounded-full w-5 h-5 flex items-center justify-center text-xs flex-shrink-0 mt-0.5">
                    3
                  </span>
                  <span>Get a detailed assessment of your performance</span>
                </li>
              </ul>
            </div>
            <p className="text-xs opacity-75">This mock interview will use 5 energy points.</p>
          </div>
        </div>

        <div className="w-full md:w-7/12 p-8 flex flex-col">
          <div className="mb-6">
            <h2 className="text-2xl font-bold text-[#0C3B5F] mb-2">Ready to practice?</h2>
            <p className="text-gray-500">
              This simulation will help you prepare for your real interview.
            </p>
          </div>

          <div className="flex-grow space-y-6">
            <div className="space-y-4">
              <div className="bg-[#f8f9fa] rounded-xl p-4 border border-gray-100">
                <h3 className="font-medium text-[#0C3B5F] flex items-center gap-2 mb-2">
                  <span className="bg-[#E12D39] text-white rounded-full w-6 h-6 flex items-center justify-center text-xs">
                    1
                  </span>
                  Language
                </h3>
                <p className="text-gray-600 text-sm">
                  Try to answer in Polish, even if your language is not Polish. You can answer in
                  any available language if needed.
                </p>
              </div>

              <div className="bg-[#f8f9fa] rounded-xl p-4 border border-gray-100">
                <h3 className="font-medium text-[#0C3B5F] flex items-center gap-2 mb-2">
                  <span className="bg-[#E12D39] text-white rounded-full w-6 h-6 flex items-center justify-center text-xs">
                    2
                  </span>
                  Feedback
                </h3>
                <p className="text-gray-600 text-sm">
                  You will receive tips based on your answers. Read them carefully and try to
                  improve your responses.
                </p>
              </div>

              <div className="bg-[#f8f9fa] rounded-xl p-4 border border-gray-100">
                <h3 className="font-medium text-[#0C3B5F] flex items-center gap-2 mb-2">
                  <span className="bg-[#E12D39] text-white rounded-full w-6 h-6 flex items-center justify-center text-xs">
                    3
                  </span>
                  Duration
                </h3>
                <p className="text-gray-600 text-sm">
                  The interview will take approximately 10-15 minutes to complete.
                </p>
              </div>
            </div>
          </div>

          <div className="mt-4">
            {error && (
              <Alert variant="destructive" className="mb-2">
                <AlertTitle>Cannot start an interview</AlertTitle>
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}

            <Button
              onClick={setupInterviewHandler}
              disabled={isPending}
              className="w-full bg-[#E12D39] hover:bg-[#c0252f] text-white rounded-xl py-6 text-lg font-medium transition-all"
            >
              {isPending ? (
                <div className="flex items-center justify-center gap-2">
                  <Loader2Icon className="h-5 w-5 animate-spin" />
                  <span>Preparing Interview...</span>
                </div>
              ) : (
                <div className="flex items-center justify-center gap-2">
                  <span>Start Interview</span>
                  <ArrowRightIcon className="h-5 w-5" />
                </div>
              )}
            </Button>
          </div>
        </div>
      </div>
    </Card>
  );
}
