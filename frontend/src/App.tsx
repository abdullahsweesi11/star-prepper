import { BrowserRouter, Routes, Route } from "react-router"
import Dashboard from "./pages/Dashboard"
import Home from "./pages/Home"
import Questions from "./pages/Questions"
import Header from "./myComponents/Header"
import ProtectedRoute from "./myComponents/ProtectedRoute"
import { AuthProvider } from "./context/AuthContext"

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Header />
                <Routes>
                    <Route index element={<Home />} />

                    <Route path="/dashboard" element={
                        <ProtectedRoute>
                            <Dashboard />
                        </ProtectedRoute>
                    } />
                    <Route path="/questions" element={
                        <Questions />
                    } />

                    <Route path="*" element={<div>Not found</div>} />
                </Routes>
            </BrowserRouter>
        </AuthProvider>
    )
}

export default App
