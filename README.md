# URL shortner 

- To run Docker image run following command
`docker run -p 8500:8100 atharva29/url-shortner:latest`
- If one wants to change port of docker, then use environment variable
`docker run -p 8500:8500 -d -e p=8500  atharva29/url-shortner:latest`

--- 

URL shortner : Shortens the URL (by hashing & encoding) and stores the mapping of long URL & short URL in text file.

---
Create short URL using `localhost:8500/create-short-url`  
- Example: 
- - Request `curl -d '{ "long_url": "www.facebook.com" }' -H 'Content-Type: application/json' localhost:8500/create-short-url`
- - Response `{"message":"short url created successfully","short_url":"http://localhost:8500/QiaaLgq2"}`

---
Redirect to URL using`lcoalhost:8500/:shortUrl`
- pass `shortUrl` in params
example
- - Request `curl localhost:8500/QiaaLgq2`
- - Response `<a href="https://www.facebook.com">Found</a>.`
---
