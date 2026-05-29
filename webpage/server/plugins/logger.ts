export default defineNitroPlugin((nitroApp) => {
  nitroApp.hooks.hook("beforeResponse", (event) => {
    const headers = getHeaders(event);
    logger.info("on before response", event.path, { headers });
  });
  nitroApp.hooks.hook("afterResponse", (event) => {
    const headers = getHeaders(event);
    logger.info("on after response", event.path, { headers });
  });
  nitroApp.hooks.hook("request", (event) => {
    logger.info("on request", event.path);
  });
});
