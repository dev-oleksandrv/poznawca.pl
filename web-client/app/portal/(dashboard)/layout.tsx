import type { BaseLayoutProps } from "@/types/ui-types";
import { DashboardHeader } from "@/components/dashboard/dashboard-header";

export default function DashboardLayout({ children }: BaseLayoutProps) {
  return (
    <div className="min-h-screen bg-gradient-to-br from-[#f8f9fa] to-[#e9ecef] flex flex-col">
      <DashboardHeader />

      <div className="flex-1 flex flex-col container mx-auto pt-2 pb-6 px-4 md:px-0">
        {children}
      </div>
    </div>
  );
}
