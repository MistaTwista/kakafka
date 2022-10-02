<script lang="ts">
    import type { kafka } from "wailsjs/go/models"

    export let topic: kafka.Topic
    export let consumerOffsets: kafka.ConsumerOffset[]
</script>

<section>
    <p>Name: {topic.name}</p>
    <p>Internal: {topic.internal}</p>
    <p>Partitions:</p>
    <table>
        <thead>
            <tr>
                <th>id</th>
                <th>replicas</th>
            </tr>
        </thead>
        <tbody>
            {#each topic.partitions as partition}
                <tr>
                    <td>
                        {partition.id}
                    </td>
                    <td>{partition.replicas.map(r => `${r.host}:${r.port}`).join(", ")}</td>
                </tr>
            {/each}
        </tbody>
    </table>
    {#if consumerOffsets !== undefined && consumerOffsets.length > 0}
        <p>Offsets:</p>
        <table>
            <thead>
                <tr>
                    <th>consumer</th>
                    <th>offsets</th>
                </tr>
            </thead>
            <tbody>
                {#each consumerOffsets as consumerOffset}
                    <tr>
                        <td>
                            {consumerOffset.consumer}
                        </td>
                        <td>
                            <table>
                                <thead>
                                    <tr>
                                        <th>partition</th>
                                        <th>offset</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#each consumerOffset.offsets as offsets}
                                        <tr>
                                            <td>
                                                {offsets.partition}
                                            </td>
                                            <td>
                                                {offsets.committedOffset}
                                            </td>
                                        </tr>
                                    {/each}
                                </tbody>
                            </table>
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {/if}
</section>

<style>
    section {
        height: 100%;
        width: 100%;
        background-color: rgb(23, 23, 23);
        display: flex;
        flex-direction: column;
    }
</style>
