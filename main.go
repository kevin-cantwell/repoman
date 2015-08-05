package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/{filename:.*\.md}`, ServeGithubMarkdown).Methods("GET")
	r.HandleFunc(`/{filename:.*}`, ServeFile).Methods("GET")
	log.Fatal(http.ListenAndServe(":9999", r))
}

func ServeFile(response http.ResponseWriter, request *http.Request) {
	filename := mux.Vars(request)["filename"]
	http.ServeFile(response, request, filename)
}

func ServeGithubMarkdown(response http.ResponseWriter, request *http.Request) {
	filename := mux.Vars(request)["filename"]

	markdown, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(response, "Failed to read markdown file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	body := map[string]string{
		"text": string(markdown),
		"mode": "gfm",
	}
	b, err := json.Marshal(body)
	if err != nil {
		http.Error(response, "Failed to build Github API request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	r, _ := http.NewRequest("POST", "https://api.github.com/markdown", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		http.Error(response, "Error connecting to Github API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != 200 {
		http.Error(response, "Error connecting to Github API: "+resp.Status, http.StatusBadGateway)
		return
	}

	markdownHTML, err := ioutil.ReadAll(resp.Body)

	response.Header().Set("Content-Type", "text/html")
	response.Write(WrapMarkdownHTML(markdownHTML))
}

func WrapMarkdownHTML(markdown []byte) []byte {
	return []byte(`
