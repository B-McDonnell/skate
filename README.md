# Skate

A personal key-value store. 🛼

Skate is simple and powerful. Use it to save and retrieve anything you’d
like—even binary data. It’s fully encrypted, backed up to the cloud (that you
can self-host if you want) and can be synced with all your machines.

```bash
# Store something (and sync it to the network)
skate set kitty meow

# Fetch something (from the local cache)
skate get kitty

# What’s in the store?
skate list

# Pull down the latest data
skate sync

# Spaces are fine
skate set "kitty litter" "smells great"

# You can store binary data, too
skate set profile-pic < my-cute-pic.jpg
skate get profile-pic > here-it-is.jpg

# Unicode also works, of course
skate set 猫咪 喵
skate get 猫咪

# For more info
skate --help
```

## Installation

Use a package manager:

```bash
# macOS or Linux
brew tap charmbracelet/tap && brew install charmbracelet/tap/skate

# Arch Linux (btw)
yay -S skate

# Nix
nix-env -iA nixpkgs.skate
```

Or download it:

* [Packages][releases] are available in Debian and RPM formats
* [Binaries][releases] are available for Linux, FreeBSD, OpenBSD, macOS, and Windows

Or just build and install it yourself. Go 1.16 or higher is required.

```bash
git clone https://github.com/charmbracelet/skate.git
cd skate
go install
```

[releases]: https://github.com/charm/skate/releases

## Other Features

### Databases

Sometimes you’ll want to separate your data into different databases:

```bash
# Database are automatically created on demand
skate set secret-boss-key@work-stuff password123

# Most commands accept a @db argument
skate set "office rumor"@work-stuff "penelope likes marmalade"
skate get "office rumor"@work-stuff
skate list @work-stuff

# Oh no, the boss is coming!
skate reset @work-stuff
```

### Linking

One of the most powerful features of Skate is its ability to link two machines
together so they have access to the same data. To link two machines together
just run:

```bash
skate link
```

And follow the instructions. Keep in mind that you'll need access to both
machines.

### Syncing

When you run `skate set`, data is encrypted and synced to the network.
When you `get`, however, data is loaded from the local cache to ensure it loads
ultra-fast. To fetch any new data from the server just run `skate sync`.

## Examples

Here are some of our favorite ways to use `skate`.

### Keep secrets out of your .bashrc

```bash
skate set gh_token GITHUB_TOKEN
echo 'export GITHUB_TOKEN=$(skate get gh_token)' >> ~/.bashrc
```

### Keep passwords in their own database

```bash
skate set github@password.db PASSWORD
skate get github@password.db
```

### Manage data with scripts

```bash
#!/bin/bash
skate set "$(date)@bookmarks.db" $1
skate list @bookmarks.db
```

What do you use `skate` for? [Let us know](mailto:vt100@charm.sh).

## Self-Hosting

Skate is backed by the Charm Cloud. By default it will use the Charm hosted
cloud, but it’s incredibly easy to self-host, even if that’s just on your
local network. For details, see the [Charm docs][selfhost].

[selfhost]: https://github.com/charmbracelet/charm#self-hosting

## Hey, Developers

Skate is built on [charm/kv](https://github.com/charmbracelet/charm/kv). If
you’d like to build a tool that includes a user key value store, check it out.

## License

[MIT](https://github.com/charmbracelet/skate/raw/master/LICENSE)

***

<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge-unrounded.jpg" width="400"></a>

Charm热爱开源 • Charm loves open source
