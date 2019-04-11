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
            <a class="addListen">+ Listen</a>
            <div class="otherOptions">
                <a class="tryAgain">Try Again</a>
                <a class="manualAdd">Search Manually</a>
            </div>
        </div>

    </div>
</template>

<script>
    export default {
        name: "AlbumNotFound",
        props:{
            listenResponse: Object
        },
        computed: {
            titleNoArtist: function () {
                let s = this.listenResponse.album_title
                s = s.split("-")
                s.shift()
                return s.join("")
            }
        }
    }
</script>

<style scoped>

</style>