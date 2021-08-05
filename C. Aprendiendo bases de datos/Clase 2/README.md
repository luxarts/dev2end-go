# Instrucciones

1. Instalar [MongoDB server](https://www.mongodb.com/try/download/community)
    > Opcional: Instalar MongoDB Compass, la interfaz gráfica para MongoDB con la cual vas a poder ver los datos de una mejor manera.
2. Descargar dependencias
   ```
   go mod download
   ```
3. Iniciar el servidor
   ```
   go run main.go
   ```

---
# API calls
## Crear usuario
```
curl --location --request POST 'localhost:9090/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Cosme fulanito",
    "password":" 1234",
    "email": "algo@algo.com"
}'
```

## Obtener usuario por ID
```
curl --location --request GET 'localhost:9090/user/:userID'
```