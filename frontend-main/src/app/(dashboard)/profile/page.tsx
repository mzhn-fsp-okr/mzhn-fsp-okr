import { me } from "@/api/auth";
import { TitleInit } from "@/components/providers/title";
import PageContent from "./client";

export default async function Page() {
  const user = await me();
  return (
    <>
      <TitleInit title="Личный кабинет" />
      <PageContent user={user} />
    </>
  );
}
