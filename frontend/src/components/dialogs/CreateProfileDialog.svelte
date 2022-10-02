<script lang="ts">
    import Dialog from "./Dialog.svelte"
    import { createEventDispatcher } from "svelte"
    import type { application } from "wailsjs/go/models"

    const dispatch = createEventDispatcher()

    let name: string
    let brokers: string

    function create() {
        dispatch("create", {
            name,
            brokers: brokers
                .split(",")
                .map(b => b.trim())
                .filter(b => b !== "")
        } as application.Profile)
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
            <p>Brokers:</p>
            <section><input type="text" bind:value={brokers} /></section>
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
