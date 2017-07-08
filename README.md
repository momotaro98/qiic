# qiic

qiic is command line tool for Qiita service data using their v1 API.

![qiic_demo](https://github.com/momotaro98/my-project-images/blob/master/qiic/demo.gif)

## Features

* get your **20 Stocked** Qiita articles
* show the articles with readable table
* open the articles in your browser with an allocated number

## Installation

### Option 1: Go Get
If you have installed go platform, executing `go get` command is easy to install.

To install or update the qiic binary into your $GOPATH as usual, run:

```bash
$ go get -u github.com/momotaro98/qiic
```

### Option 2: Binary

If you're **Mac** user, download the latest binary from the [Releases page](https://github.com/momotaro98/qiic/releases).

It's the easiest way to get started with `qiic`.

## Setup

0. **Set Environment Variable `QIITA_USERNAME`, your Qiita user name.**
0. Execute first update command.

```bash
$ qiic u
```

## Usage

### Commands:

```
u | update          update stocked articles to local
l | ls | list       list all local articles
a | access | open    generate gitignore files
```

### Basic Usage

get your 20 updated Qiita stocked article

```bash
$ qiic u  # get your 20 updated Qiita stocked articles and show them
```

You'll see the 20 articles' list
Exmaple:

```
┌────┬───────────────────────────────────────────────────┬─────────────────────┬─────┐
|A No|                       TITLE                       |         TAG         |STOCK|
|────|───────────────────────────────────────────────────|─────────────────────|─────|
|1   |Go のクロスコンパイル環境構築                            |Go                   |  354|
|────|───────────────────────────────────────────────────|─────────────────────|─────|
|2   |GitHubのリリース機能を使う                              |GitHub               |  272|
|────|───────────────────────────────────────────────────|─────────────────────|─────|
|3   |Golang Goの並列処理を学ぶ(goroutine, channel)         |Go,golang            |   28|
|────|───────────────────────────────────────────────────|─────────────────────|─────|
.
.
.
|────|───────────────────────────────────────────────────|─────────────────────|─────|
|20  |Go言語: var, init, mainが実行される順番                |Go                   |   10|
|────|───────────────────────────────────────────────────|─────────────────────|─────|
```

You can open the article in your browser with Access Number(A No)
with `qiic a [A No]`

```bash
$ qiic a 2  # Open the specified article (A No is 2) in your browser
```

### Page specifed update

You can get your 20 Qiita stocked article with page specified.

```bash
$ qiic u -p 2
```

Your 21-40th logged stocked articles will be got.

If p is 3, 41-60th alike.