"use client";

import { events } from "@/api/subcribes";
import EventList from "@/components/events/event-list";
import { useQuery } from "@tanstack/react-query";

export default function PageContent() {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["subscribe-events"],
    queryFn: events,
    placeholderData: [],
  });

  console.log(data);

  return (
    <section className="space-y-4 font-bold">
      <h1 className="text-xl">Мероприятия, на которые я подписан:</h1>
      <EventList events={data!} variant={"default"} />
    </section>
  );
}
