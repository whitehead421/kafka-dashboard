<template>
  <div>
    <span
      v-if="connected"
      class="text-sm ml-2 px-2 py-1 bg-green-200 text-green-800 rounded-full"
      >Connected</span
    >
    <span
      v-else
      class="text-sm ml-2 px-2 py-1 bg-red-200 text-red-800 rounded-full"
      >Disconnected</span
    >
  </div>
</template>

<script setup lang="ts">
interface IProps {
  connected: boolean;
}

defineProps<IProps>();
</script>
