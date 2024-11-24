import { SearchParams } from "@/api/events";
import useSearchState from "@/hooks/use-search-state";
import moment from "moment";
import { useEffect, useState } from "react";
import { replaceEmpty } from "./util";

export default function useFilters(): [
  SearchParams,
  (value: SearchParams) => void,
] {
  const [filters, setFilters] = useState<SearchParams>({
    start_date: moment().format("DD.MM.YYYY"),
  });
  const [searchFilters, setSearchFilters] = useSearchState<string>(
    "filters",
    JSON.stringify({
      start_date: moment().format("DD.MM.YYYY"),
    })
  );

  useEffect(() => {
    const filters = JSON.parse(searchFilters);
    setFilters(filters);
  }, []);

  const setFiltersLogic = (value: SearchParams) => {
    const filtered = replaceEmpty(value);
    setFilters(filtered);
    setSearchFilters(JSON.stringify(filtered));
  };

  return [filters, setFiltersLogic];
}
