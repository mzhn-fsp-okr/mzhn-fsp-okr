export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex h-screen flex-col items-center justify-between py-4">
      <div></div>
      {children}
      <div>
        <span className="text-black/50">Разработано c &hearts; </span>
        <span className="bg-team bg-clip-text text-transparent">mzhn-team</span>
      </div>
    </div>
  );
}
