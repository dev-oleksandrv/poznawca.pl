"use client";

import Link from "next/link";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { CheckCircle, XCircle, AlertTriangle, Home, RotateCcw } from "lucide-react";
import { useInterviewStore } from "@/store/interview-store";

export function InterviewChatResult() {
  const result = useInterviewStore((state) => state.result);

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
    if (score >= 80) return <CheckCircle className={`size-6 ${getScoreColor(score)}`} />;
    if (score >= 60) return <AlertTriangle className={`size-6 ${getScoreColor(score)}`} />;
    return <XCircle className={`size-6 ${getScoreColor(score)}`} />;
  };

  const getPerformanceLevel = (score: number) => {
    if (score >= 80) return "Excellent";
    if (score >= 70) return "Good";
    if (score >= 60) return "Satisfactory";
    if (score >= 50) return "Needs Improvement";
    return "Insufficient";
  };

  const getGrammarLevel = (score: number) => {
    if (score >= 80) return "Excellent grammar skills";
    if (score >= 60) return "Good grammar with some errors";
    return "Needs significant improvement";
  };

  const getAccuracyLevel = (score: number) => {
    if (score >= 80) return "Highly accurate responses";
    if (score >= 60) return "Mostly accurate with some inconsistencies";
    return "Significant factual errors present";
  };

  if (!result) {
    return null;
  }

  return (
    <Card className="flex flex-col max-h-[90vh] w-full max-w-3xl mx-auto rounded-2xl border-0 shadow-lg overflow-hidden">
      <CardHeader className="py-6 bg-gradient-to-r from-[#0C3B5F] to-[#1a5a8a] text-white">
        <div className="flex justify-between items-center">
          <div>
            <p className="text-white/80 text-sm">Interview Results</p>
            <CardTitle className="text-2xl font-bold mt-1">Performance Assessment</CardTitle>
          </div>
          <div
            className={`${getScoreBg(result.total_score)} ${getScoreColor(result.total_score)} rounded-full px-4 py-2 font-bold text-sm`}
          >
            {getPerformanceLevel(result.total_score)}
          </div>
        </div>
        <p className="text-white/90 mt-2">
          Your interview performance has been evaluated. Review your scores and feedback below.
        </p>
      </CardHeader>

      <CardContent className="p-0 flex-1 overflow-auto">
        <div className="p-6 bg-white">
          <div className="flex justify-between items-center mb-6">
            <div className="flex items-center gap-2">
              <div
                className={`${getScoreBg(result.total_score)} w-12 h-12 rounded-full flex items-center justify-center`}
              >
                {getScoreIcon(result.total_score)}
              </div>
              <div>
                <p className="text-gray-500 text-sm">Total Score</p>
                <p className={`text-2xl font-bold ${getScoreColor(result.total_score)}`}>
                  {result.total_score}%
                </p>
              </div>
            </div>
            {/*<div className="text-right">*/}
            {/*  <p className="text-gray-500 text-sm">Date</p>*/}
            {/*  <p className="text-gray-700">{result.date}</p>*/}
            {/*</div>*/}
          </div>

          <Tabs defaultValue="overview" className="w-full">
            <TabsList className="grid grid-cols-3 mb-6">
              <TabsTrigger value="overview" className="rounded-lg">
                Overview
              </TabsTrigger>
              <TabsTrigger value="grammar" className="rounded-lg">
                Grammar
              </TabsTrigger>
              <TabsTrigger value="accuracy" className="rounded-lg">
                Accuracy
              </TabsTrigger>
            </TabsList>

            <TabsContent value="overview" className="mt-0">
              <div className="space-y-6">
                <div className="bg-gray-50 rounded-xl p-4">
                  <h3 className="text-gray-700 mb-2 text-lg font-bold">Overall Feedback</h3>
                  <p className="text-gray-600 text-sm">{result.total_feedback}</p>
                </div>

                <div className="grid grid-cols-2 gap-4">
                  <div className="bg-gray-50 rounded-xl p-4">
                    <div className="flex justify-between items-center mb-2">
                      <h4 className="font-medium text-gray-700">Grammar</h4>
                      <span className={`font-bold ${getScoreColor(result.grammar_score)}`}>
                        {result.grammar_score}%
                      </span>
                    </div>
                    <div className="w-full bg-gray-200 rounded-full h-2.5 mb-2">
                      <div
                        className={`h-2.5 rounded-full ${getScoreBg(result.grammar_score)}`}
                        style={{
                          width: `${result.grammar_score}%`,
                        }}
                      ></div>
                    </div>
                  </div>

                  <div className="bg-gray-50 rounded-xl p-4">
                    <div className="flex justify-between items-center mb-2">
                      <h4 className="font-medium text-gray-700">Accuracy</h4>
                      <span className={`font-bold ${getScoreColor(result.accuracy_score)}`}>
                        {result.accuracy_score}%
                      </span>
                    </div>
                    <div className="w-full bg-gray-200 rounded-full h-2.5 mb-2">
                      <div
                        className={`h-2.5 rounded-full ${getScoreBg(result.accuracy_score)}`}
                        style={{
                          width: `${result.accuracy_score}%`,
                        }}
                      ></div>
                    </div>
                  </div>
                </div>

                <div className="bg-gray-50 rounded-xl p-4">
                  <h3 className="font-bold text-lg text-gray-700 mb-2">Next Steps</h3>
                  <ul className="space-y-2 text-gray-600">
                    <li className="flex items-start gap-2">
                      <span className="bg-[#E12D39] text-white rounded-full size-5 flex items-center justify-center text-xs flex-shrink-0 mt-0.5">
                        1
                      </span>
                      <span className="text-sm">
                        Review your feedback in the Grammar and Accuracy tabs
                      </span>
                    </li>
                    <li className="flex items-start gap-2">
                      <span className="bg-[#E12D39] text-white rounded-full size-5 flex items-center justify-center text-xs flex-shrink-0 mt-0.5">
                        2
                      </span>
                      <span className="text-sm">Practice the areas where you scored lowest</span>
                    </li>
                    <li className="flex items-start gap-2">
                      <span className="bg-[#E12D39] text-white rounded-full size-5 flex items-center justify-center text-xs flex-shrink-0 mt-0.5">
                        3
                      </span>
                      <span className="text-sm">
                        Try another mock interview to improve your skills
                      </span>
                    </li>
                  </ul>
                </div>
              </div>
            </TabsContent>

            <TabsContent value="grammar" className="mt-0">
              <div className="space-y-6">
                <div className="flex items-center gap-4">
                  <div
                    className={`${getScoreBg(
                      result.grammar_score,
                    )} w-16 h-16 rounded-full flex items-center justify-center`}
                  >
                    <span className={`text-xl font-bold ${getScoreColor(result.grammar_score)}`}>
                      {result.grammar_score}%
                    </span>
                  </div>
                  <div>
                    <h3 className="font-medium text-gray-700">Grammar Score</h3>
                    <p className="text-gray-500 text-sm">{getGrammarLevel(result.grammar_score)}</p>
                  </div>
                </div>

                <div className="bg-gray-50 rounded-xl p-4">
                  <h3 className="font-bold text-lg text-gray-700 mb-2">Grammar Feedback</h3>
                  <p className="text-gray-600 text-sm">{result.grammar_feedback}</p>
                </div>
              </div>
            </TabsContent>

            <TabsContent value="accuracy" className="mt-0">
              <div className="space-y-6">
                <div className="flex items-center gap-4">
                  <div
                    className={`${getScoreBg(
                      result.accuracy_score,
                    )} w-16 h-16 rounded-full flex items-center justify-center`}
                  >
                    <span className={`text-xl font-bold ${getScoreColor(result.accuracy_score)}`}>
                      {result.accuracy_score}%
                    </span>
                  </div>
                  <div>
                    <h3 className="font-medium text-gray-700">Accuracy Score</h3>
                    <p className="text-gray-500 text-sm">
                      {getAccuracyLevel(result.accuracy_score)}
                    </p>
                  </div>
                </div>

                <div className="bg-gray-50 rounded-xl p-4">
                  <h3 className="font-bold text-lg text-gray-700 mb-2">Accuracy Feedback</h3>
                  <p className="text-gray-600 text-sm">{result.accuracy_feedback}</p>
                </div>

                {/*<div className="bg-gray-50 rounded-xl p-4">*/}
                {/*  <h3 className="font-medium text-gray-700 mb-2">Areas for Improvement</h3>*/}
                {/*  <ul className="space-y-2">*/}
                {/*    <li className="flex items-start gap-2 text-gray-600">*/}
                {/*      <AlertTriangle className="h-4 w-4 text-amber-500 mt-0.5" />*/}
                {/*      <span>Personal information consistency</span>*/}
                {/*    </li>*/}
                {/*    <li className="flex items-start gap-2 text-gray-600">*/}
                {/*      <AlertTriangle className="h-4 w-4 text-amber-500 mt-0.5" />*/}
                {/*      <span>Details about employment history</span>*/}
                {/*    </li>*/}
                {/*    <li className="flex items-start gap-2 text-gray-600">*/}
                {/*      <AlertTriangle className="h-4 w-4 text-amber-500 mt-0.5" />*/}
                {/*      <span>Knowledge of Polish cultural facts</span>*/}
                {/*    </li>*/}
                {/*  </ul>*/}
                {/*</div>*/}
              </div>
            </TabsContent>
          </Tabs>
        </div>
      </CardContent>

      <div className="p-6 bg-gray-50 border-t border-gray-100 flex flex-wrap gap-3 justify-between">
        <Button variant="outline" className="rounded-xl" asChild>
          <Link href="/portal/interview">
            <RotateCcw className="mr-2 h-4 w-4" />
            Try Again
          </Link>
        </Button>
        <Link href="/portal">
          <Button className="rounded-xl bg-[#0C3B5F] hover:bg-[#0a3050]">
            <Home className="mr-2 h-4 w-4" />
            Back to Main
          </Button>
        </Link>
      </div>
    </Card>
  );
}
