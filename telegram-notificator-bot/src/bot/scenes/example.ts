import { Scenes } from "telegraf";

export const NAME = "example-scene";

const scene = new Scenes.BaseScene(NAME);
const regex = new RegExp(/.*/);

scene.command(["start", "cancel"], (ctx) => {
  return ctx.scene.leave();
});

scene.enter(async (ctx) => {
  ctx.reply("Hello");
});

scene.action(regex, (ctx) => {
  ctx.answerCbQuery();
  // return ctx.scene.enter(...);
});

export default scene;
