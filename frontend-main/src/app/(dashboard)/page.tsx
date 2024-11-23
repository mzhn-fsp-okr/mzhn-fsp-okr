import { sports } from "@/api/events";
import { TitleInit } from "@/components/providers/title";
import PageContent from "./components/page-content";

export default async function Home() {
  const sportsList = await sports();
  return (
    <main className="flex w-full flex-col gap-2">
      <TitleInit title="Все мероприятия" />
      <PageContent sports={sportsList.sportTypes} />
    </main>
  );
}
