<script setup lang="ts">
import type { NavigationMenuItem } from "@nuxt/ui";

const NavigationItems = ref<NavigationMenuItem[]>([
  {
    label: "Home",
    icon: "i-lucide-home",
    to: "/",
  },
  {
    label: "Blogposts",
    icon: "i-lucide-book-open",
    to: "/blogposts",
  },
]);
</script>
<template>
  <UNavigationMenu :items="NavigationItems" />
  <slot />
  <div class="mt-5 flex justify-center items-center">
    <p>2026</p>
  </div>
</template>
