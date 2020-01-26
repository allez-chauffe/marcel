<script>
  import "marcel-widget";

  export let media, uri;

  $: col = media && 100 / media.cols;
  $: row = media && 100 / media.rows;
  $: pluginsURI = `${uri}/plugins`;

  const pluginStyle = ({ frontend: { x, y, rows, cols } }) =>
    `left: ${x * col}%; top: ${y * row}%; width: ${cols *
      col}%; height: ${rows * row}%;`;
</script>

<style>
  :host,
  .media {
    position: relative;
    width: 100%;
    height: 100%;
  }
  * {
    box-sizing: border-box;
  }
  .widget {
    position: absolute;
    border: 1px solid red;
    transition: width 1s ease-in-out, height 1s ease-in-out, left 1s ease-in-out,
      top 1s ease-in-out;
  }
</style>

<svelte:options tag="marcel-media" />

{#if media}
  <div class="media">
    {#each media.plugins as plugin}
      <marcel-widget
        style={pluginStyle(plugin)}
        class="widget"
        {plugin}
        stylesvar={media.stylesvar}
        uri="{pluginsURI}/{plugin.eltName}/{plugin.instanceId}/" />
    {/each}
  </div>
{/if}
