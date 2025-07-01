import axios from "axios";

export const axiosInstance = axios.create({
    baseURL: "http://localhost:6969",   // Consider moving to .env
    withCredentials: true,              // Required for cookies
    timeout: 8000,                      // Prevent hanging requests
    headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
    },
});

// Optional: Interceptors for error logging & response normalization
axiosInstance.interceptors.response.use(
    response => response,
    error => {
        // Basic debug logging
        if (error.response) {
            console.error("Axios Error:", {
                url: error.config?.url,
                status: error.response.status,
                data: error.response.data,
            });
        } else if (error.request) {
            console.error("No response received:", error.request);
        } else {
            console.error("Axios config error:", error.message);
        }

        return Promise.reject(error);
    }
);
export const AuthedInstance =  (request: { headers: { get: (arg0: string) => any; }; }) => axiosInstance.create({
    headers: {
        Cookie: request.headers.get('cookie') ?? ''
    },
})
