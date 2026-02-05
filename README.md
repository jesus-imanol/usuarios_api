# API de Inventario de Productos y Usuarios

API RESTful para gestionar un inventario de productos y usuarios siguiendo los principios de Clean Architecture.

## Arquitectura

El proyecto sigue Clean Architecture con las siguientes capas:

- **Domain**: Contiene las entidades de negocio y las interfaces de repositorios
  - Productos y Usuarios

- **Application**: Contiene los casos de uso de la aplicación
  - Operaciones CRUD para productos y usuarios
  - Autenticación y autorización

- **Infrastructure**: Contiene las implementaciones concretas
  - Adaptadores MySQL
  - Controladores HTTP
  - Routers y dependencias

- **Core**: Utilidades compartidas
  - Configuración de base de datos
  - Utilidades de hash y upload de archivos

---

## 📋 Endpoints de la API

### Base URL
```
http://localhost:8080
```

---

## 🛍️ PRODUCTOS

### 1. Crear Producto
**Endpoint:** `POST /v1/products`

**Request Body:**
```json
{
  "name": "Laptop Dell XPS 15",
  "price": 1299.99,
  "quantity": 50
}
```

**Response (201 Created):**
```json
{
  "data": {
    "type": "products",
    "id": 1,
    "attributes": {
      "name": "Laptop Dell XPS 15",
      "price": 1299.99,
      "quantity": 50
    }
  }
}
```

---

### 2. Obtener Todos los Productos
**Endpoint:** `GET /v1/products`

**Response (200 OK):**
```json
{
  "data": [
    {
      "type": "products",
      "id": 1,
      "attributes": {
        "name": "Laptop Dell XPS 15",
        "price": 1299.99,
        "quantity": 50
      }
    },
    {
      "type": "products",
      "id": 2,
      "attributes": {
        "name": "Mouse Logitech",
        "price": 29.99,
        "quantity": 100
      }
    }
  ]
}
```

**Response cuando no hay productos (200 OK):**
```json
{
  "data": 0,
  "message": "No se encontraron productos",
  "type": "products"
}
```

---

### 3. Obtener Producto por ID
**Endpoint:** `GET /v1/products/:id`

**Ejemplo:** `GET /v1/products/1`

**Response (200 OK):**
```json
{
  "data": {
    "type": "products",
    "id": 1,
    "attributes": {
      "name": "Laptop Dell XPS 15",
      "price": 1299.99,
      "quantity": 50
    }
  }
}
```

**Response si no existe (500 Internal Server Error):**
```json
{
  "error": "error message"
}
```

---

### 4. Actualizar Producto
**Endpoint:** `PUT /v1/products/:id`

**Ejemplo:** `PUT /v1/products/1`

**Request Body:**
```json
{
  "name": "Laptop Dell XPS 15 Updated",
  "price": 1199.99,
  "quantity": 45
}
```

**Response (200 OK):**
```json
{
  "data": {
    "type": "products",
    "id": 1,
    "attributes": {
      "name": "Laptop Dell XPS 15 Updated",
      "price": 1199.99,
      "quantity": 45
    }
  }
}
```

---

### 5. Eliminar Producto
**Endpoint:** `DELETE /v1/products/:id`

**Ejemplo:** `DELETE /v1/products/1`

**Response (200 OK):**
```json
{
  "data": {
    "type": "products",
    "id": "1",
    "message": "Producto eliminado con éxito"
  }
}
```

**Response si no existe (404 Not Found):**
```json
{
  "detail": "producto con ID 1 no encontrado",
  "type": "products"
}
```

---

## 👥 USUARIOS

### 1. Registrar Usuario
**Endpoint:** `POST /v1/users`

**Request Body:**
```json
{
  "full_name": "Juan Pérez",
  "email": "juan.perez@example.com",
  "password_hash": "miPassword123",
  "gender": "M",
  "match_preference": "F",
  "city": "Ciudad de México",
  "state": "CDMX",
  "interests": "deportes, música, cine",
  "status_message": "¡Hola! Busco conocer gente nueva",
  "profile_picture": ""
}
```

**Response (201 Created):**
```json
{
  "data": {
    "type": "users",
    "id": 1,
    "attributes": {
      "full_name": "Juan Pérez",
      "email": "juan.perez@example.com",
      "profile_picture": "",
      "gender": "M",
      "city": "Ciudad de México",
      "state": "CDMX",
      "status_message": "¡Hola! Busco conocer gente nueva",
      "match_preference": "F",
      "interests": "deportes, música, cine"
    }
  }
}
```

**Headers de respuesta:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

### 2. Login de Usuario
**Endpoint:** `POST /v1/users/login`

**Request Body:**
```json
{
  "email": "juan.perez@example.com",
  "password_hash": "miPassword123"
}
```

**Response (200 OK):**
```json
{
  "data": {
    "type": "users",
    "id": 1,
    "attributes": {
      "full_name": "Juan Pérez",
      "profile_picture": "",
      "gender": "M",
      "city": "CDMX",
      "status_message": "¡Hola! Busco conocer gente nueva",
      "match_preference": "F",
      "interests": "deportes, música, cine",
      "email": "juan.perez@example.com",
      "state": "CDMX"
    }
  }
}
```

