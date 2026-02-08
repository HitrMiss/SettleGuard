<script>
  export let trustTier = 'new';
  export let bondLevel = 0;

  const tier = {
    verified: { label: 'Verified', fg: 'rgba(60,255,210,0.95)', bg: 'rgba(60,255,210,0.10)', bd: 'rgba(60,255,210,0.22)' },
    bonded:   { label: 'Bonded',   fg: 'rgba(155,120,255,0.95)', bg: 'rgba(155,120,255,0.10)', bd: 'rgba(155,120,255,0.22)' },
    new:      { label: 'New',      fg: 'rgba(255,199,102,0.95)', bg: 'rgba(255,199,102,0.08)', bd: 'rgba(255,199,102,0.22)' }
  };

  $: t = tier[trustTier] || { label: 'Unknown', fg: 'rgba(234,240,255,0.9)', bg: 'rgba(255,255,255,0.06)', bd: 'rgba(255,255,255,0.12)' };
</script>

<span
  class="inline-flex items-center gap-2 rounded-full px-3 py-1 text-xs border"
  style={"color:" + t.fg + "; background:" + t.bg + "; border-color:" + t.bd}
>
  <span class="font-semibold">{t.label}</span>
  <span class="text-slate-100/70">Bond {bondLevel}</span>
</span>
