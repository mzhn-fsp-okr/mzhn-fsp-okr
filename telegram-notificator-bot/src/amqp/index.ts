import amqp from "amqplib";
import { get } from "../bot";

import { formatMessage } from "./format_event";

type Message = {
  receiver: string;
  event: Event;
  eventType: "upcoming" | "new";
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

  await ch.assertQueue(queue, { durable: true });
  await ch.consume(
    queue,
    async (msg) => {
      const text = msg.content.toString();
      const body: Message = JSON.parse(text);

      const bot = get();
      try {
        console.log("sending message", body.receiver);
        await bot.telegram.sendMessage(body.receiver, formatMessage(body.event), { parse_mode: "Markdown" });
      } catch (error) {
        console.error("error while sending message", error);
      }
    },
    { noAck: true }
  );
};
