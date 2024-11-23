"use client";

import { sportSubscribe, sportUnsubscribe } from "@/api/subcribes";
import { Button } from "@/components/ui/button";
import { useMutation, useQuery } from "@tanstack/react-query";
import { BellMinus, BellPlus, LoaderCircle } from "lucide-react";
import { useEffect, useMemo, useState } from "react";

export function SportEntry({ id, name }: { id: string; name: string }) {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["subscribe-sports"],
    queryFn: () => [],
  });
  const [initialized, setInitialized] = useState(false);
  const [subscribed, setSubscribed] = useState(false);

  const subscribeMutation = useMutation({
    mutationFn: async () => {
      if (subscribed) await sportUnsubscribe(id);
      else await sportSubscribe(id);
      setSubscribed(!subscribed);
    },
  });

  useEffect(() => {
    if (!data) return;
    setSubscribed(data!.find((d) => d.id === id) != undefined);
    setInitialized(true);
  }, [data]);
  const disabled = useMemo(() => {
    return isLoading || subscribeMutation.isPending;
  }, [subscribeMutation.isPending, isLoading]);

  const onClick = async () => {
    if (!initialized) return;
    subscribeMutation.mutate();
  };

  if (isError) return <></>;

  return (
    <li className="flex items-center justify-between rounded border px-4 py-8">
      <p>{name}</p>
      <Button
        size="icon"
        variant="secondary"
        onClick={onClick}
        disabled={disabled}
      >
        {disabled ? (
          <LoaderCircle className="animate-spin" />
        ) : subscribed ? (
          <BellMinus />
        ) : (
          <BellPlus />
        )}
      </Button>
    </li>
  );
}
