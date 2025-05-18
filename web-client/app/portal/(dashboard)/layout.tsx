import { ReactNode } from "react";
import { DashboardHeader } from "@/components/portal-common/dashboard-header";

interface PortalDashboardLayoutProps {
  children: ReactNode;
}

export default function PortalDashboardLayout({ children }: PortalDashboardLayoutProps) {
  return (
    <div className="min-h-screen bg-gradient-to-br from-[#f8f9fa] to-[#e9ecef] flex flex-col">
      <DashboardHeader />

      <div className="flex-1 flex flex-col container mx-auto pt-2 pb-6">{children}</div>
    </div>
  );
}
