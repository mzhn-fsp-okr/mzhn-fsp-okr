"use client";

import dayGridPlugin from "@fullcalendar/daygrid";
import multiMonthPlugin from "@fullcalendar/multimonth";
import FullCalendar from "@fullcalendar/react";

import "./calendar.css";

export default function Calendar() {
  return (
    <FullCalendar
      plugins={[dayGridPlugin, multiMonthPlugin]}
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
    />
  );
}
