<script>
  import { page } from '$app/stores';
  import { onDestroy } from 'svelte';
  import { getOrder, approveOrder, reverseOrder } from '$lib/api/client.js';
  import OrderTimeline from '$lib/components/OrderTimeline.svelte';
  import IssueModal from '$lib/components/IssueModal.svelte';
  import TrustBadge from '$lib/components/TrustBadge.svelte';

  let order = null;
  let error = '';
  let busy = false;
  let issueOpen = false;

  let timer;

  async function refresh() {
    error = '';
    try {
      const id = $page.params.id;
      const res = await getOrder(id);
      order = res.order;
    } catch (e) {
      error = e.message || 'Failed to load order';
    }
  }

  async function approve() {
    busy = true;
    try {
      const id = $page.params.id;
      const res = await approveOrder(id);
      order = res.order;
    } catch (e) {
      error = e.message || 'Approve failed';
    } finally {
      busy = false;
    }
  }

  async function reverse() {
    busy = true;
    try {
      const id = $page.params.id;
      const res = await reverseOrder(id);
      order = res.order;
    } catch (e) {
      error = e.message || 'Reverse failed';
    } finally {
      busy = false;
    }
  }

  refresh();
  timer = setInterval(refresh, 2500);

  onDestroy(() => {
    if (timer) clearInterval(timer);
  });

  const settlementLabel = {
    pending_authorization: 'Pending Authorization',
    approved: 'Approved',
    finalized: 'Finalized',
    reversed: 'Reversed',
    under_review: 'Under Review'
  };

  const statusLabel = {
    placed: 'Placed',
    preparing: 'Preparing',
    picked_up: 'Picked up',
    delivered: 'Delivered',
    cancelled: 'Cancelled'
  };

  function progressFor(status) {
    if (status === 'placed') return 25;
    if (status === 'preparing') return 45;
    if (status === 'picked_up') return 70;
    if (status === 'delivered') return 100;
    if (status === 'cancelled') return 100;
    return 35;
  }

  function etaText(status) {
    if (status === 'placed') return 'ETA ~30-40 min';
    if (status === 'preparing') return 'ETA ~20-30 min';
    if (status === 'picked_up') return 'ETA ~10-20 min';
    if (status === 'delivered') return 'Delivered';
    if (status === 'cancelled') return 'Cancelled';
    return 'ETA updating…';
  }

  function toneForSettlement(s) {
    if (s === 'finalized') return 'bg-emerald-500/10 border-emerald-400/20';
    if (s === 'pending_authorization') return 'bg-amber-500/10 border-amber-400/20';
    if (s === 'under_review') return 'bg-purple-500/10 border-purple-400/20';
    if (s === 'reversed') return 'bg-red-500/10 border-red-400/20';
    return 'bg-slate-950/30 border-slate-800';
  }
</script>

