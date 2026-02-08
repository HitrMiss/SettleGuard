<script>
  import { cart } from '$lib/stores/cart.js';
  export let item;

  function add() {
    cart.add({ id: item.id, name: item.name, price: item.price, category: item.categoryId });
  }
</script>

<article class="card card-hover p-4 flex gap-4">
	<div class="thumb w-24 h-24 shrink-0 overflow-hidden">
	  {#if item.image}
		<img
		  src={item.image}
		  alt={item.name}
		  class="w-full h-full object-cover"
		  loading="lazy"
		  on:error={(e) => (e.currentTarget.style.display = 'none')}
		/>
	  {:else}
		<div class="w-full h-full grid place-items-center">
		  <div class="text-sm font-extrabold muted">Item</div>
		</div>
	  {/if}
	</div>


  <div class="min-w-0 flex-1 flex flex-col">
    <div class="flex items-start justify-between gap-3">
      <div class="min-w-0">
        <div class="font-extrabold truncate">{item.name}</div>
        <div class="muted text-sm mt-1 line-clamp-2">{item.description}</div>
      </div>
      <div class="font-extrabold">${item.price.toFixed(2)}</div>
    </div>

    <div class="mt-3 flex items-center justify-between gap-3">
      <span class="inline-flex items-center gap-2 text-xs px-3 py-1 rounded-full"
        style="background: rgba(255,138,30,0.10); border: 1px solid rgba(255,138,30,0.22); color: rgba(230,235,242,0.95);">
        <span style="color:#FF8A1E; font-weight:900;">Popular</span>
        <span class="muted">choice</span>
      </span>

      <button class="btn btn-primary w-28" on:click={add} type="button" aria-label={`Add ${item.name}`}>
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
