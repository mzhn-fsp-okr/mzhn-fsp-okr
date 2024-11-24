"use client";

import useProfile from "@/hooks/use-profile";
import { PropsWithChildren, useMemo } from "react";

interface RoleRequiredProps extends PropsWithChildren {
  role: string | string[];
  ignoreGuest?: boolean;
}

export default function RoleRequired({
  ignoreGuest,
  role,
  children,
}: RoleRequiredProps) {
  const { data, isLoading, isError } = useProfile(!ignoreGuest);

  const roles = useMemo(() => (Array.isArray(role) ? role : [role]), [role]);
  const hasPermission = useMemo(() => {
    if (isLoading || isError) return false;
    return roles.some((role) => data?.roles.includes(role));
  }, [isLoading, isError, roles]); // eslint-disable-line react-hooks/exhaustive-deps -- data?.roles will be updated when isLoading is updated

  if (!hasPermission) return <></>;
  return <>{children}</>;
}
