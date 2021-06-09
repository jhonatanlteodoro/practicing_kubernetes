## Create simple apps

Aqui vamos criar pequenos apps, cada um deles só precisa de uma rota e de um Dockerfile

### go app 1

### go app 2

### python app


# Push docker images to docker-hub

* Nota1: Nenhuma das imagens aqui deve ser usada em ambiente de produção, aqui temos só o minimo do minimo para um teste bem ridiculo... :D
* Nota2: leia novamente a nota1 para garantir a absorção :)


### Basicamente precisamos fazer um build de nossas apps, testar localmente, e se tudo estiver certo enviar para o docker hub.

1º build image
```bash
$ docker build -t app_python .
```

2º run image
```bash
$ docker run -d --publish 5000:5000 app_python:latest
```
- 2.1 check service
    ```bash
    $ curl localhost:5000
    {"from":"app python","hello":"I'm a json response"}
    ```

3º push your image to docker hub
```bash
$ docker tag app_python jhonatanlteodoro/app_python
$ docker push jhonatanlteodoro/app_python:latest
```
## Repita o processo para as outras duas imagens
