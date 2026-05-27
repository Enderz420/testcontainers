<script setup lang="ts">
const props = defineProps<Blogpost>();
</script>
<template>
  <div class="p-8 bg-purple-700/40 text-white">
    <span class="text-2xl">{{ props.title }}</span>
    <p class="text-xl">{{ props.content }}</p>

    <div class="mx-auto p-5">
      <p>Metadata</p>
      <p>Created at {{ props.created_at }}</p>
      <p>Created by {{ props.created_by }}</p>
      <p>Last updated {{ props.updated_at }}</p>
      <p>UUID: {{ props.id }}</p>
    </div>
  </div>
</template>
