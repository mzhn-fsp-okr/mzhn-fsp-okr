import { search } from "@/api/events";
import { Badge } from "@/components/ui/badge";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselThumbnailItemAlt,
  CarouselThumbnails,
} from "@/components/ui/carousel";
import { Dumbbell, Medal, Newspaper, Trophy } from "lucide-react";
import moment from "moment";
import Image from "next/image";

import banner1 from "@/assets/banners/banner-1.png";
import banner2 from "@/assets/banners/banner-2.png";
import "moment/locale/ru";

export default async function Home() {
  const events = await search({
    start_date: getFormattedDate(),
    page: 1,
    page_size: 10,
  });

  return (
    <section className="space-y-8">
      <div>
        <Carousel className="w-full text-white">
          <CarouselContent>
            <CarouselItem>
              <Image
                src={banner1}
                alt=""
                width={768}
                height={300}
                className="h-auto w-full rounded"
              />
            </CarouselItem>
            <CarouselItem>
              <Image
                src={banner2}
                alt=""
                width={768}
                height={300}
                className="h-auto w-full rounded"
              />
            </CarouselItem>
          </CarouselContent>
          <CarouselThumbnails className="flex justify-center">
            <CarouselThumbnailItemAlt />
            <CarouselThumbnailItemAlt />
          </CarouselThumbnails>
        </Carousel>
      </div>
      <div className="space-y-4 px-4">
        <h1 className="text-2xl font-bold text-white">События</h1>
        <Carousel className="w-full text-white">
          <CarouselContent className="ml-0 gap-2">
            {events.events.map((e, i) => (
              <CarouselItem
                key={i}
                className="relative rounded bg-sky-600 px-4 py-8 md:basis-1/3"
              >
                <div>
                  <p className="text-sm opacity-85">
                    {getLocalizedDateRange(e.dates.from, e.dates.to)}
                  </p>
                  <h1 className="text-md max-w-16 font-bold" title={e.name}>
                    {ellipsis(e.name, 34)}
                  </h1>
                  <p className="text-sm opacity-85">
                    {e.sportSubtype.sportType.name}
                  </p>
                </div>
                <div className="absolute right-8 top-1/2 -translate-y-1/2">
                  <Dumbbell className="size-16 text-white/15" />
                </div>
              </CarouselItem>
            ))}
          </CarouselContent>
        </Carousel>
      </div>
      <div className="space-y-4 px-4">
        <h1 className="text-2xl font-bold text-white">Спорт в СМИ</h1>
        <Carousel className="w-full text-white">
          <CarouselContent className="gap-2">
            <CarouselItem className="relative space-y-2 md:basis-1/4">
              <div className="group flex h-32 w-full items-center justify-center rounded-none bg-[#000007]">
                <Newspaper className="transition-all group-hover:scale-150" />
              </div>
              <div className="flex items-center gap-2">
                <Badge className="bg-white/5">ren.tv</Badge>
                <p className="text-xs text-muted-foreground">01 ноября 2024</p>
              </div>
              <h1>Молодежная сборная России обыграла команду из ...</h1>
              <a href="#" className="text-sky-600">
                Читать далее
              </a>
            </CarouselItem>
            <CarouselItem className="relative space-y-2 md:basis-1/4">
              <div className="group flex h-32 w-full items-center justify-center rounded-none bg-[#ff9900]">
                <Dumbbell className="transition-all group-hover:scale-150" />
              </div>
              <div className="flex items-center gap-2">
                <Badge className="bg-white/5">ren.tv</Badge>
                <p className="text-xs text-muted-foreground">01 ноября 2024</p>
              </div>
              <h1>Молодежная сборная России обыграла команду из ...</h1>
              <a href="#" className="text-sky-600">
                Читать далее
              </a>
            </CarouselItem>
            <CarouselItem className="relative space-y-2 md:basis-1/4">
              <div className="group flex h-32 w-full items-center justify-center rounded-none bg-[#221f73]">
                <Trophy className="transition-all group-hover:scale-150" />
              </div>
              <div className="flex items-center gap-2">
                <Badge className="bg-white/5">ren.tv</Badge>
                <p className="text-xs text-muted-foreground">01 ноября 2024</p>
              </div>
              <h1>Молодежная сборная России обыграла команду из ...</h1>
              <a href="#" className="text-sky-600">
                Читать далее
              </a>
            </CarouselItem>
            <CarouselItem className="relative space-y-2 md:basis-1/4">
              <div className="group flex h-32 w-full items-center justify-center rounded-none bg-[#262053]">
                <Medal className="transition-all group-hover:scale-150" />
              </div>
              <div className="flex items-center gap-2">
                <Badge className="bg-white/5">ren.tv</Badge>
                <p className="text-xs text-muted-foreground">01 ноября 2024</p>
              </div>
              <h1>Молодежная сборная России обыграла команду из ...</h1>
              <a href="#" className="text-sky-600">
                Читать далее
              </a>
            </CarouselItem>
          </CarouselContent>
        </Carousel>
      </div>
    </section>
  );
}

function getFormattedDate() {
  const now = new Date(); // Текущая дата
  const day = String(now.getDate()).padStart(2, "0"); // День с ведущим нулем
  const month = String(now.getMonth() + 1).padStart(2, "0"); // Месяц с ведущим нулем
  const year = now.getFullYear(); // Год
  return `${day}.${month}.${year}`;
}

function getLocalizedDateRange(startDate: string, endDate: string) {
  moment.locale("ru");
  const start = moment(startDate);
  const end = moment(endDate);
  if (start.isSame(end, "month")) {
    return `${start.format("DD")} - ${end.format("DD MMMM")}`;
  }
  if (start.isSame(end, "year")) {
    return `${start.format("DD MMMM")} - ${end.format("DD MMMM")}`;
  }
  return `${start.format("DD MMMM YYYY")} - ${end.format("DD MMMM YYYY")}`;
}

function ellipsis(s: string, max: number = 32) {
  if (s.length <= max) return s;
  return s.slice(0, max) + "...";
}
