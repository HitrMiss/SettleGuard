<script>
   import "$lib/styles/ghost_kitchen.css";
  import CartDrawer from '$lib/components/CartDrawer.svelte';
</script>

<div class="min-h-screen">
  <header class="sticky top-0 z-30 border-b border-white/10 bg-black/40 backdrop-blur">
    <div class="container py-4 flex items-center justify-between gap-4">
      <a href="/" class="flex items-center gap-3">
        <div class="logo">
          <span class="font-black tracking-tight text-sm">GK+</span>
        </div>
        <div class="leading-tight">
          <div class="font-semibold tracking-tight">Ghost Kitchen+</div>
          <div class="muted text-xs">Delivery � Dark-mode by design</div>
        </div>
      </a>

      <nav class="hidden sm:flex items-center gap-5">
        <a class="navlink-muted" href="/">Menu</a>
        <a class="navlink-muted" href="/cart">Cart</a>
        <a class="navlink-muted" href="/checkout">Checkout</a>
      </nav>

      <CartDrawer />
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
