"use client";

import { useRouter, useSearchParams } from "next/navigation";
import { Suspense, useEffect } from "react";

export default function PageReloader() {
  return (
    <Suspense>
      <PageReloaderCore />
    </Suspense>
  );
}

function PageReloaderCore() {
  const params = useSearchParams();
  const router = useRouter();
  useEffect(() => {
    if (!params.has("reload")) return;
    router.replace("/sign-in");
  }, [params, router]);
  return <></>;
}
