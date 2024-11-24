"use server";

import * as api from "@/api/auth";
import { createSession } from "@/lib/auth/session";

type ReturnValue = [boolean, string | null];

export async function signIn(data: api.AuthCredentials): Promise<ReturnValue> {
  try {
    const authData = await api.auth(data);
    await createSession(authData);
    return [true, null];
  } catch (error) {
    if (error instanceof Error && error.message) {
      return [false, error.message];
    }
    throw error;
  }
}
