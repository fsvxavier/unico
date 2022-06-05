# Project Name
> unico

## Indíce
- [Project Name](#project-name)
  - [Indíce](#indíce)
  - [Informação](#informação)
  - [Tecnologias](#tecnologias)
  - [Instalação](#instalação)
    - [Considerações](#considerações)
  - [Ambiente](#ambiente)
  - [Arquitetura de pastas](#arquitetura-de-pastas)
    - [Diretórios](#diretórios)
      - [api](#api)
        - [handlers](#handlers)
      - [cobertura](#cobertura)
      - [db](#db)
      - [interfaces](#interfaces)
      - [middlewares](#middlewares)
      - [migrations](#migrations)
      - [docs](#docs)
      - [usecases/mocks e repositories/mocks](#usecasesmocks-e-repositoriesmocks)
      - [models](#models)
      - [repositories](#repositories)
      - [utils](#utils)
  - [Iniciando](#iniciando)
  - [Testes](#testes)

## Informação
Projeto de teste da entrevista 

## Tecnologias
* [GoLang](https://golang.org/) - Compilador da linguagem Go
* [Go Mod](https://github.com/golang/mod) - Gerenciador de dependencias
* [Gin](https://github.com/gin-gonic/gin) - Framework go

## Instalação
Clonando o projeto
``` bash
cd $PROJECT_HOME
git clone https://github.com/fsvxavier/unico.git
```
Instalando dependências
```
$ go mod init
```
Removendo dependencias indesejadas
``` bash
$ go mod tidy
```
Baixando as dependencias para a vendor local
``` bash
$ go mod vendor
```

Iniciando o Banco de dados Mysql pelo Docker (Já com carga inicial de dados)
``` bash
$ sudo docker-compose -f docker/mysql/docker-compose.yaml up --build --abort-on-container-exit
```

### Considerações
``` bash
#
# Ambiente testado foi Ubuntu 20.04 LTS 64 Bits
# Certifique-se que efetuou a instalação correta do go
# Certifique-se que você tenha um compilador C/C++ instalado e acessível globalmente
# IDEs recomendadas: visual studio code
# Existem algumas funções Make no projeto para facilitar algumas ações
# Os exemplos de consumo da API estão disponíveis no Swagger da aplicação
#

```

## Ambiente
Configurando as variáveis de ambiente

| Nome | Descrição | Valor Padrão | Obrigatório |
| -- | -- | -- | -- |
| PORT | Porta padrão que a API irá subir | 5000 | :white_check_mark: |
| HTTP_READ_TIMEOUT | Timeout de leitura das para o httpserver | 60 | :white_check_mark: |
| HTTP_WRITE_TIMEOUT | Timeout de escrita das para o httpserver | 60 | :white_check_mark: |
| MYSQLDBPASSWORD  | Senha do usuário da aplicação da base de dados | | :white_check_mark: |
| MYSQLDBUSER  | Usuário de aplicação da base de dados | | :white_check_mark: |
| MYSQLDBHOST | Host da base de dados | | :white_check_mark: |
| MYSQLDBNAME | Nome da base de dados | | :white_check_mark: |
| MYSQLDBPORT | Porta da base de dados | | :white_check_mark: |
| MYSQLDBMAXIDLECONNS | Define o número máximo de conexões no conjunto de conexões inativas | | :white_check_mark: |
| MYSQLDBMAXOPENCONNS | Define o número máximo de conexões abertas para o banco de dados | | :white_check_mark: |
| MYSQLDBCONNMAXLIFETIME | Define a quantidade máxima de tempo que uma conexão pode ser reutilizada | | :white_check_mark: |
| APP | Nome do app | unico | :white_check_mark: |
| LOGRUS_LOG_LEVEL | Nível de severidade do log a ser impresso | INFO | :white_check_mark: |
| GIN_MODE | | INFO | :white_check_mark: |
| VERSION_APP | versão da aplicação | INFO | :white_check_mark: |
| VERSION_API | versão da API | INFO | :white_check_mark: |

## Arquitetura de pastas
### Diretórios
```bash
unico
       |-- api
           |-- handlers
       |-- cobertura
       |-- config
           |-- db
       |-- docker
           |-- mysql
       |-- docs
       |-- interfaces
       |-- middlewares
       |-- migrations
       |-- models
       |-- repository
           |-- mocks
       |-- usecases
           |-- mocks
       |-- utils
       |-- .gitignore
       |-- README.md
```

#### api
Aqui temos artefatos que para iniciar a API e servir como delivery para os clientes que a utilizam
Está camanda depende da camandas usecases, repository, models.
#### cobertura
Aqui temos artefatos gerados dos testes com respectivos arquivos de cobertura.
##### handlers
Está camada vai receber o input do cliente que está consumindo este serviço. “Trata” os dados e envia para a camada usecase.
Está camada depende da camada usecases.
Aqui temos artefatos que auxiliam as chamadas de APIS.
Como por exemplo:    
    - Endpoints;
    - Funções comuns às chamadas de API.
#### db
Está camada trata as conexões com o banco de dados.
#### docker
Está camada tem as configuraçoes da docker para teste no git
#### docs
Está camada terá todas as informações do swagger
#### interfaces
Está camada terá todos os contratos definidos nas interfaces de usecases e repositories.
#### middlewares
Está camada contém os middlewares utilizados pela aplicação.
#### usecases/mocks e repositories/mocks
Reúne todos os artefatos que geram algum mock para o sistema.
#### migrations
Está camada vai migrations da aplicação.
#### models
Está camada vai armazenar qualquer object struct. Exemplo: Cliente, Estudante, Livro.
#### repositories
Repository vai armazenar qualquer manipulador de banco de dados ou até mesmo chamado HTTP para outros serviços.
#### utils
Reúne utilitários para auxiliar nos processos comuns aos testes ou configurações do mesmo.

## Iniciando
Buildando o projeto
``` bash
# execute o comando abaixo para buildar a aplicação e garantir que nada está quebrado
$ go build
```
Executando o projeto
``` bash
$ go run main.app or ./unico
```
## Testes
```bash
# Para execução dos testes automatizados executar o comando abaixo no terminal dentro da pasta da aplicação
$ go test -v -cover ./...

# Para gerar a interface mostrando todos os arquivos e as linhas "Covered", "Not Covered" e "Not Tracked":
$ go test ./... -coverprofile cobertura/fmtcoverage.html fmt
$ go test ./... -coverprofile cobertura/cover.out
$ go tool cover -html=cobertura/cover.out # em ambiente windows tirar o = após -html
$ go tool cover -html=cobertura/cover.out -o cobertura/cover.html # em ambiente windows tirar o = após -html
$ open 'cobertura/cover.html' file # em ambiente windows abrir externamente
```
## Gerando Swagger para commit
```bash
# Caso tenha alteração nas definições do swagger é necessário executar o comando abaixo assim alterando a pasta /docs e realizar o commit da mesma
$ swag init 

# caso seja exibido erro ao de comando não encontrado para o swag executar os comandos abaixo
$ export GOPATH="$HOME/go"
$ export PATH="$GOPATH/bin:$PATH"

```

## Link swagger 
```bash

# url para o swagger local, se atentar para a porta configurada para a palicação
http://localhost:5000/swagger/index.html#/

```
