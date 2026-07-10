# Design Patterns
---
## Layered Pattern
---
> Essa pattern geralmente e utilizada para organizar programas de forma sequencial que sejam divididos em subtasks. Cada layer fornece serviços para a proxima e nao se comunicam fora de ordem.

```mermaid
flowchart TB
L4[Layer 4] ---> L3[Layer 3] ---> L2[Layer 2] ---> L1[Layer 1]
4[Presentation Layer - UI] ---> 3[Application Layer - Service] ---> 2[Business Logic Layer - Domain] ---> 1[Data Access Layer - Persistency]
```

Exemplo:
- Desktop Apps
- E-Commerce Web Apps

## Client-Server Pattern
---
> Geralmente consiste de um servidor que se comunica com varios clientes, e esse componente servidor fornece os serviços pra uso desses multiplos clientes, e espera por requests desses clientes.

```mermaid
flowchart TB

S[Server]
C1[Client 1]
C2[Client 2]
C3[Client 3]

C1 -.-> |Request| S
C2 -.-> |Request| S
C3 -.-> |Request| S

S ---> |TCP/IP| C1
S ---> |TCP/IP| C2
S ---> |TCP/IP| C3
```

Exemplo:
- Provedor de E-mail
- Compartilhamento de Documentos
- Internet Banking

## Master-Slave Pattern
---
> Esse pattern conta com dois tipos de componentes: o mestre e o trabalhador. O mestre e responsavel por dar trabalho e distribuir esse trabalho com os trabalhadores.

```mermaid
flowchart TB
M[Master]
W1[Worker 1]
W2[Worker 2]
W3[Worker 3]

M ---> W1
M ---> W2
M ---> W3
```

Exemplo:
- Database replication
- Perifericos conectados no seu computador
- Loadbalancer

## Pipe-Filter Pattern
---
> Esse pattern gerencia dados na medida em que eles sao transportados por um "pipe". Cada passo dentro desse "pipe" e fechado por um componente de filtro.

```mermaid
flowchart LR
A{Source}
X{Sink}

F1(Filter 1)
F2(Filter 2)

A ---> |Pipe 1| F1
F1 ---> |Pipe 2| F2
F2 ---> |Pipe 3| X
```

Exemplo:
- Compiladores (Cada filtro e uma etapa da compilacao)
- Workflows
- Bouncers
- Parsers
- Middlewares

## Broker Pattern
---
>Esse pattern e muito usado para oregarizar e
>estruturar sistemas distribuidos de maneira em que
os componentes sejam os mais autonomos possivel.
Existe um componente chamado de Broker responsavel
por coordenar a comunicacao entre os componentes

```mermaid
flowchart LR
C[Client]
S1[Server 1]
S2[Server 2]
S3[Server 3]
subgraph Broker
B1[Process]
B2[Process]
B3[Process]
end
subgraph Servers
S1
S2
S3
end

B1 <--> S1
B2 <--> S2
B3 <--> S3

C <-.-> Broker
```

Exemplos:
- Loadbalancer
- Message Brokers(Apache Kafka, Rabbi…)


## Peer-to-Peer Pattern
---
>Nesse pattern cada componente individual e responsavel por servir e consumir de outros componentes igualmente autosuficientes

```mermaid
flowchart LR
subgraph Peer1
C1[client]
S1[server]
end
subgraph Peer2
C2[client]
S2[server]
end

Peer1 <--> |Direct Connection| Peer2
C1 <-.-> S2
S1 <-.-> C2
```

Exemplos:
- File Sharing (Torrent, G2)
- Protocolos Multimidia
- Crypto


## Event Bus Pattern
---
>Nesse pattern a comunicacao acontece via eventos, pra isso ser eficaz voce tera emissores, consumidores e canais pra passar esses eventos (Conjunto de canais geralmente e chamado de Bus). As fontes publicam eventos nos canais e os consumidores sao notificados dos novos eventos.

```mermaid
flowchart LR
S1[Source 1]
S2[Source 2]
C1[Channel 1]
C2[Channel 2]
L1[Listener 1]
L2[Listener 2]
subgraph Bus
C1
C2
end

S1 --> C1
S2 --> C2
S1 --> C2
C1 -.-> L1
C2 -.-> L1
C2 -.-> L2
```

Exemplo:
- Servicos de Notificacao
- Comunicacao entre Microservicos

## Model View Controller Pattern (MVC)

> Esse pattern divide a aplicação em três níveis de abstração: 
> - Model: Funcionalidade Central e Dados
> - View: Display pro usuário
> - Controller: Gerencia os inputs do usuário

```mermaid
flowchart LR
V[View]
C[Controller]
M[Model]
I[Input]

I --> |Input Event| V
V --> |View Control| C
C --> |Update Model| M
C --> |Query Model| M
M -.-> |Change Notification| C
```

Exemplo:
- Web Frameworks (Django, Rails)


## Blackboard Pattern

> Esse pattern geralmente é utilizado quando não dá pra resolver o problema utilizando um outro pattern.  
> Ele conta com 3 componentes principais:
> - Blackboard: Uma entidade global de memória contendo os objetos das soluções
> - Knowledge Source: módulos especializados com suas próprias representações
> - Control Component: operadores, executam módulos

> Todos esses componentes têm acesso à blackboard e eles podem gerar novos objetos que serão adicionados à blackboard.

```mermaid
classDiagram
Blackboard <-- KnowledgeSource
KnowledgeSource <-- Control
Blackboard <-- Control
Blackboard: solutions
Blackboard: inspect()
Blackboard: update()
class KnowledgeSource{
    updateBlackboard()
}
class Control{
    loop()
    nextSource()
}
```

Exemplo:
- Reconhecimento de Fala
- Identificação de Veículo, pessoa, etc.
- Interpretação de Sonar

## Interpreter Pattern

> Esse pattern é usado pra criar um componente que interpreta um programa utilizando uma linguagem dedicada.  
Ele é focado em analisar linhas e transformar em ações.

```mermaid
flowchart LR
subgraph Requester
C
Ctx
end

subgraph Interpreter
NonterminalExpression
AbstractExpression
TerminalExpression
end

C[Client]
Ctx[Context]

subgraph AbstractExpression
AE1["interpret(Context)"]
end

subgraph TerminalExpression
TE1["interpret(Context)"]
end

subgraph NonterminalExpression
NTE1["interpret(Context)"]
end

C --> Ctx
C --> AbstractExpression
TerminalExpression --> AbstractExpression
NonterminalExpression --> AbstractExpression
NonterminalExpression -.-> AbstractExpression
```

Exemplos:
- Databases SQL
- Interpretadores de Linguagem de Programacao