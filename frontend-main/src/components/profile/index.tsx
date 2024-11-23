import { getAvatar, getFullName, getRoleFriendly, User } from "@/api/user";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/dialog";
import { DialogProps } from "@radix-ui/react-dialog";
import Link from "next/link";
import { Avatar, AvatarImage } from "../ui/avatar";
import { Button } from "../ui/button";

export interface ProfileProps extends DialogProps {
  user: User;
}

export default function Profile({ user, ...props }: ProfileProps) {
  const fullName = getFullName(user) ?? "Пользователь";
  const avatar = getAvatar(user);
  const role = getRoleFriendly(user);

  return (
    <Dialog {...props}>
      <DialogContent className="flex flex-col items-center gap-2">
        <DialogTitle className="sr-only">Профиль пользователя</DialogTitle>
        <Avatar className="size-32">
          <AvatarImage src={avatar} />
        </Avatar>
        <div className="text-center">
          <h1 className="text-lg font-bold">{fullName}</h1>
          <h2 className="text-sm text-muted-foreground">{role}</h2>
          <h2 className="text-sm text-muted-foreground">{user.email}</h2>
        </div>
        <div className="w-full space-y-2 py-4">
          <Button className="w-full" asChild>
            <Link href="/profile">В личный кабинет</Link>
          </Button>
          <Button className="w-full" variant="secondary" asChild>
            <Link href="/sign-out" prefetch={false}>
              Выход
            </Link>
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
}
