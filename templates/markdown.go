package templates

import "html/template"

type GFMPage struct {
	MarkdownBody template.HTML
}

var GFM = `
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
            
          </div>
          <div class="container">
            <div class="repository-with-sidebar repo-container new-discussion-timeline ">
              <div id="js-repo-pjax-container" class="repository-content context-loader-container" data-pjax-container>
                <div class="file-navigation js-zeroclipboard-container">
                  <div class="breadcrumb js-zeroclipboard-target">
                    <span class="repo-root js-repo-root"><span itemscope="" itemtype="http://data-vocabulary.org/Breadcrumb"><a href="/" class="" data-branch="master" data-pjax="true" itemscope="url"><span itemprop="title">root</span></a></span></span><span class="separator">/</span><strong class="final-path">README.md</strong>
                  </div>
                </div>

                <div class="file">
                  <div class="file-header">
                    <div class="file-actions">
                    </div>
                    <div class="file-info">
                        &nbsp;
                    </div>
                  </div>
                  <div id="readme" class="blob instapaper_body">
                    <article class="markdown-body entry-content" itemprop="mainContentOfPage">
                      {{.MarkdownBody}}
                    </article>
                  </div>
                </div>
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
`