<script>
    let email = '';
    let password = '';

    async function handleSignup() {
      // Handle signup logic here
    }
  </script>

  <form on:submit|preventDefault={handleSignup}>
    <input type="email" bind:value={email} placeholder="Email" required />
    <input type="password" bind:value={password} placeholder="Password" required />
    <button type="submit">Signup</button>
  </form>
