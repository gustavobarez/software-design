# Functional Options Pattern

É um pattern usado para criar structs com configurações opcionais sem precisar de construtores com muitos parâmetros.

Em vez de:
```go
NewServer(maxConn, id, tls)
```
Usamos:
```go
newServer(withTls)
```
Cada option é uma função que altera a configuração antes do objeto ser criado.

---

## Ideia principal

- Existe uma struct de configuração (Opts)
- Existem valores padrão (default)
- Existem funções (options) que modificam essa configuração
- O construtor recebe várias options e aplica depois

Fluxo:

default -> aplica options -> cria objeto final

---

## Como funciona

Definimos o tipo da option:
```go
type OptFunc func(*Opts)
```
Qualquer função com essa assinatura pode ser usada como option.

Exemplo:
```go
func withTls(opts *Opts)
```
O construtor recebe:
```go
func newServer(opts ...OptFunc)
```
E executa:
```go
for _, fn := range opts {
    fn(&o)
}
```
---

## Benefícios

- parâmetros opcionais
- valores default
- fácil extensão sem quebrar a API
