import { OpenTelemetryTransportV3 } from "@opentelemetry/winston-transport";
import winston from "winston";

export const logger = winston.createLogger({
  level: "info",
  format: winston.format.json(),
  //   defaultMeta: { service: "test-service" },
  //   transports: [new OpenTelemetryTransportV3()],
  transports: [new winston.transports.Console()],
});
