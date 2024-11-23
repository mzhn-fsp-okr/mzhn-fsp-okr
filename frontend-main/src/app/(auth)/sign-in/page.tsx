import brandImage from "@/assets/images/brand.png";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";
import SignInForm from "../components/sign-in-form";
import SupportAlert from "../components/support-alert";
import PageReloader from "./page-reloader";

export default function Page() {
  return (
    <div className="flex w-full flex-col items-center gap-4 px-4 md:w-auto">
      <Image src={brandImage} alt="logo" width={200} height={75} />
      <SignInForm />
      <div className="space-x-2">
        <SupportAlert asChild>
          <Button variant="link">Забыли пароль?</Button>
        </SupportAlert>
        <Button variant="link">
          <Link href="/register">Регистрация</Link>
        </Button>
      </div>
      <PageReloader />
    </div>
  );
}
