<script setup lang="ts">
const props = defineProps<{
  image?: string;
  name: string;
}>();
</script>
<template>
  <div class="p-5 border-2 rounded-2xl text-black bg-pink-300">
    <img v-if="image" :src="props.image" />
    <span v-else> {{ props.name }}</span>
  </div>
</template>
