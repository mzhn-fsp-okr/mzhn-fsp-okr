import client from "../util/client";

export interface Something {}

export async function example() {
  const data = await client.get<Something>(`/something`);
  return data.data;
}
