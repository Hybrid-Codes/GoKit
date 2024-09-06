<script lang="ts">
  import { goto } from "$app/navigation";
  import { onMount } from 'svelte'; // You may want to import this if needed elsewhere

  let name: string = '';
  let email: string = '';
  let password: string = '';
  let confirmPassword: string = '';
  let errors: { name: string; email: string; password: string; confirmPassword: string } = {
    name: '',
    email: '',
    password: '',
    confirmPassword: '',
  };
  let successMessage: string = '';
  let showAlert: boolean = false; // Control visibility of the alert

  function validateField(field: string) {
    if (field === 'name') {
      if (!name) {
        errors.name = 'Name is required';
      } else {
        errors.name = '';
      }
    }

    if (field === 'email') {
      if (!email) {
        errors.email = 'Email is required';
      } else {
        errors.email = '';
      }
    }

    if (field === 'password') {
      if (!password) {
        errors.password = 'Password is required';
      } else {
        errors.password = '';
      }
    }

    if (field === 'confirmPassword') {
      if (!confirmPassword) {
        errors.confirmPassword = 'Please confirm password';
      } else if (password && confirmPassword !== password) {
        errors.confirmPassword = 'Passwords do not match';
      } else {
        errors.confirmPassword = '';
      }
    }
  }

  function validateAll() {
    validateField('name');
    validateField('email');
    validateField('password');
    validateField('confirmPassword');
  }

  async function signup() {
    validateAll();

    if (Object.values(errors).some(error => error)) return;

    const response = await fetch('http://localhost:8080/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, email, password }),
    });

    if (response.ok) {
      const { token } = await response.json();
      localStorage.setItem('token', token);
      successMessage = 'You have successfully signed up!';
      showAlert = true;

      // Automatically hide the alert after 5 seconds (5000ms)
      setTimeout(() => {
        showAlert = false;
        successMessage = ''; // Optional: Clear success message
      }, 5000); // Adjust the delay as needed (5000ms = 5 seconds)
      // goto("/");
    } else {
      console.error('SignUp failed');
    }
  }
</script>

{#if showAlert}
  <div class="alert bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded relative">
    <strong class="font-bold">Success!</strong>
    <span class="block sm:inline">{successMessage}</span>
  </div>
{/if}

<form on:submit|preventDefault={signup} class="max-w-md mx-auto p-4 bg-white shadow-md rounded-lg">
  <h2 class="text-2xl font-bold mb-4">Sign Up</h2>

  <div class="relative z-0 w-full mb-6 group">
    <input type="text" id="name" name="name" bind:value={name} on:input={() => validateField('name')} required
      class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer placeholder-transparent" placeholder=" " />
    <label for="name"
      class="absolute text-sm text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 left-0 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6">
      Name
    </label>
    {#if errors.name}
      <p class="text-red-600 text-sm mt-1">{errors.name}</p>
    {/if}
  </div>

  <div class="relative z-0 w-full mb-6 group">
    <input type="email" id="email" name="email" bind:value={email} on:input={() => validateField('email')} required
      class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer placeholder-transparent" placeholder=" " />
    <label for="email"
      class="absolute text-sm text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 left-0 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6">
      Email
    </label>
    {#if errors.email}
      <p class="text-red-600 text-sm mt-1">{errors.email}</p>
    {/if}
  </div>

  <div class="relative z-0 w-full mb-6 group">
    <input type="password" id="password" name="password" bind:value={password} on:input={() => validateField('password')} required
      class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer placeholder-transparent" placeholder=" " />
    <label for="password"
      class="absolute text-sm text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 left-0 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6">
      Password
    </label>
    {#if errors.password}
      <p class="text-red-600 text-sm mt-1">{errors.password}</p>
    {/if}
  </div>

  <div class="relative z-0 w-full mb-6 group">
    <input type="password" id="confirm-password" name="confirm-password" bind:value={confirmPassword} on:input={() => validateField('confirmPassword')} required
      class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer placeholder-transparent" placeholder=" " />
    <label for="confirm-password"
      class="absolute text-sm text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 left-0 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6">
      Confirm Password
    </label>
    {#if errors.confirmPassword}
      <p class="text-red-600 text-sm mt-1">{errors.confirmPassword}</p>
    {/if}
  </div>

  <button type="submit"
    class="w-full px-4 py-2 bg-blue-600 text-white font-semibold rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
    Sign Up
  </button>
</form>
