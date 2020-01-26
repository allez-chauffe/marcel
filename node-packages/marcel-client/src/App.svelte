<script>
  import "bulma/css/bulma.css";
  import Media from "marcel-media";
  import { Message } from "svelma";
  import machine from "./machine";
  import AuthForm from "./components/auth-form.svelte";
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

<div class="container">
  <div class="content">
    {#if $machine.state === 'loggedOut'}
      <AuthForm />
    {:else if $machine.state === 'noMedia'}
      <Message active title="Attention">
        Aucun media n'est actuellement associé
      </Message>
    {:else if $machine.state === 'mediaLoaded'}
      <Media media={$machine.context.media} />
    {:else if $machine.state === 'fatalError'}
      <Message active title="Error !">{$machine.context.error}</Message>
    {/if}
  </div>

  <pre class="debugger">
    State: {$machine.state} {'\n'}Context: {JSON.stringify($machine.context, 0, 2)}
  </pre>
</div>
