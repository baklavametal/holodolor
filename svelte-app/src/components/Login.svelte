<script>
    import { view } from "../store.js";
    import { updateAuthentication } from "../App.svelte";

    let username = "";
    let password = "";

    const login = async () => {
        const payload = {
            username,
            password,
        };

        try {
            const response = await fetch("http://localhost:8088/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(payload),
            });

            const data = await response.json();

            if (data.token) {
                console.log("Login successful");
                localStorage.setItem("token", data.token);
                updateAuthentication();
                view.set("dashboard");
            } else {
                console.log("Login failed");
            }
        } catch (error) {
            console.log("An error occurred:", error);
        }
    };
</script>

<h1>Login</h1>
<form on:submit|preventDefault={login}>
    <div>
        <label for="username">Username:</label>
        <input id="username" type="text" bind:value={username} />
    </div>
    <div>
        <label for="password">Password:</label>
        <input id="password" type="password" bind:value={password} />
    </div>
    <button type="submit">Login</button>
</form>
