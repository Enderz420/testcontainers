export default defineNitroPlugin((nitroApp) => {
  nitroApp.hooks.hook("request", (event) => {
    logger.info("request received", {
      method: event.method,
      path: event.path,
    });
  });

  nitroApp.hooks.hook("afterResponse", (event) => {
    logger.info("request completed", {
      method: event.method,
      path: event.path,
      status: event.node.res.statusCode,
    });
  });
});
