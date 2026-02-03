<script>
  import menu from '$lib/data/menu.json';
  import MenuItemCard from '$lib/components/MenuItemCard.svelte';
  import OrderSummary from '$lib/components/OrderSummary.svelte';
  import { cartCount, cartSubtotal } from '$lib/stores/cart.js';

  const { restaurant, categories, items } = menu;

  let active = categories[0]?.id || 'featured';
  let q = '';

  const catIcon = (id) => {
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

<section class="section">
  <div class="shell p-6 sm:p-8">
    <div class="kicker">Built to deliver. Built to resolve.</div>

    <div class="mt-3 grid gap-6 lg:grid-cols-[1.2fr_0.8fr] lg:items-end">
      <div>
        <h1 class="h1">Ordering with discipline.</h1>
        <p class="muted mt-3 max-w-2xl">
          Ghost Kitchen+ is a demo ordering experience designed to feel real�then demonstrate SettleGuard
          settlement controls after checkout (hold, approve, unwind).
        </p>

        <div class="mt-5 flex flex-col sm:flex-row gap-3">
          <a class="btn-primary" href="/checkout">Start order</a>
          <a class="btn" href="/cart">View cart</a>
        </div>
      </div>

      <div class="card-soft p-5">
        <div class="h3">{restaurant.name}</div>
        <div class="muted text-sm mt-2">
          {restaurant.etaMin}-{restaurant.etaMax} min � Delivery ${restaurant.deliveryFee.toFixed(2)} � Service ${restaurant.serviceFee.toFixed(2)}
        </div>
        <div class="muted text-xs mt-2">{restaurant.promo}</div>

        <div class="mt-4">
          <input class="input" placeholder="Search menu�" bind:value={q} />
        </div>
      </div>
    </div>
  </div>
</section>

<section class="section pt-0">
  <div class="grid gap-6 lg:grid-cols-[1fr_22rem]">
    <div class="grid gap-4">
      <div class="shell p-3 sticky top-[110px] z-20">
        <div class="flex items-center gap-2 overflow-auto">
          {#each categories as cat}
            <button
              class={"pill " + (active === cat.id ? "pill-active" : "")}
              on:click={() => (active = cat.id)}
            >
              {catIcon(cat.id)}
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

    <aside class="hidden lg:block">
      <div class="sticky top-[110px]">
        <OrderSummary />
      </div>
    </aside>
  </div>
</section>

{#if count > 0}
  <div class="mobile-orderbar lg:hidden">
    <div class="shell mobile-orderbar-inner">
      <div class="min-w-0">
        <div class="font-semibold">Your order</div>
        <div class="muted text-xs">{count} item(s) � ${subtotal.toFixed(2)} subtotal</div>
      </div>
      <a class="btn-primary whitespace-nowrap" href="/checkout">Checkout</a>
    </div>
  </div>
{/if}
