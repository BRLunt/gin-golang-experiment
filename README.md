# GOLANG-GIN EXPERIMENT

**This is only an experiment looking into golang and the gin framework.**

Provide a service to convert markdown to html.





# The experiment
uses: 
- gin (webserver)
- markdown (markdown to html)
- bluemonday (html sanitizer) 


Test server
```bash
curl -u admin:admin -X POST http://localhost:8080/markdown -H "Content-Type: application/json" -d '{"productId": 123456, "quantity": 100}'
```
