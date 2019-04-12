<template>
    <div class="album">
        <a class="homeLink" href="/"><i class="fas fa-chevron-left"></i></a>

        <div v-if="album != undefined">

            <div v-if="album.status === 404">
                Could not find album
            </div>
            <div v-if="album.status === 200">
                <div class="albumInfo">
                    <img class="spinning" src="../assets/playing.gif">
                    <img v-bind:src="album.data.artwork">
                    <h1>{{album.data.name}}</h1>
                    <h3>{{album.data.artist.name}}</h3>
                </div>
                <div class="listens">
                    Listens: {{album.data.listens.length}}
                </div>
                <table class="tracklist">
                    <tr v-for="track in album.data.tracks">
                        <td class="position">
                            {{track.position}}
                        </td>
                        <td class="name">
                            {{track.title}}
                        </td>
                        <td class="duration">
                            {{track.duration}}
                        </td>
                    </tr>
                </table>
            </div>
        </div>

    </div>
</template>

<script>
    import AlbumRepository from "../services/AlbumRepository";

    export default {
        mounted:function () {
            AlbumRepository.find(this.$route.params.id).then((result) => {
                console.log(result)
                this.album = result
            }).catch((err) => {
                this.album = {
                    status: 404
                }
            })
        },
        name: "Album.vue",
        data:function () {
            return {
                album: undefined
            }
        }
    }
</script>

<style scoped>

</style>