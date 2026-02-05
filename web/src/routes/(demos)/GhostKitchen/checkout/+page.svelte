<script>
  import { cart, cartSubtotal, cartCount } from '$lib/stores/cart.js';
  import menu from '$lib/data/menu.json';

  const { restaurant } = menu;
  const FREE_DELIVERY_THRESHOLD = 35;
  
  // Fees
  $: baseDeliveryFee = Number(restaurant.deliveryFee ?? 0);
  $: deliveryFee = subtotal >= FREE_DELIVERY_THRESHOLD ? 0 : baseDeliveryFee;
  const serviceFee = Number(restaurant.serviceFee ?? 0);

  // Tip selection
  const presetPercents = [10, 15, 20];
  /** @type {number|'custom'} */
  let tipMode = 10;
  let customPercent = 18;

  // Delivery address (demo editable fields)
  let deliveryName = 'Jordan K.';
  let deliveryPhone = '(403) 555-0182';
  let deliveryAddress1 = '1208 5th Ave SW';
  let deliveryAddress2 = 'Unit 1406';
  let deliveryCity = 'Calgary, AB';
  let deliveryPostal = 'T2P 3N1';
  let deliveryNotes = 'Leave at door. Do not ring.';

  // Derived cart values
  $: subtotal = round2($cartSubtotal);
  $: itemCount = $cartCount;

  // "Final cart amount" BEFORE tip
  $: preTipTotal = round2(subtotal + deliveryFee + serviceFee);

  // Selected tip percentage
  $: selectedPercent =
    tipMode === 'custom' ? clampPercent(customPercent) : Number(tipMode);

  // Tip amount + final total
  $: tipAmount = round2(preTipTotal * (selectedPercent / 100));
  $: finalTotal = round2(preTipTotal + tipAmount);

  function selectPreset(p) {
    tipMode = p;
  }
  function selectCustom() {
    tipMode = 'custom';
  }
  function onCustomInput(e) {
    tipMode = 'custom';
    customPercent = e.currentTarget.value;
  }

  function clampPercent(v) {
    const n = Number(v);
    if (!Number.isFinite(n)) return 0;
    return Math.max(0, Math.min(100, n));
  }
  function round2(n) {
    return Math.round((Number(n) + Number.EPSILON) * 100) / 100;
  }

  function removeItem(id) {
    cart.remove(id);
  }
  function setQty(id, value) {
    cart.setQty(id, value);
  }

  $: deliverToLine =
    `${deliveryAddress1}${deliveryAddress2 ? `, ${deliveryAddress2}` : ''}, ${deliveryCity}`;
</script>

