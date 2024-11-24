import { LoaderCircle } from "lucide-react";

export default async function Loader() {
  return (
    <main className="flex w-full items-center justify-center">
      <LoaderCircle className="size-16 animate-spin" />
    </main>
  );
}
