import { logout } from "@/api/auth";
import { deleteSession } from "@/lib/auth/session";
import { redirect, RedirectType } from "next/navigation";

export async function GET() {
  try {
    await logout();
  } catch {
    // Ignore errors during server logout
  }
  await deleteSession();
  redirect("/sign-in?reload", RedirectType.replace);
}