<div class="grid gap-6">
  <div class="flex flex-col sm:flex-row sm:items-end sm:justify-between gap-4">
    <div>
      <div class="kicker">Ghost Kitchen+ • Order</div>
      <h1 class="h2 mt-1">Tracking & Resolution</h1>
      <div class="muted text-sm mt-1">
        This page demonstrates SettleGuard’s controls: approval holds, review, and safe unwind.
      </div>
    </div>

    <div class="flex items-center gap-2">
      <button class="btn-secondary" on:click={() => (issueOpen = true)} disabled={!order || busy}>
        Report an issue
      </button>
      <button class="btn-ghost" on:click={refresh} disabled={busy}>Refresh</button>
    </div>
  </div>

  {#if error}
    <div class="card p-4 border border-red-500/30 bg-red-500/10">
      <div class="text-red-200 font-semibold">Error</div>
      <div class="text-red-200/90 text-sm mt-1">{error}</div>
    </div>
  {/if}

  {#if order}
    <section class="card p-6">
      <div class="grid gap-6 lg:grid-cols-[1fr_0.9fr]">
        <!-- Left: tracking -->
        <div class="space-y-4">
          <div class="flex items-start justify-between gap-4">
            <div class="min-w-0">
              <div class="muted text-xs">Order ID</div>
              <div class="font-mono text-xs sm:text-sm break-all">{order.id}</div>
              <div class="mt-3 flex flex-wrap items-center gap-2">
                <span class="badge">
                  <span class="font-semibold">{statusLabel[order.status] || order.status}</span>
                  <span class="muted">{etaText(order.status)}</span>
                </span>

                <span class={"badge border " + toneForSettlement(order.settlement)}>
                  <span class="font-semibold">{settlementLabel[order.settlement] || order.settlement}</span>
                  {#if order.requiresApproval}
                    <span class="muted">approval required</span>
                  {:else}
                    <span class="muted">controlled</span>
                  {/if}
                </span>

                <TrustBadge trustTier={order.merchant.trustTier} bondLevel={order.merchant.bondLevel} />
              </div>
            </div>
          </div>

          <div class="surface p-4">
            <div class="flex items-center justify-between">
              <div class="font-semibold">Delivery progress</div>
              <div class="muted text-sm">{progressFor(order.status)}%</div>
            </div>

            <div class="mt-3 h-3 rounded-full bg-slate-900 border border-slate-800 overflow-hidden">
              <div
                class="h-full bg-slate-100/90"
                style={"width:" + progressFor(order.status) + "%"}
              ></div>
            </div>

            <div class="muted text-sm mt-3">
              <span class="text-slate-200 font-semibold">Risk:</span>
              {order.risk.level} ({order.risk.score}) •
              finality is delayed if risk is elevated.
            </div>
          </div>

          <div class="surface p-4">
            <div class="font-semibold">Items</div>
            <div class="mt-3 grid gap-2">
              {#each order.items as it}
                <div class="flex items-center justify-between rounded-xl border border-slate-800 bg-slate-950/30 p-3">
                  <div class="min-w-0">
                    <div class="font-semibold truncate">{it.qty}× {it.name}</div>
                    <div class="muted text-xs">Unit ${it.price.toFixed(2)}</div>
                  </div>
                  <div class="font-semibold">${(it.price * it.qty).toFixed(2)}</div>
                </div>
              {/each}
            </div>
          </div>
        </div>

        <!-- Right: settlement panel -->
        <aside class="space-y-4">
          <div class="surface p-5">
            <div class="font-semibold">Settlement controls</div>
            <div class="muted text-sm mt-1">
              Instant execution ? instant finality. SettleGuard gates settlement based on risk and trust.
            </div>

            <div class="mt-4 space-y-2">
              {#if order.requiresApproval && order.settlement === 'pending_authorization'}
                <button class="btn w-full" disabled={busy} on:click={approve}>
                  {busy ? 'Approving…' : 'Approve settlement (secondary signing)'}
                </button>
              {/if}

              {#if order.settlement !== 'reversed'}
                <button class="btn-danger w-full" disabled={busy} on:click={reverse}>
                  Trigger unwind (reverse)
                </button>
              {/if}
            </div>

            <div class="mt-4">
              <div class="muted text-xs mb-2">Risk reasons</div>
              <ul class="space-y-1 text-xs text-slate-300">
                {#each order.risk.reasons as r}
                  <li class="rounded-xl border border-slate-800 bg-slate-950/30 p-2">{r}</li>
                {/each}
              </ul>
            </div>
          </div>

          <div class="surface p-5">
            <div class="font-semibold">Receipt</div>
            <div class="mt-3 text-sm space-y-2">
              <div class="flex justify-between"><span class="muted">Subtotal</span><span>${order.subtotal.toFixed(2)}</span></div>
              <div class="flex justify-between"><span class="muted">Delivery</span><span>${order.deliveryFee.toFixed(2)}</span></div>
              <div class="flex justify-between"><span class="muted">Service</span><span>${order.serviceFee.toFixed(2)}</span></div>
              <div class="flex justify-between"><span class="muted">Tip</span><span>${order.tip.toFixed(2)}</span></div>
              <div class="hr pt-3 mt-3 flex justify-between font-semibold text-base">
                <span>Total</span><span>${order.total.toFixed(2)}</span>
              </div>
            </div>

            <div class="badge mt-4">
              <span class="font-semibold">Demo</span>
              <span class="muted">Use “Report an issue” to show dispute ? review ? unwind.</span>
            </div>
          </div>
        </aside>
      </div>
    </section>

    <OrderTimeline {order} />

    <IssueModal
      bind:open={issueOpen}
      orderId={order.id}
      onClose={() => (issueOpen = false)}
      onUpdated={(o) => (order = o)}
    />
  {:else}
    <div class="card p-6 muted">Loading order…</div>
  {/if}
</div>