<!DOCTYPE html>
<html lang="en" class=" is-copy-enabled">
  <head prefix="og: http://ogp.me/ns# fb: http://ogp.me/ns/fb# object: http://ogp.me/ns/object# article: http://ogp.me/ns/article# profile: http://ogp.me/ns/profile#">
    <meta charset='utf-8'>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta http-equiv="Content-Language" content="en">
    <meta name="viewport" content="width=1020">
    <meta content="origin-when-crossorigin" name="referrer" />
    
    <title>wiki/README.md at master · timehop/wiki</title>
    <link rel="search" type="application/opensearchdescription+xml" href="/opensearch.xml" title="GitHub">
    <link rel="fluid-icon" href="https://github.com/fluidicon.png" title="GitHub">
    <link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-114.png">
    <link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114.png">
    <link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-144.png">
    <link rel="apple-touch-icon" sizes="144x144" href="/apple-touch-icon-144.png">
    <meta property="fb:app_id" content="1401488693436528">

    <meta content="@github" name="twitter:site" /><meta content="summary" name="twitter:card" /><meta content="timehop/wiki" name="twitter:title" /><meta content="Contribute to wiki development by creating an account on GitHub." name="twitter:description" /><meta content="https://avatars1.githubusercontent.com/u/1259053?v=3&amp;s=400" name="twitter:image:src" />
    <meta content="GitHub" property="og:site_name" /><meta content="object" property="og:type" /><meta content="https://avatars1.githubusercontent.com/u/1259053?v=3&amp;s=400" property="og:image" /><meta content="timehop/wiki" property="og:title" /><meta content="https://github.com/timehop/wiki" property="og:url" /><meta content="Contribute to wiki development by creating an account on GitHub." property="og:description" />
    <meta name="browser-stats-url" content="https://api.github.com/_private/browser/stats">
    <meta name="browser-errors-url" content="https://api.github.com/_private/browser/errors">
    <link rel="assets" href="https://assets-cdn.github.com/">
    <link rel="web-socket" href="wss://live.github.com/_sockets/MzA3ODY0OjRlODUzZDljOTc3N2Y3M2QzNWFkYjJiYWQzM2ZhOWQzOjUwNmNlZDE5MmU1MTE5ZDJjMGZjYWYyYjNjNmM0ZDYzYTY3MjMyZDZjNzI5NjY1NjUyMDIyNWY4MzRlNzljMTc=--5ad2af2c625c426ae0002bdbb9fd038af2a8e44a">
    <meta name="pjax-timeout" content="1000">
    <link rel="sudo-modal" href="/sessions/sudo_modal">

    <meta name="msapplication-TileImage" content="/windows-tile.png">
    <meta name="msapplication-TileColor" content="#ffffff">
    <meta name="selected-link" value="repo_source" data-pjax-transient>

    <meta name="google-analytics" content="UA-3769691-2">

    <meta content="collector.githubapp.com" name="octolytics-host" /><meta content="collector-cdn.github.com" name="octolytics-script-host" /><meta content="github" name="octolytics-app-id" /><meta content="6C06CA4C:0A7D:BDD5980:55C1643B" name="octolytics-dimension-request_id" /><meta content="307864" name="octolytics-actor-id" /><meta content="kevin-cantwell" name="octolytics-actor-login" /><meta content="54a9e99ce298168945c2912a74c8a2adda41e1f8fa5be66ac3aaadef0bc0a10e" name="octolytics-actor-hash" />
    <meta content="/private/private/blob/show" data-pjax-transient="true" name="analytics-location" />
    <meta content="Rails, view, blob#show" data-pjax-transient="true" name="analytics-event" />
    <meta class="js-ga-set" name="dimension1" content="Logged In">
    <meta class="js-ga-set" name="dimension4" content="Current repo nav">
    <meta name="is-dotcom" content="true">
    <meta name="hostname" content="github.com">
    <meta name="user-login" content="kevin-cantwell">

    <link rel="icon" sizes="any" mask href="https://assets-cdn.github.com/pinned-octocat.svg">
    <meta name="theme-color" content="#4078c0">
    <link rel="icon" type="image/x-icon" href="https://assets-cdn.github.com/favicon.ico">

    <meta content="authenticity_token" name="csrf-param" />
    <meta content="T6UAYcu/kk7ROPCtRj+IeFJX1g1vxNVcxhs85+tWQeC33xRw3ZX7vU90YDAO4nMdMr656WYKA4R6B5MvbGHBug==" name="csrf-token" />
    

    <link crossorigin="anonymous" href="https://assets-cdn.github.com/assets/github/index-a6099c44cba81e3cc01a25d5aad205f2dd86c56ad656d0ee72761a3db28828c5.css" media="all" rel="stylesheet" />
    <link crossorigin="anonymous" href="https://assets-cdn.github.com/assets/github2/index-e7262cad01eef2691501230e7d70c976a8e97c8a96842f0f29a1f20f658a315d.css" media="all" rel="stylesheet" />
    
    


    <meta http-equiv="x-pjax-version" content="7b3fb0c762b1609124fa6ce9de586ac0">

      
    <meta name="description" content="Contribute to wiki development by creating an account on GitHub.">
    <meta name="go-import" content="github.com/timehop/wiki git https://github.com/timehop/wiki.git">

    <meta content="1259053" name="octolytics-dimension-user_id" /><meta content="timehop" name="octolytics-dimension-user_login" /><meta content="39162096" name="octolytics-dimension-repository_id" /><meta content="timehop/wiki" name="octolytics-dimension-repository_nwo" /><meta content="false" name="octolytics-dimension-repository_public" /><meta content="false" name="octolytics-dimension-repository_is_fork" /><meta content="39162096" name="octolytics-dimension-repository_network_root_id" /><meta content="timehop/wiki" name="octolytics-dimension-repository_network_root_nwo" />
    <link href="https://github.com/timehop/wiki/commits/master.atom?token=AASymJHRgilt7cPxQ38mSBuM6861rDr6ks6zzqTLwA%3D%3D" rel="alternate" title="Recent Commits to wiki:master" type="application/atom+xml">
  </head>


  <body class="logged_in  env-production macintosh vis-private page-blob">
    <div class="wrapper">
        <div class="header header-logged-in true" role="banner">
          <div class="container clearfix">
            <a class="header-logo-invertocat" href="https://github.com/" data-hotkey="g d" aria-label="Homepage" data-ga-click="Header, go to dashboard, icon:logo">
              <span class="mega-octicon octicon-mark-github"></span>
            </a>
          </div>
        </div>
        <div class="site" itemscope itemtype="http://schema.org/WebPage">
          <div class="pagehead repohead instapaper_ignore readability-menu ">
            <div class="container">
              <div class="clearfix">
                <h1 itemscope itemtype="http://data-vocabulary.org/Breadcrumb" class="entry-title private ">
                  <span class="author"><span itemprop="title">repoman</span></span><!--
               --><span class="path-divider">/</span><!--
               --><strong><a href="/" data-pjax="#js-repo-pjax-container">root</a></strong>
                  <span class="page-context-loader">
                    <img alt="" height="16" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-32.gif" width="16" />
                  </span>
                </h1>
              </div>
            </div>
          </div>
          <div class="container">
            <div class="repository-with-sidebar repo-container new-discussion-timeline ">
              <div class="repository-sidebar clearfix">
              



                <div class="only-with-full-nav">
                    
