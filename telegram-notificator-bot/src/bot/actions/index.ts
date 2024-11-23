import { Telegraf } from "telegraf";
import { exampleAction } from "./example";

export type Action = (bot: Telegraf) => Promise<void>;

const actions: Action[] = [exampleAction];

export default function registerActions(bot: Telegraf) {
  actions.forEach((cb) => cb(bot));
}
