<template>
  <div class="overflow-auto flex-grow-[1] max-w-full relative">
    <table
      class="whitespace-nowrap overflow-auto w-[98vw] mx-auto border border-gray-100"
    >
      <thead class="shadow-lg">
        <tr class="border-b border-gray-100">
          <th
            v-for="column in props.columns"
            :key="column.index"
            class="px-2 py-4 text-left text-xs font-medium text-gray-500"
          >
            {{ column.title }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="ticker in props.data"
          :key="ticker.Symbol"
          class="grid-row border-b border-gray-100"
          :class="{
            'flash-green': ticker.PriceChange > 0,
            'flash-red': ticker.PriceChange < 0,
            'bg-white': ticker.PriceChange === 0,
          }"
        >
          <td class="px-2 py-4 text-sm text-gray-900 w-96">
            <div class="flex gap-2 items-center">
              <img
                :src="imageLoader(ticker.Symbol)"
                :alt="ticker.Symbol"
                class="w-8 h-8"
              />
              <div class="flex flex-col gap-0.5">
                <span>
                  {{ ticker.Symbol.replace("USDT", "") }}
                </span>
                <span class="text-xs text-gray-500">
                  {{ Coins[ticker.Symbol as keyof typeof Coins] }}
                </span>
              </div>
            </div>
          </td>
          <td class="px-2 py-4 text-lg text-gray-900 w-96">
            {{ moneyFormat(ticker.Price) }}
          </td>
          <td class="px-2 py-4 text-sm text-gray-900 w-96">
            <span
              class="text-xs p-1 px-2 rounded-full"
              :class="{
                'text-green-800 bg-green-50': ticker.ChangePct > 0,
                'text-red-800 bg-red-50': ticker.ChangePct < 0,
                'text-slate-800 bg-slate-50': ticker.ChangePct === 0,
              }"
            >
              <span>{{ ticker.ChangePct.toFixed(2) }}%</span>
            </span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { moneyFormat } from "@/helpers/moneyFormat";
import { imageLoader } from "@/helpers/imageLoader";
import Coins from "@/assets/data/coins.json";

interface IProps {
  data: ProcessedTicker[];
  columns: Column[];
}

const props = defineProps<IProps>();
</script>

<style scoped>
.grid-row {
  transition: background-color 0.3s;
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
  animation: flash 0.1s;
  background-color: rgba(193, 255, 193, 0.8);
}

.flash-red {
  animation: flash 0.1s;
  background-color: rgba(255, 193, 193, 0.8);
}

table {
  position: relative;
}
th {
  background: white;
  position: sticky;
  top: 0; /* Don't forget this, required for the stickiness */
}
</style>
