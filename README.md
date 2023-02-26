# t-core

Hey! What's up? Try to run in the first time:

```sh
make setup
```

> Maybe you have problems with the path variable, be careful

Then,

```sh
make all
```

you can run it every time you make any changes

## Release

First step:

```sh
docker build -t ghcr.io/acmlira/t-core:{tag} -f build/package/Dockerfile .
```

then, publish

```sh
docker push ghcr.io/acmlira/t-core:1.0
```

finally, recycle k8s objects.
