import {useEffect, useState} from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

interface Items {
    id: number
    title: string
}

function App() {
  const [items, setItems] = useState<Items[]>()

    useEffect(() => {
        fetch("http://localhost:8080/api/items", {
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then(data => data.json())
            .then(data => setItems(data))
    }, [])

  return (
    <>
        {items?.map((el) => (
            <p key={el.id}>
                {el.title}
            </p>
        ))}
    </>
  )
}

export default App
