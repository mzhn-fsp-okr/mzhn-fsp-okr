import { z } from "zod";

const email = z.string().email();
const password = z
  .string()
  .min(4, "Пароль должен быть длиной не менее 4 символов");

const signInForm = z.object({ email, password });

export type SignInData = z.infer<typeof signInForm>;
export default signInForm;
