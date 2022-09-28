<script lang="ts">
    import Dialog from "./Dialog.svelte"
    import { createEventDispatcher } from "svelte"

    const dispatch = createEventDispatcher()

    let name: string
    let partitions: string
    let replicas: string

    function create() {
        dispatch("create", {
            name,
            partitions: Number(partitions),
            replicas: Number(replicas)
        })
    }

    function cancel() {
        dispatch("cancel")
    }
</script>

<section>
    <Dialog>
        <section>
            <p>Name:</p>
            <section><input type="text" bind:value={name} /></section>
        </section>
        <section>
            <p>Partitions:</p>
            <section><input type="text" bind:value={partitions} /></section>
        </section>
        <section>
            <p>Replicas:</p>
            <section><input type="text" bind:value={replicas} /></section>
        </section>
        <button on:click={create}>Create</button>
        <button on:click={cancel}>Cancel</button>
    </Dialog>
</section>

<style>
    section {
        z-index: 2000;
        background-color: rgb(53, 53, 53);
        border-radius: 10px;
    }
</style>
