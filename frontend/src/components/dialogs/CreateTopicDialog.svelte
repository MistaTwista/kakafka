<script lang="ts">
    import Dialog from "./Dialog.svelte"
    import { createEventDispatcher } from "svelte"
    import type { kafka } from "../../../wailsjs/go/models"

    const dispatch = createEventDispatcher()

    let name: string
    let numPartitions: string
    let replicationFactor: string

    function create() {
        dispatch("create", {
            topic: name,
            numPartitions: Number(numPartitions),
            replicationFactor: Number(replicationFactor)
        } as kafka.TopicConfig)
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
            <p>Number of Partitions:</p>
            <section><input type="text" bind:value={numPartitions} /></section>
        </section>
        <section>
            <p>Replication Factor:</p>
            <section>
                <input type="text" bind:value={replicationFactor} />
            </section>
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
