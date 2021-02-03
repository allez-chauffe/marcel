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
   display: flex;
   flex-direction: column;
   align-items: center;
 }

.input {
  margin: 0.2em;
}
</style>

{#if $machine.context.error}
  <p>{$machine.context.error}</p>
{/if}

<form on:submit|preventDefault={handleLogin}>
  <Input class="input" bind:value={login} />
  <Input class="input" type="password" bind:value={password} />
  <Button nativeType="submit" disabled={$machine.state !== 'loggedOut'}>
    Loggin
  </Button>
</form>

