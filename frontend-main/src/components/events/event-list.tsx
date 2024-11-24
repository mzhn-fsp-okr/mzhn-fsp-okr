import { SearchParams, SportEvent } from "@/api/events";
import { cn } from "@/lib/utils";
import { cva, VariantProps } from "class-variance-authority";
import EventCalendar2 from "./event-calendar2";
import EventCard from "./event-card";

const listVariants = cva("grid gap-2 ", {
  variants: {
    variant: {
      default: "grid-cols-1 sm:grid-cols-2 md:grid-cols-3",
      list: "grid-cols-1",
      calendar: "grid-cols-1",
      map: "grid-cols-1",
    },
  },
  defaultVariants: {
    variant: "default",
  },
});

export interface EventListProps extends VariantProps<typeof listVariants> {
  events: SportEvent[];
  filters?: SearchParams;
  onFiltersChange?: (v: SearchParams) => void;
}

export default function EventList({
  events,
  variant,
  filters,
  onFiltersChange,
}: EventListProps) {
  const isList = variant && ["default", "list"].includes(variant);
  const isCalendar = variant && ["calendar"].includes(variant);

  return (
    <section className={cn(listVariants({ variant }))}>
      {isList && events.map((e, i) => <EventCard key={i} event={e} />)}
      {isCalendar && (
        <EventCalendar2
          events={events}
          filters={filters!}
          onFiltersChange={onFiltersChange!}
        />
      )}
    </section>
  );
}
