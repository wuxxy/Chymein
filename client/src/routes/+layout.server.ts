import { axiosInstance } from '$lib/axios';

type StatusType = {
    database: boolean;
    is_setup: boolean;
    port: string;
    time_alive: string;
};

export const load = async () => {
    try {
        const { data } = await axiosInstance.get('/status');
        return {
            status: data as StatusType
        };
    } catch (err) {
        console.error('Failed to fetch status:', err);

        // Fallback default
        const fallback: StatusType = {
            database: false,
            is_setup: false,
            port: 'unknown',
            time_alive: '0s'
        };

        return {
            status: fallback
        };
    }
};
