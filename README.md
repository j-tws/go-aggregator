# RSS Aggregator in a CLI

A guided project by [boot.dev](https://www.boot.dev) to create a CLI RSS Aggregator application aka `gator`.

Some extra text

## Pre-requisite

You will need Go and Postgres installed in order to get the CLI working. Here are some links to get you started with Go and Postgres:

- Go: https://go.dev/doc/install
- Postgres: https://www.postgresql.org/download/

## Installation

Run `go install github.com/j-tws/go-aggregator@latest`

## Setting up config

Create a `.gatorconfig.json` file in your home folder with the following content:

```json
{
  "db_url": "postgres://username:@localhost:5432/gator?sslmode=disable"
}
```

The `db_url` key is where you will store your postgres connection string. It will usually be the format of `postgres://<your-username>:@localhost:<pg-port>/gator?sslmode=disable`

## Usage

### `register`

Before using the CLI app, you will need to register yourself. It is a very simple registration that will take a name as an input.

```bash
go-aggregator register justin
# User 'justin' successfully registered and set as current user!
```

You can't register the same name. You will receive an error if you do so.

### `users`

You can browse registered users with this command:

```bash
go-aggregator users
# * kahya
# * justin (current)
# * jason
```

### `login`

In order to save your feeds and posts, you will need to be logged in.

```bash
go-aggregator login justin
# User has been set as justin.
```

If you logged in with an unregistered name, it will respond with an error.

```bash
go-aggregator login testing123
# 2025/04/09 22:02:57 No such user in database. Please register.
```

To check if you are logged in with the correct name, run `go-aggregator users` and the name with `(current)` appended indicates that is the current user with that name logged in.

```bash
go-aggregator login justin
# User has been set as justin.

go-aggregator users
# * kahya
# * justin (current)
# * jason
```

## `addfeed` and `feeds`

To store a feed, run `go-aggregator addfeed <feed-name> <feed-url>`.
It will also prompt you to automatically follow the feed that you just added.

```bash
go-aggregator addfeed "Hacker News" "https://news.ycombinator.com/rss"
# Successfully created feed!
# Successfully created feed follow!
# Feed name: Hacker News
# Feed URL: https://news.ycombinator.com/rss
# User name: justin
```

To browse all the feeds that's been stored in the database (regardless of users) run `go-aggregator feeds`

```bash
go-aggregator feeds
# ----------
# Feed name: Hacker News
# Feed URL: https://news.ycombinator.com/rss
# Feed User Name: justin
# ----------
# Feed name: TechCrunch
# Feed URL: https://techcrunch.com/feed/
# Feed User Name: jason
# ----------
```

## `follow`

To follow a feed, run `go-aggregator follow <feed-url>`

```bash
go-aggregator follow "https://news.ycombinator.com/rss"
# Successfully followed feed!
# Feed name: Hacker News
# User name: justin
```

If you follow a feed that has yet not existed in the database, or a feed that you have already followed, it will respond with an error.

## `following`

This checks all the feeds that you are currently following.

```bash
go-aggregator following
# Current feeds for user 'jason':
# 0, Feed Name: Hacker News
# 1, Feed Name: TechCrunch
```

## `unfollow`

Use this command if you want to unfollow a feed.

```bash
go-aggregator unfollow "https://techcrunch.com/feed/"
# TechCrunch unfollowed successfully!
```

If unfollow a feed that you have yet to follow you will be responded with an error.

## `agg`

This command will accept an argument of a time interval format. The format can be of `1h`, `1h10m`, `5m10s` etc.
It will then scrape all the feeds you have followed in every time interval you have given. All the posts that you have scraped will be saved into the database as well.

```bash
go-aggregator agg 10m
# Collecting feeds every 10m0s
# Found post: Hardening the Firefox Front End with Content Security Policies
# Post successfully created: Hardening the Firefox Front End with Content Security Policies
# Found post: Apache ECharts
# Post successfully created: Apache ECharts
# Found post: The best programmers I know
# Post successfully created: The best programmers I know
# Found post: 'Sun-Like' Stars
# Post successfully created: 'Sun-Like' Stars
# Found post: PostgreSQL Full-Text Search: Fast When Done Right (Debunking the Slow Myth)
# Post successfully created: PostgreSQL Full-Text Search: Fast When Done Right (Debunking the Slow Myth)
# Found post: Obituary for Cyc
# Post successfully created: Obituary for Cyc
# Found post: Brazil's government-run payments system has become dominant
# Post successfully created: Brazil's government-run payments system has become dominant
# Found post: Linux Kernel Defence Map – Security Hardening Concepts
# Post successfully created: Linux Kernel Defence Map – Security Hardening Concepts
# Found post: Show HN: DrawDB – open-source online database diagram editor (a retro)
# Post successfully created: Show HN: DrawDB – open-source online database diagram editor (a retro)
# Found post: Tailscale has raised $160M
# Post successfully created: Tailscale has raised $160M
# Found post: The guide to reduce screen time
# Post successfully created: The guide to reduce screen time
# Found post: Dockerfmt: A Dockerfile Formatter
# Post successfully created: Dockerfmt: A Dockerfile Formatter
# Found post: A new way to make graphs more accessible to blind and low-vision readers
# Post successfully created: A new way to make graphs more accessible to blind and low-vision readers
# Found post: The Barium Experiment
# Post successfully created: The Barium Experiment
# Found post: Nonlinear soundsheet microscopy:imaging opaque organs capillary/cellular scale
# Post successfully created: Nonlinear soundsheet microscopy:imaging opaque organs capillary/cellular scale
# Found post: NTATV: Bringing Windows NT (Windows XP, Windows 2003) to the Original Apple TV
# Post successfully created: NTATV: Bringing Windows NT (Windows XP, Windows 2003) to the Original Apple TV
# Found post: Better typography with text-wrap pretty
# Post successfully created: Better typography with text-wrap pretty
# Found post: DIY experimental reactor harnesses the Birkeland-Eyde process
# Post successfully created: DIY experimental reactor harnesses the Birkeland-Eyde process
# Found post: A year of Rust in ClickHouse
# Post successfully created: A year of Rust in ClickHouse
# Found post: Ask HN: Do you still use search engines?
# Post successfully created: Ask HN: Do you still use search engines?
# Found post: An Overwhelmingly Negative and Demoralizing Force
# Post successfully created: An Overwhelmingly Negative and Demoralizing Force
# Found post: Solving a “Layton Puzzle” with Prolog
# Post successfully created: Solving a “Layton Puzzle” with Prolog
# Found post: The order of files in your ext4 filesystem does not matter
# Post successfully created: The order of files in your ext4 filesystem does not matter
# Found post: How Netflix Accurately Attributes eBPF Flow Logs
# Post successfully created: How Netflix Accurately Attributes eBPF Flow Logs
# Found post: Decomposing factorial of 300K as the product of 300K factors larger than 100K
# Post successfully created: Decomposing factorial of 300K as the product of 300K factors larger than 100K
# Found post: Analytic Combinatorics – A Worked Example
# Post successfully created: Analytic Combinatorics – A Worked Example
# Found post: Thank HN: The puzzle game I posted here 6 weeks ago got licensed by The Atlantic
# Post successfully created: Thank HN: The puzzle game I posted here 6 weeks ago got licensed by The Atlantic
# Found post: Intelligence Evolved at Least Twice in Vertebrate Animals
# Post successfully created: Intelligence Evolved at Least Twice in Vertebrate Animals
# Found post: Show HN: Coroot – eBPF-based, open source observability with actionable insights
# Post successfully created: Show HN: Coroot – eBPF-based, open source observability with actionable insights
# Found post: How to Recognize Woodpeckers by Their Drumming Sounds
# Post successfully created: How to Recognize Woodpeckers by Their Drumming Sounds
# 2025/04/09 22:23:40 Feed Hacker News collected, 30 posts found
```

## `browse`

To browse all your posts, run `go-aggregator browse <amount>`

```bash
go-aggregator browse 5
# Here are all your saved posts:
# Title: The guide to reduce screen time
# URL: https://speedbumpapp.com/en/blog/how-to-reduce-screen-time/
# Description: <a href="https://news.ycombinator.com/item?id=43630661">Comments</a>
# Published at: 2025-04-09 10:34:42 +0000 +0000
# ------------------------------
# Title: Hardening the Firefox Front End with Content Security Policies
# URL: https://attackanddefense.dev/2025/04/09/hardening-the-firefox-frontend-with-content-security-policies.html
# Description: <a href="https://news.ycombinator.com/item?id=43630388">Comments</a>
# Published at: 2025-04-09 09:34:16 +0000 +0000
# ------------------------------
# Title: 'Sun-Like' Stars
# URL: https://www.centauri-dreams.org/2025/04/08/on-sun-like-stars/
# Description: <a href="https://news.ycombinator.com/item?id=43629887">Comments</a>
# Published at: 2025-04-09 08:01:38 +0000 +0000
# ------------------------------
# Title: The best programmers I know
# URL: https://endler.dev/2025/best-programmers/
# Description: <a href="https://news.ycombinator.com/item?id=43629307">Comments</a>
# Published at: 2025-04-09 06:02:01 +0000 +0000
# ------------------------------
# Title: NTATV: Bringing Windows NT (Windows XP, Windows 2003) to the Original Apple TV
# URL: https://github.com/DistroHopper39B/NTATV
# Description: <a href="https://news.ycombinator.com/item?id=43628500">Comments</a>
# Published at: 2025-04-09 03:00:01 +0000 +0000
# ------------------------------
```
