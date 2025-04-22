# Goutil

**Goutil** es una librería de utilidades para proyectos escritos en Go, mantenida por [@pinzlab](https://github.com/pinzlab). Este repositorio centraliza funciones, helpers y fragmentos de código que se repiten en varios proyectos dentro de esta cuenta, con el objetivo de evitar duplicación y facilitar el mantenimiento.

Aunque está optimizada para uso interno, es open source y estás invitado a explorar y usar si encuentras algo útil.

> ⚠️ Nota: La estructura puede cambiar con el tiempo a medida que evoluciona el ecosistema de proyectos en esta cuenta.

---

## ¿Qué incluye?

Este repositorio contiene utilidades generales que pueden incluir:

- **Generador de scripts para Postgres**:
  - Generación de scripts consultas complejas
  - Pensado para usarse junto con [`gorm`](https://gorm.io/)


---

## 📥 Instalación

Puedes agregar `goutil` a tu proyecto con:

```bash
go get github.com/pinzlab/goutil
```

## 📦 pg

Este paquete proporciona estructuras y funciones para facilitar la generación de SQL dinámico en PostgreSQL desde código Go. Está dividido en dos partes principales:

1. **Consultas dinámicas (como búsquedas insensibles a mayúsculas/acentos)**
2. **Migraciones y creación de objetos de base de datos**

### Importar

```go
import "github.com/pinzlab/goutil/pg"
```

### 🔍 Consultas

#### 1. Ilike – Búsqueda con ILIKE y UNACCENT

La estructura `Ilike` permite crear cláusulas `WHERE` con búsquedas insensibles a mayúsculas y acentos, utilizando `ILIKE` y `UNACCENT` en PostgreSQL.

```go
	ilike := pg.NewIlike("me", "first_name", "last_name")
	query := "SELECT * FROM users WHERE " + ilike.Where
	fmt.Println(query)

	// Output: 
	// SELECT * FROM users WHERE UNACCENT(first_name) ILIKE UNACCENT(@key) OR UNACCENT(last_name) ILIKE UNACCENT(@key)

```

### 🛠️ Migraciones

#### 1.- Enum – Creación condicional de tipos ENUM

Genera un script SQL que crea un tipo `ENUM` en PostgreSQL solo si no existe. Ideal para mantener la compatibilidad en entornos de desarrollo y producción sin errores por redefinición.


```go
	enum := pg.Enum{Name: "role_enum", Values: []string{"admin", "guest"}}
	fmt.Println(enum.GetScript())

	// Output:
	//	DO $$ BEGIN
	//		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
	//			CREATE TYPE role_enum AS ENUM ('admin', 'guest');
	//		END IF;
	//	END $$;
```

#### 2.- Foreign – Claves foráneas condicionales
Crea una clave foránea en una tabla específica, agregando restricciones de integridad referencial solo si no existen. Incluye reglas ON DELETE y ON UPDATE en cascada.

```go
	foreign := pg.Foreign{
		Table:       "profile",
		Reference:   "user",
		ForeignID:   "user_id",
		ReferenceID: "id",
	}
	fmt.Println(foreign.GetScript())

	// Output:
	// DO $$ BEGIN
	// 			IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname= 'fk_profile_user_id') THEN
	// 			ALTER TABLE public.profile ADD CONSTRAINT fk_profile_user_id
	// 			FOREIGN KEY (user_id) REFERENCES user(id)
	// 			ON UPDATE CASCADE ON DELETE CASCADE;
	// 		END IF;
	// END $$;
```

#### 3.- Insert (Entity) – Inserciones seguras
Permite insertar registros en una tabla si no existen previamente, evitando duplicaciones. Se puede usar para cargar datos iniciales o hacer "seed" de forma segura.

```go
	entity := pg.Entity{
		Table:   "user",
		Check:   []string{"username"},
		Columns: []string{"username", "role"},
		Values:  [][]any{{"myadmin", "admin"}},
	}
	fmt.Println(entity.GetScript())

	// Output:
	// INSERT INTO user(username, role) SELECT 'myadmin', 'admin'
	// WHERE NOT EXISTS (SELECT 1 FROM user WHERE username = 'myadmin');
```

#### 4.- Unique – Índices únicos con borrado lógico
Crea un índice único sobre una o más columnas, pero solo para registros que no estén marcados como eliminados `WHERE dat IS NULL`. Ideal para sistemas que implementan borrado lógico (soft delete).

```go
	unique := pg.Unique{
		Table:   "user",
		Columns: []string{"email", "username"},
	}
	fmt.Println(unique.GetScript())

	// Output:
	// CREATE UNIQUE INDEX IF NOT EXISTS uni_user_email_username
	// 	ON public.user(email, username)
	// WHERE dat IS NULL;
```