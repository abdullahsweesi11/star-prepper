import { useAuth } from "@/context/AuthContext"
import LoginDropdown from "./LoginDropdown"
import RegisterDropdown from "./RegisterDropdown"
import UserMenu from "./UserMenu"

export default function AuthNav() {
    const {user} = useAuth()

    return (
        <div>
            <h5 className="size-max">
            {
                user ?
                    <UserMenu />
                :
                    <>
                    <LoginDropdown /> |
                    <RegisterDropdown />
                    </>
            }
            </h5>
        </div>
    )
}