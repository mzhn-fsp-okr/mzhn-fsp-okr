"use client";

import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import {
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@/components/ui/sidebar";
import { ChevronRight, LucideIcon } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { PropsWithChildren, ReactElement, ReactNode, useMemo } from "react";

export interface DashboardLinkProps {
  title: string;
  icon?: ReactElement<LucideIcon>;
  link: string;
  sub?: boolean;
  after?: ReactNode;
  exact?: boolean;
}

export interface DashboardMenuProps extends PropsWithChildren {
  title: string;
  icon?: ReactElement<LucideIcon>;
  link?: string;
}

export function DashboardButton({
  title,
  link,
  icon,
  sub,
  after,
  exact,
}: DashboardLinkProps) {
  const Wrapper = sub ? SidebarMenuSubItem : SidebarMenuItem;
  const Button = sub ? SidebarMenuSubButton : SidebarMenuButton;

  const path = usePathname();
  const isActive = useMemo(() => {
    if (exact) return path === link;
    return path.startsWith(link);
  }, [path, link, exact]);

  return (
    <Wrapper>
      <Button isActive={isActive} asChild>
        <Link href={link}>
          {icon}
          <span>{title}</span>
          {after}
        </Link>
      </Button>
    </Wrapper>
  );
}

export function DashboardMenu({
  children,
  link,
  ...props
}: DashboardMenuProps) {
  const path = usePathname();
  const open = useMemo(() => {
    if (!link) return false;
    return path.startsWith(link);
  }, [path, link]);

  return (
    <Collapsible className="group/collapsible" defaultOpen={open}>
      <CollapsibleTrigger asChild>
        <SidebarMenuButton>
          {props.icon}
          <span>{props.title}</span>
          <ChevronRight className="ml-auto transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90" />
        </SidebarMenuButton>
      </CollapsibleTrigger>
      <CollapsibleContent>
        <SidebarMenuSub>{children}</SidebarMenuSub>
      </CollapsibleContent>
    </Collapsible>
  );
}
