<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import type { application } from "wailsjs/go/models"
    import MenuItem from "./MenuItem.svelte"
    import PopupMenu from "./PopupMenu.svelte"

    const dispatch = createEventDispatcher()

    interface Profile extends application.Profile {
        connected: boolean
    }

    export let profiles: Profile[]

    let collapsed: boolean = false
    let selectedItemIndex: number

    let popupMenuIsActive: boolean = false
    let popupMenuCoords = { x: 0, y: 0 }
    let popupMenuForProfile: Profile

    function select(index: number) {
        selectedItemIndex = index
        dispatch("select", index)
    }

    function setSize() {
        collapsed = !collapsed
    }

    function createProlile() {
        dispatch("create-profile")
    }

    function deleteProfile() {
        dispatch("delete-profile", popupMenuForProfile)
        hidePopupMenu()
    }

    function showPopupMenu(event: MouseEvent, profile: Profile) {
        event.stopImmediatePropagation()
        event.preventDefault()

        popupMenuCoords = {
            x: event.x,
            y: event.y
        }

        popupMenuForProfile = profile
        popupMenuIsActive = true
    }

    function hidePopupMenu() {
        popupMenuIsActive = false
    }
</script>

<section class="sidebar {collapsed ? 'collapsed' : ''}">
    <ul class="profiles">
        {#each profiles as profile, index}
            <li
                class={selectedItemIndex == index && "selected"}
                on:click={() => {
                    select(index)
                }}
                on:contextmenu={e => showPopupMenu(e, profile)}
            >
                {#if profile.connected}
                    <span class="connected" />
                {/if}
                <span class="icon" />
                {#if !collapsed}
                    <span class="name">{profile.name}</span>
                {/if}
            </li>
        {/each}
    </ul>
    <section>
        <ul class="manage">
            <li on:click={() => createProlile()}>
                <span class="icon" />
                {#if !collapsed}
                    <span class="name">Add Profile</span>
                {/if}
            </li>
        </ul>
        <span class="collapser" on:click={setSize}>{collapsed ? "»" : "«"}</span>
    </section>

    {#if popupMenuIsActive}
        <PopupMenu x={popupMenuCoords.x} y={popupMenuCoords.y} on:close={hidePopupMenu}>
            <MenuItem on:click={deleteProfile}>Удалить</MenuItem>
        </PopupMenu>
    {/if}
</section>

<style>
    section {
        display: flex;
        flex-direction: column;
        justify-content: start;
    }

    .sidebar {
        height: 100%;
        width: 250px;
        background-color: rgb(53, 53, 53);
        position: relative;
        user-select: none;
    }

    section.collapsed {
        width: 70px;
    }

    .profiles {
        flex-grow: 1;
        overflow-y: auto;
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
        padding: 5px 10px;
        position: relative;
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
        box-sizing: border-box;
        background-color: rgb(46, 117, 117);
        border-radius: 5px;
    }

    .connected {
        width: 5px;
        height: 10px;
        background-color: aquamarine;
        border-top-right-radius: 5px;
        border-bottom-right-radius: 5px;
        position: absolute;
        left: 0px;
    }

    .name {
        display: inline-block;
        flex-grow: 1;
        margin-left: 10px;
        text-align: left;
    }
</style>
