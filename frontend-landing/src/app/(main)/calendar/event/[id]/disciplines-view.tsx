"use client";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ChevronsDownUp, ChevronsUpDown } from "lucide-react";
import { useState } from "react";

export default function DisciplinesView({
  disciplines,
}: {
  disciplines: string[];
}) {
  const [isOpen, setIsOpen] = useState(false);

  const needCollapsible = disciplines.length > 5;
  const alwaysShown = disciplines.slice(0, 5);
  const collapsed = disciplines.slice(5);

  return (
    <div>
      <ul className="flex flex-wrap gap-2">
        {alwaysShown.map((e, i) => (
          <li key={i}>
            <Badge variant="secondary">{e}</Badge>
          </li>
        ))}
        {isOpen &&
          collapsed.map((e, i) => (
            <li key={i}>
              <Badge variant="secondary">{e}</Badge>
            </li>
          ))}
      </ul>
      {needCollapsible && (
        <div className="py-4">
          <Button variant="secondary" onClick={() => setIsOpen(!isOpen)}>
            {isOpen ? (
              <>
                Свернуть <ChevronsDownUp />
              </>
            ) : (
              <>
                Развернуть (еще {collapsed.length}) <ChevronsUpDown />
              </>
            )}
          </Button>
        </div>
      )}
    </div>
  );
}
