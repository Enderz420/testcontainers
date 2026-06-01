import winston from "winston";

const consoleFormat = winston.format.combine(
  winston.format.colorize(),
  winston.format.timestamp({ format: "DD-MM-YYYY HH:mm:ss" }),
  winston.format.printf(({ timestamp, level, message }) => {
    return `${timestamp} ${level}: ${message}`;
  }),
);

export const logger = winston.createLogger({
  level: "error",
  format: winston.format.json(),
  //   defaultMeta: { service: "test-service" },
  //   transports: [new OpenTelemetryTransportV3()],
  transports: [
    new winston.transports.Console({
      format: consoleFormat,
      level: "info",
      silent: false,
    }),
  ],
});
