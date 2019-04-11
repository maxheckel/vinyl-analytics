import axios from "axios"

export default {
    listen(data){
        var fd = new FormData();
        fd.append('fname', 'test.wav');
        fd.append('data', data);

        return axios({
            method: 'post',
            url: process.env.VUE_APP_API_BASE_URL+'/listen',
            data: fd,
            config: { headers: {'Content-Type': 'multipart/form-data' }}
        })
    }
}