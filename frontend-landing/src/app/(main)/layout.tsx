import PageFooter from "@/components/layout/page-footer";
import PageHeader from "@/components/layout/page-header";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <section className="grid min-h-screen w-full grid-cols-1 grid-rows-[80px_1fr_auto] divide-y divide-white/5 py-4 md:grid-cols-[100px_1fr_100px] md:grid-rows-[auto_1fr_auto]">
      <PageHeader className="row-start-1 border-x md:col-start-2" />
      <main className="row-start-2 border-x border-white/5 py-4 md:col-start-2">
        {children}
      </main>
      <PageFooter className="row-start-3 border-x border-white/5 md:col-start-2" />
    </section>
  );
}
