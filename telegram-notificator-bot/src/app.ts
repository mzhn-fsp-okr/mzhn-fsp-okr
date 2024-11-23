import "dotenv/config";
import { setupAmqp } from "./amqp";
import * as bot from "./bot";
import db from "./db";
import logger from "./util/log";
(async () => {
  logger.info("Init phase start");
  db.pragma("journal_mode = WAL");

  const env = process.env;
  const tgBot = bot.setup(env);
  setupAmqp(env);
  logger.info("Init phase end");

  // Enable graceful stop
  process.once("SIGINT", () => tgBot.stop("SIGINT"));
  process.once("SIGTERM", () => tgBot.stop("SIGTERM"));

  tgBot.launch();
})();
