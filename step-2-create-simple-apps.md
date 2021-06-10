## Create simple apps

Aqui vamos criar pequenos apps, cada um deles só precisa de uma rota e de um Dockerfile

### go app 1
```go
package main

import (
	"log"
	"net/http"
)

func JsonHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(
		[]byte(`{"hello": "I'm a json response", "from": "app go 1"}`),
	)
}

func main() {
	http.HandleFunc("/", JsonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
### go app 2
```go
package main

import (
	"log"
	"net/http"
)

func JsonHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(
		[]byte(`{"hello": "I'm a json response", "from": "app go 2"}`),
	)
}

func main() {
	http.HandleFunc("/", JsonHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### python app
```python
from flask import Flask


app = Flask(__name__)


@app.route("/")
def me_api():
    return {
        "hello": "I'm a json response", "from": "app python"
    }

##############################################################
-> requirements.txt
flask==2.0.1
```


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
