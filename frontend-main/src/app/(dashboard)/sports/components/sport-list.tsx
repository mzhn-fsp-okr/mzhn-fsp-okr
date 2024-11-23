import { sports } from "@/api/events";
import { useQuery } from "@tanstack/react-query";
import { LoaderCircle } from "lucide-react";
import { useMemo } from "react";
import { SportEntry } from "./sport-entry";

export default function SportList({ search = "" }: { search: string }) {
  const { data, isFetching, isError } = useQuery({
    queryKey: ["sports"],
    queryFn: sports,
    placeholderData: { sportTypes: [] },
  });
  const entries = useMemo(() => {
    if (search == "") return data!.sportTypes;
    return data!.sportTypes.filter((e) => {
      const name = e.name.toLowerCase();
      return name.includes(search.toLowerCase());
    });
  }, [search, data]);

  if (isFetching) return <LoaderCircle className="animate-spin" />;

  return (
    <ul className="grid grid-cols-1 gap-2 sm:grid-cols-2 md:grid-cols-3">
      {entries.map((s, i) => (
        <SportEntry key={i} id={s.id} name={s.name} />
      ))}
    </ul>
  );
}
