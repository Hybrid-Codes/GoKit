<script lang="ts">
	import { goto } from "$app/navigation";

    let email = '';
    let password = '';

    async function login() {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password })
        });

        if (response.ok) {
            const { token } = await response.json();
            localStorage.setItem('token', token); // Save the token in localStorage
            // Navigate to the home page
            goto('/');
            console.log("Login Successful!")
        } else {
            console.error('Login failed');
            // Handle login failure (e.g., display an error message)
        }
    }
</script>

<form on:submit|preventDefault={login}>
    <input type="email" bind:value={email} placeholder="Email" required />
    <input type="password" bind:value={password} placeholder="Password" required />
    <button type="submit">Login</button>
</form>
