<script>
  import { Input, Button } from 'svelma';
  import machine from "../machine";
  let login, password;

  $: console.log({login, password})

  const handleLogin = () => {
    console.log(login, password)
    machine.send({ type: "login", login, password })
  };
</script>

<style>
 form {
   display: grid;
   gap: 1em;
   align-items: center;
 }

</style>

{#if $machine.context.error}
  <p>{$machine.context.error}</p>
{/if}

<form on:submit|preventDefault={handleLogin}>
  <Input bind:value={login} />
  <Input type="password" bind:value={password} />
  <Button nativeType="submit" disabled={$machine.state !== 'loggedOut'} type="is-primary">
    Loggin
  </Button>
</form>

