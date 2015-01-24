# CoBS: The Container Build Service

CoBS builds container images for alternative hardware architectures, such as ARM and POWER. It can also be used as a regular build service, similar to [Docker Hub's automated builds](http://docs.docker.com/docker-hub/builds/) and [CoreOS's Quay](https://quay.io).

How is CoBS different?

- CoBS will be open source and free to use behind your firewall.
- CoBS supports multiple CPU architectures, including ARM and POWER.
- CoBS doesn't have a clever name. Please suggest one: see Issue #1.

## Frequently Asked Questions

### Why?
Initial inspiration came from a question in #google-containers asking about Kubernetes support on ARM. I had recently acquired an ODROID-C1, so I decided to try it out. And that's where things went wrong.

Even though Docker images have an architecture property attached to them, this is currently ignored when doing a `docker run`. Since most images on Docker Hub are built for AMD64, trying to run them on ARM will result in an exec format error.

CoBS aims to bridge the gap between AMD64 images on Docker Hub and alternative hardware architectures.

### Where's the POWER?
Coming soon.

### Can the builder be run inside QEMU?
Probably. For Gopher Gala it was easiest to develop directly on an ARM machine.

### Does it support App Container/Rocket?
Soon, after Gopher Gala.

### Can't Docker/Quay fix everything by adding (and checking) a property and a couple builders?
Yup.
