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
docker build -t acmlira/t-core:{tag} -f build/package/Dockerfile .
```

then, recycle k8s objects.