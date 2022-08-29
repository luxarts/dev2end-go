# Contenidos
1. [Requisitos](#requisitos)
2. [Compilar y deplegar](#compilar-y-desplegar)
3. [Ejecutar la función](#ejecutar-funcin)
4. [Eliminar la función](#ejecutar-funcin)
5. [Actualizar la función](#actualizar-funcin)

---

## Requisitos
1. [Docker](https://www.docker.com/products/docker-desktop/)
2. [Localstack](https://github.com/localstack/localstack)
   ```
   pip install localstack
   ```
3. [AWS CLI](https://aws.amazon.com/es/cli/)
   ```
   pip install awscli
   ```
4. [AWS CLI Local](https://github.com/localstack/awscli-local)
   ```
   pip install awscli-local
   ```
5. Iniciar servicios
   ```
   localstack start -d
   ```
   
## Compilar y desplegar
1. Crear build
   ```
   GOOS=linux GOARCH=amd64 go build -o main cmd/main.go
   ```
2. Comprimir en un archivo ZIP
   ```
   zip main.zip main
   ```
3. Eliminar el build
   ```
   rm main
   ```
4. Crear función Lambda en Localstack
   ```
   awslocal lambda create-function \
      --function-name clase1 \
      --runtime go1.x \
      --role arn:aws:iam::123456::role/irrelevant \
      --handler main \
      --zip-file fileb://main.zip
   ```
5. Eliminar el zip
   ```
   rm main.zip
   ```

## Ejecutar función
```
awslocal lambda invoke \
--function-name clase1 \
--payload '{"nombre":"Clark Kent"}' \
/dev/stdout
```

## Eliminar function
```
awslocal lambda delete-function \
   --function-name clase1
```

## Actualizar función
```
awslocal lambda update-function-code \
   --function-name clase1 \
   --zip-file fileb://main.zip
```

---

###  Snippet para buildear y actualizar
```
GOOS=linux GOARCH=amd64 go build -o main cmd/main.go && 
zip main.zip main &&
awslocal lambda update-function-code \
   --function-name clase1 \
   --zip-file fileb://main.zip &&
rm main &&
rm main.zip
```