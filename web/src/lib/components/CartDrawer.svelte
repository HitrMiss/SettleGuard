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

  function toggle() {
    open = !open;
  }
  function close() {
    open = false;
  }
  function to(path) {
    open = false;
    goto(path);
  }
</script>

<button class="btn" on:click={toggle} type="button" aria-label="Open cart">
  Cart
  <span
    class="ml-2 inline-flex items-center justify-center text-xs font-extrabold px-2 py-1 rounded-full"
    style="background: rgba(255,138,30,0.14); border: 1px solid rgba(255,138,30,0.28); color: rgba(230,235,242,0.95); min-width: 2.25rem;"
  >
    {count}
  </span>
</button>

{#if open}
  <!-- overlay -->
  <div
    class="fixed inset-0 z-40"
    style="background: rgba(0,0,0,0.65); backdrop-filter: blur(8px);"
    on:click={close}
  ></div>

  <!-- drawer -->
  <aside
    class="fixed top-0 right-0 h-full z-50 w-[26rem] max-w-[92vw]"
    style="background: rgba(11,15,20,0.96); border-left: 1px solid rgba(255,255,255,0.10);"
    role="dialog"
    aria-modal="true"
    aria-label="Cart drawer"
  >
    <!-- header -->
    <div class="p-4 border-b flex items-center justify-between gap-3" style="border-color: rgba(255,255,255,0.10);">
      <div class="flex items-center gap-3 min-w-0">
        <img
          src="/brand/ghostkitchen-logo-transparent.png"
          alt="Ghost Kitchen+"
          class="h-8 w-auto shrink-0"
        />
        <div class="min-w-0">
          <div class="font-extrabold truncate">Cart</div>
          <div class="muted text-xs">{count} item(s)</div>
        </div>
      </div>

      <button class="btn-quiet shrink-0" on:click={close} type="button">Close</button>
    </div>

    <!-- body -->
    <div class="p-4 overflow-auto" style="max-height: calc(100vh - 170px);">
      {#if $cart.length === 0}
        <div class="card-soft p-4">
          <div class="font-extrabold">Cart is empty</div>
          <div class="muted text-sm mt-1">Add items from the menu to start.</div>
          <button class="btn w-full mt-4" on:click={() => to('/GhostKitchen')} type="button">Browse menu</button>
        </div>
      {:else}
        <div class="space-y-3">
          {#each $cart as item}
            <div class="card-soft p-4">
              <div class="flex items-start justify-between gap-3">
                <div class="min-w-0">
                  <div class="font-extrabold truncate">{item.qty}x {item.name}</div>
                  <div class="muted text-xs">${item.price.toFixed(2)} each</div>
                </div>
                <button class="btn btn-danger shrink-0" on:click={() => cart.remove(item.id)} type="button">
                  Remove
                </button>
              </div>

              <div class="mt-2 flex items-center justify-between gap-3">
                <div class="muted text-xs">Qty</div>
                <input
                  class="input w-24 text-right shrink-0"
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
          <div class="flex justify-between">
            <span class="muted">Subtotal</span><span class="font-extrabold">${subtotal.toFixed(2)}</span>
          </div>
          <div class="flex justify-between">
            <span class="muted">Delivery</span><span class="font-extrabold">${delivery.toFixed(2)}</span>
          </div>
          <div class="flex justify-between">
            <span class="muted">Service</span><span class="font-extrabold">${service.toFixed(2)}</span>
          </div>

          <div class="hr mt-3"></div>

          <!-- ? fixed: not using .hr as the container -->
          <div class="pt-2 flex justify-between font-extrabold text-base">
            <span>Estimated total</span><span>${estTotal.toFixed(2)}</span>
          </div>
        </div>
      {/if}
    </div>

    <!-- footer -->
    <div class="p-4 border-t" style="border-color: rgba(255,255,255,0.10);">
      <div class="grid grid-cols-2 gap-2">
        <button
          class="btn w-full"
          on:click={() => to('/GhostKitchen/cart')}
          disabled={$cart.length === 0}
          type="button"
        >
          View
        </button>

        <button
          class="btn btn-primary w-full"
          on:click={() => to('/GhostKitchen/checkout')}
          disabled={$cart.length === 0}
          type="button"
        >
          Checkout
        </button>
      </div>
    </div>
  </aside>
{/if}
