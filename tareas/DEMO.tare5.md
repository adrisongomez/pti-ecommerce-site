# Tare5: Desarrollo del backend de la Plataforma

## Requerimientos

- Diseñar y configurar bases de datos SQL y NoSQL
- Utilizar CodeIgniter 4 para estructurar el proyecto (5.3.1), definiendo rutas (5.3.2), modelos (5.3.3), vistas (5.3.4), controladores (5.3.5) y manejo de sesiones (5.3.6)
- Implementar funcionalidades CRUD (Crear, Leer, Actualizar, Eliminar) para los productos.
- Integrar bases de datos SQL (como MySQL o PostgreSQL) y NoSQL (como MongoDB) para el almacenamiento y recuperación de datos de los productos
= Asegurar que la aplicación sea segura y eficiente en la gestión de datos y sesiones.

## Notas

Para mi proyecto decide utilizar [PostgreSQL](https://www.postgresql.org/) como gestor de base datos. Es open sources, superflexibles y robusta. Ademas utilize el ORM [Prisma](https://www.google.com/search?q=prisma-go&rlz=1C5CHFA_enDO1024DO1024&oq=Prisma&gs_lcrp=EgZjaHJvbWUqBggAEEUYOzIGCAAQRRg7MgYIARBFGDsyDwgCEEUYORiLAxiABBj4BTIKCAMQABiLAxiABDIGCAQQRRg7MgYIBRBFGDwyBggGEEUYPDIGCAcQRRg80gEIMTk0N2oxajSoAgCwAgA&sourceid=chrome&ie=UTF-8) con su cliente de golang [Prisma-go](https://goprisma.org/). La ventaja de prisma contra otros ORM, este se basa en definir un schema de la base datos en archivos `.prisma` que luego este interpreta, genera las migraciones y el codigo del client, de manera que no es necesario escribir tanto codigo, solo modelar en los archivos `.prisma`. En mi caso, solo utilize un [archivo](../backends/databases/schema.prisma) y el diagrama ERM de mis entidades (HASTA AHORA) es:

<img width="824" alt="Screenshot 2025-03-27 at 4 13 55 PM" src="https://github.com/user-attachments/assets/4ba10772-fa04-40f3-8563-0dd68d123ea8" />

La aplicacion de backend esta escrita en go. Utilizando el framework de [Goa.design](https://goa.design/docs/1-introduction/). Este permite desarrollar tus APIs de una manera ma expresiva utilizando un `dsl`. Este nos permite definir, servicios, endpoints, objects... despues de definir todos los endpoints utilizando su CLI se genera el codigo de los servicios y del servidor HTTP y solo se debe implementar la parte que se encarga de escribir o leer de la base de datos. Por ejemplo, para definir el servicio de usuario se utiliza el siguiente codigo

```Go
// file backends/design/svc-users
package design

import (
	"github.com/adrisongomez/pti-ecommerce-site/backends/design/types"
	. "goa.design/goa/v3/dsl"
)

var PaginatedUser = types.PaginatedResult("user-list", types.User)

var _ = Service(servicePrefix+"user", func() {
	HTTP(func() { Path("/users") })
	Method("list", func() {
		Payload(func() {
			Attribute("pageSize", Int, "Record per page", func() {
				Minimum(10)
				Maximum(100)
				Default(10)
			})
			Attribute("after", Int, "Start listing after this resource", func() {
				Default(0)
			})
		})
		Result(PaginatedUser)
		HTTP(func() {
			GET("")
			Param("pageSize")
			Param("after")
			Response(StatusOK)
		})
	})
	Method("show", func() {
		Payload(func() {
			Attribute("userId", Int)
			Required("userId")
		})
		Result(types.User)
		HTTP(func() {
			GET("/{userId}")
			Param("userId")
			Response(StatusOK)
		})
	})
	Method("create", func() {
		Result(types.User)
		Payload(types.UserCreateInput)
		HTTP(func() {
			POST("")
			Response(StatusCreated)
		})
	})
	Method("update", func() {
		Payload(func() {
			Attribute("payload", types.UserCreateInput)
			Attribute("userId", Int)
			Required("payload", "userId")
		})
		Result(types.User)
		HTTP(func() {
			PUT("/{userId}")
			Param("userId")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("userId", Int)
			Required("userId")
		})
		Result(Boolean)
		HTTP(func() {
			DELETE("/{userId}")
			Param("userId")
			Response(StatusAccepted)
		})

	})
})
```

Luego que esta definido corro el comando de `make generate-svc` y este generara diferentes archivos en `/backends/internal` que no son versionados por que son generados. Cabe destacar que se utilizo `Makefile` como task runner en el backend. Dentro de esos archivo se define una interfaz Service que es la interfaz que debemos asegurar que se cumpla y luego se conecta el service al servicio HTTP.

```
package services

import (
	"context"
	"fmt"

	"github.com/adrisongomez/pti-ecommerce-site/backends/databases/db"
	svcuserhttp "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/http/svcuser/server"
	. "github.com/adrisongomez/pti-ecommerce-site/backends/internal/gen/svcuser"
	"github.com/adrisongomez/pti-ecommerce-site/backends/internal/utils"
	"go.uber.org/zap"
	"goa.design/goa/v3/http"
)

type UserController struct {
	client *db.PrismaClient
	logger *zap.Logger
}

func MapUserDBToOutput(model db.UserModel) *User {
	user := User{
		ID:        model.ID,
		FirstName: model.FirstName,
		LastName:  nil,
		Email:     model.Email,
		Role:      UserRole(model.Role),
		CreatedAt: model.CreatedAt.String(),
	}
	if value, ok := model.LastName(); ok {
		user.LastName = &value
	}

	if value, ok := model.UpdatedAt(); ok {
		user.UpdatedAt = utils.StringRef(value.String())
	}
	return &user
}

func (u *UserController) List(ctx context.Context, payload *ListPayload) (*UserList, error) {
	u.logger.Info("List got called With", zap.Any("payload", payload))
	usersDB, err := u.client.User.FindMany(
		db.User.DeletedAt.IsNull(),
	).
		Take(payload.PageSize).
		Skip(payload.After).Exec(ctx)
	users := []*User{}
	if err != nil {
		u.logger.Error("Error on requesting users", zap.Error(err))
		return nil, err
	}

	u.logger.Debug("DB response", zap.Any("usersDB", usersDB))
	for _, userDB := range usersDB {
		users = append(users, MapUserDBToOutput(userDB))
	}

	var rows []struct {
		Count db.BigInt `json:"count"`
	}
	err = u.client.Prisma.QueryRaw(
		"SELECT count(*) FROM project.users WHERE deleted_at IS NULL",
	).Exec(ctx, &rows)

	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("Not result from count")
	}
	count := int(rows[0].Count)

	nextCursor := utils.MinInt(count, payload.PageSize+payload.After)

	response := UserList{
		Data: users,
		PageInfo: &PageInfo{
			StartCursor:   payload.After,
			TotalResource: count,
			EndCursor:     nextCursor,
			HasMore:       count != nextCursor,
		},
	}
	return &response, nil
}

func (u *UserController) Create(ctx context.Context, payload *UserCreateInput) (*User, error) {
	u.logger.Info("User#create got called", zap.Any("payload", payload))
	changes := []db.UserSetParam{
		db.User.Role.Set(db.UserRole(payload.Role)),
	}
	if payload.LastName != nil {
		changes = append(changes, db.User.LastName.Set(*payload.LastName))
	}
	userModel, err := u.client.User.CreateOne(
		db.User.FirstName.Set(payload.FirstName),
		db.User.Email.Set(payload.Email),
		changes...,
	).Exec(ctx)

	if err != nil {
		u.logger.Error("User#create got error", zap.Error(err))
		return nil, err
	}

	return MapUserDBToOutput(*userModel), nil
}

func (u *UserController) Show(ctx context.Context, payload *ShowPayload) (*User, error) {
	u.logger.Info("User#Show got called with", zap.Any("payload", payload))
	userDB, err := u.client.User.FindUnique(db.User.ID.Equals(payload.UserID)).Exec(ctx)
	if err != nil {
		u.logger.Error("User#Show got error", zap.Error(err))
		return nil, err
	}
	return MapUserDBToOutput(*userDB), nil
}

func (u *UserController) Update(ctx context.Context, input *UpdatePayload) (*User, error) {
	payload := input.Payload
	u.logger.Info("User#create got called", zap.Any("payload", payload))
	changes := []db.UserSetParam{
		db.User.FirstName.Set(payload.FirstName),
		db.User.Email.Set(payload.Email),
		db.User.Role.Set(db.UserRole(payload.Role)),
	}
	if payload.LastName != nil {
		changes = append(changes, db.User.LastName.Set(*payload.LastName))
	}
	userModel, err := u.client.User.FindUnique(
		db.User.ID.Equals(input.UserID),
	).Update(
		changes...,
	).Exec(ctx)

	if err != nil {
		u.logger.Error("User#create got error", zap.Error(err))
		return nil, err
	}
	return MapUserDBToOutput(*userModel), nil
}

func (u *UserController) Delete(ctx context.Context, input *DeletePayload) (bool, error) {
	u.logger.Info("User#create got called", zap.Any("payload", input))

	_, err := u.client.User.FindUnique(db.User.ID.Equals(input.UserID)).Delete().Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUserService(client *db.PrismaClient) *UserController {
	logger := zap.L()
	return &UserController{client, logger}
}

func MountUserServiceSVC(mux http.Muxer, svc *UserController) {
	endpoints := NewEndpoints(svc)
	req := http.RequestDecoder
	res := http.ResponseEncoder

	handler := svcuserhttp.New(endpoints, mux, req, res, nil, nil)
	svcuserhttp.Mount(mux, handler)

	go func() {
		for _, mount := range handler.Mounts {
			zap.L().Info(fmt.Sprintf("%q mounted on %s %s\n", mount.Method, mount.Verb, mount.Pattern))
		}
	}()
}
```

Un punto muy importantes algunos endpoints como `POST products/` y `PUT products/:productId` escriben en diferentes entidades por ese motivo se utilizo transacciones para asegurar la integridad de las operaciones de manera que si encuentra un fallo en cualquiera de los queries, la base de datos no quede en un estado incompleto.

## Para correr el proyecto.
1- asegurarse exponer las variables de entornos en los archivos `.env.example`

```
make start-external-svc # Inicia un base de datos con docker
make migrate-sync 	# para sincronizar la base de datos con los archivos de migraciones
make generate-svc	# para generar los archivos de goa.design

air 			# para correr un process que realiza reload de la aplicacion cuando se realiza cambios del codigo

curl http://localhost:3030
```
 
## Demo video

https://github.com/user-attachments/assets/c1801b76-1602-4062-acf6-f688a581e08f

