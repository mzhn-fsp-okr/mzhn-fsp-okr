import { apiFetch } from "@/lib/fetch";
import { SearchParams } from "./events";

export async function invoke(message: string) {
  const result = await apiFetch<{ output: string }>(
    "/api/ai/chat/invoke",
    {
      input: {
        input: message,
      },
    },
    { method: "POST" }
  );
  return result.output as SearchParams;
}
