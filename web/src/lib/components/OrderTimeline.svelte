<script>
  export let order;

  function fmt(ts) {
    const d = new Date(ts);
    return d.toLocaleString();
  }

  const settlementLabel = {
    pending_authorization: 'Pending Authorization',
    approved: 'Approved',
    finalized: 'Finalized',
    reversed: 'Reversed',
    under_review: 'Under Review'
  };

  const typeTone = {
    execution: 'border-slate-800 bg-slate-950/30',
    authorization: 'border-amber-500/20 bg-amber-500/10',
    settlement: 'border-emerald-500/20 bg-emerald-500/10',
    fulfillment: 'border-cyan-500/20 bg-cyan-500/10',
    support: 'border-purple-500/20 bg-purple-500/10',
    risk: 'border-red-500/20 bg-red-500/10'
  };
</script>

<div class="card p-5 glow-border">
  <div class="flex items-start justify-between gap-4">
    <div>
      <div class="font-semibold">Activity</div>
      <div class="muted text-sm mt-1">
        Settlement:
        <span class="text-slate-200 font-semibold">{settlementLabel[order.settlement] || order.settlement}</span>
        • Risk:
        <span class="text-slate-200 font-semibold">{order.risk.level}</span>
        <span class="muted">({order.risk.score})</span>
      </div>
    </div>

    {#if order.requiresApproval}
      <span class="badge">
        <span class="font-semibold text-amber-200">Approval required</span>
        <span class="muted">hold active</span>
      </span>
    {/if}
  </div>

  <div class="mt-4 space-y-3">
    {#each order.timeline.slice().reverse() as ev}
      <div class={"rounded-2xl border p-4 " + (typeTone[ev.type] || "border-slate-800 bg-slate-950/30")}>
        <div class="flex items-start justify-between gap-4">
          <div class="text-sm font-semibold capitalize">{ev.type}</div>
          <div class="muted text-xs">{fmt(ev.at)}</div>
        </div>
        <div class="text-sm text-slate-200/90 mt-1">{ev.message}</div>
      </div>
    {/each}
  </div>
</div>
