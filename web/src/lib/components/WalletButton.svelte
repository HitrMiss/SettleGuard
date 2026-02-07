<script>
  import { wallet } from "$lib/stores/wallet.js";
  import { connectMetaMask, hasMetaMask } from "$lib/wallet/metamask.js";

  let busy = false;
  let error = "";

  async function connect() {
    error = "";
    busy = true;
    try {
      const { provider, address, chainId } = await connectMetaMask();
      window.__gk_provider = provider;
      wallet.set({ connected: true, address, chainId });
    } catch (e) {
      error = e?.message ?? String(e);
    } finally {
      busy = false;
    }
  }

  function disconnect() {
    // MetaMask doesn't support programmatic disconnect — we clear app state
    window.__gk_provider = null;
    wallet.set({ connected: false, address: null, chainId: null });
  }

  $: w = $wallet;
  $: short = w.address ? `${w.address.slice(0, 6)}…${w.address.slice(-4)}` : "";
</script>

{#if !w.connected}
  <button class="btn-primary" type="button" disabled={busy || !hasMetaMask()} on:click={connect}>
    {hasMetaMask() ? (busy ? "Connecting..." : "Connect MetaMask") : "Install MetaMask"}
  </button>

  {#if error}
    <div class="muted text-xs mt-2" style="color: rgba(255, 138, 30, 0.95);">{error}</div>
  {/if}
{:else}
  <div class="flex items-center gap-2">
    <span class="badge">
      <span class="muted">Wallet</span>
      <span class="font-semibold">{short}</span>
      <span class="muted">·</span>
      <span class="muted">Chain</span>
      <span class="font-semibold">{w.chainId}</span>
    </span>

    <button class="btn-quiet" type="button" on:click={disconnect}>
      Disconnect
    </button>
  </div>
{/if}
