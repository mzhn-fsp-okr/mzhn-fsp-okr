import { cookies } from "./cookie";

export async function createSession<T>(data: T) {
  const expiresAt = new Date(Date.now() + 7 * 24 * 60 * 60 * 1000);
  const encryptedData = await encrypt(data);
  cookies().set("session", encryptedData, {
    expires: expiresAt,
    sameSite: "lax",
    path: "/",
  });
}

// May be changed in future
export async function updateSession<T>(data: T) {
  return await createSession(data);
}

export async function getSession<T>() {
  const data = cookies().get("session");
  if (!data || !data.value) return undefined;
  return decrypt(data.value) as T;
}

export async function deleteSession() {
  cookies().delete("session");
}

async function encrypt<T>(data: T) {
  return JSON.stringify(data);
}

async function decrypt<T>(data: string) {
  return JSON.parse(data) as T;
}
