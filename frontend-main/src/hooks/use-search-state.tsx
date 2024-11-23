import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useState } from "react";

export default function useSearchState<T extends string, V = T>(
  key: string,
  initialValue: T,
  transformer: (value: T) => V = (value) => value as unknown as V
): [V, (value: V) => void] {
  const params = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const currentValue = (params.get(key) as T) || initialValue;
  const transformedValue = transformer(currentValue);
  const [state, setState] = useState<V>(transformedValue);

  const setSearchState = (value: V) => {
    const urlParams = new URLSearchParams(params);
    urlParams.set(key, value + "");
    setState(value);
    router.replace(`${pathname}?${urlParams.toString()}`);
  };

  return [state, setSearchState];
}
