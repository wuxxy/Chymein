import { axiosInstance } from '$lib/utils/axios';
import type { UserType } from '$lib/utils/types';
import { redirect, type ServerLoad } from '@sveltejs/kit';

export const load: ServerLoad = async ({ request }) => {
    try {
        const res = await axiosInstance.get('/Auth/me', {
            headers: {
                Cookie: request.headers.get('cookie') ?? ''
            }
        });

        const user = res.data as UserType;

        if (!user?.Username || user.Admin == false) {
            return redirect(302, '/login');
        }

        return { user };
    } catch (err) {
        throw redirect(302, '/login');
    }
};
