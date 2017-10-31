# Agenda
Agenda - Command line program by Cobra

# Usage

By using Agenda, you can create your own account and do easy
meetings managements among your partners. Make sure they are all
registerred here.

Usage:

    Agenda [command]

Available Commands:

    clearmeeting  Remove all the meetings you host.
    createmeeting Create a meeting whose host is the current user.
    destroy       A brief description of your command
    help
    help          Help about any command
    login         login account.
    logout        Current user logout
    meeting       Query meetings of a specific time interval
    modifymember  Add or remove members from your meeting
    quitmeeting   A brief description of your command
    register      Register a user account.
    removemeeting A brief description of your command
    users         A brief description of your command.....

Use `Agenda help [command]` or `Agenda [command] --help` for more information about a command.

# Implementation

## Command

By using Cobra, a command line program template package `cmd` can be easily created and coded.

    Cobra init

With command `Cobra add [command]`, a command.go file is created. After all command files have been created, we start separate coding work -- each member implements four commands.

## Interface between `entity` and `cmd`

Besides commands, we also need to implement how to manage our users and meetings. This is done by implementing `entity` package, which contains two files `meeting_manager.go` and `user_manager.go`. These two files implements all functions, including file I/O, query, insert, delete and update, which are essential to `Agenda`.

### Usage of package `entity`

All the functions and vriables are self-explanatory here.

- Constants

        const MEETING_PATH string = "data/meetings.json"
        const const USER_PATH string = "data/users.json"

- func UpdateMeeting(meetings []Meeting)
- func AddOneMeeting(m Meeting)
- func GetMeetings() []Meeting
- func GetMeeting(title string) Meeting
- func PrintOneMeeting(m Meeting)
- func GetUsers() []User
- func DeleteOneUser(username string)
- func AddOneUser(u User)
- type `Meeting`

        type Meeting struct {
        	Title     string       `json:"title"`
        	Host      string       `json:"host"`
        	Members   []SimpleUser `json:"members"`
        	Starttime string       `json:"start_time"`
        	Endtime   string       `json:"end_time"`
        }

    * func (m *Meeting) HasUser(username string) bool
- type `User`
        type User struct {
            Username string `json:"username"`
            Password string `json:"password"`
            Email    string `json:"email"`
            Phone    string `json:"phone"`
        }

- type `SimpleUser`

        type SimpleUser struct {
            Username string `json:"username"`
        }

### Examples
