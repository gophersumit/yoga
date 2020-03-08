![goreleaser](https://github.com/gophersumit/yoga/workflows/goreleaser/badge.svg)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgophersumit%2Fyoga.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgophersumit%2Fyoga?ref=badge_shield)

`Yoga` is a CLI to scaffold and generate Go powered Angular Application with SQL database.

## Installation

Download latest binaries from release page : https://github.com/gophersumit/yoga/releases

## Creating a new project

Create a new project using `yoga init`

```console
yoga init my-app
```
The command above will create a new project `my-app`.

```console
tree -I node_modules
.
├── client
│   └── angular
│       ├── angular.json
│       ├── browserslist
│       ├── e2e
│       │   ├── protractor.conf.js
│       │   ├── src
│       │   │   ├── app.e2e-spec.ts
│       │   │   └── app.po.ts
│       │   └── tsconfig.json
│       ├── karma.conf.js
│       ├── package.json
│       ├── package-lock.json
│       ├── README.md
│       ├── readme.txt
│       ├── src
│       │   ├── app
│       │   │   ├── app.component.css
│       │   │   ├── app.component.html
│       │   │   ├── app.component.spec.ts
│       │   │   ├── app.component.ts
│       │   │   ├── app.module.ts
│       │   │   └── app-routing.module.ts
│       │   ├── assets
│       │   ├── environments
│       │   │   ├── environment.prod.ts
│       │   │   └── environment.ts
│       │   ├── favicon.ico
│       │   ├── index.html
│       │   ├── main.ts
│       │   ├── polyfills.ts
│       │   ├── styles.css
│       │   └── test.ts
│       ├── static
│       │   └── readme.txt
│       ├── tsconfig.app.json
│       ├── tsconfig.json
│       ├── tsconfig.spec.json
│       └── tslint.json
├── go.mod
├── server
│   └── cmd
│       └── web
│           ├── app.go
│           └── main.go
└── yoga.json

```



## Builiding Project

use `yoga build` command to build both Angular and Go project together.

```console
$ yoga build
💪 building application 

Compiling @angular/core : es2015 as esm2015

Compiling @angular/common : es2015 as esm2015

Compiling @angular/platform-browser : es2015 as esm2015

Compiling @angular/platform-browser-dynamic : es2015 as esm2015

Compiling @angular/router : es2015 as esm2015
🤓 successfully build Angular code
🤓 successfully build Go code

$ ls
client  go.mod  server  web  yoga.json

```
## Start Application
run executable generated named `web`

```console
$ ./web 
2020/03/07 20:08:13 Web server is available on port 3030

```
<img src="https://yoga.gophersumit.com/img/homepage.png"></img>

## Generators and Support for SQL
Coming soon

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgophersumit%2Fyoga.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgophersumit%2Fyoga?ref=badge_large)