import "server-only";

import { AuthResult } from "@/api/auth";
import { redirect } from "next/navigation";
import { cache } from "react";
import { getSession } from "./session";

export const verifySession = cache(async () => {
  const session = await getSession<AuthResult>();
  if (!session) {
    redirect("/sign-in");
  }
  return { isAuth: true, session };
});
