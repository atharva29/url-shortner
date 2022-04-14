# URL shortner 

To run Docker image run following command
`docker run -p 8100:8100 atharva29/url-shortner:latest`

--- 

URL shortner : Shortens the URL (by hashing & encoding) and stores the mapping of long URL & short URL in text file.

---
Create short URL using `localhost:8100/create-short-url`  
- Example: 
- - Request `curl -d '{ "long_url": "www.facebook.com" }' -H 'Content-Type: application/json' localhost:8100/create-short-url`
- - Response `{"message":"short url created successfully","short_url":"http://localhost:8100/QiaaLgq2"}`

---
Redirect to URL using`lcoalhost:8100/:shortUrl`
- pass `shortUrl` in params
example
- URL : http://localhost:8100/hB5TnJXi
---
