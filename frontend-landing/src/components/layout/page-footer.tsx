import { cn } from "@/lib/utils";
import { HTMLAttributes } from "react";

import logo from "@/assets/images/brand-alt.png";
import Image from "next/image";

export interface PageFooterProps extends HTMLAttributes<HTMLDivElement> {}

export default function PageFooter({ className, ...props }: PageFooterProps) {
  return (
    <footer
      className={cn(
        "flex flex-col items-center gap-4 pb-4 pt-8 text-white",
        className
      )}
      {...props}
    >
      <Image src={logo} alt="Logo" width={128} height={128} />
      <div className="text-center">
        <span className="text-white/50">Разработано c &hearts; </span>
        <span className="bg-team bg-clip-text text-transparent">mzhn-team</span>
      </div>
    </footer>
  );
}
