<script lang="ts">
    import { createEventDispatcher } from "svelte"

    const dispatch = createEventDispatcher()

    interface Item {
        name: string
        connected: boolean
    }

    export let profiles: Item[]

    let collapsed: boolean = false
    let selectedItemIndex: number

    function select(index: number) {
        selectedItemIndex = index
        dispatch("select", index)
    }

    function setSize() {
        collapsed = !collapsed
    }
</script>

<section class={collapsed && "collapsed"}>
    <ul>
        {#each profiles as item, index}
            <li
                class={selectedItemIndex == index && "selected"}
                on:click={() => {
                    select(index)
                }}
            >
                <span class="icon {item.connected && "highlight"}" />
                {#if !collapsed}
                    <span class="name">{item.name}</span>
                {/if}
            </li>
        {/each}
    </ul>
    <span class="collapser" on:click={setSize}>{collapsed ? "»" : "«"}</span>
</section>

<style>
    section {
        height: 100%;
        width: 250px;
        background-color: rgb(53, 53, 53);
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }

    section.collapsed {
        width: 70px;
    }

    .collapser {
        margin: 0;
        cursor: pointer;
    }

    .collapser:hover {
        background-color: rgb(43, 43, 43);
    }

    ul {
        margin: 10px 0;
        padding: 0;
    }

    li {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
        cursor: pointer;
        padding: 0 10px;
    }

    li.selected {
        background-color: rgb(31, 31, 31);
    }

    li:hover {
        background-color: rgb(71, 71, 71);
    }

    .icon {
        display: inline-block;
        width: 50px;
        height: 50px;
        background-color: rgb(46, 117, 117);
        border-radius: 5px;
    }

    .icon.highlight {
        border: 1px solid white;
    }

    .name {
        display: inline-block;
        flex-grow: 1;
        margin-left: 10px;
        text-align: left;
    }
</style>
