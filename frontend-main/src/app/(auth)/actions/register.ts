"use server";

import * as api from "@/api/auth";
import { createSession } from "@/lib/auth/session";

type ReturnValue = [boolean, string | null];

export async function register(data: api.RegisterData): Promise<ReturnValue> {
  try {
    await api.register(data);

    const authData = await api.auth({
      email: data.email,
      password: data.password,
    });
    await createSession(authData);
    return [true, null];
  } catch (error) {
    if (error instanceof Error && error.message) {
      return [false, error.message];
    }
    throw error;
  }
}
