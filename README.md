# Goutil

**Goutil** es una librería de utilidades para proyectos escritos en Go, mantenida por [@pinzlab](https://github.com/pinzlab). Este repositorio centraliza funciones, helpers y fragmentos de código que se repiten en varios proyectos dentro de esta cuenta, con el objetivo de evitar duplicación y facilitar el mantenimiento.

Aunque está optimizada para uso interno, es open source y estás invitado a explorar y usar si encuentras algo útil.

---

## ¿Qué incluye?

Este repositorio contiene utilidades generales que pueden incluir:

- **Generador de scripts para Postgres**:
  - Generación de scripts consultas complejas
  - Pensado para usarse junto con [`gorm`](https://gorm.io/)

> ⚠️ Nota: La estructura puede cambiar con el tiempo a medida que evoluciona el ecosistema de proyectos en esta cuenta.

---

## Instalación

Puedes agregar `goutil` a tu proyecto con:

```bash
go get github.com/pinzlab/goutil
