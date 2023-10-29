<script>
	import { view } from "./store.js";
	import Landing from "./components/Landing.svelte";
	import Register from "./components/Register.svelte";
	import Login from "./components/Login.svelte";
	import Dashboard from "./components/Dashboard.svelte";

	let authenticated = false;
	let token;

	// Function to force a re-render
	export const updateAuthentication = () => {
		token = localStorage.getItem("token");
		authenticated = !!token;
	};

	// Initialize
	updateAuthentication();
</script>

<main>
	{#if authenticated}
		<Dashboard {updateAuthentication} />
	{:else if $view === "landing"}
		<Landing />
	{:else if $view === "login"}
		<Login />
	{:else if $view === "register"}
		<Register />
	{/if}
</main>
