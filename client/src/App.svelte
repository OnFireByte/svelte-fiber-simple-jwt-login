<script lang="ts">
    import { onMount } from "svelte";
    import { Link, Route, Router, navigate } from "svelte-navigator";
    import Home from "src/routes/Home.svelte";
    import Register from "src/routes/Register.svelte";
    import Login from "src/routes/Login.svelte";
    import { userData, isLogin } from "./stores";

    onMount(async () => {
        const res = await fetch("http://localhost:5000/api/user", {
            method: "GET",
            credentials: "include",
        });
        const data = await res.json();
        if (data?.status) {
            return;
        }
        userData.set(data);
        isLogin.set(true);
        navigate("/");
    });
    async function logoutHandler() {
        const _ = await fetch("http://localhost:5000/api/logout", {
            method: "GET",
            credentials: "include",
        });
        console.log("test");

        userData.set(null);
        isLogin.set(false);
        navigate("/");
    }
</script>

<svelte:head>
    <title>Svelte App</title>
    <meta name="viewport" content="width=device-width, initial-scale=1" />
</svelte:head>
<Router>
    <main>
        <nav>
            <Link to="/">Home</Link>
            {#if $isLogin}
                <button on:click={logoutHandler}>Logout</button>
            {:else}
                <Link to="login">Login</Link>
                <Link to="register">Register</Link>
            {/if}
        </nav>
        <div>
            <Route path="/" primary={false}>
                <Home />
            </Route>
            <Route path="register">
                <Register />
            </Route>
            <Route path="login">
                <Login />
            </Route>
        </div>
    </main>
</Router>

<style>
    :root {
        font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu,
            Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
    }

    main {
        text-align: center;
        padding: 1em;
        margin: 0 auto;
    }

    img {
        height: 16rem;
        width: 16rem;
    }

    h1 {
        color: #ff3e00;
        text-transform: uppercase;
        font-size: 4rem;
        font-weight: 100;
        line-height: 1.1;
        margin: 2rem auto;
        max-width: 14rem;
    }
    button {
        padding: 1rem;
        border-radius: 5px;
        border: 1px solid #ccc;
        font-size: 1.2rem;
        font-weight: 300;
        background: #ff3e00;
        color: #fff;
        cursor: pointer;
    }

    p {
        max-width: 14rem;
        margin: 1rem auto;
        line-height: 1.35;
    }
    nav {
        display: flex;
        justify-content: space-evenly;
        margin: 1rem auto;
        margin-bottom: 3rem;
    }
</style>
