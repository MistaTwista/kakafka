<script lang="ts">
    import { createEventDispatcher, onMount, onDestroy } from "svelte"

    const dispatch = createEventDispatcher()

    export let x: number
    export let y: number
    let horizontal: string
    let vertical: string

    $: style = `${horizontal}; ${vertical}`

    function mouseDown(ev: MouseEvent) {
        const self = document.getElementById("popup-menu")

        if ((ev.target as Node) === self || self.contains(ev.target as Node)) {
            return
        }

        dispatch("close")
    }

    onMount(async () => {
        window.addEventListener("mousedown", mouseDown)

        const self = document.getElementById("popup-menu")

        horizontal = x + self.clientWidth < document.body.clientWidth ? `left: ${x}px` : `left: ${x - self.clientWidth}px`
        vertical = y + self.clientHeight < document.body.clientHeight ? `top: ${y}px` : `top: ${y - self.clientHeight}px`
    })

    onDestroy(async () => {
        window.removeEventListener("mousedown", mouseDown)
    })
</script>

<section id="popup-menu" {style}>
    <slot />
</section>

<style>
    section {
        position: fixed;
        padding: 5px;
        background-color: rgb(34, 34, 34);
        border-radius: 5px;
        user-select: none;
        z-index: 3000;
    }
</style>
