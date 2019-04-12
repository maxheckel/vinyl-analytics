import Repository from './Repository'
const resource = "albums";
export default {

    create(discogsId){
        let request = {
            discogs_id: discogsId
        }
        return Repository.post(resource, request)
    },

    find(albumId){
       return Repository.get(`${resource}/${albumId}`)
    }


}