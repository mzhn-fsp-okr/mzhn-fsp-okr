import ReactQueryProvider from "@/components/providers/react-query";
import { Toaster } from "@/components/ui/toaster";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";

const font = Inter({ subsets: ["cyrillic", "latin"] });

export const metadata: Metadata = {
  title: {
    template: "%s - Твой Спорт",
    default: "Твой Спорт",
  },
  description: "Твой Спорт - Спортивный график",
  icons: {
    apple: "/apple-touch-icon.png",
  },
  manifest: "/site.webmanifest",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru">
      <body className={`${font.className} bg-zinc-900 antialiased`}>
        <ReactQueryProvider>{children}</ReactQueryProvider>
        <Toaster />
      </body>
    </html>
  );
}
