<script>
  import { api, user } from "../stores";
  let userName, password, loading, error;

  const load = task => async () => {
    loading = true;
    try {
      await task();
    } catch (err) {
      console.error(err);
      error = err.message;
    } finally {
      loading = false;
    }
  };

  const login = load(() => $api.login(userName, password));

  const refreshAuth = load(() => $api.refreshAuth());

  $: if (!$user && $api) refreshAuth();
</script>

<svelte:options tag="marcel-media-diplay-auth" />

{#if error}
  <pre>{error}</pre>
{/if}

<input disabled={loading} bind:value={userName} />
<input type="password" bind:value={password} />
<button disabled={loading} on:click={login}>Loggin</button>
