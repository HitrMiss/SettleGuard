<script>
  import { cart } from '$lib/stores/cart.js';
  export let item;

  function add() {
    cart.add({ id: item.id, name: item.name, price: item.price, category: item.categoryId });
  }
</script>

<article class="card p-4 flex gap-4">
  <div class="thumb w-24 h-24 shrink-0 grid place-items-center">
    <div class="text-sm font-semibold muted">Item</div>
  </div>

  <div class="min-w-0 flex-1 flex flex-col">
    <div class="flex items-start justify-between gap-3">
      <div class="min-w-0">
        <div class="font-semibold truncate">{item.name}</div>
        <div class="muted text-sm mt-1 line-clamp-2">{item.description}</div>
      </div>
      <div class="font-semibold">${item.price.toFixed(2)}</div>
    </div>

    <div class="mt-3 flex items-center justify-between">
      <span class="badge">
        <span class="font-semibold" style="color: rgb(var(--accent));">Popular</span>
        <span class="muted">choice</span>
      </span>

      <button class="btn" on:click={add} aria-label={`Add ${item.name}`}>
        Add
      </button>
    </div>
  </div>
</article>

<style>
  .line-clamp-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
</style>
