import Calendar from "./components/calendar";

export default function Page() {
  return (
    <section className="space-y-4 px-4">
      <h1 className="text-2xl font-bold text-white">Анонс событий</h1>
      <Calendar />
    </section>
  );
}
