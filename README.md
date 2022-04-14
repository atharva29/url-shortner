# URL shortner 

To run Docker image run following command
`docker run -p 8100:8100 atharva29/url-shortner:latest`

--- 

URL shortner : Shortens the URL (by hashing & encoding) and stores the mapping of long URL & short URL in text file.

---
`localhost:8100/create-short-url` 
example: 
- For registering URL (Post request) example.
- Request command: `curl -d '{ "long_url": "www.facebook.com" }' -H 'Content-Type: application/json' localhost:8100/create-short-url`
---
`lcoalhost:8100/:shortUrl`
- pass `shortUrl` in params
example
- URL : http://localhost:8100/hB5TnJXi
---
