<script lang="ts" setup>
import BackButton from '~icons/material-symbols/arrow-back'
import { useRouter } from 'vue-router'

const router = useRouter()
</script>

<template>
  <div class="relative h-screen p-4">
    <Button type="button" @click="router.back()" class="absolute top-4 left-4">
      <BackButton />
    </Button>
    <div class="flex flex-col items-center justify-center h-full">
      <h1>Settings</h1>
    </div>
  </div>
</template>

<style></style>