# SISTEMA DE BIBLIOTECA – PRINCIPIOS SOLID (Go)

### Requisitos
- Go 1.21+
- No se requiere DB ni dependencias externas.

### Cómo ejecutar
1. Inicializa módulo
```bash
go mod init solid-go
go mod tidy
```
2. Ejecuta la demo
```bash
go run . 
```
Si ves el warning de “VCS stamping”, puedes ejecutarlo así:

```bash
go run -buildvcs=false main.go
```