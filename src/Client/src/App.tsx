import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [json_data, setdata] = useState({ files: [], curr_path: [] })
  const [change, update_json] = useState(0)
  useEffect(() => {
    get_data()
  }, [change])

  const get_data = async () => {
    const data = await fetch('http://localhost:8000/');
    setdata(await data.json())
  }

  const post_req = (title: string, request_type: string) => {
    if (request_type === 'd') {
      const data = {
        "request": "cd",
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
    else if (request_type === 'f') {
      var link: HTMLElement = document.createElement("a")
      link.href = `http://localhost:8000/download?name=${encodeURIComponent(title)}`
      document.body.appendChild(link)
      link.click()
      link.remove();
    }
  }

  const mkdir = () => {
    var input: HTMLElement = document.createElement("input")
    input.type = 'text'
    document.body.append(input)
  }

  return (
    <>
      <h2>{json_data.curr_path}</h2>
      <button onClick={() => { post_req('', 'd') }}>refresh</button>
      <button onClick={() => { mkdir() }}>new folder</button>
      <button>new file</button>
      <ul className='file_list'>
        {json_data.files.map((data, index) => {
          return (
            <a href='#' onClick={() => { post_req(data.title, String.fromCharCode(data.type)) }}><li key={index}>{String.fromCharCode(data.type)} -- {data.title}</li></a>
          )
        })}
      </ul>
    </>
  )
}

export default App
