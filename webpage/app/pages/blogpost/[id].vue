<script setup lang="ts">
import Blogpost from "~/components/Blogpost.vue";

const { getBlogpost } = useBlogpost();

const route = useRoute();
const routeId = route.params.id?.toString() ?? "";

const id = ref(routeId);

const data = ref<Blogpost>();

onMounted(async () => {
  data.value = await getBlogpost(id.value);
});
</script>

<template>
  <div v-if="!data">Loading...</div>
  <Blogpost
    v-if="data"
    :id="data.id"
    :title="data.title"
    :content="data.content"
    :updated_at="data.updated_at"
    :created_at="data.created_at"
    :created_by="data.created_by"
  />
</template>
