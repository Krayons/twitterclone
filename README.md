<h3 align="center">Twitter Text Display Clone</h3>
<p align="center">Written in GO</p>
<p align="center">
	<a href="https://travis-ci.org/Krayons/twitterclone"><img src="https://img.shields.io/travis/Krayons/twitterclone.svg?label=linux+build"></a>
	<a href="https://goreportcard.com/report/Krayons/twitterclone"><img src="https://goreportcard.com/badge/github.com/Krayons/twitterclone"></a>
	<br>
</p>

---

## Building

To build the project GO will need to be installed and correctly configured. It is recommened to use the latest 1.x version.

```
git clone git@github.com:Krayons/twitterclone.git
cd twitterclone/cmd
go build *.go
./texttwitter
```

## Usage
```
./texttwitter -h
Usage of ./texttwitter:
  -tweets string
        file location of the tweets file (default "./tweets.txt")
  -users string
        file location of the users file (default "./users.txt")
```

## Example
![example](https://i.imgur.com/FbCtLZ5.png)

Assumptions:
* Textfile for users and tweets can fit in memory.
* Usernames are ascii and are only valid is they are contain only A-Z,a-z,0-9
* If the users file has a single user multiple times only the latest one will be assumed correct
