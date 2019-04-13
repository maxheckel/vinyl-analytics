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
                    <AlbumNotFound v-on:tryAgain="tryAgain" :listen-response="listenResponse"></AlbumNotFound>
                </div>
                <div v-if="listenResponse.album_found">

                </div>
            </div>
            <div v-if="notFound">
                <h3>Couldn't find album</h3>
            </div>
        </div>
    </div>
</template>
<script>
    import LR from '../services/ListenRepository'
    import AlbumNotFound from "./AlbumNotFound";
    export default {
        name: "Listen",
        data: function(){
            return {
                listening: false,
                hasListened: false,
                listenResponse: undefined,
                notFound: false
            }
        },
        components:{
            AlbumNotFound
        },
        methods:{
            stopListening(){
                this.listenResponse = undefined;
                this.listening = false;
            },
            tryAgain(){
                this.listenResponse = undefined;
                this.startListening()
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
                    LR.listen(data).then( (response) => {
                        self.hasListened = true;
                        self.listening = false;
                        self.listenResponse = response.data;
                        if (response.data.album_found){
                            LR.add(response.data.album_id).then((res)=>{
                                self.$router.push("/album/"+response.data.album_id)
                            })
                        }
                    }).catch((err) => {
                        self.notFound = true
                        self.stopListening()
                    })


                }


            }
        }
    }
</script>

<style scoped>

</style>