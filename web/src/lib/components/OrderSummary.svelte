<script>
  import { cart, cartSubtotal, cartCount } from '$lib/stores/cart.js';
  import menu from '$lib/data/menu.json';
  import { goto } from '$app/navigation';

  const { restaurant } = menu;

  $: subtotal = $cartSubtotal;
  $: count = $cartCount;

  const FREE_DELIVERY_THRESHOLD = 35;
  const service = restaurant.serviceFee;
  $: baseDelivery = Number(restaurant.deliveryFee ?? 0);
  $: delivery = subtotal >= FREE_DELIVERY_THRESHOLD ? 0 : baseDelivery;
  $: estTotal = subtotal + delivery + service;

  function to(path) {
    goto(path);
  }
</script>

<div class="card-soft p-5">
  <div class="flex items-center justify-between gap-3">
    <div>
      <div class="font-extrabold text-lg">Your order</div>
      <div class="muted text-sm">{count} item(s)</div>
    </div>
    <a class="btn-quiet" href="/cart">Edit</a>
  </div>

  <div class="hr mt-4"></div>

  {#if $cart.length === 0}
    <div class="muted mt-4">Add items to start an order.</div>
  {:else}
    <div class="mt-4 space-y-3 max-h-72 overflow-auto pr-1">
      {#each $cart as item}
        <div class="card-soft p-3">
          <div class="flex items-start justify-between gap-3">
            <div class="min-w-0">
              <div class="font-extrabold truncate">{item.qty}x {item.name}</div>
              <div class="muted text-xs">${item.price.toFixed(2)} each</div>
            </div>
            <button class="btn btn-danger" on:click={() => cart.remove(item.id)} type="button">Remove</button>
          </div>

          <div class="mt-2 flex items-center justify-between">
            <div class="muted text-xs">Qty</div>
            <input
              class="input w-24 text-right"
              type="number"
              min="1"
              value={item.qty}
              on:input={(e) => cart.setQty(item.id, e.target.value)}
            />
          </div>
        </div>
      {/each}
    </div>

    <div class="hr mt-4"></div>

    <div class="mt-4 text-sm space-y-2">
      <div class="flex justify-between"><span class="muted">Subtotal</span><span class="font-extrabold">${subtotal.toFixed(2)}</span></div>
      <div class="flex justify-between"><span class="muted">Delivery</span><span class="font-extrabold">${delivery.toFixed(2)}</span></div>
      <div class="flex justify-between"><span class="muted">Service</span><span class="font-extrabold">${service.toFixed(2)}</span></div>

      <div class="hr mt-3"></div>

      <div class="pt-2 flex justify-between font-extrabold text-base">
        <span>Estimated total</span><span>${estTotal.toFixed(2)}</span>
      </div>
    </div>

    <button class="btn btn-primary w-full mt-4" on:click={() => to('/GhostKitchen/checkout')} type="button">
      Checkout
    </button>
  {/if}
</div>
