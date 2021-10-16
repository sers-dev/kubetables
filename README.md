# Kubetables
Distribute IPFilter rules across all nodes in a Kubernetes cluster.

## What does it do
This repository currently only hosts the application that is meant to be run
as daemon set in a Kubernetes Cluster which monitors the newly created
`Ktbans` Custom Resources in the cluster. It creates IPFilter rules for existing CR on startup
and reacts to events like add, modify or delete and manages IPFilter rules accordingly on the host it is running on. 

Uses/requires iptables on host machines as of the current state.
Uses Kubernetes Custom Resource via etcd as datastore of rule units.

## Local docker setup
* Copy `./docker/.env.tpl` to `./docker/.env` and change values according to your setup
* Build image `docker-compose -f docker/docker-compose.yml build kubetables`
* Run `docker-compose -f docker/docker-compose.yml up -d kubetables`

## Plans for the future
* An additional tool to automatically create Ktban Custom Resources by for example evaluating fail2ban output 
will be released here as well.
* Aims to support nftables as packet filtering tool and postgres as datastore in the future.
