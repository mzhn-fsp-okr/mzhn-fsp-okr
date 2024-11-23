"use client";

import { SportEventOld } from "@/api/events";
import { getAvatar, getFullName, getRoleFriendly, User } from "@/api/user";
import EventList from "@/components/events/event-list";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import SportList from "../sports/components/sport-list";

const event = {
  id: "000",
  name: "ЧЕМПИОНАТ ЦЕНТРАЛЬНОГО ФЕДЕРАЛЬНОГО ОКРУГА",
  discipline: "АВИАМОДЕЛЬНЫЙ СПОРТ",
  count: 25,
  gender: "женщины, мужчины от 14 лет и старше",
  place: "КЛАСС F-1D, дисциплина F-1D",
  dateStart: 1732611543,
  dateEnd: 1732894444,
} satisfies SportEventOld;

export default function PageContent({ user }: { user: User }) {
  const fullName = getFullName(user) ?? "Пользователь";
  const avatar = getAvatar(user);
  const role = getRoleFriendly(user);

  return (
    <section className="flex flex-col gap-2 divide-y">
      <header className="flex flex-col items-center justify-center gap-1">
        <Avatar className="size-32 rounded">
          <AvatarImage src={avatar} />
        </Avatar>
        <h1 className="text-lg font-bold">{fullName}</h1>
        <h2 className="">{user.email}</h2>
      </header>
      <main className="py-4">
        <Tabs defaultValue="events">
          <TabsList>
            <TabsTrigger value="events">Мероприятия</TabsTrigger>
            <TabsTrigger value="sport">Дисциплины</TabsTrigger>
          </TabsList>
          <TabsContent value="events">
            <div className="py-2">
              <EventList events={[event, event]} variant={"default"} />
            </div>
          </TabsContent>
          <TabsContent value="sport">
            <div className="space-y-2 py-2">
              <SportList />
            </div>
          </TabsContent>
        </Tabs>
        {/*  */}
      </main>
    </section>
  );
}
