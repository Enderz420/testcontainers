<script setup lang="ts">
import Blogpost from "~/components/Blogpost.vue";

const { getAllBlogposts } = useBlogpost();

const data = await getAllBlogposts();
</script>
<template>
  <h1>Blogposts</h1>
  <Blogpost
    v-for="(post, index) in data.results"
    :item="post"
    :index="index"
    :key="post.id"
    :id="post.id"
    :title="post.title"
    :content="post.content"
    :updated_at="post.updated_at"
    :created_at="post.created_at"
    :created_by="post.created_by"
  />
</template>
