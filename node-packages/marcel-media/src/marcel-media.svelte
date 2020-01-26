<script>
  export let media;

  $: console.log(media);
</script>

<style>
  * {
    box-sizing: border-box;
  }
</style>

<svelte:options tag="marcel-media" />
<div class="container">
  {#if media}
    {#each media.plugins as plugin}
      <div>plugin {plugin.eltName}</div>
    {/each}
  {/if}
</div>