<div class="js-clone-url clone-url "
  data-protocol-type="http">
  <h3><span class="text-emphasized">HTTPS</span> clone URL</h3>
  <div class="input-group js-zeroclipboard-container">
    <input type="text" class="input-mini input-monospace js-url-field js-zeroclipboard-target"
           value="https://github.com/timehop/wiki.git" readonly="readonly" aria-label="HTTPS clone URL">
    <span class="input-group-button">
      <button aria-label="Copy to clipboard" class="js-zeroclipboard btn btn-sm zeroclipboard-button tooltipped tooltipped-s" data-copied-hint="Copied!" type="button"><span class="octicon octicon-clippy"></span></button>
    </span>
  </div>
</div>

  
<div class="js-clone-url clone-url open"
  data-protocol-type="ssh">
  <h3><span class="text-emphasized">SSH</span> clone URL</h3>
  <div class="input-group js-zeroclipboard-container">
    <input type="text" class="input-mini input-monospace js-url-field js-zeroclipboard-target"
           value="git@github.com:timehop/wiki.git" readonly="readonly" aria-label="SSH clone URL">
    <span class="input-group-button">
      <button aria-label="Copy to clipboard" class="js-zeroclipboard btn btn-sm zeroclipboard-button tooltipped tooltipped-s" data-copied-hint="Copied!" type="button"><span class="octicon octicon-clippy"></span></button>
    </span>
  </div>
</div>

  
<div class="js-clone-url clone-url "
  data-protocol-type="subversion">
  <h3><span class="text-emphasized">Subversion</span> checkout URL</h3>
  <div class="input-group js-zeroclipboard-container">
    <input type="text" class="input-mini input-monospace js-url-field js-zeroclipboard-target"
           value="https://github.com/timehop/wiki" readonly="readonly" aria-label="Subversion checkout URL">
    <span class="input-group-button">
      <button aria-label="Copy to clipboard" class="js-zeroclipboard btn btn-sm zeroclipboard-button tooltipped tooltipped-s" data-copied-hint="Copied!" type="button"><span class="octicon octicon-clippy"></span></button>
    </span>
  </div>
