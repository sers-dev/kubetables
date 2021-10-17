# Kubetables
Distribute IPFilter rules across all nodes in a Kubernetes cluster.

## What does it do
This repository currently only hosts the application that is meant to be run
as daemon set in a Kubernetes Cluster which monitors the newly created
`Ktbans` Custom Resources in the cluster. It creates IPFilter rules for existing CR on startup
and reacts to events like add, modify or delete and manages IPFilter rules accordingly on the host it is running on. 

Uses/requires iptables on host machines as of the current state.
Uses Kubernetes Custom Resource via etcd as datastore of rule units.

## Prerequisites
* A Kubernetes cluster is mandatory
* This application currently only works with an etcd kubernetes data storage

## Steps required in any setup
* Create custom resource definition in your cluster `kubectl apply -f kubernetes/crd/crd.yml`

## Local docker setup
* Copy `./docker/.env.tpl` to `./docker/.env` and change values according to your setup
* Build image `docker-compose -f docker/docker-compose.yml build kubetables`
* Run `docker-compose -f docker/docker-compose.yml up -d kubetables`

## Deployment
### Kubernetes
The file `kubernetes/daemonset.yaml` includes all resources necessary for the application to work
as well as the daemonset itself, which are Namespace, ServiceAccount, Role, Rolebinding.
With one apply the application should be up and running without errors:
`kubectl apply -f kubernetes/daemonset.yaml`
In `kubernetes/crd/` is the Custom Resource Definition for the newly created ktban-Resource. 
There is not only the definition but an example, that can be applied for testing purposes.

## Plans for the future
* An additional tool to automatically create Ktban Custom Resources by for example evaluating fail2ban output 
will be released here as well.
* Aims to support nftables as packet filtering tool and postgres as datastore in the future.

## Troubleshooting
* Problem: Docker container don't start, you see these lines in container logs:
```
kubetables    | panic: running [/sbin/ip6tables -t filter -S KUBETABLES 1 --wait]: exit status 3: modprobe: can't change directory to '/lib/modules': No such file or directory
kubetables    | ip6tables v1.8.7 (legacy): can't initialize ip6tables table `filter': Table does not exist (do you need to insmod?)
```
Solution: Execute on host: `sudo modprobe ip6table_filter`
Ref: https://ilhicas.com/2018/04/08/Fixing-do-you-need-insmod.html
Info: Will need to investigate this problems origin further and prevent it from happening