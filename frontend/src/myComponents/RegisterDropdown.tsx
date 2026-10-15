import {
    Popover,
    PopoverContent,
    PopoverTrigger
} from "@/components/ui/popover"
import {
    Card,
    CardDescription,
    CardHeader,
    CardContent,
    CardFooter
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useState } from "react"
import { useAuth } from "@/context/AuthContext"
import { useNavigate } from "react-router"

export default function RegisterDropdown() {

    const { register } = useAuth()
    const navigate = useNavigate()

    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmPassword, setConfirmPassword] = useState("")
    const [error, setError] = useState("")

    async function handleSubmit(event) {
        event.preventDefault()

        if (confirmPassword !== password) {
            setError("Passwords must match")
            return
        }

        const {success} = await register(email, password)
        
        if (!success) {
            setError("Registration failed")
            return
        }

        setError("")
        setEmail("")
        setPassword("")

        navigate("/dashboard")
    }

    return (
        <Popover>
            <PopoverTrigger render={<Button variant="link">Register</Button>} />
            <PopoverContent className="w-max">
                <Card className="border-0 shadow-none">
                    <CardHeader>
                        <CardDescription hidden={error === ""} className="text-red-400">{error}</CardDescription>
                    </CardHeader>
                    <CardContent>
                        <form id="register" onSubmit={handleSubmit}>
                            <div className="flex flex-col gap-6">
                                <div className="flex flex-col gap-2">
                                    <Label htmlFor="email">Email</Label>
                                    <Input
                                        id="email"
                                        name="email"
                                        type="email"
                                        placeholder="name@example.com"
                                        required
                                        value={email}
                                        onChange={event => setEmail(event.target.value)}
                                    />
                                </div>
                                <div className="grid gap-2">
                                    <Label htmlFor="password">Password</Label>
                                    <Input
                                        id="password"
                                        name="password"
                                        type="password"
                                        required
                                        value={password}
                                        onChange={event => setPassword(event.target.value)}
                                    />
                                </div>
                                <div className="grid gap-2">
                                    <Label htmlFor="confirm-password">Confirm Password</Label>
                                    <Input
                                        id="confirm-password"
                                        name="confirm-password"
                                        type="password"
                                        required
                                        value={confirmPassword}
                                        onChange={event => setConfirmPassword(event.target.value)}
                                    />
                                </div>
                            </div>
                        </form>
                    </CardContent>
                    <CardFooter>
                        <Button type="submit" form="register" className="w-full">
                            Login
                        </Button>
                    </CardFooter>
                </Card>
            </PopoverContent>
        </Popover>
    )
}