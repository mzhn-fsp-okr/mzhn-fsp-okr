import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu";
import { cn } from "@/lib/utils";
import { User } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { HTMLAttributes } from "react";
import { Button } from "../ui/button";

import logo from "@/assets/images/brand.png";

export interface PageHeaderProps extends HTMLAttributes<HTMLDivElement> {}

export default function PageHeader({ className, ...props }: PageHeaderProps) {
  return (
    <header
      className={cn(
        "flex w-full items-center justify-between rounded bg-amber-50/90 px-8 py-4",
        className
      )}
      {...props}
    >
      <Image src={logo} alt="Logo" width={128} height={128} />
      <NavigationMenu>
        <NavigationMenuList>
          <MenuItem title="Главная" href="/" />
          <MenuItem title="Календарь" href="/" />
        </NavigationMenuList>
      </NavigationMenu>
      <div>
        <Button variant="ghost">
          <User />
          Войти
        </Button>
      </div>
    </header>
  );
}

function MenuItem({ title, href }: { title: string; href: string }) {
  return (
    <NavigationMenuItem>
      <Link href={href} legacyBehavior passHref>
        <NavigationMenuLink
          className={cn(
            navigationMenuTriggerStyle(),
            "bg-transparent hover:bg-transparent hover:underline focus:bg-transparent focus:underline"
          )}
        >
          {title}
        </NavigationMenuLink>
      </Link>
    </NavigationMenuItem>
  );
}
