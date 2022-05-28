<script lang="ts">
    import { ErrorResponseType as errRes } from "./../lib/enum/errorResponseType";
    import { navigate,useFocus } from "svelte-navigator";
    import { userData, isLogin } from "../stores";

    const registerFocus = useFocus();
    let username = "";
    let password = "";
    let haveError = false;
    let errorMessage = "";

    export async function handleSubmit() {
        const res = await fetch("http://localhost:5000/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                username,
                password,
            }),
        });
        const data = await res.json();
        console.log(data);

        if (data?.status && data?.status !== "login_success") {
            errorHandler(data.status);
            return;
        }
        haveError = false;

        const res1 = await fetch("http://localhost:5000/api/user", {
            method: "GET",
            credentials: "include",
        });
        const data1 = await res1.json();
        if (data1?.status) {
            return;
        }
        userData.set(data1);
        isLogin.set(true);
        navigate("/");
    }

    function errorHandler(status: string) {
        haveError = true;
        switch (status) {
            case errRes.ErrorEmailExists:
                errorMessage = "Email already exists";
                break;

            case errRes.ErrorUsernameExists:
                errorMessage = "Username already exists";
                break;

            case errRes.ErrorPasswordNotMatch:
                errorMessage = "Password not match";
                break;

            case errRes.ErrorInvalidEmail:
                errorMessage = "Invalid email";
                break;

            default:
                errorMessage = "Unknown error";
                break;
        }
    }
</script>

<form on:submit|preventDefault={handleSubmit}>
    <input
        use:registerFocus
        type="text"
        name="username"
        placeholder="Username"
        bind:value={username}
        required
    />
    <input type="password" name="password" placeholder="Password" bind:value={password} required />
    <button type="submit">Login</button>
    {#if haveError}
        <div class="error">{errorMessage}</div>
    {/if}
</form>

<style>
    form {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        width: 100%;
        max-width: 500px;
        margin: 0 auto;
        gap: 2rem;
    }
    input {
        width: 100%;
        padding: 1rem;
        border-radius: 5px;
        border: 1px solid #ccc;
        font-size: 1.2rem;
        font-weight: 300;
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
    .error {
        color: red;
        font-size: 1.2rem;
    }
</style>
