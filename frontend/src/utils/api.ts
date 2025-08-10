import axios, {AxiosRequestConfig} from 'axios';

export async function apiFetch(url: string, options: AxiosRequestConfig = {}) {
    const requestID = crypto.randomUUID();

    return axios({
        url,
        method: options.method || 'GET',
        ...options,
        headers: {
            ...(options.headers || {}),
            'X-Request-ID': requestID,
            'Content-Type': 'application/json',
        },
    });
}
