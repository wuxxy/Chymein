import { axiosInstance } from '$lib/utils/axios';
import type { NavBarType, StatusType } from '$lib/utils/types';

export const load = async () => {
    try {
        const [statusRes, navRes] = await Promise.all([
            axiosInstance.get('/status'),
            axiosInstance.get('/data/navbar')
        ]);

        return {
            status: statusRes.data as StatusType,
            nav: navRes.data as NavBarType
        };
    } catch (err) {
        console.error('Failed to fetch app data:', err);

        const fallbackStatus: StatusType = {
            database: false,
            is_setup: false,
            connected: false,
            port: 'unknown',
            time_alive: '0s'
        };

        return { status: fallbackStatus };
    }
};
