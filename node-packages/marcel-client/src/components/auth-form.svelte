<script>
  import machine from "../machine";
  let login, password;

  function handleLoginClick() {
    machine.send({ type: "login", login, password });
  }
</script>

{#if $machine.context.error}
  <p>{$machine.context.error}</p>
{/if}

<input bind:value={login} />
<input type="password" bind:value={password} />
<button disabled={$machine.state !== 'loggedOut'} on:click={handleLoginClick}>
  Loggin
</button>
