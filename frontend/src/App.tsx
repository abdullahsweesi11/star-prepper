import { Link, BrowserRouter, Routes, Route } from "react-router"
import Dashboard from "./pages/Dashboard"
import Home from "./pages/Home"
import Questions from "./pages/Questions"
import AuthNav from "./myComponents/AuthNav"
import ProtectedRoute from "./myComponents/ProtectedRoute"
import { AuthProvider } from "./context/AuthContext"

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <header className="flex justify-between items-center bg-yellow-50 w-full h-[10dvh] px-10">

                    <div className="flex gap-10">
                        <h2 className="size-max">
                            <Link to="/">
                                <strong>STAR-Prepper</strong>
                            </Link>
                        </h2>
                        <h3>
                            <Link to="/questions">
                                Questions
                            </Link>
                        </h3>
                    </div>
                    <AuthNav />
                </header>

                <Routes>
                    <Route index element={<Home />} />
                    <Route path="/dashboard" element={
                        <ProtectedRoute>
                            <Dashboard />
                        </ProtectedRoute>
                    } />
                    <Route path="/questions" element={
                        <ProtectedRoute>
                        <Questions />
                    </ProtectedRoute>
                    } />
                </Routes>
            </BrowserRouter>
        </AuthProvider>
    )
}

export default App
