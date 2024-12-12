import { useState } from "react"
import axios from "axios"

function UrlGetter () {
    const [urlData,setUrlData] = useState("")
    const [shortUrl,setShortUrl] = useState("")

    const handleChange = (e) => {
        setUrlData(e.target.value)
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        
        try {
            console.log(urlData)
            console.log(urlData)
            const response = await axios.post("http://localhost:8080/api/data",{urlData})
            setShortUrl(response.data.shortUrl)
            console.log("Response:", response.data);
        }
        catch(error) {
            console.error("Error:", error);
        }
    }

    return (
        <>
            <div>
            <h1>Enter URL:</h1>
            <form action="" onSubmit={handleSubmit}>
                <input type="text" 
                    name="urlData"
                    value={urlData}
                    placeholder="Enter URL"
                    onChange={handleChange}
                />
                <button type="submit">Create URL</button>
            </form>
            {shortUrl && (
                <div>
                    <h2>Shorted URL:</h2>
                    <a href={shortUrl} target="_blank" rel="noopener noreferrer">{shortUrl}</a>
                </div>
            )}
        </div>
        </>
    )
}

export default UrlGetter