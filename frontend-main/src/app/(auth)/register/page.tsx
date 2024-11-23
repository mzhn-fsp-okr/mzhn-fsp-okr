import brandImage from "@/assets/images/brand.png";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";
import RegisterForm from "../components/register-form";

export default function Page() {
  return (
    <div className="flex w-full flex-col items-center gap-4 px-4 md:w-auto">
      <Image src={brandImage} alt="logo" width={200} height={75} />
      <RegisterForm />
      <div className="space-x-2">
        <Button variant="link">
          <Link href="/sign-in">Уже зарегистрированы?</Link>
        </Button>
      </div>
    </div>
  );
}
