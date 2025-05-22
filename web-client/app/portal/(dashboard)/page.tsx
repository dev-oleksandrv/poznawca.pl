import { DashboardNavigationSection } from "@/components/dashboard/dashboard-navigation-section";
import { DashboardPremiumCtaSection } from "@/components/dashboard/dashboard-premium-cta-section";
import { DashboardRecentResultsSection } from "@/components/dashboard/dashboard-recent-results-section";

export default function DashboardPage() {
  return (
    <>
      <DashboardNavigationSection />
      <DashboardPremiumCtaSection />
      <DashboardRecentResultsSection />
    </>
  );
}
