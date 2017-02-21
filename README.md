# github-bot
GitHub-Bot <img src="http://i.imgur.com/Cj4rMrS.gif" height="40" alt="Swimming Octocat" title="GH-Bot">
==========
GitHub-Bot was created and maintained by [Nyah Check](https://github.com/Ch3ck), and it's at GitHub bot to print the following and followers of any GitHub user. It follows the followers of any user and can unfollow all your following. The following is done by providing any user's following list and while authenticating with the GitHub API followers users based on that list. It uses the [Go-GitHub](github.com/google/go-github/github) library to authenticate with the GitHub API and was inspired by [Jessica Frazelle](https://github.com/jessfraz).

## Installation

## Usage
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
