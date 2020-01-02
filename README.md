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

## 🚚 Install

```sh
go get github.com/tsirysndr/go-spotify
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-spotify"
```

Construct a new Spotify client, then use the various services on the client to access different parts of the Spotify API. For example:

```Go
client := spotify.NewClient("<YOUR TOKEN AUTHORIZATION>")
result, err := client.Track.List("6rqhFgbbKwnb9MLmUQDhG6")
```

## ✨ Coverage

Currently the following services are supported:

- [x] Search for an Item
- [x] Get Audio Analysis for a Track
- [x] Get Audio Features for Several Tracks
- [x] Get Audio Features for a Track
- [x] Get Several Tracks
- [x] Get a Track
- [ ] Get Current User's Profile
- [x] Get a User's Profile
- [ ] Remove Tracks from a Playlist
- [x] Get a List of Current User's Playlists
- [x] Get a Playlist's Tracks
- [x] Get a Playlist
- [x] Get a List of a User's Playlists
- [ ] Add Tracks to a Playlist
- [ ] Create a Playlist
- [ ] Reorder or replace a Playlist's Tracks
- [ ] Change a Playlist's Details
- [ ] Get the Current User's Recently Played Tracks
- [ ] Get Information About The User's Current Playback
- [ ] Get a User's Available Devices
- [ ] Get the User's Currently Playing Track
- [ ] Skip User's Playback To Next Track
- [ ] Skip User's Playback To Previous Track
- [ ] Pause a User's Playback
- [ ] Start/Resume a User's Playback
- [ ] Set Repeat Mode On User's Playback
- [ ] Seek To Position In Currently Playing Track
- [ ] Toggle Shuffle For User's Playback
- [ ] Transfer a User's Playback
- [ ] Set Volume For User's Playback
- [ ] Get User's Top Artists and Tracks
- [ ] Remove Albums for Current User
- [ ] Remove Tracks for Current User
- [ ] Check Current User's Saved Albums
- [ ] Check Current User's Saved Tracks
- [ ] Get Current User's Saved Albums
- [ ] Get Current User's Saved Tracks 
- [ ] Save Albums for Current User
- [ ] Save Tracks for Current User
- [ ] Unfollow Artists or Users
- [ ] Unfollow a Playlist
- [ ] Check if Current User Follows Artists or Users
- [ ] Get Followed Artists
- [ ] Check if Users Follow a Playlist
- [ ] Follow Artists or Users
- [ ] Follow a Playlist
- [ ] Get Available Genre Seeds 
- [ ] Get a List of Browse Categories
- [ ] Get a Single Browse Category
- [ ] Get a Category's playlists
- [ ] Get a List of Featured Playlists
- [ ] Get a List of New Releases
- [ ] Get Recommendations Based on Seeds
- [x] Get an Artist's Albums
- [x] Get an Artist's Related Artists
- [x] Get an Artist's Top Tracks
- [x] Get an Artist
- [x] Get Several Artists
- [x] Get an Album's Tracks
- [x] Get an Album
- [x] Get Several Albums

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
