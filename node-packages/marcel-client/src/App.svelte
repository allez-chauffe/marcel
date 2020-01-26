<script>
  import "bulma/css/bulma.css";
  import "marcel-media";
  import Debugger from "./components/debugger.svelte";
  import { Message } from "svelma";
  import machine from "./machine";
  import AuthForm from "./components/auth-form.svelte";
</script>

<style>
  * {
    box-sizing: border-box;
  }
</style>

{#if $machine.state === 'loggedOut'}
  <AuthForm />
{:else if $machine.state === 'noMedia'}
  <Message active title="Attention">
    Aucun media n'est actuellement associé
  </Message>
{:else if $machine.state === 'mediaLoaded'}
  <marcel-media
    media={$machine.context.media}
    uri="{$machine.context.config.apiURI}medias/{$machine.context.media.id}" />
{:else if $machine.state === 'fatalError'}
  <Message active title="Error !">{$machine.context.error}</Message>
{/if}

<Debugger />
