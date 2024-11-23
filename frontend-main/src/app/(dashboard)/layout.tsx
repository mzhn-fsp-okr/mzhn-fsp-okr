import { Title, TitleProvider } from "@/components/providers/title";
import DashboardSidebar from "@/components/sidebar";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import CurrentTime from "@/components/util/current-time";
import AiDialog from "./components/ai-dialog";

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
            <CurrentTime className="ml-auto hidden text-sm text-muted-foreground sm:block" />
          </div>
          <div className="p-4">{children}</div>
          <AiDialog />
        </main>
      </SidebarProvider>
    </TitleProvider>
  );
}
