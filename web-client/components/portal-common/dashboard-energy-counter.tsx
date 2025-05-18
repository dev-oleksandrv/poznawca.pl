import { ZapIcon } from "lucide-react";

export function DashboardEnergyCounter() {
  return (
    <div className="flex flex-row items-center gap-2 text-[#E12D39]">
      <ZapIcon className="size-6" />

      <span className="text-xl font-black">8</span>
    </div>
  );
}
