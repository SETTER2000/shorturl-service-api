package api

import (
	pb "github.com/SETTER2000/shorturl-service-api/proto"
)

// UsersServer поддерживает все необходимые методы сервера.
type UsersServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedIShorturlServer

	// используем sync.Map для хранения пользователей
	//users sync.Map
}
