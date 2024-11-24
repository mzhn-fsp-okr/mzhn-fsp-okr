"use client";

import useProfile from "@/hooks/use-profile";
import { PropsWithChildren } from "react";

export default function GuestOnly({ children }: PropsWithChildren) {
  const { isError } = useProfile(false);
  if (isError) return <>{children}</>;
  return <></>;
}
