"use client";

import { SearchParams, SportEvent } from "@/api/events";
import dayGridPlugin from "@fullcalendar/daygrid";
import multiMonthPlugin from "@fullcalendar/multimonth";
import FullCalendar from "@fullcalendar/react";
import moment from "moment";
import { useRouter } from "next/navigation";
import { useEffect, useMemo, useState } from "react";

import "./event-calendar.css";

export interface EventCalendarProps2 {
  events: SportEvent[];
  filters: SearchParams;
  onFiltersChange: (value: SearchParams) => void;
}

export default function EventCalendar2({
  events,
  filters,
  onFiltersChange,
}: EventCalendarProps2) {
  const [dateStart, setDateStart] = useState<string>();
  const [dateEnd, setDateEnd] = useState<string>();
  const router = useRouter();

  const calendarEvents = useMemo(
    () => toCalendarEvents(events || []),
    [events]
  );

  useEffect(() => {
    onFiltersChange({
      ...filters,
      name: undefined,
      start_date: moment(dateStart, "YYYY-MM-DD").format("DD.MM.YYYY"),
      end_date: moment(dateEnd, "YYYY-MM-DD").format("DD.MM.YYYY"),
    });
  }, [dateStart, dateEnd]);

  return (
    <div className="relative">
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
