# API Doc

The main purpose of this part is to serve a Swagger-UI webapp to expose MARCEL API

## build

This will generate the server binary ("api_doc")
```shell
make 
```

##Run
Run the 
On your machine, just run theserver
```
./api_doc
```
and launch the web page in your browser: ```http://localhost:3000/```

You first need to run the backend server in order to get the swagger.json file and test the API.

## Credits

* [Swagger UI](https://swagger.io/swagger-ui/) 