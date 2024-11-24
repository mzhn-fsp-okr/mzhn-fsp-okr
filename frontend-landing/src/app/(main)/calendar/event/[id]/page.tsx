import { get } from "@/api/events";
import ButtonBack from "@/components/ui/button-back";
import { CalendarClock, Users } from "lucide-react";
import moment from "moment";
import DisciplinesView from "./disciplines-view";
import EventAvatar from "./event-avatar";

import "moment/locale/ru";

interface PageProps {
  params: {
    id: string;
  };
}

export default async function Page({ params }: PageProps) {
  const event = await get(params.id);

  const disciplines = filterDisciplines(event.description);

  return (
    <main className="flex flex-col gap-4 px-4 text-white/80">
      <section className="flex flex-col items-center justify-between gap-4 sm:flex-row">
        <div className="flex items-center justify-center gap-4 sm:justify-start">
          <EventAvatar event={event} className="size-24" />
          <div>
            <h1 className="text-2xl font-bold">{event.name}</h1>
            <h2 className="text-lg text-muted-foreground">
              {event.sportSubtype.sportType.name} - {event.sportSubtype.name}
            </h2>
          </div>
        </div>
        <div className="flex items-center gap-8">
          <p>{getEventStatus(event.dates.from, event.dates.to)}</p>
          <ul className="flex gap-2">
            <li>
              <ButtonBack />
            </li>
          </ul>
        </div>
      </section>
      <section className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <div className="flex items-center gap-8 rounded bg-blue-600 p-8 text-white transition-all hover:brightness-105">
          <Users className="size-10" />
          <div>
            <h1 className="text-lg font-bold">Участников:</h1>
            <h2 className="text-2xl">{event.participants}</h2>
          </div>
        </div>
        <div className="flex items-center gap-8 rounded bg-purple-600 p-8 text-white transition-all hover:brightness-105">
          <CalendarClock className="size-10" />
          <div>
            <h1 className="text-lg font-bold">Дата проведения:</h1>
            <h2 className="space-x-1 text-2xl">
              <span>
                {moment(event.dates.from).locale("ru").format("DD MMMM")}
              </span>
              <span>-</span>
              <span>
                {moment(event.dates.to).locale("ru").format("DD MMMM YYYY")}
              </span>
            </h2>
          </div>
        </div>
      </section>
      <section>
        <div className="space-y-2">
          <p className="font-bold">Дисциплины: </p>
          <DisciplinesView disciplines={disciplines} />
        </div>
      </section>
      <section>
        <div className="flex gap-2">
          <p className="font-bold">Кто может участвовать: </p>
          <p>{filterWhoCan(event.description)}</p>
        </div>
        <div className="flex gap-2">
          <p className="font-bold">Место проведения: </p>
          <p>{event.location}</p>
        </div>
      </section>
    </main>
  );
}

function filterWhoCan(string: string) {
  const match = string.match(/^(.*?)([A-ЯA-Z]{2})/u);
  return match ? match[1] : string;
}

function filterDisciplines(string: string) {
  const match = string.match(/([A-ЯA-Z]{2}.*)$/u);
  return (match ? match[0] : string).split(",");
}

function getEventStatus(startDate: string, endDate: string) {
  const now = moment();
  const start = moment(startDate);
  const end = moment(endDate);

  const timeToStart = moment.duration(start.diff(now));
  const timeToEnd = moment.duration(end.diff(now));
  const timeSinceEnd = moment.duration(now.diff(end));

  if (now.isBefore(start)) {
    return `Начнется через ${timeToStart.humanize()}`;
  } else if (now.isBetween(start, end)) {
    return `Идёт (осталось ${timeToEnd.humanize()})`;
  } else if (now.isAfter(end)) {
    return `Закончилось ${timeSinceEnd.humanize()} назад`;
  }
}
