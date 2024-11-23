"use client";

import { search, SportTypeAlt } from "@/api/events";
import EventList, { EventListProps } from "@/components/events/event-list";
import EventViewRadio from "@/components/events/event-view-radio";
import { Button } from "@/components/ui/button";
import { MultiSelect } from "@/components/ui/multi-select";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { useDebounce } from "@/hooks/use-debounce";
import { usePrevious } from "@/hooks/use-prev";
import useSearchState from "@/hooks/use-search-state";
import { cn } from "@/lib/utils";
import { useQuery } from "@tanstack/react-query";
import {
  ArrowLeftToLine,
  ArrowRightToLine,
  Filter,
  LoaderCircle,
} from "lucide-react";
import moment from "moment";
import { useEffect, useMemo, useState } from "react";
import FilterDialog from "./filter-dialog";
import useFilters from "./filter-hook";
import { replaceArrayEmpty } from "./util";

export default function PageContent({ sports }: { sports: SportTypeAlt[] }) {
  const [sportInit, setSportInit] = useState<boolean>(false);
  const [name, setName] = useState("");
  const sportsAvailable = toSportList(sports);

  const [filterOpen, setFilterOpen] = useState<boolean>(false);
  const [variant, setVariant] = useSearchState<string>("variant", "default");
  const [page, setPage] = useSearchState<string, number>("page", "1", (value) =>
    parseInt(value)
  );
  const [filters, setFilters] = useFilters();

  const prevVariant = usePrevious({ variant });

  const searchFilters = useMemo(
    () =>
      replaceArrayEmpty({
        ...filters,
        page,
        page_size: 21,
      }),
    [filters, page]
  );

  useEffect(() => {
    setSportInit(true);
  }, [filters]);

  useEffect(() => {
    if (prevVariant?.variant == "calendar") {
      setFilters({
        ...filters,
        start_date: moment().format("DD.MM.YYYY"),
        end_date: undefined,
      });
      setName("");
    }
  }, [variant]);

  const debouncedSetName = useDebounce((name: string) => {
    setFilters({ ...filters, name });
  }, 1000);

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

  const paginationRange = useMemo(
    () => getPaginationRange(page, data!.page_total, 3),
    [page, data]
  );
  if (isError) return <>error</>;

  return (
    <section className="relative">
      {isFetching && (
        <div className="absolute left-0 top-0 z-10 flex h-full w-full items-center justify-center bg-black/20">
          <LoaderCircle className="size-16 animate-spin text-white" />
        </div>
      )}
      <div className="flex flex-col gap-2 md:flex-row">
        {variant != "calendar" && (
          <input
            className="w-full flex-grow rounded bg-secondary p-4"
            placeholder="Поиск"
            value={name}
            onChange={(e) => {
              setName(e.target.value);
              debouncedSetName(e.target.value);
            }}
          />
        )}
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
          className="w-full md:size-14"
          variant={"secondary"}
          onClick={() => setFilterOpen(true)}
        >
          <Filter
            className={cn(Object.keys(filters).length > 0) && "text-blue-500"}
          />
          <span className="inline md:hidden">Фильтр</span>
        </Button>
      </div>

      <div className="flex flex-col justify-between gap-4 py-2 text-sm text-muted-foreground md:flex-row">
        <div>Страница: {page}</div>
        <div className="flex gap-2">
          отображать:
          <EventViewRadio value={variant} setValue={(v) => setVariant(v!)} />
        </div>
      </div>
      <EventList
        events={data!.events}
        variant={variant as EventListProps["variant"]}
        filters={filters}
        onFiltersChange={setFilters}
      />
      <Pagination className="py-2">
        <PaginationContent>
          {page != 1 && (
            <PaginationItem>
              <PaginationLink href="#" onClick={() => setPage(1)}>
                <ArrowLeftToLine />
              </PaginationLink>
            </PaginationItem>
          )}
          {data!.has_prev && (
            <PaginationItem>
              <PaginationPrevious href="#" onClick={() => setPage(page - 1)} />
            </PaginationItem>
          )}
          {paginationRange.map((p, i) => (
            <PaginationItem key={i}>
              <PaginationLink
                href="#"
                onClick={() => setPage(p)}
                isActive={page == p}
              >
                {p}
              </PaginationLink>
            </PaginationItem>
          ))}
          {data!.has_next && (
            <PaginationItem>
              <PaginationNext href="#" onClick={() => setPage(page + 1)} />
            </PaginationItem>
          )}
          {page != data!.page_total && (
            <PaginationItem>
              <PaginationLink
                href="#"
                onClick={() => setPage(data!.page_total)}
              >
                <ArrowRightToLine />
              </PaginationLink>
            </PaginationItem>
          )}
        </PaginationContent>
      </Pagination>
      <FilterDialog
        open={filterOpen}
        onOpenChange={setFilterOpen}
        filters={filters}
        onFiltersChange={setFilters}
      />
    </section>
  );
}

function getPaginationRange(
  page: number,
  totalPages: number,
  X: number
): number[] {
  const half = Math.floor(X / 2);
  let start = Math.max(page - half, 1);
  const end = Math.min(start + X - 1, totalPages);

  start = Math.max(end - X + 1, 1);
  return Array.from({ length: end - start + 1 }, (_, i) => start + i);
}

function toSportList(sports: SportTypeAlt[]) {
  return sports.map((s) => ({
    value: s.id,
    label: s.name,
  }));
}

function debounce(func: Function, timeout = 300) {
  let timer: Parameters<typeof clearTimeout>[0];
  return (...args: any) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      // @ts-ignore
      func.apply(this, args);
    }, timeout);
  };
}
