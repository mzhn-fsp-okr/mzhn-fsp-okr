import { apiFetch } from "@/lib/fetch";

export interface SportEventOld {
  id: string;
  discipline: string;
  name: string;
  gender: string;
  place: string;
  count: number;
  dateStart: number;
  dateEnd: number;
}

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
  name?: string;
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

export interface SportTypeAlt {
  id: string;
  name: string;
}

export async function search(data: SearchParams) {
  const params = new URLSearchParams();

  for (const key in data) {
    // @ts-ignore
    if (Array.isArray(data[key])) {
      // @ts-ignore
      data[key].forEach((value) => params.append(key, value));
    } else {
      // @ts-ignore
      params.append(key, data[key]);
    }
  }
  const result = await apiFetch<SearchResponse>(
    "/api/events?" + params.toString()
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
  const result = await apiFetch<{ events: SportEvent }>(`api/events/${id}`);
  return result.events;
}

export async function sports() {
  const result = await apiFetch<{ sportTypes: SportTypeAlt[] }>(
    `api/events/sports/`
  );

  return result;
}

export async function sportsSearch(name: string) {
  const result = await apiFetch<{ sportTypes: SportTypeAlt[] }>(
    `api/events/sports/?name=${name}`
  );
  return result;
}
