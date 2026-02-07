<script>
  import { cart, cartSubtotal, cartCount } from '$lib/stores/cart.js';
  import menu from '$lib/data/menu.json';

  import WalletButton from '$lib/components/WalletButton.svelte';
  import { wallet } from '$lib/stores/wallet.js';
  import { sendTxFromApi } from '$lib/wallet/metamask.js';

  const { restaurant } = menu;
  const FREE_DELIVERY_THRESHOLD = 35;

  // Go API
  const API_BASE = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';

  // Registry-configured values (keep in web/.env)
  const MERCHANT_IDENTITY =
    import.meta.env.VITE_MERCHANT_IDENTITY ??
    '0xE524B0B32282af2B6273A59159a101Fd7a79eDdb';

  const MERCHANT_PAYOUT =
    import.meta.env.VITE_MERCHANT_PAYOUT ??
    '0xE524B0B32282af2B6273A59159a101Fd7a79eDdb';

  const CATEGORY_ID =
    import.meta.env.VITE_CATEGORY_ID ??
    '0x464f4f445f424556455241474500000000000000000000000000000000000000';

  // Fees
  $: subtotal = round2(Number($cartSubtotal ?? 0));
  $: itemCount = $cartCount;

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

  // ---------------------------
  // Pay with MetaMask (Go API)
  // ---------------------------
  let paying = false;
  let payError = '';
  let txHash = '';
  let lastApiResponse = '';
  let lastTxRequest = '';
  let lastGasEstimate = '';

  function isHexString(v) {
    return typeof v === 'string' && v.startsWith('0x') && v.length >= 3;
  }

  function buildApiItems(cartArr) {
    // IMPORTANT: API expects quantity (not qty)
    return cartArr.map((line) => ({
      id: String(line.id),
      name: String(line.name ?? ''),
      price: Number(line.price ?? 0),
      quantity: Math.max(1, Number(line.qty ?? 1)),
      description: String(line.description ?? ''),
      image: String(line.image ?? '')
    }));
  }

  function parseApiJson(text) {
    try {
      return JSON.parse(text);
    } catch {
      return null;
    }
  }

  // Your API returns { targetContract, encodedAbi, amount, ... }
  // We must map that to an EIP-1193 tx request.
  function apiResponseToTxRequest(apiData, fromAddress) {
    const to = apiData?.targetContract;
    const data = apiData?.encodedAbi;

    if (!isHexString(to)) {
      throw new Error('API response missing targetContract (to).');
    }
    if (!isHexString(data)) {
      throw new Error('API response missing encodedAbi (data).');
    }

    // For contract calls that move ERC20 via contract logic, ETH value is 0.
    return {
      from: fromAddress,
      to,
      data,
      value: '0x0'
    };
  }

  // ✅ Key fix: strip gas + fee overrides so MetaMask estimates properly
  function sanitizeTxForMetamask(tx) {
    const clean = { ...tx };

    // Remove any explicit gas settings that can exceed caps
    delete clean.gas;
    delete clean.gasLimit;

    // Remove fee overrides unless you explicitly want them
    delete clean.maxFeePerGas;
    delete clean.maxPriorityFeePerGas;
    delete clean.gasPrice;

    // Ensure required fields are hex
    if (clean.value == null) clean.value = '0x0';
    if (typeof clean.value === 'number') clean.value = '0x' + clean.value.toString(16);

    return clean;
  }

  async function estimateGas(tx) {
    const eth = window?.ethereum;
    if (!eth?.request) return '';

    try {
      const est = await eth.request({
        method: 'eth_estimateGas',
        params: [tx]
      });
      return String(est);
    } catch (e) {
      // Keep the error text for debugging; don't block send attempt unless you want to
      return `estimate failed: ${e?.message ?? String(e)}`;
    }
  }

  async function payWithMetaMask() {
    payError = '';
    txHash = '';
    lastApiResponse = '';
    lastTxRequest = '';
    lastGasEstimate = '';
    paying = true;

    try {
      if ($cart.length === 0) throw new Error('Your cart is empty.');
      if (!$wallet.connected) throw new Error('Connect MetaMask first.');
      if (!$wallet.address) throw new Error('Wallet address missing.');

      // Body must match API expectations (your PS script shape + quantity)
      const body = {
        items: buildApiItems($cart),
        merchantIdentity: MERCHANT_IDENTITY,
        merchantPayout: MERCHANT_PAYOUT,
        categoryId: CATEGORY_ID
      };

      const res = await fetch(`${API_BASE}/api/checkout`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'X-User-Address': $wallet.address
        },
        body: JSON.stringify(body)
      });

      const text = await res.text();
      lastApiResponse = text;

      if (!res.ok) {
        throw new Error(`Checkout API failed: ${res.status} ${text}`);
      }

      const apiData = parseApiJson(text);
      if (!apiData) {
        throw new Error('Checkout API returned non-JSON.');
      }

      // Build tx from API response
      const rawTx = apiResponseToTxRequest(apiData, $wallet.address);

      // Sanitize (strip gas + fee overrides)
      const txReq = sanitizeTxForMetamask(rawTx);

      // Debug: show the actual tx we will send
      lastTxRequest = JSON.stringify(txReq, null, 2);

      // Optional: preflight estimate (helps diagnose reverts / chain issues)
      lastGasEstimate = await estimateGas(txReq);

      // Send via MetaMask
      const hash = await sendTxFromApi(txReq);
      txHash = hash;
    } catch (e) {
      payError = e?.message ?? String(e);
    } finally {
      paying = false;
    }
  }
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

          <div class="mt-5 grid gap-2">
            <div class="flex items-center justify-between">
              <div class="muted text-sm">Payment method</div>
              <WalletButton />
            </div>

            <button
              class="btn-primary w-full"
              type="button"
              disabled={$cart.length === 0 || !$wallet.connected || paying}
              on:click={payWithMetaMask}
            >
              {paying ? "Building tx..." : "Pay with MetaMask"}
            </button>

            {#if txHash}
              <div class="badge">
                <span class="muted">Tx</span>
                <span class="font-semibold">{txHash.slice(0, 10)}…{txHash.slice(-8)}</span>
              </div>
            {/if}

            {#if payError}
              <div class="muted text-xs" style="color: rgba(255, 138, 30, 0.95);">{payError}</div>
            {/if}

            {#if lastGasEstimate}
              <div class="muted text-xs">
                Gas estimate: <span class="font-mono break-all">{lastGasEstimate}</span>
              </div>
            {/if}

            {#if lastTxRequest}
              <details class="mt-2">
                <summary class="muted text-xs hover:text-white cursor-pointer">Tx request (debug)</summary>
                <pre class="mt-2 text-xs p-3 rounded-xl overflow-auto" style="background: rgba(0,0,0,0.35); border: 1px solid rgba(255,255,255,0.08);">{lastTxRequest}</pre>
              </details>
            {/if}

            {#if lastApiResponse}
              <details class="mt-2">
                <summary class="muted text-xs hover:text-white cursor-pointer">API response (debug)</summary>
                <pre class="mt-2 text-xs p-3 rounded-xl overflow-auto" style="background: rgba(0,0,0,0.35); border: 1px solid rgba(255,255,255,0.08);">{lastApiResponse}</pre>
              </details>
            {/if}
          </div>

          <div class="badge mt-3">
            <span class="font-semibold" style="color: rgb(var(--accent));">Demo</span>
            <span class="muted">Order completion triggers SettleGuard resolution flow.</span>
          </div>
        </div>
      </aside>
    </div>
  </div>
</section>
