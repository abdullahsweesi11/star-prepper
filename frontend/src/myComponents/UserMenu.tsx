import {
    DropdownMenu,
    DropdownMenuTrigger,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuGroup,
    DropdownMenuLabel
} from "@/components/ui/dropdown-menu"
import { Avatar, AvatarFallback } from "@/components/ui/avatar"
import { useAuth } from "@/context/AuthContext"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router"

export default function UserMenu() {
    const {user, logout} = useAuth()
    const navigate = useNavigate()

    async function handleLogout() {
        await logout()                  // TODO: handle logout fail
        navigate("/")
    }

    return (
        <DropdownMenu>
            <DropdownMenuTrigger render={
                <Button variant="ghost" className="relative h-10 w-10 p-0 rounded-full">
                    <Avatar className="h-10 w-10">
                        <AvatarFallback className="bg-gray-500 text-white">?</AvatarFallback>
                    </Avatar>
                </Button>
            } />
            <DropdownMenuContent className="w-max">
                <DropdownMenuGroup>
                    <DropdownMenuLabel>{user.email}</DropdownMenuLabel>
                </DropdownMenuGroup>
                <DropdownMenuGroup>
                    <DropdownMenuItem className="text-red-500 focus:bg-red-100 focus:text-red-600 cursor-pointer" onClick={handleLogout}>
                        Logout
                    </DropdownMenuItem>
                </DropdownMenuGroup>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}