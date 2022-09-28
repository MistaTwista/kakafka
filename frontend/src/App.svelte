<script lang="ts">
    import {
        GetConfigs,
        GetTopics,
        Connect,
        CreateTopic,
        DeleteTopic
    } from "../wailsjs/go/application/Application"
    import type { application, kafka } from "../wailsjs/go/models"
    import CreateTopicDialog from "./components/dialogs/CreateTopicDialog.svelte"
    import DeleteTopicDialog from "./components/dialogs/DeleteTopicDialog.svelte"
    import ProfilesSidebar from "./components/ProfilesSideBar.svelte"
    import TopicList from "./components/TopicList.svelte"

    interface Topic extends kafka.Topic {
        selected: boolean
        system: boolean
    }

    interface Profile extends application.Profile {
        connected: boolean
        topics: Topic[]
    }

    let profiles: Profile[] = []
    let selectedProfile: string
    $: selectedProfileIndex = profiles.findIndex(p => p.name == selectedProfile)
    $: topics = profiles[selectedProfileIndex]?.topics.map(t => ({
        name: t.name,
        system: t.system
    }))

    let topicCreateDialogIsActive: boolean
    let deletingTopic: string | undefined

    getConfigs()

    function getConfigs() {
        GetConfigs().then(cfg => {
            profiles = cfg.profiles.map(p => ({
                ...p,
                connected: false,
                topics: []
            }))
        })
    }

    function showTopicCreateDialog() {
        topicCreateDialogIsActive = true
    }

    function hideTopicCreateDialog() {
        topicCreateDialogIsActive = false
    }

    function showTopicDeleteDialog(e: CustomEvent<number>) {
        deletingTopic = topics[e.detail].name
    }

    function hideTopicDeleteDialog() {
        deletingTopic = undefined
    }

    interface CreateTopicEvent {
        name: string
        partitions: number
        replicas: number
    }

    async function createTopic(e: CustomEvent<CreateTopicEvent>) {
        const topic = await CreateTopic(
            selectedProfile,
            e.detail.name,
            e.detail.partitions,
            e.detail.replicas
        )

        const newTopics = [
            ...profiles[selectedProfileIndex].topics,
            {
                ...topic,
                selected: false,
                system: false
            }
        ]

        newTopics.sort((a, b) => a.name.localeCompare(b.name))

        profiles[selectedProfileIndex].topics = newTopics

        hideTopicCreateDialog()
    }

    async function deleteTopic() {
        await DeleteTopic(selectedProfile, deletingTopic)

        const newTopics = profiles[selectedProfileIndex].topics.filter(
            t => t.name != deletingTopic
        )

        profiles[selectedProfileIndex].topics = newTopics

        hideTopicDeleteDialog()
    }

    async function selectProfile(e: CustomEvent<number>) {
        const profile = profiles[e.detail]

        if (!profile.connected) {
            await Connect(profile.name)
        }

        profiles[e.detail].connected = true

        const list = await GetTopics(profile.name)
        profiles[e.detail].topics = list.map(t => ({
            ...t,
            selected: false,
            system: t.name.startsWith("__")
        }))
        selectedProfile = profile.name
    }
</script>

<main>
    <aside>
        <ProfilesSidebar {profiles} on:select={selectProfile} />
    </aside>
    {#if topics}
        <aside>
            <button on:click={showTopicCreateDialog}>Create topic</button>
            <TopicList {topics} on:delete={showTopicDeleteDialog} />
        </aside>
    {/if}
    <section class="content">content</section>

    {#if topicCreateDialogIsActive}
        <CreateTopicDialog
            on:create={createTopic}
            on:cancel={hideTopicCreateDialog}
        />
    {/if}

    {#if !!deletingTopic}
        <DeleteTopicDialog
            name={deletingTopic}
            on:confirm={deleteTopic}
            on:cancel={hideTopicDeleteDialog}
        />
    {/if}
</main>

<style>
    main {
        display: flex;
        flex-direction: row;
        justify-content: start;
        align-items: stretch;
        height: 100%;
        width: 100%;
    }
</style>
