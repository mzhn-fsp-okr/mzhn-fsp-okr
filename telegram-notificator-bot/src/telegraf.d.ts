import "telegraf";
import { SceneContextScene } from "telegraf/typings/scenes";

declare module "telegraf" {
  interface Context {
    session?: Record<string, any>;
    scene?: SceneContextScene<any, any>;
  }
}