**Headers de respuesta:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Response si las credenciales son inválidas (401 Unauthorized):**
```json
{
  "error": "credenciales invalidas"
}
```

---

### 3. Obtener Todos los Usuarios
**Endpoint:** `GET /v1/users`

**Response (200 OK):**
```json
{
  "data": [
    {
      "type": "users",
      "id": 1,
      "attributes": {
        "full_name": "Juan Pérez",
        "profile_picture": "",
        "gender": "M",
        "city": "CDMX",
        "status_message": "¡Hola! Busco conocer gente nueva",
        "match_preference": "F",
        "interests": "deportes, música, cine",
        "email": "juan.perez@example.com",
        "state": "CDMX"
      }
    }
  ]
}
```

**Response cuando no hay usuarios (200 OK):**
```json
{
  "data": 0,
  "message": "No se encontraron users",
  "type": "users"
}
```

---

### 4. Obtener Usuario por ID
**Endpoint:** `GET /v1/users/:id`

**Ejemplo:** `GET /v1/users/1`

**Response (200 OK):**
```json
{
  "data": {
    "type": "users",
    "id": 1,
    "attributes": {
      "full_name": "Juan Pérez",
      "email": "juan.perez@example.com",
      "profile_picture": "",
      "gender": "M",
      "city": "Ciudad de México",
      "state": "CDMX",
      "status_message": "¡Hola! Busco conocer gente nueva",
      "match_preference": "F",
      "interests": "deportes, música, cine"
    }
  }
}
```

---

### 5. Actualizar Usuario
**Endpoint:** `PUT /v1/users/:id`

**Ejemplo:** `PUT /v1/users/1`

**Request Body:**
```json
{
  "full_name": "Juan Pérez Actualizado",
  "email": "juan.perez@example.com",
  "password_hash": "newPassword123",
  "gender": "M",
  "match_preference": "F",
  "city": "Guadalajara",
  "state": "Jalisco",
  "interests": "deportes, música, cine, lectura",
  "status_message": "Actualizado mi perfil"
}
```

**Response (200 OK):**
```json
{
  "data": {
    "type": "users",
    "id": 1,
    "attributes": {
      "full_name": "Juan Pérez Actualizado",
      "email": "juan.perez@example.com",
      "gender": "M",
      "city": "Guadalajara",
      "state": "Jalisco",
      "status_message": "Actualizado mi perfil",
      "match_preference": "F",
      "interests": "deportes, música, cine, lectura"
    }
  }
}
```

---

### 6. Eliminar Usuario (Soft Delete)
**Endpoint:** `DELETE /v1/users/:id`

**Ejemplo:** `DELETE /v1/users/1`

**Response (200 OK):**
```json
{
  "data": {
    "type": "users",
    "id": "1",
    "message": "User eliminado con éxito"
  }
}
```

**Response si no existe (404 Not Found):**
```json
{
  "detail": "usuario con ID 1 no encontrado",
  "type": "users"
}
```

---

### 7. Subir Foto de Perfil
**Endpoint:** `PUT /v1/users/upload-picture/:id`

**Ejemplo:** `PUT /v1/users/upload-picture/1`

**Request:**
- **Content-Type:** `multipart/form-data`
- **Campo:** `profile_picture` (archivo de imagen)

**Response (200 OK):**
```json
{
  "data": {
    "type": "users",
    "id": 1,
    "attributes": {
      "profile_picture": "https://spontaneity-2025.s3.amazonaws.com/filename.jpg"
    }
  }
}
```

**Response en caso de error (400/500):**
```json
{
  "error": "Failed to get profile picture"
}
```

---

## 🔧 Configuración

### Variables de Entorno

Crear un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
# Base de datos MySQL
DB_HOST=localhost
DB_USER=tu_usuario
DB_PASS=tu_contraseña
DB_SCHEMA=tu_base_de_datos

# JWT Secret para autenticación
JWT_SECRET=tu_secreto_jwt_aqui

# Secret Key para hash de contraseñas
SECRET_KEY=tu_secreto_para_passwords

# AWS S3 (para subida de imágenes)
aws_access_key_id=tu_access_key
aws_secret_access_key=tu_secret_key
aws_session_token=tu_session_token
```

---

## 📊 Base de Datos

### Tabla de Productos

```sql
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Tabla de Usuarios

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    gender VARCHAR(10),
    match_preference VARCHAR(10),
    city VARCHAR(100),
    state VARCHAR(100),
    interests TEXT,
    status_message TEXT,
    profile_picture TEXT,
    deleted BOOLEAN DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

---

## 🚀 Instalación y Ejecución

### 1. Clonar el repositorio o navegar al directorio

```bash
cd inventario_productos
```

### 2. Instalar dependencias

```bash
go mod download
```

### 3. Configurar el archivo `.env`

Copiar el archivo de ejemplo y configurar las variables necesarias.

