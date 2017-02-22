GitHub-Bot
==========
GitHub-Bot was created and maintained by [Nyah Check](https://github.com/Ch3ck), and it's at GitHub bot to print the following and followers of any GitHub user. It follows the followers of any user and can unfollow all your following. The following is done by providing any user's following list and while authenticating with the GitHub API followers users based on that list. It uses the [Go-GitHub](github.com/google/go-github/github) library to authenticate with the GitHub API and was inspired by [Jessica Frazelle](https://github.com/jessfraz).

## Installation

* [Go version 1.7](https://github.com/golang/go/releases/tag/go1.7.3)

Clone Git repo:

```
$ git clone git@github.com:Ch3ck/github-bot.git
$ cd github-bot
$ go get golang.org/x/oauth2/...
$ go get github.com/Sirupsen/logrus/...
$ go get github.com/google/go-github/...

```

## Build & Run

```
$ make
```

## Usage

Create a `GitHub` token which you will use in your application.

```
$ github-bot -h
github-bot - v1.0
  -d    run in debug mode
  -seconds int
        seconds to wait before checking for new events (default 30)
  -token string
        GitHub API token
  -v    print version and exit (shorthand)
  -version
        print version and exit
```

## License

GitHub-Bot is licensed under [The GNU GPL License (GNU)](LICENSE).
