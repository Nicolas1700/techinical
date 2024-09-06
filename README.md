# Technical

Prueba de Talento

# Configuración de Variables de Entorno

Este documento describe las variables de entorno necesarias para configurar y ejecutar el servicio `techinical-api`. Asegúrate de configurar correctamente cada variable antes de iniciar la aplicación.

## Variables de entorno

### Conexión a la Base de Datos

| Variable      | Descripción                                            |
| ------------- | ------------------------------------------------------ |
| `DB_USER`     | Nombre de usuario para conectarse a la base de datos   |
| `DB_PASS`     | Contraseña del usuario de la base de datos             |
| `DB_SCHEMA`   | Esquema de la base de datos a utilizar                 |
| `DB_HOST`     | Dirección del servidor de la base de datos             |
| `DB_PORT`     | Puerto en el que la base de datos escucha conexiones   |
| `DB_DATABASE` | Nombre de la base de datos utilizada por la aplicación |

Estas variables configuran los detalles de la conexión a la base de datos PostgreSQL que utiliza la aplicación para almacenar y recuperar datos.

### Configuración del Servidor

| Variable       | Descripción                                                                               |
| -------------- | ------------------------------------------------------------------------------------------|
| `NAME_SERVICE` | Nombre del servicio API, utilizado para identificar el servicio.                          |
| `PORT_SERVICE` | Puerto en el que se ejecuta el servicio.                                                  |

Estas variables controlan el nombre del servicio y el puerto en el que la API estará disponible.

### Llave de acceso para ChatGPT
| Variable       | Descripción                                                                                                                                              |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------|
| `KEY_OPEN_IA`  | Clave API para acceder a los servicios de OpenAI. Debes ingresar tu propia clave para poder utilizar la funcionalidad de generación de nombres y números |

# Ejecutar la Aplicación

## Requisitos Previos

1. **Base de Datos:** Asegúrate de haber ejecutado el script SQL ubicado en `db/scripts.sql` sobre una base de datos PostgreSQL antes de iniciar la aplicación.
   
2. **Clave de OpenAI:** Es necesario contar con una clave válida de OpenAI, almacenada en la variable de entorno `KEY_OPEN_IA`. En los métodos POST y PATCH, si faltan ciertos datos no primarios, se utilizará la API de ChatGPT para generarlos automáticamente. 
- Consulta [cómo generar una API Key de OpenAI](https://platform.openai.com/docs/quickstart/create-and-export-an-api-key).
- Puedes consultar con el author para una key de prueba.

## Formas de Ejecución

1. **Ejecutar sin compilar:**
   Utiliza el siguiente comando para correr la aplicación directamente:
   ```bash
   go run .\main.exe

2. **Compilar y ejecutar:** Puedes compilar la aplicación utilizando el siguiente comando:
    ```bash
    go build -o technical.exe .\main.go

## Configuración de la API

La API depende de dos variables de entorno importantes:

- `NAME_SERVICE`: Nombre del servicio, por defecto `techinical-api`.
- `PORT_SERVICE`: Puerto en el que corre el servicio, por defecto `3000`.
Ejemplo:

Si `NAME_SERVICE=techinical-api` y `PORT_SERVICE=3000`, el path base de la API será:
- http://127.0.0.1:3000/techinical-api

Los subpaths disponibles son:

- `/users`
- `/videos`
- `/challenges`

Para acceder a un recurso específico, simplemente añade el subpath al base path, por ejemplo:

- http://127.0.0.1:3000/techinical-api/users

## Servicios Disponibles
Cada uno de los subpaths mencionados soporta operaciones CRUD (GET, POST, PATCH, DELETE).

# Definición de Entidades

## Users

| Atributo    | Tipo            | Descripción                                 |
|-------------|-----------------|---------------------------------------------|
| id_user     | varchar(50) PK  | Identificador del usuario (clave primaria)  |
| name_user   | varchar(50)     | Nombre del usuario                          |
| cell_phone  | int             | Número de celular                           |

## Videos

| Atributo    | Tipo           | Descripción                                |
|-------------|----------------|--------------------------------------------|
| id_video    | varchar(50) PK | Identificador del video (clave primaria)   |
| id_user     | varchar(50) FK | Relación con la entidad Users              |
| name_video  | varchar(50)    | Nombre del video                           |
| url_video   | varchar(255)   | URL del video                              |

## Challenges

| Atributo            | Tipo           | Descripción                                 |
|---------------------|----------------|---------------------------------------------|
| id_challenge        | varchar(50) PK | Identificador del desafío (clave primaria)  |
| id_video            | varchar(50) FK | Relación con la entidad Videos              |
| name_challenge      | varchar(50)    | Nombre del desafío                          |
| number_participants | int            | Número de participantes                     |

# Verbos HTTP

## GET
El método GET permite paginar los resultados utilizando los siguientes query parameters:

- `page`: Número de página a consultar.
- `limit`: Cantidad de registros a devolver por página.

Ejemplo:
- http://127.0.0.1:3000/techinical-api/users?page=1&limit=10

## POST
Al ejecutar una solicitud POST, si no se envía la llave primaria, el sistema genera automáticamente un UUID. Si faltan datos, ChatGPT completará los campos si se cumplen los requisitos previos mencionados.
- Ejemplo de Cuerpo de Solicitud:

### JSON POST
```
{
  "name_User": "Juan Pérez",
  "cell_Phone": 3220220202
}
```

## PATCH
En una solicitud PATCH, se requiere proporcionar la llave primaria. Si faltan otros campos, se generarán automáticamente, a menos que se envíen explícitamente.
- Ejemplo de Cuerpo de Solicitud:

### JSON PATCH
```
{
  "id_User": "52451097-03f0-429f-86ce-6977f2f7f1a6",
  "name_User": "Nuevo Nombre",
  "cell_Phone": 3002002020
}
```

## DELETE
Para eliminar un registro, solo se necesita enviar el ID de la entidad.
- Ejemplo de Cuerpo de Solicitud:

### JSON DELETE
```
{
  "id_User": "52451097-03f0-429f-86ce-6977f2f7f1a6"
}
```

# Comentario author
Muchas gracias por tomarte el tiempo de revisar la documentación. Con esta experiencia, he aprendido varias cosas, especialmente sobre cómo debería enfocarse la construcción de un proyecto. He comprendido la importancia de consumir recursos de manera eficiente sin crear dependencias directas entre ellos. Por ahora, este proyecto representa solo un punto de partida y, aunque aún tiene muchos aspectos por mejorar, es el comienzo de algo más grande.

¡Gracias nuevamente por tu apoyo y sugerencias!