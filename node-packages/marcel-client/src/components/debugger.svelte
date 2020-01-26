<script>
  import machine from "../machine";

  let show = true;

  window.toggleMarcelDebugger = () => (show = !show);
</script>

<style>
  .debugger {
    position: absolute;
    top: 0;
    right: 0;
  }
  pre {
    font-size: 10px;
    width: 40em;
    height: 100vh;
  }
  button {
    width: 100%;
  }
</style>

{#if show}
  <div class="debugger">
    <button on:click={window.toggleMarcelDebugger}>Hide debugguer</button>
    <pre>
      State: {$machine.state} {'\n'}Context: {JSON.stringify($machine.context, 0, 2)}
    </pre>
  </div>
{/if}
