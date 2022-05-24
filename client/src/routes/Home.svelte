<script lang="ts">
    import { userData, isLogin } from "../stores";
</script>

{#if $isLogin}
    <h1>Hello {$userData.username}!</h1>
{:else}
    <h1>Hello Wolrd!</h1>
{/if}
