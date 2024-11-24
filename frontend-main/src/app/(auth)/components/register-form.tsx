"use client";

import * as Card from "@/components/ui/card";
import * as Form from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { LoadingButton } from "@/components/ui/loading";
import { useToast } from "@/hooks/use-toast";
import registerForm, { RegisterData } from "@/lib/forms/register";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { HTMLAttributes } from "react";
import { useForm, useFormContext } from "react-hook-form";
import { register } from "../actions/register";

export default function RegisterForm() {
  return (
    <Card.Card className="w-full md:w-auto">
      <Card.CardHeader>
        <Card.CardTitle>Регистрация</Card.CardTitle>
        <Card.CardDescription>
          Добро пожаловать! Представьтесь.
        </Card.CardDescription>
      </Card.CardHeader>
      <RegisterFormProvider>
        <Card.CardContent>
          <RegisterFormContent />
        </Card.CardContent>
        <Card.CardFooter>
          <RegisterButton />
        </Card.CardFooter>
      </RegisterFormProvider>
    </Card.Card>
  );
}

function RegisterFormProvider({
  children,
  ...props
}: HTMLAttributes<HTMLFormElement>) {
  const form = useForm<RegisterData>({
    resolver: zodResolver(registerForm),
    defaultValues: {
      firstName: "",
      lastName: "",
      middleName: "",
      email: "",
      password: "",
      passwordConfirm: "",
    },
    mode: "onBlur",
  });
  const { toast } = useToast();
  const router = useRouter();

  const handleSubmit = async (data: RegisterData) => {
    const [ok, errorMessage] = await register(data);
    const message = ok ? `Приветствуем, ${data.firstName}!` : errorMessage!;
    toast({ title: message });
    if (ok) router.push("/");
  };

  return (
    <Form.Form {...form}>
      <form onSubmit={form.handleSubmit(handleSubmit)} {...props}>
        {children}
      </form>
    </Form.Form>
  );
}

function RegisterFormContent() {
  const form = useFormContext();

  return (
    <div className="grid grid-cols-1 gap-4 md:grid-cols-2 md:gap-8">
      <div className="flex flex-col gap-2 md:min-w-80">
        <Form.FormField
          control={form.control}
          name="lastName"
          render={({ field }) => (
            <Form.FormItem>
              <Form.FormLabel>Фамилия</Form.FormLabel>
              <Form.FormControl>
                <Input placeholder="Иванов" {...field} />
              </Form.FormControl>
              <Form.FormMessage />
            </Form.FormItem>
          )}
        />
        <Form.FormField
          control={form.control}
          name="firstName"
          render={({ field }) => (
            <Form.FormItem>
              <Form.FormLabel>Имя</Form.FormLabel>
              <Form.FormControl>
                <Input placeholder="Иван" {...field} />
              </Form.FormControl>
              <Form.FormMessage />
            </Form.FormItem>
          )}
        />
        <Form.FormField
          control={form.control}
          name="middleName"
          render={({ field }) => (
            <Form.FormItem>
              <Form.FormLabel>Отчество</Form.FormLabel>
              <Form.FormControl>
                <Input placeholder="Иванович" {...field} />
              </Form.FormControl>
              <Form.FormMessage />
            </Form.FormItem>
          )}
        />
      </div>
      <div className="flex flex-col gap-2 md:min-w-80">
        <Form.FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <Form.FormItem>
              <Form.FormLabel>Почта</Form.FormLabel>
              <Form.FormControl>
                <Input placeholder="user@site.ru" {...field} />
              </Form.FormControl>
              <Form.FormMessage />
            </Form.FormItem>
          )}
        />
        <Form.FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <Form.FormItem>
              <Form.FormLabel>Пароль</Form.FormLabel>
              <Form.FormControl>
                <Input type="password" placeholder="●●●●●●●●" {...field} />
              </Form.FormControl>
              <Form.FormMessage />
            </Form.FormItem>
          )}
        />
        <Form.FormField
          control={form.control}
          name="passwordConfirm"
          render={({ field }) => (
            <Form.FormItem>
              <Form.FormLabel>Подтвердите пароль</Form.FormLabel>
              <Form.FormControl>
                <Input type="password" placeholder="●●●●●●●●" {...field} />
              </Form.FormControl>
              <Form.FormMessage />
            </Form.FormItem>
          )}
        />
      </div>
    </div>
  );
}

function RegisterButton() {
  const { formState } = useFormContext();
  return (
    <LoadingButton
      className="w-full"
      type="submit"
      loading={formState.isSubmitting}
    >
      Регистрация
    </LoadingButton>
  );
}
