import { apiFetch } from "@/lib/fetch";

export async function tgCheck() {
  const result = await apiFetch<{ telegram: string | null }>(
    "/integrations/",
    undefined,
    {},
    process.env.NEXT_PUBLIC_API_URL
  );

  return result.telegram != null;
}

export async function tgCode() {
  const result = await apiFetch<{ token: string; expireAt: string }>(
    "/integrations/telegram/code",
    undefined,
    {},
    process.env.NEXT_PUBLIC_API_URL
  );
  return result;
}
