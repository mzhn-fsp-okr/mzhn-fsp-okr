import { me } from "@/api/auth";
import { tgCheck } from "@/api/integrations";
import { TitleInit } from "@/components/providers/title";
import PageContent from "./client";

export default async function Page() {
  const user = await me();
  const isTgConnected = await tgCheck();

  return (
    <>
      <TitleInit title="Личный кабинет" />
      <PageContent user={user} isTg={isTgConnected} />
    </>
  );
}
