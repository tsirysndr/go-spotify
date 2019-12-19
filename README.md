<h1 align="center">Welcome to go-spotify 👋</h1>
<h3 align="center">*** Work In Progress ***</h3>
<p align="center">
  <a href="https://github.com/tsirysndr/go-spotify/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-spotify.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-spotify">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-spotify">
  <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-spotify">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-spotify">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-spotify">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-spotify">
  <a href="https://github.com/tsirysndr/go-spotify/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>

> go-spotify is a Go client library for accessing the [Spotify API](https://developer.spotify.com/web-api/)

## Install

```sh
go get github.com/tsirysndr/go-spotify
```

## Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-spotify"
```

Construct a new Spotify client, then use the various services on the client to access different parts of the Spotify API. For example:

```Go
client := spotify.NewClient("<YOUR TOKEN AUTHORIZATION>")
result, err := client.Track.ListTracks("6rqhFgbbKwnb9MLmUQDhG6")
```

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
