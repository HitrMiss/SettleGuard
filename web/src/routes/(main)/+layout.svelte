<script>
    let { children } = $props();
</script>

<nav class="bg-gray-900 text-white p-4">
    <b>Admin Panel</b>
    <a href="/">Home</a>
    <a href="/users">Users</a>
</nav>

<main class="p-8">
    {@render children()}
</main>