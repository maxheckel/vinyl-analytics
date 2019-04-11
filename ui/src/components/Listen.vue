<template>
    <div class="listen">
        <div v-show="listenResponse == undefined">
            <button class="main-button" v-show="!listening" @click="startListening">
                Listen
            </button>
            <div class="listening" v-show="listening">
                <div class="lds-ring"><div></div><div></div><div></div><div></div></div>
            </div>
        </div>
        <div v-show="!listening">
            <div v-if="listenResponse !== undefined">
                <i class="fas fa-chevron-left" @click="stopListening"></i>
                <div v-if="listenResponse.album_found === false">
                    <AlbumNotFound :listen-response="listenResponse"></AlbumNotFound>
                </div>
                <div v-if="listenResponse.album_found">
                    <AlbumFound :listen-response="listenResponse"></AlbumFound>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
    import LR from '../services/ListenRepository'
    import AlbumFound from './AlbumFound'
    import AlbumNotFound from "./AlbumNotFound";
    export default {
        name: "Listen",
        data: function(){
            return {
                listening: false,
                hasListened: false,
                listenResponse: undefined
            }
        },
        components:{
            AlbumNotFound,
            AlbumFound
        },
        methods:{
            stopListening(){
                this.listenResponse = undefined;
                this.listening = false;
            },
            startListening(){
                let id = val => document.getElementById(val)

                navigator.mediaDevices.getUserMedia({audio:true})
                    .then(stream => {handlerFunction(stream)})
                var self = this
                function handlerFunction(stream) {
                    var rec = new MediaRecorder(stream);
                    rec.start();
                    self.listening = true;

                    setTimeout(function () {
                        rec.stop();
                    }, 12000)
                    var audioChunks = [];
                    rec.ondataavailable = e => {
                        audioChunks.push(e.data);
                        if (rec.state == "inactive"){
                            let blob = new Blob(audioChunks,{type:'audio/mpeg-3'});
                            sendData(blob)
                        }
                    }
                }
                function sendData(data) {
                    LR.listen(data).then(function (response) {
                        self.hasListened = true;
                        self.listening = false;
                        console.log(response.data)
                        self.listenResponse = response.data;
                    })


                }


            }
        }
    }
</script>

<style scoped>

</style>