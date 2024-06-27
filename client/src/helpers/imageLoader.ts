export const imageLoader = (symbol: string) => {
  let symbolWithoutUSDT = symbol.replace("USDT", "");
  return `/img/icons/${symbolWithoutUSDT.toLowerCase()}.svg`;
};
