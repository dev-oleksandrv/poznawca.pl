import { Card, CardContent, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowLeftIcon } from "lucide-react";
import Link from "next/link";

export function DashboardInterviewSetupError() {
  return (
    <Card className="p-4">
      <CardTitle>
        <h3 className="text-xl font-black">No available interviewers right now</h3>
      </CardTitle>
      <CardContent className="p-0 mt-2">
        <p>Unfortunately, we don't have available interviewers right now.</p>
        <Button
          className="bg-[#E12D39] hover:bg-[#c0252f] text-white rounded-xl py-5 font-medium transition-all mt-1"
          asChild
        >
          <Link href="/portal">
            <ArrowLeftIcon />
            <span>Back</span>
          </Link>
        </Button>
      </CardContent>
    </Card>
  );
}
