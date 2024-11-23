import { Action } from ".";
import { notificationService } from "../../grpc/notification-service";

export const exampleAction: Action = async (bot) => {
  bot.start(async (ctx) => {
    const key = ctx.payload;
    const chatId = String(ctx.chat.id);

    const res = await notificationService.linkTelegram(chatId, key);

    await ctx.reply(res);
  });
};
