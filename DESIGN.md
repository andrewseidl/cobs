# Design

## Overview
There are three basic components: servers, hunters, and builders. The servers oversee everything and provide the backend API and frontend website. The hunters are responsible for finding the requested Dockerfile, modifying for the requested architectures, and tar'ing everything up. The builders build the images and push to the registry.

Users are allowed to request that a certain repository be built. When this happens, the hunters will attempt to find a Dockerfile for that repository. If one is found, it will configured and sent to the builders. If a relevant Dockerfile is not found, the user will be asked for more information.

By default, built images will reside in the `cobsarch` namespace, where `arch` is replaced with the results of `uname -m`. Repository names will be prepended with their original namespace. For example, if a user requests an ARMv7 build of `dockerfile/rethinkdb`, the resulting repo will be named `cobsarmv7l/dockerfile-rethinkdb`.

## Servers
### Frontend

### Backend

## Hunters

## Builders
