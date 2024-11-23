import { cn } from "@/lib/utils";
import {
  ToggleGroup,
  ToggleGroupItem,
  ToggleGroupSingleProps,
} from "@radix-ui/react-toggle-group";

export interface EventViewRadioProps
  extends Omit<ToggleGroupSingleProps, "value" | "type"> {
  value: string | null | undefined;
  setValue: (value: string | null | undefined) => void;
}

export default function EventViewRadio({
  value,
  setValue,
  className,
  ...props
}: EventViewRadioProps) {
  return (
    <ToggleGroup
      type="single"
      value={value ?? ""}
      onValueChange={(e) => {
        if (!e) setValue(value);
        else setValue(e);
      }}
      className={cn("space-x-2", className)}
      {...props}
    >
      <ToggleGroupItem
        value="default"
        className="hover:underline aria-checked:text-blue-500"
      >
        сетка
      </ToggleGroupItem>
      <ToggleGroupItem
        value="list"
        className="hover:underline aria-checked:text-blue-500"
      >
        список
      </ToggleGroupItem>
      <ToggleGroupItem
        value="calendar"
        className="hover:underline aria-checked:text-blue-500"
      >
        календарь
      </ToggleGroupItem>
      <ToggleGroupItem
        value="map"
        className="hover:underline aria-checked:text-blue-500"
      >
        карта
      </ToggleGroupItem>
    </ToggleGroup>
  );
}
