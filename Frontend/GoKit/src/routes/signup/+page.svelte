<script>
  import { goto } from '$app/navigation';
  let name = '';
  let email = '';
  let password = '';

  async function signup() {
      const response = await fetch('http://localhost:8080/signup', { // Use the SvelteKit endpoint instead of directly calling the Go backend
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
          },
          body: JSON.stringify({ name, email, password })
      });

      if (response.ok) {
          const data = await response.json();
          localStorage.setItem('token', data.token); // Save the token in localStorage
          goto('/'); // Navigate to the home page
      } else {
          console.error('SignUp failed', await response.text()); // Log the error message
          // Handle signup failure (e.g., display an error message to the user)
      }
  }
</script>

<form on:submit={signup}>
  <label>
      Name:
      <input type="text" bind:value={name} required />
  </label>
  <label>
      Email:
      <input type="email" bind:value={email} required />
  </label>
  <label>
      Password:
      <input type="password" bind:value={password} required />
  </label>
  <button type="submit">Sign Up</button>
</form>
