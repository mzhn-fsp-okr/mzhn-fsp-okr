import { TitleInit } from "@/components/providers/title";
import PageContent from "./components/page-content";

export default function Page() {
  return (
    <section>
      <TitleInit title="Мои мероприятия" />
      <PageContent />
    </section>
  );
}
