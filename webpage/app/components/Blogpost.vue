<script setup lang="ts">
const props = defineProps<Blogpost>();
</script>
<template>
  <div>
    <span>{{ props.title }}</span>
    <p>{{ props.content }}</p>

    <div>
      <p>Metadata</p>
      <p>Created at {{ props.created_at }}</p>
      <p>Created by {{ props.created_by }}</p>
      <p>Last updated {{ props.updated_at }}</p>
      <p>UUID: {{ props.id }}</p>
    </div>
  </div>
</template>
