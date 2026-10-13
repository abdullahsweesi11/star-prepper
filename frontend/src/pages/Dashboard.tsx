import {
    Card,
    CardAction,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"

export default function Dashboard() {
    return <div className="py-10 px-20">
        <div>
            <header className="flex justify-between mb-2">
                <h2>My Stories</h2>
                <Button variant="ghost" className="px-2 ">
                    Filter
                </Button>
            </header>

            <Input className="mb-5" placeholder="Search"/>

            <h4 className="text-gray-500 italic my-10 place-self-center">No stories available</h4>

            {/* <div className="flex flex-col gap-2">
                <Card className="w-full max-w-full">
                    <CardHeader>
                        <CardTitle>IntervU</CardTitle>
                        <CardDescription>First-year team project</CardDescription>
                        <CardAction>
                            <Button className="bg-yellow-500 text-black">Open</Button>
                        </CardAction>
                    </CardHeader>
                </Card>
                <Card className="w-full max-w-full">
                    <CardHeader>
                        <CardTitle>Eventbrite</CardTitle>
                        <CardDescription>Second-year event-booking agile team project</CardDescription>
                        <CardAction>
                            <Button className="bg-yellow-500 text-black">Open</Button>
                        </CardAction>
                    </CardHeader>
                </Card>
                <Card className="w-full max-w-full">
                    <CardHeader>
                        <CardTitle>GUH 2025 Logistics</CardTitle>
                        <CardDescription>Organised catering logistics for 250-person hackathon</CardDescription>
                        <CardAction>
                            <Button className="bg-yellow-500 text-black">Open</Button>
                        </CardAction>
                    </CardHeader>
                </Card>
            </div> */}
        </div>
    </div>
}