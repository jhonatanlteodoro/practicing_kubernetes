# Crie um ambiente para hostear nosso pgsql

## Crie um novo namespace

*database.yaml :*
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: database
  labels:
    name: batabase
```

## Crie um novo PV e um PVC para ser usado no pgsql dentro do seu novo namespace
obs: esses recrusos devem ser usados dentro do nosso worker-2

### PV
*pv_local_storage.yaml*
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: local-pv-postgresql
  namespace: database
spec:
  capacity:
    storage: 10Gi
  volumeMode: Filesystem
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: local-storage
  local:
    path: /mnt/disks/ssd2
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - worker-2
```
```bash
$ kubectl create -f pv_local_storage.yaml
```

### PVC
*pvc_local_storage.yaml*
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: local-pvc-postgresql
  namespace: database
spec:
  accessModes:
    - ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 2Gi
  storageClassName: local-storage
```
```bash
$ kubectl create -f pvc_local_storage.yaml
```

## Crie um deployment para seu postgresql

*postgre_deployment.yaml*
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
 name: postgres
 namespace: database
spec:
 strategy:
   rollingUpdate:
     maxSurge: 1
     maxUnavailable: 1
   type: RollingUpdate
 replicas: 1
 selector:
   matchLabels:
     app: postgres
 template:
   metadata:
     labels:
       app: postgres
   spec:
     containers:
       - name: postgres
         image: postgres:10
         resources:
           limits:
             cpu: "0.5"
             memory: "700Mi"
           requests:
             cpu: "0.5"
             memory: "700Mi"
         ports:
           - containerPort: 5432
         env:
           - name: POSTGRES_PASSWORD
             value: password
           - name: PGDATA
             value: /var/lib/postgresql/data/pgdata
         volumeMounts:
           - mountPath: /var/lib/postgresql/data
             name: postgredb
     volumes:
       - name: postgredb
         persistentVolumeClaim:
           claimName: local-pvc-postgresql
```
```bash
$ kubectl create -f postgre_deployment.yaml
```

## Crie um banco de dados e uma tabela para ser usado por nossa golang app 2

```bash
$ kubectl get pods -n database
NAME                        READY   STATUS    RESTARTS   AGE
postgres-5c8c5d8d98-9gzjn   1/1     Running   0          36m
$ kubectl exec -it postgres-5c8c5d8d98-9gzjn -- psql -U postgres # Entrando no cli do pgsql dentro do pod
$ create database requests_goapp; # criando um novo db
$ \c requests_goapp; # setando o db criado para o seu contexto
$ CREATE TABLE requests (request_id SERIAL PRIMARY_KEY, date bigint NOT NULL); # criando uma tabela exemplo
$ \q # saindo do cli
```
