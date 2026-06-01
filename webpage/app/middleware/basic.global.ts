export default defineNuxtRouteMiddleware((to, from) => {
  logger.info("user navigation", {
    to: to.path,
    from: from.path,
  });
});
