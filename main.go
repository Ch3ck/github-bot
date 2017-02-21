//Filename: Main.go
//Author: Nyah Check
//Purpose: GitHub Bot to increase following and print following list.
//Token: 91b804cf541f1e923004b11e95af94249192b54c
//Licence: GNU PL 2017


package main

import (
    "encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/oauth2"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
)

const (
	// LOGGER is what is printed for help/info output
	LOGGER = "github-bot - %s\n"
	// VERSION is the binary version.
	VERSION = "v1.0"
)

var (
	token    string
	interval string
    kmd      string
	lastChecked time.Time
    
	debug   bool
	version bool
)

type UserData struct {
    Login             string
	ID                int
	HTMLURL           string
	Location          string
	Email             string
}

func init() {
	// parse flags
	flag.StringVar(&token, "token", "", "GitHub API token")
	flag.StringVar(&interval, "interval", "30s", "check interval (ex. 5ms, 10s, 1m, 3h)")

	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(LOGGER, VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		fmt.Printf("%s", VERSION)
		os.Exit(0)
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if token == "" {
		token = "91b804cf541f1e923004b11e95af94249192b54c"
	}
}


func main() {
	//usr := "augusshire"
	var ticker *time.Ticker
	// On ^C, or SIGTERM handle exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			ticker.Stop()
			logrus.Infof("Received %s, exiting.", sig.String())
			os.Exit(0)
		}
	}()

	// Create the http client.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	// Create the github client.
	client := github.NewClient(tc)
	

	// Get the authenticated user, the empty string being passed let's the GitHub
	// API know we want ourself.
	user, _, err := client.Users.Get("")
	if err != nil {
		logrus.Fatal(err)
	}
	username := *user.Login

	logrus.Infof("Bot started for user %s.", username)
	getFollowing(client, username)
	getFollowers(client, username)
	//followUsers(client, usr)
    //unFollow(client, username)
}

// getFollowers iterates over all followers received for user.
func getFollowers(client *github.Client, username string) error {
    followers, _, err := client.Users.ListFollowers( username,nil)
	if err != nil {
		return err
	}

    //writes user details to file.
    saveData("logs/followers.json", followers)
	for _, flwr := range followers {
        fmt.Printf("%+v", flwr)
	}

	return nil
}


//saveData to file.
func saveData(file string, data []*github.User) (error) {
    in, err := os.Open(file)
     if err != nil {
           return err
     }
     defer in.Close()
    
    //serialize the data
    newdata :=  []UserData{}
    out, er := json.Marshal(data)
    err = json.Unmarshal(out, &newdata)
    out, er = json.Marshal(newdata)
    er = ioutil.WriteFile(file, out, 0644)
    //fmt.Fprintf(in, out)

    return er
}


// getFollowing iterates over the list of following and writes to file using a gob object
func getFollowing(client *github.Client, username string) error {
	
    following, _, err := client.Users.ListFollowing(username, nil)//to test properly whether to parse resp instead inloop
	if err != nil {
		return err
	}

    saveData("logs/following.json", following)//writes details to file.
	for _, flwg := range following {
        logrus.Infof("%+v", flwg)
	}

	return nil
}


// followUsers, gets the list of followers for a particular user and followers them on GitHub.
// This requires authentication with the API.
func followUsers(client *github.Client, username string) error {

    usrs, _, err := client.Users.ListFollowing(username, nil) //to test properly whether to parse resp instead inloop
	if err != nil {
		return err
	}

	for _, usr := range usrs {
		//Follow user
		res, e := client.Users.Follow(*usr.Login)
        if err != nil {
            panic(e.Error())
        }
        
        fmt.Printf("%+v", res)
	}

	return nil
}


// Unfollow all GitHub users on one's follower list.
func unFollow(client *github.Client, username string) error {
	
    usrs, _, err := client.Users.ListFollowing(username, nil)
	if err != nil {
		return err
	}

	for _, usr := range usrs {
		//Follow user
		res, e := client.Users.Unfollow(*usr.Login)
        if err != nil {
            panic(e.Error())
        }
        
        logrus.Infof("%+v", res)
	}

	return nil
}


func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
