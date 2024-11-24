"use client";

import { tgCode } from "@/api/integrations";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { useQuery } from "@tanstack/react-query";
import { LoaderCircle } from "lucide-react";
import { useMemo } from "react";

export default function TelegramDialog() {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["tg-code"],
    queryFn: tgCode,
  });

  const link = useMemo(() => {
    if (!data) return "";
    return `https://t.me/yoursport_notificator_bot?start=${data.token}`;
  }, [data]);
  const qrUrl = useMemo(() => {
    if (!data) return "";
    return `https://quickchart.io/qr?size=512&text=https%3A%2F%2Ft.me%2Fyoursport_notificator_bot%3Fstart%3D${data.token}`;
  }, [data]);

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Подключить</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Подключение к Telegram</DialogTitle>
          <DialogDescription>
            Отсканируйте код чтобы быть в курсе спортивных событий
          </DialogDescription>
        </DialogHeader>
        <div className="flex flex-col items-center gap-4">
          {isLoading && (
            <div className="flex size-32 animate-pulse items-center justify-center bg-secondary md:size-64">
              <LoaderCircle className="animate-spin" />
            </div>
          )}
          {isError && (
            <div className="flex size-32 animate-pulse items-center justify-center bg-secondary md:size-64">
              <p className="text-red-600">ОШИБКА</p>
            </div>
          )}
          {!isLoading && !isError && (
            <>
              <img src={qrUrl} alt="qr" className="size-32 md:size-64" />
              <p>или</p>
              <a
                href={link}
                target="_blank"
                className="text-blue-500 underline"
              >
                перейдите по ссылке
              </a>
            </>
          )}
          <p className="text-sm text-muted-foreground">
            Отсканировали? Обновите страницу!
          </p>
        </div>
      </DialogContent>
    </Dialog>
  );
}
