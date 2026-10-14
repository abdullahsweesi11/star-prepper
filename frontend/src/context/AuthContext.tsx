import api from "../api/axios"
import { createContext, useContext, useEffect, useState } from "react"

const AuthContext = createContext(null)

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        api.get("/auth/me")
            .then((res) => setUser(res.data.user))
            .catch(() => setUser(null))
            .finally(() => setLoading(false))
    }, [])

    const login = async (email: string, password: string): Promise<{success: boolean}> => {
        try {
            const response = await api.post("/auth/login", { email, password })

            const user = response?.data?.user
            if (!user) {
                console.error("Login failed: Invalid response received from server")
                return {success: false}
            }
            setUser(response.data.user)
            console.log("Login successful")
            return {success: true}
        }
        catch (error) {
            const errorMessage = error.response?.data?.message || error.message || "Unexpected error occurred"
            console.error("Login failed:", errorMessage)

            return {success: false}
        }
    }

    const register = async (email: string, password: string): Promise<{success:boolean}> =>
    {
        try {
            const response = await api.post("/auth/register", { email, password })

            const user = response?.data?.user
            if (!user) {
                console.error("Registration failed: Invalid response received from server")
                return {success: false}
            }
            setUser(response.data.user)
            console.log("Registration successful")
            return {success: true}
        }
        catch (error) {
            const errorMessage = error.response?.data?.message || error.message || "Unexpected error occurred"
            console.error("Login failed:", errorMessage)

            return {success: false}
        }
    }

    const logout = async (): Promise<{success: boolean}> => {
        try {
            await api.post("/auth/logout")
            setUser(null)

            console.log("Logout successful")
            return {success: true}
        } catch (error) {
            const errorMessage = error.response?.data?.message || error.message || "Unexpected error occurred"
            console.error("Logout failed:", errorMessage)
            return {success: false}
        }
    }

    return <AuthContext.Provider value={{user, loading, login, register, logout}}>
        {children}
    </AuthContext.Provider>
}

export const useAuth = () => useContext(AuthContext)