import { writable } from "svelte/store";

export const wallet = writable({
  connected: false,
  address: null,
  chainId: null,
  approval: false // are we allowed to access the users USDC?
});
