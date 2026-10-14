import { Link } from "react-router";
import { useAuth } from "@/context/AuthContext";
import AuthNav from "./AuthNav";

export default function Header() {
    const { user } = useAuth()

    return (
        <header className="flex justify-between items-center bg-yellow-50 w-full h-[10dvh] px-10">

            <div className="flex gap-10">
                <h2 className="size-max">
                    <Link to={user ? "/dashboard" : "/"}>
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
    )
}