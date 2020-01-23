<script>
  import "bulma/css/bulma.css";
  import machine from "./stores/machine";
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

<svelte:options tag="marcel-media-display" />

<pre>{$machine.state} {JSON.stringify($machine.context, 0, 2)}</pre>

{#if $machine.state === 'loggedOut'}
  <AuthForm />
{/if}
