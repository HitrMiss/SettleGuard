import { writable } from "svelte/store";

export const wallet = writable({
  connected: false,
  address: null,
  chainId: null
});
