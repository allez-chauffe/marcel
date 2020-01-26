<script>
  export let media;
</script>

<style>
  .container {
    display: flex;
  }
  .content {
    flex: 1;
  }
  .debugger {
    font-size: 10px;
    width: 40em;
  }
</style>

<svelte:options tag="marcel-media" />

<div class="container">
  <div class="content">
    {#if media}
      <p>Here is the media !</p>
    {:else}
      <p>No media</p>
    {/if}
  </div>

  <pre class="debugger">{media}</pre>
</div>
