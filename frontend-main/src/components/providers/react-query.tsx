"use client";

import { clientConfig } from "@/lib/react-query";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React, { useState } from "react";

export default function ReactQueryProvider({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  const [queryClient] = useState(() => new QueryClient(clientConfig));
  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}
