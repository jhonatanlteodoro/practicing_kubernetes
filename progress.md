# Acompanhando o processo

### 1º Etapa
* instalar e configurar o cluster kubernetes
  - Subir o ambiente no virtual box com vagrant
  - configurar os nodes
* criar pequenos apps de exemplo em go e python
  -  Criar uma aplicação simples em python
  -  Criar duas aplicações simples em golang
* criar pods exemplos para esses apps
  -  Criar um pod de forma declarativa para o app python
  -  Criar um pod de forma declarativa para o app1 golang
  -  Criar um pod de forma declarativa para o app2 golang

### 2º Etapa
*Golang app1 :*
* Criar um endpoint que possibilite a leitura de um json file
* Criar um endpoint que possibilite a escrita de um item em um json
* Criar um endpoint de healthcheck

*Cluster :*
* Adicionar um volume com permissão de leitura e escrita para guardar nosso json file do golang app1

*Golang app1 :*
* Modificar o pod deployment para que ele use o volume criado no passo anterior

### 3º Etapa
*Cluster :*
* Criar um namespace chamado databases para hospedar nosso querido postgresql
* Criar um volume para read/write para o postgresql
* Criar um volume claim para usar o volume criado para o postgresql
* Criar um config map para armazenar config do postgresql
* Criar uma secret para armazenar a senha do postgresql
* Criar um deployment para o postgresql
* Criar um service para o postgresql


