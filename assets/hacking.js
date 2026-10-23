const queryString = window.location.search;
const urlParams = new URLSearchParams(queryString);
const hash = urlParams.get("hash")

hackPassword(hash)

async function hackPassword(hash) {
  url = 'http://localhost:8080/guess'
  data = { "hash": hash }
   console.log(JSON.stringify(data))
try {
  const response = await fetch('http://localhost:8080/guess', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data)
  });

  if (response.ok) {
    const result = await response.json()
    console.log("Password is:", result.password)
    password = encodeURIComponent(result.password)
    window.location.replace(`/success?password=${password}`)
  } else {
    const errorMsg = "faild to hack password"
    fmtError = encodeURIComponent(errorMsg)
    window.location.replace(`/error?msg=${fmtError}`)
    console.log("failed to crack hash")
  }
} catch (error) {
  console.error("Failed to send request:", error)
}
}
