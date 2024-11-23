import { apiFetch } from "@/lib/fetch";

export interface SportSubtype {
  id: string;
  name: string;
  sportType: {
    id: string;
    name: string;
  };
}

export interface SportEvent {
  id: string;
  ekpId: string;
  sportSubtype: SportSubtype;
  name: string;
  description: string;
  dates: {
    from: string;
    to: string;
  };
  location: string;
  participants: number;
}

export interface SearchParams {
  page?: number;
  page_size?: number;
  start_date?: string;
  end_date?: string;
  sport_type_id?: string[];
  sport_subtype_id?: string[];
  min_age?: number;
  max_age?: number;
  sex?: boolean;
  min_participants?: number;
  max_participants?: number;
  location?: string;
}

export interface SearchResponse {
  events: SportEvent[];
  total: number;
  page: number;
  page_size: number;
  page_total: number;
  has_prev: boolean;
  has_next: boolean;
}

export async function search(data: SearchParams) {
  const result = await apiFetch<SearchResponse>(
    "/api/events/?" +
      new URLSearchParams(data as unknown as Record<string, string>).toString()
  );

  result.events = result.events ?? [];
  result.page = data.page ?? 1;
  result.page_size = data.page_size ?? 0;
  result.page_total = Math.ceil(result.total / result.page_size);
  result.has_prev = data.page != 1;
  result.has_next = data.page != result.page_total;

  return result;
}

export async function get(id: string) {
  const result = await apiFetch<{ events: SportEvent }>(`/api/events/${id}`);
  return result.events;
}
