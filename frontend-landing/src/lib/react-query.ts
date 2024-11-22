import { QueryClientConfig } from "@tanstack/react-query";

export const clientConfig = {
  defaultOptions: {
    queries: {
      staleTime: 60 * 1000,
      throwOnError(error, query) {
        return true;
      },
      retry(failureCount, error) {
        if (error.message === "NEXT_REDIRECT") return false;
        return true;
      },
    },
  },
} satisfies QueryClientConfig;
