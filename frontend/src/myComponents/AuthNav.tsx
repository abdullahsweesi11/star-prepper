import { useAuth } from "@/context/AuthContext"
import { Button } from "@/components/ui/button"

export default function AuthNav() {
    const {user} = useAuth()

    return (
        <div>
            <h5 className="size-max">
            {
                user ?
                    user.email
                :
                    <>
                    <Button variant="link">Login</Button> |
                    <Button variant="link">Register</Button>
                    </>
            }
            </h5>
        </div>
    )
}