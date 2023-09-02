# Configuracion inicial

Este proyecto sigue una estructura organizada de carpetas para mantener el código fuente limpio y fácil de navegar. Aquí se describe el propósito de cada carpeta:

- **/config**: Contiene archivos de configuración, como configuraciones de bases de datos o configuraciones de la aplicación.

- **/controllers**: En esta carpeta se encuentran los controladores de la aplicación, que manejan la lógica de negocio y las solicitudes del cliente.

- **/middleware**: Almacena los archivos de middleware, que son funciones que se ejecutan antes o después de las rutas para realizar acciones como la autenticación o la validación de datos.

- **/models**: Aquí se definen los modelos de datos de tu aplicación, que representan las estructuras de datos.

- **/routes**: Contiene las definiciones de rutas de la aplicación, que mapean las URL a controladores específicos.

# Módulos Utilizados en la Aplicación

Este proyecto utiliza varios módulos de terceros para llevar a cabo diversas funcionalidades. 
Breve descripcion de los modulos y link de la documentación: 

### 1. [Fiber](https://docs.gofiber.io)

   - **Versión:** v2.49.1
   - **Descripción:** Fiber es un framework web de Go diseñado para ser rápido y eficiente. Se utiliza para manejar las rutas y solicitudes HTTP en la aplicación.

### 2. [JWT (JSON Web Tokens)](https://github.com/golang-jwt/jwt)

   - **Versión:** v3.2.2+incompatible
   - **Descripción:** La biblioteca JWT se utiliza para generar, firmar y verificar tokens JWT. Estos tokens se utilizan comúnmente para la autenticación y la autorización de usuarios en aplicaciones web y API.

### 3. [Godotenv](https://github.com/joho/godotenv)

   - **Versión:** v1.5.1
   - **Descripción:** Godotenv se utiliza para cargar variables de entorno desde un archivo [`.env`](#variables-de-entorno) en el entorno de desarrollo. Esto permite gestionar configuraciones sensibles de manera segura.

### 4. [crypto](https://pkg.go.dev/golang.org/x/crypto)

   - **Versión:** v0.12.0
   - **Descripción:** El paquete crypto proporciona funciones criptográficas básicas. Es comúnmente utilizado para operaciones de seguridad, como la generación de hash o la encriptación de contraseñas.

### 5. [GORM](https://gorm.io/docs/) y [GORM Postgres Driver](https://pkg.go.dev/gorm.io/driver/postgres)

   - **Versión de GORM:** v1.25.4
   - **Versión del controlador Postgres:** v1.5.2
   - **Descripción:** GORM es un ORM (Object-Relational Mapping) para Go que se utiliza para interactuar con bases de datos relacionales. El controlador de Postgres es específico para trabajar con bases de datos PostgreSQL.


 **run go mod tidy** Para asegurarse que las dependencias esten al dia.

# Variables de entorno

Ejemplo archivo .env utilizado:

- */.env*
```
PORT=3000
DB_URL="host=localhost user=postgres password=<Contraseña> dbname=<NombreDeLaDB> port=5432 sslmode=disable"
JWT_SECRET=proyectointegrador
```

# Aplicacion de ejemplo

Aplicacion simple de login y registro de usuarios, donde cada usuario puede añadir tareas a completar.