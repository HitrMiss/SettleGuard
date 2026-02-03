<script>
  import { cart, cartCount, cartSubtotal } from '$lib/stores/cart.js';
  import menu from '$lib/data/menu.json';
  import { goto } from '$app/navigation';

  const { restaurant } = menu;

  let open = false;
  $: count = $cartCount;
  $: subtotal = $cartSubtotal;

  const delivery = restaurant.deliveryFee;
  const service = restaurant.serviceFee;
  $: estTotal = subtotal + delivery + service;

  function toggle() { open = !open; }
  function close() { open = false; }
  function to(path) { open = false; goto(path); }
</script>

<button class="btn" on:click={toggle} aria-label="Open cart">
  Cart
  <span class="ml-2 badge"><span class="font-semibold">{count}</span></span>
</button>

{#if open}
  <div class="drawer-overlay" on:click={close}></div>

  <aside class="drawer" role="dialog" aria-modal="true" aria-label="Cart drawer">
    <div class="drawer-header">
      <div class="flex items-center gap-3">
        <div class="logo"><span class="font-black tracking-tight text-sm">GK+</span></div>
        <div>
          <div class="font-semibold">Cart</div>
          <div class="muted text-xs">{count} item(s)</div>
        </div>
      </div>
      <button class="btn-quiet" on:click={close}>Close</button>
    </div>

    <div class="drawer-body">
      {#if $cart.length === 0}
        <div class="shell p-4">
          <div class="font-semibold">Cart is empty</div>
          <div class="muted text-sm mt-1">Add items from the menu to start.</div>
          <button class="btn w-full mt-4" on:click={() => to('/')}>Browse menu</button>
        </div>
      {:else}
        <div class="space-y-3">
          {#each $cart as item}
            <div class="card-soft p-4">
              <div class="flex items-start justify-between gap-3">
                <div class="min-w-0">
                  <div class="font-semibold truncate">{item.qty}× {item.name}</div>
                  <div class="muted text-xs">${item.price.toFixed(2)} each</div>
                </div>
                <button class="btn-danger" on:click={() => cart.remove(item.id)}>Remove</button>
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
          <div class="flex justify-between"><span class="muted">Subtotal</span><span>${subtotal.toFixed(2)}</span></div>
          <div class="flex justify-between"><span class="muted">Delivery</span><span>${delivery.toFixed(2)}</span></div>
          <div class="flex justify-between"><span class="muted">Service</span><span>${service.toFixed(2)}</span></div>
          <div class="hr pt-3 mt-3 flex justify-between font-semibold text-base">
            <span>Estimated total</span><span>${estTotal.toFixed(2)}</span>
          </div>
        </div>
      {/if}
    </div>

    <div class="drawer-footer">
      <div class="grid grid-cols-2 gap-2">
        <button class="btn" on:click={() => to('/cart')} disabled={$cart.length === 0}>View</button>
        <button class="btn-primary" on:click={() => to('/checkout')} disabled={$cart.length === 0}>Checkout</button>
      </div>
    </div>
  </aside>
{/if}
