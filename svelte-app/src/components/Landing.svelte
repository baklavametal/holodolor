<script>
    import { view } from "../store.js";

    const showLogin = () => {
        view.set("login");
    };

    const showRegister = () => {
        view.set("register");
    };
</script>

<h1>HOLODOLOR</h1>
<button on:click={showLogin}>Login</button>
<button on:click={showRegister}>Register</button>
