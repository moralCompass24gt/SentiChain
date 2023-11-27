import "@/styles/globals.css";
import "tailwindcss/tailwind.css";
import "daisyui/dist/full.css";

import type { AppProps } from "next/app";
import { Header } from "@/components/Header";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <main className="relative flex flex-col flex-1">
        <Component {...pageProps} />
      </main>
    </div>
  );
}
