import { useState } from "react"

import { initWasm } from "../../wasm"

export const HomePage = () => {
    const [count, setCount] = useState(0)
    
    const increment = async () => {
        const { add } = initWasm.getMethods()
        setCount(add(count, 1))
    }

    return (
        <>
            <h1>{count}</h1>
            <button onClick={increment}>
                increment
            </button>
        </>
    )
}