</div>



  <div class="clone-options">You can clone with
    <form accept-charset="UTF-8" action="/users/set_protocol?protocol_selector=http&amp;protocol_type=push" class="inline-form js-clone-selector-form is-enabled" data-form-nonce="08775a3354e7c6e73c74f48b232de5d7f4658854" data-remote="true" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="hp6nmIozVW+3OkqectrHPenjad5br0o8IHqMqo8g3a6u10+CGv1sBVdKiCpXIcxpMGk2TKMPaMA47FdqD90Pxw==" /></div><button class="btn-link js-clone-selector" data-protocol="http" type="submit">HTTPS</button></form>, <form accept-charset="UTF-8" action="/users/set_protocol?protocol_selector=ssh&amp;protocol_type=push" class="inline-form js-clone-selector-form is-enabled" data-form-nonce="08775a3354e7c6e73c74f48b232de5d7f4658854" data-remote="true" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="qM43FtlWKrFA/pCLh3A8/ujDUqICL6f+K+/9RQXYUgzlh4SJbtwKxMGmlcKeuYzUedKpe3kV0KCyhTdaI9VGrw==" /></div><button class="btn-link js-clone-selector" data-protocol="ssh" type="submit">SSH</button></form>, or <form accept-charset="UTF-8" action="/users/set_protocol?protocol_selector=subversion&amp;protocol_type=push" class="inline-form js-clone-selector-form is-enabled" data-form-nonce="08775a3354e7c6e73c74f48b232de5d7f4658854" data-remote="true" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="ng9R0qKroD9dgCx+sy7tkb1sL+pHsmmenC1Ll6xKGmRVzpPXQNLjY3Pl0OBULUhtHUw7jMizsvqz2cCoH8f3bQ==" /></div><button class="btn-link js-clone-selector" data-protocol="subversion" type="submit">Subversion</button></form>.
    <a href="https://help.github.com/articles/which-remote-url-should-i-use" class="help tooltipped tooltipped-n" aria-label="Get help on which URL is right for you.">
      <span class="octicon octicon-question"></span>
    </a>
  </div>
    <a href="https://mac.github.com" class="btn btn-sm sidebar-button" title="Save timehop/wiki to your computer and use it in GitHub Desktop." aria-label="Save timehop/wiki to your computer and use it in GitHub Desktop.">
      <span class="octicon octicon-device-desktop"></span>
      Clone in Desktop
    </a>

                  <a href="/timehop/wiki/archive/master.zip"
                     class="btn btn-sm sidebar-button"
                     aria-label="Download the contents of timehop/wiki as a zip file"
                     title="Download the contents of timehop/wiki as a zip file"
                     rel="nofollow">
                    <span class="octicon octicon-cloud-download"></span>
                    Download ZIP
                  </a>
                </div>
          </div>
          <div id="js-repo-pjax-container" class="repository-content context-loader-container" data-pjax-container>

            

<a href="/timehop/wiki/blob/0bc5474bfcbaf0122d7f0ce0138d22a09996930a/README.md" class="hidden js-permalink-shortcut" data-hotkey="y">Permalink</a>

<!-- blob contrib key: blob_contributors:v21:3822e55dacae8728cd213ce8a01b1b7c -->

  <div class="file-navigation js-zeroclipboard-container">
    
