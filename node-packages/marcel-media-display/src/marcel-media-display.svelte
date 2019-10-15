<script>
  import "bulma/css/bulma.css";
  import { setContext } from "svelte";
  import {
    config,
    api,
    user,
    client,
    media,
    connection,
    loadConfig
  } from "./stores";
  import AuthForm from "./components/auth-form.svelte";
  import * as toast from "./utils/toast";

  // Props //
  export let apiuri, websocketuri;
  // Props //

  $: loadConfig({ apiURI: apiuri, websocketURI: websocketuri });
  $: if ($user) $api.loadClient();
  $: if ($client) $api.loadMedia($client.mediaID);
</script>

<svelte:options tag="marcel-media-display" />

{#if !$user}
  <AuthForm />
{:else if $client && !$client.mediaID}
  <p>Aucun média n'est associé à ce client</p>
{:else if $media}
  <p>Media</p>
{:else}
  <p>loading...</p>
{/if}

{#if $connection}
  <p>Client connected</p>
{/if}
