<script lang="ts">
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();
    let content = "";

    async function handleSubmit() {
        await fetch(`${import.meta.env.VITE_SERVER_URL}/api/note`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                content,
            }),
        });
        content = "";
        dispatch("submit");
    }
</script>

<form on:submit|preventDefault={handleSubmit}>
    <input type="text" name="content" placeholder="Content" bind:value={content} />
    <input type="submit" value="Add Note" class="btn" />
</form>

<style>
    form {
        display: flex;
        gap: 1rem;
        justify-content: center;
    }
    input {
        width: 30rem;
        padding: 1rem;
        border-radius: 5px;
        border: 1px solid #ccc;
        font-size: 1.2rem;
        font-weight: 300;
    }
    .btn {
        padding: 1rem;
        border-radius: 5px;
        border: 1px solid #ccc;
        font-size: 1.2rem;
        font-weight: 300;
        background: #ff3e00;
        color: #fff;
        cursor: pointer;
        width: 10rem;
    }
</style>
