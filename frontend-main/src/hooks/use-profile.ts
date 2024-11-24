import { me } from "@/api/auth";
import { useQuery } from "@tanstack/react-query";

export default function useProfile(throwOnFail: boolean = true) {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["me"],
    queryFn: me,
    throwOnError(error, query) {
      if (error.message === "NEXT_REDIRECT") return throwOnFail;
      return true;
    },
  });
  return { data, isLoading, isError };
}
