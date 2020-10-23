# Malicious ONBUILD demo

> Demo resources to show how `ONBUILD` instruction could be used to build malicious container images.
> <br><br> 

*Go demo binary build with [saschagrunert/demo](https://github.com/saschagrunert/demo) framkework.*

---

## Demo info
- Malicious container image:
  - `containers/malicious`

- SFTP:
  - Uses `mayth/simple-upload-server` docker image

- Innocent container images:
  - `containers/inno`, `containers/inno-2`, `containers/inno-3`

- Go binary:
  - Files: `main.go`, `go.mod`, `go.sum` and `pkg/*.go`
  - Binary already build: `malicious-onbuild-demo`
    ```
    ➜ ./malicious-onbuild-demo --h
    NAME:
      Malicious ONBUILD demo - Malicious ONBUILD instruction demo

    USAGE:
      malicious-onbuild-demo [global options] command [command options] [arguments...]

    AUTHOR:
      Andoni Alonso <andonialonsof@gmail.com>

    COMMANDS:
      help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
      --all, -l                     run all demos (default: false)
      --auto, -a                    run the demo in automatic mode, where every step gets executed automatically (default: false)
      --auto-timeout auto, -t auto  the timeout to be waited when auto is enabled (default: 3s)
      --continuously, -c            run the demos continuously without any end (default: false)
      --hide-descriptions, -d       hide descriptions between the steps (default: false)
      --immediate, -i               immediately output without the typewriter animation (default: false)
      --skip-steps value, -s value  skip the amount of initial steps within the demo (default: 0)
      -0, --onbuild_basic           Explain what is ONBUILD instruction (default: false)
      -1, --show_malicious          Show onbuild_malicious container (default: false)
      -2, --use_malicious_1         Use onbuild_malicious container I (default: false)
      -3, --what_happened           What happened using onbuild_malicious container (default: false)
      -4, --use_malicious_2         Use onbuild_malicious container II (default: false)
      -5, --conclusions             Conclusions about ONBUILD instruction (default: false)
      --help, -h                    show help (default: false)
    ```


If you want to play without the demo binary you can use check the `Makefile`:
```
➜ make
Available rules:

build-go            Build go binary
build-innocent      Build client image from local malicious image
build-innocent-2    Build client from local malicious image + use.dockerignore
build-innocent-3    Build client from Dockerhub malicious image + use.dockerignore
build-malicious     Build local malicious image
open-upload         Shows SFTP content
run                 Run all demo steps
run-fast            Run all demo steps (without typewritter animation)
server-up           Start SFTP local server
```