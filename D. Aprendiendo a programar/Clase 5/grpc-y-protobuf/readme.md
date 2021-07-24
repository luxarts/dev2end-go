# Qué es gRPC

# Qué es Protobuf

---

# Cómo usar gRPC y protobuf

### Instalación
1. Instalar gRPC
2. Instalar protobuf
3. Instalar `protoc-gen-go-grpc`
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
4. Instalar `protoc-gen-go`
```
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

### Generar mensajes y servicios
```
 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
 proto/pelicula.proto
```

