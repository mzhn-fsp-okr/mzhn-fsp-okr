import { Title, TitleProvider } from "@/components/providers/title";
import DashboardSidebar from "@/components/sidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import CurrentTime from "@/components/util/current-time";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <TitleProvider>
      <SidebarProvider>
        <DashboardSidebar />
        <main className="w-full">
          <div className="flex items-center gap-4 border-b px-4 py-2">
            <div className="border-r pr-2">
              <SidebarTrigger />
            </div>
            <Title className="text-sm" />
            <CurrentTime className="text-muted-foreground ml-auto hidden text-sm sm:block" />
          </div>
          <div className="p-4">{children}</div>
        </main>
      </SidebarProvider>
    </TitleProvider>
  );
}
