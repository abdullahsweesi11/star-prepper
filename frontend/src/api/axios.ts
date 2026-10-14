import {type AxiosInstance} from "axios"
import axios from "axios"

const api: AxiosInstance = axios.create({
    baseURL: "http://localhost:5000/api",       // TODO: change
    withCredentials: true
})

export default api