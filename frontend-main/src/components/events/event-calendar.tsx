"use client";

import { search, SportEvent } from "@/api/events";
import dayGridPlugin from "@fullcalendar/daygrid";
import multiMonthPlugin from "@fullcalendar/multimonth";
import FullCalendar from "@fullcalendar/react";
import { useQuery } from "@tanstack/react-query";
import { LoaderCircle } from "lucide-react";
import moment from "moment";
import { useRouter } from "next/navigation";
import { useMemo, useState } from "react";

import "./event-calendar.css";

export interface EventCalendarProps {}

export default function EventCalendar({}: EventCalendarProps) {
  const [dateStart, setDateStart] = useState<string>();
  const [dateEnd, setDateEnd] = useState<string>();
  const router = useRouter();

  const { data, isFetching, isError } = useQuery({
    queryKey: ["events", dateStart, dateEnd],
    queryFn: async () =>
      await search({
        start_date: moment(dateStart).format("DD.MM.YYYY"),
        end_date: moment(dateEnd).format("DD.MM.YYYY"),
      }),
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
  const calendarEvents = useMemo(
    () => toCalendarEvents(data!.events || []),
    [data]
  );

  return (
    <div className="relative">
      {isFetching && (
        <div className="absolute left-0 top-0 z-10 flex h-full w-full items-center justify-center bg-black/20">
          <LoaderCircle className="size-16 animate-spin text-white" />
        </div>
      )}
      {isError && (
        <div className="absolute left-0 top-0 z-10 flex h-full w-full items-center justify-center bg-black/50">
          <p className="text-red-500">При загрузке произошла ошибка</p>
        </div>
      )}
      <FullCalendar
        plugins={[dayGridPlugin, multiMonthPlugin]}
        timeZone="Europe/Moskow"
        events={calendarEvents}
        locale="ru"
        buttonText={{
          today: "Сегодня",
          month: "Месяц",
          week: "Неделя",
          year: "Год",
        }}
        headerToolbar={{
          left: "prev,next,today",
          center: "title",
          right: "dayGridMonth,dayGridWeek,multiMonthYear",
        }}
        datesSet={(dates) => {
          setDateStart(dates.startStr);
          setDateEnd(dates.endStr);
        }}
        eventClick={(arg) => router.push(arg.event.extendedProps.href)}
      />
    </div>
  );
}

function toCalendarEvents(events: SportEvent[]) {
  return events.map((e) => ({
    id: e.id,
    title: e.name,
    start: moment(e.dates.from).format("YYYY-MM-DD"),
    end: moment(e.dates.to).format("YYYY-MM-DD"),
    href: `/event/${e.id}`,
    color: stringToHSL(e.sportSubtype.sportType.name),
  }));
}

function stringToHSL(str: string): string {
  let hash = 0;
  for (let i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash);
  }
  const h = Math.abs(hash % 360);
  const s = 70;
  const l = 60;
  return `hsl(${h}, ${s}%, ${l}%)`;
}
