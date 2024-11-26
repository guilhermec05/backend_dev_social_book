# backend_dev_social_book


neste código foi implementado com golang, banco de dados mysql, com dockerização, autenticação com jwt token.

## requisitos
Possuir docker e docker compose instalado em sua máquina







# Instalação

comando de docker para inicializar, acesse a pasta /.docker.

```bash
  docker-compose up --build
```
ou se preferir usar em background.
```bash
  docker-compose up --build -d 
```
Depois de executar o docker, em qualquer qualquer sgbd que estiver instalado em sua maquina, conecta ao banco com as credenciais estão no ficheiro /.docker/docker-compose.yml, logo apos execute os comando que estão no ficheiro /sql/db.sql 

### observação 
pode ser que não consiga usar a api logo na execução do docker, se isto acontecer, no local do ficheiro docker-composer.yml, segue os passos abaixo:

- acesso o docker com o comando abaixo.
```bash
  docker-compose exec app sh
```
- execute aplicação
```bash
  go run main.go
```


## url para usar a aplicação

http://localhost:3001

