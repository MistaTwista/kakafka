<script lang="ts">
    import { createEventDispatcher } from "svelte"

    const dispatch = createEventDispatcher()

    interface Item {
        name: string
        system: boolean
    }

    export let topics: Item[]

    function select(index: number) {
        dispatch("select", index)
    }

    function deleteTopic(index: number) {
        dispatch("delete", index)
    }
</script>

<div>
    <ul>
        {#each topics as item, index}
            <li
                on:click={() => {
                    select(index)
                }}
            >
                <span>{item.name}</span>
                {#if !item.system}
                    <button on:click={() => deleteTopic(index)}>x</button>
                {/if}
            </li>
        {/each}
    </ul>
</div>

<style>
    div {
        height: 100%;
        width: 100%;
        background-color: rgb(53, 53, 53);
        display: flex;
        flex-direction: column;
    }

    ul {
        margin: 0;
        padding: 0;
    }

    li {
        display: block;
        cursor: pointer;
        padding: 10px;
        text-align: left;
    }

    li:hover {
        background-color: rgb(71, 71, 71);
    }

    span {
        text-align: left;
    }
</style>
