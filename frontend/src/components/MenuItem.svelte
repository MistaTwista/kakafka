<script lang="ts">
    import { createEventDispatcher } from "svelte"

    const dispatch = createEventDispatcher()

    function click() {
        dispatch("click")
    }
</script>

<section on:click={() => click()}>
    <slot />
</section>

<style>
    section {
        padding: 5px;
        background-color: rgb(34, 34, 34);
        border-radius: 5px;
    }

    section:hover {
        background-color: rgb(46, 46, 46);
    }
</style>
