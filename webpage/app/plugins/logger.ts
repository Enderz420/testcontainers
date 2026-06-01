export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.config.errorHandler = (error, instance, info) => {
    logger.error("error:", {
      error,
      instance,
      info,
    });
  };
});
