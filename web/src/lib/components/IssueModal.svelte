<script>
  import { disputeOrder } from '$lib/api/client.js';

  export let open = false;
  export let orderId = '';
  export let onClose = () => {};
  export let onUpdated = (order) => {};

  let issueType = 'missing_item';
  let details = '';
  let busy = false;
  let error = '';

  async function submit() {
    busy = true;
    error = '';
    try {
      const res = await disputeOrder(orderId, { issueType, details });
      onUpdated(res.order);
      details = '';
      onClose();
    } catch (e) {
      error = e.message || 'Failed';
    } finally {
      busy = false;
    }
  }
</script>

{#if open}
  <div class="fixed inset-0 z-50 flex items-end sm:items-center justify-center">
    <div class="drawer-overlay" on:click={onClose}></div>

    <div class="relative w-full sm:max-w-xl bg-slate-950 border border-slate-800 rounded-t-3xl sm:rounded-3xl p-5 glow-border">
      <div class="flex items-start justify-between gap-4">
        <div>
          <div class="font-semibold text-lg">Report an issue</div>
          <div class="muted text-sm mt-1">
            Simulates SettleGuard resolution: review, approval hold, or safe unwind.
          </div>
        </div>
        <button class="btn-ghost" on:click={onClose}>Close</button>
      </div>

      <div class="mt-5 grid gap-3">
        <label class="block">
          <div class="muted text-sm mb-1">Issue type</div>
          <select class="input" bind:value={issueType}>
            <option value="missing_item">Missing item</option>
            <option value="wrong_order">Wrong order</option>
            <option value="not_received">Didn’t receive order</option>
            <option value="suspected_fraud">Suspected fraud</option>
          </select>
        </label>

        <label class="block">
          <div class="muted text-sm mb-1">Details (optional)</div>
          <textarea class="input" rows="3" bind:value={details} placeholder="Add a note for support…"></textarea>
        </label>

        {#if error}
          <div class="rounded-2xl border border-red-500/30 bg-red-500/10 p-4">
            <div class="text-red-200 font-semibold">Couldn’t submit</div>
            <div class="text-red-200/90 text-sm mt-1">{error}</div>
          </div>
        {/if}

        <button class="btn w-full" disabled={busy} on:click={submit}>
          {busy ? 'Submitting…' : 'Submit issue'}
        </button>

        <div class="muted text-xs">
          Tip: choose “Suspected fraud” to trigger an automatic unwind in the demo.
        </div>
      </div>
    </div>
  </div>
{/if}
