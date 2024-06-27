<template>
  <div
    class="md:w-80 w-52 flex justify-between p-2 border rounded-md text-slate-800"
    :class="[flashClass]"
  >
    <div class="flex justify-between items-center">
      <div class="flex gap-2 items-center">
        <img
          :src="imageLoader(ticker.Symbol)"
          :alt="ticker.Symbol"
          class="w-8 h-8"
        />
        <div class="flex flex-col">
          <div class="flex items-center gap-0.5">
            <h2 class="font-bold md:text-xl text-md">{{ ticker.Symbol }}</h2>
            <span class="font-light text-gray-400 md:block hidden">USDT</span>
          </div>
          <span class="text-xs text-gray-400">Binance</span>
        </div>
      </div>
    </div>
    <div class="flex flex-col items-end leading-none">
      <span class="md:text-lg font-bold">
        {{ moneyFormat(ticker.Price) }}
      </span>
      <span
        class="text-xs font-semibold p-1 px-2 rounded-full w-min"
        :class="{
          'text-green-600': ticker.ChangePct > 0,
          'text-red-600': ticker.ChangePct < 0,
          'text-slate-600': ticker.ChangePct === 0,
        }"
      >
        <span
          >{{ ticker.ChangePct > 0 ? "+" : ""
          }}{{ ticker.ChangePct.toFixed(2) }}%</span
        >
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { moneyFormat } from "@/helpers/moneyFormat";
import { imageLoader } from "@/helpers/imageLoader";
import { computed } from "vue";
interface IProps {
  ticker: ProcessedTicker;
}

const props = defineProps<IProps>();

const flashClass = computed(() => {
  if (props.ticker.PriceChange > 0) {
    return "flash-green";
  } else if (props.ticker.PriceChange < 0) {
    return "flash-red";
  } else {
    return "";
  }
});
</script>

<style scoped>
.ticker-card {
  transition: background-color 1s;
}

@keyframes flash {
  0% {
    background-color: transparent;
  }
  50% {
    background-color: inherit;
  }
  100% {
    background-color: transparent;
  }
}

.flash-green {
  animation: flash 0.3s;
  background-color: rgba(193, 255, 193, 0.6);
  border-color: rgba(193, 255, 193, 1);
}

.flash-red {
  animation: flash 0.3s;
  background-color: rgba(255, 193, 193, 0.6);
  border-color: rgba(255, 193, 193, 1);
}
</style>
