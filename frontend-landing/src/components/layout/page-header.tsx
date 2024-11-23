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

import brand from "@/assets/images/brand.png";
import logo from "@/assets/images/logo.png";

export interface PageHeaderProps extends HTMLAttributes<HTMLDivElement> {}

export default function PageHeader({ className, ...props }: PageHeaderProps) {
  return (
    <header
      className={cn(
        "flex w-full items-center justify-between rounded bg-[#e9e9e9] px-8 py-4",
        className
      )}
      {...props}
    >
      <Link href="/">
        <Image
          src={brand}
          alt="Logo"
          width={128}
          height={128}
          className="hidden sm:block"
        />
        <Image
          src={logo}
          alt="Logo"
          width={32}
          height={32}
          className="sm:hidden"
        />
      </Link>
      <NavigationMenu>
        <NavigationMenuList>
          <MenuItem title="Главная" href="/" />
          <MenuItem title="Календарь" href="/calendar" />
        </NavigationMenuList>
      </NavigationMenu>
      <div>
        <Button variant="ghost" asChild>
          <Link href={process.env.NEXT_PUBLIC_CABINET_URL!} target="_blank">
            <User />
            <span className="hidden sm:inline">Войти</span>
          </Link>
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