<div class="select-menu js-menu-container js-select-menu left">
  <span class="btn btn-sm select-menu-button js-menu-target css-truncate" data-hotkey="w"
    data-ref="master"
    title="master"
    role="button" aria-label="Switch branches or tags" tabindex="0" aria-haspopup="true">
    <i>Branch:</i>
    <span class="js-select-button css-truncate-target">master</span>
  </span>

  <div class="select-menu-modal-holder js-menu-content js-navigation-container" data-pjax aria-hidden="true">

    <div class="select-menu-modal">
      <div class="select-menu-header">
        <span class="select-menu-title">Switch branches/tags</span>
        <span class="octicon octicon-x js-menu-close" role="button" aria-label="Close"></span>
      </div>

      <div class="select-menu-filters">
        <div class="select-menu-text-filter">
          <input type="text" aria-label="Find or create a branch…" id="context-commitish-filter-field" class="js-filterable-field js-navigation-enable" placeholder="Find or create a branch…">
        </div>
        <div class="select-menu-tabs">
          <ul>
            <li class="select-menu-tab">
              <a href="#" data-tab-filter="branches" data-filter-placeholder="Find or create a branch…" class="js-select-menu-tab" role="tab">Branches</a>
            </li>
            <li class="select-menu-tab">
              <a href="#" data-tab-filter="tags" data-filter-placeholder="Find a tag…" class="js-select-menu-tab" role="tab">Tags</a>
            </li>
          </ul>
        </div>
      </div>

      <div class="select-menu-list select-menu-tab-bucket js-select-menu-tab-bucket" data-tab-filter="branches" role="menu">

        <div data-filterable-for="context-commitish-filter-field" data-filterable-type="substring">


            <a class="select-menu-item js-navigation-item js-navigation-open "
               href="/timehop/wiki/blob/add-roadmap/README.md"
               data-name="add-roadmap"
               data-skip-pjax="true"
               rel="nofollow">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <span class="select-menu-item-text css-truncate-target" title="add-roadmap">
                add-roadmap
              </span>
            </a>
            <a class="select-menu-item js-navigation-item js-navigation-open "
               href="/timehop/wiki/blob/adds-devsetup/README.md"
               data-name="adds-devsetup"
               data-skip-pjax="true"
               rel="nofollow">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <span class="select-menu-item-text css-truncate-target" title="adds-devsetup">
                adds-devsetup
              </span>
            </a>
            <a class="select-menu-item js-navigation-item js-navigation-open "
               href="/timehop/wiki/blob/adds-postgis-install/README.md"
               data-name="adds-postgis-install"
               data-skip-pjax="true"
               rel="nofollow">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <span class="select-menu-item-text css-truncate-target" title="adds-postgis-install">
                adds-postgis-install
              </span>
            </a>
            <a class="select-menu-item js-navigation-item js-navigation-open "
               href="/timehop/wiki/blob/backend-pagerduty/README.md"
               data-name="backend-pagerduty"
               data-skip-pjax="true"
               rel="nofollow">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <span class="select-menu-item-text css-truncate-target" title="backend-pagerduty">
                backend-pagerduty
              </span>
            </a>
            <a class="select-menu-item js-navigation-item js-navigation-open "
               href="/timehop/wiki/blob/backend-testing/README.md"
               data-name="backend-testing"
               data-skip-pjax="true"
               rel="nofollow">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <span class="select-menu-item-text css-truncate-target" title="backend-testing">
                backend-testing
              </span>
            </a>
            <a class="select-menu-item js-navigation-item js-navigation-open selected"
               href="/timehop/wiki/blob/master/README.md"
               data-name="master"
               data-skip-pjax="true"
               rel="nofollow">
              <span class="select-menu-item-icon octicon octicon-check"></span>
              <span class="select-menu-item-text css-truncate-target" title="master">
                master
              </span>
            </a>
        </div>

          <form accept-charset="UTF-8" action="/timehop/wiki/branches" class="js-create-branch select-menu-item select-menu-new-item-form js-navigation-item js-new-item-form" data-form-nonce="08775a3354e7c6e73c74f48b232de5d7f4658854" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="RF3nenM6Vm/PsiRj93BY/2RZqtfEfQ3f9Jg7jdnp8K+aSrJnsP5Z4TSoeuB52jgbx6dSNjplWDkOUpLPQ6Ta/w==" /></div>
            <span class="octicon octicon-git-branch select-menu-item-icon"></span>
            <div class="select-menu-item-text">
              <span class="select-menu-item-heading">Create branch: <span class="js-new-item-name"></span></span>
              <span class="description">from ‘master’</span>
            </div>
            <input type="hidden" name="name" id="name" class="js-new-item-value">
            <input type="hidden" name="branch" id="branch" value="master">
            <input type="hidden" name="path" id="path" value="README.md">
</form>
      </div>

      <div class="select-menu-list select-menu-tab-bucket js-select-menu-tab-bucket" data-tab-filter="tags">
        <div data-filterable-for="context-commitish-filter-field" data-filterable-type="substring">


        </div>

        <div class="select-menu-no-results">Nothing to show</div>
      </div>

    </div>
  </div>
