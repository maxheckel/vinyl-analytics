<template>
    <div class="listen">
        <button class="main-button" v-show="!listening" @click="startListening">
            Listen
        </button>
        <div class="listening" v-show="listening">
            Listening...
        </div>
        <div v-show="!listening">
            <div v-if="listenResponse !== undefined">

            </div>
        </div>
    </div>
</template>
<script>
    import LR from '../services/ListenRepository'
    export default {
        name: "Listen",
        data: function(){
            return {
                listening: false,
                hasListened: false,
                listenResponse: undefined
            }
        },
        methods:{
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