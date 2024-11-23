import { me, User } from "@/api/auth";
import { PropsWithChildren } from "react";

export default async function GuestOnly({ children }: PropsWithChildren) {
  let user: User;
  try {
    user = await me();
  } catch (e) {
    return <>{children}</>;
  }
  return <></>;
}
