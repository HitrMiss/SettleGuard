<script>
  import menu from '$lib/data/menu.json';
  import { cart, cartSubtotal } from '$lib/stores/cart.js';
  import { goto } from '$app/navigation';

  const { restaurant } = menu;

  const FREE_DELIVERY_THRESHOLD = 35;

  $: subtotal = Number($cartSubtotal ?? 0);

  // Base fees as numbers
  $: baseDeliveryFee = Number(restaurant.deliveryFee ?? 0);
  $: serviceFee = Number(restaurant.serviceFee ?? 0);

  // Waive delivery when subtotal meets threshold
  $: deliveryFee = subtotal >= FREE_DELIVERY_THRESHOLD ? 0 : baseDeliveryFee;

  $: estTotal = round2(subtotal + deliveryFee + serviceFee);

  function round2(n) {
    return Math.round((Number(n) + Number.EPSILON) * 100) / 100;
  }
</script>

<div class="grid gap-6">
  <div class="flex items-end justify-between">
    <div>
      <h1 class="h2">Your cart</h1>
      <div class="muted text-sm mt-1">Ghost Kitchen+ demo storefront</div>
    </div>
    <button class="btn-ghost" on:click={() => cart.clear()} disabled={$cart.length === 0}>Clear</button>
  </div>

  <div class="grid gap-6 md:grid-cols-[1.4fr_0.6fr]">
    <div class="card p-6">
      {#if $cart.length === 0}
        <div class="surface p-5">
          <div class="font-semibold">Cart is empty</div>
          <div class="muted mt-1">Return to the menu to add items.</div>
          <div class="mt-4">
            <a class="btn" href="/GhostKitchen">Browse menu</a>
          </div>
        </div>
      {:else}
        <div class="space-y-3">
          {#each $cart as item}
            <div class="surface p-4">
              <div class="flex items-start justify-between gap-4">
                <div class="min-w-0">
                  <div class="font-semibold truncate">{item.name}</div>
                  <div class="muted text-sm">${item.price.toFixed(2)} each</div>
                </div>
                <button class="btn-ghost" on:click={() => cart.remove(item.id)}>Remove</button>
              </div>

              <div class="mt-3 flex items-center justify-between">
                <div class="muted text-sm">Qty</div>
                <input
                  class="input w-28 text-right"
                  type="number"
                  min="1"
                  value={item.qty}
                  on:input={(e) => cart.setQty(item.id, e.target.value)}
                />
              </div>
            </div>
          {/each}
        </div>
      {/if}
    </div>

    <aside class="card p-6">
      <div class="font-semibold">Estimated total</div>
      <div class="muted text-sm mt-1">Fees shown are demo values.</div>

      <div class="mt-4 space-y-2 text-sm">
        <div class="flex justify-between"><span class="muted">Subtotal</span><span>${subtotal.toFixed(2)}</span></div>
        <div class="flex justify-between"><span class="muted">Delivery</span><span>${deliveryFee.toFixed(2)}</span></div>
        <div class="flex justify-between"><span class="muted">Service</span><span>${serviceFee.toFixed(2)}</span></div>
        <div class="hr pt-3 mt-3 flex justify-between font-semibold">
          <span>Total</span><span>${estTotal.toFixed(2)}</span>
        </div>
      </div>

      <button class="btn w-full mt-5" on:click={() => goto('/GhostKitchen/checkout')} disabled={$cart.length === 0}>
        Continue to checkout
      </button>

      <div class="muted text-xs mt-3">
        SettleGuard demo: final settlement may require approval depending on trust + risk.
      </div>
    </aside>
  </div>
</div>
