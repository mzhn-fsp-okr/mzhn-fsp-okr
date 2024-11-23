import { sports } from "@/api/events";
import { Button } from "@/components/ui/button";
import { useQuery } from "@tanstack/react-query";
import { BellPlus, LoaderCircle } from "lucide-react";
import { useMemo } from "react";

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
    <ul className="grid grid-cols-2 gap-2 md:grid-cols-3">
      {entries.map((s, i) => (
        <SportEntry key={i} id={s.id} name={s.name} />
      ))}
    </ul>
  );
}

function SportEntry({ id, name }: { id: string; name: string }) {
  return (
    <li className="flex items-center justify-between rounded border px-4 py-8">
      <p>{name}</p>
      <Button size="icon" variant="secondary">
        <BellPlus />
      </Button>
    </li>
  );
}
