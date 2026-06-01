import winston from "winston";

export const logger = winston.createLogger({
  level: "info",
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.json(),
  ), //   defaultMeta: { service: "test-service" },
  //   transports: [new OpenTelemetryTransportV3()],
  transports: [new winston.transports.Console()],
});
