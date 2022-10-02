<script lang="ts">
    import Dialog from "./Dialog.svelte"
    import { createEventDispatcher } from "svelte"
    import type { application, kafka } from "wailsjs/go/models"

    const dispatch = createEventDispatcher()

    export let profile: application.Profile

    function confirm() {
        dispatch("confirm")
    }

    function cancel() {
        dispatch("cancel")
    }
</script>

<section>
    <Dialog>
        <section>
            <p>Delete profile {profile.name}</p>
        </section>
        <button on:click={confirm}>Delete</button>
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
