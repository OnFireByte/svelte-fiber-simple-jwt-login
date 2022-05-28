<script lang="ts">
    import AddNote from "src/components/AddNote.svelte";
    import NoteBlock from "src/components/NoteBlock.svelte";

    import { userData, isLogin } from "src/stores";
    import type Note from "src/types/note";
    import { onMount } from "svelte";

    let notes = new Map<string, Note[]>();

    $: noteActive = notes.get("active") || [];
    $: noteSuccess = notes.get("success") || [];

    export async function fetchNote() {
        const res = await fetch("http://localhost:5000/api/note?group=true", {
            method: "GET",
            credentials: "include",
        });
        const data = await res.json();
        if (data?.status) {
            return null;
        }
        notes = new Map<string, Note[]>(Object.entries(data));
    }

    onMount(fetchNote);
</script>

{#if $isLogin}
    <h1>Hello {$userData.username}!</h1>
    <AddNote on:submit={fetchNote} />
    <h2>Active</h2>
    <div class="note-container">
        {#each noteActive as data}
            <NoteBlock {data} on:change={fetchNote} />
        {/each}
    </div>
    <h2>Success</h2>
    <div class="note-container">
        {#each noteSuccess as data}
            <NoteBlock {data} on:change={fetchNote} />
        {/each}
    </div>
{:else}
    <h1>Hello Wolrd!</h1>
{/if}

<style>
    .note-container {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        align-items: center;
    }
</style>
