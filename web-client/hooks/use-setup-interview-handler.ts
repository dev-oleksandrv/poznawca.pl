import { useState } from "react";
import { useRouter } from "next/navigation";

export function useSetupInterviewHandler() {
  const router = useRouter();

  const [error, setError] = useState<string | null>("");
  const [isPending, setIsPending] = useState<boolean>(false);

  const setupInterviewHandler = async () => {
    setIsPending(true);

    // TODO: Implement
  };

  return {
    error,
    isPending,
    setupInterviewHandler,
  };
}
