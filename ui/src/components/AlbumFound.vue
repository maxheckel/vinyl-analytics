<template>
    <div class="albumFound">
        {{listenResponse.album_title}}
    </div>
</template>

<script>
    export default {
        name: "AlbumFound",
        props:{
            listenResponse: Object
        }
    }
</script>

<style scoped>

</style>