<script>
  const nav = [
    { label: 'Product', href: '#product' },
    { label: 'Partners', href: '#partners' }
  ];

  const bullets = [
    'Control settlement finality without slowing checkout.',
    'Reduce fraud and refund abuse using hold windows.',
    'Policy-driven routing based on trust and behavior.',
    'Auditability for operations and support teams.'
  ];
</script>

<!--
  Goal: Everything (header/nav/button/dividers/sections) should not extend beyond the
  right edge of the "largest content box" (the Product bullets card = max-w-3xl).
  So we use a single narrower container for the whole page.
-->
<div class="min-h-screen text-white" style="background: var(--sg-bg);">
  <div class="pointer-events-none fixed inset-0 sg-bg"></div>

  <!-- HEADER -->
  <header class="sticky top-0 z-40 backdrop-blur-xl bg-[rgba(7,12,22,0.88)] border-b border-white/10">
    <div class="sg-topline"></div>

    <!-- Use a tight container (max-w-3xl) so header never exceeds content width -->
    <div class="mx-auto max-w-3xl px-5 sm:px-6 h-20 flex items-center justify-between gap-6">
      <!-- BIG logo -->
      <a href="/" class="flex items-center">
        <img
          src="/brand/settleguard-logo.png"
          alt="SettleGuard"
          class="h-32 sm:h-40 md:h-48 w-auto select-none translate-y-2 -ml-10"
          draggable="false"
        />
      </a>

      <div class="flex items-center gap-6">
        <nav class="hidden md:flex items-center gap-8 text-sm text-white/70">
          {#each nav as item}
            <a class="hover:text-white transition" href={item.href}>{item.label}</a>
          {/each}
        </nav>

        <a class="sg-btn sg-btn-primary" href="/GhostKitchen">View demo →</a>
      </div>
    </div>
  </header>

  <main class="relative" id="top">
    <!-- Divider -->
    <div class="mx-auto max-w-3xl px-5 sm:px-6 mt-10">
      <div class="sg-divider"></div>
    </div>

    <!-- HERO -->
    <section class="mx-auto max-w-3xl px-5 sm:px-6 scroll-mt-28 py-16">
      <div class="max-w-2xl">
        <div class="sg-pill">
          <span class="sg-dot"></span>
          Payment safety layer
        </div>

        <h1 class="mt-6 text-4xl sm:text-5xl font-semibold tracking-tight leading-[1.08]">
          Fraud & refund protection
          <span class="block text-white/70">for instant payment experiences.</span>
        </h1>

        <p class="mt-6 text-lg text-white/70 leading-relaxed">
          SettleGuard wraps your existing payment flow with a settlement control layer.
          Payments feel instant while finality remains safe and configurable.
        </p>

        <div class="mt-10 flex flex-col sm:flex-row gap-4">
          <a class="sg-btn sg-btn-primary" href="#product">View product</a>
          <a class="sg-btn sg-btn-secondary" href="#partners">View partners →</a>
        </div>

        <div class="mt-10"></div>
      </div>
    </section>

    <!-- PRODUCT line -->
    <div class="mx-auto max-w-3xl px-5 sm:px-6">
      <div class="sg-divider"></div>
    </div>

    <!-- PRODUCT -->
    <section id="product" class="mx-auto max-w-3xl px-5 sm:px-6 scroll-mt-28 py-16">
      <div class="max-w-2xl">
        <div class="text-xs uppercase tracking-[0.20em] text-white/55">Product</div>
        <h2 class="mt-3 text-2xl sm:text-3xl font-semibold tracking-tight">
          Clean controls that fit your stack
        </h2>
        <p class="mt-4 text-white/70 leading-relaxed">
          Keep your processors and rails. SettleGuard adds configurable settlement timing
          where fraud and refund risk demand it.
        </p>
      </div>

      <!-- This is the "largest content box" width we are matching everywhere -->
      <div class="mt-12 sg-card p-8 max-w-3xl">
        <ul class="space-y-5 text-white/75 leading-relaxed">
          {#each bullets as b}
            <li class="flex gap-4">
              <span class="mt-[0.55rem] h-1.5 w-1.5 rounded-full bg-[rgba(34,211,238,0.85)] flex-none"></span>
              <span>{b}</span>
            </li>
          {/each}
        </ul>
      </div>

      <div class="mt-12"></div>
    </section>

    <!-- PARTNERS line -->
    <div class="mx-auto max-w-3xl px-5 sm:px-6">
      <div class="sg-divider"></div>
    </div>

    <!-- PARTNERS -->
    <section id="partners" class="mx-auto max-w-3xl px-5 sm:px-6 scroll-mt-28 py-16">
      <div class="max-w-2xl">
        <div class="text-xs uppercase tracking-[0.20em] text-white/55">Partners</div>
        <h2 class="mt-3 text-2xl sm:text-3xl font-semibold tracking-tight">
          Demonstrated in real flows
        </h2>
        <p class="mt-4 text-white/70 leading-relaxed">
          SettleGuard behavior is validated inside practical checkout experiences
          before broader integration.
        </p>
      </div>

      <div class="mt-12 sg-card p-8 flex flex-col sm:flex-row items-start sm:items-center justify-between gap-6 max-w-3xl">
        <div>
          <div class="text-sm text-white/60">Demonstration partner</div>
          <div class="mt-1 text-lg font-semibold">GhostKitchen+</div>
          <div class="mt-1 text-sm text-white/70">
            Storefront checkout demo illustrating staged settlement behavior.
          </div>
        </div>

        <a class="sg-btn sg-btn-secondary" href="/GhostKitchen">View demo →</a>
      </div>

      <div class="mt-10"></div>
    </section>

    <footer class="border-t border-white/10 bg-black/15 mt-12">
      <div class="mx-auto max-w-3xl px-5 sm:px-6 py-10 flex items-center justify-between">
        <div class="text-sm text-white/55">
          © {new Date().getFullYear()} SettleGuard. All rights reserved.
        </div>
      </div>
    </footer>
  </main>
</div>
