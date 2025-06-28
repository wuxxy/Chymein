import {axiosInstance} from "$lib/axios";

export const load = (async()=>{
    return {
        status: axiosInstance.get('/status'),
    }
})