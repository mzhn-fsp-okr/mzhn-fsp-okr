import { AuthResult } from "@/api/auth";
import { redirect } from "next/navigation";
import { cookies } from "./auth/cookie";
import { getSession } from "./auth/session";

export async function apiFetch<T>(
  url: string,
  body?: unknown,
  options: RequestInit = {},
  baseUrl?: string
): Promise<T> {
  baseUrl ??= process.env.NEXT_PUBLIC_SELF_URL!;
  const authData = await getSession<AuthResult>();

  const headers = new Headers(options.headers ?? {});
  headers.set("Cookie", cookies().toString());
  headers.set("Content-Type", "application/json");

  if (authData) {
    headers.set("Authorization", `Bearer ${authData.accessToken}`);
  }

  const requestOptions = {
    body: JSON.stringify(body),
    ...options,
    headers,
  };
  const response = await fetch(`${baseUrl}${url}`, requestOptions);

  if (!response.ok) {
    if (response.status === 401) {
      redirect("/sign-out");
    }

    let errorMessage = await response.text();
    try {
      const errorData = JSON.parse(errorMessage);
      if (errorData["error"]) errorMessage = errorData["error"];
    } catch {}
    throw new Error(errorMessage || "Network error");
  }
  return await response.json();
}
