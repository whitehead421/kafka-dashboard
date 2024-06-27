<template>
  <div class="flex flex-col gap-2 h-screen">
    <AppHeader :connected @close-socket="closeSok" />
    <p class="px-4 leading-none text-slate-900 font-light text-sm">
      This project demonstrates a real-time dashboard built by Vue, Go & Kafka.
      Also includes a WebSocket server to stream stock prices to the client.
    </p>
    <div class="p-4 flex flex-wrap gap-x-4 gap-y-2 justify-center">
      <TickerCard
        v-for="ticker in sortedTickers"
        :key="ticker.Symbol"
        :ticker="ticker"
      />
    </div>
    <CurrencyGrid v-if="connected" :data="sortedTickers" :columns="columns" />
    <AppFooter />
  </div>
</template>

<script setup lang="ts">
import AppHeader from './components/AppHeader.vue';
import AppFooter from './components/AppFooter.vue';
import TickerCard from './components/TickerCard.vue';
import { computed, ref } from 'vue';
import CurrencyGrid from './components/CurrencyGrid.vue';

const connected = ref(false);
const tickers = ref<ProcessedTicker[]>([]);

const columns: Column[] = [
  { title: 'Symbol', index: 'Symbol' },
  { title: 'Last Price', index: 'Price' },
  { title: 'Change', index: 'Change' },
];

// Establish WebSocket connection
const socket = new WebSocket('ws://localhost:8080/ws');

const closeSok = () => {
  socket.close();
};

socket.onopen = () => {
  console.log('WebSocket connected');
  connected.value = true;
};

socket.onclose = () => {
  console.log('WebSocket disconnected');
  connected.value = false;
};

socket.onerror = (error) => {
  console.error('WebSocket error:', error);
};

socket.onmessage = (event) => {
  const ticker = JSON.parse(event.data);

  // Ensure that previous price is a number
  const previousTicker = tickers.value.find((t) => t.Symbol === ticker.Symbol);
  const previousPrice = previousTicker ? previousTicker.Price : 0;

  const currentPrice = ticker.Price;

  // Check if currentPrice is a valid number
  if (isNaN(currentPrice)) {
    console.error('Received invalid current price:', ticker.Price);
    return;
  }

  const priceChange = currentPrice - previousPrice;

  // Add a property to indicate the price change
  ticker.PriceChange = priceChange;

  // Update the tickers array but keep the order by following the type rules
  const existingTickerIndex = tickers.value.findIndex(
    (t) => t.Symbol === ticker.Symbol
  );
  if (existingTickerIndex !== -1) {
    tickers.value[existingTickerIndex] = ticker;
  } else {
    tickers.value.push(ticker);
  }
};

// Computed property to sort tickers by price
const sortedTickers = computed(() => {
  return [...tickers.value].sort((a, b) => b.Price - a.Price);
});
</script>
