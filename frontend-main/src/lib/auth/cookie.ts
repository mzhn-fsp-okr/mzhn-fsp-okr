import Cookies from "js-cookie";
import type {
  ResponseCookie,
  ResponseCookies,
} from "next/dist/compiled/@edge-runtime/cookies";

export class ClientCookieStorage implements ResponseCookies {
  get size(): number {
    return document.cookie.split(";").length;
  }
  get(...args: [name: string] | [ResponseCookie]): ResponseCookie | undefined {
    const value = args[0];
    if (typeof value !== "string") {
      throw new Error("Method not implemented.");
    }
    const cookie = Cookies.get(value);
    if (!cookie) return undefined;
    return { name: value, value: cookie };
  }
  getAll(): ResponseCookie[] {
    const cookie = Cookies.get();
    return Object.entries(cookie).map((v) => ({ name: v[0], value: v[1] }));
  }
  has(name: string): boolean {
    return Cookies.get(name) !== undefined;
  }
  set(
    ...args:
      | [key: string, value: string, cookie?: ResponseCookie]
      | [options: ResponseCookie]
  ): this {
    const key = args[0] as string;
    const value = args[1] as string;
    const options = args[2] as ResponseCookie | undefined;
    // @ts-expect-error Temporary disable, need to fix
    Cookies.set(key, value, options);
    return this;
  }
  delete(
    ...args:
      | [key: string]
      | [options: Omit<ResponseCookie, "value" | "expires">]
  ): this {
    const value = args[0];
    if (typeof value !== "string") {
      throw new Error("Method not implemented.");
    }
    Cookies.remove(value);
    return this;
  }
  clear(): this {
    this.getAll().forEach((cookie) => this.delete(cookie.name));
    return this;
  }
  toString(): string {
    return document.cookie;
  }
  [Symbol.iterator](): IterableIterator<[string, ResponseCookie]> {
    function* iterator(arr: ResponseCookie[]) {
      for (const item of arr) {
        yield { [item.name]: item };
      }
    }
    return iterator(this.getAll()) as IterableIterator<
      [string, ResponseCookie]
    >;
  }
}

export function cookies(): ResponseCookies {
  if (typeof window === "undefined") {
    // eslint-disable-next-line @typescript-eslint/no-require-imports
    const { cookies: c } = require("next/headers");
    return c();
  } else {
    return new ClientCookieStorage();
  }
}
