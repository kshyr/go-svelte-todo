<script lang="ts">
import { onMount } from "svelte";
const endpoint = "http://localhost:3000"
let tasks = [];

  onMount(async function () {
    const response = await fetch(endpoint);
    const data = await response.json();
    tasks = data
  });
</script>

<main class="min-h-screen bg-neutral-900 flex justify-center items-center">
  <section class="max-w-2xl w-full bg-neutral-700">
  {#each tasks as task}
    <div class="flex gap-4 justify-around"><span>Title: {task.title}</span> <span>Status: {task.completed}</span></div>
  {/each}
  </section>
  
</main>

<style lang="postcss">

</style>
