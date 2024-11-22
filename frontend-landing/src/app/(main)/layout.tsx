import PageFooter from "@/components/layout/page-footer";
import PageHeader from "@/components/layout/page-header";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <section className="grid min-h-screen w-full divide-y divide-white/5 py-4 sm:grid-cols-[100px_1fr_100px] sm:grid-rows-[auto_1fr_auto]">
      <PageHeader className="col-start-2 row-start-1 border-x" />
      <main className="col-start-2 row-start-2 border-x border-white/5 py-4">
        {children}
      </main>
      <PageFooter className="col-start-2 row-start-3 border-x border-white/5" />
    </section>
  );
}
