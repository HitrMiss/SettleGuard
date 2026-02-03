<script>
  import menu from '$lib/data/menu.json';
  import { cart, cartSubtotal } from '$lib/stores/cart.js';
  import { createOrder } from '$lib/api/client.js';
  import { goto } from '$app/navigation';

  const { restaurant, merchant } = menu;

  let address = '123 Demo Street';
  let instructions = 'Leave at door';
  let tip = 2.0;
  let busy = false;
  let error = '';

  const tipPresets = [0, 2, 4, 6];

  $: subtotal = $cartSubtotal;
  $: deliveryFee = restaurant.deliveryFee;
  $: serviceFee = restaurant.serviceFee;
  $: total = round2(subtotal + deliveryFee + serviceFee + Number(tip || 0));

  function round2(n) {
    return Math.round(n * 100) / 100;
  }

  function setTip(v) {
    tip = v;
  }

  async function placeOrder() {
    error = '';
    if ($cart.length === 0) {
      error = 'Cart is empty.';
      return;
    }
    if (!address.trim()) {
      error = 'Address is required.';
      return;
    }

    busy = true;
    try {
      const payload = {
        restaurantName: restaurant.name,
        merchant,
        items: $cart.map((x) => ({ id: x.id, name: x.name, price: x.price, qty: x.qty, category: x.category })),
        subtotal,
        deliveryFee,
        serviceFee,
        tip: Number(tip || 0),
        total,
        address,
        instructions
      };

      const res = await createOrder(payload);
      cart.clear();
      await goto(`/order/${res.order.id}`);
    } catch (e) {
      error = e.message || 'Failed';
    } finally {
      busy = false;
    }
  }
</script>

<div class="grid gap-6">
  <div>
    <h1 class="h2">Checkout</h1>
    <div class="muted text-sm mt-1">Ghost Kitchen+ • place an order, then report an issue to demo SettleGuard</div>
  </div>

  <div class="grid gap-6 md:grid-cols-[1fr_0.8fr]">
    <div class="card p-6 space-y-5">
      <div class="surface p-4">
        <div class="font-semibold">Delivery details</div>
        <div class="muted text-sm mt-1">Keep it simple — this is a demo flow.</div>

        <div class="mt-4 grid gap-3">
          <label class="block">
            <div class="muted text-sm mb-1">Address</div>
            <input class="input" bind:value={address} />
          </label>

          <label class="block">
            <div class="muted text-sm mb-1">Instructions</div>
            <input class="input" bind:value={instructions} />
          </label>
        </div>
      </div>

      <div class="surface p-4">
        <div class="font-semibold">Tip</div>
        <div class="muted text-sm mt-1">Use presets to quickly change the total (affects risk scoring).</div>

        <div class="mt-3 flex flex-wrap gap-2">
          {#each tipPresets as t}
            <button class={"chip " + (Number(tip || 0) === t ? "chip-active" : "")} on:click={() => setTip(t)}>
              ${t.toFixed(0)}
            </button>
          {/each}
          <div class="flex-1 min-w-[140px]">
            <input class="input" type="number" step="0.5" min="0" bind:value={tip} />
          </div>
        </div>
      </div>

      {#if error}
        <div class="surface p-4 border border-red-500/30">
          <div class="text-red-300 font-semibold">Couldn’t place order</div>
          <div class="text-red-200/90 text-sm mt-1">{error}</div>
        </div>
      {/if}

      <button class="btn w-full" disabled={busy || $cart.length === 0} on:click={placeOrder}>
        {busy ? 'Placing order…' : 'Place order'}
      </button>

      <div class="muted text-xs">
        Demo behavior: payment “executes instantly,” but final settlement may require approval depending on risk.
      </div>
    </div>

    <aside class="card p-6">
      <div class="font-semibold">Receipt</div>
      <div class="muted text-sm mt-1">
        ETA {restaurant.etaMin}-{restaurant.etaMax} min • Fees are fixed demo values
      </div>

      <div class="mt-4 space-y-3">
        {#if $cart.length === 0}
          <div class="surface p-4">
            <div class="font-semibold">No items</div>
            <div class="muted text-sm mt-1">Add items from the menu first.</div>
            <div class="mt-4">
              <a class="btn-secondary" href="/">Browse menu</a>
            </div>
          </div>
        {:else}
          <div class="surface p-4">
            <div class="font-semibold">Items</div>
            <div class="mt-3 space-y-2 text-sm">
              {#each $cart as it}
                <div class="flex justify-between gap-3">
                  <span class="muted">{it.qty}× {it.name}</span>
                  <span>${(it.price * it.qty).toFixed(2)}</span>
                </div>
              {/each}
            </div>
          </div>

          <div class="surface p-4 text-sm space-y-2">
            <div class="flex justify-between"><span class="muted">Subtotal</span><span>${subtotal.toFixed(2)}</span></div>
            <div class="flex justify-between"><span class="muted">Delivery</span><span>${deliveryFee.toFixed(2)}</span></div>
            <div class="flex justify-between"><span class="muted">Service</span><span>${serviceFee.toFixed(2)}</span></div>
            <div class="flex justify-between"><span class="muted">Tip</span><span>${Number(tip || 0).toFixed(2)}</span></div>
            <div class="hr pt-3 mt-3 flex justify-between font-semibold text-base">
              <span>Total</span><span>${total.toFixed(2)}</span>
            </div>
          </div>

          <div class="badge">
            <span class="font-semibold">Note</span>
            <span class="muted">Higher totals may trigger approval required.</span>
          </div>
        {/if}
      </div>
    </aside>
  </div>
</div>
