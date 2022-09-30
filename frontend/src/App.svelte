<script lang="ts">
    import { GetConfigs, GetTopics, Connect, CreateTopic, DeleteTopic, ConsumerOffsets } from "../wailsjs/go/application/Application"
    import type { kafka, application } from "../wailsjs/go/models"
    import CreateTopicDialog from "./components/dialogs/CreateTopicDialog.svelte"
    import DeleteTopicDialog from "./components/dialogs/DeleteTopicDialog.svelte"
    import ProfilesSidebar from "./components/ProfilesSideBar.svelte"
    import TopicList from "./components/TopicList.svelte"
    import Topic from "./components/Topic.svelte"
    import { LogDebug } from "../wailsjs/runtime"

    interface Profile extends application.Profile {
        connected: boolean
        topics: kafka.Topic[]
    }

    let profiles: Profile[] = []
    let selectedProfileIndex: number | undefined
    $: selectedProfile = selectedProfileIndex !== undefined ? profiles[selectedProfileIndex] : undefined

    $: showTopicsBar = !!selectedProfile?.connected
    $: topics = selectedProfile?.topics || []

    let selectedTopicIndex: number | undefined
    let selectedTopicConsumerOffsets: kafka.ConsumerOffset[] | undefined
    $: selectedTopic = selectedTopicIndex !== undefined ? topics[selectedTopicIndex] : undefined

    $: showTopicPanel = selectedTopic !== undefined

    let topicCreateDialogIsActive: boolean

    let indexOfTopicForDelete: number | undefined
    $: topicForDelete = indexOfTopicForDelete !== undefined && topics[indexOfTopicForDelete]
    $: topicDeleteDialogIsActive = indexOfTopicForDelete !== undefined

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
        indexOfTopicForDelete = e.detail
    }

    function hideTopicDeleteDialog() {
        indexOfTopicForDelete = undefined
    }

    async function createTopic(e: CustomEvent<kafka.TopicConfig>) {
        const topic = await CreateTopic(selectedProfile.name, e.detail)

        const newTopicList: kafka.Topic[] = [...topics, topic]

        newTopicList.sort((a, b) => a.name.localeCompare(b.name))

        profiles[selectedProfileIndex].topics = newTopicList

        hideTopicCreateDialog()
    }

    async function deleteTopic() {
        await DeleteTopic(selectedProfile.name, topicForDelete.name)

        const newTopics = topics.filter(t => t.name != topicForDelete.name)

        profiles[selectedProfileIndex].topics = newTopics

        hideTopicDeleteDialog()
    }

    async function selectProfile(e: CustomEvent<number>) {
        selectedTopicConsumerOffsets = undefined

        const profileIndex = e.detail

        const profile = profiles[profileIndex]

        if (!profile.connected) {
            await Connect(profile.name)
        }

        profiles[profileIndex].connected = true

        const topicList = await GetTopics(profile.name, false)

        profiles[profileIndex].topics = topicList

        selectedProfileIndex = profileIndex
    }

    async function selectTopic(e: CustomEvent<number>) {
        const topicIndex = e.detail

        const topic = topics[topicIndex]

        selectedTopicIndex = topicIndex

        const offsets = await ConsumerOffsets(selectedProfile.name, topic.name)
        
        selectedTopicConsumerOffsets = offsets
    }
</script>

<main>
    <aside>
        <ProfilesSidebar {profiles} on:select={selectProfile} />
    </aside>
    {#if showTopicsBar}
        <aside>
            <button on:click={showTopicCreateDialog}>Create topic</button>
            <TopicList {topics} on:select={selectTopic} on:delete={showTopicDeleteDialog} />
        </aside>
    {/if}
    {#if showTopicPanel}
        <Topic topic={selectedTopic} consumerOffsets={selectedTopicConsumerOffsets} />
    {:else}
        <section>select topic...</section>
    {/if}

    {#if topicCreateDialogIsActive}
        <CreateTopicDialog on:create={createTopic} on:cancel={hideTopicCreateDialog} />
    {/if}

    {#if topicDeleteDialogIsActive}
        <DeleteTopicDialog topic={topicForDelete} on:confirm={deleteTopic} on:cancel={hideTopicDeleteDialog} />
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
