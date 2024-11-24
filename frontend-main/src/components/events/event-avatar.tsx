import { SportEvent } from "@/api/events";
import { cn } from "@/lib/utils";
import { Avatar, AvatarProps } from "@radix-ui/react-avatar";
import { AvatarImage } from "../ui/avatar";

export interface EventAvatarProps extends AvatarProps {
  event: SportEvent;
}

export default function EventAvatar({
  event,
  className,
  ...props
}: EventAvatarProps) {
  const src = `https://api.dicebear.com/9.x/glass/svg?seed=${event.id}`;
  return (
    <Avatar className={cn("size-10 rounded", className)} {...props}>
      <AvatarImage src={src} className="rounded" />
    </Avatar>
  );
}