</div>

    <div class="btn-group right">
      <a href="/timehop/wiki/find/master"
            class="js-show-file-finder btn btn-sm empty-icon tooltipped tooltipped-nw"
            data-pjax
            data-hotkey="t"
            aria-label="Quickly jump between files">
        <span class="octicon octicon-list-unordered"></span>
      </a>
      <button aria-label="Copy file path to clipboard" class="js-zeroclipboard btn btn-sm zeroclipboard-button tooltipped tooltipped-s" data-copied-hint="Copied!" type="button"><span class="octicon octicon-clippy"></span></button>
    </div>

    <div class="breadcrumb js-zeroclipboard-target">
      <span class="repo-root js-repo-root"><span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/timehop/wiki" class="" data-branch="master" data-pjax="true" itemscope="url"><span itemprop="title">wiki</span></a></span></span><span class="separator">/</span><strong class="final-path">README.md</strong>
    </div>
  </div>


  <div class="commit file-history-tease">
    <div class="file-history-tease-header">
        <img alt="@kevin-cantwell" class="avatar" height="24" src="https://avatars0.githubusercontent.com/u/307864?v=3&amp;s=48" width="24" />
        <span class="author"><a href="/kevin-cantwell" rel="contributor">kevin-cantwell</a></span>
        <time datetime="2015-07-16T19:30:27Z" is="relative-time">Jul 16, 2015</time>
        <div class="commit-title">
            <a href="/timehop/wiki/commit/3e59cb53106f9c1b169a8e6c4605fb1fcf9c25f6" class="message" data-pjax="true" title="Specify clubhouse integration">Specify clubhouse integration</a>
        </div>
    </div>

    <div class="participation">
      <p class="quickstat">
        <a href="#blob_contributors_box" rel="facebox">
          <strong>1</strong>
           contributor
        </a>
      </p>
      
    </div>
    <div id="blob_contributors_box" style="display:none">
      <h2 class="facebox-header">Users who have contributed to this file</h2>
      <ul class="facebox-user-list">
          <li class="facebox-user-list-item">
            <img alt="@kevin-cantwell" height="24" src="https://avatars0.githubusercontent.com/u/307864?v=3&amp;s=48" width="24" />
            <a href="/kevin-cantwell">kevin-cantwell</a>
          </li>
      </ul>
    </div>
  </div>

<div class="file">
  <div class="file-header">
    <div class="file-actions">

      <div class="btn-group">
        <a href="/timehop/wiki/raw/master/README.md" class="btn btn-sm " id="raw-url">Raw</a>
          <a href="/timehop/wiki/blame/master/README.md" class="btn btn-sm js-update-url-with-hash">Blame</a>
        <a href="/timehop/wiki/commits/master/README.md" class="btn btn-sm " rel="nofollow">History</a>
      </div>

        <a class="octicon-btn tooltipped tooltipped-nw"
           href="https://mac.github.com"
           aria-label="Open this file in GitHub for Mac"
           data-ga-click="Repository, open with desktop, type:mac">
            <span class="octicon octicon-device-desktop"></span>
        </a>

            <form accept-charset="UTF-8" action="/timehop/wiki/edit/master/README.md" class="inline-form" data-form-nonce="08775a3354e7c6e73c74f48b232de5d7f4658854" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="QnBOsA9t9Hh6xUd5CY3HrRuwwEMI6LZfCOf54LCrnOe7Dp/Ab1zC299LmjbbbOzww3xe3Crrg89YeupysmENHA==" /></div>
              <button class="octicon-btn tooltipped tooltipped-n" type="submit" aria-label="Edit this file" data-hotkey="e" data-disable-with>
                <span class="octicon octicon-pencil"></span>
              </button>
</form>
          <form accept-charset="UTF-8" action="/timehop/wiki/delete/master/README.md" class="inline-form" data-form-nonce="08775a3354e7c6e73c74f48b232de5d7f4658854" method="post"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /><input name="authenticity_token" type="hidden" value="xQjLFJvMxhNUFegbHaQHDIGE7xyTO3+H/mLEcS7LFblLRL3S+Z391ZUfxumfnWRMI0RbAgRVQ/IBzVrNK9m0jQ==" /></div>
            <button class="octicon-btn octicon-btn-danger tooltipped tooltipped-n" type="submit" aria-label="Delete this file" data-disable-with>
              <span class="octicon octicon-trashcan"></span>
            </button>
</form>    </div>

    <div class="file-info">
        56 lines (35 sloc)
        <span class="file-info-divider"></span>
      3.749 kB
    </div>
  </div>
  
  <div id="readme" class="blob instapaper_body">
    <article class="markdown-body entry-content" itemprop="mainContentOfPage">
` + string(markdown) + `
    </article>
  </div>

</div>

