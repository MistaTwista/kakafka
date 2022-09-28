<script lang="ts">
    import Overlay from "../Overlay.svelte"
</script>

<section>
    <section class="window-container">
        <section class="window">
            <slot />
        </section>
    </section>
    <Overlay />
</section>

<style>
    .window-container {
        z-index: 2000;
        display: flex;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        justify-content: center;
        align-items: center;
    }

    .window {
        z-index: 2001;
        background-color: rgb(53, 53, 53);
        border-radius: 10px;
    }
</style>
