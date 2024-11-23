import { apiFetch } from "@/lib/fetch";
import { SearchResponse } from "./events";

export async function eventSubscribe(id: string) {
  return await apiFetch(
    "/api/subscriptions/event/subscribe",
    {
      eventId: id,
    },
    { method: "POST" }
  );
}

export async function eventUnsubscribe(id: string) {
  return await apiFetch(
    "/api/subscriptions/event/unsubscribe",
    {
      eventId: id,
    },
    { method: "POST" }
  );
}

export async function events() {
  const result = await apiFetch<SearchResponse>(
    "/subscriptions/event/",
    undefined,
    {},
    process.env.NEXT_PUBLIC_API_URL
  );
  return result.events;
}
