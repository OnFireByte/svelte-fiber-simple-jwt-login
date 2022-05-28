<script lang="ts">
    import type Note from "src/types/note";

    export let data: Note = {
        uuid: "",
        content: "",
        status: "",
    };

    async function handleDelete() {
        const res = await fetch("http://localhost:5000/api/note", {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                uuid: data.uuid,
            }),
        });
        console.log(res);

        const json = res.json();
        console.log(json);
    }
</script>

<div class="block">
    <div class="content">
        {data.content}
    </div>
    <div class="close" on:click={handleDelete}>X</div>
</div>

<style>
    .block {
        background-color: #f0f0f0;
        padding: 1.5rem;
        display: flex;
        width: 30rem;
        justify-content: space-between;
        align-items: center;
    }
    .close {
        cursor: pointer;
        background-color: #ff3e00;
        padding: 0.5rem;
        width: 0.75rem;
        height: 0.75rem;
        border-radius: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-weight: 700;
    }
</style>
