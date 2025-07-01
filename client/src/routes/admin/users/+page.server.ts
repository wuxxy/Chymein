import {AuthedInstance, axiosInstance} from "$lib/utils/axios";

export const load = async ({request}) => {
    try {
        return {
            fetchUsers: AuthedInstance(request).get('/Admin/users')
        };
    } catch (err) {
        console.error("Axios Error:", err.response?.data || err.message);
        // handle or rethrow
        throw new Error("Unauthorized access to /Admin/users");
    }
}