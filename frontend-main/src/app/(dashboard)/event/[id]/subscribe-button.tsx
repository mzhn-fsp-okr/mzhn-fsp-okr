"use client";

import { events, eventSubscribe, eventUnsubscribe } from "@/api/subcribes";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { useMutation, useQuery } from "@tanstack/react-query";
import { CalendarMinus, CalendarPlus, LoaderCircle } from "lucide-react";
import { useEffect, useMemo, useState } from "react";

export default function SubscribeButton({ id }: { id: string }) {
  const { data, isLoading, isError } = useQuery({
    queryKey: ["subscribe-events"],
    queryFn: events,
  });
  const [initialized, setInitialized] = useState(false);
  const [subscribed, setSubscribed] = useState(false);

  const subscribeMutation = useMutation({
    mutationFn: async () => {
      if (subscribed) await eventUnsubscribe(id);
      else await eventSubscribe(id);
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
    <Button
      className={cn("size-12", subscribed && "bg-blue-400")}
      variant={"secondary"}
      onClick={onClick}
      disabled={disabled}
    >
      {disabled ? (
        <LoaderCircle className="animate-spin" />
      ) : subscribed ? (
        <CalendarMinus />
      ) : (
        <CalendarPlus />
      )}
    </Button>
  );
}
