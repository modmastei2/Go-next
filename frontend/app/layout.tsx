import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Go-next Shop - Full-stack Application",
  description: "A full-stack shop application built with Next.js and Golang",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="antialiased">
        {children}
      </body>
    </html>
  );
}
