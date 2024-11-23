import { SQLite } from "@telegraf/session/sqlite";
import { session, Telegraf } from "telegraf";
import db from "../db";
import logger from "../util/log";
import registerActions from "./actions";
import stage from "./scenes";

let tgBot: Telegraf | null = null;

/** Create Telegraf bot instance */
export function setup(env: NodeJS.ProcessEnv) {
  const bot = new Telegraf(env.BOT_TOKEN);

  const store = SQLite({ database: db });
  bot.use(session({ store }));
  bot.use(stage.middleware());

  registerActions(bot);
  logger.info("Telegraf bot initialized");

  tgBot = bot;
  return bot;
}

export function get() {
  if (!tgBot) throw new Error("Bot is not initialized");
  return tgBot;
}
