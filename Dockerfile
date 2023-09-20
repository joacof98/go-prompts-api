# Usar la imagen base oficial de Golang
FROM golang:1.20.1 AS build

# Establecer el directorio de trabajo en el GOPATH
WORKDIR /go/src/app

# Copiar los archivos necesarios para descargar las dependencias
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el resto de la aplicación
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o epicprompts cmd/epicprompts/main.go

# Crear una imagen mínima basada en Scratch para ejecutar la aplicación
FROM scratch

# Copiar el ejecutable compilado desde la etapa de construcción anterior
COPY --from=build /go/src/app/epicprompts /

# Exponer el puerto en el que escucha tu aplicación (si es necesario)
EXPOSE 8080

# Comando para ejecutar la aplicación cuando se inicie el contenedor
CMD ["/epicprompts"]
