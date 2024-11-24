"use client";

import { useState } from "react";
import SportList from "./sport-list";

export default function PageContent() {
  const [search, setSearch] = useState("");

  return (
    <section className="space-y-2">
      <div className="flex gap-2">
        <input
          className="w-full flex-grow rounded bg-secondary p-4"
          placeholder="Поиск"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
      </div>
      <SportList search={search} />
    </section>
  );
}