### 4. Crear las tablas en MySQL

Ejecutar los scripts SQL proporcionados anteriormente.

### 5. Ejecutar la aplicación

```bash
go run main.go
```

La API estará disponible en: **http://localhost:8080**

---

## 🧪 Ejemplos de Uso con cURL

### Productos

**Crear un producto:**
```bash
curl -X POST http://localhost:8080/v1/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop HP",
    "price": 899.99,
    "quantity": 25
  }'
```

**Obtener todos los productos:**
```bash
curl http://localhost:8080/v1/products
```

**Obtener un producto por ID:**
```bash
curl http://localhost:8080/v1/products/1
```

**Actualizar un producto:**
```bash
curl -X PUT http://localhost:8080/v1/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop HP Actualizado",
    "price": 849.99,
    "quantity": 20
  }'
```

**Eliminar un producto:**
```bash
curl -X DELETE http://localhost:8080/v1/products/1
```

### Usuarios

**Registrar un usuario:**
```bash
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "María García",
    "email": "maria@example.com",
    "password_hash": "password123",
    "gender": "F",
    "match_preference": "M",
    "city": "Monterrey",
    "state": "Nuevo León",
    "interests": "viajes, fotografía",
    "status_message": "Explorando el mundo",
    "profile_picture": ""
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "maria@example.com",
    "password_hash": "password123"
  }'
```

**Subir foto de perfil:**
```bash
curl -X PUT http://localhost:8080/v1/users/upload-picture/1 \
  -F "profile_picture=@/ruta/a/tu/imagen.jpg"
```

---

## 📦 Tecnologías y Dependencias

- **Go** 1.24.0+
- **Gin Framework** - Framework web
- **MySQL Driver** - Conexión a base de datos
- **godotenv** - Manejo de variables de entorno
- **bcrypt** - Hash de contraseñas
- **JWT-Go** - Autenticación con tokens
- **AWS SDK Go** - Integración con Amazon S3
- **CORS** - Configuración de Cross-Origin Resource Sharing

---

## 📁 Estructura del Proyecto

```
inventario_productos/
├── main.go
├── go.mod
├── go.sum
├── .env
├── README.md
└── src/
    ├── core/
    │   └── db_mysql.go
    ├── products/
    │   ├── domain/
    │   │   ├── entities/
    │   │   │   └── Product.go
    │   │   └── repositories/
    │   │       └── product_repository.go
    │   ├── application/
    │   │   ├── CreateProduct_usecase.go
    │   │   ├── UpdateProduct_usecase.go
    │   │   ├── GetProducts_usecase.go
    │   │   ├── GetProductById_usecase.go
    │   │   └── DeleteProduct_usecase.go
    │   └── infraestructure/
    │       ├── adapters/
    │       │   └── MySQL.go
    │       ├── controllers/
    │       ├── routers/
    │       └── dependencies_product/
    └── users/
        ├── domain/
        │   ├── entities/
        │   │   └── User.go
        │   └── repositories/
        │       └── user_repository.go
        ├── application/
        │   ├── RegisterUser_usecase.go
        │   ├── LoginUser_usecase.go
        │   ├── ListUser_usecase.go
        │   ├── GetUserById_usecase.go
        │   ├── UpdateUser_usecase.go
        │   ├── DeleteUser_usecase.go
        │   └── UploadPictureUser_usecase.go
        └── infraestructure/
            ├── adapters/
            │   └── MySQL.go
            ├── utils/
            │   ├── HashPassword.go
            │   └── UploadFiles.go
            ├── controllers/
            ├── routers/
            └── dependencies_user/
```

---

## 🔒 Seguridad

- Las contraseñas se almacenan hasheadas usando **bcrypt**
- Los tokens JWT tienen una expiración de 72 horas
- Se implementa CORS para control de acceso
- Las imágenes se suben a AWS S3 con configuración segura

---

## 📝 Notas Importantes

1. **Formato de Respuesta:** Todas las respuestas siguen el estándar JSON API con estructura `data`, `type`, `id`, y `attributes`
2. **Eliminación Suave:** Los usuarios se eliminan de forma lógica (soft delete) con el campo `deleted`
3. **Validación:** Se recomienda agregar validaciones adicionales en el frontend
4. **HTTPS:** En producción, usar HTTPS para todas las comunicaciones
5. **AWS S3:** Configurar el bucket con las políticas de acceso adecuadas

---

## 🐛 Manejo de Errores

La API retorna los siguientes códigos de estado HTTP:

- **200 OK** - Operación exitosa
- **201 Created** - Recurso creado exitosamente
- **400 Bad Request** - Datos de entrada inválidos
- **401 Unauthorized** - Credenciales inválidas
- **404 Not Found** - Recurso no encontrado
- **500 Internal Server Error** - Error del servidor

---

## 👨‍💻 Autor

Proyecto desarrollado siguiendo los principios de Clean Architecture y mejores prácticas de Go.

---

## 📄 Licencia

Este proyecto es de código abierto y está disponible bajo la licencia que determines.
