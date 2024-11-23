import { z } from "zod";

const lastName = z
  .string()
  .min(2, "Фамилия должна быть не менее, чем 2 символа");
const firstName = z.string().min(2, "Имя должно быть не менее, чем 2 символа");
const middleName = z
  .string()
  .min(2, "Отчество должно быть не менее, чем 2 символа");
const email = z.string().email();
const password = z
  .string()
  .min(4, "Пароль должен быть длиной не менее 4 символов");
const passwordConfirm = z.string();

const registerForm = z
  .object({
    lastName,
    firstName,
    middleName,
    email,
    password,
    passwordConfirm,
  })
  .refine((data) => data.password == data.passwordConfirm, {
    message: "Пароли должны совпадать",
    path: ["passwordConfirm"],
  });

export type RegisterData = z.infer<typeof registerForm>;
export default registerForm;
