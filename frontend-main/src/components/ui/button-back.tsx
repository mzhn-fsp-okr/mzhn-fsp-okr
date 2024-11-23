"use client";

import { ArrowLeft } from "lucide-react";
import { useRouter } from "next/navigation";
import { Button } from "./button";

export default function ButtonBack() {
  const router = useRouter();
  return (
    <Button
      className="size-12"
      variant={"secondary"}
      onClick={() => router.back()}
    >
      <ArrowLeft />
    </Button>
  );
}
