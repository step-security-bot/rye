# Rye language 🌾

[![Build and Test](https://github.com/refaktor/rye/actions/workflows/build.yml/badge.svg)](https://github.com/refaktor/rye/actions/workflows/build.yml)
[![golangci-lint](https://github.com/refaktor/rye/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/refaktor/rye/actions/workflows/golangci-lint.yml)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/refaktor/rye/badge)](https://securityscorecards.dev/viewer/?uri=github.com/refaktor/rye)
[![Go Reference](https://pkg.go.dev/badge/github.com/refaktor/rye.svg)](https://pkg.go.dev/github.com/refaktor/rye)
[![Go Report Card](https://goreportcard.com/badge/github.com/refaktor/rye)](https://goreportcard.com/report/github.com/refaktor/rye)
[![GitHub Release](https://img.shields.io/github/release/refaktor/rye.svg?style=flat)](https://github.com/refaktor/rye/releases/latest)

- [Rye language 🌾](#rye-language-)
  - [What is Rye](#what-is-rye)
  - [Status: Alpha](#status-alpha)
  - [Overview](#overview)
    - [Some specifics](#some-specifics)
    - [Introductions](#introductions)
    - [Examples](#examples)
    - [Rye vs. Python](#rye-vs-python)
  - [Modules](#modules)
    - [Base - official integrations](#base---official-integrations)
    - [Contrib - will be community / third party integrations](#contrib---will-be-community--third-party-integrations)
  - [Follow development](#follow-development)
    - [Rye blog](#rye-blog)
    - [Ryelang reddit](#ryelang-reddit)
    - [Github](#github)
  - [Getting Rye](#getting-rye)
    - [Binaries](#binaries)
    - [Docker images](#docker-images)
      - [Binary Docker image](#binary-docker-image)
      - [Dev Docker image](#dev-docker-image)
    - [Building Rye](#building-rye)
      - [Build Rye with specific modules](#build-rye-with-specific-modules)
      - [Build WASM version of Rye](#build-wasm-version-of-rye)
  - [Related links](#related-links)
  - [Questions, contact](#questions-contact)

*visit **[ryelang.org](https://ryelang.org/)**, **[our blog](https://ryelang.org/blog/)** or join our **[reddit group](https://reddit.com/r/ryelang/)** for latest examples and development updates*

## What is Rye

Rye is a high level, dynamic **programming language** based on ideas from **Rebol**, flavored by
Factor, Linux shells and Golang. It's still an experiment in language design, but it should slowly become more and
more useful in real world.

It features a Golang based interpreter and console and could also be seen as (modest) Golang's scripting companion as
Go's libraries are quite easy to integrate.

## Status: Alpha

Core ideas of the language are formed. Most experimenting, at least for this stage is done.
Right now, focus is on making the core and runtime useful for anyone who might decide to try it.

That means I am improving the Rye console, documentation and stabilizing core functions.

## Overview

### Some specifics

Rye is **homoiconic**, Rye's code is also Rye's data.

```red
data: { name "Jim" score 33 }
code: { print "Hello " + data/name }
if data/score > 25 code
; outputs: Hello Jim
print second data + ", " + second code
; outputs: Jim, Hello
```

Rye has **no keywords or special forms**. Ifs, Loops, even Function
definiton words are just functions. This means you can have more of them,
load them at library level and create your own.

```red
if-jim: fn { name code } { if name = "jim" code }

visitor: "jim"
if-jim visitor { print "Welcome in!" }
; prints: Welcome in!
```

Rye has no statements. Everything is an **expression**, returns 
something. Even asignment returns a value, so you can assign
inline. Either (an if/else like function) returns the result of the evaluated
block and can be used like an ternary operator.

```red
direction: 'in
full-name: join3 name: "Jane" " " surname: "Doe"  
print either direction = 'in { "Hi" +_ full-name } { "Bye bye" +_ name }
; outputs: Hi Jane Doe
```

Rye has **more syntax types** than your average language.
And it has generic methods for short function names. *Get* and *send*
below dispatch on the *kind* of first argument (http uri and an email address).

```red
email: jim@example.com
content: html->text get http://www.example.com
send email "title" content
; sends email with contents of the webpage
```

Rye has an option of forward code flow. It has a concept of 
**op and pipe words**. Every function can
be called as ordinary function or as op/pipe word. It also 
has a concept of **injected blocks** like *for* below.

```red
http://www.example.com/ .get .html->text :content
jim@example.com .send "title" content
; sends email with contents of the webpage
{ "Jane" "Jim" } |for { .prn }
; outputs: Jane Jim
get http://www.example.com/users.json
|parse-json |for { -> "name" |capitalize |print }
; outputs capitalized names of users
```

Rye has **higher order like functions**, but they come in what
would usually be special forms (here these are still just functions).

```red
nums: { 1 2 3 }
map nums { + 30 }
; returns { 31 32 33 }
filter nums { :x all { x > 1  x < 3 } }
; returns: { 2 }
strs: { "one" "two" "three" }
{ 3 1 2 3 } |filter { > 1 } |map { <-- strs } |for { .prn }
; prints: three two three
```
*[more about HOF-s](https://ryelang.blogspot.com/2022/02/higher-order-functions-test.html)*

Rye has some different ideas about **failure handling**. This
is a complex subject, so look at other docs about it. Remember it's
all still experimental.

```red
read-all %mydata.json |check { 404 "couldn't read the file" }
|parse-json |check { 500 "couldn't parse JSON" }
-> "score" |fix { 50 } |print1 "Your score: {}"
; outputs: Your score: 50
; if file is there and json is OK, but score field is missing
```

Most languages return with an explicit keyword *return*. Rye, like lisps 
always returns the result of the **last expression**. But Rye also has
so called **returning words** which for visual clarity start with **^** 
and always or conditionally return to caller.

```red
add-nums: fn { a b } { a + b }
digits: fn1 { = 0 |either { "zero" } { "one" } }
percentage: fn { a b } { 100 * a |/ b |^fix { "ERR" } |concat "%" }
percentage 33 77  ; returns: 42%
percentage 42 0   ; returns: ERR
unlock-jim: fn1 { = "Jim" |^if { print "Hi Jim" } print "Locked" }
unlock-jim "Jim"  ; prints: Hi Jim
unlock-jim "Jane" ; prints: Locked
```

Rye has **multiple dialects**, specific interpreters of Rye values. 
Two examples of this are the validation and SQL dialects.

```red
dict { "name" "jane" "surname" "boo" }
|validate { name: required score: optional 0 integer } |probe
; prints: #[ name: "jane" score: 0 ]

name: "James"
db: open sqlite://main.db
exec db { insert into pals ( name ) values ( ?name ) }
```
*more dialects: [html parsing](https://ryelang.blogspot.com/2021/04/html-parsing-dialect-demo.html), [conversion](https://ryelang.blogspot.com/2021/12/convert-function-and-conversion-dialect.html)*

Rye has a concept of **kinds with validators and converters**.

```red
person: kind 'person { name: required string calc { .capitalize } }
user: kind user { user-name: required }
converter person user { user-name: :name }
dict { "name" "jameson" } >> person >> user |print
; prints: <Context (user): user-name: <String: Jameson>>
```

Rye's **scope/context** is a first class Rye value and by this very malleable.
One of the results of this are isolated contexts.

```red
iso: isolate { print: ?print plus: ?add }
do-in iso { 100 .plus 11 |print }
; prints 111
do-in iso { 2 .add 3 |print }
; Error: Word add not found
 ```

### Introductions

These are all work in progress, also need a refresher, but can maybe offer some insight:

  * [Meet Rye](https://refaktor.github.io/rye/TOUR_0.html)
  * [Introduction for Python programmers](https://refaktor.github.io/rye/INTRO_1.html)
 
### Examples

These pages are littered with examples. You can find them on this **README** page, on our **blog**, on **Introductions**, but also:

  * [Examples folder](./examples/) - unorganized so far, will change
  * [Solutions to simple compression puzzle](https://github.com/otobrglez/compression-puzzle/tree/master/src/rye)

### Rye vs. Python

Python is the *lingua franca* and the *measuring stone* amongst dynamic programming languages, where Rye is also searching it's place. 
When learning about something new it makes sense to approach it from a familiar place.

  * [Less variables, more flows example vs Python](https://ryelang.blogspot.com/2021/11/less-variables-more-flows-example-vs.html)
  * [Simple compression puzzle - from Python to Rye solution](https://github.com/otobrglez/compression-puzzle/blob/master/src/rye/compress_jm_rec_steps.rye)
  
## Modules

The author of Factor once said that at the end *it's not about the language, but the libraries*. I can only agree, adding *libraries, and distribution*. Rye is still an experiment in language design, so it doesn't have anything like production level libraries. But to test the language with practical problems in mind, or because I needed something for my use of Rye, there are quite many modules already made. 

### Base - official integrations
  * Core builtin functions ⭐⭐⭐  🧪~80%
  * Bcrypt - password hashing
  * Bson - binary (j)son
  * Crypto - cryptographic functions ⭐
  * Email - email generation and parsing
  * Html - html parsing
  * Http - http servers and clients ⭐⭐
  * IO (!) - can be excluded at build time ⭐⭐ 
  * Json - json parsing ⭐⭐ 
  * Mysql - database ⭐
  * Postgresql - database ⭐
  * Psutil - linux process management
  * Regexp - regular expressions ⭐ 🧪~50%
  * Smtpd - smtp server (receiver)
  * Sqlite - database ⭐⭐
  * Sxml - sax XML like streaming dialect
  * Validation - validation dialect ⭐⭐ 🧪~50% 
  * Webview - Webview GUI
   
### Contrib - will be community / third party integrations
  * Amazon AWS
  * Bleve full text search 
  * OpenAI - OpenAI API
  * Postmark - email sending service
  * Telegram bot - telegram bots
  * Ebitengine - 2d game engine

legend: ⭐ priority , 🧪 tests
    
## Follow development

### Rye blog

For most up-to date information on the language and it's development visit our **[old](http://ryelang.blogspot.com/)** and **[new blog](http://ryelang.org/blog)**.

### Ryelang reddit

This is another place for updates and also potental discussions. You are welcome to join **[our reddit group](https://reddit.com/r/ryelang/)**. 

### Github

If code speaks to you, our Github page is the central location for all things Rye. You are welcome to collaborate, post Issues or PR-s, there is tons of things to do and improve :)

## Getting Rye

Rye is developed on Linux, but has also been compiled on macOS and Docker. If you need additional architecture or OS, post an Issue.

### Binaries

You can find precompiled Binaries for **Linux** and **macOS** under [Releases](https://github.com/refaktor/rye/releases).

Docker images are published under [Packages](https://github.com/refaktor/rye/pkgs/container/rye).

### Docker images

#### Binary Docker image

This image includes Linux, Rye binary ready for use and Emacs-nox editor.

Docker image: **[ghcr.io/refaktor/rye](https://github.com/refaktor/rye/pkgs/container/rye)**

Run it via:

```bash
docker run -ti ghcr.io/refaktor/rye
```

#### Dev Docker image

The repository comes with a local [Docker image](./.docker/Dockerfile) that builds rye and allows you to do so.

```bash
docker build -t refaktor/rye -f .docker/Dockerfile .
```

Run 🏃‍♂️ the rye REPL with:

```bash
docker run -ti refaktor/rye
```

### Building Rye

Use official documentation or lines below to install Golang 1.19.3 https://go.dev/doc/install

    wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    go version
    
Clone the main branch from the Rye repository. There is a submodule (a different repo) for contributed packages, hence the additional flag is needed:

    git clone https://github.com/refaktor/rye.git && cd rye

Build the tiny version of Rye

    go build -tags "b_tiny"

Run the rye file:

    ./rye hello.rye

Run the Rye Console

    ./rye

Install build-esential if you don't already have it, for packages that require cgo (like sqlite):

    sudo apt install build-essential

#### Build Rye with specific modules

Rye has many bindings, that you can determine at build time, so (currently) you get a specific Rye binary for your specific project. This is an example of a build with many bindings. 
I've been working on a way to make this more elegant and systematic, but the manual version is:

    go build -tags "b_tiny,b_sqlite,b_http,b_sql,b_postgres,b_openai,b_bson,b_crypto,b_smtpd,b_mail,b_postmark,b_bcrypt,b_telegram"
	
#### Build WASM version of Rye

Rye can also work inside your browser or any other WASM container. I will add examples, html pages and info about it, but to build it:

    GOOS=js GOARCH=wasm go build -tags "b_tiny,b_norepl" -o wasm/rye.wasm main_wasm.go

## Related links

  [**Rebol**](http://www.rebol.com) - Rebol's author Carl Sassenrath invented or combined together 90% of concepts that Rye builds upon.
  
  [Factor](https://factorcode.org/) - Factor from Slava Pestov taught me new fluidity and that variables are *no-good*, but stack shuffle words are even worse ;)
  
  [Red](https://red-lang.org) - Another language inspired by Rebol from well known Rebol developer  DocKimbel and his colleagues. A concrete endaviour, with it's low level language, compiler, GUI, ...
  
  [Oldes' Rebol 3](https://oldes.github.io/Rebol3/) - Rebol3 fork maintained by Oldes (from Amanita Design), tries to resolve issues without unnecessarily changing the language itself.
  
  [Arturo](https://arturo-lang.io/) - Another unique language that builds on Rebol's core ideas.

  [Charm](https://github.com/tim-hardcastle/Charm) - Not related to Rebol, but an interesting Go based language with some similar runtime ideas and challenges.
  
  [Ren-c](https://github.com/metaeducation/ren-c) - Rebol 3 fork maintained by HostileFork, more liberal with changes to the language. 
  
 ## Questions, contact
  
You can post an **[Issue](https://github.com/refaktor/rye/issues)** visit github **[Discussions](https://github.com/refaktor/rye/discussions)** or contact me through <a href="mailto:janko .itm+rye[ ]gmail .com">gmail</a> or <a href="https://twitter.com/refaktor">twitter</a>.</p>

