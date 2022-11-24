# ChessDotCom

ChessDotCom is a minimal wrapper for the official Chess.com [REST API](https://www.chess.com/news/view/published-data-api).

# Overview

This package is designed to be incredibly lightweight and simple to use, mirroring the offering of the official Chess.com API.


## Installation

ChessDotCom has been developed and fully tested with Go version `1.19`. While it will likely work with other versions, it is not recomended.


Install the package:
```
go get github.com/A-ndy-git/chessdotcom
````

Import into your project:
```
import (
  "github.com/A-ndy-git/chessdotcom"
)
```

## Usage

You <i>probably</i> shouldn't. ChessDotCom is primarily a personal education project. Nevertheless, its simple enough to provide value if you need it.

Create a new client instance, and you are done! Simples!

```
client := chessdotcom.New()

// Client is instantiated and ready to be used. E.g.
res, err := client.PlayerProfile("magnuscarlsen")
...
```

# Documentation

Automatically created `GoDoc` docs are available [here](https://go.pkg/). High level methods are listed below.

## Player

- <b>Profile</b> - <i>Get additional details about a player in a game</i>
```
client.PlayerProfile("magnuscarlsen")
```
------

- <b>Stats</b> - <i>Get ratings, win/loss, and other stats about a player's game play, tactics, lessons and Puzzle Rush score.</i>
```
client.PlayerStats("magnuscarlsen")
```
------

- <b>Games</b> - <i>Array of Daily Chess games that a player is currently playing.</i>
```
client.PlayerGames("magnuscarlsen")
```
------

- <b>Game Archive</b> - <i>Array of monthly archives available for this player.</i>
```
client.PlayerGameArchive("magnuscarlsen")
```
------

- <b>Games to Move</b> - <i>Array of Daily Chess games where it is the player's turn to act.</i>
```
client.PlayerGamesToMove("magnuscarlsen")
```
------

- <b>Archive PGN</b> - <i>standard multi-game format PGN containing all games for a month.</i>
```
client.PlayerAchivePGN("magnuscarlsen", "2022", "01")
```
------

- <b>Game Monthly Archive</b> - <i>Array of Live and Daily Chess games that a player has finished.</i>
```
client.PlayerGameMonthlyArchive("magnuscarlsen", "2022", "01")
```
------

- <b>Clubs</b> - <i>List of clubs the player is a member of, with joined date and last activity date.</i>
```
client.PlayerClubs("magnuscarlsen")
```
------

- <b>Tournaments</b> - <i>List of tournaments the player is registered, is attending or has attended in the past.</i>
```
client.PlayerTournaments("magnuscarlsen")
```
------

- <b>Titled Players</b> - <i>List of titled-player usernames.</i>
```
client.PlayerTitledPlayers()
```
------

## Club
## Tournament
## Team Match
## Country
## Daily Puzzle
## Point System Config
