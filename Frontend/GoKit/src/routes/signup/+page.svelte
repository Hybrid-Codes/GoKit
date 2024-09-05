<script>
    let name = '';
    let email = '';
    let password = '';
    let errors = { name: '', email: '', password: '' };
    let successMessage = '';
    let failMessage = '';

    function validate() {
      errors = { name: '', email: '', password: '' };
      if (!name) errors.name = 'Name is required';
      if (!email) errors.email = 'Email is required';
      if (!password) errors.password = 'Password is required';
    }

    async function signup() {
      validate();
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
      } else {
        console.error('SignUp failed');
        failMessage = 'Incorrect username or password.';
      }
    }
  </script>

  {#if successMessage}
  <div class="p-4 mb-4 text-sm text-green-700 bg-green-100 rounded-lg" role="alert">
    {successMessage}
  </div>
  {/if}

  <!-- {#if failMessage}
  <div class="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg" role="alert">
    {failMessage}
  </div>
  {/if} -->

  <form on:submit|preventDefault={signup} class="max-w-md mx-auto p-4 bg-white shadow-md rounded-lg">
    <h2 class="text-2xl font-bold mb-4">Sign Up</h2>

    <div class="relative z-0 w-full mb-6 group">
      <input type="text" id="name" name="name" bind:value={name} required
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
      <input type="email" id="email" name="email" bind:value={email} required
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
      <input type="password" id="password" name="password" bind:value={password} required
        class="block py-2.5 px-0 w-full text-sm text-gray-900 bg-transparent border-0 border-b-2 border-gray-300 appearance-none focus:outline-none focus:ring-0 focus:border-blue-600 peer placeholder-transparent" placeholder=" " />
      <label for="password"
        class="absolute text-sm text-gray-500 duration-300 transform -translate-y-6 scale-75 top-3 left-0 -z-10 origin-[0] peer-focus:left-0 peer-focus:text-blue-600 peer-placeholder-shown:scale-100 peer-placeholder-shown:translate-y-0 peer-focus:scale-75 peer-focus:-translate-y-6">
        Password
      </label>
      {#if errors.password}
        <p class="text-red-600 text-sm mt-1">{errors.password}</p>
      {/if}
    </div>

    <button type="submit"
      class="w-full px-4 py-2 bg-blue-600 text-white font-semibold rounded-md shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500">
      Sign Up
    </button>
  </form>
