"use client";

import Link from "next/link";

export function DashboardHeader() {
  return (
    <header className="flex-none flex items-center container mx-auto py-3 px-4 md:px-0">
      <div className="flex-1 flex items-center justify-start">
        <Link
          href="/portal"
          className="text-[#E12D39] hover:text-[#c0252f] text-xl font-bold transition-colors"
        >
          poznawca.pl
        </Link>
      </div>

      <div className="flex flex-1 justify-end items-center"></div>
    </header>
  );
}
