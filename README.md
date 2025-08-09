# Dragon Ball API

Una API REST desarrollada en Go que proporciona información sobre personajes de Dragon Ball. Esta API actúa como un proxy y transformador de datos de la [Dragon Ball API pública](https://dragonball-api.com/).

## 🚀 Características

- **Consulta de personajes**: Obtén información detallada de personajes de Dragon Ball
- **Paginación**: Soporte para paginación con parámetros `limit` y `page`
- **Gestión de favoritos**: Funcionalidad para manejar personajes favoritos
- **Arquitectura limpia**: Separación clara entre handlers, services, models y mappers
- **API RESTful**: Endpoints siguiendo las mejores prácticas REST

## 📁 Estructura del Proyecto

```
drogonball-api/
├── handlers/           # Controladores HTTP
│   ├── get_characters.go
│   └── add_favorite.go
├── models/            # Estructuras de datos
│   └── character.go
├── routers/           # Configuración de rutas
│   ├── routers.go
│   └── v1.go
├── services/          # Lógica de negocio
│   └── get_characters.go
├── mappers/           # Transformación de datos
│   └── response.go
├── storage/           # Capa de persistencia (vacía por ahora)
├── main.go           # Punto de entrada de la aplicación
├── go.mod            # Dependencias del módulo
└── README.md         # Documentación del proyecto
```

## 🛠️ Tecnologías Utilizadas

- **Go 1.24.5**: Lenguaje de programación principal
- **Gin**: Framework web HTTP para Go
- **HTTP Client**: Para consumir APIs externas

## 📋 Prerequisitos

- Go 1.21 o superior
- Conexión a internet (para consumir la API externa)

## ⚡ Instalación y Ejecución

1. **Clonar el repositorio**
   ```bash
   git clone <repository-url>
   cd drogonball-api
   ```

2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

4. **La API estará disponible en**: `http://localhost:8080`

## 📚 Endpoints

### Obtener Personajes

```http
GET /api/v1/characters?limit={limit}&page={page}
```

**Parámetros de consulta:**
- `limit` (int): Número de personajes por página
- `page` (int): Número de página a consultar

**Ejemplo de request:**
```bash
curl "http://localhost:8080/api/v1/characters?limit=10&page=1"
```

**Ejemplo de response:**
```json
{
  "items": [
    {
      "id": 1,
      "name": "Goku",
      "ki": "60.000.000",
      "maxKi": "90 Septillion",
      "race": "Saiyan",
      "gender": "Male",
      "description": "El protagonista de la serie...",
      "image": "https://dragonball-api.com/characters/goku_normal.webp",
      "affiliation": "Z Fighter",
      "deletedAt": null
    }
  ],
  "links": {
    "first": "https://dragonball-api.com/api/characters?page=1&limit=10",
    "previous": "",
    "next": "https://dragonball-api.com/api/characters?page=2&limit=10",
    "last": "https://dragonball-api.com/api/characters?page=6&limit=10"
  }
}
```

### Gestionar Favoritos

```http
POST /api/v1/favorites
```

**Ejemplo de request:**
```bash
curl -X POST "http://localhost:8080/api/v1/favorites"
```

**Ejemplo de response:**
```json
{
  "message": "characters"
}
```

## 🏗️ Arquitectura

La aplicación sigue una arquitectura en capas:

1. **Handlers**: Manejan las requests HTTP y responses
2. **Services**: Contienen la lógica de negocio
3. **Models**: Definen las estructuras de datos
4. **Mappers**: Transforman datos entre diferentes formatos
5. **Routers**: Configuran las rutas y middleware

### Flujo de datos:

```
HTTP Request → Router → Handler → Service → External API
                                     ↓
HTTP Response ← Mapper ← Handler ← Service Response
```

## 🔧 Desarrollo

### Agregar nuevos endpoints

1. Crear el handler en `handlers/`
2. Agregar la lógica de negocio en `services/`
3. Definir modelos en `models/` si es necesario
4. Registrar la ruta en `routers/v1.go`

### Estructura de un handler típico:

```go
func MyHandler(c *gin.Context) {
    // Validar parámetros
    // Llamar al service
    // Manejar errores
    // Retornar response
}
```

## 🐛 Manejo de Errores

La API maneja varios tipos de errores:

- **400 Bad Request**: Parámetros inválidos
- **500 Internal Server Error**: Errores del servidor o API externa
- **Timeout**: La API externa tiene un timeout de 15 segundos

## 🔄 API Externa

Esta API consume datos de: `https://dragonball-api.com/api/characters`

## 📝 TODO

- [ ] Implementar autenticación
- [ ] Agregar sistema de base de datos para favoritos
- [ ] Implementar caché
- [ ] Agregar tests unitarios
- [ ] Documentación con Swagger
- [ ] Logging estructurado
- [ ] Métricas y monitoreo

## 🤝 Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit tus cambios (`git commit -am 'Agregar nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo `LICENSE` para más detalles.

