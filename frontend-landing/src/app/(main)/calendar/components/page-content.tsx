"use client";

import { search, SportTypeAlt } from "@/api/events";
import { Button } from "@/components/ui/button";
import { MultiSelect } from "@/components/ui/multi-select";
import { cn } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";
import { Filter, LoaderCircle } from "lucide-react";
import { useEffect, useMemo, useState } from "react";
import EventCalendar2 from "./calendar";
import FilterDialog from "./filter-dialog";
import useFilters from "./filter-hook";
import { replaceArrayEmpty } from "./util";

export default function PageContent({ sports }: { sports: SportTypeAlt[] }) {
  const [sportInit, setSportInit] = useState<boolean>(false);
  const sportsAvailable = toSportList(sports);

  const [filterOpen, setFilterOpen] = useState<boolean>(false);
  const [filters, setFilters] = useFilters();

  const searchFilters = useMemo(
    () =>
      replaceArrayEmpty({
        ...filters,
      }),
    [filters]
  );

  useEffect(() => {
    setSportInit(true);
  }, [filters]);

  const { data, isFetching, isError } = useQuery({
    queryKey: ["events", searchFilters],
    queryFn: async () => await search(searchFilters),
    placeholderData: {
      events: [],
      total: 0,
      page: 1,
      page_size: 0,
      page_total: 0,
      has_next: false,
      has_prev: false,
    },
  });

  if (isError) return <>error</>;

  return (
    <section className="relative">
      {isFetching && (
        <div className="absolute left-0 top-0 z-10 flex h-full w-full items-center justify-center bg-black/20">
          <LoaderCircle className="size-16 animate-spin text-white" />
        </div>
      )}
      <div className="flex gap-2 py-4">
        {sportInit && (
          <MultiSelect
            options={sportsAvailable}
            defaultValue={filters.sport_type_id}
            onValueChange={(sports) => {
              setFilters({ ...filters, sport_type_id: sports });
            }}
            variant={"secondary"}
            placeholder="Виды спорта"
          />
        )}
        <Button
          className="size-14"
          variant={"secondary"}
          onClick={() => setFilterOpen(true)}
        >
          <Filter
            className={cn(Object.keys(filters).length > 0) && "text-blue-500"}
          />
        </Button>
      </div>
      <EventCalendar2
        events={data!.events}
        filters={filters}
        onFiltersChange={setFilters}
      />

      <FilterDialog
        open={filterOpen}
        onOpenChange={setFilterOpen}
        filters={filters}
        onFiltersChange={setFilters}
      />
    </section>
  );
}

function toSportList(sports: SportTypeAlt[]) {
  return sports.map((s) => ({
    value: s.id,
    label: s.name,
  }));
}