<section class="section">
  <div class="shell p-6 sm:p-8">
    <div class="kicker">Checkout</div>
    <h1 class="h2 mt-2">Review order</h1>
    <p class="muted mt-2">
      Tip is calculated as a percentage of your final cart amount (pre-tip).
    </p>

    <div class="rail mt-6"></div>

    <div class="mt-6 grid gap-6 lg:grid-cols-[1fr_22rem]">
      <!-- LEFT -->
      <div class="grid gap-4">
        <!-- Delivery address (REQUIRED) -->
        <div class="card p-5">
          <div class="flex items-start justify-between gap-3">
            <div>
              <div class="font-semibold">Delivery address</div>
              <div class="muted text-sm mt-1">
                Confirm where your order will be delivered.
              </div>
            </div>

            <span class="badge">
              <span class="muted">ETA</span>
              <span class="font-semibold">{restaurant.etaMin}-{restaurant.etaMax} min</span>
            </span>
          </div>

          <div class="hr mt-4"></div>

          <div class="mt-4 grid gap-3 sm:grid-cols-2">
            <div>
              <label class="muted text-xs">Name</label>
              <input class="input mt-1" bind:value={deliveryName} placeholder="Full name" />
            </div>

            <div>
              <label class="muted text-xs">Phone</label>
              <input class="input mt-1" bind:value={deliveryPhone} placeholder="(###) ###-####" />
            </div>

            <div class="sm:col-span-2">
              <label class="muted text-xs">Address line 1</label>
              <input class="input mt-1" bind:value={deliveryAddress1} placeholder="Street address" />
            </div>

            <div>
              <label class="muted text-xs">Address line 2</label>
              <input class="input mt-1" bind:value={deliveryAddress2} placeholder="Apt / Unit / Buzz code" />
            </div>

            <div>
              <label class="muted text-xs">City / Province</label>
              <input class="input mt-1" bind:value={deliveryCity} placeholder="City, Province" />
            </div>

            <div>
              <label class="muted text-xs">Postal code</label>
              <input class="input mt-1" bind:value={deliveryPostal} placeholder="A1A 1A1" />
            </div>

            <div class="sm:col-span-2">
			  <label class="muted text-xs">Delivery instructions</label>
			  <textarea
				class="input mt-1 resize-none min-h-[96px]"
				bind:value={deliveryNotes}
				placeholder="e.g., Leave at door, do not ring, call on arrival"
				rows="3"
			  ></textarea>
			</div>
          </div>
        </div>

        <!-- Items -->
        <div class="card p-5">
          <div class="flex items-center justify-between">
            <div class="font-semibold">Items</div>
            <div class="muted text-sm">{itemCount} item(s)</div>
          </div>

          <div class="hr mt-4"></div>

          {#if $cart.length === 0}
            <div class="muted mt-4">Your cart is empty.</div>
          {:else}
            <div class="mt-4 space-y-3">
              {#each $cart as item}
                <div class="card-soft p-4">
                  <div class="flex items-start justify-between gap-3">
                    <div class="min-w-0">
                      <div class="font-semibold truncate">{item.name}</div>
                      <div class="muted text-xs">${item.price.toFixed(2)} each</div>
                    </div>
                    <button class="btn-danger" type="button" on:click={() => removeItem(item.id)}>
                      Remove
                    </button>
                  </div>

                  <div class="mt-3 flex items-center justify-between">
                    <div class="muted text-xs">Quantity</div>
                    <input
                      class="input w-24 text-right"
                      type="number"
                      min="1"
                      value={item.qty}
                      on:input={(e) => setQty(item.id, e.currentTarget.value)}
                    />
                  </div>
                </div>
              {/each}
            </div>
          {/if}
        </div>

        <!-- Tip -->
        <div class="card p-5">
          <div class="flex items-start justify-between gap-3">
            <div>
              <div class="font-semibold">Tip</div>
              <div class="muted text-sm mt-1">
                Based on final cart amount (pre-tip): ${preTipTotal.toFixed(2)}
              </div>
            </div>

            <div class="badge">
              <span class="muted">Selected</span>
              <span class="font-semibold" style="color: rgb(var(--accent));">{selectedPercent}%</span>
            </div>
          </div>

          <div class="mt-4 flex flex-wrap gap-2">
            {#each presetPercents as p}
              <button
                type="button"
                class={tipMode === p ? "pill pill-active" : "pill"}
                on:click={() => selectPreset(p)}
              >
                {p}%
              </button>
            {/each}

            <button
              type="button"
              class={tipMode === 'custom' ? "pill pill-active" : "pill"}
              on:click={selectCustom}
            >
              Custom
            </button>
          </div>

          {#if tipMode === 'custom'}
            <div class="mt-4 grid gap-2 sm:grid-cols-[1fr_180px] sm:items-center">
              <div class="muted text-sm">
                Custom tip percentage (0–100). Tip updates automatically.
              </div>

              <div class="flex items-center gap-2">
                <input
                  class="input text-right"
                  type="number"
                  min="0"
                  max="100"
                  step="1"
                  value={customPercent}
                  on:input={onCustomInput}
                  aria-label="Custom tip percent"
                />
                <div class="muted font-semibold">%</div>
              </div>
            </div>
          {/if}

          <div class="hr mt-4"></div>

          <div class="mt-4 flex items-center justify-between">
            <div class="muted">Tip amount</div>
            <div class="font-semibold">${tipAmount.toFixed(2)}</div>
          </div>
        </div>
      </div>

      <!-- RIGHT: Summary -->
      <aside class="grid gap-4">
        <div class="shell p-5">
          <div class="flex items-start justify-between gap-3">
            <div>
              <div class="font-semibold text-lg">Summary</div>
              <div class="muted text-xs mt-1">Deliver to: {deliverToLine}</div>
            </div>
            <span class="badge">
              <span class="muted">Order</span>
              <span class="font-semibold">{itemCount} item(s)</span>
            </span>
          </div>

          <div class="hr mt-4"></div>

          <div class="mt-4 text-sm space-y-2">
            <div class="flex justify-between">
              <span class="muted">Subtotal</span>
              <span>${subtotal.toFixed(2)}</span>
            </div>

            <div class="flex justify-between">
              <span class="muted">Delivery</span>
              <span>${deliveryFee.toFixed(2)}</span>
            </div>

            <div class="flex justify-between">
              <span class="muted">Service</span>
              <span>${serviceFee.toFixed(2)}</span>
            </div>

            <div class="hr mt-3"></div>

            <div class="flex justify-between font-semibold">
              <span class="muted">Pre-tip total</span>
              <span>${preTipTotal.toFixed(2)}</span>
            </div>

            <div class="flex justify-between">
              <span class="muted">Tip ({selectedPercent}%)</span>
              <span>${tipAmount.toFixed(2)}</span>
            </div>

            <div class="hr mt-3"></div>

            <div class="flex justify-between font-semibold text-base">
              <span>Total</span>
              <span>${finalTotal.toFixed(2)}</span>
            </div>
          </div>

          <button class="btn-primary w-full mt-4" type="button" disabled={$cart.length === 0}>
            Place order (demo)
          </button>

          <div class="badge mt-3">
            <span class="font-semibold" style="color: rgb(var(--accent));">Demo</span>
            <span class="muted">Order completion triggers SettleGuard resolution flow.</span>
          </div>
        </div>
      </aside>
    </div>
  </div>
</section>
