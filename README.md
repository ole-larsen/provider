alias swagger='sudo docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger version

swagger generate server -A provider-service -f ./schema/swagger.yml

swagger generate client -A provider-service -f ./schema/swagger.yml
 
swagger serve ./schema/swagger.yml

swagger validate ./schema/swagger.yml

or 

```docker build -t olelarsen/provider --build-arg PORT=5000 .```
```docker push olelarsen/provider```
```docker run --name provider --env PORT=5555 -p 5000:5000 olelarsen/provider```

# For refresh:
https://www.oauth.com/oauth2-servers/making-authenticated-requests/refreshing-an-access-token/

```
POST /oauth/token HTTP/1.1
Host: authorization-server.com
 
grant_type=refresh_token
&refresh_token=xxxxxxxxxxx
&client_id=xxxxxxxxxx
&client_secret=xxxxxxxxxx
```

and receive

```
{
  "access_token": "BWjcyMzY3ZDhiNmJkNTY",
  "refresh_token": "Srq2NjM5NzA2OWJjuE7c",
  "token_type": "Bearer",
  "expires": 3600
}
```