"use client";

import { useMemo } from "react";
import { Popover, PopoverContent, PopoverTrigger } from "./popover";

interface TextWithPopoverProps {
  content: string;
  contentClassName?: string;
  maxLength?: number;
}

export function TextWithPopover({
  content,
  contentClassName = "",
  maxLength = 150,
}: TextWithPopoverProps) {
  const [normalizedContent, isNormalized] = useMemo(() => {
    const cleanedContent = content.trim();
    if (cleanedContent.length <= maxLength) {
      return [cleanedContent, false];
    }
    return [`${cleanedContent.slice(0, maxLength)}...`, true];
  }, [content, maxLength]);

  return (
    <p className={contentClassName}>
      {normalizedContent}
      {isNormalized && (
        <Popover>
          <PopoverTrigger asChild>
            <span className="ml-1 cursor-pointer text-gray-800">Read more</span>
          </PopoverTrigger>
          <PopoverContent className="w-96 max-w-screen">
            <p className={contentClassName}>{content}</p>
          </PopoverContent>
        </Popover>
      )}
    </p>
  );
}
