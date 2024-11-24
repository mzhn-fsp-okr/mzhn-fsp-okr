"use client";

import { invoke } from "@/api/ai";
import { sportsSearch } from "@/api/events";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";
import { useMutation } from "@tanstack/react-query";
import { Bot, LoaderCircle, MessageCircle, Send, User } from "lucide-react";
import { FormEvent, PropsWithChildren, useState } from "react";
import { replaceEmpty } from "./util";

interface Message {
  bot: boolean;
  text: string;
  link?: string;
}

const possibleMessages = [
  "Вот что я нашёл по твоему запросу:",
  "Нашёл тебе несколько соревнований, смотри:",
  "Как тебе эти соревнования? Смотри:",
  "Вот список соревнований по твоему запросу:",
];

export default function AiDialog() {
  const [message, setMessage] = useState("");
  const [messageHistory, setMessageHistory] = useState<Message[]>([
    {
      bot: true,
      text: "Привет! Я помогу тебе подыскать соревнования!\n\nНапример: Подскажи мне соревнования по плаванию на следующей неделе",
    },
  ]);

  const clear = () => {
    setMessageHistory([
      {
        bot: true,
        text: "Привет! Я помогу тебе подыскать соревнования!\n\nНапример: Подскажи мне соревнования по плаванию на следующей неделе",
      },
    ]);
  };

  const invokeChat = useMutation({
    mutationFn: async (text: string) => {
      const response = await invoke(text);
      const index = Math.round(Math.random() * 4);

      const filters = replaceEmpty(response);
      if (filters.sport) {
        const values = await sportsSearch(filters.sport);
        filters.sport_type_id = values.sportTypes.map((v) => v.id);
        if (filters.sport_type_id.length == 0) {
          delete filters.sport_type_id;
        }
        delete filters.sport;
      }

      if (filters.name == "СОБЫТИЕ") {
        delete filters.name;
      }

      const url = `/?filters=${JSON.stringify(filters)}&page=1`;

      setMessageHistory([
        ...messageHistory,
        {
          bot: true,
          text: possibleMessages[index],
          link: url,
        },
      ]);
      setMessage("");
    },
  });

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setMessageHistory([
      ...messageHistory,
      {
        bot: false,
        text: message,
      },
    ]);
    await invokeChat.mutate(message);
  };

  return (
    <Dialog
      onOpenChange={() => {
        clear();
      }}
    >
      <DialogTrigger asChild>
        <Button className="fixed bottom-4 right-4 size-14" title="Открыть ИИ">
          <MessageCircle />
        </Button>
      </DialogTrigger>
      <DialogContent className="w-full md:min-w-[800px]">
        <DialogHeader>
          <DialogTitle>ИИ-помощник</DialogTitle>
        </DialogHeader>
        <div className="space-y-4 rounded bg-gray-100 px-4 py-4">
          {messageHistory.map((h, i) => (
            <ChatMessage key={i} bot={h.bot} link={h.link}>
              <p>{h.text}</p>
            </ChatMessage>
          ))}
        </div>
        <form className="flex items-center gap-2" onSubmit={onSubmit}>
          <Input
            className="flex-grow py-6"
            placeholder="Введите сообщение"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            disabled={invokeChat.isPending}
          />
          <Button
            className="size-12"
            type="submit"
            disabled={message.length <= 5 || messageHistory.length > 1}
          >
            {invokeChat.isPending ? (
              <LoaderCircle className="animate-spin" />
            ) : (
              <Send />
            )}
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}

interface ChatMessageProps extends PropsWithChildren {
  bot?: boolean;
  link?: string;
}

function ChatMessage({ children, bot, link }: ChatMessageProps) {
  return (
    <div className={cn("flex w-full text-sm", !bot && "justify-end")}>
      <div
        className={cn(
          "flex items-center gap-4 md:max-w-[50%]",
          bot ? "" : "flex-row-reverse justify-start"
        )}
      >
        <div
          className={cn(
            "flex aspect-square size-10 items-center justify-center rounded-full",
            bot ? "bg-blue-400" : "bg-green-300"
          )}
        >
          {bot ? <Bot /> : <User />}
        </div>
        <div className="whitespace-pre-line rounded bg-white p-2">
          {children}
          <br />
          {link && (
            <a href={link} className="py-4 text-blue-400 underline">
              Просмотреть
            </a>
          )}
        </div>
      </div>
    </div>
  );
}
