<script lang="ts">
    import { navigate,useFocus } from "svelte-navigator";

    const registerFocus = useFocus();

    import { ErrorResponseType as errRes } from "./../lib/enum/errorResponseType";
    let username = "";
    let email = "";
    let password = "";
    let confirmPassword = "";
    let haveError = false;
    let errorMessage = "";

    export async function handleSubmit() {
        const res = await fetch("http://localhost:5000/api/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",

            body: JSON.stringify({
                username,
                email,
                password,
                confirmPassword,
            }),
        });
        const data = await res.json();

        if (data?.status) {
            errorHandler(data.status);
            return;
        }
        haveError = false;
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
    <input type="text" name="email" placeholder="Email" bind:value={email} required />
    <input type="password" name="password" placeholder="Password" bind:value={password} required />
    <input
        type="password"
        name="confirmPassword"
        placeholder="Confirm Password"
        bind:value={confirmPassword}
        required
    />
    <button type="submit">Register</button>
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
