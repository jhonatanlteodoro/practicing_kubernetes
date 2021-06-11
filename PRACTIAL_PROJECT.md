# Objective: Learning/Practice Kubernetes


## Cluster
- Requirements
  *  2 Nodes + 1 Master Node
  *  Volume to read
  *  Volume readwrite
  *  NGINX ingress controller

## 3 APPS
* Golang APP
  - Requirements
    *  Namespace default
    *  volume to read write
    *  volume just to read
    *  configmaps
    *  secret key
    *  endpoint to read file
    *  endpoint to write file
    *  endpoint to helthcheck
    *  rolling update

* Golang APP 2
  - Requirements
    *  Namespace custom
    *  configmaps
    *  secret key
    *  endpoint to read content from a postgresql
    *  endpoint to write content to a postgresql
    *  endpoint to helthcheck
    *  rolling update

* Python APP
  - Requirements
    *  Namespace custom
    *  secret key
    *  rolling update
    *  endpoint to read redis
    *  endpoint to write redis
    *  configmaps
    *  secrets

## Redis
- Requirements
  *  Environment production-like

## Jenkins
- Requirements
  *  Environment production-like

## PostgreSQL
- Requirements
  *  Simple db setup that should provide a small table to read and write content for some apps
