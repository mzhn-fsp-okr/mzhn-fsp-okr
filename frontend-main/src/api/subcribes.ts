import { apiFetch } from "@/lib/fetch";
import moment from "moment";
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

  return result.events.map((e) => ({
    ...e,
    sportSubtype: {
      ...e.sportSubtype,
      // @ts-expect-error ---
      sportType: e.sportSubtype.parent,
    },
    dates: {
      // @ts-expect-error ---
      from: moment(e.dates.date_from, "DD.MM.YYYY").format("YYYY-MM-DD"),
      // @ts-expect-error ---
      to: moment(e.dates.date_to, "DD.MM.YYYY").format("YYYY-MM-DD"),
    },
  }));
}

export async function sportSubscribe(id: string) {
  return await apiFetch(
    "/api/subscriptions/sport/subscribe",
    {
      sportId: id,
    },
    { method: "POST" }
  );
}

export async function sportUnsubscribe(id: string) {
  return await apiFetch(
    "/api/subscriptions/sport/unsubscribe",
    {
      sportId: id,
    },
    { method: "POST" }
  );
}

export async function sports() {
  const result = await apiFetch<SearchResponse>(
    "/subscriptions/sport/", // ????????????????
    undefined,
    {},
    process.env.NEXT_PUBLIC_API_URL
  );
  return result.events;
}
