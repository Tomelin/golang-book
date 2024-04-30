# Entendendo o uso de context no Golang


## Introdução
O **CONTEXT** no Golang é um pacote fundamental, principalmente falando em goroutines.  Podemos dizer que o context tem muitas utilidades dentro de um programa GO, "carregar" parametros de um lado para outro, tempo de execução, .....  ***"Também podemos dizer que o context pode controlar um ciclo de vida"***, ou seja, ele pode determinar quando o processo é cancelado. Veremos alguns exemplo a seguir.

Para termos um desenvolviento claro, a documentação do Go nos recomenda o seguinte:
* **NÃO** recomenda ter context em struct e sim colocar nas funções que realmente não necessárias;
* **NÃO** passar context nulo (nil) e sim usar o conext.TODO que veremos em seguida;
* Usar valores no context em solicitão de processos e APIs, e demais request e não passar como paramteros opcionais;
* Usar context em processos de GOROUTINE;
* Pode-de repassar o context em multiplas funcões.


> Programs that use Contexts should follow these rules to keep interfaces consistent across packages and enable static analysis tools to check context propagation:
>
> Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:
>
> Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.
>
> Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
> 
> The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
>
> -- <cite>[Package context documentation](https://pkg.go.dev/context)</cite>

## Tipo de context
Agora vamos enetder quais tipos de context existem e quando devemos usar cada um deles.  Após explicarmos cada tipode de context, vamos mostrar os exemplos:

* **context.TODO** é um context vazio e deve se usar quando o context não estiver pronto para uso;
* **context.Background** é um context vazio, que não tem o proposito de ser cancelado ou carregar parametros dentro dele, muito utilizado em funções main, init e testes;
* **context.Cause** é usado para retorno de outro context cancelado Com o Cause é possivel entender a causa do context cancelado;
* **context.WithCancel** é muito utilizado em channels, para indicar quando um channel foi finalizado e ter por padrão retornar o context "PAI";
* **context.WithCancelCause** se comporta igual ao WithCancel, mas ao invés de retornar o context "PAI", retorna a causa do cancelamento;
* **context.WithDeadline** esse context tem por caracteristica ser finalizado num tempo futuro especificado e quando finalizado, retorna o context "PAI";
* **context.WithDeadlineCause** o mesmo que o context WithDeadline, porém retorna a causa do cancelamento ao invés de retornar o context "PAI";
* **context.WithTimeout** muito parecido com o WithDeadline, porém aqui estamos especificando o tempo máximo de execução, muito utilizado para conexões externas, tais como database, APIs externas, message queue, ....  Quando finalizado esse context, o mesmo retorna o context "PAI";
* **context.WithTimeoutCause** o mesmo que o context WithTimeout, porém retornando a causa do cancelamento da execução;
* **context.WithValue** esse context tem o objetivo de carregar valores "key: value", recomenda-se que a chave (key) seja um tipo diferente de string:
  > The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions between packages using context. Users of WithValue should define their own types for keys. To avoid allocating when assigning to an interface{}, context keys often have concrete type struct{}. Alternatively, exported context key variables' static type should be a pointer or interface.
  >  [WithValue](https://pkg.go.dev/context#WithValue).
* **context.WithoutCancel** esse context não tem cicllo de vida (timeout ou deadline) , mas após finalizado retornar o context "PAI".


## Exemplos
Agora que já conhecemos os tipos de context existentes, podemos ver o exemplo de cada deles:

### TODO
```
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
```

### BACKGROUND
``` 
// context.Background
// é um context vazio, que não tem o proposito de ser cancelado ou carregar parametros dentro dele, muito utilizado em funções main, init e testes;
func background() context.Context {

	// context.Background retorna context.Context
	ctx := context.Background()

	// print do context vazio
	log.Println(ctx)

	// vamos retornar um context nessa função, para ser usado em outros exemplos
	retur
```

### WithValue
``` 
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
```

### WithDeadline
``` 
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
```

### WithTimeout
``` 
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
```

