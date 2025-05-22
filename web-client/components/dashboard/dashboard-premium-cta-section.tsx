import { Card, CardContent } from "@/components/ui/card";
import { AwardIcon, BarChart3Icon, CrownIcon, ZapIcon } from "lucide-react";
import { Button } from "@/components/ui/button";

export function DashboardPremiumCtaSection() {
  return (
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
  );
}
