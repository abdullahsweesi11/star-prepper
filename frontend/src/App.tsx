import { Button } from "@/components/ui/button"
import { Navbar } from "@/myComponents/nav"

function App() {
  return (
    <>
    <Navbar />
    <div className="flex min-h-svh flex-col items-center justify-center">
      <Button onClick={() => alert("You clicked me!")}>Click me</Button>
    </div>
    </>
  )
}

export default App
