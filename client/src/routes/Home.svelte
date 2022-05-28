<script lang="ts">
    import App from "src/App.svelte";
    import type Note from "src/types/note";
    import { onMount } from "svelte";
</script>

{#if $isLogin}
    <h1>Hello {$userData.username}!</h1>
{:else}
    <h1>Hello Wolrd!</h1>
{/if}
