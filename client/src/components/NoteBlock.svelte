<script lang="ts">
    import type Note from "src/types/note";
    import { createEventDispatcher } from "svelte";
    const dispatch = createEventDispatcher();

    export let data: Note = {
        uuid: "",
        content: "",
        status: "",
    };

    async function handleDelete() {
        await fetch("http://localhost:5000/api/note", {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                uuid: data.uuid,
            }),
        });

        dispatch("change");
    }

    enum status {
        success = "success",
        active = "active",
    }
    function swapStatus(s: status) {
        switch (s) {
            case status.success:
                return status.active;
            case status.active:
                return status.success;
        }
    }
    async function handleStatus(s: status) {
        await fetch("http://localhost:5000/api/note", {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                uuid: data.uuid,
                status: s,
            }),
        });

        dispatch("change");
    }
</script>

<div class="block">
    <div class="content">
        {data.content}
    </div>
    <div class="right">
        <div
            class="swap"
            on:click={() => {
                handleStatus(swapStatus(status[data.status]));
            }}
        >
            swap
        </div>
        <div class="close" on:click={handleDelete}>X</div>
    </div>
</div>

<style>
    .block {
        background-color: #f0f0f0;
        padding: 1.5rem;
        display: flex;
        width: 30rem;
        justify-content: space-between;
        align-items: center;
        border-radius: 10px;
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
    .right {
        display: flex;
        gap: 1rem;
        align-items: center;
    }
    .swap {
        cursor: pointer;
        background-color: #cbcbcb;
        padding: 0.5rem;
        height: 0.75rem;
        border-radius: 200px;
        display: flex;
        align-items: center;
        justify-content: center;
    }
</style>
