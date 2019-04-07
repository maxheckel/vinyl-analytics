<template>
    <div class="listen">
        <button v-show="!listening" @click="startListening">
            Listen
        </button>
        <div class="listening" v-show="listening">
            Listening...
        </div>
        <div v-show="!listening">
            <div v-if="listenResponse.metadata.music[0] !== undefined && listenResponse.metadata.music !== null">
                I heard {{listenResponse.metadata.music[0].title}}
            </div>
            <div v-if="hasListened && listenResponse.metadata.music === null ">
                Could not determine song
            </div>
        </div>
    </div>
</template>
<script>
    export default {
        name: "Listen",
        data: function(){
            return {
                listening: false,
                hasListened: false,
                listenResponse: {
                    metadata:{
                        music:[]
                    }
                }
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
                    var fd = new FormData();
                    fd.append('fname', 'test.wav');
                    fd.append('data', data);

                    var request = new XMLHttpRequest();
                    request.open("POST", "http://localhost:8081/listen");
                    request.send(fd);

                    request.onreadystatechange = function() {
                        if (request.readyState == 4 && request.status == 200)
                            self.hasListened = true;
                            self.listening = false;
                            var response = JSON.parse(request.responseText)
                            self.listenResponse = response
                    }


                }


            }
        }
    }
</script>

<style scoped>

</style>