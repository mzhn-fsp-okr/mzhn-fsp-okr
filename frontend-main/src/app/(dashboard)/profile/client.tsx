"use client";

import { getAvatar, getFullName, User } from "@/api/user";
import { Avatar, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Send } from "lucide-react";
import TelegramDialog from "./tg-dialog";

export default function PageContent({
  user,
  isTg,
}: {
  user: User;
  isTg: boolean;
}) {
  const fullName = getFullName(user) ?? "Пользователь";
  const avatar = getAvatar(user);

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
        <Tabs defaultValue="notifications">
          <TabsList>
            <TabsTrigger value="notifications">Уведомления</TabsTrigger>
          </TabsList>
          <TabsContent value="notifications">
            <main className="flex">
              <div className="flex items-center gap-10 rounded border px-8 py-4">
                <Send className="size-10 text-blue-500" />
                <div className="flex flex-col gap-2">
                  <p className="text-lg font-bold">Telegram</p>
                  {isTg ? (
                    <Button disabled>Подключено</Button>
                  ) : (
                    <TelegramDialog />
                  )}
                </div>
              </div>
            </main>
          </TabsContent>
        </Tabs>
        {/*  */}
      </main>
    </section>
  );
}
