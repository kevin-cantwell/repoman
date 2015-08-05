# Why Make This?

I was unable to find a convenient way to preview markdown files before committing to a Github-hosted repo, so I built a file server that renders markdown as similarly to Github as possible.

# Overview

The primary purpose of this tool is to serve markdown files over HTTP and render them in a Github-like way. This includes support for relative links and [GFM mode](https://developer.github.com/v3/markdown/#render-an-arbitrary-markdown-document). Markdown files (ie: those with the `.md` extension) are posted to Github's API for translation to html. This helps to ensure that the markdown is rendered as closesly as possible to how Github would do it on their site. Other files are served directly, byte for byte. 

