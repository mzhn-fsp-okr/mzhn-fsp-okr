"use client";

import { getAvatar, getFullName, getRoleFriendly, User } from "@/api/user";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

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
            <div className="py-2"></div>
          </TabsContent>
          <TabsContent value="sport">
            <div className="space-y-2 py-2"></div>
          </TabsContent>
        </Tabs>
        {/*  */}
      </main>
    </section>
  );
}
