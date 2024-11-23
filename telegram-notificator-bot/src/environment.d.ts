declare global {
  namespace NodeJS {
    interface ProcessEnv {
      BOT_TOKEN: string;
      API_URL: string;
      AMQP_NOTIFICATION_EXCHANGE: string;
      AMQP_TELEGRAM_QUEUE: string;

      AMQP_HOST: string;
      AMQP_PORT: number;
      AMQP_USER: string;
      AMQP_PASS: string;

      NOTIFICATION_GRPC_HOST: string;
      NOTIFICATION_GRPC_PORT: number;
    }
  }
}

export {};
