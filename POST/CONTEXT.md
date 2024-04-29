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

### Cause

### WithCancel

### WithCancelCause

### BACKGROUND

### WithDeadline

### WithTimeout

### WithValue

### WithValue