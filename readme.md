# ChessDotCom

ChessDotCom is a minimal wrapper for the official Chess.com [REST API](https://www.chess.com/news/view/published-data-api). Key features:

- All responses are fully typed.
- Logical beautification of field names. (e.g. `@id` becomes `ID`)
- Built-in handling of PGN data types.
- Internal request management to prevent API throttling.
- <b>ZERO</b> external dependencies!

# Overview

This package is designed to be incredibly lightweight and simple to use, mirroring the offering of the official Chess.com API.


## Installation

ChessDotCom has been developed and fully tested with Go version `1.19`. While it will likely work with other versions, it is not recomended.


Install the package:
```go
go get github.com/A-ndy-git/chessdotcom
````

Import into your project:
```go
import (
  "github.com/A-ndy-git/chessdotcom"
)
```

## Usage

You <i>probably</i> shouldn't. ChessDotCom is primarily a personal education project. Nevertheless, its simple enough to provide value if you need it.

Create a new client instance, and you are done! Simples!

```go
client := chessdotcom.New()

// Client is instantiated and ready to be used. E.g.
res, err := client.PlayerProfile("magnuscarlsen")
...
```

# Documentation

Automatically created `GoDoc` docs are available [here](https://go.pkg/). High level methods are listed below.

## Player

- <b>Profile</b> - <i>Get additional details about a player in a game</i>
```go
res, err := client.PlayerProfile("erik")
if err != nil {
  // handle error
}

// res ->
  {
    "Avatar": "https://images.chesscomfiles.com/uploads/v1/user/41.5434c4ff.200x200o.5b102889d835.jpeg",
    "PlayerId": 41,
    "ID": "https://api.chess.com/pub/player/erik",
    "URL": "https://www.chess.com/member/erik",
    "Name": "Erik",
    "Username": "erik",
    "Followers": 4340,
    "Country": "https://api.chess.com/pub/country/US",
    "Location": "Bay Area, CA",
    "LastOnline": 1669316678,
    "Joined": 1178556600,
    "Status": "staff",
    "IsStreamer": false,
    "Verified": true
  }
```
------

- <b>Stats</b> - <i>Get ratings, win/loss, and other stats about a player's game play, tactics, lessons and Puzzle Rush score.</i>
```go
client.PlayerStats("magnuscarlsen")
```
------

- <b>Games</b> - <i>Array of Daily Chess games that a player is currently playing.</i>
```go
client.PlayerGames("magnuscarlsen")
```
------

- <b>Game Archive</b> - <i>Array of monthly archives available for this player.</i>
```go
client.PlayerGameArchive("magnuscarlsen")
```
------

- <b>Games to Move</b> - <i>Array of Daily Chess games where it is the player's turn to act.</i>
```go
client.PlayerGamesToMove("magnuscarlsen")
```
------

- <b>Archive PGN</b> - <i>standard multi-game format PGN containing all games for a month.</i>
```go
client.PlayerAchivePGN("magnuscarlsen", "2022", "01")
```
------

- <b>Game Monthly Archive</b> - <i>Array of Live and Daily Chess games that a player has finished.</i>
```go
client.PlayerGameMonthlyArchive("magnuscarlsen", "2022", "01")
```
------

- <b>Clubs</b> - <i>List of clubs the player is a member of, with joined date and last activity date.</i>
```go
client.PlayerClubs("magnuscarlsen")
```
------

- <b>Tournaments</b> - <i>List of tournaments the player is registered, is attending or has attended in the past.</i>
```go
client.PlayerTournaments("magnuscarlsen")
```
------

- <b>Titled Players</b> - <i>List of titled-player usernames.</i>
```go
client.PlayerTitledPlayers()
```
------

## Club
## Tournament
## Team Match
## Country
## Daily Puzzle
## Point System Config