<a href="#jump-to-line" rel="facebox[.linejump]" data-hotkey="l" style="display:none">Jump to Line</a>
<div id="jump-to-line" style="display:none">
  <form accept-charset="UTF-8" action="" class="js-jump-to-line-form" method="get"><div style="margin:0;padding:0;display:inline"><input name="utf8" type="hidden" value="&#x2713;" /></div>
    <input class="linejump-input js-jump-to-line-field" type="text" placeholder="Jump to line&hellip;" aria-label="Jump to line" autofocus>
    <button type="submit" class="btn">Go</button>
</form></div>

          </div>
        </div>
        <div class="modal-backdrop"></div>
      </div>
  </div>


    </div><!-- /.wrapper -->

      <div class="container">
  <div class="site-footer" role="contentinfo">
    <ul class="site-footer-links right">
        <li><a href="https://status.github.com/" data-ga-click="Footer, go to status, text:status">Status</a></li>
      <li><a href="https://developer.github.com" data-ga-click="Footer, go to api, text:api">API</a></li>
      <li><a href="https://training.github.com" data-ga-click="Footer, go to training, text:training">Training</a></li>
      <li><a href="https://shop.github.com" data-ga-click="Footer, go to shop, text:shop">Shop</a></li>
        <li><a href="https://github.com/blog" data-ga-click="Footer, go to blog, text:blog">Blog</a></li>
        <li><a href="https://github.com/about" data-ga-click="Footer, go to about, text:about">About</a></li>
        <li><a href="https://help.github.com" data-ga-click="Footer, go to help, text:help">Help</a></li>

    </ul>

    <a href="https://github.com" aria-label="Homepage">
      <span class="mega-octicon octicon-mark-github" title="GitHub"></span>
</a>
    <ul class="site-footer-links">
      <li>&copy; 2015 <span title="0.06342s from github-fe141-cp1-prd.iad.github.net">GitHub</span>, Inc.</li>
        <li><a href="https://github.com/site/terms" data-ga-click="Footer, go to terms, text:terms">Terms</a></li>
        <li><a href="https://github.com/site/privacy" data-ga-click="Footer, go to privacy, text:privacy">Privacy</a></li>
        <li><a href="https://github.com/security" data-ga-click="Footer, go to security, text:security">Security</a></li>
        <li><a href="https://github.com/contact" data-ga-click="Footer, go to contact, text:contact">Contact</a></li>
    </ul>
  </div>
</div>


    <div class="fullscreen-overlay js-fullscreen-overlay" id="fullscreen_overlay">
  <div class="fullscreen-container js-suggester-container">
    <div class="textarea-wrap">
      <textarea name="fullscreen-contents" id="fullscreen-contents" class="fullscreen-contents js-fullscreen-contents" placeholder="" aria-label=""></textarea>
      <div class="suggester-container">
        <div class="suggester fullscreen-suggester js-suggester js-navigation-container"></div>
      </div>
    </div>
  </div>
  <div class="fullscreen-sidebar">
    <a href="#" class="exit-fullscreen js-exit-fullscreen tooltipped tooltipped-w" aria-label="Exit Zen Mode">
      <span class="mega-octicon octicon-screen-normal"></span>
    </a>
    <a href="#" class="theme-switcher js-theme-switcher tooltipped tooltipped-w"
      aria-label="Switch themes">
      <span class="octicon octicon-color-mode"></span>
    </a>
  </div>
</div>



    
    
    

    <div id="ajax-error-message" class="flash flash-error">
      <span class="octicon octicon-alert"></span>
      <a href="#" class="octicon octicon-x flash-close js-ajax-error-dismiss" aria-label="Dismiss error"></a>
      Something went wrong with that request. Please try again.
    </div>


      <script crossorigin="anonymous" src="https://assets-cdn.github.com/assets/frameworks-60fa9d481f93b9638a55282fc13cd1e893f5da608855190c2259c5b35883105c.js"></script>
      <script async="async" crossorigin="anonymous" src="https://assets-cdn.github.com/assets/github/index-4864e54542a25a0f1dc884a414b7fee9b624d1717ce4a61500d1e23907a794ac.js"></script>
      
      
  </body>
</html>
`)
}
