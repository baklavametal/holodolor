<script>
    let events = [];

    // Function passed down from App.svelte to update authentication state
    export let updateAuthentication;

    import { view } from "../store.js";

    const fetchEvents = async () => {
        const token = localStorage.getItem("token");
        if (!token) return;

        try {
            const response = await fetch("http://localhost:8088/events", {
                method: "GET",
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
            });

            if (response.ok) {
                events = await response.json();
            } else {
                console.error("Failed to fetch events");
            }
        } catch (error) {
            console.error("An error occurred:", error);
        }
    };

    const logout = () => {
        localStorage.removeItem("token");
        view.set("landing"); // Navigate back to landing page
        updateAuthentication(); // Update authentication state
    };

    // Fetch events when the component mounts
    fetchEvents();
</script>

<h1>Dashboard</h1>
<button on:click={logout}>Logout</button>
<ul>
    {#each events as event}
        <li>
            <strong>{event.type}</strong> at {event.timestamp}
            {#if event.intensity}
                <span> (Intensity: {event.intensity})</span>
            {/if}
        </li>
    {/each}
</ul>
