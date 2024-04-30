package main

import (
	"context"
	"log"
	"time"
)

type contexKey string

// criando as const para utilizarmos no context.WithValue
const (
	timeoutKey contexKey = contexKey("timeout")
	deadKey    contexKey = contexKey("dead")
)

// mais exeplos podem serem visto no site do Golang
// https://pkg.go.dev/context#pkg-examples
// https://go.dev/blog/context
func main() {

	log.Println(">>>>>context.TODO <<<<<")
	todo()

	log.Println("\n\n >>>>>context.Background <<<<<")
	ctx := background()

	log.Println("\n\n >>>>>context.WithValue <<<<<")
	ctx = value(ctx)

	log.Println("\n\n >>>>>context.WithDeadline <<<<<")
	deadline(ctx)

	log.Println("\n\n >>>>>context.WithTimeout <<<<<")
	timeout(ctx)
}

// context.TODO examplo
// é um context vazio e deve se usar quando o context não estiver pronto para uso
func todo() {

	// context.TODO retorna context.Context
	ctx := context.TODO()

	// print do context vazio
	log.Println(ctx)
}

// context.Background
// é um context vazio, que não tem o proposito de ser cancelado ou carregar parametros dentro dele, muito utilizado em funções main, init e testes;
func background() context.Context {

	// context.Background retorna context.Context
	ctx := context.Background()

	// print do context vazio
	log.Println(ctx)

	// vamos retornar um context nessa função, para ser usado em outros exemplos
	return ctx
}

// context.WithDeadline
// esse context tem por caracteristica ser finalizado num tempo futuro especificado e quando finalizado, retorna o context "PAI";
func deadline(ctx context.Context) {

	// pegando o valor do context withValue e transformando para inteiro
	dead := ctx.Value(deadKey).(int)

	// context.WithDeadline retorna (context.Context, context.CancelFunc)
	// observa que nesse exemplo, passamos o context e estamos adicionando o horário de agora + 2 segundos (valor do deadKey)
	ctxDL, cancel := context.WithDeadline(ctx, time.Now().UTC().Add(time.Second*time.Duration(dead)))
	defer cancel()

	// nesse o print, o context irá retornar o horário (time) para o deadline
	log.Println(ctxDL, "")

	// pegando o hoário de expiração, para ser comparar
	t, _ := ctxDL.Deadline()

	// aqui faremos um for comparando se já chegamos no time deseja e se chegou, faremos o cancel do context e saimos do for
	for {

		// estamos validando se o erro do dealine é diferente de nil.
		if ctxDL.Err() != nil {
			log.Println("error is: ", ctxDL.Err().Error())
			log.Println("time dealine", t.Format(time.RFC3339))
			log.Println("time now", time.Now().UTC().Format(time.RFC3339))
			break
		} else {
			log.Println(t.Format(time.RFC3339))
		}

		time.Sleep(time.Second * 1)
	}

}

// context.WithTimeout
// muito parecido com o WithDeadline, porém aqui estamos especificando o tempo máximo de execução, muito utilizado para conexões externas, tais como database, APIs externas, message queue, ....  Quando finalizado esse context, o mesmo retorna o context "PAI";
func timeout(ctx context.Context) {

	// pegando o valor do context withValue e transformando para inteiro
	t := ctx.Value(timeoutKey).(int)

	// context.WithTimeout retorna (context.Context, context.CancelFunc)
	// observa que nesse exemplo, passamos o context e estamos adicionado um time de 2 segundos
	ctxOut, cancel := context.WithTimeout(ctx, time.Second*time.Duration(t))
	defer cancel()

	// nesse o print, o context irá retornar o horário (time) para o deadline
	log.Println(ctxOut, "")

	// faremos um for comparando se já chegamos no time deseja e se chegou, faremos o cancel do context e saimos do for
	for {

		// estamos validando se o erro do context WithTimeout é diferente de nil
		if ctxOut.Err() != nil {
			log.Println("error is: ", ctxOut.Err().Error())

			break
		}

		time.Sleep(time.Second * 1)
	}
}

func value(ctx context.Context) context.Context {

	// estamos adicionando uma key com o nome de timeout, para usarmos no context WithTimeout
	ctx = context.WithValue(ctx, timeoutKey, 2)

	// estamos adicionando uma key com o nome de dead, para usarmos no context WithDeadline
	ctx = context.WithValue(ctx, deadKey, 2)

	// para acessarmos o valor da chave, basta usar da seguinte forma:
	log.Println(ctx.Value(timeoutKey))
	log.Println(ctx.Value(deadKey))

	return ctx
}
