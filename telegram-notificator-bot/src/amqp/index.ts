import amqp from "amqplib";
import { get } from "../bot";

import { TelegramError } from "telegraf";
import { formatNewEventMessage } from "./format-new-event";
import { formatNewSubMessage, formatSportSubMessage } from "./format-new-sub";
import { formatMessage as formatUpcomingMessage } from "./format-upcoming-event";

type Message = {
  receiver: string;
  event: Event;
  type: "upcoming" | "new" | "new-sub";
};

type Event = {
  id: string;
  ekpId: string;
};

export const setupAmqp = async (env: NodeJS.ProcessEnv) => {
  console.log("setup amqp");

  const host = env.AMQP_HOST || "localhost";
  const port = env.AMQP_PORT || 5672;
  const user = env.AMQP_USER || "guest";
  const pass = env.AMQP_PASS || "guest";

  const queue = env.AMQP_TELEGRAM_QUEUE || "telegram-queue";

  const conn = await amqp.connect(`amqp://${user}:${pass}@${host}:${port}`);

  const ch = await conn.createChannel();

  process.once("SIGINT", async () => {
    await ch.close();
    await conn.close();
  });

  let isPaused = false;

  await ch.assertQueue(queue, { durable: true });
  const consume = async (msg) => {
    const text = msg.content.toString();
    const body: Message = JSON.parse(text);

    const bot = get();
    try {
      console.log("sending message", body.receiver, body.type);
      if (body.type == "upcoming") {
        await bot.telegram.sendMessage(
          body.receiver,
          formatUpcomingMessage(body.event),
          { parse_mode: "Markdown" }
        );
      } else if (body.type == "new-sub") {
        await bot.telegram.sendMessage(
          body.receiver,
          body.event
            ? formatNewSubMessage(body.event)
            : // @ts-ignore
              formatSportSubMessage(body.sport),
          { parse_mode: "Markdown" }
        );
      } else if (body.type == "new") {
        await bot.telegram.sendMessage(
          body.receiver,
          formatNewEventMessage(body.event),
          { parse_mode: "Markdown" }
        );
      }

      ch.ack(msg);
    } catch (error) {
      if (error instanceof TelegramError) {
        if (error.code == 429 && !isPaused) {
          const delay = Number(error.message.split("retry after")[1]);

          console.log("rate limit reached, waiting", delay, "sec");
          isPaused = true;

          await ch.cancel("message-sender");

          setTimeout(async () => {
            console.log("resuming consumer...");
            isPaused = false;
            ch.consume(queue, consume, {
              noAck: false,
              consumerTag: "message-sender",
            });
          }, delay * 1000);
        } else {
          ch.nack(msg, false, true);
          console.error("unexprected error telegram", error);
        }
      } else {
        console.error("error while sending message", error);
      }
    }
  };

  ch.consume(queue, consume, {
    noAck: false,
    consumerTag: "message-sender",
  });
};
