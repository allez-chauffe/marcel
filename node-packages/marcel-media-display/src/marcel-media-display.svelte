<script>
  import "bulma/css/bulma.css";
  import { Message } from "svelma";
  import machine from "./machine";
  import { debounce } from "./utils/function";
  import AuthForm from "./components/auth-form.svelte";

  // Props //
  export let apiuri, websocketuri;
  // Props //

  const configChanged = debounce(config =>
    machine.send({ type: "configChanged", data: config })
  );

  $: configChanged({ apiURI: apiuri, websocketURI: websocketuri });
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

<link
  rel="stylesheet"
  href="https://cdn.jsdelivr.net/npm/bulma@0.8.0/css/bulma.min.css" />

<svelte:options tag="marcel-media-display" />

<div class="container">
  <div class="content">
    {#if $machine.state === 'loggedOut'}
      <AuthForm />
    {:else if $machine.state === 'noMedia'}
      <p>Aucun media n'est actulement associé</p>
    {:else if $machine.state === 'mediaLoaded'}
      <p>Le media !</p>
    {:else if $machine.state === 'fatalError'}
      <Message active title="Error !">{$machine.context.error}</Message>
    {/if}
  </div>

  <pre class="debugger">
    State: {$machine.state} {'\n'}Context: {JSON.stringify($machine.context, 0, 2)}
  </pre>
</div>
