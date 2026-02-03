import { writable, derived } from 'svelte/store';

function createCart() {
  const { subscribe, set, update } = writable([]);

  return {
    subscribe,
    add(item) {
      update((items) => {
        const found = items.find((x) => x.id === item.id);
        if (found) return items.map((x) => (x.id === item.id ? { ...x, qty: x.qty + 1 } : x));
        return [...items, { ...item, qty: 1 }];
      });
    },
    remove(id) {
      update((items) => items.filter((x) => x.id !== id));
    },
    setQty(id, qty) {
      const q = Math.max(1, Number(qty || 1));
      update((items) => items.map((x) => (x.id === id ? { ...x, qty: q } : x)));
    },
    clear() {
      set([]);
    }
  };
}

export const cart = createCart();

export const cartCount = derived(cart, ($cart) => $cart.reduce((acc, x) => acc + x.qty, 0));
export const cartSubtotal = derived(cart, ($cart) =>
  round2($cart.reduce((acc, x) => acc + x.price * x.qty, 0))
);

function round2(n) {
  return Math.round(n * 100) / 100;
}
