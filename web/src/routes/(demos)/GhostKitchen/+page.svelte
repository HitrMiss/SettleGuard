<script>
  import menu from '$lib/data/menu.json';
  import MenuItemCard from '$lib/components/MenuItemCard.svelte';
  import OrderSummary from '$lib/components/OrderSummary.svelte';
  import { cartCount, cartSubtotal } from '$lib/stores/cart.js';

  const { restaurant, categories, items } = menu;

  let active = categories[0]?.id || 'featured';
  let q = '';

  const labelFor = (id) => {
    if (id === 'featured') return 'Featured';
    if (id === 'burgers') return 'Burgers';
    if (id === 'sides') return 'Sides';
    if (id === 'drinks') return 'Drinks';
    if (id === 'dessert') return 'Dessert';
    return 'Menu';
  };

  $: filtered = items.filter((x) => {
    const matchesCat = active ? x.categoryId === active : true;
    const matchesQ =
      !q ||
      x.name.toLowerCase().includes(q.toLowerCase()) ||
      x.description.toLowerCase().includes(q.toLowerCase());
    return matchesCat && matchesQ;
  });

  $: count = $cartCount;
  $: subtotal = $cartSubtotal;
</script>

<section class="card p-6 sm:p-8">
  <div class="text-xs uppercase tracking-widest muted">Ghost Kitchen+ / {restaurant.name}</div>

  <div class="mt-3 grid gap-6 lg:grid-cols-[1.2fr_0.8fr] lg:items-end">
    <div>
      <h1 class="text-3xl sm:text-5xl font-extrabold tracking-tight">Ordering with style.</h1>
      <p class="muted mt-3 max-w-2xl">
        A demo ordering experience that feels real—then showcases SettleGuard resolution after checkout.
      </p>

      <div class="mt-5 flex flex-col sm:flex-row gap-3">
        <a class="btn btn-primary" href="/GhostKitchen/checkout">Checkout</a>
        <a class="btn" href="/GhostKitchen/cart">View cart</a>
      </div>
    </div>

    <div class="card-soft p-5">
      <div class="text-lg font-extrabold">{restaurant.name}</div>
      <div class="muted text-sm mt-2">
        {restaurant.etaMin}-{restaurant.etaMax} min • Delivery ${restaurant.deliveryFee.toFixed(2)} • Service ${restaurant.serviceFee.toFixed(2)}
      </div>
      <div class="muted text-xs mt-2">{restaurant.promo}</div>

      <div class="mt-4">
        <input class="input" placeholder="Search menu…" bind:value={q} />
      </div>
    </div>
  </div>
</section>

<!-- ✅ Key change: lg:items-start prevents the row from stretching left column based on right column height -->
<section class="mt-6 grid gap-6 lg:grid-cols-[minmax(0,1fr)_360px] lg:items-start">
  <!-- ✅ Key change: content-start prevents the left grid from distributing extra height as extra gap -->
  <div class="grid gap-4 min-w-0 content-start self-start">
    <!-- FIXED-HEIGHT CATEGORY STRIP (no layout shift) -->
    <div
      class="card-soft sticky z-20"
      style="
        top: 96px;
        height: 64px;
        padding: 10px;
        display: flex;
        align-items: center;
        overflow: hidden;
      "
    >
      <div
        style="
          display: flex;
          gap: 8px;
          overflow-x: auto;
          overflow-y: hidden;
          white-space: nowrap;
          width: 100%;
          padding-bottom: 2px;
        "
      >
        {#each categories as cat}
          <button
            class={"pill " + (active === cat.id ? "pill-active" : "")}
            on:click={() => (active = cat.id)}
            type="button"
            style="height: 44px; flex: 0 0 auto;"
          >
            {labelFor(cat.id)}
          </button>
        {/each}
      </div>
    </div>

    <div class="grid gap-3 sm:grid-cols-2">
      {#each filtered as item}
        <MenuItemCard {item} />
      {/each}
    </div>
  </div>

  <!-- ✅ self-start so the aside never forces stretching behavior -->
  <aside class="hidden lg:block self-start" style="width: 360px;">
    <div class="sticky self-start" style="top: 96px; width: 360px;">
      <OrderSummary />
    </div>
  </aside>
</section>

{#if count > 0}
  <div class="fixed bottom-3 left-1/2 -translate-x-1/2 z-30 w-[min(560px,calc(100vw-24px))] lg:hidden">
    <div class="card-soft px-3 py-2 flex items-center justify-between gap-2" style="background: rgba(12,15,20,0.78); backdrop-filter: blur(10px);">
      <div class="min-w-0">
        <div class="font-extrabold">Your order</div>
        <div class="muted text-xs">{count} item(s) • ${subtotal.toFixed(2)} subtotal</div>
      </div>
      <a class="btn btn-primary whitespace-nowrap" href="/GhostKitchen/checkout">Checkout</a>
    </div>
  </div>
{/if}
