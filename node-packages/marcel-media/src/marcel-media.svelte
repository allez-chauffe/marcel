<script>
  import _ from "immutadot";
  export let media, uri;

  $: col = media && 100 / media.cols;
  $: row = media && 100 / media.rows;
  $: pluginsURI = `${uri}/plugins`;

  const pluginStyle = ({ frontend: { x, y, rows, cols } }) => {
    console.log({ x, y, cols, rows, col, row });
    return `left: ${x * col}%; top: ${y * row}%; width: ${cols *
      col}%; height: ${rows * row}%;`;
  };

  const pluginURI = ({ eltName, instanceId }) =>
    `${pluginsURI}/${eltName}/${instanceId}`;

  $: console.log(uri, col, row, pluginsURI, media);

  window.setPluginPosition = position => {
    media = _.update(media, "plugins[0].frontend", frontend => ({
      ...frontend,
      ...position
    }));
  };
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
  .plugin {
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
      <iframe
        title="{plugin.eltName}#{plugin.instanceId}"
        class="plugin"
        style={pluginStyle(plugin)}
        src="{pluginsURI}/{plugin.eltName}/{plugin.instanceId}/" />
    {/each}
  </div>
{/if}
