swagger generate server -A provider-service -f ./schema/swagger.yml

swagger generate client -A provider-service -f ./schema/swagger.yml
 
swagger serve ./schema/swagger.yml

swagger validate ./schema/swagger.yml

or 

```docker build -t olelarsen/provider --build-arg PORT=5000 .```
```docker push olelarsen/provider```
```docker run --name provider --env PORT=5555 -p 5000:5000 olelarsen/provider```