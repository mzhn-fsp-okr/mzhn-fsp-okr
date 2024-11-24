import { SportEvent } from "@/api/events";
import { cn } from "@/lib/utils";
import { ArrowRight, Pin } from "lucide-react";
import Link from "next/link";
import EventAvatar from "./event-avatar";

import moment from "moment";
import "moment/locale/ru";

export interface EventCardProps {
  event: SportEvent;
  pinned?: boolean;
}

export default function EventCard({ event, pinned }: EventCardProps) {
  const link = `/event/${event.id}`;
  return (
    <Link
      href={link}
      className={cn(
        "group relative flex flex-col items-center gap-4 rounded border bg-white p-6 text-center transition-all hover:border-blue-500 sm:flex-row sm:items-start sm:text-start",
        pinned && "border-blue-500"
      )}
    >
      <EventAvatar event={event} />
      <div className="flex max-w-[60%] flex-col items-center text-center sm:items-start sm:text-start">
        <h1 className="text-sm font-bold" title={event.name}>
          {ellipsis(event.name, 36)}
        </h1>
        <h2 className="text-xs text-muted-foreground">
          {event.sportSubtype.sportType.name}
        </h2>
        <p className="space-x-1 whitespace-nowrap text-xs text-muted-foreground">
          <span>{moment(event.dates.from).locale("ru").format("DD MMMM")}</span>
          <span>-</span>
          <span>{moment(event.dates.to).locale("ru").format("DD MMMM")}</span>
        </p>
      </div>
      {pinned && (
        <div className="absolute right-2 top-2 rotate-45 text-muted-foreground">
          <Pin className="size-4" />
        </div>
      )}
      <div className="absolute right-0 top-0 flex h-full items-center justify-center bg-blue-300 px-6 py-4 text-white opacity-0 transition-all group-hover:opacity-100">
        <ArrowRight />
      </div>
    </Link>
  );
}

function ellipsis(s: string, max: number = 32) {
  if (s.length <= max) return s;
  return s.slice(0, max) + "...";
}
