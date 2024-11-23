import { me, User } from "@/api/auth";
import { PropsWithChildren } from "react";

interface RoleRequiredProps extends PropsWithChildren {
  role: string | string[];
  ignoreGuest?: boolean;
}

export default async function RoleRequired({
  ignoreGuest,
  role,
  children,
}: RoleRequiredProps) {
  let user: User;
  try {
    user = await me();
  } catch (e) {
    if (ignoreGuest) return <></>;
    throw e;
  }

  const roles = Array.isArray(role) ? role : [role];
  const hasPermission = roles.some((role) => user.roles.includes(role));

  if (!hasPermission) return <></>;
  return <>{children}</>;
}
