swagger generate server -A provider-service -f ./schema/swagger.yml

swagger generate client -A provider-service -f ./schema/swagger.yml
 
swagger serve ./schema/swagger.yml

swagger validate ./schema/swagger.yml

or 

```make build```