<script>
    let email = '';
    let password = '';

    async function handleLogin() {
      // Handle login logic here
    }
  </script>

  <form on:submit|preventDefault={handleLogin}>
    <input type="email" bind:value={email} placeholder="Email" required />
    <input type="password" bind:value={password} placeholder="Password" required />
    <button type="submit">Login</button>
  </form>
