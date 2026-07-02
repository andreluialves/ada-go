// Trace: GET /usuarios/123

// └── handler_http                         120ms
//
//	└── buscar_usuario                   115ms
//	    ├── validar_parametros             2ms
//	    ├── consultar_cache                8ms
//	    │   └── cache_miss
//	    ├── consultar_banco               90ms
//	    └── montar_resposta                5ms
package main

import (
	"context"

	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
)

func consultarBanco(ctx context.Context, userID string) error {
	// Aqui seria uma chamada real ao banco.
	// Por enquanto, só vamos simular sucesso.
	return nil
}

func buscarUsuario(ctx context.Context, userID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "buscar_usuario")
	defer span.Finish()

	span.SetTag("user.id", userID)

	err := consultarBanco(ctx, userID)
	if err != nil {
		span.SetTag("error", true)
		span.LogFields(
			otlog.String("event", "erro_ao_buscar_usuario"),
			otlog.String("message", err.Error()),
		)
		return err
	}

	span.LogFields(
		otlog.String("event", "usuario_encontrado"),
	)

	return nil
}

func main() {
	buscarUsuario(context.Background(), "123456")
}
