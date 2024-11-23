import { sports } from "@/api/events";
import PageContent from "./components/page-content";

export default async function Page() {
  const sportsList = await sports();
  return (
    <section className="space-y-4 px-4">
      <h1 className="text-2xl font-bold text-white">Анонс событий</h1>
      <PageContent sports={sportsList.sportTypes} />
    </section>
  );
}
