<template>
    <div class="albumNotFound">
        <div class="info">
            <div class="art">
                <img v-bind:src="listenResponse.image">
            </div>
            <div class="details">
                <span class="albumTitle">{{titleNoArtist}}</span>
                <span class="artistTitle">{{listenResponse.artist_title}}</span>
                <span class="songTitle">{{listenResponse.song_title}}</span>
            </div>
        </div>
        <div class="controls">
            <div v-show="!addingListen">
                <a @click="createAlbumAndAddListen" class="addListen">+ Listen</a>

                <div class="otherOptions">
                    <a @click="tryAgain" class="tryAgain">Try Again</a>
                    <a class="manualAdd">Search Manually</a>
                </div>
            </div>
            <div v-show="addingListen">
                <h1>Adding listen...</h1>
            </div>

        </div>

    </div>
</template>

<script>
    import AlbumRepository from "../services/AlbumRepository";
    import ListenRepository from "../services/ListenRepository";

    export default {
        name: "AlbumNotFound",
        props:{
            listenResponse: Object
        },
        data: function(){
            return {
                addingListen: false
            }
        },
        computed: {
            titleNoArtist: function () {
                let s = this.listenResponse.album_title
                s = s.split("-")
                s.shift()
                return s.join("")
            }
        },
        methods: {
            tryAgain(){
                this.$emit('tryAgain')
            },
            createAlbumAndAddListen(){
                this.addingListen = true
                AlbumRepository.create(this.listenResponse.discogs_id).then((result) => {
                    ListenRepository.add(result.data.id).then(() => {
                        this.$router.push("/album/"+result.data.id)
                    })
                })
            }
        }
    }
</script>

<style scoped>

</style>