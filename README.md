# technical

Prueba tecnica Talent pitch

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
