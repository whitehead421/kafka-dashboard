interface ProcessedTicker {
  Symbol: string;
  Name: string;
  Price: number;
  Change: number;
  ChangePct: number;
  Volume: number;
  PriceChange: number;
}

interface Column {
  index: string;
  title: string;
}
