<script>
   import "$lib/styles/ghost_kitchen.css";
  import CartDrawer from '$lib/components/CartDrawer.svelte';
</script>

<div class="min-h-screen">
  <header
    class="sticky top-0 z-30 backdrop-blur border-b"
    style="border-color: rgba(255,255,255,0.08); background: rgba(11,15,20,0.85);"
  >
    <div class="container py-4 flex items-center justify-between gap-4 flex-nowrap">
      <a href="/GhostKitchen" class="flex items-center gap-3 group shrink-0">
        <img
          src="/brand/ghostkitchen-logo-transparent.png"
          alt="Ghost Kitchen+"
          class="h-9 w-auto"
		  style="transform: scale(4.5);"
        />
      </a>

      <nav class="hidden md:flex items-center gap-6 min-w-0">
        <a class="navlink-muted" href="/GhostKitchen">Menu</a>
        <a class="navlink-muted" href="/GhostKitchen/cart">Cart</a>
        <a class="navlink-muted" href="/GhostKitchen/checkout">Checkout</a>
      </nav>

      <div class="shrink-0">
        <CartDrawer />
      </div>
    </div>
  </header>

  <main class="container py-8">
    <slot />
  </main>

  <footer class="container pb-10">
    <div class="hr mt-10"></div>
    <div class="muted text-sm mt-6">
      Ghost Kitchen+ is a demo storefront used to showcase SettleGuard scam resolution and controlled settlement.
    </div>
  </footer>
</div>
