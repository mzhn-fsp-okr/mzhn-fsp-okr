"use client";

import { HTMLAttributes, useEffect, useState } from "react";

export default function CurrentTime(props: HTMLAttributes<HTMLDivElement>) {
  const [time, setTime] = useState(new Date());
  useEffect(() => {
    const id = setInterval(() => setTime(new Date()), 500);
    return () => clearInterval(id);
  }, []);
  return (
    <p {...props} suppressHydrationWarning>
      {formatTime(time)}
    </p>
  );
}

function formatTime(time: Date) {
  const padTo2Digits = (num: number) => String(num).padStart(2, "0");

  const hours = padTo2Digits(time.getHours());
  const minutes = padTo2Digits(time.getMinutes());
  const day = padTo2Digits(time.getDate());
  const month = padTo2Digits(time.getMonth() + 1);
  const year = time.getFullYear();

  return `${hours}:${minutes} ${day}.${month}.${year}`;
}
