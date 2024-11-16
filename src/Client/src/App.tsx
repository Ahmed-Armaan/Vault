import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [json_data, setdata] = useState([])
  const [change, update_json] = useState(0)
  useEffect(() => {
    get_data()
  }, [change])

  const get_data = async () => {
    const data = await fetch('http://localhost:8000/');
    setdata(await data.json())
  }

  const post_req = (title: string, request: string) => {
    var req: string = (request === 'f') ? 'open' : 'cd'
    const data = {
      "request": req,
      "path": title
    }

    fetch('http://localhost:8000/req', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', },
      body: JSON.stringify(data),
    })
      .then(Response => Response.json)
      .then(() => {
        update_json((change + 1) % 2)
      })
      .catch(Error => {
        console.log(`Error: ${Error}`)
      })
  }

  return (
    <>
      <ul className='file_list'>
        {json_data.map((data, index) => {
          return (
            <a href='#' onClick={() => { post_req(data.title, String.fromCharCode(data.type)) }}><li key={index}>{String.fromCharCode(data.type)} -- {data.title}</li></a>
          )
        })}
      </ul>
    </>
  )
}

export default App
