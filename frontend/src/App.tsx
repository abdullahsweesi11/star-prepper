import { Link, BrowserRouter, Routes, Route } from "react-router"
import Dashboard from "./pages/Dashboard"
import Questions from "./pages/Questions"
import AuthNav from "./myComponents/AuthNav"

function App() {
    return (
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
                <Route path="/" element={<Dashboard />} />
                <Route path="/questions" element={<Questions />} />
            </Routes>
        </BrowserRouter>
    )
}

export default App
