# Why Make This?

I was unable to find a convenient way to preview markdown files before committing to a Github-hosted repo, so I built a file server that renders markdown as similarly to Github as possible.

# Overview

The primary purpose of this tool is to serve markdown files over HTTP and render them in a Github-like way. This includes support for relative links and [GFM mode](https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document). Markdown files (ie: those with the `.md` extension) are posted to Github's API for translation to html. This helps to ensure that the markdown is rendered as closesly as possible to how Github would do it on their site. Other files are served directly, byte for byte. 

# Installation

From your terminal, run:

```bash
$ go get -u github.com/kevin-cantwell/repoman/cmd/repoman
```

# Usage

Change directories into the root of your repo, then run:

```bash
$ repoman
```

The server is available at `localhost:9999`.

# TODOs

* ~~Render directory listings with Github styles, just like on Github.~~
  * ~~Sort files directory-first~~
* ~~Auto-render anything named README.md below directory listings if the directory contains such a filename, just like on Github.~~
* ~~Fix breadcrumbs to accurately display the current path to a given file.~~
* ~~Render code files with syntax highlighting and Github styles, just like on Github.~~
  * TODO: Support all Prism-supported languages
* Render binary files and images with Github styles, just like on Github.