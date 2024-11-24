import { ChannelCredentials } from "@grpc/grpc-js";
import { GrpcTransport } from "@protobuf-ts/grpc-transport";
import { LinkTelegramRequest } from "./proto/notifications-service"; // Adjust path based on generated file location
import { NotificationServiceClient } from "./proto/notifications-service.client"; // Adjust path based on generated file location

class NotificationService {
  client: NotificationServiceClient;

  constructor() {
    const host = `${process.env.NOTIFICATION_SERVICE_HOST}:${process.env.NOTIFICATION_SERVICE_PORT}`;
    // Define the gRPC server URL
    const transport = new GrpcTransport({
      host: host,
      channelCredentials: ChannelCredentials.createInsecure(),
    });

    // Create a client for the NotificationService
    this.client = new NotificationServiceClient(transport);
  }

  async linkTelegram(chatId: string, token: string) {
    const request: LinkTelegramRequest = {
      chatId,
      token,
    };

    let response: any = undefined;
    try {
      response = await this.client.linkTelegram(request);
      console.log("Successfully linked Telegram:", response);
      return "ok";
    } catch (error) {
      if (error.code === "RESOURCE_EXHAUSTED") {
        console.error("failed to link telegram: token expired");
        return "token expired";
      }
      console.error("Failed to link Telegram:", error);
    }
  }
}

export const notificationService = new NotificationService();
