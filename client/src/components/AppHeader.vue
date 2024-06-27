<template>
  <div class="flex p-4 justify-between items-center">
    <div class="flex items-center gap-2">
      <div>
        <img src="/img/kafka.svg" alt="Kafka" class="w-8 h-8" />
      </div>
      <h1 class="font-semibold text-xl">Kafka Dashboard</h1>
    </div>
    <SocketStatus :connected="connected" @click="emits('close-socket')" />
  </div>
</template>

<script setup lang="ts">
import SocketStatus from "./SocketStatus.vue";

interface IProps {
  connected: boolean;
}

interface IEmits {
  (event: "close-socket"): void;
}

const emits = defineEmits<IEmits>();
defineProps<IProps>();
</script>
