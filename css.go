package repoman

var GithubCSS = []byte(`
/*! normalize.css v3.0.3 | MIT License | github.com/necolas/normalize.css */

html {
    font-family: sans-serif;
    -webkit-text-size-adjust: 100%;
    -ms-text-size-adjust: 100%;
    text-size-adjust: 100%
}
body {
    margin: 0
}
article,
aside,
details,
figcaption,
figure,
footer,
header,
hgroup,
main,
menu,
nav,
section,
summary {
    display: block
}
audio,
canvas,
progress,
video {
    display: inline-block;
    vertical-align: baseline
}
audio:not([controls]) {
    display: none;
    height: 0
}
[hidden],
template {
    display: none
}
a {
    background-color: transparent
}
a:active,
a:hover {
    outline: 0
}
abbr[title] {
    border-bottom: 1px dotted
}
b,
strong {
    font-weight: bold
}
dfn {
    font-style: italic
}
h1 {
    font-size: 2em;
    margin: 0.67em 0
}
mark {
    background: #ff0;
    color: #000
}
small {
    font-size: 80%
}
sub,
sup {
    font-size: 75%;
    line-height: 0;
    position: relative;
    vertical-align: baseline
}
sup {
    top: -0.5em
}
sub {
    bottom: -0.25em
}
img {
    border: 0
}
svg:not(:root) {
    overflow: hidden
}
figure {
    margin: 1em 40px
}
hr {
    box-sizing: content-box;
    height: 0
}
pre {
    overflow: auto
}
code,
kbd,
pre,
samp {
    font-family: monospace, monospace;
    font-size: 1em
}
button,
input,
optgroup,
select,
textarea {
    color: inherit;
    font: inherit;
    margin: 0
}
button {
    overflow: visible
}
button,
select {
    text-transform: none
}
button,
html input[type="button"],
input[type="reset"],
input[type="submit"] {
    -webkit-appearance: button;
    cursor: pointer
}
button[disabled],
html input[disabled] {
    cursor: default
}
button::-moz-focus-inner,
input::-moz-focus-inner {
    border: 0;
    padding: 0
}
input {
    line-height: normal
}
input[type="checkbox"],
input[type="radio"] {
    box-sizing: border-box;
    padding: 0
}
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
    height: auto
}
input[type="search"] {
    -webkit-appearance: textfield;
    box-sizing: content-box
}
input[type="search"]::-webkit-search-cancel-button,
input[type="search"]::-webkit-search-decoration {
    -webkit-appearance: none
}
fieldset {
    border: 1px solid #c0c0c0;
    margin: 0 2px;
    padding: 0.35em 0.625em 0.75em
}
legend {
    border: 0;
    padding: 0
}
textarea {
    overflow: auto
}
optgroup {
    font-weight: bold
}
table {
    border-collapse: collapse;
    border-spacing: 0
}
td,
th {
    padding: 0
}
* {
    box-sizing: border-box
}
input,
select,
textarea,
button {
    font: 13px/1.4 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol"
}
body {
    font: 13px/1.4 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol";
    color: #333;
    background-color: #fff
}
a {
    color: #4078c0;
    text-decoration: none
}
a:hover,
a:active {
    text-decoration: underline
}
hr,
.rule {
    height: 0;
    margin: 15px 0;
    overflow: hidden;
    background: transparent;
    border: 0;
    border-bottom: 1px solid #ddd
}
hr:before,
.rule:before {
    display: table;
    content: ""
}
hr:after,
.rule:after {
    display: table;
    clear: both;
    content: ""
}
h1,
h2,
h3,
h4,
h5,
h6 {
    margin-top: 15px;
    margin-bottom: 15px;
    line-height: 1.1
}
h1 {
    font-size: 30px
}
h2 {
    font-size: 21px
}
h3 {
    font-size: 16px
}
h4 {
    font-size: 14px
}
h5 {
    font-size: 12px
}
h6 {
    font-size: 11px
}
small {
    font-size: 90%
}
blockquote {
    margin: 0
}
.lead {
    margin-bottom: 30px;
    font-size: 20px;
    font-weight: 300;
    color: #555
}
.text-muted {
    color: #767676
}
.text-danger {
    color: #bd2c00
}
.text-emphasized {
    font-weight: bold;
    color: #333
}
ul,
ol {
    padding: 0;
    margin-top: 0;
    margin-bottom: 0
}
ol ol,
ul ol {
    list-style-type: lower-roman
}
ul ul ol,
ul ol ol,
ol ul ol,
ol ol ol {
    list-style-type: lower-alpha
}
dd {
    margin-left: 0
}
tt,
code {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px
}
pre {
    margin-top: 0;
    margin-bottom: 0;
    font: 12px Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.container {
    width: 980px;
    margin-right: auto;
    margin-left: auto
}
.container:before {
    display: table;
    content: ""
}
.container:after {
    display: table;
    clear: both;
    content: ""
}
.columns {
    margin-right: -10px;
    margin-left: -10px
}
.columns:before {
    display: table;
    content: ""
}
.columns:after {
    display: table;
    clear: both;
    content: ""
}
.column {
    float: left;
    padding-right: 10px;
    padding-left: 10px
}
.one-third {
    width: 33.333333%
}
.two-thirds {
    width: 66.666667%
}
.one-fourth {
    width: 25%
}
.one-half {
    width: 50%
}
.three-fourths {
    width: 75%
}
.one-fifth {
    width: 20%
}
.four-fifths {
    width: 80%
}
.single-column {
    padding-right: 10px;
    padding-left: 10px
}
.table-column {
    display: table-cell;
    width: 1%;
    padding-right: 10px;
    padding-left: 10px;
    vertical-align: top
}
fieldset {
    padding: 0;
    margin: 0;
    border: 0
}
label {
    font-size: 13px;
    font-weight: bold
}
.form-control,
input[type="text"],
input[type="password"],
input[type="email"],
input[type="number"],
input[type="tel"],
input[type="url"],
textarea {
    min-height: 34px;
    padding: 7px 8px;
    font-size: 13px;
    color: #333;
    vertical-align: middle;
    background-color: #fff;
    background-repeat: no-repeat;
    background-position: right 8px center;
    border: 1px solid #ccc;
    border-radius: 3px;
    outline: none;
    box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.075)
}
.form-control.focus,
.form-control:focus,
input[type="text"].focus,
input[type="text"]:focus,
.focused .drag-and-drop,
input[type="password"].focus,
input[type="password"]:focus,
input[type="email"].focus,
input[type="email"]:focus,
input[type="number"].focus,
input[type="number"]:focus,
input[type="tel"].focus,
input[type="tel"]:focus,
input[type="url"].focus,
input[type="url"]:focus,
textarea.focus,
textarea:focus {
    border-color: #51a7e8;
    box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.075), 0 0 5px rgba(81, 167, 232, 0.5)
}
input.input-contrast,
.input-contrast {
    background-color: #fafafa
}
input.input-contrast:focus,
.input-contrast:focus {
    background-color: #fff
}
::-webkit-input-placeholder {
    color: #aaa
}
::-moz-placeholder {
    color: #aaa
}
:-ms-input-placeholder {
    color: #aaa
}
::placeholder {
    color: #aaa
}
input.input-mini {
    min-height: 26px;
    padding-top: 4px;
    padding-bottom: 4px;
    font-size: 12px
}
input.input-large {
    padding: 6px 10px;
    font-size: 16px
}
.input-block {
    display: block;
    width: 100%
}
.input-monospace {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
dl.form {
    margin: 15px 0
}
dl.form input[type="text"],
dl.form input[type="password"],
dl.form input[type="email"],
dl.form input[type="url"],
dl.form textarea {
    background-color: #fafafa
}
dl.form input[type="text"]:focus,
dl.form .focused .drag-and-drop,
.focused dl.form .drag-and-drop,
dl.form input[type="password"]:focus,
dl.form input[type="email"]:focus,
dl.form input[type="url"]:focus,
dl.form textarea:focus {
    background-color: #fff
}
dl.form>dt {
    margin: 0 0 6px
}
dl.form>dt label {
    position: relative
}
dl.form.flattened>dt {
    float: left;
    margin: 0;
    line-height: 32px
}
dl.form.flattened>dd {
    line-height: 32px
}
dl.form>dd input[type="text"],
dl.form>dd input[type="password"],
dl.form>dd input[type="email"],
dl.form>dd input[type="url"] {
    width: 440px;
    max-width: 100%;
    margin-right: 5px;
    background-position-x: 98%
}
dl.form>dd input.shorter {
    width: 130px
}
dl.form>dd input.short {
    width: 250px
}
dl.form>dd input.long {
    width: 100%
}
dl.form>dd textarea {
    width: 100%;
    height: 200px;
    min-height: 200px
}
dl.form>dd textarea.short {
    height: 50px;
    min-height: 50px
}
dl.form>dd h4 {
    margin: 4px 0 0
}
dl.form>dd h4.is-error {
    color: #bd2c00
}
dl.form>dd h4.is-success {
    color: #6cc644
}
dl.form>dd h4+p.note {
    margin-top: 0
}
dl.form.required>dt>label:after {
    padding-left: 5px;
    color: #9f1006;
    content: "*"
}
dl.form .success,
dl.form .error,
dl.form .indicator {
    display: none;
    font-size: 12px;
    font-weight: bold
}
dl.form.loading {
    opacity: 0.5
}
dl.form.loading .indicator {
    display: inline
}
dl.form.loading .spinner {
    display: inline-block;
    vertical-align: middle
}
dl.form.successful .success {
    display: inline;
    color: #390
}
dl.form.errored>dt label {
    color: #900
}
dl.form.errored .error {
    display: inline;
    color: #900
}
dl.form.errored dd.error,
dl.form.errored dd.warning {
    display: inline-block;
    padding: 5px;
    font-size: 11px;
    color: #494620;
    background: #f7ea57;
    border: 1px solid #c0b536;
    border-top-color: #fff;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
dl.form.warn .warning {
    display: inline;
    color: #900
}
dl.form.warn dd.warning {
    display: inline-block;
    padding: 5px;
    font-size: 11px;
    color: #494620;
    background: #f7ea57;
    border: 1px solid #c0b536;
    border-top-color: #fff;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
dl.form .form-note {
    display: inline-block;
    padding: 5px;
    margin-top: -1px;
    font-size: 11px;
    color: #494620;
    background: #f7ea57;
    border: 1px solid #c0b536;
    border-top-color: #fff;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.note {
    min-height: 17px;
    margin: 4px 0 2px;
    font-size: 12px;
    color: #767676
}
.note .spinner {
    margin-right: 3px;
    vertical-align: middle
}
.form-checkbox {
    padding-left: 20px;
    margin: 15px 0;
    vertical-align: middle
}
.form-checkbox label em.highlight {
    position: relative;
    left: -4px;
    padding: 2px 4px;
    font-style: normal;
    background: #fffbdc;
    border-radius: 3px
}
.form-checkbox input[type=checkbox],
.form-checkbox input[type=radio] {
    float: left;
    margin: 2px 0 0 -20px;
    vertical-align: middle
}
.form-checkbox .note {
    display: block;
    margin: 0;
    font-size: 12px;
    font-weight: normal;
    color: #666
}
.hfields {
    margin: 15px 0
}
.hfields:before {
    display: table;
    content: ""
}
.hfields:after {
    display: table;
    clear: both;
    content: ""
}
.hfields dl.form {
    float: left;
    margin: 0 30px 0 0
}
.hfields dl.form>dt label {
    display: inline-block;
    margin: 5px 0 0;
    color: #666
}
.hfields dl.form>dt img {
    position: relative;
    top: -2px
}
.hfields .btn {
    float: left;
    margin: 28px 25px 0 -20px
}
.hfields select {
    margin-top: 5px
}
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
    margin: 0;
    -webkit-appearance: none
}
.input-group {
    display: table
}
.input-group input {
    position: relative;
    width: 100%
}
.input-group input:focus {
    z-index: 2
}
.input-group input[type="text"]+.btn {
    margin-left: 0
}
.input-group.inline {
    display: inline-table
}
.input-group input,
.input-group-button {
    display: table-cell
}
.input-group-button {
    width: 1%;
    vertical-align: middle
}
.input-group input:first-child,
.input-group-button:first-child .btn {
    border-top-right-radius: 0;
    border-bottom-right-radius: 0
}
.input-group-button:first-child .btn {
    margin-right: -1px
}
.input-group input:last-child,
.input-group-button:last-child .btn {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0
}
.input-group-button:last-child .btn {
    margin-left: -1px
}
.form-actions:before {
    display: table;
    content: ""
}
.form-actions:after {
    display: table;
    clear: both;
    content: ""
}
.form-actions .btn {
    float: right
}
.form-actions .btn+.btn {
    margin-right: 5px
}
.form-warning {
    padding: 8px 10px;
    margin: 10px 0;
    font-size: 14px;
    color: #333;
    background: #ffffe2;
    border: 1px solid #e7e4c2;
    border-radius: 4px
}
.form-warning p {
    margin: 0;
    line-height: 1.5
}
.form-warning strong {
    color: #000
}
.form-warning a {
    font-weight: bold
}
.status-indicator {
    font: normal normal 16px/1 "octicons";
    display: inline-block;
    text-decoration: none;
    -webkit-font-smoothing: antialiased;
    margin-left: 5px
}
.status-indicator-success:before {
    color: #6cc644;
    content: "\f03a"
}
.status-indicator-failed:before {
    color: #bd2c00;
    content: "\f02d"
}
.flash-messages {
    margin-top: 15px;
    margin-bottom: 15px
}
.flash,
.flash-global {
    position: relative;
    font-size: 14px;
    line-height: 1.6;
    color: #246;
    background-color: #e2eef9;
    border: solid 1px #bac6d3
}
.flash.flash-warn,
.flash-global.flash-warn {
    color: #4c4a42;
    background-color: #fff9ea;
    border-color: #dfd8c2
}
.flash.flash-error,
.flash-global.flash-error {
    color: #911;
    background-color: #fcdede;
    border-color: #d2b2b2
}
.flash .flash-close,
.flash-global .flash-close {
    float: right;
    padding: 17px;
    margin-top: -15px;
    margin-right: -15px;
    margin-left: 20px;
    color: inherit;
    text-decoration: none;
    cursor: pointer;
    opacity: 0.6
}
.flash .flash-close:hover,
.flash-global .flash-close:hover {
    opacity: 1
}
.flash p:last-child,
.flash-global p:last-child {
    margin-bottom: 0
}
.flash .flash-action,
.flash-global .flash-action {
    float: right;
    margin-top: -4px;
    margin-left: 20px
}
.flash a,
.flash-global a {
    font-weight: bold
}
.flash {
    padding: 15px;
    border-radius: 3px
}
.flash+.flash {
    margin-top: 5px
}
.flash-with-icon {
    padding-left: 40px
}
.flash-with-icon>.octicon {
    float: left;
    margin-top: 3px;
    margin-left: -25px
}
.flash-global {
    padding: 10px;
    margin-top: -1px;
    border-width: 1px 0
}
.flash-global h2,
.flash-global p {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 14px;
    line-height: 1.4
}
.flash-global .flash-action {
    margin-top: 5px
}
.flash-title {
    margin-top: 0;
    margin-bottom: 5px
}
.avatar {
    display: inline-block;
    overflow: hidden;
    line-height: 1;
    vertical-align: middle;
    border-radius: 3px
}
.avatar-small {
    border-radius: 2px
}
.avatar-link {
    float: left;
    line-height: 1
}
.avatar-group-item {
    display: inline-block;
    margin-bottom: 3px
}
.avatar-parent-child {
    position: relative
}
.avatar-child {
    position: absolute;
    right: -15%;
    bottom: -9%;
    border-radius: 2px;
    box-shadow: -2px -2px 0 rgba(255, 255, 255, 0.8)
}
.blankslate {
    position: relative;
    padding: 30px;
    text-align: center;
    background-color: #fafafa;
    border: 1px solid #e5e5e5;
    border-radius: 3px;
    box-shadow: inset 0 0 10px rgba(0, 0, 0, 0.05)
}
.blankslate.clean-background {
    background: none;
    border: 0;
    box-shadow: none
}
.blankslate.capped {
    border-radius: 0 0 3px 3px
}
.blankslate.spacious {
    padding: 100px 60px 120px
}
.blankslate.has-fixed-width {
    width: 485px;
    margin: 0 auto
}
.blankslate.large-format h3 {
    margin: 0.75em 0;
    font-size: 20px
}
.blankslate.large-format p {
    font-size: 16px
}
.blankslate.large-format p.has-fixed-width {
    width: 540px;
    margin: 0 auto;
    text-align: left
}
.blankslate.large-format .mega-octicon {
    width: 40px;
    height: 40px;
    font-size: 40px;
    color: #aaa
}
.blankslate.large-format .octicon-inbox {
    font-size: 48px;
    line-height: 40px
}
.blankslate code {
    padding: 2px 5px 3px;
    font-size: 14px;
    background: #fff;
    border: 1px solid #eee;
    border-radius: 3px
}
.blankslate>.mega-octicon {
    color: #aaa
}
.blankslate .mega-octicon+.mega-octicon {
    margin-left: 10px
}
.tabnav+.blankslate {
    margin-top: 20px
}
.blankslate .context-loader.large-format-loader {
    padding-top: 50px
}
.counter {
    display: inline-block;
    padding: 2px 5px;
    font-size: 11px;
    font-weight: bold;
    line-height: 1;
    color: #666;
    background-color: #eee;
    border-radius: 20px
}
.btn {
    position: relative;
    display: inline-block;
    padding: 6px 12px;
    font-size: 13px;
    font-weight: bold;
    line-height: 20px;
    color: #333;
    white-space: nowrap;
    vertical-align: middle;
    cursor: pointer;
    background-color: #eee;
    background-image: -webkit-linear-gradient(#fcfcfc, #eee);
    background-image: linear-gradient(#fcfcfc, #eee);
    border: 1px solid #d5d5d5;
    border-radius: 3px;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    -webkit-appearance: none
}
.btn i {
    font-style: normal;
    font-weight: 500;
    opacity: 0.6
}
.btn .octicon {
    vertical-align: text-top
}
.btn .counter {
    text-shadow: none;
    background-color: #e5e5e5
}
.btn:focus {
    text-decoration: none;
    border-color: #51a7e8;
    outline: none;
    box-shadow: 0 0 5px rgba(81, 167, 232, 0.5)
}
.btn:focus:hover,
.btn.selected:focus {
    border-color: #51a7e8
}
.btn:hover,
.btn:active,
.btn.zeroclipboard-is-hover,
.btn.zeroclipboard-is-active {
    text-decoration: none;
    background-color: #ddd;
    background-image: -webkit-linear-gradient(#eee, #ddd);
    background-image: linear-gradient(#eee, #ddd);
    border-color: #ccc
}
.btn:active,
.btn.selected,
.btn.zeroclipboard-is-active {
    background-color: #dcdcdc;
    background-image: none;
    border-color: #b5b5b5;
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.15)
}
.btn.selected:hover {
    background-color: #cfcfcf
}
.btn:disabled,
.btn:disabled:hover,
.btn.disabled,
.btn.disabled:hover {
    color: rgba(102, 102, 102, 0.5);
    cursor: default;
    background-color: rgba(229, 229, 229, 0.5);
    background-image: none;
    border-color: rgba(197, 197, 197, 0.5);
    box-shadow: none
}
.btn-primary {
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.15);
    background-color: #60b044;
    background-image: -webkit-linear-gradient(#8add6d, #60b044);
    background-image: linear-gradient(#8add6d, #60b044);
    border-color: #5ca941
}
.btn-primary .counter {
    color: #60b044;
    background-color: #fff
}
.btn-primary:hover {
    color: #fff;
    background-color: #569e3d;
    background-image: -webkit-linear-gradient(#79d858, #569e3d);
    background-image: linear-gradient(#79d858, #569e3d);
    border-color: #4a993e
}
.btn-primary:active,
.btn-primary.selected {
    text-shadow: 0 1px 0 rgba(0, 0, 0, 0.15);
    background-color: #569e3d;
    background-image: none;
    border-color: #418737
}
.btn-primary.selected:hover {
    background-color: #4c8b36
}
.btn-primary:disabled,
.btn-primary:disabled:hover,
.btn-primary.disabled,
.btn-primary.disabled:hover {
    color: #fefefe;
    background-color: #add39f;
    background-image: -webkit-linear-gradient(#c3ecb4, #add39f);
    background-image: linear-gradient(#c3ecb4, #add39f);
    border-color: #b9dcac #b9dcac #a7c89b
}
.btn-danger {
    color: #900
}
.btn-danger:hover {
    color: #fff;
    background-color: #b33630;
    background-image: -webkit-linear-gradient(#dc5f59, #b33630);
    background-image: linear-gradient(#dc5f59, #b33630);
    border-color: #cd504a
}
.btn-danger:active,
.btn-danger.selected {
    color: #fff;
    background-color: #b33630;
    background-image: none;
    border-color: #9f312c
}
.btn-danger.selected:hover {
    background-color: #9f302b
}
.btn-danger:disabled,
.btn-danger:disabled:hover,
.btn-danger.disabled,
.btn-danger.disabled:hover {
    color: #cb7f7f;
    background-color: #efefef;
    background-image: -webkit-linear-gradient(#fefefe, #efefef);
    background-image: linear-gradient(#fefefe, #efefef);
    border-color: #e1e1e1
}
.btn-danger:hover .counter,
.btn-danger:active .counter,
.btn-danger.selected .counter {
    color: #b33630;
    background-color: #fff
}
.btn-outline {
    color: #4078c0;
    background-color: #fff;
    background-image: none;
    border: 1px solid #e5e5e5
}
.btn-outline .counter {
    background-color: #eee
}
.btn-outline:hover,
.btn-outline:active,
.btn-outline.selected,
.btn-outline.zeroclipboard-is-hover,
.btn-outline.zeroclipboard-is-active {
    color: #fff;
    background-color: #4078c0;
    background-image: none;
    border-color: #4078c0
}
.btn-outline:hover .counter,
.btn-outline:active .counter,
.btn-outline.selected .counter,
.btn-outline.zeroclipboard-is-hover .counter,
.btn-outline.zeroclipboard-is-active .counter {
    color: #4078c0;
    background-color: #fff
}
.btn-outline.selected:hover {
    background-color: #396cad
}
.btn-outline:disabled,
.btn-outline:disabled:hover,
.btn-outline.disabled,
.btn-outline.disabled:hover {
    color: #767676;
    background-color: #fff;
    background-image: none;
    border-color: #e5e5e5
}
.btn-with-count {
    float: left;
    border-top-right-radius: 0;
    border-bottom-right-radius: 0
}
.btn-sm {
    padding: 2px 10px
}
.hidden-text-expander {
    display: block
}
.hidden-text-expander.inline {
    position: relative;
    top: -1px;
    display: inline-block;
    margin-left: 5px;
    line-height: 0
}
.hidden-text-expander a {
    display: inline-block;
    height: 12px;
    padding: 0 5px;
    font-size: 12px;
    font-weight: bold;
    line-height: 6px;
    color: #555;
    text-decoration: none;
    vertical-align: middle;
    background: #ddd;
    border-radius: 1px
}
.hidden-text-expander a:hover {
    text-decoration: none;
    background-color: #ccc
}
.hidden-text-expander a:active {
    color: #fff;
    background-color: #4183c4
}
.social-count {
    float: left;
    padding: 2px 7px;
    font-size: 11px;
    font-weight: bold;
    line-height: 20px;
    color: #333;
    vertical-align: middle;
    background-color: #fff;
    border: 1px solid #ddd;
    border-left: 0;
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px
}
.social-count:hover,
.social-count:active {
    text-decoration: none
}
.social-count:hover {
    color: #4078c0;
    cursor: pointer
}
.btn-block {
    display: block;
    width: 100%;
    text-align: center
}
.btn-group {
    display: inline-block;
    vertical-align: middle
}
.btn-group:before {
    display: table;
    content: ""
}
.btn-group:after {
    display: table;
    clear: both;
    content: ""
}
.btn-group .btn {
    position: relative;
    float: left
}
.btn-group .btn:not(:first-child):not(:last-child) {
    border-radius: 0
}
.btn-group .btn:first-child {
    border-top-right-radius: 0;
    border-bottom-right-radius: 0
}
.btn-group .btn:last-child {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0
}
.btn-group .btn:hover,
.btn-group .btn:active,
.btn-group .btn.selected {
    z-index: 2
}
.btn-group .btn:focus {
    z-index: 3
}
.btn-group .btn+.btn {
    margin-left: -1px
}
.btn-group .btn+.button_to,
.btn-group .button_to+.btn,
.btn-group .button_to+.button_to {
    margin-left: -1px
}
.btn-group .button_to {
    float: left
}
.btn-group .button_to .btn {
    border-radius: 0
}
.btn-group .button_to:first-child .btn {
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px
}
.btn-group .button_to:last-child .btn {
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px
}
.btn-group+.btn-group,
.btn-group+.btn {
    margin-left: 5px
}
.btn-link {
    display: inline-block;
    padding: 0;
    font-size: inherit;
    color: #4078c0;
    white-space: nowrap;
    cursor: pointer;
    background-color: transparent;
    border: 0;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    -webkit-appearance: none
}
.btn-link:hover,
.btn-link:focus {
    text-decoration: underline
}
.btn-link:focus {
    outline: none
}
.menu {
    margin-bottom: 15px;
    list-style: none;
    background-color: #fff;
    border: 1px solid #d8d8d8;
    border-radius: 3px
}
.menu-item {
    position: relative;
    display: block;
    padding: 8px 10px;
    text-shadow: 0 1px 0 #fff;
    border-bottom: 1px solid #eee
}
.menu-item:first-child {
    border-top: 0;
    border-top-right-radius: 2px;
    border-top-left-radius: 2px
}
.menu-item:first-child:before {
    border-top-left-radius: 2px
}
.menu-item:last-child {
    border-bottom: 0;
    border-bottom-right-radius: 2px;
    border-bottom-left-radius: 2px
}
.menu-item:last-child:before {
    border-bottom-left-radius: 2px
}
.menu-item:hover {
    text-decoration: none;
    background-color: #f9f9f9
}
.menu-item.selected {
    font-weight: bold;
    color: #222;
    cursor: default;
    background-color: #fff
}
.menu-item.selected:before {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    width: 2px;
    content: "";
    background-color: #d26911
}
.menu-item .octicon {
    margin-right: 5px;
    width: 16px;
    color: #333;
    text-align: center
}
.menu-item .counter {
    float: right;
    margin-left: 5px
}
.menu-item .menu-warning {
    float: right;
    color: #d26911
}
.menu-item .avatar {
    float: left;
    margin-right: 5px
}
.menu-item.alert .counter {
    color: #bd2c00
}
.menu-heading {
    display: block;
    padding: 8px 10px;
    margin-top: 0;
    margin-bottom: 0;
    font-size: 13px;
    font-weight: bold;
    line-height: 20px;
    color: #555;
    background-color: #f7f7f7;
    border-bottom: 1px solid #eee
}
.menu-heading:hover {
    text-decoration: none
}
.menu-heading:first-child {
    border-top-right-radius: 2px;
    border-top-left-radius: 2px
}
.menu-heading:last-child {
    border-bottom-right-radius: 2px;
    border-bottom-left-radius: 2px;
    border-bottom: 0
}
.tabnav {
    margin-top: 0;
    margin-bottom: 15px;
    border-bottom: 1px solid #ddd
}
.tabnav .counter {
    margin-left: 5px
}
.tabnav-tabs {
    margin-bottom: -1px
}
.tabnav-tab {
    display: inline-block;
    padding: 8px 12px;
    font-size: 14px;
    line-height: 20px;
    color: #666;
    text-decoration: none;
    border: 1px solid transparent;
    border-bottom: 0
}
.tabnav-tab.selected {
    color: #333;
    background-color: #fff;
    border-color: #ddd;
    border-radius: 3px 3px 0 0
}
.tabnav-tab:hover {
    text-decoration: none
}
.tabnav-extra {
    display: inline-block;
    padding-top: 10px;
    margin-left: 10px;
    font-size: 12px;
    color: #666
}
.tabnav-extra>.octicon {
    margin-right: 2px
}
a.tabnav-extra:hover {
    color: #4078c0;
    text-decoration: none
}
.tabnav-btn {
    margin-left: 10px
}
.filter-list {
    list-style-type: none
}
.filter-list.small .filter-item {
    padding: 4px 10px;
    margin: 0 0 2px;
    font-size: 12px
}
.filter-list.pjax-active .filter-item {
    color: #767676;
    background-color: transparent
}
.filter-list.pjax-active .filter-item.pjax-active {
    color: #fff;
    background-color: #4078c0
}
.filter-item {
    position: relative;
    display: block;
    padding: 8px 10px;
    margin-bottom: 5px;
    overflow: hidden;
    font-size: 14px;
    color: #767676;
    text-decoration: none;
    text-overflow: ellipsis;
    white-space: nowrap;
    cursor: pointer;
    border-radius: 3px
}
.filter-item:hover {
    text-decoration: none;
    background-color: #eee
}
.filter-item.selected {
    color: #fff;
    background-color: #4078c0
}
.filter-item.selected .octicon-remove-close {
    float: right;
    opacity: 0.8
}
.filter-item .count {
    float: right;
    font-weight: bold
}
.filter-item .bar {
    position: absolute;
    top: 2px;
    right: 0;
    bottom: 2px;
    z-index: -1;
    display: inline-block;
    background-color: #f1f1f1
}
.state {
    display: inline-block;
    padding: 4px 8px;
    font-weight: bold;
    line-height: 20px;
    color: #fff;
    text-align: center;
    border-radius: 3px;
    background-color: #999
}
.state-open,
.state-proposed,
.state-reopened {
    background-color: #6cc644
}
.state-merged {
    background-color: #6e5494
}
.state-closed {
    background-color: #bd2c00
}
.state-renamed {
    background-color: #fffa5d
}
.tooltipped {
    position: relative
}
.tooltipped:after {
    position: absolute;
    z-index: 1000000;
    display: none;
    padding: 5px 8px;
    font: normal normal 11px/1.5 Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol";
    color: #fff;
    text-align: center;
    text-decoration: none;
    text-shadow: none;
    text-transform: none;
    letter-spacing: normal;
    word-wrap: break-word;
    white-space: pre;
    pointer-events: none;
    content: attr(aria-label);
    background: rgba(0, 0, 0, 0.8);
    border-radius: 3px;
    -webkit-font-smoothing: subpixel-antialiased
}
.tooltipped:before {
    position: absolute;
    z-index: 1000001;
    display: none;
    width: 0;
    height: 0;
    color: rgba(0, 0, 0, 0.8);
    pointer-events: none;
    content: "";
    border: 5px solid transparent
}
.tooltipped:hover:before,
.tooltipped:hover:after,
.tooltipped:active:before,
.tooltipped:active:after,
.tooltipped:focus:before,
.tooltipped:focus:after {
    display: inline-block;
    text-decoration: none
}
.tooltipped-multiline:hover:after,
.tooltipped-multiline:active:after,
.tooltipped-multiline:focus:after {
    display: table-cell
}
.tooltipped-s:after,
.tooltipped-se:after,
.tooltipped-sw:after {
    top: 100%;
    right: 50%;
    margin-top: 5px
}
.tooltipped-s:before,
.tooltipped-se:before,
.tooltipped-sw:before {
    top: auto;
    right: 50%;
    bottom: -5px;
    margin-right: -5px;
    border-bottom-color: rgba(0, 0, 0, 0.8)
}
.tooltipped-se:after {
    right: auto;
    left: 50%;
    margin-left: -15px
}
.tooltipped-sw:after {
    margin-right: -15px
}
.tooltipped-n:after,
.tooltipped-ne:after,
.tooltipped-nw:after {
    right: 50%;
    bottom: 100%;
    margin-bottom: 5px
}
.tooltipped-n:before,
.tooltipped-ne:before,
.tooltipped-nw:before {
    top: -5px;
    right: 50%;
    bottom: auto;
    margin-right: -5px;
    border-top-color: rgba(0, 0, 0, 0.8)
}
.tooltipped-ne:after {
    right: auto;
    left: 50%;
    margin-left: -15px
}
.tooltipped-nw:after {
    margin-right: -15px
}
.tooltipped-s:after,
.tooltipped-n:after {
    -webkit-transform: translateX(50%);
    -ms-transform: translateX(50%);
    transform: translateX(50%)
}
.tooltipped-w:after {
    right: 100%;
    bottom: 50%;
    margin-right: 5px;
    -webkit-transform: translateY(50%);
    -ms-transform: translateY(50%);
    transform: translateY(50%)
}
.tooltipped-w:before {
    top: 50%;
    bottom: 50%;
    left: -5px;
    margin-top: -5px;
    border-left-color: rgba(0, 0, 0, 0.8)
}
.tooltipped-e:after {
    bottom: 50%;
    left: 100%;
    margin-left: 5px;
    -webkit-transform: translateY(50%);
    -ms-transform: translateY(50%);
    transform: translateY(50%)
}
.tooltipped-e:before {
    top: 50%;
    right: -5px;
    bottom: 50%;
    margin-top: -5px;
    border-right-color: rgba(0, 0, 0, 0.8)
}
.tooltipped-multiline:after {
    width: -webkit-max-content;
    width: -moz-max-content;
    width: max-content;
    max-width: 250px;
    word-break: break-word;
    word-wrap: normal;
    white-space: pre-line;
    border-collapse: separate
}
.tooltipped-multiline.tooltipped-s:after,
.tooltipped-multiline.tooltipped-n:after {
    right: auto;
    left: 50%;
    -webkit-transform: translateX(-50%);
    -ms-transform: translateX(-50%);
    transform: translateX(-50%)
}
.tooltipped-multiline.tooltipped-w:after,
.tooltipped-multiline.tooltipped-e:after {
    right: 100%
}
@media screen and (min-width: 0\0) {
    .tooltipped-multiline:after {
        width: 250px
    }
}
.tooltipped-sticky:before,
.tooltipped-sticky:after {
    display: inline-block
}
.tooltipped-sticky.tooltipped-multiline:after {
    display: table-cell
}
.fullscreen-overlay-enabled.dark-theme .tooltipped:after {
    color: #000;
    background: rgba(255, 255, 255, 0.8)
}
.fullscreen-overlay-enabled.dark-theme .tooltipped .tooltipped-s:before,
.fullscreen-overlay-enabled.dark-theme .tooltipped .tooltipped-se:before,
.fullscreen-overlay-enabled.dark-theme .tooltipped .tooltipped-sw:before {
    border-bottom-color: rgba(255, 255, 255, 0.8)
}
.fullscreen-overlay-enabled.dark-theme .tooltipped.tooltipped-n:before,
.fullscreen-overlay-enabled.dark-theme .tooltipped.tooltipped-ne:before,
.fullscreen-overlay-enabled.dark-theme .tooltipped.tooltipped-nw:before {
    border-top-color: rgba(255, 255, 255, 0.8)
}
.fullscreen-overlay-enabled.dark-theme .tooltipped.tooltipped-e:before {
    border-right-color: rgba(255, 255, 255, 0.8)
}
.fullscreen-overlay-enabled.dark-theme .tooltipped.tooltipped-w:before {
    border-left-color: rgba(255, 255, 255, 0.8)
}
.flex-table {
    display: table
}
.flex-table-item {
    display: table-cell;
    width: 1%;
    white-space: nowrap;
    vertical-align: middle
}
.flex-table-item-primary {
    width: 99%
}
.css-truncate.css-truncate-target,
.css-truncate .css-truncate-target {
    display: inline-block;
    max-width: 125px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    vertical-align: top
}
.css-truncate.expandable.zeroclipboard-is-hover .css-truncate-target,
.css-truncate.expandable.zeroclipboard-is-hover.css-truncate-target,
.css-truncate.expandable:hover .css-truncate-target,
.css-truncate.expandable:hover.css-truncate-target {
    max-width: 10000px !important
}
.sunken-menu {
    position: relative;
    padding-top: 15px;
    padding-bottom: 15px;
    background-image: -webkit-linear-gradient(left, #f6f6f6 0%, #fff 8px);
    background-image: linear-gradient(to right, #f6f6f6 0%, #fff 8px);
    box-shadow: inset 1px 0 0 #eee
}
.sunken-menu:before,
.sunken-menu:after {
    position: absolute;
    left: 0;
    width: 100%;
    height: 15px;
    content: "";
    background-color: transparent
}
.sunken-menu:before {
    top: 0;
    background-color: transparent;
    background-image: -webkit-linear-gradient(#fff, rgba(255, 255, 255, 0));
    background-image: linear-gradient(#fff, rgba(255, 255, 255, 0))
}
.sunken-menu:after {
    bottom: 0;
    background-color: transparent;
    background-image: -webkit-linear-gradient(rgba(255, 255, 255, 0), #fff);
    background-image: linear-gradient(rgba(255, 255, 255, 0), #fff)
}
.sunken-menu-separator {
    position: relative;
    height: 5px;
    margin: 8px 0 6px 1px;
    background-image: -webkit-radial-gradient(farthest-side at left top, #f4f4f4, rgba(244, 244, 244, 0));
    background-image: radial-gradient(farthest-side at left top, #f4f4f4, rgba(244, 244, 244, 0))
}
.sunken-menu-separator:before {
    position: absolute;
    top: 0;
    width: 100%;
    height: 1px;
    content: "";
    background-image: -webkit-linear-gradient(left, #eee 70%, #fff 100%);
    background-image: linear-gradient(to right, #eee 70%, #fff 100%);
    border-top: 1px solid #eee\9
}
.sunken-menu-group {
    list-style-type: none
}
.sunken-menu-item {
    display: block;
    padding: 8px 10px;
    margin-top: 5px;
    margin-bottom: 5px;
    border: 1px solid transparent;
    outline: 0
}
.sunken-menu-item .counter {
    position: absolute;
    top: 8px;
    right: 10px
}
.sunken-menu-item .octicon {
    left: -1px;
    width: 16px;
    color: #999;
    text-align: center
}
.sunken-menu-item .mini-loader {
    position: absolute;
    top: 9px;
    left: 11px;
    display: none
}
.sunken-menu-item:focus,
.sunken-menu-item:hover {
    text-decoration: none;
    box-shadow: inset 2px 0 0 #ccc
}
.sunken-menu-item:focus .octicon,
.sunken-menu-item:hover .octicon {
    color: #333
}
.sunken-menu-item.selected {
    font-weight: bold;
    color: #333;
    background-color: #fff;
    border-color: #eee #eee #eee transparent;
    border-radius: 0 3px 3px 0;
    box-shadow: 1px 1px 3px rgba(0, 0, 0, 0.05)
}
.sunken-menu-item.selected:after {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    width: 3px;
    content: "";
    background-color: #d26911;
    border-radius: 0 3px 3px 0
}
.sunken-menu-item.selected .octicon {
    color: #333
}
.sunken-menu-item.is-loading .mini-loader {
    display: inline-block;
    -webkit-animation: mini-nav-loader, 0.4s, forwards;
    animation: mini-nav-loader, 0.4s, forwards
}
.sunken-menu-item.is-loading .octicon {
    color: #fff
}
@-webkit-keyframes mini-nav-loader {
    0%, 90% {
        opacity: 0
    }
    100% {
        opacity: 1
    }
}
@keyframes mini-nav-loader {
    0%, 90% {
        opacity: 0
    }
    100% {
        opacity: 1
    }
}
@font-face {
    font-family: 'octicons';
    src: url(/assets/octicons/octicons/octicons.eot?#iefix) format("embedded-opentype"), url(/assets/octicons/octicons/octicons.woff) format("woff"), url(/assets/octicons/octicons/octicons.ttf) format("truetype"), url(/assets/octicons/octicons/octicons.svg#octicons) format("svg");
    font-weight: normal;
    font-style: normal
}
.octicon,
.mega-octicon {
    font: normal normal normal 16px/1 octicons;
    display: inline-block;
    text-decoration: none;
    text-rendering: auto;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
.mega-octicon {
    font-size: 32px
}
.octicon-alert:before {
    content: '\f02d'
}
.octicon-alignment-align:before {
    content: '\f08a'
}
.octicon-alignment-aligned-to:before {
    content: '\f08e'
}
.octicon-alignment-unalign:before {
    content: '\f08b'
}
.octicon-arrow-down:before {
    content: '\f03f'
}
.octicon-arrow-left:before {
    content: '\f040'
}
.octicon-arrow-right:before {
    content: '\f03e'
}
.octicon-arrow-small-down:before {
    content: '\f0a0'
}
.octicon-arrow-small-left:before {
    content: '\f0a1'
}
.octicon-arrow-small-right:before {
    content: '\f071'
}
.octicon-arrow-small-up:before {
    content: '\f09f'
}
.octicon-arrow-up:before {
    content: '\f03d'
}
.octicon-beer:before {
    content: '\f069'
}
.octicon-book:before {
    content: '\f007'
}
.octicon-bookmark:before {
    content: '\f07b'
}
.octicon-briefcase:before {
    content: '\f0d3'
}
.octicon-broadcast:before {
    content: '\f048'
}
.octicon-browser:before {
    content: '\f0c5'
}
.octicon-bug:before {
    content: '\f091'
}
.octicon-calendar:before {
    content: '\f068'
}
.octicon-check:before {
    content: '\f03a'
}
.octicon-checklist:before {
    content: '\f076'
}
.octicon-chevron-down:before {
    content: '\f0a3'
}
.octicon-chevron-left:before {
    content: '\f0a4'
}
.octicon-chevron-right:before {
    content: '\f078'
}
.octicon-chevron-up:before {
    content: '\f0a2'
}
.octicon-circle-slash:before {
    content: '\f084'
}
.octicon-circuit-board:before {
    content: '\f0d6'
}
.octicon-clippy:before {
    content: '\f035'
}
.octicon-clock:before {
    content: '\f046'
}
.octicon-cloud-download:before {
    content: '\f00b'
}
.octicon-cloud-upload:before {
    content: '\f00c'
}
.octicon-code:before {
    content: '\f05f'
}
.octicon-color-mode:before {
    content: '\f065'
}
.octicon-comment-add:before,
.octicon-comment:before {
    content: '\f02b'
}
.octicon-comment-discussion:before {
    content: '\f04f'
}
.octicon-credit-card:before {
    content: '\f045'
}
.octicon-dash:before {
    content: '\f0ca'
}
.octicon-dashboard:before {
    content: '\f07d'
}
.octicon-database:before {
    content: '\f096'
}
.octicon-device-camera:before {
    content: '\f056'
}
.octicon-device-camera-video:before {
    content: '\f057'
}
.octicon-device-desktop:before {
    content: '\f27c'
}
.octicon-device-mobile:before {
    content: '\f038'
}
.octicon-diff:before {
    content: '\f04d'
}
.octicon-diff-added:before {
    content: '\f06b'
}
.octicon-diff-ignored:before {
    content: '\f099'
}
.octicon-diff-modified:before {
    content: '\f06d'
}
.octicon-diff-removed:before {
    content: '\f06c'
}
.octicon-diff-renamed:before {
    content: '\f06e'
}
.octicon-ellipsis:before {
    content: '\f09a'
}
.octicon-eye-unwatch:before,
.octicon-eye-watch:before,
.octicon-eye:before {
    content: '\f04e'
}
.octicon-file-binary:before {
    content: '\f094'
}
.octicon-file-code:before {
    content: '\f010'
}
.octicon-file-directory:before {
    content: '\f016'
}
.octicon-file-media:before {
    content: '\f012'
}
.octicon-file-pdf:before {
    content: '\f014'
}
.octicon-file-submodule:before {
    content: '\f017'
}
.octicon-file-symlink-directory:before {
    content: '\f0b1'
}
.octicon-file-symlink-file:before {
    content: '\f0b0'
}
.octicon-file-text:before {
    content: '\f011'
}
.octicon-file-zip:before {
    content: '\f013'
}
.octicon-flame:before {
    content: '\f0d2'
}
.octicon-fold:before {
    content: '\f0cc'
}
.octicon-gear:before {
    content: '\f02f'
}
.octicon-gift:before {
    content: '\f042'
}
.octicon-gist:before {
    content: '\f00e'
}
.octicon-gist-secret:before {
    content: '\f08c'
}
.octicon-git-branch-create:before,
.octicon-git-branch-delete:before,
.octicon-git-branch:before {
    content: '\f020'
}
.octicon-git-commit:before {
    content: '\f01f'
}
.octicon-git-compare:before {
    content: '\f0ac'
}
.octicon-git-merge:before {
    content: '\f023'
}
.octicon-git-pull-request-abandoned:before,
.octicon-git-pull-request:before {
    content: '\f009'
}
.octicon-globe:before {
    content: '\f0b6'
}
.octicon-graph:before {
    content: '\f043'
}
.octicon-heart:before {
    content: '\2665'
}
.octicon-history:before {
    content: '\f07e'
}
.octicon-home:before {
    content: '\f08d'
}
.octicon-horizontal-rule:before {
    content: '\f070'
}
.octicon-hourglass:before {
    content: '\f09e'
}
.octicon-hubot:before {
    content: '\f09d'
}
.octicon-inbox:before {
    content: '\f0cf'
}
.octicon-info:before {
    content: '\f059'
}
.octicon-issue-closed:before {
    content: '\f028'
}
.octicon-issue-opened:before {
    content: '\f026'
}
.octicon-issue-reopened:before {
    content: '\f027'
}
.octicon-jersey:before {
    content: '\f019'
}
.octicon-jump-down:before {
    content: '\f072'
}
.octicon-jump-left:before {
    content: '\f0a5'
}
.octicon-jump-right:before {
    content: '\f0a6'
}
.octicon-jump-up:before {
    content: '\f073'
}
.octicon-key:before {
    content: '\f049'
}
.octicon-keyboard:before {
    content: '\f00d'
}
.octicon-law:before {
    content: '\f0d8'
}
.octicon-light-bulb:before {
    content: '\f000'
}
.octicon-link:before {
    content: '\f05c'
}
.octicon-link-external:before {
    content: '\f07f'
}
.octicon-list-ordered:before {
    content: '\f062'
}
.octicon-list-unordered:before {
    content: '\f061'
}
.octicon-location:before {
    content: '\f060'
}
.octicon-gist-private:before,
.octicon-mirror-private:before,
.octicon-git-fork-private:before,
.octicon-lock:before {
    content: '\f06a'
}
.octicon-logo-github:before {
    content: '\f092'
}
.octicon-mail:before {
    content: '\f03b'
}
.octicon-mail-read:before {
    content: '\f03c'
}
.octicon-mail-reply:before {
    content: '\f051'
}
.octicon-mark-github:before {
    content: '\f00a'
}
.octicon-markdown:before {
    content: '\f0c9'
}
.octicon-megaphone:before {
    content: '\f077'
}
.octicon-mention:before {
    content: '\f0be'
}
.octicon-microscope:before {
    content: '\f089'
}
.octicon-milestone:before {
    content: '\f075'
}
.octicon-mirror-public:before,
.octicon-mirror:before {
    content: '\f024'
}
.octicon-mortar-board:before {
    content: '\f0d7'
}
.octicon-move-down:before {
    content: '\f0a8'
}
.octicon-move-left:before {
    content: '\f074'
}
.octicon-move-right:before {
    content: '\f0a9'
}
.octicon-move-up:before {
    content: '\f0a7'
}
.octicon-mute:before {
    content: '\f080'
}
.octicon-no-newline:before {
    content: '\f09c'
}
.octicon-octoface:before {
    content: '\f008'
}
.octicon-organization:before {
    content: '\f037'
}
.octicon-package:before {
    content: '\f0c4'
}
.octicon-paintcan:before {
    content: '\f0d1'
}
.octicon-pencil:before {
    content: '\f058'
}
.octicon-person-add:before,
.octicon-person-follow:before,
.octicon-person:before {
    content: '\f018'
}
.octicon-pin:before {
    content: '\f041'
}
.octicon-playback-fast-forward:before {
    content: '\f0bd'
}
.octicon-playback-pause:before {
    content: '\f0bb'
}
.octicon-playback-play:before {
    content: '\f0bf'
}
.octicon-playback-rewind:before {
    content: '\f0bc'
}
.octicon-plug:before {
    content: '\f0d4'
}
.octicon-repo-create:before,
.octicon-gist-new:before,
.octicon-file-directory-create:before,
.octicon-file-add:before,
.octicon-plus:before {
    content: '\f05d'
}
.octicon-podium:before {
    content: '\f0af'
}
.octicon-primitive-dot:before {
    content: '\f052'
}
.octicon-primitive-square:before {
    content: '\f053'
}
.octicon-pulse:before {
    content: '\f085'
}
.octicon-puzzle:before {
    content: '\f0c0'
}
.octicon-question:before {
    content: '\f02c'
}
.octicon-quote:before {
    content: '\f063'
}
.octicon-radio-tower:before {
    content: '\f030'
}
.octicon-repo-delete:before,
.octicon-repo:before {
    content: '\f001'
}
.octicon-repo-clone:before {
    content: '\f04c'
}
.octicon-repo-force-push:before {
    content: '\f04a'
}
.octicon-gist-fork:before,
.octicon-repo-forked:before {
    content: '\f002'
}
.octicon-repo-pull:before {
    content: '\f006'
}
.octicon-repo-push:before {
    content: '\f005'
}
.octicon-rocket:before {
    content: '\f033'
}
.octicon-rss:before {
    content: '\f034'
}
.octicon-ruby:before {
    content: '\f047'
}
.octicon-screen-full:before {
    content: '\f066'
}
.octicon-screen-normal:before {
    content: '\f067'
}
.octicon-search-save:before,
.octicon-search:before {
    content: '\f02e'
}
.octicon-server:before {
    content: '\f097'
}
.octicon-settings:before {
    content: '\f07c'
}
.octicon-log-in:before,
.octicon-sign-in:before {
    content: '\f036'
}
.octicon-log-out:before,
.octicon-sign-out:before {
    content: '\f032'
}
.octicon-split:before {
    content: '\f0c6'
}
.octicon-squirrel:before {
    content: '\f0b2'
}
.octicon-star-add:before,
.octicon-star-delete:before,
.octicon-star:before {
    content: '\f02a'
}
.octicon-steps:before {
    content: '\f0c7'
}
.octicon-stop:before {
    content: '\f08f'
}
.octicon-repo-sync:before,
.octicon-sync:before {
    content: '\f087'
}
.octicon-tag-remove:before,
.octicon-tag-add:before,
.octicon-tag:before {
    content: '\f015'
}
.octicon-telescope:before {
    content: '\f088'
}
.octicon-terminal:before {
    content: '\f0c8'
}
.octicon-three-bars:before {
    content: '\f05e'
}
.octicon-thumbsdown:before {
    content: '\f0db'
}
.octicon-thumbsup:before {
    content: '\f0da'
}
.octicon-tools:before {
    content: '\f031'
}
.octicon-trashcan:before {
    content: '\f0d0'
}
.octicon-triangle-down:before {
    content: '\f05b'
}
.octicon-triangle-left:before {
    content: '\f044'
}
.octicon-triangle-right:before {
    content: '\f05a'
}
.octicon-triangle-up:before {
    content: '\f0aa'
}
.octicon-unfold:before {
    content: '\f039'
}
.octicon-unmute:before {
    content: '\f0ba'
}
.octicon-versions:before {
    content: '\f064'
}
.octicon-remove-close:before,
.octicon-x:before {
    content: '\f081'
}
.octicon-zap:before {
    content: '\26A1'
}
.date_selector {
    width: 225px;
    text-align: left;
    text-decoration: none;
    z-index: 9;
    display: none
}
.date_selector .month_nav,
.date_selector .year_nav {
    margin-top: 5px;
    margin-bottom: 5px;
    padding: 0;
    display: block;
    position: relative;
    text-align: center;
    line-height: 20px
}
.date_selector .month_nav {
    float: left;
    width: 55%
}
.date_selector .year_nav {
    float: right;
    width: 35%
}
.date_selector .date-button {
    position: absolute;
    top: 0;
    width: 18px;
    height: 18px;
    padding: 4px;
    color: #4078c0;
    font-size: 12px;
    cursor: pointer;
    line-height: 12px
}
.date_selector .prev {
    left: 0
}
.date_selector .next {
    right: 0
}
.date_selector table {
    width: 100%;
    clear: both
}
.date_selector tr {
    font-size: 0
}
.date_selector th,
.date_selector td {
    width: 32px;
    height: 32px;
    line-height: 28px;
    padding: 0;
    text-align: center;
    font-weight: normal;
    display: inline-block;
    font-size: 12px;
    margin-top: -1px;
    margin-left: -1px
}
.date_selector td {
    border: 1px solid #ccc;
    color: #4078c0;
    background: #fff;
    cursor: default
}
.date_selector td.today {
    background: #eee
}
.date_selector td.selected,
.date_selector td.selectable_day:hover {
    background: #4078c0;
    color: #fff;
    border-color: #33609a;
    z-index: 10;
    position: relative;
    cursor: pointer
}
.date_selector td.unselected_month {
    color: #ccc
}
.jcrop-holder {
    direction: ltr;
    text-align: left;
    -ms-touch-action: none;
    touch-action: none
}
.jcrop-vline,
.jcrop-hline {
    background: #fff url(https://github.com/images/spinners/Jcrop.gif);
    font-size: 0;
    position: absolute
}
.jcrop-vline {
    height: 100%;
    width: 1px !important
}
.jcrop-vline.right {
    right: 0
}
.jcrop-hline {
    height: 1px !important;
    width: 100%
}
.jcrop-hline.bottom {
    bottom: 0
}
.jcrop-tracker {
    height: 100%;
    width: 100%;
    -webkit-tap-highlight-color: transparent;
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
.jcrop-handle {
    background-color: #333;
    border: 1px #eee solid;
    width: 7px;
    height: 7px;
    font-size: 1px
}
.jcrop-handle.ord-n {
    left: 50%;
    margin-left: -4px;
    margin-top: -4px;
    top: 0
}
.jcrop-handle.ord-s {
    bottom: 0;
    left: 50%;
    margin-bottom: -4px;
    margin-left: -4px
}
.jcrop-handle.ord-e {
    margin-right: -4px;
    margin-top: -4px;
    right: 0;
    top: 50%
}
.jcrop-handle.ord-w {
    left: 0;
    margin-left: -4px;
    margin-top: -4px;
    top: 50%
}
.jcrop-handle.ord-nw {
    left: 0;
    margin-left: -4px;
    margin-top: -4px;
    top: 0
}
.jcrop-handle.ord-ne {
    margin-right: -4px;
    margin-top: -4px;
    right: 0;
    top: 0
}
.jcrop-handle.ord-se {
    bottom: 0;
    margin-bottom: -4px;
    margin-right: -4px;
    right: 0
}
.jcrop-handle.ord-sw {
    bottom: 0;
    left: 0;
    margin-bottom: -4px;
    margin-left: -4px
}
.jcrop-dragbar.ord-n,
.jcrop-dragbar.ord-s {
    height: 7px;
    width: 100%
}
.jcrop-dragbar.ord-e,
.jcrop-dragbar.ord-w {
    height: 100%;
    width: 7px
}
.jcrop-dragbar.ord-n {
    margin-top: -4px
}
.jcrop-dragbar.ord-s {
    bottom: 0;
    margin-bottom: -4px
}
.jcrop-dragbar.ord-e {
    margin-right: -4px;
    right: 0
}
.jcrop-dragbar.ord-w {
    margin-left: -4px
}
.jcrop-light .jcrop-vline,
.jcrop-light .jcrop-hline {
    background: #fff;
    opacity: 0.7 !important
}
.jcrop-light .jcrop-handle {
    background-color: #000;
    border-color: #fff;
    border-radius: 3px
}
.jcrop-dark .jcrop-vline,
.jcrop-dark .jcrop-hline {
    background: #000;
    opacity: 0.7 !important
}
.jcrop-dark .jcrop-handle {
    background-color: #fff;
    border-color: #000;
    border-radius: 3px
}
.jcrop-holder img,
img.jcrop-preview {
    max-width: none
}
.code-frequency .addition {
    fill-opacity: 1;
    fill: #6cc644
}
.code-frequency .deletion {
    fill-opacity: 1;
    fill: #bd2c00
}
.cadd {
    font-weight: bold;
    color: #6cc644
}
.cdel {
    font-weight: bold;
    color: #bd2c00
}
.commit-activity-graphs .dots {
    display: none
}
#commit-activity-master {
    margin-top: 20px
}
.is-graph-loading #commit-activity-master {
    display: none
}
rect {
    shape-rendering: crispedges
}
rect.max {
    fill: #ffc644
}
g.bar {
    fill: #1db34f
}
g.mini {
    fill: #f17f49
}
g.active rect {
    fill: #bd380f
}
circle.focus {
    fill: #555
}
.dot text {
    stroke: none;
    fill: #555
}
.tint-box {
    border-radius: 6px;
    background: #f3f3f3;
    position: relative;
    margin-bottom: 10px
}
.tint-box.transparent {
    background: #fff
}
.tint-box .activity {
    margin-top: 0;
    padding-top: 100px
}
.contrib-data {
    margin: 0 0 10px;
    list-style: none;
    padding: 0
}
#contributors .capped-card .avatar {
    float: left;
    width: 32px;
    height: 32px;
    margin-right: 5px
}
#contributors .capped-card h3 {
    font-weight: normal
}
#contributors .capped-card .ameta {
    display: block;
    font-size: 12px;
    color: #ccc
}
#contributors .capped-card .rank {
    float: right;
    color: #767676;
    font-size: 13px
}
#contributors .capped-card .cmt {
    color: #767676
}
#contributors .capped-card path {
    fill: #f17f49
}
#contributors .capped-card .midlabel {
    fill: #ccc
}
.d {
    color: #bd2c00
}
.a {
    color: #6cc644
}
.axis {
    fill: #aaa;
    font-size: 10px
}
.axis line {
    shape-rendering: crispedges;
    stroke: #eee
}
.axis path {
    display: none
}
.axis .zero line {
    stroke-width: 1.5;
    stroke: #4078c0;
    stroke-dasharray: 3 3
}
.graphs .is-graph-loading {
    min-height: 500px
}
.graphs.wheader h2 {
    padding: 1px
}
.graphs .area {
    fill: #1db34f;
    fill-opacity: 0.5
}
.graphs .path {
    stroke: #1db34f;
    stroke-width: 2px;
    stroke-opacity: 1;
    fill: none
}
.graphs .dot {
    fill: #1db34f;
    stroke: #16873c;
    stroke-width: 2px
}
.graphs .dot.padded {
    stroke: #fff;
    stroke-width: 1px
}
.graphs .dot.padded circle:hover {
    fill: #4078c0
}
.graphs .d3-tip {
    fill: #333
}
.graphs .d3-tip text {
    fill: #fff;
    font-size: 11px
}
.graphs .dir {
    font-size: 12px;
    font-weight: normal;
    color: #555;
    line-height: 100%;
    padding-top: 5px;
    float: right
}
.graphs .selection rect {
    fill: #333;
    fill-opacity: 0.1;
    stroke: #333;
    stroke-width: 1px;
    stroke-opacity: 0.4;
    shape-rendering: crispedges;
    stroke-dasharray: 3 3
}
.graph-filter h3 {
    display: inline-block;
    margin: 10px 0 0;
    font-weight: 300;
    font-size: 24px
}
.graph-filter .info {
    margin-top: 5px;
    margin-bottom: 20px;
    color: #767676
}
.graph-filter .select-menu {
    float: right;
    margin-top: 13px
}
h2.ghead:after {
    content: ".";
    height: 0;
    display: block;
    visibility: hidden;
    clear: both
}
.graph-canvas .activity {
    text-align: center;
    width: 400px;
    margin: 100px auto 0;
    color: #444;
    border-radius: 3px;
    padding: 10px
}
.graph-canvas .error {
    color: #900;
    background: #feeaea;
    padding: 10px;
    border-radius: 3px
}
.graph-canvas .dots {
    margin: 0 auto
}
.graph-canvas>.activity {
    display: none
}
.graph-loading,
.graph-error,
.graph-no-usable-data,
.graph-empty {
    display: none
}
.graph-canvas.is-graph-loading>.activity,
.graph-canvas.is-graph-without-usable-data>.activity,
.graph-canvas.is-graph-empty>.activity {
    display: block
}
.is-graph-loading .graph-loading,
.is-graph-empty .graph-empty,
.is-graph-without-usable-data .graph-no-usable-data,
.is-graph-load-error .graph-error {
    display: block
}
.svg-tip {
    padding: 10px;
    background: rgba(0, 0, 0, 0.8);
    color: #bbb;
    font-size: 12px;
    position: absolute;
    z-index: 99999;
    text-align: center;
    border-radius: 3px
}
.svg-tip strong {
    color: #ddd
}
.svg-tip.is-visible {
    display: block
}
.svg-tip:after {
    box-sizing: border-box;
    position: absolute;
    left: 50%;
    height: 5px;
    width: 5px;
    bottom: -10px;
    margin: 0 0 0 -5px;
    content: " ";
    border: 5px solid transparent;
    border-top-color: rgba(0, 0, 0, 0.8)
}
.svg-tip.comparison {
    text-align: left;
    pointer-events: none;
    padding: 0
}
.svg-tip.comparison .title {
    display: block;
    padding: 10px;
    margin: 0;
    line-height: 1;
    font-weight: bold;
    pointer-events: none
}
.svg-tip.comparison ul {
    list-style: none;
    margin: 0;
    white-space: nowrap
}
.svg-tip.comparison li {
    display: inline-block;
    padding: 10px
}
.svg-tip.comparison li:first-child {
    border-top: 3px solid #1db34f;
    border-right: 1px solid #333
}
.svg-tip.comparison li:last-child {
    border-top: 3px solid #1d7fb3
}
.svg-tip-one-line {
    white-space: nowrap
}
.day-name {
    fill: #555
}
circle.day {
    stroke-width: 0;
    fill: #444
}
circle.day:hover {
    fill: #4078c0
}
line.axis {
    stroke-width: 1;
    stroke: #eee;
    shape-rendering: crispedges
}
line.axis.even {
    stroke: #e0e0e0
}
.traffic-graph {
    min-height: 150px
}
.traffic-graph .activity {
    margin-top: 0
}
.traffic-graph .activity .dots {
    margin-top: 40px
}
.traffic-graph .path {
    fill: none;
    stroke-width: 2
}
.traffic-graph path.total {
    stroke: #1db34f
}
.traffic-graph path.unique {
    stroke: #1d7fb3
}
.traffic-graph .x.axis .tick:first-child line {
    stroke: #1db34f;
    stroke-width: 2px
}
.traffic-graph .y line {
    stroke: #1db34f
}
.traffic-graph .y.unique line {
    stroke: #1d7fb3
}
.traffic-graph .overlay {
    fill-opacity: 0
}
.uniques-graph .x.axis .tick:nth-child(14) line {
    stroke: #1d7fb3;
    stroke-width: 2px
}
.svg-tip .date {
    color: #fff
}
#top-domains .dots {
    margin: 167px auto 0;
    display: block
}
#top-domains .favicon {
    width: 16px;
    height: 16px
}
table.capped-list {
    width: 100%;
    line-height: 100%
}
table.capped-list th {
    text-align: left;
    padding: 8px;
    border-bottom: 1px solid #ddd;
    background: #f4f4f4
}
table.capped-list td {
    padding: 8px;
    border-bottom: 1px solid #eee;
    font-size: 12px
}
table.capped-list th.middle,
table.capped-list td.middle {
    text-align: center
}
table.capped-list .favicon {
    width: 16px;
    height: 16px;
    vertical-align: middle;
    margin: 0 5px
}
table.capped-list .octicon {
    margin-right: 10px;
    vertical-align: -1px;
    color: #555
}
table.capped-list tr:nth-child(even) {
    background-color: #fcfcfc
}
table.capped-list.mini-icons .mini-icon {
    margin-right: 5px;
    color: #555
}
.capped-list-label {
    overflow: hidden;
    white-space: nowrap;
    max-width: 200px;
    text-overflow: ellipsis
}
.traffic-graph-stats {
    border-top: 1px solid #ddd
}
.traffic-graph-stats .summary-stats {
    width: 100%
}
.traffic-graph-stats .summary-stats:before {
    display: table;
    content: ""
}
.traffic-graph-stats .summary-stats:after {
    display: table;
    clear: both;
    content: ""
}
.traffic-graph-stats .summary-stats li {
    width: 50%;
    display: block;
    float: left;
    padding-bottom: 10px
}
.totals circle {
    fill: #1db34f;
    stroke: #fff;
    stroke-width: 2
}
.uniques circle {
    fill: #1d7fb3;
    stroke: #fff;
    stroke-width: 2
}
.top-lists .is-loading {
    text-align: center;
    margin: 40px
}
ul.web-views li {
    width: 140px
}
ul.clones li {
    width: 170px
}
.markdown-body {
    overflow: hidden;
    font-family: "Helvetica Neue", Helvetica, "Segoe UI", Arial, freesans, sans-serif;
    font-size: 16px;
    line-height: 1.6;
    word-wrap: break-word
}
.markdown-body>*:first-child {
    margin-top: 0 !important
}
.markdown-body>*:last-child {
    margin-bottom: 0 !important
}
.markdown-body a:not([href]) {
    color: inherit;
    text-decoration: none
}
.markdown-body .absent {
    color: #c00
}
.markdown-body .anchor {
    position: absolute;
    top: 0;
    left: 0;
    display: block;
    padding-right: 6px;
    padding-left: 30px;
    margin-left: -30px
}
.markdown-body .anchor:focus {
    outline: none
}
.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
    position: relative;
    margin-top: 1em;
    margin-bottom: 16px;
    font-weight: bold;
    line-height: 1.4
}
.markdown-body h1 .octicon-link,
.markdown-body h2 .octicon-link,
.markdown-body h3 .octicon-link,
.markdown-body h4 .octicon-link,
.markdown-body h5 .octicon-link,
.markdown-body h6 .octicon-link {
    display: none;
    color: #000;
    vertical-align: middle
}
.markdown-body h1:hover .anchor,
.markdown-body h2:hover .anchor,
.markdown-body h3:hover .anchor,
.markdown-body h4:hover .anchor,
.markdown-body h5:hover .anchor,
.markdown-body h6:hover .anchor {
    padding-left: 8px;
    margin-left: -30px;
    text-decoration: none
}
.markdown-body h1:hover .anchor .octicon-link,
.markdown-body h2:hover .anchor .octicon-link,
.markdown-body h3:hover .anchor .octicon-link,
.markdown-body h4:hover .anchor .octicon-link,
.markdown-body h5:hover .anchor .octicon-link,
.markdown-body h6:hover .anchor .octicon-link {
    display: inline-block
}
.markdown-body h1 tt,
.markdown-body h1 code,
.markdown-body h2 tt,
.markdown-body h2 code,
.markdown-body h3 tt,
.markdown-body h3 code,
.markdown-body h4 tt,
.markdown-body h4 code,
.markdown-body h5 tt,
.markdown-body h5 code,
.markdown-body h6 tt,
.markdown-body h6 code {
    font-size: inherit
}
.markdown-body h1 {
    padding-bottom: 0.3em;
    font-size: 2.25em;
    line-height: 1.2;
    border-bottom: 1px solid #eee
}
.markdown-body h1 .anchor {
    line-height: 1
}
.markdown-body h2 {
    padding-bottom: 0.3em;
    font-size: 1.75em;
    line-height: 1.225;
    border-bottom: 1px solid #eee
}
.markdown-body h2 .anchor {
    line-height: 1
}
.markdown-body h3 {
    font-size: 1.5em;
    line-height: 1.43
}
.markdown-body h3 .anchor {
    line-height: 1.2
}
.markdown-body h4 {
    font-size: 1.25em
}
.markdown-body h4 .anchor {
    line-height: 1.2
}
.markdown-body h5 {
    font-size: 1em
}
.markdown-body h5 .anchor {
    line-height: 1.1
}
.markdown-body h6 {
    font-size: 1em;
    color: #777
}
.markdown-body h6 .anchor {
    line-height: 1.1
}
.markdown-body p,
.markdown-body blockquote,
.markdown-body ul,
.markdown-body ol,
.markdown-body dl,
.markdown-body table,
.markdown-body pre {
    margin-top: 0;
    margin-bottom: 16px
}
.markdown-body hr {
    height: 4px;
    padding: 0;
    margin: 16px 0;
    background-color: #e7e7e7;
    border: 0 none
}
.markdown-body ul,
.markdown-body ol {
    padding-left: 2em
}
.markdown-body ul.no-list,
.markdown-body ol.no-list {
    padding: 0;
    list-style-type: none
}
.markdown-body ul ul,
.markdown-body ul ol,
.markdown-body ol ol,
.markdown-body ol ul {
    margin-top: 0;
    margin-bottom: 0
}
.markdown-body li>p {
    margin-top: 16px
}
.markdown-body dl {
    padding: 0
}
.markdown-body dl dt {
    padding: 0;
    margin-top: 16px;
    font-size: 1em;
    font-style: italic;
    font-weight: bold
}
.markdown-body dl dd {
    padding: 0 16px;
    margin-bottom: 16px
}
.markdown-body blockquote {
    padding: 0 15px;
    color: #777;
    border-left: 4px solid #ddd
}
.markdown-body blockquote>:first-child {
    margin-top: 0
}
.markdown-body blockquote>:last-child {
    margin-bottom: 0
}
.markdown-body table {
    display: block;
    width: 100%;
    overflow: auto;
    word-break: normal;
    word-break: keep-all
}
.markdown-body table th {
    font-weight: bold
}
.markdown-body table th,
.markdown-body table td {
    padding: 6px 13px;
    border: 1px solid #ddd
}
.markdown-body table tr {
    background-color: #fff;
    border-top: 1px solid #ccc
}
.markdown-body table tr:nth-child(2n) {
    background-color: #f8f8f8
}
.markdown-body img {
    max-width: 100%;
    box-sizing: border-box
}
.markdown-body .emoji {
    max-width: none
}
.markdown-body span.frame {
    display: block;
    overflow: hidden
}
.markdown-body span.frame>span {
    display: block;
    float: left;
    width: auto;
    padding: 7px;
    margin: 13px 0 0;
    overflow: hidden;
    border: 1px solid #ddd
}
.markdown-body span.frame span img {
    display: block;
    float: left
}
.markdown-body span.frame span span {
    display: block;
    padding: 5px 0 0;
    clear: both;
    color: #333
}
.markdown-body span.align-center {
    display: block;
    overflow: hidden;
    clear: both
}
.markdown-body span.align-center>span {
    display: block;
    margin: 13px auto 0;
    overflow: hidden;
    text-align: center
}
.markdown-body span.align-center span img {
    margin: 0 auto;
    text-align: center
}
.markdown-body span.align-right {
    display: block;
    overflow: hidden;
    clear: both
}
.markdown-body span.align-right>span {
    display: block;
    margin: 13px 0 0;
    overflow: hidden;
    text-align: right
}
.markdown-body span.align-right span img {
    margin: 0;
    text-align: right
}
.markdown-body span.float-left {
    display: block;
    float: left;
    margin-right: 13px;
    overflow: hidden
}
.markdown-body span.float-left span {
    margin: 13px 0 0
}
.markdown-body span.float-right {
    display: block;
    float: right;
    margin-left: 13px;
    overflow: hidden
}
.markdown-body span.float-right>span {
    display: block;
    margin: 13px auto 0;
    overflow: hidden;
    text-align: right
}
.markdown-body code,
.markdown-body tt {
    padding: 0;
    padding-top: 0.2em;
    padding-bottom: 0.2em;
    margin: 0;
    font-size: 85%;
    background-color: rgba(0, 0, 0, 0.04);
    border-radius: 3px
}
.markdown-body code:before,
.markdown-body code:after,
.markdown-body tt:before,
.markdown-body tt:after {
    letter-spacing: -0.2em;
    content: "\00a0"
}
.markdown-body code br,
.markdown-body tt br {
    display: none
}
.markdown-body del code {
    text-decoration: inherit
}
.markdown-body pre>code {
    padding: 0;
    margin: 0;
    font-size: 100%;
    word-break: normal;
    white-space: pre;
    background: transparent;
    border: 0
}
.markdown-body .highlight {
    margin-bottom: 16px
}
.markdown-body .highlight pre,
.markdown-body pre {
    padding: 16px;
    overflow: auto;
    font-size: 85%;
    line-height: 1.45;
    background-color: #f7f7f7;
    border-radius: 3px
}
.markdown-body .highlight pre {
    margin-bottom: 0;
    word-break: normal
}
.markdown-body pre {
    word-wrap: normal
}
.markdown-body pre code,
.markdown-body pre tt {
    display: inline;
    max-width: initial;
    padding: 0;
    margin: 0;
    overflow: initial;
    line-height: inherit;
    word-wrap: normal;
    background-color: transparent;
    border: 0
}
.markdown-body pre code:before,
.markdown-body pre code:after,
.markdown-body pre tt:before,
.markdown-body pre tt:after {
    content: normal
}
.markdown-body kbd {
    display: inline-block;
    padding: 3px 5px;
    font-size: 11px;
    line-height: 10px;
    color: #555;
    vertical-align: middle;
    background-color: #fcfcfc;
    border: solid 1px #ccc;
    border-bottom-color: #bbb;
    border-radius: 3px;
    box-shadow: inset 0 -1px 0 #bbb
}
.pl-c {
    color: #969896
}
.pl-c1,
.pl-s .pl-v {
    color: #0086b3
}
.pl-e,
.pl-en {
    color: #795da3
}
.pl-s .pl-s1,
.pl-smi {
    color: #333
}
.pl-ent {
    color: #63a35c
}
.pl-k {
    color: #a71d5d
}
.pl-pds,
.pl-s,
.pl-s .pl-pse .pl-s1,
.pl-sr,
.pl-sr .pl-cce,
.pl-sr .pl-sra,
.pl-sr .pl-sre {
    color: #183691
}
.pl-v {
    color: #ed6a43
}
.pl-id {
    color: #b52a1d
}
.pl-ii {
    background-color: #b52a1d;
    color: #f8f8f8
}
.pl-sr .pl-cce {
    color: #63a35c;
    font-weight: bold
}
.pl-ml {
    color: #693a17
}
.pl-mh,
.pl-mh .pl-en,
.pl-ms {
    color: #1d3e81;
    font-weight: bold
}
.pl-mq {
    color: #008080
}
.pl-mi {
    color: #333;
    font-style: italic
}
.pl-mb {
    color: #333;
    font-weight: bold
}
.pl-md {
    background-color: #ffecec;
    color: #bd2c00
}
.pl-mi1 {
    background-color: #eaffea;
    color: #55a532
}
.pl-mdr {
    color: #795da3;
    font-weight: bold
}
.pl-mo {
    color: #1d3e81
}
.ace_gutter {
    background: #ffffff;
    color: #999999
}
.ace_print-margin {
    width: 1px;
    background: #e8e8e8
}
.ace-github-light {
    background-color: #ffffff;
    color: #333333
}
.ace_cursor {
    color: #000000
}
.ace_marker-layer .ace_selection {
    background: #c8c8fa
}
.ace_multiselect .ace_selection.ace_start {
    box-shadow: 0 0 3px 0px #ffffff;
    border-radius: 2px
}
.ace_marker-layer .ace_step {
    background: #c6dbae
}
.ace_marker-layer .ace_bracket {
    margin: -1px 0 0 -1px;
    border: 1px solid #c0c0c0
}
.ace_marker-layer .ace_active-line {
    background: #f5f5f5
}
.ace_gutter-active-line {
    background-color: #f5f5f5
}
.ace_marker-layer .ace_selected-word {
    border: 1px solid #c8c8fa
}
.ace_fold {
    background-color: #a71d5d;
    border-color: #333333
}
.ace_keyword {
    color: #a71d5d
}
.ace_constant {
    color: #0086b3
}
.ace_support {
    color: #0086b3
}
.ace_support.ace_constant {
    color: #0086b3
}
.ace_support.ace_type {
    color: #a71d5d
}
.ace_storage {
    color: #a71d5d
}
.ace_storage.ace_type {
    color: #a71d5d
}
.ace_invalid.ace_illegal {
    text-decoration: underline;
    font-style: italic;
    color: #f8f8f8;
    background-color: #b52a1d
}
.ace_invalid.ace_deprecated {
    text-decoration: underline;
    font-style: italic;
    color: #b52a1d
}
.ace_string {
    color: #183691
}
.ace_string.ace_regexp {
    color: #183691
}
.ace_comment {
    color: #969896
}
.ace_variable {
    color: #ed6a43
}
.ace_entity.ace_name {
    color: #795da3
}
.ace_entity.ace_name.ace_tag {
    color: #63a35c
}
.ace_markup.ace_heading {
    color: #1d3e81
}
.ace_markup.ace_list {
    color: #693a17
}
body {
    word-wrap: break-word
}
.focus-content {
    width: 620px
}
#site-container>.container:first-child {
    margin-top: 20px
}
.emoji-icon {
    display: inline-block;
    height: 20px;
    width: 20px;
    vertical-align: middle;
    background-repeat: no-repeat;
    background-size: 20px 20px
}
.labels {
    position: relative
}
.label {
    display: inline-block;
    padding: 3px 4px;
    font-size: 11px;
    font-weight: bold;
    line-height: 1;
    color: #fff;
    border-radius: 2px;
    box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.12)
}
.label:hover {
    text-decoration: none
}
.label-admin {
    color: #666;
    background-color: #eee
}
.label-generic {
    margin-top: -1px;
    margin-bottom: -1px;
    color: #767676;
    font-weight: normal;
    background-color: transparent;
    border: 1px solid #eee;
    box-shadow: none
}
.label-recommended {
    color: #60b044;
    border: 1px solid #60b044
}
.label-neutral {
    background-color: #767676
}
.label-private {
    color: #a1882b;
    background-color: #ffefc6
}
a.label-link {
    border: 1px solid transparent
}
a.label-link:hover {
    text-decoration: none
}
.label-membership-pending {
    background-color: #c9510c
}
.label-review {
    color: #4c4a42;
    background-color: #fceb9b
}
.label-success {
    background-color: #6cc644
}
.label-coming-soon {
    background-color: #f93
}
.label-active {
    background-color: #6cc644
}
.facebox {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 100;
    padding-bottom: 40px
}
.facebox ul {
    margin-left: 25px;
    margin-bottom: 15px
}
.facebox .facebox-staff-links {
    margin: -15px -15px 15px;
    padding: 10px 15px;
    background-color: #f5f5f5;
    border-bottom: 1px solid #e5e5e5
}
.facebox .facebox-staff-links li {
    display: inline-block;
    margin-right: 10px;
    color: #767676;
    list-style: none
}
.facebox .facebox-staff-links a {
    font-weight: bold
}
.facebox pre {
    padding: 10px;
    background-color: #eee;
    border: 1px solid #ddd;
    border-radius: 3px
}
.facebox .shortcuts {
    width: 860px
}
.facebox .facebox-user-list {
    margin-left: 0;
    margin-bottom: 0;
    max-height: 400px;
    overflow: auto
}
.facebox .lineprofiler {
    width: 900px
}
.facebox .lineprofiler pre {
    overflow-x: scroll;
    white-space: pre;
    word-wrap: normal
}
.facebox-popup {
    position: relative;
    background-color: #fff;
    border: 1px solid rgba(0, 0, 0, 0.25);
    border-radius: 5px;
    box-shadow: 0 0 18px rgba(0, 0, 0, 0.4);
    background-clip: padding-box
}
.facebox-content {
    width: 455px;
    padding: 15px
}
.facebox-content:before {
    display: table;
    content: ""
}
.facebox-content:after {
    display: table;
    clear: both;
    content: ""
}
.facebox-close {
    position: absolute;
    top: 8px;
    right: 5px;
    padding: 10px;
    -webkit-appearance: none;
    background-color: transparent;
    border: 0;
    opacity: 0.25;
    cursor: pointer
}
.facebox-close:hover {
    opacity: 1
}
.facebox-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%
}
.facebox-overlay-hide {
    z-index: -100
}
.facebox-overlay-active {
    z-index: 99;
    background-color: #000
}
.facebox-loading {
    min-height: 64px;
    background-image: url(https://github.com/images/spinners/octocat-spinner-64.gif);
    background-position: center center;
    background-repeat: no-repeat
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    .facebox-loading {
        background-image: url(https://github.com/images/spinners/octocat-spinner-128.gif);
        background-size: 64px 64px
    }
}
.facebox-header {
    margin: -15px -15px 15px;
    padding: 15px;
    border-bottom: 1px solid #e5e5e5;
    font-size: 18px;
    font-weight: normal
}
.facebox-footer {
    margin: 0 -15px -15px;
    padding: 10px 15px;
    border-top: 1px solid #e5e5e5;
    border-bottom-right-radius: 5px;
    border-bottom-left-radius: 5px;
    background: #fafafa;
    text-align: right
}
.facebox-footer .help {
    margin: 0;
    text-align: center;
    color: #767676
}
.facebox-alert,
.facebox-danger {
    margin: -16px -15px 15px;
    padding: 10px 15px;
    border-style: solid;
    border-width: 1px 0
}
.facebox-alert {
    color: #796620;
    background-color: #f8eec7;
    border-color: #f2e09a
}
.facebox-danger {
    padding-left: 40px;
    color: #9c342e;
    background-color: #f7d9d7;
    border-color: #f2c4c2
}
.facebox-danger .octicon {
    float: left;
    margin-left: -25px
}
.facebox-separator {
    margin: 20px -15px
}
.facebox-staff-search .hfields {
    margin-top: 0;
    margin-bottom: 0
}
.facebox-staff-search .hfields input[type="text"] {
    width: 340px;
    margin-right: 0
}
.facebox-staff-search .hfields .btn {
    margin-top: 29px;
    margin-right: 0
}
.facebox-staff-search .status-check-list {
    float: none;
    margin: 15px 0 0
}
#facebox .billing-credit-cards {
    margin: 0 0 15px
}
#facebox .billing-credit-cards li {
    margin: 0 4px 0 0
}
.keyboard-shortcuts {
    float: right;
    font-size: 11px;
    color: #767676
}
.keyboard-shortcuts .mini-icon {
    position: relative;
    top: 2px;
    margin-left: 5px
}
.keyboard-mappings {
    font-size: 12px;
    color: #555
}
.keyboard-mappings th {
    padding-top: 25px;
    font-size: 14px;
    line-height: 1.5;
    color: #333;
    text-align: left
}
.keyboard-mappings tbody:first-child tr:first-child th {
    padding-top: 0
}
.keyboard-mappings td {
    padding-top: 3px;
    padding-bottom: 3px;
    vertical-align: top;
    line-height: 20px
}
.keyboard-mappings .keys {
    padding-right: 10px;
    color: #767676;
    text-align: right;
    white-space: nowrap
}
.keyboard-mappings .platform-mac {
    display: none
}
.macintosh .keyboard-mappings .platform-mac {
    display: inline
}
.macintosh .keyboard-mappings .platform-other {
    display: none
}
.facebox-user-list-item {
    padding: 3px 0;
    list-style: none;
    font-weight: bold;
    vertical-align: middle
}
.facebox-user-list-item a {
    color: #000
}
.facebox-user-list-item img {
    margin-right: 5px;
    border-radius: 3px;
    vertical-align: middle
}
.linejump .linejump-input {
    width: 340px;
    background-color: #fafafa
}
.linejump .linejump-input,
.linejump .btn {
    font-size: 16px;
    padding: 10px 15px
}
.linejump+.facebox-close {
    top: 18px
}
.repo-transfer-tip {
    margin-bottom: 0
}
.user-mention,
.team-mention {
    font-weight: bold;
    color: #333;
    white-space: nowrap
}
dl.form>dd input[type="text"].is-autocheck-loading,
dl.form>dd input[type="text"].is-autocheck-successful,
dl.form>dd input[type="text"].is-autocheck-errored,
dl.form>dd input[type="password"].is-autocheck-loading,
dl.form>dd input[type="password"].is-autocheck-successful,
dl.form>dd input[type="password"].is-autocheck-errored,
dl.form>dd input[type="email"].is-autocheck-loading,
dl.form>dd input[type="email"].is-autocheck-successful,
dl.form>dd input[type="email"].is-autocheck-errored {
    padding-right: 30px
}
dl.form>dd input[type="text"].is-autocheck-loading,
dl.form>dd input[type="password"].is-autocheck-loading,
dl.form>dd input[type="email"].is-autocheck-loading {
    background-image: url(https://github.com/images/spinners/octocat-spinner-16.gif)
}
dl.form>dd input[type="text"].is-autocheck-successful,
dl.form>dd input[type="password"].is-autocheck-successful,
dl.form>dd input[type="email"].is-autocheck-successful {
    background-image: url(https://github.com/images/modules/ajax/success.png)
}
dl.form>dd input[type="text"].is-autocheck-errored,
dl.form>dd input[type="password"].is-autocheck-errored,
dl.form>dd input[type="email"].is-autocheck-errored {
    background-image: url(https://github.com/images/modules/ajax/error.png)
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    dl.form>dd input[type="text"].is-autocheck-loading,
    dl.form>dd input[type="text"].is-autocheck-successful,
    dl.form>dd input[type="text"].is-autocheck-errored,
    dl.form>dd input[type="password"].is-autocheck-loading,
    dl.form>dd input[type="password"].is-autocheck-successful,
    dl.form>dd input[type="password"].is-autocheck-errored,
    dl.form>dd input[type="email"].is-autocheck-loading,
    dl.form>dd input[type="email"].is-autocheck-successful,
    dl.form>dd input[type="email"].is-autocheck-errored {
        background-size: 16px 16px
    }
    dl.form>dd input[type="text"].is-autocheck-loading,
    dl.form>dd input[type="password"].is-autocheck-loading,
    dl.form>dd input[type="email"].is-autocheck-loading {
        background-image: url(https://github.com/images/spinners/octocat-spinner-32.gif)
    }
    dl.form>dd input[type="text"].is-autocheck-successful,
    dl.form>dd input[type="password"].is-autocheck-successful,
    dl.form>dd input[type="email"].is-autocheck-successful {
        background-image: url(https://github.com/images/modules/ajax/success@2x.png)
    }
    dl.form>dd input[type="text"].is-autocheck-errored,
    dl.form>dd input[type="password"].is-autocheck-errored,
    dl.form>dd input[type="email"].is-autocheck-errored {
        background-image: url(https://github.com/images/modules/ajax/error@2x.png)
    }
}
.form-cards {
    height: 31px;
    margin: 0 0 15px
}
.form-cards .card {
    float: left;
    width: 47px;
    height: 31px;
    text-indent: -9999px;
    background-image: url(https://github.com/images/modules/pricing/credit-cards-@1x.png);
    background-position: 0 0;
    opacity: 0.6
}
.form-cards .card.visa {
    background-position: 0 0
}
.form-cards .card.amex {
    background-position: -50px 0
}
.form-cards .card.mastercard {
    background-position: -100px 0
}
.form-cards .card.discover {
    background-position: -150px 0
}
.form-cards .card.jcb {
    background-position: -200px 0
}
.form-cards .card.dinersclub {
    background-position: -250px 0
}
.form-cards .card.enabled {
    opacity: 1
}
.form-cards .card.disabled {
    opacity: 0.2
}
.form-cards>.cards {
    margin: 0
}
.form-cards>.cards>li {
    float: left;
    margin: 0 4px 0 0;
    list-style-type: none
}
.form-cards>.cards>li.text {
    font-size: 11px;
    line-height: 31px;
    color: #767676
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    .form-cards>.cards .card {
        background-image: url(https://github.com/images/modules/pricing/credit-cards-@2x.png);
        background-size: 300px 31px
    }
}
.status-indicator-loading {
    position: relative;
    top: 3px;
    width: 16px;
    height: 16px;
    margin-top: -4px;
    background: url(https://github.com/images/spinners/octocat-spinner-32-EAF2F5.gif) 0 0 no-repeat;
    background-size: 16px
}
.inline-form {
    display: inline-block
}
.inline-form .btn-plain {
    background-color: transparent;
    border: 0
}
html.no-dnd-uploads .drag-and-drop {
    min-height: 32px
}
html.no-dnd-uploads .drag-and-drop .default {
    display: none
}
html.no-dnd-uploads .upload-enabled textarea {
    border-bottom: 1px solid #ddd
}
.drag-and-drop {
    padding: 7px 10px;
    margin: 0;
    font-size: 13px;
    line-height: 16px;
    color: #767676;
    background-color: #fafafa;
    border: 1px solid #ccc;
    border-top: 0;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.drag-and-drop .default,
.drag-and-drop .loading,
.drag-and-drop .error {
    display: none
}
.drag-and-drop .error {
    color: #bd2c00
}
.drag-and-drop img {
    vertical-align: top
}
.word-upload-callout {
    position: absolute;
    float: left;
    padding: 0;
    width: 365px;
    bottom: -138px;
    z-index: 9999;
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 3px;
    box-shadow: 0 2px 10px #e2e2e2
}
.word-upload-callout .callout-close {
    position: absolute;
    top: 10px;
    right: 10px
}
.word-upload-callout .callout-text {
    padding: 15px 15px 15px 85px
}
.word-upload-callout .note {
    clear: both;
    padding: 10px;
    border-top: 1px solid #ddd;
    font-size: 12px
}
.word-upload-callout .callout-title {
    margin: 0 0 5px
}
.word-upload-callout .callout-message {
    margin: 0
}
.word-upload-callout .callout-icons {
    padding-top: 15px;
    padding-left: 15px;
    float: left
}
.word-upload-callout .octicon-file-pdf {
    color: #cc0323;
    position: relative;
    top: 5px;
    left: 3px
}
.word-upload-callout .docx {
    color: #174b90
}
.word-upload-callout .pptx {
    color: #cc3b26;
    position: absolute;
    left: 20px;
    top: 48px
}
.word-upload-callout::before,
.word-upload-callout::after {
    border: 10px solid transparent;
    content: " ";
    position: absolute;
    height: 0;
    width: 0
}
.word-upload-callout::before {
    border-bottom-color: #fff;
    position: absolute;
    top: -19px;
    left: 15px;
    z-index: 2
}
.word-upload-callout::after {
    border-bottom-color: #ccc;
    position: absolute;
    top: -20px;
    left: 15px;
    z-index: 1
}
.is-default .drag-and-drop .default {
    display: inline-block
}
.is-uploading .drag-and-drop .loading {
    display: inline-block
}
.is-bad-file .drag-and-drop .bad-file {
    display: inline-block
}
.is-too-big .drag-and-drop .too-big {
    display: inline-block
}
.is-empty .drag-and-drop .empty {
    display: inline-block
}
.is-bad-browser .drag-and-drop .bad-browser {
    display: inline-block
}
.drag-and-drop-error-info {
    font-weight: normal;
    color: #767676
}
.drag-and-drop-error-info a {
    color: #4078c0
}
.is-failed .drag-and-drop .failed-request {
    display: inline-block
}
.manual-file-chooser {
    position: absolute;
    width: 240px;
    padding: 5px;
    margin-left: -80px;
    cursor: pointer;
    opacity: 0.0001
}
.manual-file-chooser:hover+.manual-file-chooser-text {
    text-decoration: underline
}
.btn .manual-file-chooser {
    top: 0;
    padding: 0;
    line-height: 34px
}
.upload-enabled textarea {
    display: block;
    border-bottom: 1px dashed #ddd;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0
}
.focused .drag-and-drop {
    box-shadow: rgba(81, 167, 232, 0.5) 0 0 3px
}
.dragover textarea,
.dragover .drag-and-drop {
    box-shadow: #c9ff00 0 0 3px
}
.previewable-comment-form {
    position: relative
}
.previewable-comment-form .tabnav {
    position: relative;
    padding: 10px 10px 0
}
.previewable-comment-form .comment {
    border: 1px solid #cacaca
}
.previewable-comment-form .comment-form-error {
    margin-bottom: 10px
}
.previewable-comment-form .write-content,
.previewable-comment-form .preview-content {
    display: none;
    padding: 0 0 10px;
    margin: 0 10px
}
.previewable-comment-form.write-selected .write-content,
.previewable-comment-form.preview-selected .preview-content {
    display: block
}
.previewable-comment-form textarea {
    display: block;
    width: 100%;
    min-height: 100px;
    max-height: 500px;
    padding: 10px;
    resize: vertical
}
.previewable-comment-form textarea.fullscreen-contents:focus {
    border: 0;
    box-shadow: none
}
div.composer {
    margin-top: 0;
    border: 0
}
.composer .comment-form-textarea {
    height: 200px;
    min-height: 200px
}
.composer .tabnav {
    margin: 0 0 10px
}
h2.account {
    margin: 15px 0 0;
    font-size: 18px;
    font-weight: normal;
    color: #666
}
p.explain {
    position: relative;
    font-size: 12px;
    color: #666
}
p.explain strong {
    color: #333
}
p.explain .octicon {
    margin-right: 5px;
    color: #bbb
}
p.explain .minibutton {
    top: -4px;
    float: right
}
.callout {
    padding: 10px;
    margin: 15px 0;
    font-size: 13px;
    color: #4c4a42;
    background-color: #fff9ea;
    border-color: #dfd8c2;
    border-radius: 3px
}
.callout strong {
    font-weight: bold;
    color: #000
}
.callout h2 {
    margin: 0;
    font-size: 16px;
    font-weight: 300
}
.callout p:last-child {
    margin-bottom: 0
}
.callout p:first-child {
    margin-top: 0
}
.callout hr {
    width: 100%;
    padding: 0 10px;
    margin: 10px 0 10px -10px;
    background: transparent;
    border-top: 1px solid #e5e2c8;
    border-bottom: 1px solid #fff
}
.infotip {
    padding: 10px;
    margin: 15px 0;
    font-size: 12px;
    color: #4c4a42;
    background-color: #fff9ea;
    border: 1px solid #dfd8c2;
    border-radius: 3px
}
.infotip p {
    margin: 0
}
.infotip p+p {
    margin-top: 15px
}
.dashboard-notice {
    position: relative;
    padding: 15px 15px 15px 55px;
    margin-bottom: 20px;
    font-size: 14px;
    background-color: #fafafa;
    border: solid 1px #d8d8d8;
    border-radius: 3px
}
.dashboard-notice .dismiss {
    position: absolute;
    top: 10px;
    right: 10px;
    width: 16px;
    height: 16px;
    color: #bbb;
    cursor: pointer
}
.dashboard-notice .dismiss:hover {
    color: #666
}
.dashboard-notice .mega-octicon {
    position: absolute;
    top: 15px;
    left: 15px
}
.dashboard-notice .octicon-organization {
    color: #4078c0
}
.dashboard-notice h2 {
    margin-top: 9px;
    margin-bottom: 16px;
    font-size: 18px;
    font-weight: normal;
    color: #000
}
.dashboard-notice p {
    margin-top: 0
}
.dashboard-notice p.no-title {
    padding-right: 5px;
    margin-top: 0
}
.dashboard-notice .inset-figure {
    float: right;
    margin-bottom: 15px;
    margin-left: 20px
}
.dashboard-notice ul {
    margin-left: 18px
}
.dashboard-notice li {
    padding-bottom: 15px
}
.dashboard-notice .coupon {
    padding: 10px;
    margin: 15px 0;
    font-size: 20px;
    font-weight: bold;
    text-align: center;
    background: #fff;
    border: 1px dashed #d1e5ff
}
kbd {
    display: inline-block;
    padding: 3px 5px;
    font: 11px Consolas, "Liberation Mono", Menlo, Courier, monospace;
    line-height: 10px;
    color: #555;
    vertical-align: middle;
    background-color: #fcfcfc;
    border: solid 1px #ccc;
    border-bottom-color: #bbb;
    border-radius: 3px;
    box-shadow: inset 0 -1px 0 #bbb
}
.badmono {
    font-family: sans-serif;
    font-weight: bold
}
.select-menu-button:after {
    display: inline-block;
    width: 0;
    height: 0;
    content: "";
    vertical-align: -2px;
    border: 4px solid;
    border-right-color: transparent;
    border-left-color: transparent;
    border-bottom-color: transparent
}
.select-menu-button.icon-only {
    padding-left: 7px
}
.select-menu-button.primary:after {
    border-top-color: #fff
}
.select-menu-button.primary:after:active {
    background-color: #4a993e
}
.select-menu .spinner {
    float: left;
    margin: 4px 0 0 -24px
}
.select-menu.active .select-menu-modal-holder {
    display: block
}
.select-menu.select-menu-modal-right {
    position: relative
}
.select-menu.select-menu-modal-right .select-menu-modal-holder {
    right: 0
}
.select-menu .select-menu-clear-item {
    display: block
}
.select-menu .select-menu-clear-item .octicon {
    color: inherit
}
.select-menu .select-menu-clear-item+.select-menu-no-results {
    display: none !important
}
.select-menu.is-loading .select-menu-loading-overlay {
    display: block
}
.select-menu.is-loading .select-menu-modal {
    min-height: 200px
}
.select-menu-loading-overlay {
    display: none;
    text-indent: 100%;
    height: 100%;
    width: 100%;
    position: absolute;
    top: 0;
    z-index: 5;
    border-radius: 5px;
    border: 1px solid transparent;
    background-color: rgba(255, 255, 255, 0.8);
    -webkit-animation: pulse 2s infinite linear;
    animation: pulse 2s infinite linear
}
.select-menu-loading-overlay:before {
    position: absolute;
    left: 50%;
    top: 50%;
    margin: -16px 0 0 -16px;
    width: 32px;
    content: "\f008";
    font: normal normal 32px/1 "octicons";
    display: inline-block;
    text-decoration: none;
    -webkit-font-smoothing: antialiased;
    text-indent: 0
}
@-webkit-keyframes pulse {
    0% {
        color: rgba(170, 170, 170, 0.1)
    }
    10% {
        color: #aaaaaa
    }
    100% {
        color: rgba(170, 170, 170, 0.1)
    }
}
@keyframes pulse {
    0% {
        color: rgba(170, 170, 170, 0.1)
    }
    10% {
        color: #aaaaaa
    }
    100% {
        color: rgba(170, 170, 170, 0.1)
    }
}
.select-menu-modal-holder {
    position: absolute;
    display: none;
    z-index: 21
}
.select-menu-modal {
    position: relative;
    width: 300px;
    margin-top: 4px;
    margin-bottom: 20px;
    overflow: hidden;
    font-size: 12px;
    color: #666;
    background-color: #fff;
    background-clip: padding-box;
    border: 1px solid rgba(200, 200, 200, 0.4);
    border-radius: 3px;
    box-shadow: 0 3px 12px rgba(0, 0, 0, 0.15)
}
.select-menu-header {
    padding: 8px 10px;
    background: #f5f5f5;
    border-bottom: 1px solid rgba(200, 200, 200, 0.4)
}
.select-menu-header .select-menu-title {
    font-weight: bold;
    color: #333;
    text-shadow: 0 1px 0 #fff
}
.select-menu-header .octicon {
    display: block;
    float: right;
    color: #ccc;
    cursor: pointer
}
.select-menu-header .octicon:hover {
    color: #555
}
.select-menu-filters {
    background-color: #f8f8f8
}
.select-menu-text-filter {
    padding: 10px 10px 0
}
.select-menu-text-filter:first-child:last-child {
    padding-bottom: 10px;
    border-bottom: 1px solid #ddd
}
.select-menu-text-filter input {
    display: block;
    width: 100%;
    max-width: 100%;
    padding: 5px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.select-menu-text-filter input::-webkit-input-placeholder {
    color: #aaa
}
.select-menu-text-filter input::-moz-placeholder {
    color: #aaa
}
.select-menu-text-filter input:-ms-input-placeholder {
    color: #aaa
}
.select-menu-text-filter input::placeholder {
    color: #aaa
}
.select-menu-tabs {
    height: 33px;
    border-bottom: 1px solid #ddd
}
.select-menu-tabs ul {
    overflow: hidden;
    padding: 14px 10px 0
}
.select-menu-tabs .select-menu-tab {
    display: inline-block
}
.select-menu-tabs a {
    height: 20px;
    padding: 5px 8px;
    font-size: 11px;
    font-weight: bold;
    color: #888;
    text-decoration: none;
    line-height: 20px;
    border-radius: 3px 3px 0 0;
    cursor: pointer
}
.select-menu-tabs a:hover {
    color: #333
}
.select-menu-tabs a.selected {
    padding: 4px 5px;
    border: 1px solid #ddd;
    border-bottom: 1px solid #fff;
    background-color: #fff;
    color: #333
}
.select-menu-list {
    max-height: 400px;
    overflow: auto;
    position: relative;
    line-height: 1.4
}
.select-menu-list.select-menu-tab-bucket {
    display: none
}
.select-menu-list.select-menu-tab-bucket.selected {
    display: block
}
.select-menu-list.is-showing-new-item-form .select-menu-new-item-form {
    display: table
}
.select-menu-list.is-showing-new-item-form .select-menu-no-results,
.select-menu-list.is-showing-new-item-form .select-menu-clear-item {
    display: none
}
.select-menu-item {
    cursor: pointer;
    border-bottom: 1px solid #eee;
    display: table;
    table-layout: fixed;
    width: 100%;
    overflow: hidden;
    color: inherit
}
.select-menu-item:hover {
    text-decoration: none
}
.select-menu-item.select-menu-item-template {
    display: none
}
.select-menu-item.disabled,
.select-menu-item.disabled.selected {
    color: #767676
}
.select-menu-item.disabled .select-menu-item-gravatar,
.select-menu-item.disabled.selected .select-menu-item-gravatar {
    opacity: 0.5
}
.select-menu-item .octicon {
    vertical-align: middle
}
.select-menu-item .octicon-check {
    visibility: hidden
}
.select-menu-item input[type="radio"] {
    display: none
}
.select-menu-item .select-menu-item-icon {
    display: table-cell;
    color: transparent;
    vertical-align: top;
    padding: 8px 0 8px 8px;
    width: 24px;
    text-align: center
}
.select-menu-item.navigation-focus,
.select-menu-item.navigation-focus.selected,
.select-menu-item.navigation-focus.select-menu-action,
.select-menu-item.navigation-focus .description-inline {
    background-color: #4078c0;
    color: #fff
}
.select-menu-item.navigation-focus>.octicon,
.select-menu-item.navigation-focus.selected>.octicon,
.select-menu-item.navigation-focus.select-menu-action>.octicon,
.select-menu-item.navigation-focus .description-inline>.octicon {
    color: #fff
}
.select-menu-item.navigation-focus .text-danger,
.select-menu-item.navigation-focus .description,
.select-menu-item.navigation-focus.selected .text-danger,
.select-menu-item.navigation-focus.selected .description,
.select-menu-item.navigation-focus.select-menu-action .text-danger,
.select-menu-item.navigation-focus.select-menu-action .description,
.select-menu-item.navigation-focus .description-inline .text-danger,
.select-menu-item.navigation-focus .description-inline .description {
    color: #fff
}
.select-menu-item>.octicon-dash {
    display: none
}
.select-menu-item.indeterminate>.octicon-check {
    display: none
}
.select-menu-item.indeterminate>.octicon-dash {
    display: table-cell
}
.select-menu-item.select-menu-action,
.select-menu-item.selected {
    color: #333
}
.select-menu-item.select-menu-action .description,
.select-menu-item.selected .description {
    color: #666
}
.select-menu-item.select-menu-action .octicon-check,
.select-menu-item.selected .octicon-check {
    visibility: visible
}
.select-menu-item.select-menu-action>.octicon,
.select-menu-item.selected>.octicon {
    color: #333
}
.select-menu-item.select-menu-action .select-menu-item-text {
    font-weight: bold
}
.select-menu[data-multiple] .select-menu-item:active {
    background-color: transparent !important
}
.select-menu-item a {
    color: inherit;
    text-decoration: none
}
.select-menu-item .hidden-select-button-text {
    display: none
}
.select-menu-item .css-truncate-target {
    display: table-cell;
    max-width: 100%
}
form.select-menu-item>div:first-child {
    display: none !important
}
.select-menu-item.last-visible,
.select-menu-list:last-child .select-menu-item:last-child {
    border-bottom: 0;
    border-radius: 0 0 3px 3px
}
.select-menu-actions .select-menu-item:hover {
    background-color: #4078c0;
    color: #fff
}
.select-menu-actions .select-menu-item:hover>.octicon {
    color: #fff
}
.select-menu-actions .select-menu-item:hover .description {
    color: #fff
}
.select-menu-no-results {
    padding: 9px;
    display: none;
    cursor: auto;
    color: #767676
}
.select-menu-list.filterable-empty .select-menu-no-results,
.select-menu-no-results:only-child {
    display: block
}
.select-menu-button-gravatar,
.select-menu-item-gravatar {
    overflow: hidden;
    line-height: 0;
    width: 20px
}
.select-menu-button-gravatar img,
.select-menu-item-gravatar img {
    height: 20px;
    width: 20px;
    display: inline-block;
    border-radius: 3px
}
.select-menu-item-gravatar {
    display: table-cell;
    padding: 6px 0 6px 8px;
    vertical-align: top;
    width: 28px
}
.select-menu-button-gravatar {
    float: left;
    margin-right: 5px
}
.select-menu-item-text {
    display: table-cell;
    vertical-align: top;
    padding: 8px;
    text-align: left
}
.select-menu-item-text:first-child {
    margin-left: 5px
}
.select-menu-item-text .description {
    color: #767676;
    font-size: 12px;
    max-width: 265px;
    display: block;
    margin-top: 3px
}
.select-menu-item-text .description-inline {
    color: #767676;
    font-size: 10px
}
.select-menu-item-heading {
    display: block;
    margin-top: 0;
    margin-bottom: 0;
    font-size: 14px;
    font-weight: bold;
    line-height: 1.1
}
.select-menu-item-heading .description {
    font-weight: normal;
    display: inline
}
.select-menu-footer {
    padding: 8px;
    font-weight: bold;
    border-top: 1px solid #eee
}
.select-menu-footer a {
    display: inline-block;
    margin-top: 1px;
    vertical-align: top
}
.select-menu-footer .octicon {
    color: #666
}
.select-menu-new-item-form {
    display: none
}
.select-menu-new-item-form .octicon {
    color: #4078c0
}
.modal-backdrop {
    display: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
body.menu-active .modal-backdrop {
    display: block;
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    z-index: 20
}
.dropdown {
    position: relative
}
.dropdown-caret {
    display: inline-block;
    width: 0;
    height: 0;
    content: "";
    vertical-align: -2px;
    border: 4px solid;
    border-right-color: transparent;
    border-left-color: transparent;
    border-bottom-color: transparent
}
.dropdown-menu {
    position: absolute;
    top: 100%;
    left: 0;
    z-index: 100;
    width: 160px;
    margin-top: 2px;
    padding-top: 5px;
    padding-bottom: 5px;
    background-color: #fff;
    background-clip: padding-box;
    border: 1px solid rgba(0, 0, 0, 0.15);
    border-radius: 4px;
    box-shadow: 0 3px 12px rgba(0, 0, 0, 0.15)
}
.dropdown-menu:before {
    position: absolute;
    display: inline-block;
    content: "";
    border: 8px solid transparent;
    border-bottom-color: rgba(0, 0, 0, 0.15)
}
.dropdown-menu:after {
    position: absolute;
    display: inline-block;
    content: "";
    border: 7px solid transparent;
    border-bottom-color: #fff
}
.dropdown-item {
    display: block;
    padding: 4px 10px 4px 15px;
    color: #333;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis
}
.dropdown-item:hover,
.dropdown-item.zeroclipboard-is-hover {
    color: #fff;
    text-decoration: none;
    background-color: #4078c0
}
.dropdown-item:hover>.octicon,
.dropdown-item.zeroclipboard-is-hover>.octicon {
    color: inherit;
    opacity: 1
}
.dropdown-signout {
    width: 100%;
    text-align: left;
    background: none;
    border: 0
}
.dropdown-divider {
    height: 1px;
    margin: 8px 1px;
    background-color: #e5e5e5
}
.dropdown-header {
    padding: 4px 15px;
    font-size: 12px;
    color: #767676
}
.dropdown-menu-content {
    display: none
}
.dropdown-menu-w {
    left: auto;
    right: 100%;
    width: auto;
    margin-top: 0;
    margin-right: 10px
}
.dropdown-menu-w:before {
    top: 10px;
    right: -16px;
    left: auto;
    border-color: transparent;
    border-left-color: rgba(0, 0, 0, 0.15)
}
.dropdown-menu-w:after {
    top: 11px;
    right: -14px;
    left: auto;
    border-color: transparent;
    border-left-color: #fff
}
.dropdown-menu-e {
    left: 100%;
    width: auto;
    margin-top: 0;
    margin-left: 10px
}
.dropdown-menu-e:before {
    top: 10px;
    left: -16px;
    border-color: transparent;
    border-right-color: rgba(0, 0, 0, 0.15)
}
.dropdown-menu-e:after {
    top: 11px;
    left: -14px;
    border-color: transparent;
    border-right-color: #fff
}
.dropdown-menu-sw {
    left: auto;
    right: 0
}
.dropdown-menu-sw:before {
    top: -16px;
    left: auto;
    right: 9px
}
.dropdown-menu-sw:after {
    top: -14px;
    left: auto;
    right: 10px
}
.dropdown-menu-se:before {
    top: -16px;
    left: 9px
}
.dropdown-menu-se:after {
    top: -14px;
    left: 10px
}
include-fragment,
poll-include-fragment {
    display: block
}
.pagination:before {
    display: table;
    content: ""
}
.pagination:after {
    display: table;
    clear: both;
    content: ""
}
.pagination a,
.pagination span,
.pagination em {
    position: relative;
    float: left;
    margin-left: -1px;
    font-size: 13px;
    font-weight: bold;
    font-style: normal;
    padding: 7px 12px;
    color: #4078c0;
    white-space: nowrap;
    vertical-align: middle;
    cursor: pointer;
    background: #fff;
    border: 1px solid #e5e5e5;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
.pagination a:first-child,
.pagination span:first-child,
.pagination em:first-child {
    margin-left: 0;
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px
}
.pagination a:last-child,
.pagination span:last-child,
.pagination em:last-child {
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px
}
.pagination a:hover,
.pagination a:focus,
.pagination span:hover,
.pagination span:focus,
.pagination em:hover,
.pagination em:focus {
    z-index: 2;
    background-color: #e7e7e7;
    border-color: #e5e5e5;
    text-decoration: none
}
.pagination .selected {
    z-index: 3
}
.pagination .current,
.pagination .current:hover {
    z-index: 3;
    color: #fff;
    background-color: #4078c0;
    border-color: #4078c0
}
.pagination .gap,
.pagination .disabled,
.pagination .gap:hover,
.pagination .disabled:hover {
    background-color: #fafafa;
    cursor: default;
    color: #d3d3d3
}
.ajax_paginate {
    display: block;
    margin-top: 20px
}
.ajax_paginate a {
    float: none;
    display: block;
    padding: 6px;
    text-align: center
}
.ajax_paginate.loading a {
    text-indent: -3000px;
    background-color: #eaeaea;
    background-image: url(https://github.com/images/spinners/octocat-spinner-16px-EAF2F5.gif);
    background-repeat: no-repeat;
    background-position: center center;
    border-color: #c5c5c5
}
@media screen and (-webkit-min-device-pixel-ratio: 2),
screen and (max--moz-device-pixel-ratio: 2) {
    .ajax_paginate.loading a {
        background-image: url(https://github.com/images/spinners/octocat-spinner-32-EAF2F5.gif);
        background-size: 16px auto
    }
}
.paginate-container {
    margin-top: 20px;
    margin-bottom: 15px;
    text-align: center
}
.paginate-container .pagination {
    display: inline-block
}
.tab-size[data-tab-size="1"] {
    -moz-tab-size: 1;
    -o-tab-size: 1;
    tab-size: 1
}
.tab-size[data-tab-size="2"] {
    -moz-tab-size: 2;
    -o-tab-size: 2;
    tab-size: 2
}
.tab-size[data-tab-size="3"] {
    -moz-tab-size: 3;
    -o-tab-size: 3;
    tab-size: 3
}
.tab-size[data-tab-size="4"] {
    -moz-tab-size: 4;
    -o-tab-size: 4;
    tab-size: 4
}
.tab-size[data-tab-size="5"] {
    -moz-tab-size: 5;
    -o-tab-size: 5;
    tab-size: 5
}
.tab-size[data-tab-size="6"] {
    -moz-tab-size: 6;
    -o-tab-size: 6;
    tab-size: 6
}
.tab-size[data-tab-size="7"] {
    -moz-tab-size: 7;
    -o-tab-size: 7;
    tab-size: 7
}
.tab-size[data-tab-size="8"] {
    -moz-tab-size: 8;
    -o-tab-size: 8;
    tab-size: 8
}
.tab-size[data-tab-size="9"] {
    -moz-tab-size: 9;
    -o-tab-size: 9;
    tab-size: 9
}
.tab-size[data-tab-size="10"] {
    -moz-tab-size: 10;
    -o-tab-size: 10;
    tab-size: 10
}
.tab-size[data-tab-size="11"] {
    -moz-tab-size: 11;
    -o-tab-size: 11;
    tab-size: 11
}
.tab-size[data-tab-size="12"] {
    -moz-tab-size: 12;
    -o-tab-size: 12;
    tab-size: 12
}
.header {
    padding-top: 10px;
    padding-bottom: 10px;
    min-width: 1000px;
    background-color: #f5f5f5;
    border-bottom: 1px solid #e5e5e5
}
.header-logged-out {
    padding-top: 15px;
    padding-bottom: 15px
}
.read-only-mode-banner {
    background-color: #f8e45f;
    border-bottom-color: #f6dc2e;
    text-align: center
}
.header-logo-invertocat {
    float: left;
    margin-right: 10px;
    margin-left: -2px;
    color: #333;
    white-space: nowrap
}
.header-logo-invertocat .octicon-mark-github {
    float: left;
    width: 28px;
    height: 28px;
    font-size: 28px
}
.header-logo-invertocat:hover {
    color: #4078c0;
    text-decoration: none
}
.logo-subbrand {
    float: left;
    margin-left: 6px;
    font-size: 16px;
    font-weight: bold;
    line-height: 28px
}
.header-logo-wordmark {
    position: relative;
    float: left;
    height: 26px;
    margin-right: 15px;
    color: #333
}
.header-logo-wordmark:hover {
    color: #4078c0
}
.header-logo-wordmark .octicon-logo-github {
    height: 26px;
    font-size: 32px
}
.notification-indicator .mail-status {
    background-image: -webkit-linear-gradient(#7aa1d3, #4078c0);
    background-image: linear-gradient(#7aa1d3, #4078c0);
    position: absolute;
    top: -2px;
    right: 2px;
    z-index: 2;
    display: none;
    width: 14px;
    height: 14px;
    color: #fff;
    text-align: center;
    background-clip: padding-box;
    border-radius: 50%;
    border: 2px solid #f3f3f3
}
.notification-indicator .mail-status.unread {
    display: inline-block
}
.notification-indicator:hover .mail-status {
    background-color: #4078c0
}
.site-search {
    position: relative;
    float: left
}
.site-search form {
    position: relative;
    float: left;
    width: 360px;
    margin-right: 10px
}
.site-search .form-control {
    position: relative;
    width: 100%;
    min-height: 26px;
    padding: 2px;
    font-size: 12px;
    display: block;
    line-height: 17px
}
.site-search .chromeless-input {
    background: none;
    border: 0;
    box-shadow: none;
    min-height: 22px;
    display: inline-block;
    font-size: 12px;
    padding: 3px 5px;
    line-height: 16px;
    width: 250px
}
.site-search .chromeless-input:focus {
    border: 0;
    box-shadow: none
}
.site-search .chromeless-input::-ms-clear {
    display: none
}
.site-search .scope-badge {
    display: none;
    padding: 0 5px;
    line-height: 22px;
    font-size: 12px;
    font-weight: normal;
    color: #767676;
    background-color: #eee;
    border-radius: 2px;
    vertical-align: middle
}
.site-search.repo-scope .scope-badge {
    display: inline-block
}
.site-search.repo-scope .form-control.focus .scope-badge {
    background-color: #e1eaf5;
    color: #4078c0
}
.header-nav {
    list-style: none
}
.header-nav-item {
    float: left
}
.header-nav-item.active .dropdown-menu-content {
    display: block
}
.header-nav-item.active .tooltipped:before,
.header-nav-item.active .tooltipped:after {
    display: none
}
.header-nav-item .dropdown-menu {
    margin-top: 13px;
    width: 180px
}
.header-nav-link {
    display: block;
    padding: 4px 8px;
    font-size: 13px;
    font-weight: bold;
    line-height: 20px;
    color: #333
}
.header-nav-link:hover,
.header-nav-link:focus {
    color: #4078c0;
    text-decoration: none
}
.header-nav-link:hover .dropdown-caret,
.header-nav-link:focus .dropdown-caret {
    border-top-color: #4078c0
}
.user-nav {
    margin-right: -8px
}
.user-nav .header-nav-link {
    height: 28px
}
.user-nav .octicon {
    width: 18px;
    height: 18px;
    text-align: center
}
.user-nav .octicon-inbox {
    font-size: 20px
}
.user-nav .octicon-plus {
    float: left;
    width: 16px;
    height: 18px;
    font-size: 18px
}
.user-nav .avatar {
    float: left;
    margin-right: 5px
}
.header-nav-current-user {
    padding-bottom: 0;
    font-size: inherit
}
.header-nav-current-user .css-truncate-target {
    max-width: 100%
}
.header-actions {
    float: right;
    margin-top: -3px;
    margin-bottom: -3px
}
.header-actions .btn {
    margin-left: 5px
}
.enterprise .header {
    background-color: #2a2c2e;
    border-bottom-color: #121213
}
.is-stats .enterprise .header {
    box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.05)
}
.enterprise .header-logo-wordmark,
.enterprise .header-logo-invertocat,
.enterprise .header-nav-link {
    color: #c8c8ca
}
.enterprise .header-logo-wordmark:hover,
.enterprise .header-logo-wordmark:focus,
.enterprise .header-logo-invertocat:hover,
.enterprise .header-logo-invertocat:focus,
.enterprise .header-nav-link:hover,
.enterprise .header-nav-link:focus {
    color: #fafafa
}
.enterprise .header-nav-link:hover .dropdown-caret,
.enterprise .header-nav-link:focus .dropdown-caret {
    border-top-color: #fafafa
}
.enterprise .notification-indicator .mail-status {
    border-color: #2a2c2e
}
.enterprise .notification-indicator:hover .mail-status {
    background-color: #d26911
}
.enterprise .header-actions .btn {
    border: 0;
    box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.25), 0 1px 1px rgba(0, 0, 0, 0.5)
}
.enterprise .site-search .scope-badge {
    background-color: #5b5f63;
    color: #c8c8ca
}
.enterprise .site-search .form-control {
    color: #fafafa;
    background-color: #4f5256;
    border-color: #121213
}
.enterprise .site-search .form-control.focus {
    background-color: #55595d;
    border-color: #000;
    box-shadow: inset 0 1px 0 rgba(0, 0, 0, 0.075)
}
.enterprise .site-search .form-control.focus .scope-badge {
    background-color: #676c71;
    color: #fff
}
.enterprise .chromeless-input {
    color: #fff
}
.unsupported-browser {
    padding: 15px 0;
    color: #211e14;
    background-image: -webkit-linear-gradient(#feefae, #fae692);
    background-image: linear-gradient(#feefae, #fae692);
    border-bottom: 1px solid #b3a569
}
.unsupported-browser .container {
    background: url(https://github.com/images/icons/ie-notice.png) no-repeat 0 5px
}
.unsupported-browser h5 {
    font-size: 13px;
    margin: 5px 0 2px;
    padding-left: 48px
}
.unsupported-browser p {
    margin: 0;
    padding-left: 48px
}
.unsupported-browser .btn {
    float: right;
    margin-top: 5px;
    margin-left: 8px
}
.mobile-banner button.switch-to-mobile {
    display: block;
    width: 100%;
    padding: 30px 0 45px;
    border: 0;
    background-color: #444;
    color: #eaeaea;
    font-size: 60px;
    text-align: center;
    font-weight: bold
}
.mobile-banner button.switch-to-mobile .mega-octicon {
    position: relative;
    top: -8px;
    margin-right: 5px;
    color: #ddd;
    font-size: 48px
}
.accessibility-aid {
    height: 1px;
    width: 1px;
    clip: rect(1px, 1px, 1px, 1px);
    overflow: hidden;
    position: absolute;
    margin: 0
}
.accessibility-aid:focus {
    height: auto;
    width: auto;
    clip: auto;
    z-index: 1;
    top: 0;
    line-height: 49px;
    padding: 0 10px;
    background: #f5f5f5;
    font-weight: bold;
    text-decoration: none;
    color: #333
}
.is-stats .accessibility-aid:focus {
    top: 34px
}
.site-footer {
    position: relative;
    margin-top: 40px;
    padding-top: 40px;
    padding-bottom: 40px;
    font-size: 12px;
    line-height: 1.5;
    color: #767676;
    border-top: 1px solid #eee
}
.site-footer:before {
    display: table;
    content: ""
}
.site-footer:after {
    display: table;
    clear: both;
    content: ""
}
.site-footer .octicon-mark-github {
    position: absolute;
    top: 38px;
    left: 50%;
    height: 24px;
    width: 24px;
    margin-left: -12px;
    font-size: 24px;
    color: #ccc
}
.site-footer .octicon-mark-github:hover {
    color: #bbb
}
.site-footer-links {
    margin: 0;
    list-style: none
}
.site-footer-links li {
    display: inline-block;
    line-height: 16px
}
.site-footer-links li+li {
    margin-left: 10px
}
.billing-plans tbody td {
    width: 25%;
    vertical-align: middle
}
.billing-plans .current {
    background-color: #f2ffed
}
.billing-plans .name {
    font-size: 14px;
    font-weight: bold;
    color: #333
}
.billing-plans .coupon {
    font-size: 12px
}
.billing-plans .coupon td {
    color: #fff;
    background-color: #6cc644
}
.billing-plans .coupon .text-right {
    white-space: nowrap
}
.billing-plans .coupon.expiring td {
    background-color: #df6e00
}
.billing-plans .coupon.expiring .coupon-label:after {
    border-bottom-color: #df6e00
}
.billing-plans tbody>.selected {
    background-color: #fdffce
}
.coupon-label {
    position: relative;
    padding: 9px;
    margin: -9px
}
.coupon-label:after {
    position: absolute;
    bottom: 100%;
    left: 15px;
    width: 0;
    height: 0;
    pointer-events: none;
    content: " ";
    border: solid transparent;
    border-width: 5px;
    border-bottom-color: #6cc644
}
.boxed-group-table .toggle-currency {
    font-size: 11px;
    font-weight: normal
}
.is-hidden,
.has-removed-contents {
    display: none
}
.currency-notice {
    margin-bottom: 10px
}
.org-login {
    margin-top: -30px;
    margin-bottom: 30px
}
.org-login img {
    width: 450px;
    padding: 1px;
    margin: 10px -25px;
    border: 1px solid #ccc
}
.plan-notice {
    border-top: 1px solid #eee;
    margin-bottom: 0;
    padding: 10px
}
.auth-form {
    width: 400px;
    margin: 60px auto
}
.auth-form .note {
    margin: 15px 0;
    text-align: center
}
.auth-form-header {
    position: relative;
    padding: 10px 20px;
    margin: 0;
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.3);
    background-color: #829aa8;
    border: 1px solid #768995;
    border-radius: 3px 3px 0 0
}
.auth-form-header h1 {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 16px
}
.auth-form-header h1 a {
    color: #fff
}
.auth-form-header .octicon {
    position: absolute;
    top: 10px;
    right: 20px;
    color: rgba(0, 0, 0, 0.4);
    text-shadow: 0 1px 0 rgba(255, 255, 255, 0.1)
}
.auth-form-body {
    padding: 20px;
    font-size: 14px;
    background-color: #fff;
    border: 1px solid #d8dee2;
    border-top: 0;
    border-radius: 0 0 3px 3px
}
.auth-form-body .input-block {
    margin-top: 5px;
    margin-bottom: 15px
}
.auth-form-body p {
    margin: 0 0 10px
}
.two-factor-help {
    position: relative;
    padding: 10px 10px 10px 36px;
    margin: 60px 0 auto auto;
    border: 1px solid #eaeaea;
    border-radius: 3px
}
.two-factor-help .octicon-device-mobile {
    position: absolute;
    top: 10px;
    left: 10px
}
.two-factor-help .octicon-key {
    position: absolute;
    left: 10px
}
.two-factor-help .btn-sm {
    float: right
}
.two-factor-help ul {
    list-style-type: none
}
.u2f-login-spinner {
    float: left;
    margin-right: 5px
}
.flash.sms-error,
.flash.sms-success {
    display: none;
    margin: 0 0 10px
}
.is-sent .sms-success {
    display: block
}
.is-sent .sms-error {
    display: none
}
.is-not-sent .sms-success {
    display: none
}
.is-not-sent .sms-error {
    display: block
}
.autocomplete-results {
    position: absolute;
    z-index: 99;
    display: none;
    max-height: 20em;
    overflow-y: auto;
    font-size: 13px;
    list-style: none;
    background: #fff;
    border: 1px solid #c1c1c1;
    border-radius: 3px;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.3)
}
.autocomplete-results .no-results {
    display: none
}
.autocomplete-group {
    width: 100%;
    overflow: hidden
}
.autocomplete-item {
    display: block;
    padding: 5px;
    overflow: hidden;
    font-weight: bold;
    text-decoration: none;
    text-overflow: ellipsis;
    white-space: nowrap;
    cursor: pointer
}
.autocomplete-item.selected,
.autocomplete-item.navigation-focus {
    color: #fff;
    text-decoration: none;
    background-color: #4078c0
}
.autocomplete-item.selected .organization-member,
.autocomplete-item.selected .ldap-group-dn,
.autocomplete-item.navigation-focus .organization-member,
.autocomplete-item.navigation-focus .ldap-group-dn {
    color: #f2f2f2
}
.autocomplete-item .secondary-label {
    font-weight: normal
}
.autocomplete-item .organization-member {
    float: right;
    padding-top: 1px;
    color: #808080
}
.suggester-container {
    position: absolute;
    top: 0;
    left: 0;
    z-index: 30
}
.suggester {
    position: relative;
    top: 0;
    left: 0;
    display: none;
    min-width: 180px;
    margin-top: 20px;
    background: #fff;
    border: 1px solid #ddd;
    border-radius: 3px;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.1)
}
.suggester.active {
    display: block
}
.suggester ul {
    padding: 0;
    margin: 0;
    list-style: none
}
.suggester li {
    display: block;
    padding: 5px 10px;
    font-weight: bold;
    border-bottom: 1px solid #ddd
}
.suggester li small {
    font-weight: normal;
    color: #767676
}
.suggester li:last-child {
    border-bottom: 0;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.suggester li:first-child a {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.suggester li.navigation-focus {
    color: #fff;
    text-decoration: none;
    background: #4078c0
}
.suggester li.navigation-focus small {
    color: #fff
}
.breadcrumb {
    margin-bottom: 10px;
    font-size: 18px;
    color: #767676
}
.breadcrumb .separator:before,
.breadcrumb .separator:after {
    content: " "
}
.breadcrumb strong.final-path {
    color: #000
}
.breadcrumb .zeroclipboard-button {
    display: inline-block;
    margin-left: 5px
}
.breadcrumb .repo-root {
    font-weight: bold
}
.editor-license-template,
.editor-gitignore-template {
    position: relative;
    top: 3px;
    display: none;
    float: right;
    font-size: 14px
}
.editor-license-template.is-visible,
.editor-gitignore-template.is-visible {
    display: block
}
.editor-license-template .select-menu-git-ignore,
.editor-license-template .select-menu-license-picker,
.editor-gitignore-template .select-menu-git-ignore,
.editor-gitignore-template .select-menu-license-picker {
    right: 0
}
.editor-abort {
    display: inline;
    font-size: 14px
}
.blob-interaction-bar {
    position: relative;
    background-color: #f2f2f2;
    border-bottom: 1px solid #e5e5e5
}
.blob-interaction-bar:before {
    display: table;
    content: ""
}
.blob-interaction-bar:after {
    display: table;
    clear: both;
    content: ""
}
.blob-interaction-bar .octicon-search {
    position: absolute;
    top: 10px;
    left: 10px;
    font-size: 12px;
    color: #767676
}
input.blob-filter {
    padding: 4px 20px 5px 30px;
    width: 100%;
    font-size: 12px;
    border: 0;
    border-radius: 0;
    outline: none
}
input.blob-filter:focus {
    outline: none
}
.markdown-body .csv-data td,
.markdown-body .csv-data th {
    padding: 5px;
    overflow: hidden;
    font-size: 12px;
    line-height: 1;
    text-align: left;
    white-space: nowrap
}
.markdown-body .csv-data .blob-num {
    padding: 10px 8px 9px;
    text-align: right;
    background: #fff;
    border: 0
}
.markdown-body .csv-data tr {
    border-top: 0
}
.markdown-body .csv-data th {
    font-weight: bold;
    background: #f8f8f8;
    border-top: 0
}
.too-long-message {
    display: none;
    color: #cea61b
}
.is-too-long-error .too-long-message {
    display: block
}
.check-for-fork {
    display: inline-block
}
.check-for-fork img {
    vertical-align: text-bottom
}
.check-for-fork.is-error .check-for-fork-loading {
    display: none
}
.check-for-fork.is-error .check-for-fork-error {
    display: inline-block
}
.check-for-fork-error {
    display: none
}
.file-commit-form {
    padding-left: 64px
}
.file-commit-form .commit-form-avatar {
    float: left;
    margin-left: -64px;
    border-radius: 4px
}
.file-commit-form .commit-form {
    position: relative;
    padding: 15px;
    margin-bottom: 10px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.file-commit-form .commit-form:after,
.file-commit-form .commit-form:before {
    position: absolute;
    top: 11px;
    left: -16px;
    right: 100%;
    width: 0;
    height: 0;
    display: block;
    content: " ";
    border-color: transparent;
    border-style: solid solid outset;
    pointer-events: none
}
.file-commit-form .commit-form:after {
    border-width: 7px;
    border-right-color: #fff;
    margin-top: 1px;
    margin-left: 2px
}
.file-commit-form .commit-form:before {
    border-right-color: #ddd;
    border-width: 8px
}
.file-commit-form .commit-message {
    min-height: 100px
}
.file-commit-form-heading {
    margin-top: 0;
    margin-bottom: 10px
}
.quick-pull-choice .form-checkbox {
    padding-left: 25px;
    margin: 10px 0
}
.quick-pull-choice .form-checkbox label {
    font-weight: normal
}
.quick-pull-choice .form-checkbox .octicon {
    width: 16px;
    margin-right: 3px;
    text-align: center
}
.quick-pull-choice .form-checkbox .icon-shield {
    width: 16px;
    margin-top: 1px;
    margin-right: 3px;
    vertical-align: text-bottom
}
.quick-pull-choice dl.form,
.quick-pull-choice .form-checkbox:last-child {
    margin-bottom: 0
}
.quick-pull-choice .quick-pull-branch-name {
    display: none;
    padding-left: 48px;
    margin-top: 5px
}
.quick-pull-choice .new-branch-name-input {
    position: relative;
    margin-top: 5px
}
.quick-pull-choice .new-branch-name-input input {
    width: 240px;
    padding-left: 26px;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.quick-pull-choice .new-branch-name-input .quick-pull-new-branch-icon {
    position: absolute;
    top: 9px;
    left: 10px;
    color: #b0c4ce
}
.quick-pull-choice.will-normalize-ref .quick-pull-normalization-info {
    display: inline-block
}
.quick-pull-choice.will-create-branch .quick-pull-branch-name {
    display: inline-block
}
.quick-pull-normalization-info {
    position: absolute;
    top: 34px;
    left: 0;
    z-index: 10;
    display: none;
    padding: 5px;
    font-size: 11px;
    color: #494620;
    background: #f7ea57;
    border: 1px solid #c0b536;
    border-top-color: #fff;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.merge-pr {
    padding-top: 10px;
    margin: 20px 0 0;
    border-top: 1px solid #ddd
}
.merge-pr.open .merge-branch-form {
    display: block
}
.merge-pr.open .branch-action {
    display: none
}
.status-heading {
    margin-top: 0;
    margin-bottom: 1px
}
.build-statuses-list {
    max-height: 0;
    padding: 0;
    margin: 15px -15px -16px -55px;
    overflow-y: auto;
    border: solid #eee;
    border-width: 1px 0 0;
    -webkit-transition: max-height 0.25s ease-in-out;
    transition: max-height 0.25s ease-in-out
}
.statuses-toggle-opened {
    display: none
}
.build-status-item {
    padding: 10px 15px 10px 53px;
    background-color: #fafafa;
    border-bottom: 1px solid #eee
}
.build-status-item:last-child {
    border-bottom: 0
}
.build-status-item .css-truncate-target {
    width: 75%;
    max-width: 75%
}
.status-meta {
    color: #767676
}
.branch-action-item-icon {
    float: left;
    margin-left: -40px
}
.build-status-icon {
    width: 16px;
    text-align: center
}
.build-status-details {
    margin-left: 10px
}
.merge-pr-more-commits {
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 64px;
    font-size: 12px;
    color: #767676
}
.branch-action {
    padding-left: 64px;
    margin-top: 15px;
    margin-bottom: 15px
}
.branch-action .merge-branch-heading {
    margin-bottom: 2px
}
.branch-action .delete-branch-failure {
    display: none;
    padding: 15px
}
.branch-action.error .delete-branch-failure {
    display: block
}
.branch-action.error .merge-message {
    display: none
}
.branch-action-icon {
    float: left;
    width: 48px;
    height: 48px;
    margin-left: -64px;
    line-height: 48px;
    color: #fff;
    text-align: center;
    border-radius: 3px
}
.branch-action-body {
    position: relative;
    background-color: #fff;
    border: 1px solid #ddd;
    border-radius: 3px
}
.branch-action-body:after,
.branch-action-body:before {
    position: absolute;
    top: 11px;
    left: -16px;
    right: 100%;
    width: 0;
    height: 0;
    display: block;
    content: " ";
    border-color: transparent;
    border-style: solid solid outset;
    pointer-events: none
}
.branch-action-body:after {
    border-width: 7px;
    border-right-color: #fff;
    margin-top: 1px;
    margin-left: 2px
}
.branch-action-body:before {
    border-right-color: #ddd;
    border-width: 8px
}
.branch-action-body .spinner {
    display: block;
    float: left;
    width: 32px;
    height: 32px;
    margin-right: 15px;
    background: url(https://github.com/images/spinners/octocat-spinner-32.gif) no-repeat
}
.branch-action-body .merge-message,
.branch-action-body .merge-branch-form {
    padding: 15px;
    background-color: #fafafa;
    border-top: solid 1px #e5e5e5;
    border-bottom-right-radius: 2px;
    border-bottom-left-radius: 2px
}
.post-merge-message {
    padding: 15px
}
.branch-action-item {
    padding: 15px 15px 15px 55px;
    line-height: 1.4
}
.branch-action-item+.branch-action-item {
    border-top: 1px solid #e5e5e5
}
.branch-action-item.open>.build-statuses-list {
    max-height: 215px;
    margin-bottom: -15px
}
.branch-action-item.open .statuses-toggle-opened {
    display: inline
}
.branch-action-item.open .statuses-toggle-closed {
    display: none
}
.branch-action-state-clean .branch-action-icon {
    background-color: #6cc644
}
.branch-action-state-clean .branch-action-body {
    border-color: #95c97e
}
.branch-action-state-clean .branch-action-body:before {
    border-right-color: #95c97e
}
.branch-action-state-unknown .branch-action-icon,
.branch-action-state-unstable .branch-action-icon {
    background-color: #cea61b
}
.branch-action-state-unknown .branch-action-body,
.branch-action-state-unstable .branch-action-body {
    border-color: #e2cc7a
}
.branch-action-state-unknown .branch-action-body:before,
.branch-action-state-unstable .branch-action-body:before {
    border-right-color: #e2cc7a
}
.branch-action-state-merged .branch-action-icon {
    background-color: #6e5494
}
.branch-action-state-merged .branch-action-body {
    border-color: #cbc0db
}
.branch-action-state-merged .branch-action-body:before {
    border-right-color: #cbc0db
}
.branch-action-state-dirty .branch-action-icon,
.branch-action-state-closed-dirty .branch-action-icon {
    background-color: #888
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    .branch-action-body .spinner {
        background-image: url(https://github.com/images/spinners/octocat-spinner-64.gif);
        background-size: 32px 32px
    }
}
.merge-branch-form {
    display: none;
    margin: 15px 0
}
.merge-branch-form .commit-form {
    border-color: #95c97e
}
.merge-branch-form .commit-form:before {
    border-right-color: #95c97e
}
.merge-branch-form.error .commit-form {
    border-color: #e97a74
}
.merge-branch-form.error .commit-form:before {
    border-right-color: #e97a74
}
.merge-branch-form .merge-form-failed {
    display: none
}
.merge-branch-form.error .merge-form-failed {
    display: block
}
.merge-branch-form.error .merge-form-contents {
    display: none
}
.merge-button-matrix-merge-form .merge-branch-form {
    display: block
}
.merge-button-matrix-merge-form .merge-branch-form .merge-form-contents {
    display: block
}
.merge-button-matrix-merge-failed .merge-branch-form {
    display: block
}
.merge-button-matrix-merge-failed .merge-branch-form .merge-form-failed {
    display: block
}
.merge-button-matrix-merge-failed .merge-branch-form .merge-form-contents {
    display: none
}
.input-add-reviewers {
    max-width: 340px
}
.completeness-indicator {
    width: 30px;
    height: 30px;
    text-align: center
}
.completeness-indicator .octicon {
    display: block;
    margin-top: 7px;
    margin-right: auto;
    margin-left: auto
}
.completeness-indicator-success {
    color: #fff;
    background-color: #6cc644;
    border-radius: 50%
}
.completeness-indicator-info {
    color: #888
}
.completeness-indicator-info .mega-octicon {
    font-size: 34px;
    line-height: 30px
}
.completeness-indicator-error {
    color: #fff;
    background-color: #bd2c00;
    border-radius: 50%
}
.completeness-indicator-problem {
    color: #fff;
    background-color: #888;
    border-radius: 50%
}
.completeness-indicator-blank {
    color: #aaa;
    background-color: #f3f3f3;
    border-radius: 50%
}
.completeness-indicator-blank .octicon {
    margin-top: 6px
}
.merge-help-container {
    padding-right: 15px;
    padding-left: 15px;
    margin: 15px -15px -16px -55px
}
.merge-help-container .merge-branch-manually {
    padding-top: 15px;
    padding-bottom: 15px;
    margin-top: 0
}
p.recently-touched-branches-description {
    margin: 0;
    font-size: 11px;
    color: #888
}
.recently-touched-branches {
    padding: 0;
    margin: 5px 0 10px;
    color: #4c4a42;
    background-color: #fff9ea;
    border: solid 1px #dfd8c2;
    border-radius: 3px
}
.recently-touched-branches li {
    height: 36px;
    padding: 5px;
    margin: 0;
    line-height: 23px;
    list-style-type: none;
    border-bottom: 1px solid #e5e2c8
}
.recently-touched-branches li:last-child {
    border-bottom: 0
}
.recently-pushed-branch-actions {
    float: right
}
.recently-pushed-branch-details {
    display: inline-block;
    margin: 0 0 0 7px;
    font-size: 13px;
    line-height: 26px;
    color: #a19e7f
}
.recently-pushed-branch-details a {
    color: #6b694f
}
.recently-pushed-branch-details .css-truncate-target {
    max-width: 400px
}
.branch-name {
    display: inline-block;
    padding: 2px 6px;
    font: 12px Consolas, "Liberation Mono", Menlo, Courier, monospace;
    color: rgba(0, 0, 0, 0.5);
    background-color: rgba(209, 227, 237, 0.5);
    border-radius: 3px
}
.branch-name .octicon {
    margin: 1px -2px 0 0;
    color: #b0c4ce
}
a.branch-name {
    color: #4078c0
}
.range-editor {
    position: relative;
    padding: 5px 15px 5px 40px;
    margin-top: 15px;
    margin-bottom: 15px;
    line-height: 26px;
    background-color: #fafafa;
    border: 1px solid #e5e5e5;
    border-radius: 3px
}
.range-editor .dots {
    font-size: 16px
}
.range-editor .select-menu {
    position: relative;
    display: inline-block
}
.range-editor .select-menu .btn-sm {
    vertical-align: top
}
.range-editor .select-menu.fork-suggester {
    display: none
}
.range-editor .branch-name {
    line-height: 22px
}
.range-editor .branch .css-truncate-target,
.range-editor .fork-suggester .css-truncate-target {
    max-width: 180px
}
.range-editor .pre-mergability {
    display: inline-block;
    padding: 5px;
    line-height: 26px;
    vertical-align: middle
}
.range-editor.is-cross-repo .select-menu.fork-suggester {
    display: inline-block
}
.range-editor-icon {
    float: left;
    margin-top: 10px;
    margin-left: -25px;
    color: #767676
}
.gh-header-new-pr {
    margin-bottom: 15px
}
.gh-header-new-pr .gh-header-meta {
    margin-top: 5px;
    border-bottom: 0;
    padding-bottom: 0
}
.gh-header-new-pr .branch-name {
    display: inline
}
.compare-pr-header {
    display: none
}
.is-pr-composer-expanded .compare-show-header {
    display: none
}
.is-pr-composer-expanded .compare-pr-header {
    display: block
}
.range-cross-repo-pair {
    display: inline-block;
    padding: 5px;
    white-space: nowrap
}
ul.comparison-list {
    width: 350px;
    margin: 25px auto 15px;
    font-size: 14px;
    text-align: left;
    background: #fff;
    border: 1px solid #ddd;
    border-radius: 3px
}
ul.comparison-list>li {
    padding: 7px 10px;
    list-style-type: none;
    border-top: 1px solid #eee
}
ul.comparison-list>li a {
    font-weight: bold
}
ul.comparison-list>li em {
    float: right;
    font-style: normal;
    color: #767676
}
ul.comparison-list>li .octicon {
    position: relative;
    top: 1px;
    color: #aaa
}
ul.comparison-list>li .css-truncate-target {
    max-width: 200px
}
ul.comparison-list>li.title {
    font-size: 12px;
    font-weight: bold;
    color: #aaa;
    text-transform: uppercase;
    background: #fafafa;
    border-top: 0;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.recently-touched-branches-wrapper {
    margin: 15px 0
}
.starring-container .unstarred,
.starring-container.on .starred {
    display: block
}
.starring-container.on .unstarred,
.starring-container .starred {
    display: none
}
.starring-container.loading {
    opacity: 0.5
}
.user-following-container .follow,
.user-following-container.on .unfollow {
    display: inline-block
}
.user-following-container.on .follow,
.user-following-container .unfollow {
    display: none
}
.user-following-container.loading {
    opacity: 0.5
}
.members .user-following-container {
    float: right
}
.close-button {
    background: transparent;
    border: 0;
    padding: 0;
    outline: none
}
.code-list .file-box {
    border: 1px solid #ddd;
    border-radius: 3px
}
.code-list em {
    background-color: rgba(255, 255, 140, 0.5);
    padding: 3px;
    border-radius: 3px;
    font-weight: bold;
    font-style: normal;
    color: #333
}
.code-list .title {
    margin: -3px 0 10px 38px;
    min-height: 24px;
    font-weight: bold;
    line-height: 1.2
}
.code-list .repo-specific .title,
.code-list .repo-specific .full-path {
    margin-left: 0
}
.code-list .match-count,
.code-list .updated-at {
    margin: 0;
    font-size: 11px;
    color: #999;
    font-weight: normal
}
.code-list .language {
    float: right;
    color: rgba(51, 51, 51, 0.75);
    font-size: 12px;
    margin-left: 10px
}
.code-list .avatar {
    float: left
}
.code-list .code-list-item+.code-list-item {
    margin-top: 20px;
    padding-top: 20px;
    border-top: 1px solid #eee;
    margin-bottom: 10px
}
.code-list .blob-num {
    padding: 0
}
.code-list .blob-num:before {
    content: normal
}
.code-list .blob-num a {
    color: inherit;
    padding: 0 10px
}
.code-list .blob-num a:hover {
    color: #4078c0
}
.code-list .blob-code {
    white-space: pre-wrap
}
.code-list .divider .blob-num,
.code-list .divider .blob-code {
    background-color: #f8fafd;
    padding-top: 0;
    padding-bottom: 0;
    cursor: default
}
.code-list .divider .blob-num {
    background-color: #f0f5fa;
    line-height: 15px;
    padding: 0 10px;
    height: 18px
}
.code-list .full-path {
    margin: 0 0 0 40px
}
.code-list .full-path .octicon-repo {
    color: #767676
}
.code-list .full-path .octicon-lock {
    color: #e9dba4
}
.code-list .full-path a {
    color: #999
}
.code-list-item-private .file-box {
    border: 1px solid #fadda5
}
.code-list-item-private .blob-num {
    background-color: #fff9ea;
    border-right: 1px solid #fadda5
}
.code-list-item-private .blob-num a {
    color: #a1882b
}
.code-list-item-private .divider .blob-num,
.code-list-item-private .divider .blob-code {
    background-color: #fff9ea;
    color: #a1882b
}
.codesearch-head {
    padding-bottom: 20px
}
.codesearch-head.pagehead h1 {
    float: left;
    width: 250px;
    line-height: 33px
}
.advanced-search-form h3 {
    margin-top: 20px
}
.advanced-search-form .flattened dt {
    width: 230px
}
.advanced-search-form .flattened dt label {
    font-weight: normal
}
.advanced-search-form .flattened dd {
    margin-left: 250px
}
.advanced-search-form .form-checkbox {
    margin-left: 250px
}
.advanced-search-form fieldset {
    border-bottom: 1px solid #f1f1f1;
    padding-bottom: 20px;
    margin-bottom: 30px
}
.codesearch-results .large-format-loader {
    padding-top: 5%
}
.codesearch-results .repo-list {
    margin-top: -20px
}
.codesearch-results .repo-list-name {
    font-weight: normal
}
.codesearch-results .repo-list-name a,
.codesearch-results .code-list .title a {
    word-wrap: break-word
}
.codesearch-results .repo-list-name em,
.codesearch-results .repo-list-description em {
    padding: 3px;
    font-style: normal;
    font-weight: bold;
    background-color: rgba(255, 255, 140, 0.5);
    border-radius: 3px
}
.meta-search-links {
    margin-top: 20px
}
.meta-search-links a {
    margin-right: 10px
}
.codesearch-aside .menu .octicon {
    width: 16px;
    text-align: center;
    margin-right: 5px
}
.codesearch-aside .meta-search-links {
    margin-top: 20px
}
.codesearch-aside .meta-search-links a {
    margin-right: 10px
}
.codesearch-aside .filter-list {
    border-bottom: 1px solid #f1f1f1;
    margin-bottom: 20px;
    padding-bottom: 20px
}
.codesearch-aside .filter-list li {
    position: relative
}
.codesearch-aside .filter-list li span.bar {
    background: #f1f1f1;
    display: inline-block;
    position: absolute;
    z-index: -1;
    top: 2px;
    bottom: 2px;
    right: 0
}
.simple-search-page {
    width: 740px;
    padding-top: 100px;
    padding-bottom: 100px
}
.simple-search-page h2 {
    font-weight: normal
}
.simple-search-page h2 .mega-octicon {
    vertical-align: middle
}
.search-form-fluid {
    position: relative
}
.search-form-fluid .flex-table-item-primary {
    padding-right: 10px
}
.search-form-fluid .completed-query {
    position: absolute;
    top: 7px;
    left: 8px;
    right: 8px;
    z-index: 1;
    margin: 0;
    overflow: hidden;
    white-space: nowrap
}
.search-form-fluid .completed-query span {
    opacity: 0
}
.search-form-fluid .search-page-label {
    position: relative;
    display: block;
    font-weight: normal;
    cursor: text
}
.search-form-fluid .search-page-label.focus .completed-query {
    opacity: 0.6
}
.search-form-fluid .search-page-input {
    position: relative;
    z-index: 2;
    min-height: 0;
    margin: 0;
    padding: 0;
    background: none;
    border: 0;
    box-shadow: none
}
.search-form-fluid .search-page-input:focus {
    box-shadow: none
}
.token-warning {
    position: absolute;
    top: 10px;
    right: 85px;
    color: #000
}
.sort-bar {
    border-bottom: 1px solid #f1f1f1;
    margin-bottom: 20px;
    padding-bottom: 20px
}
.sort-bar .sort-label {
    padding-right: 5px;
    font-weight: 200;
    font-size: 13px;
    color: #666
}
.sort-bar .select-menu {
    float: right
}
.sort-bar h3 {
    margin: 0
}
.file-editor-textarea {
    width: 100%;
    padding: 5px 4px;
    font: 12px Consolas, "Liberation Mono", Menlo, Courier, monospace;
    resize: vertical;
    border: 0;
    border-radius: 0;
    outline: none
}
.container-preview .tabnav-tabs {
    margin: -6px 0 -6px -11px
}
.container-preview .tabnav-tabs .tabnav-tab {
    padding: 12px 15px;
    border-radius: 0
}
.container-preview .tabnav-tabs>.selected:first-child {
    border-top-left-radius: 3px
}
.container-preview .tabnav-tabs .selected {
    font-weight: bold
}
.container-preview.show-code .commit-create,
.container-preview.show-code .actions {
    display: block
}
.container-preview.show-code .commit-preview,
.container-preview.show-code .loading-preview-msg,
.container-preview.show-code .no-changes-preview-msg,
.container-preview.show-code .error-preview-msg {
    display: none
}
.container-preview:not(.show-code) .commit-create,
.container-preview:not(.show-code) .actions {
    display: none
}
.container-preview.loading-preview .loading-preview-msg {
    display: block
}
.container-preview.loading-preview .no-changes-preview-msg,
.container-preview.loading-preview .error-preview-msg,
.container-preview.loading-preview .commit-preview {
    display: none
}
.container-preview.show-preview .commit-preview {
    display: block
}
.container-preview.show-preview .loading-preview-msg,
.container-preview.show-preview .no-changes-preview-msg,
.container-preview.show-preview .error-preview-msg {
    display: none
}
.container-preview.no-changes-preview .no-changes-preview-msg {
    display: block
}
.container-preview.no-changes-preview .loading-preview-msg,
.container-preview.no-changes-preview .error-preview-msg,
.container-preview.no-changes-preview .commit-preview {
    display: none
}
.container-preview.error-preview .error-preview-msg {
    display: block
}
.container-preview.error-preview .loading-preview-msg,
.container-preview.error-preview .no-changes-preview-msg,
.container-preview.error-preview .commit-preview {
    display: none
}
.container-preview p.preview-msg {
    padding: 30px;
    font-size: 16px
}
.ace_editor.ace-github-light {
    position: relative;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px;
    line-height: 18px
}
.ace_editor.ace-github-light .ace_scroller.ace_scroll-left {
    box-shadow: none
}
.ace_gutter {
    border-right: 1px solid #eee
}
.ace_gutter-layer {
    min-width: 50px
}
.ace_nobold .ace_line>span {
    font-weight: normal !important
}
.ace_marker-layer .ace_step {
    background-color: #fcff00
}
.ace_marker-layer .ace_stack {
    background-color: #a4e565
}
.ace_marker-layer .ace_selected-word {
    background-color: #fafaff
}
.ace_indent-guide {
    box-shadow: inset -1px 0 0 rgba(0, 0, 0, 0.1)
}
.details-collapse .collapse {
    position: relative;
    display: none;
    height: 0;
    overflow: hidden;
    -webkit-transition: height 0.35s ease-in-out;
    transition: height 0.35s ease-in-out
}
.details-collapse.open .collapse {
    display: block;
    height: auto;
    overflow: visible
}
.comment .email-format {
    line-height: 1.5
}
.comment .context-loader {
    display: none
}
.previewable-edit .previewable-comment-form {
    display: none
}
.previewable-edit .previewable-comment-form .form-actions {
    margin-bottom: 10px;
    margin-right: 10px
}
.previewable-edit.is-comment-editing .timeline-comment-header {
    display: none
}
.is-comment-editing .timeline-comment-actions,
.is-comment-editing .edit-comment-hide {
    display: none
}
.is-comment-editing .previewable-comment-form {
    display: block
}
.is-comment-loading .context-loader {
    display: block
}
.is-comment-loading .previewable-comment-form {
    opacity: 0.5
}
.is-comment-stale .comment-form-stale {
    display: block
}
.is-comment-stale .comment-content {
    padding-top: 10px
}
.is-comment-stale .previewable-comment-form {
    opacity: 0.75
}
.comment-body {
    width: 100%;
    padding: 15px;
    overflow: visible;
    font-size: 14px
}
.comment-body:before {
    display: table;
    content: ""
}
.comment-body:after {
    display: table;
    clear: both;
    content: ""
}
.comment-body .highlight {
    overflow: visible !important;
    background-color: transparent
}
.comment-form-textarea {
    width: 100%;
    max-width: 100%;
    height: 100px;
    min-height: 100px;
    margin: 0;
    font-size: 14px;
    line-height: 1.6;
    -webkit-transform: translateZ(0);
    transform: translateZ(0)
}
.comment-form-textarea.dragover {
    border: solid 1px #4078c0
}
.discussion-topic-header {
    position: relative;
    padding: 10px;
    word-wrap: break-word
}
.comment-form-error,
.comment-form-stale {
    display: none;
    padding: 5px 10px;
    margin: 0 10px;
    font-weight: bold;
    color: #900;
    background-color: #ffeaea;
    border: 1px solid #e2a0a0
}
.comment-form-error.comment-form-bottom,
.comment-form-stale.comment-form-bottom {
    margin-bottom: 10px
}
.comment-form-stale {
    margin-top: 0
}
.email-format {
    line-height: 1.5em !important
}
.email-format div {
    white-space: pre-wrap
}
.email-format .email-hidden-reply {
    display: none;
    white-space: pre-wrap
}
.email-format .email-quoted-reply,
.email-format .email-signature-reply {
    padding: 0 15px;
    margin: 15px 0;
    color: #767676;
    border-left: 4px solid #ddd
}
.email-format .email-hidden-toggle a {
    display: inline-block;
    height: 12px;
    padding: 0 9px;
    font-size: 12px;
    font-weight: bold;
    line-height: 6px;
    color: #555;
    text-decoration: none;
    vertical-align: middle;
    background: #ddd;
    border-radius: 1px
}
.email-format .email-hidden-toggle a:hover {
    background-color: #ccc
}
.email-format .email-hidden-toggle a:active {
    color: #fff;
    background-color: #4078c0
}
.comment-email-format div {
    white-space: normal
}
.comment-email-format .email-hidden-reply {
    display: none;
    white-space: normal
}
.comment-email-format blockquote,
.comment-email-format p {
    margin: 0
}
.blankslate.conversation-limited {
    padding: 20px 0 10px;
    margin: 15px
}
.locked-conversation .write-tab,
.locked-conversation .preview-tab {
    color: #ccc
}
.commit-sha {
    padding: 0.2em 0.4em;
    font-size: 90%;
    font-weight: normal;
    background-color: #f5f5f5;
    border: 1px solid #eee;
    border-radius: 0.2em
}
.commit-partial-notice {
    margin-top: 20px;
    margin-bottom: 20px
}
.commit-paginate-container {
    float: right;
    margin: -5px 0 0;
    text-align: inherit
}
.commit .commit-title,
.commit .commit-title a {
    color: #4e575b
}
.commit .commit-title.blank,
.commit .commit-title.blank a {
    color: #9cabb1
}
.commit .commit-title .issue-link {
    font-weight: bold;
    color: #4078c0
}
.commit .sha-block,
.commit .sha {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px
}
.commit.open .commit-desc {
    display: block
}
.commit-link {
    font-weight: normal;
    color: #4078c0
}
.commit-email-flash {
    display: inline
}
.commit-desc {
    display: none
}
.commit-desc pre {
    max-width: 700px;
    margin: 10px 0;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 13px;
    line-height: 1.45;
    color: #596063;
    white-space: pre-wrap
}
.commit-desc+.commit-branches {
    padding-top: 8px;
    margin-top: 2px;
    border-top: solid 1px #d1e2eb
}
.commit-tease {
    padding: 8px 8px 0;
    background: #e6f1f6;
    border: 1px solid #b7c7cf;
    border-bottom-color: #d8e6ec;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.commit-tease p.commit-title {
    margin: 0 0 6px
}
.commit-tease .commit-desc {
    margin: -3px 0 10px
}
.commit-tease .commit-desc pre {
    font-size: 11px
}
.commit-tease .commit-meta {
    padding: 8px;
    margin-right: -8px;
    margin-left: -8px;
    background: #fff;
    border-top: 1px solid #d8e6ec
}
.commit-tease .commit-meta .loader-loading {
    margin: 0 0 -9px
}
.commit-tease .zeroclipboard-link {
    float: right;
    margin-top: -2px;
    margin-left: 5px
}
.commit-tease .sha-block {
    float: right;
    color: #888
}
.commit-tease .sha-block>.sha {
    color: #444
}
.commit-tease .sha-block>a {
    color: #444;
    text-decoration: none
}
.commit-tease .authorship {
    margin: -2px 0 -4px -4px;
    font-size: 12px;
    color: #767676
}
.commit-tease .authorship a {
    font-weight: bold;
    color: #444;
    text-decoration: none
}
.commit-tease .authorship a:hover {
    text-decoration: underline
}
.commit-tease .authorship .avatar {
    margin: -2px 3px 0 0
}
.commit-tease .authorship .author-name {
    color: #444
}
.commit-tease .authorship .committer {
    display: block;
    margin-left: 30px;
    font-size: 11px
}
.comment-count {
    float: right;
    margin-top: 1px;
    font-size: 11px;
    color: #7f9199
}
.comment-count .octicon {
    margin-left: 5px;
    vertical-align: middle
}
.commits-listing {
    position: relative;
    padding-bottom: 20px;
    margin-bottom: 15px
}
.commits-listing:before {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 14px;
    z-index: -1;
    display: block;
    width: 2px;
    content: "";
    background-color: #f3f3f3
}
.commits-listing .discussion-item-icon {
    margin-right: 5px;
    margin-left: -1px
}
.commits-listing .timeline-commits {
    padding-left: 8px;
    margin-bottom: 20px
}
.commits-listing .timeline-commits:last-child {
    margin-bottom: 0
}
.commits-listing-padded {
    padding-left: 39px
}
.commit-group {
    margin-top: 10px;
    list-style-type: none
}
.commit-group-title {
    margin-top: 15px;
    margin-left: -31px;
    color: #767676
}
.commit-group-title .octicon-git-commit {
    margin-right: 17px;
    color: #ccc;
    background: #fff
}
.commits-list-item.navigation-focus {
    background: #f7fbfc
}
.commits-list-item .commit-title {
    margin: 0;
    font-size: 15px;
    font-weight: bold;
    color: #333
}
.commits-list-item .commit-meta {
    margin-top: 1px;
    font-weight: normal;
    color: #767676
}
.commits-list-item .status .octicon {
    height: 14px;
    line-height: 14px
}
.commits-list-item .commit-author {
    color: #767676
}
.commits-list-item .octicon-arrow-right {
    margin: 0 3px
}
.commits-list-item .btn-outline {
    margin-top: 2px
}
.commits-list-item .commit-desc pre {
    padding-left: 8px;
    margin-top: 5px;
    margin-bottom: 10px;
    font-size: 12px;
    color: #596063;
    border-left: 1px solid #e5e5e5
}
.commits-list-item .commit-desc pre a {
    word-break: break-word
}
.commits-comments-link {
    margin-top: 9px;
    color: #767676;
    vertical-align: middle
}
.commits-comments-link:hover {
    color: #4183c4;
    text-decoration: none
}
.commit-avatar-cell {
    width: 47px
}
.commit-avatar-cell.table-list-cell {
    padding-right: 0
}
.commit-indicator {
    margin-left: 5px
}
.commit-links-cell {
    width: 230px;
    text-align: right
}
.commit-links-group {
    margin-right: 5px
}
.timeline-commits {
    width: 100%;
    margin-top: 5px;
    border-collapse: separate
}
.timeline-commits+.timeline-commits {
    margin-top: 15px
}
.timeline-commits td {
    padding-top: 4px;
    padding-right: 8px;
    padding-bottom: 4px;
    font-size: 12px;
    line-height: 16px;
    vertical-align: top;
    background-color: transparent
}
.discussion-item .timeline-commits .commit-author {
    display: none
}
.timeline-commits .commit-gravatar {
    width: 16px;
    padding-left: 10px
}
.timeline-commits .commit-author {
    width: 200px;
    padding-right: 20px;
    white-space: nowrap
}
.timeline-commits .author {
    font-weight: bold;
    color: #555
}
.timeline-commits .commit-message {
    min-height: 0;
    max-width: 550px
}
.timeline-commits .commit-message a {
    color: #555
}
.timeline-commits .commit-message a:hover {
    color: #4078c0
}
.timeline-commits .commit-desc pre {
    padding-left: 10px;
    font-size: 11px;
    color: #767676;
    border-left: 1px solid #eee;
    overflow: visible
}
.timeline-commits .hidden-text-expander {
    margin-top: 3px;
    margin-left: 0;
    vertical-align: top
}
.timeline-commits .hidden-text-expander a {
    height: 13px;
    background-color: #eee
}
.timeline-commits .hidden-text-expander a:hover {
    color: #fff;
    background-color: #4078c0
}
.timeline-commits .commit-meta {
    text-align: right;
    white-space: nowrap
}
.timeline-commits .commit-meta .status {
    width: 16px;
    text-align: center
}
.timeline-commits .commit-meta .status.status-pending {
    color: #cea61b
}
.timeline-commits .commit-meta .octicon {
    margin-right: 1px;
    margin-left: 1px
}
.commit-icon {
    display: table-cell;
    width: 16px;
    color: #ccc
}
.commit-icon .octicon {
    background-color: #fff
}
.commit-id {
    color: #bbb
}
.commit-id:hover {
    color: #4078c0
}
.full-commit {
    padding: 8px 8px 0;
    margin: 10px 0;
    background: #e6f1f6;
    border: 1px solid #c5d5dd;
    border-radius: 3px
}
.full-commit:first-child {
    margin-top: 0
}
.full-commit .btn-outline {
    background: none;
    border: 1px solid #cedee5
}
.full-commit .btn-outline:hover {
    color: #4078c0;
    border: 1px solid #4078c0
}
.full-commit p.commit-title {
    margin: 0 0 8px;
    font-size: 18px;
    font-weight: bold;
    color: #213f4d;
    text-shadow: 0 1px rgba(255, 255, 255, 0.5)
}
.full-commit .branches-list {
    display: inline;
    margin-right: 10px;
    margin-left: 2px;
    vertical-align: middle;
    list-style: none
}
.full-commit .branches-list li {
    display: inline-block;
    padding-left: 3px;
    font-weight: bold;
    color: #596063
}
.full-commit .branches-list li:before {
    padding-right: 6px;
    font-weight: normal;
    content: "+"
}
.full-commit .branches-list li:first-child {
    padding-left: 0
}
.full-commit .branches-list li:first-child:before {
    padding-right: 0;
    content: ""
}
.full-commit .branches-list li.loading {
    font-weight: normal;
    color: #818c90
}
.full-commit .branches-list li.pull-request {
    font-weight: normal;
    color: #818c90
}
.full-commit .branches-list li.pull-request:before {
    margin-left: -8px;
    content: ""
}
.full-commit .branches-list li.pull-request-error {
    margin-bottom: -1px
}
.full-commit .branches-list li a {
    color: inherit
}
.full-commit .commit-meta {
    padding: 8px;
    margin-right: -8px;
    margin-left: -8px;
    background: #fff;
    border-top: 1px solid #d8e6ec;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.full-commit .sha-block {
    float: right;
    margin-left: 15px;
    font-size: 12px;
    color: #888
}
.full-commit.merge-commit .sha-block {
    clear: right
}
.full-commit.merge-commit .sha-block+.sha-block {
    margin-top: 2px
}
.full-commit .sha-block>.sha {
    color: #444
}
.full-commit .sha-block>a {
    color: #444;
    text-decoration: none;
    border-bottom: 1px dotted #ccc
}
.full-commit .sha-block>a:hover {
    border-bottom: 1px solid #444
}
.full-commit .authorship {
    margin-top: -2px;
    margin-bottom: -4px;
    margin-left: -4px;
    font-size: 14px;
    color: #767676
}
.full-commit .authorship .gravatar {
    margin-top: -2px;
    margin-right: 3px;
    vertical-align: middle;
    border-radius: 3px
}
.full-commit .authorship a {
    font-weight: bold;
    color: #444;
    text-decoration: none
}
.full-commit .authorship a:hover {
    text-decoration: underline
}
.full-commit .authorship .author-name {
    color: #444
}
.full-commit .authorship .hint a {
    color: #4078c0
}
.full-commit .authorship .committer {
    display: block;
    margin-top: -2px;
    margin-left: 34px;
    font-size: 12px
}
.full-commit .commit-desc {
    display: block;
    margin: -5px 0 10px
}
.full-commit .commit-desc pre {
    max-width: 100%;
    overflow: visible;
    text-shadow: 0 1px rgba(255, 255, 255, 0.5);
    word-wrap: break-word
}
.branches-tag-list {
    display: inline;
    margin-right: 10px;
    margin-left: 2px;
    vertical-align: middle;
    list-style: none
}
.branches-tag-list .more-commit-details,
.branches-tag-list.open .hidden-text-expander {
    display: none
}
.branches-tag-list.open .more-commit-details {
    display: inline-block
}
.branches-tag-list li {
    display: inline-block;
    padding-left: 3px
}
.branches-tag-list li:first-child {
    padding-left: 0;
    font-weight: bold;
    color: #596063
}
.branches-tag-list li.loading {
    font-weight: normal;
    color: #818c90
}
.branches-tag-list li.abbrev-tags {
    cursor: pointer
}
.branches-tag-list li a {
    color: inherit
}
.branches-tag-list li .hidden-text-expander a {
    background-color: #dae5eb
}
.branches-tag-list li .hidden-text-expander a:hover {
    background-color: #d1dbe0
}
.commit-branches {
    min-height: 18px;
    margin-top: -6px;
    margin-bottom: 8px;
    font-size: 12px;
    color: #818c90;
    vertical-align: middle
}
.commit-branches .octicon {
    vertical-align: middle
}
.commit-loader .loader-error {
    display: none;
    margin: 0;
    font-size: 12px;
    font-weight: bold;
    color: #bd2c00
}
.commit-loader.error .loader-loading {
    display: none
}
.commit-loader.error .loader-error {
    display: block
}
.historical-banner {
    padding: 15px 20px 15px 130px;
    margin-bottom: 20px;
    overflow: hidden;
    color: #333;
    background: #fff;
    border: 1px solid #e5e5e5;
    border-radius: 5px
}
.historical-banner h2 {
    margin: 0 0 5px
}
.historical-banner p {
    margin: 0
}
.historical-banner .illustration {
    position: absolute;
    top: 12px;
    left: 20px;
    color: rgba(0, 0, 0, 0.1)
}
.roses-divider {
    margin-bottom: 20px;
    text-align: center
}
.file-history-tease {
    margin-bottom: 10px;
    font-size: 14px;
    color: #7b878c;
    background: #e7ecee;
    border: 1px solid #d2d9dd;
    border-radius: 3px
}
.file-history-tease .author a {
    font-weight: bold;
    color: #000
}
.file-history-tease .commit-title {
    display: inline
}
.file-history-tease .sha {
    font-size: 13px
}
.file-history-tease .loader-loading,
.file-history-tease .loader-error {
    margin: 0
}
.file-history-tease .loader-loading img,
.file-history-tease .loader-error img {
    vertical-align: middle
}
.file-history-tease .participation {
    padding: 5px 8px;
    font-size: 12px;
    font-weight: normal;
    line-height: 20px;
    color: #666;
    background-color: #fff;
    border-top: 1px solid #d8e6ec;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.file-history-tease .participation:before {
    display: table;
    content: ""
}
.file-history-tease .participation:after {
    display: table;
    clear: both;
    content: ""
}
.file-history-tease .participation .quickstat {
    float: left;
    margin: 0 10px 0 0
}
.file-history-tease .participation .quickstat strong {
    color: #000
}
.file-history-tease .participation a {
    color: #888;
    text-decoration: none
}
.file-history-tease .participation .avatar-link {
    margin-right: 3px
}
.file-history-tease-header {
    padding: 5px 8px;
    line-height: 24px
}
.file-history-tease-header .avatar {
    float: left;
    margin-right: 5px
}
.commit-comments-heading h3 {
    display: inline-block;
    margin-right: 15px
}
.commit-build-statuses {
    position: relative;
    display: inline-block;
    text-align: left
}
.commit-build-statuses.active .dropdown-menu-content {
    display: block
}
.commit-build-statuses.active .tooltipped:before,
.commit-build-statuses.active .tooltipped:after {
    display: none
}
.commit-build-statuses .dropdown-menu {
    min-width: 400px;
    max-width: 500px;
    padding-top: 0;
    padding-bottom: 0
}
.commit-build-statuses .dropdown-menu .build-statuses-list {
    max-height: 170px;
    border-bottom: 0
}
.commit-build-statuses .dropdown-menu-w,
.commit-build-statuses .dropdown-menu-e {
    top: -11px
}
.commit-build-statuses .build-status-item:last-child {
    border-radius: 0 0 2px 2px
}
.commit-ref {
    position: relative;
    display: inline-block;
    padding: 0 5px;
    border-radius: 3px;
    font: 0.75em/2 Consolas, "Liberation Mono", Menlo, Courier, monospace;
    color: #336479;
    white-space: nowrap;
    background-color: #e8f0f8
}
.commit-ref .user {
    color: #598a9f
}
a.commit-ref:hover {
    text-shadow: -1px -1px 0 rgba(0, 0, 0, 0.2);
    background-image: -webkit-linear-gradient(#74a4d4, #2a5177);
    background-image: linear-gradient(#74a4d4, #2a5177);
    border-color: #2a5177;
    text-decoration: none
}
.compare-cutoff,
.diff-cutoff {
    margin: 5px 0;
    padding: 8px 0;
    font-weight: bold;
    text-align: center;
    border-radius: 3px;
    color: #4c4a42;
    background-color: #fff9ea;
    border: solid 1px #dfd8c2
}
.table-of-contents {
    margin: 15px 0
}
.table-of-contents li {
    padding: 7px 0;
    list-style-type: none
}
.table-of-contents li+li {
    border-top: 1px solid #eee
}
.table-of-contents li>.octicon {
    margin-right: 3px;
    vertical-align: -1px
}
.table-of-contents .octicon-diff-removed {
    color: #bd2c00
}
.table-of-contents .octicon-diff-renamed {
    color: #677a85
}
.table-of-contents .octicon-diff-modified {
    color: #d0b44c
}
.table-of-contents .octicon-diff-added {
    color: #6cc644
}
.toc-diff-stats {
    padding-left: 20px;
    line-height: 26px
}
.toc-diff-stats .octicon {
    float: left;
    margin-top: 3px;
    margin-left: -20px;
    color: #ccc
}
.toc-diff-stats .btn-link {
    font-weight: bold
}
.toc-diff-stats+.content {
    padding-top: 5px
}
span.no-nl-marker {
    position: relative;
    color: #bd2c00;
    vertical-align: middle
}
.symlink .no-nl-marker {
    display: none
}
.existing-pull {
    margin: 10px 0
}
.existing-pull .list-group-item:before {
    display: table;
    content: ""
}
.existing-pull .list-group-item:after {
    display: table;
    clear: both;
    content: ""
}
.existing-pull .existing-pull-contents {
    float: left;
    width: 680px
}
.existing-pull .existing-pull-button {
    float: right;
    margin-top: 3px
}
.existing-pull .existing-pull-number {
    font-weight: normal;
    color: #aaa
}
.existing-pull .css-truncate {
    max-width: 700px
}
.existing-pull .css-truncate p {
    display: inline
}
.compare-pr-placeholder {
    margin: 10px 0;
    padding: 15px;
    font-size: 14px;
    color: #4c4a42;
    background-color: #fff9ea;
    border: solid 1px #dfd8c2;
    border-radius: 3px
}
.compare-pr-placeholder p {
    margin: 7px 0;
    color: #6d6c60
}
.compare-pr-placeholder .btn {
    margin-right: 10px;
    margin-bottom: -2px
}
.compare-pr-placeholder .btn .octicon {
    vertical-align: -1px
}
.compare-pr-placeholder .help-link {
    margin-top: 5px;
    margin-right: -3px;
    padding: 3px;
    color: #9c997d;
    text-decoration: none
}
.compare-pr .new-pr-form {
    display: none
}
.compare-pr .contributing {
    display: none
}
.compare-pr.open .compare-pr-placeholder {
    display: none
}
.compare-pr.open .new-pr-form {
    display: block
}
.compare-pr.open .contributing {
    display: block
}
.contributing {
    padding: 15px;
    margin-bottom: 15px;
    font-size: 14px;
    color: #4c4a42;
    background-color: #fff9ea;
    border: 1px solid #dfd8c2;
    border-radius: 3px
}
#contact-github textarea {
    height: 100px;
    resize: vertical
}
#contact-github .contact-checklist {
    margin: 20px 0
}
#contact-github .contact-checklist>li {
    margin: 15px 0 15px 18px;
    list-style-position: outside
}
.heartocat {
    display: block;
    margin: 50px auto 0
}
.documentation-results-wrapper {
    position: relative;
    top: -19px
}
.documentation-results {
    position: absolute;
    top: 0;
    z-index: 2;
    width: 400px;
    margin-top: 5px;
    clear: both;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.2)
}
.documentation-results ul {
    width: 100%
}
.documentation-results ul li:first-child a {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.documentation-results .documentation-results-footer a {
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px
}
.documentation-results a {
    outline: none;
    padding: 5px 10px;
    display: block;
    color: #333;
    font-weight: bold;
    cursor: pointer;
    text-decoration: none;
    border: solid #ddd;
    border-width: 0 1px 1px;
    background-color: #fff
}
.documentation-results a:hover {
    background-color: #3586c3;
    color: #fff
}
.documentation-results a.selected {
    background-color: #3586c3;
    color: #fff
}
ul.documentation-results-group {
    list-style-type: none
}
.contact-form-extras {
    display: none
}
.context-loader-container .context-loader {
    display: none
}
.context-loader-container .context-loader.is-context-loading {
    display: block;
    white-space: nowrap
}
.context-loader-container .context-loader-overlay {
    opacity: 1;
    -webkit-transition: opacity 0.25s ease-in-out;
    transition: opacity 0.25s ease-in-out
}
.context-loader-container .context-loader-overlay.is-context-loading {
    opacity: 0.5
}
.page-context-loader {
    margin-left: 10px;
    display: none
}
.page-context-loader.is-context-loading {
    display: inline-block
}
body.disables-context-loader .page-context-loader,
body.disables-context-loader .context-loader {
    display: none !important
}
.contributions-tab {
    margin-top: 20px
}
.contributions-tab .simple-conversation-list .state {
    margin-top: 3px
}
.calendar-graph {
    padding: 5px 0 0;
    height: 126px;
    text-align: center
}
.calendar-graph.days-selected rect.day {
    opacity: 0.5
}
.calendar-graph.days-selected rect.day.active {
    opacity: 1
}
.calendar-graph .activity {
    margin-top: 0
}
.calendar-graph .dots {
    margin: 20px auto 0;
    width: 64px;
    height: 64px
}
.calendar-graph text.month {
    font-size: 10px;
    fill: #aaa
}
.calendar-graph text.wday {
    fill: #ccc;
    font-size: 9px
}
#contributions-calendar rect.day {
    shape-rendering: crispedges
}
#contributions-calendar rect.day.empty:hover {
    stroke: none
}
#contributions-calendar rect.day:hover {
    stroke: #555;
    stroke-width: 1px
}
.contrib-column {
    padding: 15px 0;
    text-align: center;
    border-left: 1px solid #ddd;
    border-top: 1px solid #ddd;
    font-size: 11px
}
.contrib-column-first {
    border-left: 0
}
.contrib-number {
    font-weight: 300;
    line-height: 1.3em;
    font-size: 24px;
    display: block;
    color: #333
}
.contrib-footer {
    font-size: 11px;
    padding: 0 10px 12px
}
.contrib-legend {
    font-size: 11px;
    color: #767676;
    float: right
}
.contrib-legend .legend {
    display: inline-block;
    list-style: none;
    margin: 0 5px;
    position: relative;
    bottom: -1px
}
.contrib-legend .legend li {
    display: inline-block;
    width: 10px;
    height: 10px
}
.new-user-contrib-intro {
    border-top: solid 1px #ddd;
    padding: 5px 20px;
    font-size: 16px
}
.contrib-square {
    color: #d6e685;
    font-size: 22px;
    line-height: 1
}
.contribution-activity h2 {
    font-size: 18px;
    font-weight: normal;
    margin: 30px 0 15px
}
.contribution-activity .select-menu-button {
    position: relative;
    top: -4px
}
.contribution-activity.loading .contribution-activity-listing {
    display: none
}
.contribution-activity.loading .contribution-activity-spinner {
    display: block
}
.contribution-activity-spinner {
    margin: 20px auto 0;
    width: 64px;
    height: 64px;
    display: none
}
ul.simple-conversation-list a.meta {
    color: #767676
}
li.contribution {
    list-style: none;
    padding: 10px 0
}
li.contribution:last-child {
    border-bottom: 0
}
li.contribution h3 {
    font-size: 14px;
    display: inline-block;
    margin: 0
}
li.contribution .cmeta {
    display: block;
    font-size: 12px
}
li.contribution .cmt {
    color: #767676
}
li.contribution .d {
    color: #c00
}
li.contribution .a {
    color: #8cac29
}
li.contribution .num {
    color: #767676
}
.subscribe-feed {
    display: inline-block;
    color: #333
}
.subscribe-feed .octicon {
    margin-right: 5px
}
.new-user-panel {
    position: relative;
    padding: 18px;
    margin-bottom: 30px;
    font-size: 16px;
    border: dashed 2px #ccc;
    border-radius: 3px
}
.new-user-panel-close {
    position: absolute;
    top: 10px;
    right: 18px;
    color: #ccc
}
.new-user-panel-close:hover {
    color: #666
}
.new-user-intro {
    margin: 0 120px;
    font-size: 36px;
    font-weight: normal;
    line-height: 1.3;
    text-align: center
}
.new-user-heading-small {
    margin: 5px 170px 20px;
    font-size: 20px;
    color: #888;
    text-align: center
}
.button-hello-world {
    font-size: 16px;
    padding: 10px 50px
}
.welcome-guides {
    position: relative;
    padding: 30px;
    margin-bottom: 30px;
    border: 2px dashed #ddd;
    border-radius: 3px
}
.welcome-guides h1 {
    margin-top: 0;
    margin-bottom: 0;
    text-align: center
}
.welcome-guides .lead {
    margin-top: 0;
    margin-bottom: 20px;
    text-align: center
}
.welcome-guides .guides-list-item {
    color: #767676
}
.welcome-guides .guides-list-item:hover {
    color: #4078c0;
    text-decoration: none
}
.welcome-guides .guides-list-item p {
    margin-bottom: 0
}
.welcome-guides .dismiss-guides {
    position: absolute;
    top: 30px;
    right: 28px;
    display: block;
    padding-left: 5px;
    color: #ccc
}
.welcome-guides .dismiss-guides:hover {
    color: #4078c0
}
.guides-bootcamp {
    text-align: center
}
.guides-bootcamp p {
    margin-top: 0;
    margin-bottom: 0
}
.guides-bootcamp .guides-list-item {
    padding-right: 15px;
    padding-left: 15px
}
.guides-bootcamp .guides-image {
    display: block;
    width: auto;
    height: 100px;
    margin: 10px auto 20px
}
.guides-bootcamp .guides-list-title {
    margin-top: 0;
    margin-bottom: 5px;
    color: #4078c0
}
.bootcamp {
    margin: 0 0 20px
}
.bootcamp h1 {
    position: relative;
    padding: 8px 10px;
    margin: 0;
    font-size: 16px;
    font-weight: bold;
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.4);
    background-color: #829aa8;
    border: 1px solid #677c89;
    border-bottom-color: #6b808d;
    border-radius: 3px 3px 0 0
}
.bootcamp h1 a {
    color: #fff;
    text-decoration: none
}
.bootcamp .octicon-x {
    position: relative;
    top: -2px;
    font-size: 16px;
    line-height: 16px
}
.bootcamp .dismiss-bootcamp {
    position: absolute;
    top: 9px;
    right: 9px;
    display: block;
    width: 16px;
    height: 16px;
    background-repeat: no-repeat;
    background-position: 0 0
}
.bootcamp .dismiss-bootcamp:hover {
    background-position: 0 -19px
}
.bootcamp .bootcamp-body {
    padding: 10px 1%;
    overflow: hidden;
    background-color: #e9f1f4;
    border-color: #e9f1f4 #d8dee2 #d8dee2;
    border-style: solid;
    border-width: 1px;
    border-radius: 0 0 3px 3px
}
.bootcamp ul li {
    background-image: -webkit-linear-gradient(#fffff5, #f5f3b4);
    background-image: linear-gradient(#fffff5, #f5f3b4);
    position: relative;
    display: block;
    float: left;
    width: 24.25%;
    height: 215px;
    margin: 0 1% 0 0;
    overflow: hidden;
    font-size: 13px;
    font-weight: normal;
    color: #666;
    border: 1px solid #dfddb5;
    border-radius: 3px;
    box-shadow: 0 1px 0 #fff
}
.bootcamp ul li.be-social {
    margin-right: 0
}
.bootcamp ul li:hover {
    background-image: -webkit-linear-gradient(#fcfce9, #f1eea3);
    background-image: linear-gradient(#fcfce9, #f1eea3);
    border: 1px solid #d6d4ad
}
.bootcamp ul li a {
    color: #666;
    text-decoration: none
}
.bootcamp .image {
    position: relative;
    display: block;
    height: 133px;
    background-repeat: no-repeat;
    background-position: center center
}
.bootcamp .setup .image {
    background-image: url(https://github.com/images/modules/dashboard/bootcamp/octocat_setup.png);
    background-size: 129px 96px
}
.bootcamp .create-a-repo .image {
    background-image: url(https://github.com/images/modules/dashboard/bootcamp/octocat_repo.png);
    background-size: 129px 96px
}
.bootcamp .fork-a-repo .image {
    background-image: url(https://github.com/images/modules/dashboard/bootcamp/octocat_fork.png);
    background-size: 178px 96px
}
.bootcamp .be-social .image {
    background-image: url(https://github.com/images/modules/dashboard/bootcamp/octocat_collabocats.png);
    background-position: center 27px;
    background-size: 207px 96px
}
.bootcamp .desc {
    position: relative;
    z-index: 2;
    padding: 10px 15px 20px;
    overflow: hidden;
    text-align: center;
    background-repeat: no-repeat
}
.bootcamp .desc>h2 {
    padding: 0;
    margin: 0 0 5px;
    font-size: 15px;
    color: #393939
}
.bootcamp .desc p {
    padding: 0;
    margin: 0;
    line-height: 1.2em
}
.bootcamp .step-number {
    position: absolute;
    top: -1px;
    left: 10px;
    font-size: 36px;
    font-weight: bold;
    color: #e4e1a8;
    opacity: 0.75
}
.user-repos .mini-repo-list-item {
    padding-right: 6px
}
.user-repos .mini-repo-list-item .repo-and-owner {
    max-width: 100%
}
.user-repos .mini-repo-list-item .owner {
    max-width: 145px
}
#example_octofication {
    float: right;
    width: 335px;
    margin: 0
}
.octofication {
    margin-bottom: 15px
}
.octofication .message {
    min-height: 56px;
    padding: 10px 10px 10px 50px;
    border: solid 1px #4078c0;
    border-radius: 3px
}
.octofication .message h3 {
    margin: 1px 20px 3px 0;
    font-size: 14px;
    line-height: 1.2
}
.octofication .message p {
    padding: 0;
    margin: 0;
    font-size: 12px;
    color: #555
}
.octofication .message p+p {
    margin-top: 15px
}
.octofication .broadcast-icon {
    position: relative;
    float: left;
    margin-left: -40px;
    color: #4078c0
}
.octofication .broadcast-icon-mask {
    position: absolute;
    top: 0;
    width: 10px;
    height: 16px;
    background-color: #fff;
    opacity: 0;
    -webkit-animation: broadCastMaskFade 1s ease-in-out 2s 2;
    animation: broadCastMaskFade 1s ease-in-out 2s 2
}
.octofication .broadcast-icon-mask.left {
    left: 0
}
.octofication .broadcast-icon-mask.right {
    right: 0
}
.octofication .notice-dismiss {
    position: relative;
    top: -2px;
    float: right;
    color: #bbb
}
.octofication .notice-dismiss:hover {
    color: #666
}
.octofication-more {
    margin: 5px 0;
    font-size: 11px;
    text-align: right
}
@-webkit-keyframes broadCastMaskFade {
    0% {
        opacity: 0
    }
    30% {
        opacity: 1
    }
    70% {
        opacity: 1
    }
    100% {
        opacity: 0
    }
}
@keyframes broadCastMaskFade {
    0% {
        opacity: 0
    }
    30% {
        opacity: 1
    }
    70% {
        opacity: 1
    }
    100% {
        opacity: 0
    }
}
.github-jobs-promotion {
    margin-bottom: 15px
}
.github-jobs-promotion p {
    background-image: -webkit-linear-gradient(#f5fbff, #e4f0ff);
    background-image: linear-gradient(#f5fbff, #e4f0ff);
    position: relative;
    padding: 10px 18px;
    font-size: 12px;
    color: #1b3650;
    text-align: center;
    border: 1px solid #cee0e7;
    border-radius: 3px
}
.github-jobs-promotion p a {
    color: #1b3650
}
.github-jobs-promotion a.jobs-logo {
    display: block;
    font-size: 11px;
    color: #767676;
    text-align: center
}
.github-jobs-promotion a.jobs-logo:hover {
    text-decoration: none
}
.github-jobs-promotion a.jobs-logo strong {
    display: inline-block;
    width: 62px;
    height: 12px;
    text-indent: -9999px;
    vertical-align: middle;
    background: url(https://github.com/images/modules/jobs/logo.png) 0 0 no-repeat;
    background-size: 62px auto
}
.github-jobs-promotion .job-location {
    white-space: nowrap
}
.github-jobs-promotion a.octicon-info {
    position: absolute;
    right: 5px;
    bottom: 5px;
    color: #a9b8be;
    text-decoration: none;
    cursor: pointer;
    opacity: 0.8
}
.github-jobs-promotion p:hover .octicon-info {
    opacity: 1
}
#dashboard h1 {
    margin-bottom: 0.5em;
    font-size: 160%
}
#dashboard h1 a {
    font-size: 70%;
    font-weight: normal
}
#dashboard .notice {
    padding: 15px;
    margin-top: 0;
    margin-bottom: 0;
    text-align: center
}
.news .account-switcher {
    margin-bottom: 20px
}
.news .release {
    margin-top: 0;
    margin-bottom: 0
}
.news blockquote {
    color: #666
}
.news h1 {
    margin-bottom: 0
}
.news .alert {
    position: relative;
    padding: 0 0 1em 45px;
    overflow: hidden;
    border-top: 1px solid #f1f1f1
}
.news .alert .commits {
    padding-left: 40px
}
.news .alert .css-truncate.css-truncate-target,
.news .alert .css-truncate .css-truncate-target {
    max-width: 180px
}
.news .alert p {
    margin: 0
}
.news .alert .markdown-body blockquote {
    padding: 0 0 0 40px;
    border: 0 none
}
.news .alert .mega-octicon {
    position: absolute;
    top: 14px;
    left: 0;
    width: 32px;
    height: 32px;
    padding: 3px;
    color: #bbb
}
.news .alert .mega-octicon::before {
    color: inherit
}
.news .alert .octicon {
    width: 16px;
    height: 16px;
    color: #bbb
}
.news .alert .body {
    padding: 1em 0 0;
    overflow: hidden;
    font-size: 14px;
    border-bottom: 0
}
.news .alert .time {
    font-size: 12px;
    color: #bbb
}
.news .alert .title {
    padding: 0;
    font-weight: bold
}
.news .alert .title .subtle {
    color: #bbb
}
.news .alert .gravatar {
    float: left;
    margin-right: 0.6em;
    line-height: 0;
    background-color: #fff;
    border-radius: 3px
}
.news .alert .simple>.octicon {
    position: absolute !important;
    left: 11px;
    width: 16px;
    height: 16px
}
.news .alert .simple .title {
    display: inline-block;
    font-size: 13px;
    font-weight: normal;
    color: #666
}
.news .alert .simple .time {
    display: inline-block
}
.news .alert .branch-link,
.news .alert .pull-info {
    display: inline-block;
    padding: 3px 7px;
    margin-top: 5px;
    font-size: 12px;
    color: rgba(0, 0, 0, 0.5);
    background: #e8f1f6;
    border-radius: 3px
}
.news .alert .branch-link em,
.news .alert .pull-info em {
    font-style: normal;
    font-weight: bold
}
.news .alert .branch-link {
    position: relative;
    top: -2px;
    margin: 0;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    color: #4183c4
}
.news .alert .branch-link .octicon {
    display: none
}
.news .alert:first-child {
    border-top: 0
}
.news .alert:first-child .body {
    padding-top: 0
}
.news .alert:first-child .mega-octicon {
    top: 0
}
.news .git_hub .done {
    color: #666;
    text-decoration: line-through
}
.news .commits li {
    margin-top: 0.15em;
    list-style-type: none
}
.news .commits li.more {
    padding-top: 2px;
    font-size: 11px
}
.news .commits li .committer {
    display: none;
    padding-left: 0.5em
}
.news .commits li img {
    margin: 0 1px 0 0;
    vertical-align: middle;
    background-color: #fff;
    border-radius: 2px
}
.news .commits li img.emoji {
    padding: 0;
    margin: 0;
    border: 0
}
.news .commits li .message {
    display: inline-block;
    max-width: 390px;
    margin-top: 2px;
    overflow: hidden;
    font-size: 13px;
    line-height: 1.3;
    text-overflow: ellipsis;
    white-space: nowrap;
    vertical-align: top
}
.news div.message,
.news li blockquote {
    display: inline;
    font-size: 13px;
    color: #666
}
.release-assets {
    padding-left: 40px
}
.release-assets li {
    margin-top: 0.15em;
    list-style-type: none
}
.release-assets .more {
    padding-top: 2px;
    font-size: 11px
}
.news-full,
.page-profile .news {
    float: none;
    width: auto
}
.activity-tab .blankslate {
    margin-top: 10px
}
.activity-tab .news .markdown-body blockquote,
.activity-tab .news .alert .commits {
    padding-left: 0
}
.activity-tab .news a.gravatar,
.activity-tab .news div.gravatar {
    display: none
}
.saml-signed-out-notice {
    position: relative;
    width: 450px;
    padding: 10px 10px 10px 70px;
    margin: 50px auto 30px;
    border: 1px solid #eee;
    border-radius: 3px
}
.saml-signed-out-notice .mega-octicon {
    position: absolute;
    top: 30px;
    left: 20px;
    color: #ddd
}
.saml-signed-out-notice h3 {
    margin-bottom: 0
}
.saml-signed-out-notice p {
    margin-top: 5px
}
.survey-box.simple-box {
    max-width: 400px;
    position: fixed;
    right: 25px;
    bottom: -20px;
    z-index: 25;
    padding-bottom: 0;
    border-bottom: 0;
    border-radius: 4px 4px 0 0;
    background-color: #f9f9f9;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
    -webkit-animation: slideUp 0.5s ease 1.25s both;
    animation: slideUp 0.5s ease 1.25s both;
    -webkit-transform: translate3d(0, 0, 0);
    transform: translate3d(0, 0, 0)
}
.survey-box.simple-box .simple-box-title {
    margin-bottom: -7px;
    padding-bottom: 0;
    border-bottom: 0;
    font-size: 14px;
    line-height: 1.3
}
.survey-box.simple-box .simple-box-footer {
    margin: 0 -15px;
    padding: 7px;
    background-color: #fff
}
.survey-box.simple-box .close-button {
    margin-top: -5px;
    color: #aaa
}
@-webkit-keyframes slideUp {
    0% {
        -webkit-transform: translate3d(0, 500px, 0);
        transform: translate3d(0, 500px, 0)
    }
    100% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@keyframes slideUp {
    0% {
        -webkit-transform: translate3d(0, 500px, 0);
        transform: translate3d(0, 500px, 0)
    }
    100% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
.dashboards-overview-lead {
    width: 700px
}
.dashboards-overview-cards .boxed-group {
    margin: 10px 0;
    width: 100%
}
.dashboards-overview-cards .boxed-group path {
    stroke: #1db34f;
    stroke-opacity: 0.5
}
.dashboards-overview-cards .blankslate {
    border: 0;
    background-color: #fff;
    box-shadow: none;
    padding-top: 47px
}
.dashboards-overview-cards .octicon-arrow-down {
    color: #bd2c00
}
.dashboards-overview-cards .octicon-arrow-up {
    color: #1db34f
}
.dashboards-overview-cards .graph-canvas .dots {
    padding: 43px 0
}
.dashboards-overview-cards .summary-stats {
    height: 78px
}
.dashboards-overview-cards .summary-stats .created_at {
    color: #1db34f
}
.dashboards-overview-cards .summary-stats .closed_at,
.dashboards-overview-cards .summary-stats .merged_at {
    color: #4078c0
}
.dashboards-overview-cards .summary-stats .totals-num {
    margin: 0 7px
}
.dashboards-overview-cards .summary-stats .single {
    width: 100%
}
.dashboards-overview-graph {
    height: 160px
}
.dashboards-overview-graph .path {
    fill: none;
    stroke-width: 2
}
.dashboards-overview-graph path.created_at {
    stroke: #1db34f
}
.dashboards-overview-graph path.merged_at,
.dashboards-overview-graph path.closed_at {
    stroke: #1d7fb3
}
.dashboards-overview-graph .y line {
    stroke: #1db34f
}
.dashboards-overview-graph .y.unique line {
    stroke: #1d7fb3
}
.dashboards-overview-graph .overlay {
    fill-opacity: 0
}
.created_at circle {
    fill: #1db34f;
    stroke: #fff;
    stroke-width: 2
}
.merged_at circle,
.closed_at circle {
    fill: #1d7fb3;
    stroke: #fff;
    stroke-width: 2
}
.diffstat {
    font-size: 12px;
    font-weight: bold;
    color: #666;
    white-space: nowrap;
    cursor: default
}
.diffstat-bar {
    display: inline-block;
    margin-left: 3px;
    font-size: 16px;
    color: #ddd;
    letter-spacing: 1px;
    text-align: left;
    text-decoration: none;
    font: normal normal 16px/1 "octicons";
    display: inline-block;
    text-decoration: none;
    -webkit-font-smoothing: antialiased
}
.discussion-timeline {
    position: relative;
    width: 760px;
    float: left
}
.discussion-timeline:before {
    display: block;
    content: "";
    position: absolute;
    top: 0;
    bottom: 0;
    left: 79px;
    width: 2px;
    background-color: #f3f3f3;
    z-index: -1
}
.discussion-timeline .email-hidden-container {
    margin: 3px 0
}
.discussion-sidebar {
    position: -webkit-sticky;
    position: sticky;
    top: 0;
    z-index: 21;
    float: right;
    width: 140px
}
.discussion-sidebar-wide {
    width: 200px
}
.discussion-sidebar-item {
    padding-top: 15px;
    font-size: 12px;
    color: #767676
}
.discussion-sidebar-item .btn .octicon {
    margin-right: 0
}
.discussion-sidebar-item .btn-block {
    margin-bottom: 8px
}
.discussion-sidebar-item+.discussion-sidebar-item {
    margin-top: 15px;
    border-top: 1px solid #eee
}
.discussion-sidebar-item .select-menu {
    position: relative
}
.discussion-sidebar-item .select-menu-modal-holder {
    top: 25px;
    left: auto;
    right: -1px
}
.discussion-sidebar-heading {
    margin-top: 0;
    margin-bottom: 10px;
    font-size: 12px;
    line-height: 16px;
    color: #767676
}
.discussion-sidebar-toggle {
    padding: 5px 0;
    margin: -5px 0 5px
}
.discussion-sidebar-toggle .octicon {
    float: right;
    padding: 5px;
    margin: -6px -5px -5px 5px;
    color: #ccc
}
.discussion-sidebar-toggle:hover {
    color: #4078c0;
    text-decoration: none;
    cursor: pointer
}
.discussion-sidebar-toggle:hover .octicon {
    color: inherit
}
button.discussion-sidebar-toggle {
    display: block;
    width: 100%;
    font-weight: bold;
    text-align: left;
    background: none;
    border: 0
}
.sidebar-labels .labels .label {
    display: block;
    max-width: 100%;
    padding: 6px 10px;
    font-size: 12px;
    box-shadow: none
}
.sidebar-labels .labels .label+.label {
    margin-top: 3px
}
.sidebar-milestone .progress-bar {
    margin-bottom: 2px;
    border-radius: 2px;
    height: 8px
}
.milestone-name {
    display: block;
    margin-top: 5px;
    font-weight: bold;
    color: #555
}
.milestone-name .css-truncate-target {
    max-width: 100%
}
.milestone-name:hover {
    color: #4078c0;
    text-decoration: none
}
.sidebar-assignee .css-truncate-target {
    max-width: 110px
}
.sidebar-assignee .avatar {
    margin-top: -1px;
    margin-right: 2px;
    border-radius: 2px
}
.sidebar-assignee .assignee {
    color: #555;
    font-weight: bold
}
.sidebar-assignee .assignee:hover {
    color: #4078c0;
    text-decoration: none
}
.sidebar-notifications {
    position: relative
}
.sidebar-notifications .thread-subscription-status {
    margin: 0;
    padding: 0;
    border: 0
}
.sidebar-notifications .thread-subscription-status .thread-subscribe-form {
    display: block
}
.sidebar-notifications .thread-subscription-status .mega-octicon {
    display: none
}
.sidebar-notifications .thread-subscription-status .reason {
    padding: 0;
    margin: 5px 0 0;
    font-size: 11px;
    color: #767676
}
.sidebar-notifications .thread-subscription-status .btn-sm {
    display: block;
    width: 100%
}
.participation .participant-avatar {
    float: left;
    margin: 3px 0 0 3px
}
.participation a {
    color: #767676
}
.participation a:hover {
    color: #4078c0;
    text-decoration: none
}
.participation-avatars {
    margin-left: -3px
}
.participation-avatars:before {
    display: table;
    content: ""
}
.participation-avatars:after {
    display: table;
    clear: both;
    content: ""
}
.participation-more {
    float: left;
    margin: 6px 0 0;
    line-height: 14px
}
.lock-toggle-link {
    color: #767676;
    font-weight: bold
}
.lock-toggle-link:hover {
    color: #4078c0;
    text-decoration: none
}
.inline-comment-form .form-actions,
.timeline-new-comment .form-actions {
    padding: 0 10px 10px
}
.gh-header-actions {
    float: right;
    margin-top: 3px
}
.gh-header-actions .btn-sm {
    float: left;
    margin-left: 5px
}
.gh-header-actions .btn-sm .octicon {
    margin-right: 0
}
.gh-header .gh-header-edit {
    display: none
}
.gh-header.open .gh-header-show {
    display: none
}
.gh-header.open .gh-header-edit {
    display: block
}
.gh-header-title {
    margin-top: 0;
    margin-bottom: 0;
    margin-right: 150px;
    font-weight: normal;
    line-height: 1.1;
    word-wrap: break-word
}
.gh-header-no-access .gh-header-title {
    margin-right: 0
}
.gh-header-number {
    font-weight: 300;
    color: #aaa;
    letter-spacing: -1px
}
.gh-header-edit {
    margin-top: -5px
}
.gh-header-edit:before {
    display: table;
    content: ""
}
.gh-header-edit:after {
    display: table;
    clear: both;
    content: ""
}
.gh-header-edit .edit-issue-title {
    float: left;
    width: 760px;
    padding: 6px 10px;
    margin-right: 10px;
    font-size: 16px;
    background-color: #fafafa
}
.gh-header-edit .edit-issue-title:focus {
    background-color: #fff
}
.gh-header-edit .btn {
    float: left;
    padding: 7px 15px
}
.gh-header-edit .btn-link {
    float: left;
    margin: 9px 10px
}
.gh-header-meta {
    margin-top: 9px;
    font-size: 14px;
    line-height: 20px;
    color: #767676;
    padding-bottom: 20px;
    border-bottom: 1px solid #eee
}
.gh-header.issue .gh-header-meta {
    margin-bottom: 15px
}
.gh-header.pull .gh-header-meta {
    border-bottom: 0;
    padding-bottom: 0
}
.gh-header-meta .flex-table-item {
    vertical-align: top
}
.gh-header-meta .flex-table-item-primary {
    padding-top: 4px;
    white-space: normal;
    word-wrap: break-word
}
.gh-header-meta .flex-table-item-primary .commit-ref .css-truncate-target,
.gh-header-meta .flex-table-item-primary .commit-ref:hover .css-truncate-target {
    max-width: 780px !important
}
.gh-header-meta .state {
    margin-right: 8px
}
.gh-header-meta .avatar {
    float: left;
    margin-top: -3px;
    margin-right: 5px
}
.gh-header-meta .author {
    color: #555;
    font-weight: bold
}
.gh-header-meta .noun {
    text-transform: lowercase
}
.tabnav-pr {
    margin: 15px 0 20px;
    border-color: #e5e5e5
}
.tabnav-pr .tabnav-tab {
    position: relative;
    padding: 9px 14px;
    font-size: 13px;
    color: #767676
}
.tabnav-pr .tabnav-tab.selected {
    color: #333;
    border-color: #e5e5e5
}
.tabnav-pr .diffstat-bar {
    padding-bottom: 3px
}
.timeline-comment-wrapper>.timeline-comment:after,
.timeline-comment-wrapper>.timeline-comment:before,
.timeline-new-comment .timeline-comment:after,
.timeline-new-comment .timeline-comment:before {
    position: absolute;
    top: 11px;
    left: -16px;
    right: 100%;
    width: 0;
    height: 0;
    display: block;
    content: " ";
    border-color: transparent;
    border-style: solid solid outset;
    pointer-events: none
}
.timeline-comment-wrapper>.timeline-comment:after,
.timeline-new-comment .timeline-comment:after {
    border-width: 7px;
    border-right-color: #f7f7f7;
    margin-top: 1px;
    margin-left: 2px
}
.timeline-comment-wrapper>.timeline-comment:before,
.timeline-new-comment .timeline-comment:before {
    border-right-color: #ddd;
    border-width: 8px
}
.timeline-comment-wrapper {
    position: relative;
    padding-left: 64px;
    margin-top: 15px;
    margin-bottom: 15px;
    border-top: 2px solid #fff;
    border-bottom: 2px solid #fff
}
.timeline-comment-wrapper:first-child {
    margin-top: 0
}
.discussion-timeline-actions .timeline-comment-wrapper:first-child {
    margin-top: 15px
}
.timeline-comment-wrapper .timeline-comment.current-user:after,
.timeline-comment-wrapper .timeline-comment.current-user:before {
    position: absolute;
    top: 11px;
    left: -16px;
    right: 100%;
    width: 0;
    height: 0;
    display: block;
    content: " ";
    border-color: transparent;
    border-style: solid solid outset;
    pointer-events: none
}
.timeline-comment-wrapper .timeline-comment.current-user:after {
    border-width: 7px;
    border-right-color: #f2f8fa;
    margin-top: 1px;
    margin-left: 2px
}
.timeline-comment-wrapper .timeline-comment.current-user:before {
    border-right-color: #bfccd1;
    border-width: 8px
}
.timeline-comment-wrapper .timeline-comment.unread-item:after,
.timeline-comment-wrapper .timeline-comment.unread-item:before {
    position: absolute;
    top: 11px;
    left: -16px;
    right: 100%;
    width: 0;
    height: 0;
    display: block;
    content: " ";
    border-color: transparent;
    border-style: solid solid outset;
    pointer-events: none
}
.timeline-comment-wrapper .timeline-comment.unread-item:after {
    border-width: 7px;
    border-right-color: #fff9ea;
    margin-top: 1px;
    margin-left: 2px
}
.timeline-comment-wrapper .timeline-comment.unread-item:before {
    border-right-color: #dfd8c2;
    border-width: 8px
}
.timeline-comment-avatar {
    float: left;
    margin-left: -64px;
    border-radius: 3px
}
.timeline-comment {
    position: relative;
    background-color: #fff;
    border: 1px solid #ddd;
    border-radius: 3px
}
.timeline-comment.will-transition-once {
    -webkit-transition: border-color 0.65s ease-in-out;
    transition: border-color 0.65s ease-in-out
}
.timeline-comment.will-transition-once .timeline-comment-header {
    -webkit-transition: background-color 0.65s ease, border-bottom-color 0.65s ease-in-out;
    transition: background-color 0.65s ease, border-bottom-color 0.65s ease-in-out
}
.timeline-comment.will-transition-once .timeline-comment-label {
    -webkit-transition: border-color 0.65s ease-in-out;
    transition: border-color 0.65s ease-in-out
}
.timeline-comment.will-transition-once:before,
.timeline-comment.will-transition-once:after {
    -webkit-transition: border-right-color 0.65s ease-in-out;
    transition: border-right-color 0.65s ease-in-out
}
.timeline-comment.current-user {
    border-color: #bfccd1
}
.timeline-comment.current-user .timeline-comment-header {
    background-color: #f2f8fa;
    border-bottom-color: #dde4e6
}
.timeline-comment.current-user .timeline-comment-label {
    border-color: #bfccd1
}
.timeline-comment.current-user .previewable-comment-form .comment-form-head.tabnav {
    color: #8e9597;
    background-color: #f2f8fa;
    border-bottom-color: #e1edf1
}
.timeline-comment.unread-item {
    border-color: #dfd8c2
}
.timeline-comment.unread-item .timeline-comment-header {
    background-color: #fff9ea;
    border-bottom-color: #f1ede3
}
.timeline-comment.unread-item .timeline-comment-label {
    border-color: #dfd8c2
}
.timeline-comment.unread-item .previewable-comment-form .comment-form-head.tabnav {
    color: #8e9597;
    background-color: #f2f8fa;
    border-bottom-color: #e1edf1
}
.timeline-comment:empty {
    display: none
}
.timeline-comment .comment+.comment {
    border-top: 1px solid #e5e5e5
}
.timeline-comment .comment+.comment:before,
.timeline-comment .comment+.comment:after {
    display: none
}
.timeline-comment .comment+.comment .timeline-comment-header {
    border-top-left-radius: 0;
    border-top-right-radius: 0
}
.timeline-comment-header {
    padding-left: 15px;
    padding-right: 15px;
    color: #767676;
    background-color: #f7f7f7;
    border-bottom: 1px solid #eee;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.timeline-comment-header .author {
    font-weight: bold;
    color: #555
}
.timeline-comment-header .timestamp {
    white-space: nowrap;
    color: inherit
}
.timeline-comment-header code {
    word-break: break-all
}
.comment-type-icon {
    color: inherit
}
.timeline-comment-label {
    float: right;
    margin: 8px 0 0 10px;
    padding: 2px 5px;
    font-size: 12px;
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 3px
}
.timeline-comment-label-spammy {
    color: #bd2c00;
    border-color: #bd2c00
}
.timeline-comment-header-text {
    max-width: 78%;
    padding-top: 10px;
    padding-bottom: 10px
}
.timeline-comment-header-text code a {
    color: #555
}
.timeline-comment-header-avatar {
    float: left;
    margin-top: 10px;
    margin-right: 5px
}
.timeline-comment-actions {
    float: right;
    margin-right: -5px;
    margin-left: 10px
}
.timeline-comment-action {
    display: inline-block;
    padding: 10px 5px;
    color: inherit;
    opacity: 0.5
}
.timeline-comment-action:hover {
    opacity: 1;
    color: #4078c0;
    text-decoration: none
}
.timeline-comment-action .octicon-check {
    height: 16px;
    font-size: 18px
}
.compare-tab-comments .timeline-comment-actions {
    display: none
}
.discussion-item-ref .commit-gravatar {
    padding-left: 2px;
    padding-right: 5px
}
.discussion-item-ref .task-progress {
    display: block;
    margin-bottom: -2px
}
.discussion-item-ref .task-progress .progress-bar {
    margin-bottom: 0
}
.discussion-item-ref .task-progress .octicon {
    font-size: 16px
}
.discussion-item-ref .discussion-item-body .title {
    margin-top: 10px
}
.discussion-item-ref .state {
    padding: 1px 5px;
    margin-top: -4px;
    margin-left: 8px;
    font-size: 12px
}
.discussion-item-ref .state .octicon {
    width: 1em;
    font-size: 14px
}
.timeline-new-comment {
    max-width: 780px;
    margin-bottom: 0
}
.timeline-new-comment .comment-form-head {
    margin-bottom: 10px
}
.timeline-new-comment .previewable-comment-form .comment-body {
    padding: 5px 5px 15px;
    border-bottom: 1px solid #eee
}
.discussion-item {
    position: relative;
    margin: 15px 0 15px 79px;
    padding-left: 25px
}
.discussion-item+.discussion-item {
    padding-top: 15px;
    border-top: 1px solid #f5f5f5
}
.discussion-item .author {
    color: #555;
    font-weight: bold
}
.discussion-item .timestamp {
    color: inherit;
    white-space: nowrap
}
.discussion-item .label-color {
    padding: 2px 4px;
    font-size: 12px;
    font-weight: bold;
    border-radius: 2px;
    box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.12)
}
.discussion-item .label-color a:hover {
    text-decoration: none
}
.discussion-item.open .discussion-item-details {
    display: block
}
.discussion-item.open .discussion-item-toggler-opened {
    display: inline
}
.discussion-item.open .discussion-item-toggler-closed {
    display: none
}
.discussion-item-details {
    display: none
}
.discussion-item-toggler-opened {
    display: none
}
.discussion-item-icon {
    float: left;
    width: 32px;
    height: 32px;
    margin-top: -7px;
    margin-left: -40px;
    line-height: 28px;
    color: #767676;
    text-align: center;
    background-color: #f3f3f3;
    border: 2px solid #fff;
    border-radius: 50%
}
.discussion-item-icon.octicon-pencil {
    font-size: 14px
}
.discussion-item-header {
    min-height: 30px;
    padding-top: 5px;
    padding-bottom: 5px;
    color: #767676;
    line-height: 20px;
    word-wrap: break-word
}
.discussion-item-header .avatar {
    float: left;
    margin-top: 2px;
    margin-right: 5px
}
.discussion-item-header .discussion-item-private {
    vertical-align: -1px
}
.discussion-item-header:last-child {
    padding-bottom: 0
}
.discussion-item-header .commit-ref {
    font-size: 85%;
    vertical-align: baseline
}
.discussion-item-header .btn-outline {
    float: right;
    padding: 4px 8px;
    margin-top: -5px;
    margin-left: 10px
}
.discussion-item-body {
    margin-top: 5px
}
.discussion-item-footer {
    padding-left: 21px;
    font-size: 12px
}
.discussion-item-link {
    color: #767676
}
.discussion-item-link:hover {
    color: #4078c0
}
.discussion-item-entity {
    font-weight: bold;
    color: #333
}
.discussion-item-entity:hover {
    color: #4078c0;
    text-decoration: none
}
.discussion-item-ref-title {
    margin-top: 0;
    margin-bottom: 0;
    line-height: 1.2
}
.discussion-item-ref-title .issue-num {
    font-weight: normal;
    color: #767676
}
.discussion-item-ref-title .title-link {
    color: #333
}
.discussion-item-ref-title .title-link:hover {
    color: #4078c0;
    text-decoration: none
}
.discussion-item-ref-title .title-link:hover .issue-num {
    color: inherit
}
.discussion-item-context-icon {
    display: inline-block;
    line-height: 22px;
    margin-top: -2px;
    margin-left: 10px
}
.discussion-item-help {
    color: #767676
}
.discussion-item-help:hover {
    color: #4078c0
}
.discussion-item-private {
    color: #a1882b
}
.discussion-item-rollup-ref .state {
    margin-top: 2px
}
.discussion-item-rollup-ref .discussion-item-context-icon {
    margin-top: 2px
}
.discussion-item-reopened .discussion-item-icon {
    color: #fff;
    background-color: #6cc644
}
.discussion-item-closed .discussion-item-icon {
    color: #fff;
    background-color: #bd2c00
}
.discussion-item-head_ref_deleted .discussion-item-icon {
    padding-left: 2px;
    color: #fff;
    background-color: #767676
}
.discussion-item-locked .discussion-item-icon,
.discussion-item-unlocked .discussion-item-icon {
    color: #fff;
    background-color: #333
}
.discussion-item .renamed-was,
.discussion-item .renamed-is {
    color: #333;
    font-weight: bold
}
.discussion-commits .discussion-item-icon {
    padding-top: 1px
}
.discussion-commits .discussion-item-body {
    margin-top: 0;
    margin-left: -31px
}
.discussion-item-toggle-open {
    display: none
}
.discussion-item-toggle {
    float: right;
    color: #767676
}
.discussion-item-toggle:hover {
    color: #4078c0;
    text-decoration: none
}
.discussion-item-toggle .octicon {
    vertical-align: middle
}
.outdated-diff-comment-container .discussion-item-body {
    display: none
}
.outdated-diff-comment-container.open .discussion-item-body,
.outdated-diff-comment-container.open .discussion-item-toggle-open {
    display: block
}
.outdated-diff-comment-container.open .discussion-item-toggle-closed {
    display: none
}
.new-discussion-timeline .previewable-comment-form .comment-form-head.tabnav {
    background: #f7f7f7;
    padding: 6px 10px 0;
    border-radius: 3px 3px 0 0
}
.new-discussion-timeline .previewable-comment-form .draft-indicator {
    position: relative;
    top: -1px
}
.new-discussion-timeline .previewable-comment-form .comment {
    border: 0
}
.new-discussion-timeline .previewable-comment-form .comment-body {
    padding: 5px 5px 15px;
    border-bottom: 1px solid #eee;
    background-color: transparent
}
.new-discussion-timeline .previewable-comment-form .timeline-comment .timeline-comment-actions {
    display: none
}
.new-discussion-timeline .closed-banner {
    position: relative;
    margin: 15px 0 -15px;
    height: 19px;
    overflow: visible;
    background: #f3f3f3;
    border-radius: 0;
    border-bottom: 15px solid #fff
}
.new-discussion-timeline .composer .timeline-comment {
    margin-bottom: 10px
}
.new-discussion-timeline .composer .timeline-comment:after {
    border-right-color: #fff
}
.new-discussion-timeline .composer .comment-form-head.tabnav {
    padding-top: 0;
    background-color: #fff
}
.new-discussion-timeline .composer.has-checklist .form-actions {
    padding-top: 10px;
    border-top: 1px solid #eee;
    border-radius: 0 0 3px 3px;
    background-color: #fafafa
}
.discussion-timeline-actions {
    border-top: 2px solid #f3f3f3;
    background-color: #fff
}
.discussion-timeline-actions .merge-pr {
    padding-top: 0;
    border-top: 0
}
.discussion-timeline-actions .thread-subscription-status {
    margin-top: 20px
}
.discussion-timeline-actions .thread-subscription-status .mega-octicon {
    display: none
}
.discussion-item-merged .discussion-item-icon {
    padding-left: 2px;
    color: #fff;
    background-color: #6e5494
}
.discussion-item-merged.open .discussion-item-footer {
    display: none
}
.discussion-item-merged.open .discussion-item-details {
    margin-top: 5px;
    margin-bottom: 10px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.discussion-item-merged.open .discussion-item-details-header {
    padding: 12px 15px;
    margin-top: 0;
    margin-bottom: 0;
    font-size: inherit;
    border-top: 1px solid #ddd
}
.discussion-item-merged.open .discussion-item-details-header:first-child {
    border-top: 0
}
.discussion-item-merged.open .build-statuses-list {
    max-height: 370px;
    margin: 0;
    border-top-color: #ddd
}
.discussion-item-merged.open .build-status-item {
    padding-left: 15px
}
.donut-chart>.error,
.donut-chart>.failure {
    fill: #bd2c00
}
.donut-chart>.pending {
    fill: #cea61b
}
.donut-chart>.success {
    fill: #6cc644
}
.survey-question-form .other-text-form {
    display: none;
    margin-top: 0
}
.survey-question-form.is-other-selected .other-text-form {
    display: inline-block
}
.setup-header .large-file-storage-header {
    font-size: 44px
}
.early-acccess-setup-form .form {
    margin-top: 0;
    margin-bottom: 30px
}
.early-acccess-setup-form select {
    display: block;
    width: 200px
}
.early-access-setup-list {
    padding: 0 15px 15px;
    margin: 0;
    font-size: 14px
}
.early-access-setup-list .early-access-setup-list-item {
    margin-top: 10px;
    margin-left: 20px
}
.early-access-setup-list .early-access-setup-list-item:first-child {
    margin-top: 0
}
.early-access-thanks-wrapper {
    position: relative;
    height: 80vh;
    z-index: 1;
    margin-bottom: -41px;
    background-color: #fcfcfc;
    border-bottom: 1px solid #ddd
}
.early-access-thanks-content {
    position: relative;
    top: 50%;
    margin: 0 auto;
    width: 500px;
    -webkit-transform: translateY(-50%);
    -ms-transform: translateY(-50%);
    transform: translateY(-50%)
}
.early-access-thanks-content .simple-box {
    padding: 30px;
    font-size: 16px
}
.early-access-thanks-title {
    margin-top: 0;
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    font-weight: normal
}
.early-access-thanks-lead {
    margin-top: 0;
    margin-bottom: 0
}
.eap-error-state-title {
    margin-top: 0
}
.explore-head .container {
    position: relative
}
.explore-content {
    margin-top: -15px
}
.explore-content .blankslate {
    margin-top: 15px
}
.repo-collection>ul {
    list-style-type: none;
    background: #f7f7f7;
    border: 1px solid #ddd;
    border-radius: 3px
}
.repo-collection .author-gravatar {
    float: left;
    margin-right: 10px;
    background: #fff;
    border-radius: 3px
}
.collection-stat {
    float: right;
    margin-left: 10px;
    font-size: 12px;
    color: #444
}
.collection-stat .octicon {
    margin-right: 5px;
    color: #a7a7a7
}
.collection-item {
    position: relative;
    float: left;
    width: 50%;
    height: 70px;
    padding: 15px
}
.collection-item .octicon-x {
    position: absolute;
    top: 10px;
    right: 10px;
    color: #ccc;
    text-decoration: none
}
.collection-item .repo-name {
    display: block;
    font-size: 16px;
    font-weight: bold
}
.collection-item .css-truncate-target {
    max-width: 380px
}
.collection-item .repo-description {
    margin: 0
}
.explore-collection h2 {
    margin: 0 0 10px;
    font-size: 18px;
    font-weight: normal;
    color: #2a2a2a
}
.explore-collection h2 .select-menu {
    position: relative;
    display: inline-block
}
.explore-collection h2 .select-menu-button {
    font-weight: bold;
    cursor: pointer
}
.explore-collection h2 .mega-octicon {
    vertical-align: middle
}
.explore-collection .see-more-link {
    float: right;
    margin-top: 7px
}
.explore-page .see-more-link {
    font-size: 14px;
    color: inherit
}
.explore-page .see-more-link .octicon {
    margin-left: 5px
}
.explore-page.marketing-section {
    border-bottom: 0
}
.explore-page.marketing-section .thread-subscription-status {
    border: 0
}
.explore-page.marketing-section .signed-out-comment {
    margin-left: 0
}
.explore-page .language-filter-list {
    margin-bottom: 10px
}
.explore-section {
    position: relative;
    padding: 40px 0;
    border-bottom: 1px solid #eee
}
.explore-section:nth-child(even) {
    background: #f9f9f9
}
.explore-section:nth-child(even) .repo-collection>ul {
    background: #fff
}
.explore-section:first-child {
    padding-top: 0
}
.explore-section:nth-child(odd):last-child {
    padding-bottom: 0;
    border-bottom: 0
}
.explore-pjax-container {
    position: relative
}
.user-leaderboard-list .follow-list-info {
    margin-top: 12px;
    margin-bottom: 0;
    font-size: 12px;
    color: #666
}
.user-leaderboard-list .follow-list-info .css-truncate.css-truncate-target {
    max-width: none
}
.user-leaderboard-list .repo-list-item {
    padding-top: 10px;
    padding-bottom: 0;
    padding-left: 21px;
    border-top: 0
}
.user-leaderboard-list .repo-list-item .repo-description,
.user-leaderboard-list .repo-list-item .repo-and-owner {
    max-width: 530px
}
.user-leaderboard-list .repo-list-item .repo {
    color: #5c5c5c
}
.leaderboard-list {
    margin: 0;
    list-style-type: none
}
.user-leaderboard-list-name {
    margin: 0;
    font-size: 18px;
    font-weight: normal
}
.user-leaderboard-list-name .full-name {
    margin-left: 5px;
    font-weight: bold;
    color: #5c5c5c
}
.repo-snipit {
    display: inline-block;
    margin-top: 7px
}
.repo-snipit:hover {
    text-decoration: none
}
.repo-snipit .octicon {
    font-size: 14px;
    color: #767676
}
.repo-snipit-name {
    max-width: 200px;
    color: #666
}
.repo-snipit-description {
    max-width: 300px;
    color: #767676
}
.repo-snipit:hover .repo-snipit-name,
.repo-snipit:hover .repo-snipit-description {
    color: #4078c0
}
.leaderboard-action {
    float: right;
    margin-top: -3px;
    margin-left: 10px
}
.leaderboard-list-rank {
    position: absolute;
    top: 25px;
    left: 0;
    width: 20px;
    font-size: 11px;
    font-weight: 300;
    color: #b9b9b9;
    text-align: right;
    text-transform: uppercase
}
.leaderboard-list-item {
    position: relative;
    padding-top: 20px;
    padding-bottom: 20px;
    padding-left: 35px;
    border-bottom: 1px solid #eee
}
.leaderboard-list-item:last-child {
    border-bottom: 0
}
.leaderboard-gravatar {
    float: left;
    width: 48px;
    height: 48px;
    border-radius: 3px
}
.leaderboard-list-content {
    min-height: 48px;
    margin-left: 58px
}
.collection-page .signed-out-comment {
    margin-left: 0
}
.explore-mail-tease {
    padding-top: 20px;
    overflow: hidden;
    background: #202021 url(https://github.com/images/modules/home/octicons-bg.png) center repeat;
    border-bottom: 1px solid #ddd
}
.explore-mail-tease h3 {
    color: #fff;
    text-align: center
}
.explore-mail-tease img {
    margin-bottom: -5px
}
.newsletter-frequency-choice {
    display: -webkit-box;
    display: -webkit-flex;
    display: -ms-flexbox;
    display: flex;
    margin: 40px 0;
    list-style-type: none;
    -webkit-flex-flow: row wrap;
    -ms-flex-flow: row wrap;
    flex-flow: row wrap;
    -webkit-justify-content: space-around;
    -ms-flex-pack: distribute;
    justify-content: space-around
}
.newsletter-frequency-choice .choice {
    position: relative;
    width: 25%
}
.newsletter-frequency-choice .choice>label {
    display: block;
    height: 100%;
    margin: 0 10px;
    font-weight: normal;
    text-align: center;
    cursor: pointer;
    background: #fff;
    border: 3px solid #eee;
    border-radius: 4px
}
.newsletter-frequency-choice .choice:hover label {
    border-color: #4078c0
}
.newsletter-frequency-choice .choice:hover h3 {
    color: #fff;
    background: #4078c0;
    border-color: #4078c0
}
.newsletter-frequency-choice .choice.selected label {
    border-color: #6cc644;
    box-shadow: 0 0 5px rgba(0, 0, 0, 0.2)
}
.newsletter-frequency-choice .choice.selected h3 {
    color: #fff;
    background: #6cc644;
    border-color: #6cc644
}
.newsletter-frequency-choice .choice.selected p {
    color: #333
}
.newsletter-frequency-choice .choice .notice {
    position: absolute;
    right: 0;
    bottom: 1em;
    left: 0;
    z-index: -1;
    font-weight: bold;
    color: #6cc644;
    text-align: center;
    opacity: 0
}
.newsletter-frequency-choice .choice .notice.visible {
    bottom: -2em;
    opacity: 1;
    -webkit-transition: opacity 0.15s ease-in-out;
    transition: opacity 0.15s ease-in-out
}
.newsletter-frequency-choice h3 {
    padding: 10px;
    margin: 0;
    font-weight: normal;
    background: #fafafa;
    border-bottom: 1px solid #eee
}
.newsletter-frequency-choice h3 input {
    position: relative;
    top: -2px;
    margin: 0 3px 0 -19px
}
.newsletter-frequency-choice p {
    height: 7em;
    margin: 15px;
    color: #767676;
    text-align: left
}
.explore-signup-entice {
    position: relative;
    padding: 15px;
    font-size: 14px;
    background: #f7f7f7;
    border: 1px solid #ddd;
    border-radius: 3px
}
.explore-signup-entice h3 {
    margin-bottom: 10px;
    font-size: 18px
}
.explore-signup-entice-inner {
    position: absolute;
    top: 3px;
    right: 3px;
    bottom: 3px;
    left: 3px;
    padding-top: 30px;
    text-align: center;
    background: rgba(247, 247, 247, 0.9)
}
.explore-signup-entice-wrapper {
    max-width: 500px;
    padding: 5px;
    margin: 0 auto;
    background: rgba(247, 247, 247, 0.6)
}
.explore-signup-cta {
    margin-right: -10px;
    font-size: 13px;
    vertical-align: middle
}
.explore-signup-cta a {
    font-weight: bold
}
.explore-signup-cta .btn {
    position: relative;
    top: -1px
}
@-webkit-keyframes fadein {
    0% {
        opacity: 0
    }
    100% {
        opacity: 1
    }
}
@keyframes fadein {
    0% {
        opacity: 0
    }
    100% {
        opacity: 1
    }
}
.explore-marketing-header {
    margin: 10px auto 30px;
    text-align: center
}
.explore-marketing-header.is-animating {
    -webkit-animation: fadein 1s;
    animation: fadein 1s
}
.explore-marketing-header h2 {
    margin: 0 0 5px;
    font-size: 32px;
    font-weight: normal
}
.explore-marketing-header .lead {
    margin: 5px 0 0
}
.linux .show-mac,
.macintosh .show-mac {
    display: block
}
.linux .hide-mac,
.macintosh .hide-mac {
    display: none
}
.windows .show-mac {
    display: none
}
.homepage .container {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif
}
.homepage .header-logged-out .primary {
    display: none
}
.homepage .site-footer {
    border-top: 0;
    margin-top: 0
}
.marketing-section-depth {
    position: absolute;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 10;
    height: 30px;
    background-image: -webkit-linear-gradient(transparent, rgba(0, 0, 0, 0.15));
    background-image: linear-gradient(transparent, rgba(0, 0, 0, 0.15));
    box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.25)
}
.marketing-section-signup {
    text-shadow: 0 1px 3px #222;
    background: #202021 url(https://github.com/images/modules/home/octicons-bg.png) center repeat;
    padding-top: 80px;
    padding-bottom: 80px;
    position: relative
}
.marketing-section-signup .heading {
    margin-top: 0;
    padding-top: 10px;
    font-size: 70px;
    font-weight: normal;
    line-height: 1;
    color: #fff;
    letter-spacing: -1px
}
.marketing-section-signup .subheading {
    margin: 10px 0 0;
    font-size: 21px;
    line-height: 1.5;
    color: #fff
}
.marketing-section-signup .subheading a {
    font-weight: 500
}
.form-signup-home {
    float: right;
    width: 320px;
    margin-left: 40px
}
.form-signup-home .text-muted:last-child {
    margin-bottom: 0
}
.form-signup-home dl.form {
    position: relative;
    margin-top: 0;
    margin-bottom: 10px
}
.form-signup-home dl.form dd+.text-muted {
    margin-top: 5px
}
.form-signup-home .btn,
.form-signup-home dl.form input[type="text"],
.form-signup-home dl.form input[type="password"] {
    padding: 10px;
    font-size: 16px;
    border-radius: 5px
}
.form-signup-home .btn {
    border: 0
}
.form-signup-home .btn:focus {
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05), 0 0 12px rgba(255, 255, 255, 0.75)
}
.form-signup-home dl.form input[type="text"],
.form-signup-home dl.form input[type="password"] {
    width: 100%;
    margin-right: 0;
    border-color: #fff
}
.form-signup-home dl.form input[type="text"]:focus,
.form-signup-home dl.form .focused .drag-and-drop,
.focused .form-signup-home dl.form .drag-and-drop,
.form-signup-home dl.form input[type="password"]:focus {
    background-color: #fff;
    border-color: #fff;
    box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.075), 0 0 12px rgba(255, 255, 255, 0.75)
}
.form-signup-home dl.form.errored dd.error,
.form-signup-home dl.form.errored dd.warning {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    z-index: 5;
    margin-top: 2px;
    font-weight: normal;
    padding: 10px;
    border: 0;
    text-align: left;
    border-radius: 3px;
    color: #fff;
    background-color: #bf1515;
    font-size: 13px;
    text-shadow: none
}
.form-signup-home dl.form.errored dd.error:after,
.form-signup-home dl.form.errored dd.warning:after {
    position: absolute;
    bottom: 100%;
    z-index: 15;
    left: 10px;
    border: 5px solid rgba(191, 21, 21, 0);
    content: " ";
    height: 0;
    width: 0;
    pointer-events: none;
    border-bottom-color: #bf1515
}
.form-signup-home dl.form.errored dd.error:empty,
.form-signup-home dl.form.errored dd.warning:empty {
    display: none
}
.form-signup-home dl.form dd input.is-autocheck-successful,
.form-signup-home dl.form dd input.is-autocheck-errored,
.form-signup-home dl.form dd input.is-autocheck-loading {
    background-image: none
}
.form-signup-home dl.successed:after,
.form-signup-home dl.errored:after {
    position: absolute;
    top: 15px;
    right: 10px;
    text-shadow: none;
    font: normal normal 16px/1 "octicons";
    display: inline-block;
    text-decoration: none;
    -webkit-font-smoothing: antialiased
}
.form-signup-home dl.successed:after {
    content: "\f03a";
    color: #6cc644
}
.form-signup-home dl.errored:after {
    content: "\f02d";
    color: #bd2c00
}
.form-signup-home dl.is-loading:after {
    position: absolute;
    top: 15px;
    right: 10px;
    display: block;
    width: 16px;
    height: 16px;
    content: "";
    background-image: url(https://github.com/images/spinners/octocat-spinner-16px.gif)
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    .form-signup-home dl.is-loading:after {
        background-image: url(https://github.com/images/spinners/octocat-spinner-32.gif);
        background-size: 16px 16px
    }
}
.text-muted .notice-highlight {
    color: #fff
}
.text-center {
    text-align: center
}
.marketing-section-img {
    display: block;
    max-width: 980px;
    margin: 40px auto;
    border-radius: 5px;
    border: 1px solid rgba(0, 0, 0, 0.25);
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.15)
}
.marketing-section-enterprise {
    overflow: hidden;
    max-height: 375px;
    padding-bottom: 20px;
    margin-top: -1px;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    background-image: -webkit-linear-gradient(#202d5f, #614381);
    background-image: linear-gradient(#202d5f, #614381);
    border-bottom: 0;
    box-shadow: inset 0 10px 20px rgba(0, 0, 0, 0.1);
    color: #fff
}
.marketing-section-enterprise.marketing-inline {
    margin-bottom: 30px
}
.marketing-section-enterprise h1>a {
    color: inherit;
    text-decoration: none
}
.marketing-section-enterprise .marketing-header .lead {
    color: #cecbda
}
.btn-marketing {
    padding: 9px 15px;
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 18px;
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.25);
    background-color: #1d6ac8;
    background-image: -webkit-linear-gradient(#45b3f3, #1d6ac8);
    background-image: linear-gradient(#45b3f3, #1d6ac8);
    border: 0;
    box-shadow: 0 1px 1px rgba(0, 0, 0, 0.25)
}
.btn-marketing:hover {
    color: #fff;
    background-image: -webkit-linear-gradient(#2da9f1, #1a5eb2);
    background-image: linear-gradient(#2da9f1, #1a5eb2)
}
.btn-marketing:active {
    text-shadow: 0 1px 0 rgba(0, 0, 0, 0.15);
    background-color: #1a5eb2;
    background-image: none
}
@-webkit-keyframes cloud1animation {
    0% {
        -webkit-transform: translateX(-50px);
        transform: translateX(-50px)
    }
    100% {
        -webkit-transform: translateX(0);
        transform: translateX(0)
    }
}
@keyframes cloud1animation {
    0% {
        -webkit-transform: translateX(-50px);
        transform: translateX(-50px)
    }
    100% {
        -webkit-transform: translateX(0);
        transform: translateX(0)
    }
}
@-webkit-keyframes cloud2animation {
    0% {
        -webkit-transform: translateX(50px);
        transform: translateX(50px)
    }
    100% {
        -webkit-transform: translateX(0);
        transform: translateX(0)
    }
}
@keyframes cloud2animation {
    0% {
        -webkit-transform: translateX(50px);
        transform: translateX(50px)
    }
    100% {
        -webkit-transform: translateX(0);
        transform: translateX(0)
    }
}
.issue-list em {
    font-weight: bold;
    background-color: rgba(255, 255, 140, 0.5);
    padding: 3px;
    border-radius: 3px;
    font-style: normal
}
.issue-list .title {
    padding: 0;
    min-height: 24px;
    font-weight: normal;
    font-size: 18px;
    line-height: 24px;
    margin: 0 80px 10px 0;
    word-wrap: break-word
}
.issue-list .title .mega-octicon {
    position: absolute;
    top: -4px;
    left: 0;
    width: 32px;
    color: #888;
    text-align: center
}
.issue-list .title .closed.mega-octicon {
    color: #bd2c00
}
.issue-list .title .open.mega-octicon {
    color: #6cc644
}
.issue-list .title .merged.mega-octicon {
    color: #6e5494
}
.issue-list .description {
    margin: 0 0 10px;
    line-height: 20px;
    overflow: hidden
}
.issue-list-meta {
    margin: 0;
    list-style-type: none;
    font-size: 11px;
    color: #999
}
.issue-list-meta:before {
    display: table;
    content: ""
}
.issue-list-meta:after {
    display: table;
    clear: both;
    content: ""
}
.issue-list-meta>li {
    display: inline-block;
    margin-right: 10px
}
.issue-list-meta a {
    color: #333
}
.issue-list-meta .octicon {
    color: #838383;
    vertical-align: middle
}
.issue-list-item {
    border-bottom: 1px solid #f1f1f1;
    padding: 0 0 20px 40px;
    margin: 0 0 20px;
    position: relative
}
.labels-list .blankslate {
    display: none
}
.labels-list .table-list-header {
    display: block
}
.labels-list.is-empty .blankslate {
    display: block
}
.labels-list.is-empty .table-list-header {
    display: none
}
.labels-list-item .table-list-cell {
    width: 100%
}
.labels-list-item .label {
    display: inline-block;
    height: 34px;
    padding: 0 10px;
    margin-right: 5px;
    font-size: 16px;
    font-weight: bold;
    line-height: 34px;
    text-align: center;
    border-radius: 3px;
    -webkit-transition: all 0.2s linear;
    transition: all 0.2s linear
}
.labels-list-item .label .octicon {
    margin-right: 3px
}
.labels-list-item .label:hover {
    opacity: 0.85
}
.labels-list-item.open .label,
.labels-list-item.open .label-description,
.labels-list-item.open .labels-list-action {
    display: none
}
.labels-list-item.open .label-delete {
    display: block;
    text-align: left
}
.labels-list-item.edit .label,
.labels-list-item.edit .label-description,
.labels-list-item.edit .labels-list-action {
    display: none
}
.labels-list-item.edit .label-edit {
    display: block
}
.label-description {
    padding: 8px 10px;
    color: #767676
}
.label-delete-confirmation {
    line-height: 34px
}
.labels-list-actions {
    margin-left: 60px
}
.labels-list-action {
    float: left;
    display: block;
    color: #767676;
    padding: 8px 10px
}
.labels-list-action .octicon {
    margin-right: 2px
}
.labels-list-action .octicon-pencil {
    font-size: 14px
}
.labels-list-action:hover {
    color: #4078c0;
    cursor: pointer
}
.new-label {
    display: none;
    padding: 10px;
    margin-bottom: 15px;
    background-color: #fafafa;
    border: 1px solid #e5e5e5;
    border-radius: 3px
}
.new-label .label-edit {
    display: block
}
.new-label .label-edit:before {
    display: table;
    content: ""
}
.new-label .label-edit:after {
    display: table;
    clear: both;
    content: ""
}
.new-label-actions {
    float: right
}
.open .new-label {
    display: block
}
.label-spinner {
    display: none;
    float: left;
    margin-left: -35px;
    margin-top: 9px
}
.label-edit:before {
    display: table;
    content: ""
}
.label-edit:after {
    display: table;
    clear: both;
    content: ""
}
.label-edit label {
    display: block;
    margin-bottom: 5px
}
.label-edit .error {
    float: left;
    margin-top: 8px;
    margin-left: 10px;
    color: #f00
}
.label-edit.is-valid .color-editor .octicon-check {
    display: block
}
.label-edit.loading .label-spinner {
    display: block
}
.color-editor {
    position: relative;
    float: left;
    width: 100px
}
.color-editor.open .label-colors {
    display: block
}
.color-editor-bg {
    position: absolute;
    left: 0;
    z-index: 10;
    width: 20px;
    height: 20px;
    margin-top: 7px;
    margin-left: 7px;
    cursor: pointer;
    border-radius: 3px
}
input.color-editor-input {
    width: 100px;
    padding-left: 34px;
    border-color: #ccc !important
}
input.color-editor-input:focus {
    border-color: #51a7e8 !important
}
input.color-editor-input:focus ~ .label-colors {
    display: block
}
.invalid-color-indicator {
    display: none;
    position: absolute;
    top: 7px;
    left: 7px;
    z-index: 11;
    width: 20px;
    height: 20px;
    line-height: 20px;
    font-weight: bold;
    color: #fff;
    text-align: center
}
.label-edit-name {
    width: 40%;
    float: left;
    margin-right: 10px
}
.label-colors {
    float: left;
    display: none;
    width: auto;
    padding-left: 5px;
    padding-right: 5px
}
.label-edit,
.label-delete {
    display: none
}
.label-delete-form {
    display: inline
}
.label-delete-form.loading .label-delete-spinner {
    display: block
}
.label-delete-spinner {
    display: none;
    margin-top: 10px;
    margin-right: 10px;
    float: left
}
.color-chooser {
    display: table-row;
    height: 25px;
    list-style: none
}
.color-chooser li {
    display: table-cell;
    width: 1%
}
.color-chooser li:hover {
    position: relative;
    z-index: 2;
    outline: 2px solid #fff;
    box-shadow: 0 0 5px 2px rgba(0, 0, 0, 0.25)
}
.color-chooser .color-cooser-color {
    display: block;
    width: 25px;
    height: 25px;
    text-align: center;
    cursor: pointer
}
.repository-lang-stats {
    position: relative
}
.repository-lang-stats ol.repository-lang-stats-numbers li {
    display: table-cell;
    width: 1%;
    padding: 11px 5px;
    text-align: center;
    white-space: nowrap;
    border-bottom: 0
}
.repository-lang-stats ol.repository-lang-stats-numbers li span.percent {
    float: none
}
.repository-lang-stats ol.repository-lang-stats-numbers li>a,
.repository-lang-stats ol.repository-lang-stats-numbers li>span {
    font-weight: bold;
    color: #999;
    text-decoration: none
}
.repository-lang-stats ol.repository-lang-stats-numbers li .lang {
    color: #333
}
.repository-lang-stats ol.repository-lang-stats-numbers li .language-color {
    display: inline-block;
    width: 10px;
    height: 10px;
    border-radius: 50%
}
.repository-lang-stats ol.repository-lang-stats-numbers li a:hover {
    background: transparent
}
.stats-switcher-viewport {
    height: 38px;
    overflow: hidden
}
.stats-switcher-viewport .stats-switcher-wrapper {
    position: relative;
    top: 0;
    -webkit-transition: top 0.25s ease-in-out;
    transition: top 0.25s ease-in-out
}
.stats-switcher-viewport.is-revealing-lang-stats .stats-switcher-wrapper {
    top: -38px
}
.repository-lang-stats-graph {
    display: table;
    width: 100%;
    overflow: hidden;
    white-space: nowrap;
    cursor: pointer;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    border: 1px solid #ddd;
    border-top: 0;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.repository-lang-stats-graph .language-color {
    display: table-cell;
    line-height: 8px;
    text-indent: -9999px
}
.list-group-item {
    position: relative;
    display: block;
    margin-bottom: -1px;
    padding: 8px 10px 10px 40px;
    border: 1px solid #e5e5e5
}
.list-group-item a:hover {
    text-decoration: none
}
.list-group-item:first-child {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.list-group-item:last-child {
    margin-bottom: 0;
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px
}
.list-group-item.closed {
    background-color: #fcfcfc
}
.list-group-item.selectable {
    padding-left: 60px
}
.list-group-item.selected {
    background-color: #ffffef
}
.list-group-item.navigation-focus {
    background-color: #f5f9fc
}
.list-group-item .list-group-item-summary a {
    color: #767676
}
.list-group-item .list-group-item-summary a.quiet {
    color: #999
}
.list-group-item .status {
    position: relative;
    top: 2px;
    margin-right: -9px;
    float: right
}
.list-group-item .type-icon {
    vertical-align: middle;
    position: relative;
    top: 1px;
    width: 16px;
    text-align: center
}
.list-group-item .assignee {
    float: right
}
.list-group-item .assignee img {
    display: block;
    border-radius: 2px
}
.list-group-item .labels {
    display: inline-block;
    top: -2px;
    margin-bottom: -2px;
    margin-left: 4px
}
.list-group-item-name {
    margin: 0 60px 2px 0;
    font-size: 15px;
    line-height: 1.3;
    word-wrap: break-word
}
.list-group-item-name .type-icon {
    float: left;
    margin-top: 1px;
    margin-left: -24px
}
.list-group-item-link {
    color: #333
}
.closed.octicon,
.reverted.octicon {
    color: #bd2c00
}
.open.octicon {
    color: #6cc644
}
.merged.octicon {
    color: #6e5494
}
.list-group-item-summary {
    margin-top: 2px
}
.list-group-item-summary p {
    margin: 0 0 5px
}
.standalone .list-group-item-summary p {
    margin-bottom: 0
}
.animated-ellipsis-container {
    display: inline-block;
    overflow: hidden;
    height: 12px;
    width: 12px;
    -webkit-transform: translateZ(0);
    transform: translateZ(0)
}
.animated-ellipsis-container>.animated-ellipsis {
    overflow: hidden;
    display: inline-block;
    vertical-align: bottom
}
@-webkit-keyframes ellipsis {
    from {
        width: 2px
    }
    to {
        width: 12px
    }
}
@keyframes ellipsis {
    from {
        width: 2px
    }
    to {
        width: 12px
    }
}
.large-loading-area {
    text-align: center;
    padding: 100px 0
}
.context-loader.large-format-loader {
    position: fixed;
    display: none;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    margin: 0;
    width: auto;
    padding: 190px 0 0;
    font-weight: normal;
    background: rgba(255, 255, 255, 0.8);
    border: 0;
    z-index: 9999;
    text-align: center;
    color: #767676
}
.context-loader.large-format-loader.is-loading {
    display: block
}
.request-reviewers {
    padding: 0 10px 10px;
    border-top: 1px solid #ddd
}
.request-reviewers h3 {
    margin-bottom: 5px
}
.request-reviewers .label {
    vertical-align: middle
}
.request-reviewers .input-block {
    display: inline-block;
    max-width: 348px;
    margin-right: 10px
}
.request-reviewers-heading .label {
    margin-top: -2px;
    margin-left: 3px
}
.request-reviewers-autocomplete .autocomplete-results {
    border: 0
}
.request-reviewers-autocomplete .autocomplete-results .typeahead-result {
    display: block;
    width: 300px
}
.request-reviewers-autocomplete .octicon-jersey {
    float: left;
    width: 24px;
    height: 24px;
    margin-right: 10px;
    margin-left: -34px;
    line-height: 24px;
    color: #767676;
    text-align: center
}
.request-reviewers-list {
    margin-top: 10px
}
.request-reviewers-list:before {
    display: table;
    content: ""
}
.request-reviewers-list:after {
    display: table;
    clear: both;
    content: ""
}
.requested-reviewer-item {
    padding-bottom: 10px
}
.review-bar .octicon-check {
    float: left;
    margin-right: 10px;
    color: #6cc644
}
.review-bar .complete-review-actions .review-bar-form {
    float: right
}
.review-bar .complete-review-actions p {
    margin-top: 6px
}
.checklist {
    border-top: 1px solid #eee;
    padding: 15px 10px
}
.checklist.is-empty .checklist-items,
.checklist.is-empty .checklist-form {
    display: none
}
.checklist.is-empty .checklist-callout {
    display: block
}
.checklist .checklist-items,
.checklist .checklist-form {
    display: block
}
.checklist .checklist-callout {
    display: none
}
.comment .checklist {
    padding: 15px;
    border-radius: 0 0 3px 3px;
    background-color: #fcfcfc
}
.checklist-form {
    clear: both;
    display: none;
    position: relative
}
.checklist-form .suggester-container {
    top: 4px
}
.checklist-form dl.form {
    margin: 0
}
dl.form>dd input.checklist-input {
    width: 576px;
    margin-right: 0
}
dl.form>dd input.checklist-input[disabled] {
    color: #999
}
.checklist-callout .octicon-checklist {
    vertical-align: bottom
}
.checklist-callout .btn {
    top: -4px
}
.checklist-callout-copy {
    margin: 0;
    font-size: 12px;
    color: #666
}
.checklist-items {
    display: none
}
.checklist-item {
    margin-bottom: 10px;
    list-style-type: none;
    color: #767676;
    -webkit-transition: opacity 0.2s ease 0s;
    transition: opacity 0.2s ease 0s
}
.checklist-item.is-transitioning {
    opacity: 0.5
}
.checklist-item.is-closed .checklist-item-status-closed,
.checklist-item.is-reopened .checklist-item-status-reopened {
    display: inline;
    opacity: 1
}
.checklist-item-status {
    position: absolute;
    margin-left: 15px;
    opacity: 0;
    -webkit-transition: opacity 0.35s ease 0s;
    transition: opacity 0.35s ease 0s
}
.checklist-item-checkbox-form {
    display: inline
}
.checklist-progress {
    position: relative;
    top: 1px;
    background-color: #6cc644;
    height: 1px;
    -webkit-transition: width 0.35s ease 0s;
    transition: width 0.35s ease 0s
}
#quick-issue-modal {
    display: none
}
.quick-issue-modal-footer {
    margin-bottom: 0
}
.quick-issue-thanks {
    display: none;
    font-size: 18px
}
.quick-issue-link {
    margin-left: 30px
}
.quick-issue-body {
    display: block;
    width: 100%
}
.quick-issue-form {
    position: relative
}
.quick-issue-form .suggestions {
    margin-left: 0;
    margin-bottom: 0
}
.quick-issue-form .drag-and-drop {
    font-size: 10px
}
.clearfix:before {
    display: table;
    content: ""
}
.clearfix:after {
    display: table;
    clear: both;
    content: ""
}
.right {
    float: right
}
.left {
    float: left
}
.centered {
    display: block;
    float: none;
    margin-left: auto;
    margin-right: auto
}
.text-right {
    text-align: right
}
.text-left {
    text-align: left
}
.danger {
    color: #c00
}
.mute {
    color: #000
}
.text-diff-added {
    color: #55a532
}
.text-diff-deleted {
    color: #bd2c00
}
.text-open,
.text-success {
    color: #6cc644
}
.text-closed {
    color: #bd2c00
}
.text-reverted {
    color: #bd2c00
}
.text-merged {
    color: #6e5494
}
.text-renamed {
    color: #fffa5d
}
.text-pending {
    color: #cea61b
}
.text-error,
.text-failure {
    color: #bd2c00
}
.muted-link {
    color: #767676
}
.muted-link:hover {
    color: #4078c0;
    text-decoration: none
}
.hidden {
    display: none
}
.warning {
    padding: 0.5em;
    margin-bottom: 0.8em;
    font-weight: bold;
    background-color: #fffccc
}
.error_box {
    padding: 1em;
    font-weight: bold;
    background-color: #ffebe8;
    border: 1px solid #dd3c10
}
`)

var Github2CSS = []byte(`
.marketing .pagehead h1 {
    font-size: 30px
}
.marketing .pagehead p {
    margin-top: 4px;
    margin-bottom: 0;
    font-size: 14px;
    color: #767676
}
.marketing .pagehead ul.actions {
    margin-top: 10px
}
.marketing h2 .secure {
    float: right;
    padding: 1px 0;
    font-size: 11px;
    font-weight: bold;
    text-transform: uppercase;
    color: #6cc644
}
.marketing .questions p {
    font-size: 14px
}
.marketing-header {
    margin-bottom: 40px
}
.marketing-header h1 {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 42px;
    font-weight: 300
}
.marketing-header .lead {
    color: #767676;
    max-width: 750px;
    margin: 10px auto 0
}
.marketing-header .btn {
    margin-top: 15px;
    padding: 12px 20px;
    font-size: 18px;
    font-weight: normal;
    border-radius: 6px
}
.marketing-section {
    position: relative;
    padding-top: 80px;
    padding-bottom: 80px;
    border-bottom: 1px solid #e5e5e5;
    text-align: center;
    font-size: 16px;
    line-height: 1.5
}
.marketing-section:before {
    display: table;
    content: ""
}
.marketing-section:after {
    display: table;
    clear: both;
    content: ""
}
.marketing-section h3 {
    font-size: 21px;
    font-weight: normal
}
.marketing-section-stripe {
    background-color: #f5f5f5
}
.marketing-hero-octicon {
    width: 100px;
    height: 100px;
    border-radius: 50px;
    text-align: center;
    border: solid 1px #e5e5e5;
    margin: 0 auto 15px
}
.marketing-hero-octicon .mega-octicon {
    color: #4078c0;
    font-size: 48px;
    line-height: 100px
}
.marketing-hero-octicon .octicon-checklist {
    position: relative;
    right: -3px
}
.marketing-grid {
    font-size: 14px
}
.marketing-grid .column {
    padding: 20px 25px 40px
}
.marketing-grid p {
    margin: 0 auto;
    max-width: 90%;
    color: #5a5a5a
}
.marketing-grid .mega-octicon {
    color: #4078c0
}
.read-it {
    padding-top: 50px;
    text-align: center;
    border-top: 1px solid #eee
}
.contact-form .input-block {
    margin-top: 10px;
    margin-bottom: 10px
}
.contact-form textarea {
    height: 200px;
    resize: vertical
}
.hanging-list li,
.hanging-icon-list li {
    margin: 10px 0;
    font-size: 14px
}
.hanging-list li {
    margin-left: 12px;
    list-style-position: inside
}
.hanging-icon-list li {
    padding-left: 25px;
    list-style-type: none
}
.hanging-icon-list .octicon {
    float: left;
    margin-left: -20px;
    color: #767676
}
.hanging-icon-list .octicon-check {
    color: #6cc644
}
.hanging-icon-list .octicon-x {
    color: #bd2c00
}
.logos-page h3 {
    font-size: 18px
}
.logos-download {
    position: relative;
    display: block;
    float: left;
    width: 32%;
    height: 290px;
    margin-bottom: 30px;
    padding-top: 20px;
    text-align: center;
    border: 1px solid #ddd;
    border-radius: 6px
}
.logos-download+.logos-download {
    margin-left: 2%
}
.logos-download .gh-logo {
    margin-top: 70px
}
.logos-download .gh-octocat {
    margin-top: 10px
}
.logos-download-link {
    position: absolute;
    right: 0;
    bottom: 0;
    left: 0;
    display: block;
    padding: 15px 20px;
    font-size: 16px;
    font-weight: bold;
    background-color: #f5f5f5;
    border-top: 1px solid #ddd;
    border-radius: 0 0 5px 5px
}
.logos-download-link .octicon {
    vertical-align: 2px
}
.logos-download:hover {
    text-decoration: none
}
.logos-download:hover .logos-download-link {
    background-color: #eee
}
.nonprofit-head {
    padding: 100px 0 120px;
    border-bottom: 1px solid #eee;
    text-align: center;
    position: relative;
    overflow: hidden
}
.nonprofit-head .title {
    font-weight: 300;
    font-size: 30px;
    color: #767676;
    margin-bottom: 20px;
    display: inline-block;
    border-bottom: 1px solid #ccc
}
.nonprofit-head .title .mega-octicon {
    color: #333
}
.nonprofit-head .logo {
    vertical-align: middle
}
.nonprofit h1 {
    font-weight: 300;
    font-size: 28px;
    line-height: 1.5em;
    position: relative
}
.nonprofit h2 {
    font-weight: normal
}
.heart {
    width: 12px;
    height: 12px;
    background: #83d6c0;
    box-shadow: 140px 30px 0 #efa, 120px -120px 0 #aded84, 220px -60px 0 #ded, 30px 240px 0 #ada, 60px -60px 0 #d76666, 60px -30px 0 #ff846f, 60px 0 0 #f9a7a7, 60px 30px 0 #ffc8c8, 60px 60px 0 #ffd8d8, 30px 60px 0 #baf2ca, 30px 30px 0 #98eaac, 30px 0 0 #80d896, 30px -30px 0 #6dd085, 30px -60px 0 #55be6f, 0 -60px 0 #4cc2a7, 0 -30px 0 #73d3b9, 0 30px 0 #93e3cd, 0 60px 0 #adf9e4, -30px 60px 0 #ffe1b9, -30px 30px 0 #ffd194, -30px 0 0 #ffc86f, -60px 0 0 #fd9ff0, -60px 30px 0 #ffbaf7, -60px 60px 0 #fccdf7, -180px 60px 0 #9df;
    position: absolute;
    left: 50%;
    top: 40%;
    margin-left: 400px;
    -webkit-transform: rotate(45deg);
    -ms-transform: rotate(45deg);
    transform: rotate(45deg)
}
.heart.left {
    margin-left: -400px;
    -webkit-transform: rotate(-45deg), scaleX(-1);
    -ms-transform: rotate(-45deg), scaleX(-1);
    transform: rotate(-45deg), scaleX(-1)
}
.octo-earth {
    position: absolute;
    left: 50%;
    bottom: -150px;
    margin-left: -120px;
    -webkit-animation: rotate 20s infinite linear;
    animation: rotate 20s infinite linear
}
@-webkit-keyframes rotate {
    0% {
        -webkit-transform: rotate(0);
        transform: rotate(0)
    }
    100% {
        -webkit-transform: rotate(-360deg);
        transform: rotate(-360deg)
    }
}
@keyframes rotate {
    0% {
        -webkit-transform: rotate(0);
        transform: rotate(0)
    }
    100% {
        -webkit-transform: rotate(-360deg);
        transform: rotate(-360deg)
    }
}
.nonprofit-steps {
    margin-left: 30px;
    font-size: 20px;
    font-weight: 300
}
.nonprofit-steps li {
    margin-bottom: 10px
}
.nonprofit-steps ul {
    margin: 15px 0 0 20px;
    list-style: square
}
.nonprofit-section {
    padding: 50px 0;
    background: #f5f5f5
}
.nonprofit-section h1 {
    text-align: center
}
.nonprofit-section .dialog {
    width: 640px;
    background: #fff;
    margin: 30px auto 0;
    padding: 30px
}
.nonprofit-section .dialog h2:first-child {
    margin-top: 0
}
.nonprofit-section .dialog p:last-child {
    margin-bottom: 0
}
.dialog.edu-callout {
    border: 5px solid #aec;
    padding: 25px
}
.dialog.edu-callout p {
    margin-top: 0
}
.dialog.edu-callout .mega-octicon {
    float: left;
    padding-top: 10px;
    padding-bottom: 10px;
    margin-right: 15px;
    color: #418f65
}
.add-on-table {
    border-collapse: separate;
    border: 1px solid #e0e0e0;
    border-radius: 3px;
    margin: 20px 0 40px
}
.add-on-table td {
    padding: 10px;
    border-bottom: 1px solid #e0e0e0
}
.add-on-table tr:last-child td,
.add-on-table td[rowspan] {
    border-bottom: 0
}
.add-on-table .add-on-name {
    border-right: 1px solid #e0e0e0;
    border-bottom: 0;
    padding-left: 30px;
    padding-right: 30px;
    text-align: center;
    width: 1%;
    white-space: nowrap
}
.add-on-table .add-on-name .btn:not(:last-child) {
    margin-right: 8px
}
.add-on-plan {
    float: left;
    text-align: center;
    background: #f0f5fa;
    min-width: 160px;
    border-radius: 3px;
    margin-right: 15px;
    padding: 10px;
    font-weight: bold
}
.lfs-plan-price {
    margin-top: 8px;
    color: #7a7a7a
}
.lfs-details {
    padding: 7px 0 6px;
    list-style: none
}
.lfs-details:first-child {
    border-bottom: 1px solid #eee
}
.lfs-data {
    width: 120px;
    color: #767676;
    float: right
}
.add-on-logo {
    display: block;
    width: 240px;
    margin: 10px auto
}
.integrations {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif
}
.integrations .site-footer {
    border-top: 1px solid #e5e5e5;
    margin-top: -99px
}
.integrations .lead {
    color: #c1c2c2;
    font-size: 18px;
    font-weight: 300
}
.integrations-intro {
    background: #ecf9f8;
    background-image: -webkit-linear-gradient(45deg, #ecf9f8 0%, #fbf4fc 100%);
    background-image: linear-gradient(45deg, #ecf9f8 0%, #fbf4fc 100%);
    padding: 60px 0
}
.integrations-intro .marketing-header {
    margin-bottom: 0
}
.integrations-intro h1 {
    font-size: 48px;
    font-weight: 200;
    margin-bottom: 20px
}
.integrations-intro .lead {
    max-width: 540px;
    color: rgba(87, 87, 87, 0.7)
}
.integrations-intro .marketing-graphic {
    margin: 52px 0 0;
    position: relative
}
.integrations-intro .in-page-link {
    content: " ";
    display: block;
    position: absolute;
    top: 0;
    bottom: 0;
    height: 260px
}
.integrations-intro .in-page-link.build {
    width: 240px;
    left: 98px
}
.integrations-intro .in-page-link.collaborate {
    width: 181px;
    left: 338px
}
.integrations-intro .in-page-link.deploy {
    width: 331px;
    right: 0
}
.integrations-list .lead {
    color: #767676
}
.integrations-list .lead.narrow {
    max-width: 520px
}
.integrations-list .marketing-hero-octicon {
    width: 75px;
    height: 75px;
    margin: -30px auto 30px;
    border-width: 5px
}
.integrations-list .marketing-hero-octicon .mega-octicon {
    font-size: 34px;
    line-height: 66px
}
.integrations-list .octicon-blue {
    border-color: #229eba
}
.integrations-list .octicon-blue .mega-octicon {
    color: #229eba
}
.integrations-list .octicon-turquoise {
    border-color: #75bbb6
}
.integrations-list .octicon-turquoise .mega-octicon {
    color: #75bbb6
}
.integrations-list .octicon-purple {
    border-color: #976999
}
.integrations-list .octicon-purple .mega-octicon {
    color: #976999
}
.integrations-outro {
    background: #ecf9f8;
    background-image: -webkit-linear-gradient(45deg, #ecf9f8 0%, #fbf4fc 100%);
    background-image: linear-gradient(45deg, #ecf9f8 0%, #fbf4fc 100%);
    padding-top: 118px;
    padding-bottom: 218px
}
.integrations-outro .lead {
    width: 550px;
    color: rgba(87, 87, 87, 0.7)
}
.btn.outro-button {
    font-size: 20px;
    font-weight: 200;
    background: #75bbb6;
    border: 0;
    padding: 14px 65px;
    color: #fff;
    text-decoration: none;
    text-shadow: none;
    -webkit-transition: opacity ease 0.1s;
    transition: opacity ease 0.1s
}
.btn.outro-button:hover {
    opacity: 0.8
}
.integrations-logo-container {
    min-height: 80px;
    margin-bottom: 20px
}
.features-next .lead strong {
    color: #444
}
.features-next .native-mobile-screens {
    list-style-type: none;
    border-bottom: solid 1px #d9d9d9;
    line-height: 0;
    padding-left: 0
}
.features-next .native-mobile-screens li {
    display: inline;
    margin: 0 5px 30px
}
.team-org-chart {
    margin: 30px auto;
    width: 470px
}
.team-org-chart .mega-octicon {
    vertical-align: middle
}
.team-org-group {
    border: solid 1px #ccc;
    background-color: #fff;
    text-align: center;
    font-size: 16px;
    padding: 10px;
    margin-bottom: 13px;
    border-radius: 3px
}
.team-org-group strong {
    color: #333
}
.team-org-team {
    width: 147px;
    height: 120px;
    display: inline-block;
    vertical-align: top
}
.team-org-team+.team-org-team {
    margin-left: 10px
}
.octicon-jersey-red {
    color: #bd2c00
}
.octicon-jersey-green {
    color: #6cc644
}
.octicon-jersey-orange {
    color: #c9510c
}
.team-org-members {
    margin-top: 15px
}
.team-org-members .octicon {
    color: #aaa
}
.team-org-repos .mega-octicon {
    color: #bbb;
    margin: 0 5px
}
.team-animation {
    -webkit-animation-duration: 12s;
    animation-duration: 12s;
    -webkit-animation-iteration-count: infinite;
    animation-iteration-count: infinite
}
.team-design {
    -webkit-animation-name: teamDesign;
    animation-name: teamDesign
}
.team-dev {
    -webkit-animation-name: teamDev;
    animation-name: teamDev
}
.team-marketing {
    -webkit-animation-name: teamMarketing;
    animation-name: teamMarketing
}
.team-dev-design {
    -webkit-animation-name: teamDevDesign;
    animation-name: teamDevDesign
}
.team-dev-design-marketing {
    -webkit-animation-name: teamDevDesignMarketing;
    animation-name: teamDevDesignMarketing
}
.features-section img {
    max-width: 100%
}
.features-section code {
    font-size: 0.9em;
    padding: 3px 5px;
    border-radius: 2px;
    background-color: #e7e7e7
}
.features-section p {
    max-width: 750px;
    margin-right: auto;
    margin-left: auto
}
.features-repo-count {
    white-space: nowrap
}
.features-content-right {
    float: right;
    width: 470px;
    text-align: left
}
.features-content-left {
    float: left;
    width: 470px;
    text-align: left
}
.diagram-icon {
    position: absolute;
    border-radius: 50px;
    border: solid 4px #4078c0;
    background-color: #fff;
    width: 53px;
    height: 53px;
    text-align: center;
    line-height: 55px;
    color: #4078c0
}
.diagram-icon-small {
    position: absolute;
    color: #4078c0;
    margin-top: 2px\9
}
.diagram-icon-branch {
    top: -13px;
    left: 81px;
    -webkit-animation: bounceIn 0.6s ease-in-out 0.25s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 0.25s 1 normal both
}
.diagram-icon-pr {
    top: 89px;
    left: 405px;
    -webkit-animation: bounceIn 0.6s ease-in-out 1.8s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 1.8s 1 normal both
}
.diagram-icon-merge {
    top: -13px;
    left: 843px;
    -webkit-animation: bounceIn 0.6s ease-in-out 3.7s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 3.7s 1 normal both
}
.diagram-icon-commit-1 {
    top: 101px;
    left: 240px;
    -webkit-animation: bounceIn 0.6s ease-in-out 1.3s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 1.3s 1 normal both;
    background-color: #fff
}
.diagram-icon-commit-2 {
    top: 101px;
    left: 295px;
    -webkit-animation: bounceIn 0.6s ease-in-out 1.4s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 1.4s 1 normal both;
    background-color: #fff
}
.diagram-icon-commit-3 {
    top: 101px;
    left: 350px;
    -webkit-animation: bounceIn 0.6s ease-in-out 1.5s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 1.5s 1 normal both;
    background-color: #fff
}
.diagram-icon-discussion-1 {
    top: 79px;
    left: 488px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2s 1 normal both;
    opacity: 0.3
}
.diagram-icon-commit-4 {
    top: 101px;
    left: 515px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.1s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.1s 1 normal both;
    background-color: #fff
}
.diagram-icon-discussion-2 {
    top: 131px;
    left: 542px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.2s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.2s 1 normal both;
    opacity: 0.3
}
.diagram-icon-commit-5 {
    top: 101px;
    left: 570px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.3s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.3s 1 normal both;
    background-color: #fff
}
.diagram-icon-discussion-3 {
    top: 79px;
    left: 597px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.4s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.4s 1 normal both;
    opacity: 0.3
}
.diagram-icon-commit-6 {
    top: 101px;
    left: 625px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.5s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.5s 1 normal both;
    background-color: #fff
}
.diagram-icon-discussion-4 {
    top: 131px;
    left: 652px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.6s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.6s 1 normal both;
    opacity: 0.3
}
.diagram-icon-commit-7 {
    top: 101px;
    left: 680px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.7s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.7s 1 normal both;
    background-color: #fff
}
.diagram-icon-discussion-5 {
    top: 79px;
    left: 707px;
    -webkit-animation: bounceIn 0.6s ease-in-out 2.8s 1 normal both;
    animation: bounceIn 0.6s ease-in-out 2.8s 1 normal both;
    opacity: 0.3
}
.features-branch-diagram {
    position: relative;
    margin-top: 40px;
    margin-bottom: 50px
}
.preload .diagram-animation {
    -webkit-animation: none !important;
    animation: none !important;
    opacity: 0
}
.mobile .diagram-animation {
    -webkit-animation: none !important;
    animation: none !important;
    opacity: 1
}
.features-highlight {
    margin: 8px 0;
    display: inline-block;
    background-color: #e7e7e7;
    padding: 10px;
    border-radius: 3px
}
.features-highlight i {
    font-style: normal;
    color: #4078c0
}
.features-callout {
    border: solid 1px #eee;
    border-radius: 3px;
    padding: 10px;
    margin-top: 15px;
    margin-bottom: 45px;
    display: inline-block;
    color: #767676;
    font-size: 14px;
    line-height: 1.4;
    text-align: left;
    width: 450px
}
.features-callout>p {
    margin-top: 0;
    margin-bottom: 0
}
.features-callout .left {
    margin-right: 10px;
    vertical-align: center
}
.features-copy-minor {
    font-size: 12px;
    color: #555
}
.svn-callout {
    clear: both;
    padding-top: 30px;
    padding-left: 217px;
    line-height: 0.8;
    text-align: left
}
.svn-callout-heading {
    font-size: 18px;
    margin-bottom: 0;
    color: #444
}
.svn-callout-logo {
    margin-left: -70px;
    margin-top: -1px;
    float: left
}
@-webkit-keyframes teamDev {
    3% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    27% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    30% {
        border-color: #ccc;
        color: #bbb
    }
}
@keyframes teamDev {
    3% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    27% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    30% {
        border-color: #ccc;
        color: #bbb
    }
}
@-webkit-keyframes teamDesign {
    34% {
        border-color: #ccc;
        color: #bbb
    }
    37% {
        border-color: #6cc644;
        color: #6cc644
    }
    60% {
        border-color: #6cc644;
        color: #6cc644
    }
    63% {
        border-color: #ccc;
        color: #bbb
    }
}
@keyframes teamDesign {
    34% {
        border-color: #ccc;
        color: #bbb
    }
    37% {
        border-color: #6cc644;
        color: #6cc644
    }
    60% {
        border-color: #6cc644;
        color: #6cc644
    }
    63% {
        border-color: #ccc;
        color: #bbb
    }
}
@-webkit-keyframes teamMarketing {
    67% {
        border-color: #ccc;
        color: #bbb
    }
    70% {
        border-color: #c9510c;
        color: #c9510c
    }
    94% {
        border-color: #c9510c;
        color: #c9510c
    }
}
@keyframes teamMarketing {
    67% {
        border-color: #ccc;
        color: #bbb
    }
    70% {
        border-color: #c9510c;
        color: #c9510c
    }
    94% {
        border-color: #c9510c;
        color: #c9510c
    }
}
@-webkit-keyframes teamDevDesign {
    3% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    27% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    30% {
        border-color: #ccc;
        color: #bbb
    }
    34% {
        border-color: #ccc;
        color: #bbb
    }
    37% {
        border-color: #6cc644;
        color: #6cc644
    }
    60% {
        border-color: #6cc644;
        color: #6cc644
    }
    63% {
        border-color: #ccc;
        color: #bbb
    }
}
@keyframes teamDevDesign {
    3% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    27% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    30% {
        border-color: #ccc;
        color: #bbb
    }
    34% {
        border-color: #ccc;
        color: #bbb
    }
    37% {
        border-color: #6cc644;
        color: #6cc644
    }
    60% {
        border-color: #6cc644;
        color: #6cc644
    }
    63% {
        border-color: #ccc;
        color: #bbb
    }
}
@-webkit-keyframes teamDevDesignMarketing {
    3% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    27% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    30% {
        border-color: #ccc;
        color: #bbb
    }
    34% {
        border-color: #ccc;
        color: #bbb
    }
    37% {
        border-color: #6cc644;
        color: #6cc644
    }
    60% {
        border-color: #6cc644;
        color: #6cc644
    }
    63% {
        border-color: #ccc;
        color: #bbb
    }
    67% {
        border-color: #ccc;
        color: #bbb
    }
    70% {
        border-color: #c9510c;
        color: #c9510c
    }
    94% {
        border-color: #c9510c;
        color: #c9510c
    }
}
@keyframes teamDevDesignMarketing {
    3% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    27% {
        border-color: #bd2c00;
        color: #bd2c00
    }
    30% {
        border-color: #ccc;
        color: #bbb
    }
    34% {
        border-color: #ccc;
        color: #bbb
    }
    37% {
        border-color: #6cc644;
        color: #6cc644
    }
    60% {
        border-color: #6cc644;
        color: #6cc644
    }
    63% {
        border-color: #ccc;
        color: #bbb
    }
    67% {
        border-color: #ccc;
        color: #bbb
    }
    70% {
        border-color: #c9510c;
        color: #c9510c
    }
    94% {
        border-color: #c9510c;
        color: #c9510c
    }
}
@-webkit-keyframes bounceIn {
    0% {
        opacity: 0;
        -webkit-transform: scale(0.3);
        transform: scale(0.3)
    }
    50% {
        opacity: 1;
        -webkit-transform: scale(1.05);
        transform: scale(1.05)
    }
    70% {
        -webkit-transform: scale(0.9);
        transform: scale(0.9)
    }
    100% {
        -webkit-transform: scale(1);
        transform: scale(1)
    }
}
@keyframes bounceIn {
    0% {
        opacity: 0;
        -webkit-transform: scale(0.3);
        transform: scale(0.3)
    }
    50% {
        opacity: 1;
        -webkit-transform: scale(1.05);
        transform: scale(1.05)
    }
    70% {
        -webkit-transform: scale(0.9);
        transform: scale(0.9)
    }
    100% {
        -webkit-transform: scale(1);
        transform: scale(1)
    }
}
@-webkit-keyframes fadeIn {
    0% {
        opacity: 0
    }
    100% {
        opacity: 1
    }
}
@keyframes fadeIn {
    0% {
        opacity: 0
    }
    100% {
        opacity: 1
    }
}
.segmented-nav-tab {
    display: none;
    margin-top: 40px
}
.segmented-nav-tab:before {
    display: table;
    content: ""
}
.segmented-nav-tab:after {
    display: table;
    clear: both;
    content: ""
}
.segmented-nav-tab.active {
    display: block
}
.octicon-list {
    list-style: none;
    margin-left: 26px;
    margin-bottom: 28px
}
.octicon-list li {
    margin-bottom: 20px;
    position: relative
}
.octicon-list .octicon {
    position: absolute;
    margin-left: -22px;
    top: 3px;
    color: #4078c0
}
.table-list {
    display: table;
    table-layout: fixed;
    width: 100%;
    color: #999;
    border-bottom: 1px solid #e5e5e5
}
.table-list-bordered .table-list-cell:first-child {
    border-left: 1px solid #eee
}
.table-list-bordered .table-list-cell:last-child {
    border-right: 1px solid #eee
}
.table-list-item {
    position: relative;
    list-style: none;
    display: table-row
}
.table-list-item.unread .table-list-cell:first-child {
    box-shadow: 2px 0 0 #4078c0 inset
}
.table-list-cell {
    position: relative;
    display: table-cell;
    padding: 8px 10px;
    font-size: 12px;
    vertical-align: top;
    border-top: 1px solid #eee
}
.table-list-cell.flush-left {
    padding-left: 0
}
.table-list-cell.flush-right {
    padding-right: 0
}
.table-list-cell-checkbox {
    width: 30px;
    padding-left: 0;
    padding-right: 0;
    text-align: center
}
.select-toggle-check {
    margin-top: 7px
}
.table-list-cell-type {
    padding-top: 10px;
    padding-left: 0;
    padding-right: 0;
    width: 20px;
    text-align: center
}
.table-list-cell-type>a {
    display: inline-block
}
.table-list-cell-type .octicon {
    margin-top: 3px
}
.table-list-cell-type:first-child {
    padding-left: 10px
}
.table-list-cell-avatar {
    padding-left: 0;
    padding-right: 0;
    width: 16px
}
.table-list-header {
    position: relative;
    margin-top: 20px;
    margin-bottom: -1px;
    background-color: #f8f8f8;
    border: 1px solid #e5e5e5;
    border-radius: 3px 3px 0 0
}
.table-list-header:before {
    display: table;
    content: ""
}
.table-list-header:after {
    display: table;
    clear: both;
    content: ""
}
.table-list-header .btn-link {
    position: relative;
    display: inline-block;
    padding-top: 13px;
    padding-bottom: 13px;
    font-weight: normal
}
.table-list-header .table-list-header-actions {
    margin-top: 8px;
    margin-right: 10px
}
.table-list-header .table-list-header-action {
    display: inline-block;
    vertical-align: middle
}
.table-list-heading {
    margin-left: 10px
}
.table-list-header-select-all {
    float: left;
    width: 30px;
    padding: 12px 10px;
    margin-right: 5px;
    margin-left: -1px;
    text-align: center
}
.table-list-header-meta {
    display: inline-block;
    padding-top: 13px;
    padding-bottom: 13px;
    color: #767676
}
.table-list-filters:first-child .table-list-header-toggle:first-child {
    padding-left: 10px
}
.table-list-header-toggle.states .selected {
    font-weight: bold
}
.table-list-header-toggle .btn-link {
    color: #767676
}
.table-list-header-toggle .btn-link .octicon {
    margin-right: 2px
}
.table-list-header-toggle .btn-link:hover {
    color: #222;
    text-decoration: none
}
.table-list-header-toggle .btn-link.selected,
.table-list-header-toggle .btn-link.selected:hover {
    color: #222
}
.table-list-header-toggle .btn-link+.btn-link {
    margin-left: 10px
}
.table-list-header-toggle .btn-link:disabled,
.table-list-header-toggle .btn-link.disabled {
    pointer-events: none;
    opacity: 0.5
}
.table-list-header-toggle .select-menu {
    position: relative
}
.table-list-header-toggle .select-menu-item.selected {
    font-weight: bold
}
.table-list-header-toggle .select-menu-button {
    padding-right: 15px;
    padding-left: 15px
}
.table-list-header-toggle .select-menu-button:hover,
.table-list-header-toggle .select-menu-button.selected,
.table-list-header-toggle .select-menu-button.selected:hover {
    color: #222
}
.table-list-header-toggle .select-menu-modal-holder {
    right: 10px
}
.table-list-header-toggle .select-menu-modal-holder .select-menu-modal {
    margin-top: -1px
}
.table-list-triage {
    display: none
}
.triage-mode .table-list-non-triage,
.triage-mode .table-list-filters {
    display: none
}
.triage-mode .table-list-triage {
    display: block
}
.subhead {
    padding-bottom: 10px;
    margin-top: 10px;
    margin-bottom: 15px;
    border-bottom: 1px solid #e5e5e5
}
.subhead-spacious {
    margin-top: 45px
}
.subhead-heading {
    margin-top: 0;
    margin-bottom: 0;
    font-weight: normal
}
.subhead-description {
    margin-top: 5px;
    margin-bottom: 0;
    font-size: 14px;
    color: #767676
}
.subnav {
    margin-bottom: 20px
}
.subnav:before {
    display: table;
    content: ""
}
.subnav:after {
    display: table;
    clear: both;
    content: ""
}
.subnav>.right {
    margin-left: 10px
}
.subnav-bordered {
    padding-bottom: 20px;
    border-bottom: 1px solid #eee
}
.subnav-flush {
    margin-bottom: 0
}
.subnav-item {
    position: relative;
    float: left;
    padding: 7px 14px;
    font-weight: bold;
    color: #666;
    border: 1px solid #e5e5e5
}
.subnav-item+.subnav-item {
    margin-left: -1px
}
.subnav-item:hover,
.subnav-item:focus {
    text-decoration: none;
    background-color: #f5f5f5
}
.subnav-item.selected,
.subnav-item.selected:hover,
.subnav-item.selected:focus {
    z-index: 2;
    color: #fff;
    background-color: #4078c0;
    border-color: #4078c0
}
.subnav-item:first-child {
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px
}
.subnav-item:last-child {
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px
}
.subnav-search {
    position: relative;
    margin-left: 10px
}
input.subnav-search-input {
    width: 320px;
    padding-left: 30px;
    color: #767676;
    border-color: #d5d5d5
}
input.subnav-search-input-wide {
    width: 500px
}
.subnav-search-icon {
    position: absolute;
    top: 0;
    left: 1px;
    display: block;
    width: 30px;
    height: 34px;
    line-height: 34px;
    color: #ccc;
    text-align: center;
    pointer-events: none
}
.subnav-search-context .btn {
    color: #555;
    border-top-right-radius: 0;
    border-bottom-right-radius: 0
}
.subnav-search-context .btn:hover,
.subnav-search-context .btn:focus,
.subnav-search-context .btn:active,
.subnav-search-context .btn.selected {
    z-index: 2
}
.subnav-search-context+.subnav-search {
    margin-left: -1px
}
.subnav-search-context+.subnav-search .subnav-search-input {
    border-top-left-radius: 0;
    border-bottom-left-radius: 0
}
.subnav-search-context .select-menu-modal-holder {
    z-index: 30
}
.subnav-search-context .select-menu-modal {
    width: 220px
}
.subnav-search-context .select-menu-item-icon {
    color: inherit
}
.subnav-divider-right {
    padding-right: 10px;
    border-right: 1px solid #eee
}
.subnav-spacer-right {
    padding-right: 10px
}
.boxed-group {
    position: relative;
    border-radius: 3px;
    margin-bottom: 30px
}
.boxed-group .counter {
    color: #fff;
    background-color: #babec0
}
.boxed-group.flush .boxed-group-inner {
    padding: 0
}
.boxed-group.condensed .boxed-group-inner {
    font-size: 12px;
    padding: 0
}
.boxed-group>h3,
.boxed-group .heading {
    background-color: #f5f5f5;
    margin: 0;
    border-radius: 3px 3px 0 0;
    border: 1px solid #d8d8d8;
    border-bottom: 0;
    padding: 9px 10px 10px;
    font-size: 14px;
    line-height: 17px;
    display: block
}
.boxed-group>h3 a,
.boxed-group .heading a {
    color: inherit
}
.boxed-group>h3 a.boxed-group-breadcrumb,
.boxed-group .heading a.boxed-group-breadcrumb {
    color: #666;
    font-weight: normal;
    text-decoration: none
}
.boxed-group>h3 .avatar,
.boxed-group .heading .avatar {
    margin-top: -4px
}
.boxed-group .tabnav.heading {
    padding: 0
}
.boxed-group .tabnav.heading .tabnav-tab.selected {
    border-top: 0
}
.boxed-group .tabnav.heading li:first-child .selected {
    border-left-color: #fff;
    border-top-left-radius: 3px
}
.boxed-group .tabnav-tab {
    border-radius: 0;
    border-top: 0
}
.boxed-group code.heading {
    font-size: 12px
}
.boxed-group.dangerzone>h3 {
    background-color: #df3e3e;
    border: 1px solid #a00;
    color: #fff;
    text-shadow: 0 -1px 0 #900
}
.boxed-group.dangerzone .boxed-group-inner {
    border-top: 0
}
.boxed-group.condensed>h3 {
    padding: 6px 6px 7px;
    font-size: 12px
}
.boxed-group.condensed>h3 .octicon {
    padding: 0 6px 0 2px
}
.one-half .boxed-group,
.dashboard-sidebar .boxed-group {
    margin-bottom: 20px
}
.boxed-group .bleed-flush {
    width: 100%;
    padding: 0 10px;
    margin-left: -10px
}
.boxed-group .compact {
    margin-top: 10px;
    margin-bottom: 10px
}
.boxed-group-inner {
    padding: 1px 10px;
    background: #fff;
    border: 1px solid #d8d8d8;
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px;
    color: #666;
    font-size: 13px
}
.boxed-group-inner .markdown-body {
    padding: 20px 10px 10px;
    font-size: 13px
}
.boxed-group-inner.markdown-body {
    padding-top: 10px;
    padding-bottom: 10px
}
.boxed-group-inner.seamless {
    padding: 0
}
.boxed-group-inner h4 {
    margin: 15px 0 -5px;
    font-size: 14px;
    color: #000
}
.boxed-group-inner .tabnav {
    margin-left: -10px;
    margin-right: -10px;
    padding-left: 10px;
    padding-right: 10px
}
.boxed-group-inner .tabnav-tab.selected {
    border-top: 1px solid #ddd
}
.boxed-group-inner .help {
    clear: both;
    margin: 1em -10px 0;
    padding: 1em 10px 1em 35px;
    border-top: 1px solid #ddd;
    color: #767676
}
.boxed-group-inner .help .octicon {
    margin-left: -25px;
    margin-right: 5px
}
.boxed-group-inner .boxed-group-list+.help {
    margin-top: 0
}
.boxed-group-inner .flash-global {
    margin-left: -10px;
    margin-right: -10px;
    border-top: 0
}
.boxed-action {
    float: right;
    margin-left: 10px
}
.boxed-group-action {
    float: right;
    margin: 6px 10px 0 0;
    position: relative;
    z-index: 2
}
.boxed-group-action.flush {
    margin-top: 0;
    margin-right: 0
}
.boxed-group-action>button {
    background-color: transparent;
    border: 0;
    -webkit-appearance: none
}
.boxed-group-icon {
    padding: 4px;
    color: #767676
}
.field-with-errors {
    display: inline
}
.compact-options {
    margin: -6px 0 13px
}
.compact-options>li {
    margin: 0 12px 0 0;
    display: inline-block;
    list-style-type: none;
    font-weight: bold
}
.compact-options>li label {
    float: left
}
.compact-options>li .spinner {
    float: left;
    width: 16px;
    height: 16px;
    margin-left: 5px;
    display: block
}
.boxed-group-list {
    list-style: none;
    margin: 0
}
.boxed-group-list:first-child>li:first-child {
    border-top: 0
}
.boxed-group-list>li {
    display: block;
    margin-left: -10px;
    margin-right: -10px;
    padding: 5px 10px;
    line-height: 23px;
    border-bottom: 1px solid #e5e5e5
}
.boxed-group-list>li:hover {
    background: #ffe
}
.boxed-group-list>li:first-child {
    border-top: 1px solid #ddd
}
.boxed-group-list>li:last-of-type {
    border-bottom: 0
}
.boxed-group-list>li.selected {
    background: #e5f9e2
}
.boxed-group-list>li.approved .btn-sm,
.boxed-group-list>li.rejected .btn-sm {
    display: none
}
.boxed-group-list>li.approved:before {
    margin-right: 5px;
    font: normal normal 16px/1 "octicons";
    display: inline-block;
    text-decoration: none;
    -webkit-font-smoothing: antialiased;
    content: "\f03a";
    color: #5ec051
}
.boxed-group-list>li.rejected:before {
    margin-right: 5px;
    font: normal normal 16px/1 "octicons";
    display: inline-block;
    text-decoration: none;
    -webkit-font-smoothing: antialiased;
    content: "\f050";
    color: #bc0000
}
.boxed-group-list>li.rejected a {
    text-decoration: line-through
}
.boxed-group-list>li img {
    margin-top: -2px;
    margin-right: 4px;
    vertical-align: middle;
    border-radius: 3px
}
.boxed-group-list>li .btn-sm {
    float: right;
    margin: -1px 0 0 10px
}
.boxed-group-list>li .btn-group {
    float: right
}
.boxed-group-list>li .btn-group .btn-sm {
    float: left
}
.boxed-group.flush .boxed-group-list li {
    margin-left: 0;
    width: auto;
    padding-left: 0;
    padding-right: 0
}
.boxed-group-list.standalone {
    margin-top: -1px
}
.boxed-group-list.standalone>li:first-child {
    border-top: 0
}
.boxed-group-table {
    width: 100%;
    text-align: left
}
.boxed-group-table tr:last-child td {
    border-bottom: 0
}
.boxed-group-table th {
    padding: 9px;
    border-bottom: 1px solid #eee;
    background-color: #fafafa
}
.boxed-group-table td {
    padding: 9px;
    border-bottom: 1px solid #eee;
    vertical-align: top
}
#ajax-error-message {
    display: none;
    position: fixed;
    top: -200px;
    left: 50%;
    width: 974px;
    z-index: 9999;
    margin: 0 3px;
    margin-left: -487px;
    -webkit-transition: top 0.5s ease-in-out;
    transition: top 0.5s ease-in-out
}
#ajax-error-message.visible {
    top: 0
}
#ajax-error-message>.octicon-alert {
    vertical-align: text-top
}
.boxed-group-success,
.boxed-group-warning,
.boxed-group-info {
    padding: 10px 15px;
    margin: -1px -10px 0;
    border-style: solid;
    border-width: 1px 0
}
.boxed-group-success .btn-sm,
.boxed-group-warning .btn-sm,
.boxed-group-info .btn-sm {
    margin: -5px 0
}
.boxed-group-success:first-child,
.boxed-group-warning:first-child,
.boxed-group-info:first-child {
    border-top: 0
}
.boxed-group-success {
    color: #22662c;
    background-color: #e2f9e5;
    border-color: #bad3be
}
.boxed-group-warning {
    color: #4c4a42;
    background-color: #fff9ea;
    border-color: #dfd8c2
}
.boxed-group-info {
    color: inherit;
    border-color: inherit
}
.avatar-stack .avatar {
    display: inline-block;
    width: 20px;
    height: 20px;
    margin-right: -15px;
    border-radius: 2px;
    -webkit-transition: margin 0.2s ease-in-out;
    transition: margin 0.2s ease-in-out;
    background-color: #fff;
    border-right: 1px solid #fff;
    z-index: 2;
    position: relative
}
.avatar-stack .avatar:first-child {
    z-index: 3
}
.avatar-stack .avatar:last-child {
    margin-right: 0;
    z-index: 1
}
.avatar-stack:hover .avatar {
    margin-right: 3px
}
.avatar-stack:hover .avatar:last-child {
    margin-right: 0
}
.conversation-list-heading {
    margin: 35px 0 10px;
    height: 0;
    text-align: center;
    font-size: 16px;
    font-weight: normal;
    color: #999;
    border-bottom: 1px solid #ddd
}
.conversation-list-heading .inner {
    display: inline-block;
    position: relative;
    top: -10px;
    padding: 0 5px;
    background: #fff
}
.simple-conversation-list {
    margin: 15px 0;
    font-size: 13px;
    color: #999
}
.simple-conversation-list>li {
    margin: 0;
    padding: 11px 0 8px;
    list-style-type: none;
    border-top: 1px solid #eee
}
.simple-conversation-list>li:first-child {
    border-top: 0
}
.simple-conversation-list>li .title {
    font-weight: bold
}
.simple-conversation-list>li .num {
    color: #999
}
.simple-conversation-list>li .state {
    margin-right: 3px;
    margin-top: -3px;
    padding-top: 2px;
    padding-bottom: 2px
}
.simple-conversation-list>li .meta {
    float: right;
    margin-left: 10px
}
.simple-conversation-list.varied-states>li {
    padding-left: 90px
}
.simple-conversation-list.varied-states>li:before {
    display: table;
    content: ""
}
.simple-conversation-list.varied-states>li:after {
    display: table;
    clear: both;
    content: ""
}
.simple-conversation-list.varied-states>li .state {
    float: left;
    width: 80px;
    margin-left: -90px
}
.pagehead {
    position: relative;
    padding-top: 20px;
    padding-bottom: 20px;
    margin-bottom: 20px;
    border-bottom: 1px solid #eee
}
.pagehead.admin {
    background: url(https://github.com/images/modules/pagehead/background-yellowhatch-v3.png) 0 0 repeat-x
}
.pagehead ul.pagehead-actions {
    z-index: 21;
    float: right;
    margin: 0
}
.pagehead ul.pagehead-actions .feed-icon {
    margin-top: 5px
}
.pagehead .path-divider {
    margin: 0 0.25em
}
.pagehead h1 {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 20px;
    font-weight: normal;
    line-height: 28px
}
.pagehead h1 strong {
    font-weight: bold
}
.pagehead h1 .avatar {
    margin-top: -2px;
    margin-right: 9px;
    margin-bottom: -2px
}
.pagehead .account-switcher {
    margin-top: -3px;
    margin-bottom: -3px
}
.pagehead-actions>li {
    float: left;
    margin: 0 10px 0 0;
    font-size: 11px;
    color: #333;
    list-style-type: none
}
.pagehead-actions>li:last-child {
    margin-right: 0
}
.pagehead-actions .octicon-mute {
    color: #c00
}
.pagehead-actions .select-menu {
    position: relative
}
.pagehead-actions .select-menu:before {
    display: table;
    content: ""
}
.pagehead-actions .select-menu:after {
    display: table;
    clear: both;
    content: ""
}
.pagehead-actions .select-menu-modal-holder {
    top: 100%
}
.context-loader {
    position: absolute;
    top: 0;
    left: 50%;
    z-index: 20;
    width: 154px;
    padding: 10px 10px 10px 30px;
    margin-left: -75px;
    font-size: 12px;
    font-weight: bold;
    color: #666;
    background: url(https://github.com/images/spinners/octocat-spinner-16px.gif) 10px 50% no-repeat #eee;
    border: 1px solid #ddd;
    border-top: 1px solid #fff;
    border-radius: 5px;
    border-top-left-radius: 0;
    border-top-right-radius: 0
}
@media screen and (-webkit-min-device-pixel-ratio: 2),
screen and (max--moz-device-pixel-ratio: 2) {
    .context-loader {
        background: url(https://github.com/images/spinners/octocat-spinner-32-EAF2F5.gif) 10px 50% no-repeat #eee;
        background-size: 16px auto
    }
}
.pagehead-nav {
    float: right;
    margin-bottom: -20px
}
.pagehead-nav-item {
    float: left;
    padding: 6px 10px 21px;
    margin-left: 20px;
    font-size: 14px;
    color: #767676
}
.pagehead-nav-item:hover {
    color: #333;
    text-decoration: none
}
.pagehead-nav-item.selected {
    color: #333;
    border-bottom: 2px solid #d26911
}
.pagehead-nav-item+.btn-outline {
    margin-top: -1px;
    margin-left: 20px
}
.pagehead-tabs {
    margin-top: 20px;
    margin-bottom: -21px;
    overflow: auto
}
.pagehead-tabs-item {
    float: left;
    padding: 8px 15px 11px;
    color: #666;
    white-space: nowrap;
    border: solid transparent;
    border-width: 3px 1px 1px;
    border-radius: 3px 3px 0 0
}
.pagehead-tabs-item .octicon {
    color: #bbb;
    vertical-align: text-top
}
.pagehead-tabs-item .counter {
    color: #777
}
.pagehead-tabs-item:hover {
    color: #333;
    text-decoration: none
}
.pagehead-tabs-item.selected {
    font-weight: bold;
    color: #333;
    background-color: #fff;
    border-color: #d26911 #e5e5e5 transparent
}
.pagehead-tabs-item.selected>.octicon {
    color: inherit
}
.progress-bar {
    display: block;
    overflow: hidden;
    background-color: #eee;
    border-radius: 3px;
    height: 15px
}
.progress-bar .progress {
    height: 100%;
    display: block;
    background-color: #6cc644
}
.protip {
    margin-top: 20px;
    text-align: center
}
.protip code {
    padding: 2px;
    background-color: #f4f4f4;
    border-radius: 3px
}
.protip-callout {
    padding: 8px 10px;
    margin-bottom: 20px;
    color: #4c4a42;
    text-align: left;
    border: solid 1px #eee;
    border-radius: 3px
}
.repo-list {
    position: relative
}
.repo-list .participation-graph {
    position: absolute;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: -1
}
.repo-list .participation-graph.disabled {
    display: none
}
.repo-list .participation-graph .bars {
    position: absolute;
    bottom: 0
}
.repo-list-item {
    position: relative;
    padding-top: 30px;
    padding-bottom: 30px;
    list-style: none;
    border-bottom: 1px solid #eee
}
.repo-list-name {
    margin: 0 0 8px;
    font-size: 20px;
    line-height: 1.2
}
.repo-list-name .prefix,
.repo-list-name .slash {
    font-weight: normal
}
.repo-list-name .slash {
    margin-right: -4px;
    margin-left: -4px
}
.repo-list-description {
    max-width: 550px;
    margin-top: 8px;
    margin-bottom: 0;
    font-size: 14px;
    color: #666
}
.repo-list-stats {
    margin-top: 6px;
    float: right;
    font-size: 12px;
    font-weight: bold;
    color: #888
}
.repo-list-stats .repo-list-stat-item {
    margin-left: 8px;
    display: inline-block;
    color: #888;
    text-decoration: none
}
.repo-list-stats .repo-list-stat-item:hover {
    color: #4078c0
}
.repo-list-stats .repo-list-stat-item>.octicon {
    font-size: 14px
}
.repo-list-info {
    display: inline-block;
    height: 100%;
    margin-top: 0;
    margin-bottom: 0;
    font-size: 12px;
    color: #888;
    vertical-align: middle
}
.repo-list-info .octicon {
    margin-top: -3px;
    font-size: 12px;
    vertical-align: middle
}
.repo-list-meta {
    display: block;
    margin-top: 8px;
    margin-bottom: 0;
    font-size: 13px;
    color: #888
}
.repo-list-meta .avatar {
    margin-top: -2px
}
.repo-list-meta a:hover {
    text-decoration: none
}
.task-list-item {
    list-style-type: none
}
.task-list-item label {
    font-weight: normal
}
.task-list-item.enabled label {
    cursor: pointer
}
.task-list-item+.task-list-item {
    margin-top: 3px
}
.task-list-item-checkbox {
    margin: 0 0.35em 0.25em -1.6em;
    vertical-align: middle
}
.about-header {
    height: 300px;
    background-color: #111;
    background-image: url(https://github.com/images/modules/about/about-header.jpg);
    background-position: 50%;
    background-size: cover
}
.about-header.team {
    background-image: url(https://github.com/images/modules/about/team-header.jpg)
}
.about-header.press {
    background-image: url(https://github.com/images/modules/about/press-header.jpg)
}
.about-header.jobs {
    background-image: url(https://github.com/images/modules/about/jobs-header.jpg)
}
.about-menu {
    margin-bottom: 40px;
    border-bottom: 1px solid #eee
}
.about-menu-link {
    float: left;
    width: 25%;
    padding: 20px 20px 17px;
    font-size: 18px;
    color: #767676;
    text-align: center;
    border-bottom: 3px solid #fff
}
.about-menu-link:hover {
    color: #4078c0;
    text-decoration: none;
    border-bottom-color: #f5f5f5
}
.about-menu-link.selected {
    font-weight: bold;
    color: #333;
    border-bottom-color: #d26911
}
.about-content {
    padding: 0;
    font-size: 18px;
    font-weight: 300;
    line-height: 1.5
}
.about-content p {
    margin-top: 0
}
.about-content hr {
    margin-top: 30px;
    margin-bottom: 30px;
    border-bottom-color: #eee
}
.about-lead {
    padding-right: 8%;
    padding-left: 8%;
    margin-bottom: 40px;
    font-size: 24px;
    text-align: center
}
.social-callout-twitter {
    padding-right: 18px;
    margin-top: 40px;
    margin-bottom: 20px
}
.social-callout-twitter:hover .social-callout-twitter-logo {
    background-image: url(https://github.com/images/icons/twitter-white.png)
}
.social-callout-twitter-logo {
    display: inline-block;
    width: 32px;
    height: 32px;
    vertical-align: middle;
    background: url(https://github.com/images/icons/twitter.png) 0 0 no-repeat;
    background-size: 32px auto
}
.press-mentions {
    margin-top: 10px
}
.press-mentions li {
    margin-bottom: 15px;
    list-style-type: none
}
.press-mentions cite {
    display: block;
    font-size: 13px;
    font-style: normal;
    font-weight: normal;
    color: #666
}
.press-date {
    margin-top: 25px;
    margin-bottom: 0;
    color: #808080
}
.press-info {
    margin: 20px 0 30px
}
.press-info:before {
    display: table;
    content: ""
}
.press-info:after {
    display: table;
    clear: both;
    content: ""
}
.press-info-link {
    position: relative;
    display: block;
    padding: 10px 15px;
    font-weight: normal;
    line-height: 2;
    background-color: #fafafa;
    border: 1px solid #ddd
}
.press-info-link:first-child {
    border-top-left-radius: 4px;
    border-top-right-radius: 4px
}
.press-info-link:last-child {
    border-bottom-right-radius: 4px;
    border-bottom-left-radius: 4px
}
.press-info-link+.press-info-link {
    margin-top: -1px
}
.press-info-link .mega-octicon {
    width: 32px;
    margin-right: 7px;
    color: #aec0d0;
    text-align: center;
    vertical-align: middle
}
.press-info-link:hover {
    z-index: 2;
    color: #fff;
    text-decoration: none;
    background-color: #4078c0;
    border-color: #4078c0
}
.press-info-link:hover .mega-octicon {
    color: inherit
}
.hubbers-list {
    margin: -6px;
    list-style: none
}
.hubbers-list:before {
    display: table;
    content: ""
}
.hubbers-list:after {
    display: table;
    clear: both;
    content: ""
}
.hubbers-list-item {
    display: block;
    float: left;
    width: 128px;
    height: 128px;
    margin: 6px;
    text-align: center
}
.hubbers-list-item img {
    display: block;
    width: 100%;
    height: 100%;
    background-color: #fff
}
.hubbers-list-item .hubber-name {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    display: block;
    padding-top: 100px;
    font-size: 12px;
    font-weight: bold;
    color: #fff;
    text-align: center;
    text-shadow: 0 0 4px #000;
    background-color: rgba(0, 0, 0, 0.25);
    background-image: -webkit-linear-gradient(rgba(0, 0, 0, 0.01) 50%, rgba(0, 0, 0, 0.25));
    background-image: linear-gradient(rgba(0, 0, 0, 0.01) 50%, rgba(0, 0, 0, 0.25));
    opacity: 0;
    -webkit-transition: opacity 0.25s ease-in-out;
    transition: opacity 0.25s ease-in-out
}
.hubbers-list-item a {
    position: absolute;
    display: block;
    width: 128px;
    height: 128px
}
.hubbers-list-item a:hover .hubber-name {
    opacity: 1
}
.jobs-open-positions {
    padding: 20px;
    margin-top: 10px;
    background-color: #f1f6fb;
    border: solid 1px #d0e5f8;
    border-radius: 3px
}
.jobs-open-positions ul {
    margin-top: 24px;
    list-style: none
}
.account-membership-form .become-a-member,
.account-membership-form .already-a-member {
    display: none
}
.account-membership-form.is-member .already-a-member {
    display: block
}
.account-membership-form.is-not-member .become-a-member {
    display: block
}
.cvv-hint {
    position: relative;
    padding-right: 15px
}
.cvv-hint:hover .cvv-hint-tooltip {
    display: block
}
.cvv-hint-tooltip {
    display: none;
    position: absolute;
    border: 1px solid #d0d0d0;
    padding: 15px;
    z-index: 1000;
    background-color: #fff;
    left: 100%;
    top: -150px;
    box-shadow: 0 0 5px #ebebeb, 0 0 5px #ebebeb, 0 0 5px #ebebeb, 0 0 5px #ebebeb, 0 0 5px #ebebeb
}
.credit-card {
    border: 1px solid #ddd;
    width: 250px;
    padding: 20px;
    height: 150px;
    position: relative;
    margin-top: 5px;
    border-radius: 10px
}
.credit-card.amex {
    margin-top: 15px
}
.credit-card.amex .title {
    color: #fff;
    font-family: "Arial Black", "Arial Bold", Gadget, sans-serif;
    text-shadow: 1px 0 0 #ddd, -1px 0 0 #ddd, 0 1px 0 #ddd, 0 -1px 0 #ddd, -1px -1px 0 #ddd, 1px 1px 0 #ddd, -1px 1px 0 #ddd, 1px -1px 0 #ddd;
    position: relative;
    z-index: 1;
    top: -5px;
    text-align: center;
    letter-spacing: 1px;
    -webkit-transform: scale(1.3, 1);
    -ms-transform: scale(1.3, 1);
    transform: scale(1.3, 1)
}
.credit-card.amex .card-number {
    margin-top: 40px;
    font-size: 15px;
    display: inline-block;
    white-space: nowrap;
    position: relative
}
.credit-card.amex .gladiator {
    position: absolute;
    left: 50%;
    top: 50px;
    margin-left: -35px;
    height: 80px;
    width: 70px;
    border: 3px solid #fff;
    box-shadow: 0 0 1px #aaa;
    border-top-left-radius: 35px 40px;
    border-bottom-left-radius: 35px 40px;
    border-top-right-radius: 35px 40px;
    border-bottom-right-radius: 35px 40px;
    background-color: #e0e0e0
}
.credit-card.normal .strap {
    background-color: #555;
    height: 20px;
    margin: -5px -20px 15px
}
.credit-card.normal .signature {
    display: inline-block;
    background-color: #e9e9e9;
    white-space: nowrap;
    font-family: "Brush Script MT", cursive;
    height: 30px;
    width: 150px;
    font-size: 17px;
    color: #aaa;
    letter-spacing: -1px;
    line-height: 33px;
    text-indent: 10px
}
.credit-card .cvv {
    left: -7px;
    top: -10px;
    border: 2px solid #f00;
    padding: 2px 5px;
    line-height: 1;
    font-family: monospace;
    font-size: 10px;
    border-top-left-radius: 20px 10px;
    border-bottom-left-radius: 20px 10px;
    border-top-right-radius: 20px 10px;
    border-bottom-right-radius: 20px 10px;
    text-align: center;
    position: relative;
    display: inline-block
}
.credit-card .cvv span {
    position: absolute;
    right: 100%;
    margin-right: 5px;
    color: #767676
}
.credit-card .text {
    text-transform: uppercase;
    font-size: 7px;
    display: block;
    line-height: 1.1;
    font-weight: bold;
    font-family: monospace
}
.billing-addon-items table input {
    width: 5em
}
.billing-addon-items td {
    vertical-align: middle;
    border-bottom: 0
}
.billing-addon-items td.fixed {
    width: 150px
}
.billing-addon-items td.black {
    color: #000
}
.billing-addon-items tr {
    border-bottom: 1px solid #eee
}
.billing-addon-items tr:last-child {
    border-bottom-width: 0
}
.billing-addon-items tr:nth-child(even) {
    background-color: #fafafa
}
.billing-addon-items tr.total-row {
    background-color: #fff;
    color: #bd2c00
}
.billing-addon-items tr.dark-row {
    border-bottom-width: 1px;
    background-color: #fafafa
}
.billing-addon-items .new-addon-items {
    margin-left: 5px
}
.billing-addon-items .addon-cost {
    color: #767676
}
.billing-addon-items .discounted-original-price {
    color: #666
}
.billing-addon-items .form-submit,
.billing-addon-items .payment-method {
    margin-left: 10px
}
.billing-addon-items .payment-summary {
    margin-left: 10px;
    margin-right: 10px
}
.billing-credit-card .javascript-disabled-overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
    display: none;
    background-color: #fff;
    opacity: 0.5
}
.billing-credit-card.disabled .javascript-disabled-overlay,
.billing-credit-card.unsupported .javascript-disabled-overlay {
    display: block
}
.billing-actions {
    padding-bottom: 10px
}
.help.billing-next-payment-help {
    margin-top: 0
}
.billing-extra-box {
    border-left: 3px solid #eee;
    padding-left: 10px;
    margin: 10px 0 0
}
.billing-section {
    border-bottom: 1px solid #eee;
    padding: 15px 10px;
    line-height: 1.5em
}
.billing-section.oneliner {
    padding-bottom: 14px
}
.billing-section.oneliner .action-button {
    margin-top: -4px;
    margin-bottom: -4px
}
.billing-section p {
    margin: 10px 0 0
}
.billing-section .disabled-message {
    color: #bd2c00
}
.billing-section .action-button {
    float: right;
    margin-bottom: 5px;
    margin-left: 10px
}
.billing-section .octicon-btn {
    float: right;
    padding: 4px;
    margin-left: 5px
}
.billing-section .section-label {
    width: 85px;
    color: #767676;
    font-weight: normal;
    text-align: right;
    position: absolute
}
.billing-section .section-content {
    color: #333;
    margin-left: 100px
}
.billing-section .pending-invitations-link,
.billing-section .subtle-link {
    color: #767676
}
.billing-section:last-child {
    border-bottom: 0
}
.billing-section.info-section {
    background: #f9f9f9;
    border-bottom: 0;
    color: #767676;
    overflow: hidden
}
.billing-section.info-section .octicon-info {
    font-size: 30px;
    color: #ddd
}
.billing-section .usage-bar {
    margin: 5px 0 0;
    background: #eee;
    border-radius: 20px;
    max-width: 304px;
    width: 100%
}
.billing-section .usage-bar.exceeded .progress {
    background: #bd2c00
}
.billing-section .usage-bar .progress {
    max-width: 100%;
    border-radius: 20px;
    height: 5px;
    background: #67d07c;
    position: relative
}
.billing-section .usage-bar .progress.no-highlight {
    background: #999
}
.billing-usage-summary {
    margin-bottom: 20px
}
.billing-data-usage-meter {
    margin-bottom: 15px
}
.billing-data-usage-meter:last-child {
    margin-bottom: 5px
}
.packs-table .desc {
    width: 1%;
    white-space: nowrap
}
.pack-upgrade-plus {
    color: #6cc644;
    font-weight: bold;
    font-size: 16px;
    text-align: center
}
.lfs-data-icon {
    width: 15px;
    color: #767676;
    margin-right: 5px;
    text-align: center
}
.lfs-data-icon.dark {
    color: #333
}
.setup-wrapper .paypal-container {
    margin-bottom: 30px
}
.setup-wrapper .paypal-logged-in .paypal-container {
    margin-bottom: 10px
}
.payment-methods {
    position: relative
}
.payment-methods .selected-payment-method {
    display: none
}
.payment-methods .selected-payment-method:before {
    display: table;
    content: ""
}
.payment-methods .selected-payment-method:after {
    display: table;
    clear: both;
    content: ""
}
.payment-methods .selected-payment-method.active {
    display: block
}
.payment-methods .pay-with-header {
    margin: 5px 0
}
.payment-methods .pay-with-paypal .setup-creditcard-form,
.payment-methods .pay-with-paypal .paypal-form-actions,
.payment-methods .pay-with-paypal .terms,
.payment-methods .pay-with-paypal .paypal-signed-in,
.payment-methods .pay-with-paypal .paypal-down-flash,
.payment-methods .pay-with-paypal .loading-paypal-spinner {
    display: none
}
.payment-methods.paypal-loading .loading-paypal-spinner {
    display: block
}
.payment-methods.paypal-down .paypal-down-flash {
    display: block
}
.payment-methods.paypal-logged-in .paypal-sign-in {
    display: none
}
.payment-methods.paypal-logged-in .setup-creditcard-form,
.payment-methods.paypal-logged-in .paypal-form-actions,
.payment-methods.paypal-logged-in .terms,
.payment-methods.paypal-logged-in .paypal-signed-in {
    display: block
}
.payment-methods.has-paypal-account .paypal-sign-in {
    display: none
}
.payment-methods.has-paypal-account .paypal-signed-in {
    display: block
}
.contact-us {
    background-color: #fff9ea;
    border-bottom: 1px solid #eee;
    font-size: 12px;
    padding: 6px 10px
}
.contact-us .octicon {
    margin-right: 5px;
    color: #767676
}
.paypal-label {
    font-weight: bold;
    margin: 15px 0 10px
}
.paypal-container {
    margin-bottom: 15px;
    vertical-align: top;
    display: inline-block;
    background-color: #f9f9f9;
    border-radius: 4px
}
#braintree-paypal-loggedin {
    background-position: 12px 50% !important;
    border: 1px solid #ddd !important;
    padding: 11px 16px !important;
    border-radius: 4px
}
#bt-pp-name {
    margin-left: 20px !important
}
#bt-pp-email {
    margin-left: 15px !important
}
#bt-pp-cancel {
    font-size: 0 !important;
    color: #a00 !important;
    text-decoration: none !important;
    font-family: octicons !important;
    line-height: 1 !important;
    display: inline-block;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale
}
#bt-pp-cancel:before {
    font-size: 16px !important;
    content: "\f081"
}
.payment-history .no-payments {
    margin: 0;
    border-top: 0
}
.payment-history .id,
.payment-history .date,
.payment-history .method,
.payment-history .receipt,
.payment-history .status,
.payment-history .description,
.payment-history .amount {
    width: 1%;
    white-space: nowrap
}
.payment-history .receipt {
    text-align: center
}
.payment-history .currency,
.payment-history .status {
    color: #767676
}
.payment-history .status-icon {
    width: 14px;
    text-align: center
}
.payment-history .succeeded .status {
    color: #6cc644
}
.payment-history .refunded,
.payment-history .failed {
    background: #f9f9f9
}
.payment-history .refunded td,
.payment-history .failed td {
    opacity: 0.5
}
.payment-history .refunded .receipt,
.payment-history .refunded .status,
.payment-history .failed .receipt,
.payment-history .failed .status {
    opacity: 1
}
.payment-history .refunded .status {
    color: #999
}
.payment-history .failed .status {
    color: #bd2c00
}
.paypal-icon {
    vertical-align: middle;
    margin: 0 2px 0 1px
}
.inline-form-action {
    display: inline
}
.boxed-group .boxed-group-content {
    margin: 10px
}
.billing-container {
    width: 500px;
    margin-top: 80px;
    margin-bottom: 80px
}
.billing-step {
    margin-bottom: 50px
}
.billing-step-title {
    font-weight: 400;
    padding-bottom: 10px;
    border-bottom: 1px solid #e0e0e0;
    margin-bottom: 15px
}
.currency-container .local-currency,
.currency-container .local-currency-block {
    display: none
}
.currency-container.open .local-currency {
    display: inline
}
.currency-container.open .local-currency-block {
    display: block
}
.currency-container.open .default-currency {
    display: none
}
.plan-chooser {
    margin: 10px auto 20px
}
.plan-chooser.on-free .toggle-currency,
.plan-chooser.on-free .currency-notice,
.plan-chooser.on-free .plan-price-group,
.plan-chooser.on-free .coupon-notice {
    display: none
}
.strong-label {
    font-weight: bold;
    margin-bottom: 5px;
    display: inline-block
}
.plan-chooser-repo-menu {
    margin-bottom: 15px
}
.plan-chooser-repo-menu .price-label {
    display: none
}
.plan-chooser-repo-menu .btn-block {
    text-align: left
}
.discounted-original-price,
.per-repo.has-coupon .original-price {
    text-decoration: line-through;
    color: #767676;
    font-weight: normal
}
.billing-managers-abilities-list {
    list-style: none
}
.billing-managers-abilities-list li {
    margin-bottom: 6px
}
.billing-managers-abilities-list .octicon {
    width: 24px;
    text-align: center
}
.billing-managers-abilities-list .octicon-check {
    color: #6cc644
}
.billing-managers-abilities-list .octicon-x {
    color: #bd2c00
}
.billing-manager-input {
    width: 300px
}
.billing-manager-banner {
    border-bottom: 1px solid #eee;
    background: #f9f9f9;
    padding: 30px 20px;
    overflow: hidden;
    margin-bottom: 30px
}
.billing-manager-banner .container {
    position: relative
}
.billing-manager-banner-text {
    color: #555;
    font-size: 14px;
    margin-left: 210px
}
.billing-manager-banner-text .btn {
    margin-top: 8px;
    margin-right: 8px
}
.billing-manager-banner-title {
    font-weight: bold;
    font-size: 12px;
    color: #767676
}
.billing-manager-icon {
    font-size: 180px;
    position: absolute;
    top: -35px;
    left: 0;
    color: #e0e0e0
}
.golden-ticket-banner {
    text-align: center;
    margin-top: 30px;
    margin-bottom: 15px;
    border-top: 1px solid #e6d445
}
.golden-ticket {
    margin-top: -30px;
    height: 60px
}
.golden-ticket-button {
    padding: 30px 20px;
    font-size: 18px;
    font-weight: normal;
    float: left;
    width: 50%
}
.golden-ticket-button .mega-octicon {
    vertical-align: middle;
    margin-right: 10px
}
.golden-ticket-button:first-child {
    border-radius: 3px 0 0 3px
}
.golden-ticket-button:last-child {
    border-left: 0;
    border-radius: 0 3px 3px 0
}
.golden-ticket-confirm .setup-header {
    text-align: center;
    border: 0
}
.heat[data-heat="1"] {
    background-color: #ffeca7
}
.heat[data-heat="2"] {
    background-color: #ffdd8c
}
.heat[data-heat="3"] {
    background-color: #ffdd7c
}
.heat[data-heat="4"] {
    background-color: #fba447
}
.heat[data-heat="5"] {
    background-color: #f68736
}
.heat[data-heat="6"] {
    background-color: #f37636
}
.heat[data-heat="7"] {
    background-color: #ca6632
}
.heat[data-heat="8"] {
    background-color: #c0513f
}
.heat[data-heat="9"] {
    background-color: #a2503a
}
.heat[data-heat="10"] {
    background-color: #793738
}
.blame-breadcrumb .css-truncate-target {
    max-width: 680px
}
.blame-commit,
.blame-commit+.blame-line {
    border-top: 1px solid #e9e9e9
}
.blame-container {
    margin-top: -1px
}
.blame-blob-num {
    background-color: #fdfdfd
}
.blame-commit-info {
    position: relative;
    width: 350px;
    min-width: 350px;
    max-width: 350px;
    padding: 8px 10px;
    vertical-align: top
}
.blame-commit-avatar {
    float: left;
    margin-right: 5px
}
.blame-commit-title {
    font-weight: bold;
    color: #333;
    max-width: 230px;
    display: inline-block;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    line-height: 1.1;
    vertical-align: top
}
.blame-sha {
    font: 11px Consolas, "Liberation Mono", Menlo, Courier, monospace;
    float: right
}
.blame-commit-meta {
    color: #767676;
    font-size: 12px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    line-height: 1.1
}
.line-age {
    width: 2px;
    padding: 0 1px
}
.line-age-legend {
    float: right;
    margin-top: -25px;
    font-size: 12px;
    color: #767676
}
.line-age-legend ol {
    display: inline-block;
    list-style: none;
    margin: 0 5px
}
.line-age-legend ol li {
    display: inline-block;
    width: 8px;
    height: 10px
}
#blog-main .pagehead h1 {
    margin-top: 3px;
    font-size: 24px
}
.blog-title {
    color: #333
}
.blog-search {
    position: relative;
    float: right
}
.blog-search .blog-search-input {
    padding-left: 28px;
    width: 200px
}
.blog-search .octicon-search {
    position: absolute;
    left: 7px;
    top: 9px;
    z-index: 5;
    color: #767676
}
.blog-search-results em {
    background-color: #faffa6;
    padding: 0.1em
}
.blog-aside {
    float: right;
    width: 200px
}
.blog-aside .btn {
    margin-bottom: 20px;
    text-align: center
}
.blog-aside .menu-container {
    float: none;
    margin-bottom: 30px
}
.blog-aside .rss {
    display: inline-block;
    margin-left: 5px;
    color: #999
}
.blog-aside .rss .octicon {
    float: left;
    margin-right: 5px;
    color: #c9510c
}
.blog-content {
    width: 685px;
    font-family: "Helvetica Neue", Helvetica, Arial, freesans, sans-serif
}
.blog-content h1,
.blog-content h2,
.blog-content h3 {
    font-weight: 500
}
.blog-content .markdown-body h2 {
    font-size: 20px
}
.blog-content .markdown-body h3 {
    font-size: 18px
}
.blog-draft-indicator {
    color: #bd2c00
}
.blog-post {
    margin-bottom: 60px
}
.blog-post-meta {
    list-style: none;
    margin-bottom: 20px;
    color: #999
}
.blog-post-meta .meta-item {
    display: inline;
    padding-right: 20px
}
.blog-post-meta a {
    color: #999
}
.blog-post-meta .octicon,
.blog-post-meta .author-avatar {
    vertical-align: top;
    border-radius: 3px
}
.blog-post-title {
    margin-top: 0;
    margin-bottom: 10px;
    font-size: 32px
}
#blog-home {
    color: #ccc;
    font-size: 15px;
    font-weight: 100;
    margin-right: 10px;
    margin-left: -25px;
    vertical-align: middle
}
#blog-home:hover {
    color: #767676
}
.blog-post-body {
    font-size: 16px;
    line-height: 1.6;
    color: #444
}
.blog-post-body img {
    padding: 3px;
    border: 1px solid #d8d8d8
}
.blog-post-body img.emoji {
    border: 0;
    padding: 0
}
.blog-post-body iframe {
    border: 0;
    width: 100%
}
.blog-feedback {
    margin: 50px 0;
    background-color: #fafafa;
    border: 1px solid #ddd;
    border-bottom-color: #ccc;
    border-radius: 3px;
    box-shadow: inset 0 1px 0 #fff, 0 1px 5px #f1f1f1
}
.blog-feedback-header {
    margin: 0;
    padding: 10px;
    border-bottom: 1px solid #ddd;
    box-shadow: 0 1px 0 #fff;
    font-size: 14px;
    font-weight: bold
}
.blog-feedback-header.with-twitter {
    background: url(https://github.com/images/icons/twitter.png) 648px 1px no-repeat;
    background-size: 32px auto
}
.blog-feedback-description {
    margin: 0;
    padding: 10px;
    color: #767676
}
.branches .page-header {
    margin-bottom: 20px
}
.branches .clear-search {
    display: none
}
.branches .loading-overlay {
    display: none;
    position: absolute;
    top: 0;
    width: 100%;
    height: 100%;
    padding-top: 50px;
    z-index: 20;
    background-color: rgba(255, 255, 255, 0.7);
    text-align: center
}
.branches .loading-overlay .spinner {
    display: inline-block
}
.branches.is-search-mode .clear-search {
    display: inline-block
}
.branches.is-loading .loading-overlay {
    display: block
}
.branches .status {
    display: inline-block;
    width: 16px;
    text-align: center
}
.branches .status .octicon {
    position: relative;
    top: 2px
}
.branches .pull-request-link {
    top: 0;
    display: inline;
    padding: 2px 5px;
    line-height: 1em
}
.branches .branch-actions {
    float: right;
    position: relative;
    top: -3px;
    right: -4px
}
.branches .branch-actions form {
    display: inline
}
.branches .branch-actions .octicon {
    width: 16px;
    text-align: center
}
.branch-groups {
    position: relative
}
.branch-group {
    margin-bottom: 20px;
    width: 100%;
    border-radius: 3px
}
.branch-group:before {
    display: table;
    content: ""
}
.branch-group:after {
    display: table;
    clear: both;
    content: ""
}
.branch-group-heading {
    border: 1px solid #ddd;
    border-bottom: 0;
    padding: 6px 12px;
    background: #f5f5f5;
    text-shadow: 0 1px 0 #fff
}
.branch-group-heading+.branch-summary {
    border-top: 1px solid #ddd
}
.branch-group-heading .branch-name {
    background: #767676;
    color: #fff;
    text-shadow: none
}
.branch-group-name {
    color: #767676;
    font-weight: bold
}
.branch-group-heading:first-child,
.branch-summary:first-child {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.branch-group-heading:last-child,
.branch-summary:last-child {
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px
}
.branches-view-switcher {
    display: inline-block;
    vertical-align: middle
}
.branch-search {
    position: relative;
    float: right;
    vertical-align: middle
}
.branch-search .clear-search {
    position: absolute;
    right: 12px;
    top: 9px;
    color: #999
}
.branch-search-field {
    width: 250px;
    padding-right: 25px
}
.no-results-message {
    padding: 12px;
    border: 1px solid #ddd;
    border-radius: 0 0 3px 3px;
    color: rgba(0, 0, 0, 0.5);
    text-align: center
}
.branch-summary {
    padding: 12px;
    border: 1px solid #ddd;
    border-bottom: 0;
    color: rgba(0, 0, 0, 0.5)
}
.branch-summary:last-child {
    border-bottom: 1px solid #ddd
}
.branch-summary .branch-spinner {
    display: none;
    vertical-align: text-bottom
}
.branch-summary.loading .branch-delete-icon {
    display: none
}
.branch-summary.loading .branch-spinner {
    display: inline-block
}
.branch-summary.is-deleted .existing-branch-summary {
    display: none
}
.branch-summary.is-deleted .deleted-branch-summary {
    display: block
}
.deleted-branch-summary {
    display: none
}
.deleted-branch-summary .css-truncate-target {
    max-width: 500px
}
.deleted-branch-summary .branch-name {
    opacity: 0.5;
    text-decoration: line-through
}
.deleted-branch-summary .branch-spinner {
    float: right;
    position: relative;
    top: 4px;
    right: 5px
}
.pr-details {
    display: inline-block;
    width: 144px;
    text-align: right
}
.pr-details .state {
    padding: 1px 5px;
    font-size: 12px;
    width: 75px;
    margin-left: 5px;
    text-decoration: none
}
.pr-details .state .octicon {
    font-size: 14px
}
.branch-delete {
    display: inline-block;
    color: #bd2c00;
    margin: 4px 2px 0 8px
}
.branch-delete.disabled {
    color: #ddd
}
.more-branches {
    display: block;
    padding: 6px;
    width: 100%;
    border: 1px solid #dae5eb;
    border-radius: 0 0 3px 3px;
    text-align: center;
    text-decoration: none;
    background: #f1f7fa;
    color: #4078c0
}
.more-branches:hover {
    background: #e6f1f6
}
.more-branches .octicon {
    position: relative;
    top: 1px;
    margin-left: 5px
}
.branch-details {
    display: inline-block;
    width: 490px;
    margin-right: 10px
}
.branch-details .css-truncate-target {
    max-width: 240px
}
.branch-details .icon-shield {
    margin-top: 1px;
    margin-right: 2px;
    opacity: 0.45
}
.branch-meta {
    color: #aaa;
    font-size: 12px;
    line-height: 20px
}
.default-label {
    width: 150px;
    text-align: center;
    display: inline-block;
    vertical-align: top
}
.default-label .sha {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.default-label .sha .ellipses {
    color: inherit;
    font-family: Helvetica, arial, nimbussansl, liberationsans, freesans, clean, sans-serif, "Segoe UI Emoji", "Segoe UI Symbol"
}
.default-label .sha .octicon {
    padding-right: 4px
}
@-webkit-keyframes branch-bar-slide {
    from {
        width: 0%
    }
    to {
        width: 100%
    }
}
@keyframes branch-bar-slide {
    from {
        width: 0%
    }
    to {
        width: 100%
    }
}
.branch-a-b-count {
    display: inline-block;
    vertical-align: middle
}
.branch-a-b-count .count-half {
    float: left;
    width: 90px;
    text-align: right;
    padding-bottom: 6px;
    position: relative
}
.branch-a-b-count .count-half:last-child {
    text-align: left;
    border-left: 1px solid #bbb
}
.branch-a-b-count .count-value {
    position: relative;
    top: -1px;
    display: block;
    padding: 0 3px;
    font-size: 10px
}
.branch-a-b-count .bar {
    position: absolute;
    min-width: 3px;
    height: 4px
}
.branch-a-b-count .meter {
    position: absolute;
    height: 4px;
    background-color: #ccc;
    -webkit-animation: branch-bar-slide 0.3s ease forwards 0.5s;
    animation: branch-bar-slide 0.3s ease forwards 0.5s
}
.branch-a-b-count .meter.zero {
    background-color: transparent
}
.branch-a-b-count .bar-behind {
    right: 0;
    border-radius: 3px 0 0 3px
}
.branch-a-b-count .bar-behind .meter {
    right: 0;
    border-radius: 3px 0 0 3px
}
.branch-a-b-count .bar-ahead {
    left: 0;
    border-radius: 0 3px 3px 0
}
.branch-a-b-count .bar-ahead .meter {
    border-radius: 0 3px 3px 0
}
.branch-a-b-count .bar-ahead.even,
.branch-a-b-count .bar-behind.even {
    background: #eee;
    min-width: 2px
}
.capped-cards {
    list-style: none
}
.capped-cards .capped-card {
    float: left;
    width: 450px
}
.capped-card {
    border-radius: 2px;
    border: 1px solid #ddd;
    list-style: none;
    margin: 10px
}
.capped-card:before {
    display: table;
    content: ""
}
.capped-card:after {
    display: table;
    clear: both;
    content: ""
}
.capped-card:nth-child(odd) {
    margin-left: 0
}
.capped-card:nth-child(even) {
    margin-right: 0
}
.capped-card h3 {
    margin: 0;
    border-bottom: 1px solid #eee;
    line-height: 100%;
    padding: 10px
}
.capped-card>p {
    border-bottom: 1px solid #eee;
    color: #767676;
    display: block;
    font-size: 15px;
    line-height: 100%;
    margin: 0;
    padding: 0 10px 10px
}
.capped-card-content {
    background: #f7f7f7;
    display: block
}
.capped-card-content:before {
    display: table;
    content: ""
}
.capped-card-content:after {
    display: table;
    clear: both;
    content: ""
}
.clone-url h5 {
    margin-top: 0;
    margin-bottom: 10px
}
.clone-url .input-group {
    width: 100%
}
.clone-url input.input-mini {
    font-size: 11px;
    color: #767676
}
.commit-form {
    position: relative;
    padding: 15px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.commit-form:after,
.commit-form:before {
    position: absolute;
    top: 11px;
    left: -16px;
    right: 100%;
    width: 0;
    height: 0;
    display: block;
    content: " ";
    border-color: transparent;
    border-style: solid solid outset;
    pointer-events: none
}
.commit-form:after {
    border-width: 7px;
    border-right-color: #fff;
    margin-top: 1px;
    margin-left: 2px
}
.commit-form:before {
    border-right-color: #ddd;
    border-width: 8px
}
.commit-form .input-block {
    margin-top: 10px;
    margin-bottom: 10px
}
.commit-form-avatar {
    float: left;
    margin-left: -64px;
    border-radius: 3px
}
.commit-form-actions:before {
    display: table;
    content: ""
}
.commit-form-actions:after {
    display: table;
    clear: both;
    content: ""
}
.commit-form-actions .btn {
    margin-right: 5px
}
.commit-form-actions .check-for-fork {
    line-height: 34px
}
.copyable-terminal {
    position: relative;
    padding: 10px 55px 10px 10px;
    background-color: #f7f7f7;
    border-radius: 3px
}
.copyable-terminal-content {
    overflow: auto
}
.copyable-terminal-button {
    position: absolute;
    top: 5px;
    right: 5px
}
.copyable-terminal-button .zeroclipboard-button {
    float: right
}
.copyable-terminal-button .zeroclipboard-button .octicon {
    padding-left: 1px;
    margin: 0 auto
}
.logged_out.enter-coupon {
    background-color: #f9f9f9
}
.logged_out.enter-coupon .coupon-form-body {
    margin-bottom: -20px;
    background-image: none;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.075), inset 1px 0 #fff, 0 0 200px #fff
}
.logged_out.enter-coupon .header-logged-out {
    background-color: #fff
}
.logged_out.enter-coupon .site-footer {
    border-top: 0
}
.coupons .setup-plans td img {
    vertical-align: middle;
    margin-top: -2px
}
.coupons .coupon-form-body {
    width: 230px;
    margin: 100px auto 60px;
    padding: 20px;
    font-size: 14px;
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    text-align: center;
    background-color: #fff;
    background-image: -webkit-linear-gradient(#fefefe, #fafafa);
    background-image: linear-gradient(#fefefe, #fafafa);
    border: 1px solid #ccc;
    border-radius: 3px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.075), inset 1px 0 #fff
}
.coupons .coupon-form-body .input-block {
    margin-bottom: 15px
}
.coupons .coupon-form-body .btn {
    display: block;
    width: 100%
}
.coupon-icon {
    width: 80px;
    height: 80px;
    margin: 0 auto 15px;
    color: #4078c0;
    border: 1px solid #dedede;
    border-radius: 40px
}
.coupon-icon .mega-octicon {
    margin-left: 2px;
    font-size: 48px;
    line-height: 80px
}
.coupon-signin-title {
    margin-top: 40px
}
.coupon-title {
    margin-bottom: 20px;
    font-weight: 500
}
.coupons-list-options {
    margin-bottom: 15px
}
.coupons-list-options .select-menu,
.coupons-list-options .btn-group {
    display: inline-block;
    margin-right: 10px
}
.coupons-list-options .pagination {
    float: right;
    margin: 0
}
dl.form.developer-select-account {
    margin-top: 0
}
.developer-wrapper .setup-info-module .features-list {
    margin-left: 16px
}
.developer-wrapper .setup-info-module .features-list .octicon {
    margin-left: -17px
}
.developer-thanks h2 {
    font-size: 38px;
    font-weight: normal
}
.developer-thanks .hook {
    margin-top: 2px;
    margin-bottom: 30px;
    font-size: 18px;
    font-weight: 300;
    color: #666
}
.developer-thanks-image {
    position: relative;
    bottom: -45px;
    float: left;
    width: 400px
}
.developer-thanks-section {
    margin: 130px 0 0 470px
}
.developer-next-steps {
    list-style: none;
    font-size: 18px;
    font-weight: 300
}
.developer-next-steps li {
    margin-top: 10px
}
.developer-next-steps li:first-child {
    margin-top: 0
}
.developer-next-steps .mega-octicon {
    position: relative;
    top: 5px;
    margin-right: 10px;
    font-size: 32px;
    color: #6cc644
}
.blob-wrapper {
    overflow-x: auto;
    overflow-y: hidden;
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px
}
.diff-table {
    width: 100%;
    border-collapse: separate
}
.diff-table .line-comments {
    padding: 10px;
    vertical-align: top
}
.diff-table .line-comments:first-child+.empty-cell {
    border-left-width: 1px
}
.diff-table tr:not(:last-child) .line-comments {
    border-top: 1px solid #eee;
    border-bottom: 1px solid #eee
}
.blob-num {
    width: 1%;
    min-width: 50px;
    white-space: nowrap;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px;
    line-height: 18px;
    color: rgba(0, 0, 0, 0.3);
    vertical-align: top;
    text-align: right;
    border: solid #eee;
    border-width: 0 1px 0 0;
    cursor: pointer;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    padding-left: 10px;
    padding-right: 10px
}
.blob-num:hover {
    color: rgba(0, 0, 0, 0.6)
}
.blob-num:before {
    content: attr(data-line-number)
}
.blob-num.non-expandable {
    cursor: default
}
.blob-num.non-expandable:hover {
    color: rgba(0, 0, 0, 0.3)
}
.blob-code {
    position: relative;
    padding-left: 10px;
    padding-right: 10px;
    vertical-align: top
}
.blob-code-inner {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px;
    color: #333;
    white-space: pre;
    overflow: visible;
    word-wrap: normal
}
.blob-code-inner .x-first {
    border-top-left-radius: 0.2em;
    border-bottom-left-radius: 0.2em
}
.blob-code-inner .x-last {
    border-top-right-radius: 0.2em;
    border-bottom-right-radius: 0.2em
}
.soft-wrap .diff-table {
    table-layout: fixed
}
.soft-wrap .blob-code {
    padding-left: 18px;
    text-indent: -7px
}
.soft-wrap .blob-code-inner {
    word-wrap: break-word;
    white-space: pre-wrap
}
.soft-wrap .no-nl-marker {
    display: none
}
.soft-wrap .add-line-comment {
    margin-left: -28px
}
.blob-num-hunk,
.blob-code-hunk,
.blob-num-expandable,
.blob-code-expandable {
    vertical-align: middle;
    color: rgba(0, 0, 0, 0.3);
    border-color: #d2dff0
}
.blob-num-hunk,
.blob-num-expandable {
    background-color: #edf2f9
}
.blob-code-hunk,
.blob-code-expandable {
    padding-top: 4px;
    padding-bottom: 4px;
    background-color: #f4f7fb;
    border-width: 1px 0
}
.blob-expanded .blob-num,
.blob-expanded .blob-code {
    background-color: #fafafa
}
.blob-expanded+tr:not(.blob-expanded) .blob-num,
.blob-expanded+tr:not(.blob-expanded) .blob-code {
    border-top: 1px solid #eee
}
.blob-expanded .blob-num-hunk {
    border-top: 1px solid #eee
}
tr:not(.blob-expanded)+.blob-expanded .blob-num,
tr:not(.blob-expanded)+.blob-expanded .blob-code {
    border-top: 1px solid #eee
}
.blob-num-expandable {
    padding: 0;
    font-size: 12px;
    text-align: center
}
.blob-num-expandable .diff-expander {
    display: block;
    width: auto;
    height: auto;
    margin-right: -1px;
    padding: 4px 11px 4px 10px;
    cursor: pointer;
    color: #767676
}
.blob-num-expandable .diff-expander:hover {
    color: #fff;
    text-shadow: none;
    background-color: #4078c0;
    border-color: #4078c0
}
.blob-code-addition {
    background-color: #eaffea
}
.blob-code-addition .x {
    background-color: #a6f3a6
}
.blob-num-addition {
    background-color: #dbffdb;
    border-color: #c1e9c1
}
.blob-code-deletion {
    background-color: #ffecec
}
.blob-code-deletion .x {
    background-color: #f8cbcb
}
.blob-num-deletion {
    background-color: #ffdddd;
    border-color: #f1c0c0
}
.selected-line.blob-code {
    background-color: #f8eec7
}
.selected-line.blob-code .x {
    background-color: transparent
}
.selected-line.blob-num {
    background-color: #f6e8b5;
    border-color: #f0db88
}
.add-line-comment {
    position: relative;
    z-index: 5;
    float: left;
    width: 20px;
    height: 20px;
    margin: -1px -10px -1px -20px;
    line-height: 20px;
    color: #fff;
    text-align: center;
    text-indent: 0;
    cursor: pointer;
    background-color: #4078c0;
    background-image: -webkit-linear-gradient(#5386c6, #4078c0);
    background-image: linear-gradient(#5386c6, #4078c0);
    border-radius: 3px;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.15);
    opacity: 0;
    -webkit-transform: scale(0.8, 0.8);
    -ms-transform: scale(0.8, 0.8);
    transform: scale(0.8, 0.8);
    -webkit-transition: -webkit-transform 0.1s ease-in-out;
    transition: transform 0.1s ease-in-out
}
.add-line-comment:hover {
    -webkit-transform: scale(1, 1);
    -ms-transform: scale(1, 1);
    transform: scale(1, 1)
}
.is-hovered .add-line-comment {
    opacity: 1
}
.add-line-comment.octicon-check {
    opacity: 1;
    background: #333
}
.inline-comment-form {
    border: 1px solid #ddd;
    border-radius: 3px
}
.inline-review-comment {
    margin-top: 0 !important;
    margin-bottom: 10px !important
}
.inline-review-comment .gc:first-child+tr .blob-num,
.inline-review-comment .gc:first-child+tr .blob-code {
    padding-top: 5px
}
.inline-review-comment tr:last-child {
    border-bottom-left-radius: 2px;
    border-bottom-right-radius: 2px
}
.inline-review-comment tr:last-child .blob-num,
.inline-review-comment tr:last-child .blob-code {
    padding-bottom: 8px
}
.inline-review-comment tr:last-child .blob-num:first-child,
.inline-review-comment tr:last-child .blob-code:first-child {
    border-bottom-left-radius: 2px
}
.inline-review-comment tr:last-child .blob-num:last-child,
.inline-review-comment tr:last-child .blob-code:last-child {
    border-bottom-right-radius: 2px
}
.timeline-inline-comments {
    width: 100%;
    table-layout: fixed
}
.timeline-inline-comments .inline-comments,
.show-inline-notes .inline-comments {
    display: table-row
}
.inline-comments {
    display: none
}
.inline-comments.is-collapsed {
    display: none
}
.inline-comments .line-comments.is-collapsed {
    visibility: hidden
}
.inline-comments .line-comments+.blob-num {
    border-left-width: 1px
}
.inline-comments .timeline-comment {
    margin-bottom: 10px
}
.inline-comments .inline-comment-form,
.inline-comments .inline-comment-form-container {
    max-width: 780px
}
.inline-comments .ajax-indicator {
    display: none;
    vertical-align: bottom
}
.inline-comments form.loading .ajax-indicator {
    display: inline-block
}
.inline-comments .comment-resolved-by {
    margin-left: 10px
}
.inline-comments .comment-resolved-by .username {
    font-weight: bold;
    color: #333
}
.comment-holder {
    max-width: 780px
}
.line-comments+.line-comments,
.empty-cell+.line-comments {
    border-left: 1px solid #eee
}
.inline-comment-form-container .inline-comment-form,
.inline-comment-form-container.open .inline-comment-form-actions {
    display: none
}
.inline-comment-form-container .inline-comment-form-actions,
.inline-comment-form-container.open .inline-comment-form {
    display: block
}
body.split-diff .header>.container,
body.split-diff .repohead>.container,
body.split-diff .site>.container,
body.split-diff .gist-content-wrapper>.container {
    width: 100%;
    padding-left: 30px;
    padding-right: 30px
}
body.split-diff .repository-with-sidebar {
    padding-right: 60px
}
body.split-diff .repository-with-sidebar .repository-sidebar {
    margin-right: -60px
}
body.split-diff .repository-content {
    width: 100%
}
body.split-diff .new-pr-form {
    max-width: 980px
}
body.split-diff .new-pr-form .discussion-sidebar {
    width: 200px
}
.file-diff-split {
    table-layout: fixed
}
.file-diff-split .blob-code+.blob-num {
    border-left-width: 1px
}
.file-diff-split .blob-code-inner {
    white-space: pre-wrap;
    word-wrap: break-word
}
.file-diff-split .empty-cell {
    cursor: default;
    background-color: #fafafa;
    border-right-color: #eee
}
.ghe-license-status {
    padding: 40px 0;
    font-size: 16px;
    text-align: center
}
.ghe-license-status .octocat {
    width: 225px;
    margin-bottom: 20px
}
.ghe-license-status h1 {
    margin-bottom: 10px
}
.ghe-license-status p {
    margin-top: 0;
    margin-bottom: 5px;
    color: #767676
}
.ghe-license-expiry-icon {
    margin: 5px 10px 0 0;
    color: #ddb38a
}
.enterprise .flash-global {
    max-height: 90px;
    overflow-y: scroll
}
.fakelogin {
    text-align: center;
    font-size: 14px;
    line-height: 34px;
    background-image: -webkit-linear-gradient(#dc5f59, #b33630);
    background-image: linear-gradient(#dc5f59, #b33630);
    border-bottom: 1px solid #900;
    color: #fff;
    position: fixed;
    top: 0;
    z-index: 1000;
    width: 100%;
    text-shadow: 0 -1px 0 rgba(153, 0, 0, 0.25)
}
.fakelogin+.header {
    margin-top: 35px
}
.fakelogin+#serverstats {
    margin-top: 35px
}
.fakelogin .cancel-impersonation {
    color: #fff;
    text-decoration: underline
}
.file {
    position: relative;
    margin-top: 20px;
    margin-bottom: 15px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.file .data.empty {
    padding: 5px 10px;
    color: #767676
}
.file .data.suppressed,
.file.open .image {
    display: none
}
.file.open .data.suppressed {
    display: block
}
.file .image {
    position: relative;
    padding: 30px;
    text-align: center;
    background-color: #ddd
}
.file .image table {
    margin: 0 auto
}
.file .image td {
    text-align: center;
    color: #888;
    padding: 0 5px;
    vertical-align: top
}
.file .image td img {
    max-width: 100%
}
.file .image .border-wrap {
    position: relative;
    display: inline-block;
    line-height: 0;
    background-color: #fff;
    border: 1px solid #767676
}
.file .image a {
    display: inline-block;
    line-height: 0
}
.file .image img,
.file .image canvas {
    max-width: 600px;
    background: url(https://github.com/images/modules/commit/trans_bg.gif) right bottom #eee;
    border: 1px solid #fff
}
.file .image .view img,
.file .image .view canvas {
    position: relative;
    top: 0;
    right: 0;
    background: url(https://github.com/images/modules/commit/trans_bg.gif) right bottom #eee;
    max-width: inherit
}
.file .image .view>span {
    vertical-align: middle
}
.file .hidden {
    display: none !important
}
.file .empty {
    background: none
}
.file-header {
    padding: 5px 10px;
    background-color: #f7f7f7;
    border-bottom: 1px solid #d8d8d8;
    border-top-left-radius: 2px;
    border-top-right-radius: 2px
}
.file-header:before {
    display: table;
    content: ""
}
.file-header:after {
    display: table;
    clear: both;
    content: ""
}
.file-actions {
    float: right;
    padding-top: 3px
}
.file-actions select {
    margin-left: 5px
}
.file-info {
    float: left;
    line-height: 32px;
    font-size: 12px;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.file-info-divider {
    display: inline-block;
    width: 1px;
    height: 18px;
    margin-right: 3px;
    margin-left: 3px;
    vertical-align: middle;
    background-color: #ddd
}
.file-mode {
    text-transform: capitalize
}
.show-file-notes {
    display: none
}
.has-inline-notes .show-file-notes {
    display: inline-block;
    margin-right: 10px;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
.file-blankslate {
    border: 0;
    border-radius: 0 0 2px 2px
}
.octicon-btn {
    display: inline-block;
    margin-left: 5px;
    padding: 5px;
    line-height: 1;
    color: #767676;
    vertical-align: middle;
    background: transparent;
    border: 0;
    outline: none
}
.octicon-btn:hover {
    color: #4078c0
}
.octicon-btn.disabled {
    color: #bbb;
    cursor: default
}
.octicon-btn.disabled:hover {
    color: #bbb
}
.octicon-btn-danger:hover {
    color: #bd2c00
}
.enable-fullscreen.btn-sm {
    display: inline-block;
    margin-top: 4px;
    margin-left: 5px;
    padding: 0 6px
}
.enable-fullscreen.btn-sm .octicon {
    margin-right: 0
}
.new-file .enable-fullscreen {
    margin-left: 11px
}
.write-content {
    position: relative
}
.write-content .enable-fullscreen {
    position: absolute;
    top: 5px;
    right: 16px;
    color: #333;
    opacity: 0.5;
    line-height: 1em
}
.write-content .enable-fullscreen:hover {
    opacity: 1
}
.fullscreen-overlay {
    display: none;
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: #fff;
    text-shadow: none;
    z-index: 1000
}
.fullscreen-overlay .fullscreen-container {
    max-width: 800px;
    height: 100%;
    margin: 0 auto;
    padding: 30px 0
}
.fullscreen-overlay .textarea-wrap {
    width: 100%;
    height: 100%;
    position: relative
}
.fullscreen-overlay textarea {
    width: 100%;
    height: 100%;
    padding: 20px;
    border: 0;
    background: #fff;
    color: darkgray;
    font-size: 21px;
    line-height: 1.6em;
    resize: none;
    -webkit-transition: color 0.15s ease-in-out;
    transition: color 0.15s ease-in-out;
    box-shadow: none
}
.fullscreen-overlay textarea:focus,
.fullscreen-overlay textarea:hover {
    outline: none;
    color: #333
}
.fullscreen-overlay .fullscreen-sidebar {
    position: absolute;
    top: 30px;
    right: 30px;
    text-align: right;
    z-index: 1002
}
.fullscreen-overlay .fullscreen-sidebar .exit-fullscreen,
.fullscreen-overlay .fullscreen-sidebar .theme-switcher {
    color: #c3c3c3;
    float: right;
    clear: right;
    margin-bottom: 15px
}
.fullscreen-overlay .fullscreen-sidebar .exit-fullscreen:hover,
.fullscreen-overlay .fullscreen-sidebar .theme-switcher:hover {
    color: #333;
    text-shadow: 0 0 10px #fff
}
.fullscreen-overlay .fullscreen-sidebar .theme-switcher {
    margin-right: 8px
}
.fullscreen-overlay.dark-theme {
    background: #1d1f21
}
.fullscreen-overlay.dark-theme textarea {
    background: #1d1f21;
    color: #a4b1b1
}
.fullscreen-overlay.dark-theme textarea:focus,
.fullscreen-overlay.dark-theme textarea:hover {
    color: #dbe0e0
}
.fullscreen-overlay.dark-theme .fullscreen-sidebar {
    color: #dbe0e0
}
.fullscreen-overlay.dark-theme .fullscreen-sidebar .exit-fullscreen,
.fullscreen-overlay.dark-theme .fullscreen-sidebar .theme-switcher {
    color: #a4b1b1
}
.fullscreen-overlay.dark-theme .fullscreen-sidebar .exit-fullscreen:hover,
.fullscreen-overlay.dark-theme .fullscreen-sidebar .theme-switcher:hover {
    color: #dbe0e0;
    text-shadow: 0 0 10px #000
}
.fullscreen-overlay .suggester-container {
    top: 5px;
    left: 0
}
.fullscreen-overlay-enabled .wrapper,
.fullscreen-overlay-enabled #footer {
    display: none
}
.fullscreen-overlay-enabled .fullscreen-overlay {
    display: block
}
.previewable-comment-form .upload-enabled .fullscreen-overlay textarea {
    max-height: none;
    border: 0;
    border-radius: 0
}
#gollum-editor {
    margin: 10px 0 50px;
    padding: 10px 0 0;
    border: 0
}
#gollum-editor .comment-form-head.tabnav {
    border: 1px solid #ddd
}
#gollum-editor #gollum-editor-body {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    line-height: 22px;
    margin: 13px 0 5px;
    height: 390px;
    resize: vertical
}
#gollum-editor #gollum-editor-body+.collapsed,
#gollum-editor #gollum-editor-body+.expanded {
    border-top: 1px solid #ddd;
    margin-top: 7px
}
#gollum-editor .collapsed,
#gollum-editor .expanded {
    border-bottom: 1px solid #ddd;
    display: block;
    overflow: hidden;
    padding: 10px 0 5px
}
#gollum-editor .collapsed a.button,
#gollum-editor .expanded a.button {
    border: 1px solid #ddd;
    color: #333;
    display: block;
    float: left;
    height: 25px;
    overflow: hidden;
    margin: 2px 5px 7px 0;
    padding: 0;
    text-shadow: 0 1px 0 #fff;
    width: 25px;
    background-image: -webkit-linear-gradient(#fafafa, #eaeaea);
    background-image: linear-gradient(#fafafa, #eaeaea);
    border-radius: 3px
}
#gollum-editor .collapsed a.button:hover,
#gollum-editor .expanded a.button:hover {
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.3);
    text-decoration: none;
    background-image: -webkit-linear-gradient(#599bdc, #3072b3);
    background-image: linear-gradient(#599bdc, #3072b3)
}
#gollum-editor .collapsed a.button span,
#gollum-editor .expanded a.button span {
    margin: 4px
}
#gollum-editor .collapsed h4,
#gollum-editor .expanded h4 {
    font-size: 16px;
    float: left;
    margin: 0;
    padding: 6px 0 0 4px;
    text-shadow: 0 -1px 0 #fff
}
#gollum-editor .collapsed a.button span.octicon-triangle-right {
    display: inline-block
}
#gollum-editor .collapsed textarea,
#gollum-editor .collapsed a.button span.octicon-triangle-down {
    display: none
}
#gollum-editor .expanded a.button span.octicon-triangle-down {
    display: inline-block
}
#gollum-editor .expanded a.button span.octicon-triangle-right {
    display: none
}
#gollum-editor .expanded textarea {
    border: 1px solid #ddd;
    clear: both;
    display: block;
    font-size: 12px;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    height: 84px;
    margin: 8px 0;
    padding: 6px;
    width: 883px;
    resize: vertical
}
#gollum-editor a.gollum-minibutton,
#gollum-editor a.gollum-minibutton:visited {
    border: 1px solid #d4d4d4;
    color: #333;
    cursor: pointer;
    display: block;
    font-size: 12px;
    font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
    font-weight: bold;
    margin: 0 0 0 9px;
    padding: 5px 12px;
    text-shadow: 0 1px 0 #fff;
    background-image: -webkit-linear-gradient();
    background-image: linear-gradient();
    border-radius: 3px
}
#gollum-editor a.gollum-minibutton:hover,
#gollum-editor a.gollum-minibutton:visited:hover {
    border-color: #518cc6 #518cc6 #2a65a0;
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.3);
    text-decoration: none;
    background-image: -webkit-linear-gradient(#599bdc, #3072b3);
    background-image: linear-gradient(#599bdc, #3072b3)
}
.singleline {
    display: block;
    margin: 20px 0
}
.singleline label {
    display: block;
    margin-bottom: 6px
}
#gollum-editor-title-field {
    border-bottom: 0;
    margin: 0 0 10px
}
.gollum-editor-page-title {
    font-weight: bold;
    margin-top: 0
}
.gollum-editor-page-title.ph {
    color: #000
}
#gollum-editor-function-bar {
    height: 26px;
    padding-bottom: 10px;
    border-bottom: 1px solid #ddd;
    margin: 10px 0;
    border: 0
}
#gollum-editor-function-bar #gollum-editor-function-buttons {
    display: none;
    float: left
}
#gollum-editor-function-bar.active #gollum-editor-function-buttons {
    display: block
}
#gollum-editor-function-bar #gollum-editor-format-selector {
    padding-top: 5px;
    float: left;
    margin-left: 20px
}
#gollum-editor-function-bar #gollum-editor-format-selector select {
    margin: 0
}
#gollum-editor-function-bar #gollum-editor-format-selector label {
    color: #767676;
    font-size: 11px;
    font-weight: bold;
    line-height: 17px;
    padding: 0 5px 0 0
}
#gollum-editor-function-buttons .btn-sm {
    width: 30px;
    padding-left: 0;
    padding-right: 0;
    text-align: center
}
#gollum-editor-function-buttons .btn-sm .octicon {
    margin-right: 0
}
#gollum-error-message {
    display: none;
    padding-top: 12px;
    font-size: 1.8em;
    color: #f33
}
#gollum-editor-help {
    overflow: hidden;
    padding: 0;
    border: 1px solid #ddd;
    border-radius: 3px
}
#gollum-editor-help-parent,
#gollum-editor-help-list {
    display: block;
    float: left;
    height: 170px;
    list-style-type: none;
    overflow: auto;
    margin: 0;
    padding: 10px 0;
    width: 160px;
    border-right: 1px solid #eee
}
#gollum-editor-help-parent li,
#gollum-editor-help-list li {
    font-size: 12px;
    line-height: 1.6;
    margin: 0;
    padding: 0
}
#gollum-editor-help-parent li a,
#gollum-editor-help-list li a {
    border: 1px solid transparent;
    border-width: 1px 0;
    display: block;
    font-weight: bold;
    padding: 2px 12px;
    text-shadow: 0 -1px 0 #fff
}
#gollum-editor-help-parent li a:hover,
#gollum-editor-help-list li a:hover {
    background: #fff;
    border-color: #f0f0f0;
    text-decoration: none;
    box-shadow: none
}
#gollum-editor-help-parent li a.selected,
#gollum-editor-help-list li a.selected {
    border: 1px solid #eee;
    border-bottom-color: #e7e7e7;
    border-width: 1px 0;
    background: #fff;
    color: #000;
    box-shadow: 0 1px 2px #f0f0f0
}
#gollum-editor-help-list {
    background: #fafafa
}
#gollum-editor-help-wrapper {
    background: #fff;
    overflow: auto;
    height: 170px;
    padding: 10px
}
#gollum-editor-help-content {
    font-size: 12px;
    margin: 0 10px 0 5px;
    padding: 0;
    line-height: 1.8
}
#gollum-editor-help-content p {
    margin: 0 0 10px;
    padding: 0
}
.ie #gollum-editor .singleline input {
    padding-top: 0.25em;
    padding-bottom: 0.75em
}
#gollum-footer {
    font-size: 12px;
    line-height: 19px
}
#gollum-dialog-dialog h4 {
    border-bottom: 1px solid #ddd;
    color: #333;
    font-size: 16px;
    line-height: normal;
    font-weight: bold;
    margin: 0 0 12px;
    padding: 0 0 6px;
    text-shadow: 0 -1px 0 #f7f7f7
}
#gollum-dialog-dialog-body {
    font-size: 12px;
    line-height: 16px;
    margin: 0;
    padding: 0
}
#gollum-dialog-dialog-body fieldset {
    display: block;
    border: 0;
    margin: 0;
    overflow: hidden;
    padding: 0 12px
}
#gollum-dialog-dialog-body fieldset .field {
    margin: 0 0 18px;
    padding: 0
}
#gollum-dialog-dialog-body fieldset .field:last-child {
    margin: 0 0 12px
}
#gollum-dialog-dialog-body fieldset label {
    color: #666;
    display: block;
    font-size: 14px;
    font-weight: bold;
    line-height: 1.6;
    margin: 0;
    padding: 0;
    min-width: 80px
}
#gollum-dialog-dialog-body fieldset input[type="text"] {
    display: block;
    margin: 3px 0 0;
    width: 100%
}
#gollum-dialog-dialog-body fieldset input.code {
    font-family: 'Monaco', 'Courier New', Courier, monospace
}
#gollum-dialog-dialog-buttons {
    border-top: 1px solid #ddd;
    overflow: hidden;
    margin: 14px 0 0;
    padding: 12px 0 0
}
a.gollum-minibutton,
a.gollum-minibutton:visited {
    border: 1px solid #d4d4d4;
    color: #333;
    cursor: pointer;
    display: inline;
    font-size: 12px;
    font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
    font-weight: bold;
    float: right;
    width: auto;
    margin: 0 0 0 9px;
    padding: 4px 12px;
    text-shadow: 0 1px 0 #fff;
    background-image: -webkit-linear-gradient();
    background-image: linear-gradient();
    border-radius: 3px
}
a.gollum-minibutton:hover,
a.gollum-minibutton:visited:hover {
    border-color: #518cc6 #518cc6 #2a65a0;
    color: #fff;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.3);
    text-decoration: none;
    background-image: -webkit-linear-gradient(#599bdc, #3072b3);
    background-image: linear-gradient(#599bdc, #3072b3)
}
.wiki-wrapper .ie #gollum-editor {
    padding-bottom: 1em
}
.wiki-wrapper #wiki-content .enable-fullscreen {
    right: 4px
}
.wiki-wrapper #wiki-content .previewable-comment-form.write-selected .write-content,
.wiki-wrapper #wiki-content .previewable-comment-form.preview-selected .preview-content {
    padding: 0;
    margin: 0
}
.wiki-wrapper #wiki-content .comment-body {
    padding: 5px 0 20px
}
.wiki-wrapper hr {
    margin: 25px 0 20px
}
.wiki-wrapper.comment-body {
    width: 920px
}
.hooks-listing .boxed-group-action.select-menu {
    z-index: auto
}
.hook-item a:hover {
    text-decoration: none
}
.hook-item .item-status {
    float: left;
    margin-right: 8px;
    width: 16px;
    text-align: center
}
.hook-item .description {
    color: #999
}
.hook-item .description .css-truncate-target {
    max-width: 160px
}
.hook-item .icon-for-success,
.hook-item .icon-for-failure,
.hook-item .icon-for-pending,
.hook-item .icon-for-inactive {
    display: none
}
.hook-item.success .icon-for-success {
    display: inline-block;
    color: #6cc644
}
.hook-item.failure .icon-for-failure {
    display: inline-block;
    color: #bd2c00
}
.hook-item.pending .icon-for-pending {
    display: inline-block;
    color: #999
}
.hook-item.inactive .icon-for-inactive {
    display: inline-block;
    color: #999
}
.hook-url.css-truncate-target {
    max-width: 360px
}
.hook-events-field .hook-event-selector {
    display: none
}
.hook-events-field.is-custom .hook-event-selector {
    display: block
}
.hook-event-selector {
    margin-left: 10px
}
.hook-event {
    display: inline-block;
    width: 310px;
    margin: 0;
    padding: 5px 0 5px 30px
}
.hook-event .note {
    font-size: 11px;
    margin: 0;
    color: #aaa
}
.hook-event-choice {
    font-weight: normal
}
.hook-form.is-ssl .ssl-hook-fields {
    display: block
}
.hook-form .ssl-hook-fields {
    display: none
}
.hook-form .ssl-hook-fields #disable-ssl-verification-modal,
.hook-form .ssl-hook-fields .enable-ssl-verification {
    display: none
}
.hook-form .ssl-hook-fields.is-not-verifying-ssl .enable-ssl-verification {
    display: block
}
.hook-form .ssl-hook-fields.is-not-verifying-ssl .disable-ssl-verification {
    display: none
}
.hook-form .disable-ssl-verification .actions {
    margin-top: -4px
}
.hook-form .invalid-url-notice {
    display: none;
    padding: 7px 4px
}
.hook-form .invalid-url-notice .octicon-alert {
    position: relative;
    top: 1px
}
.hook-form.is-invalid-url .invalid-url-notice {
    display: block
}
.hooks-oap-warning {
    margin-top: 0
}
.hooks-oap-warning ul {
    margin: 10px 0
}
.hooks-oap-warning ul li {
    margin-left: 16px
}
.hook-secret .hook-secret-standin {
    display: block
}
.hook-secret .hook-secret-field {
    display: none
}
.hook-secret.open .hook-secret-standin {
    display: none
}
.hook-secret.open .hook-secret-field {
    display: block
}
.hook-deliveries-list .loading-message {
    display: block
}
.hook-deliveries-list .error-message {
    display: none
}
.hook-deliveries-list.is-error .loading-message {
    display: none
}
.hook-deliveries-list.is-error .error-message {
    display: block
}
.hook-deliveries-list .spinner {
    display: inline-block;
    vertical-align: top;
    margin: 0
}
.hook-deliveries-list .hook-delivery-item:hover {
    background-color: transparent
}
.hook-deliveries-list .item-status {
    display: inline-block;
    margin-right: 5px;
    width: 16px;
    text-align: center
}
.hook-deliveries-list .item-status .icon-for-success,
.hook-deliveries-list .item-status .icon-for-failure,
.hook-deliveries-list .item-status .icon-for-pending {
    display: none
}
.hook-deliveries-list .item-status.success {
    visibility: visible;
    color: #6cc644
}
.hook-deliveries-list .item-status.success .icon-for-success {
    display: inline-block
}
.hook-deliveries-list .item-status.failure {
    color: #bd2c00
}
.hook-deliveries-list .item-status.failure .icon-for-failure {
    display: inline-block
}
.hook-deliveries-list .item-status.pending {
    color: #999
}
.hook-deliveries-list .item-status.pending .icon-for-pending {
    display: inline-block
}
.hook-deliveries-pagination-loading-message {
    display: none
}
.hook-deliveries-pagination-loading-message .animated-ellipsis-container {
    text-align: left
}
.hook-deliveries-pagination.loading .hook-deliveries-pagination-button {
    display: none
}
.hook-deliveries-pagination.loading .hook-deliveries-pagination-loading-message {
    display: block
}
.boxed-group-list li.hook-delivery-item {
    padding: 10px
}
.hook-delivery-item .hook-delivery-details {
    display: none
}
.hook-delivery-item .hook-delivery-details .loading-message,
.hook-delivery-item .hook-delivery-details .error-message {
    display: none
}
.hook-delivery-item .hook-delivery-details.is-loading .loading-message {
    display: block
}
.hook-delivery-item .hook-delivery-details.has-error .error-message {
    display: block
}
.hook-delivery-item.open .hook-delivery-details {
    display: block
}
.hook-delivery-item .loading-message {
    text-align: center
}
.hook-delivery-time {
    float: right;
    margin-right: 10px;
    color: #999;
    font-size: 10px
}
a.hook-delivery-summary {
    text-decoration: none
}
.hook-delivery-guid {
    display: inline-block;
    padding: 2px 6px;
    color: rgba(0, 0, 0, 0.5);
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px;
    background-color: rgba(209, 227, 237, 0.5);
    border-radius: 3px
}
.hook-delivery-guid .octicon {
    margin: 1px -2px 0 0;
    color: #b0c4ce
}
.hook-delivery-actions {
    padding-top: 1px
}
.boxed-group-list>li.hook-delivery-item .btn-sm {
    margin: 0
}
.boxed-group-list>li.hook-delivery-item .hook-delivery-details .redeliver.btn-sm {
    margin: 5px 0 0
}
.hook-deliveries-list .error-message,
.hook-delivery-details .error-message {
    margin: 10px 0;
    padding: 7px 4px
}
.hook-deliveries-list .error-message .octicon,
.hook-delivery-details .error-message .octicon {
    position: relative;
    top: 1px
}
.boxed-group span.animated-ellipsis-container,
.boxed-group span.animated-ellipsis {
    padding: 0
}
.boxed-group .animated-ellipsis-container {
    line-height: 1.3
}
.hook-delivery-details {
    clear: right
}
.hook-delivery-details .error-message {
    margin-bottom: 0
}
.hook-delivery-details .tabnav-tabcontent {
    display: none
}
.hook-delivery-details .tabnav-tabcontent.selected {
    display: block
}
.hook-delivery-details hr {
    margin: 10px 0
}
.hook-delivery-details pre {
    padding: 7px 12px;
    margin: 10px 0;
    background-color: #f8f8f8;
    border: 1px solid #ddd;
    border-radius: 3px;
    font-size: 13px;
    line-height: 1.5;
    overflow: auto
}
.hook-delivery-details .tabnav {
    margin: 10px 0
}
.hook-delivery-details h4.remote-call-header {
    border-bottom: 1px solid #999;
    margin: 20px 0 10px
}
.hook-delivery-response-status {
    display: inline-block;
    padding: 4px 6px 3px;
    color: #fff;
    background-color: #bd2c00;
    border-radius: 3px;
    font-size: 10px;
    line-height: 1.1;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.hook-delivery-response-status[data-response-status^="2"] {
    background-color: #6cc644
}
.redelivery-dialog .pending-message {
    display: block
}
.redelivery-dialog .failure-message {
    display: none
}
.redelivery-dialog.failed {
    color: #9c2400;
    background-image: -webkit-linear-gradient(#f8d8d8, #efd0d0);
    background-image: linear-gradient(#f8d8d8, #efd0d0);
    border-color: #da9797
}
.redelivery-dialog.failed .pending-message {
    display: none
}
.redelivery-dialog.failed .failure-message {
    display: block
}
.redelivering-hook-delivery .error-message {
    display: none
}
.redelivering-hook-delivery.is-error .loading-message {
    display: none
}
.redelivering-hook-delivery.is-error .error-message {
    display: block
}
.test-hook-message .success-message,
.test-hook-message .error-message {
    display: none;
    margin-top: 10px
}
.test-hook-message.success .success-message {
    display: block
}
.test-hook-message.error .error-message {
    display: block
}
.issues-reset-query-wrapper {
    margin-bottom: 20px
}
.issues-reset-query {
    font-weight: bold;
    color: #767676
}
.issues-reset-query .octicon-x {
    width: 20px;
    height: 20px;
    margin-right: 3px;
    line-height: 20px;
    color: #fff;
    text-align: center;
    background-color: #767676;
    border-radius: 3px
}
.issues-reset-query:hover {
    color: #4078c0;
    text-decoration: none
}
.issues-reset-query:hover .octicon-x {
    background-color: #4078c0
}
.table-list-milestones .table-list-cell {
    padding: 15px 20px
}
.table-list-milestones .stat {
    display: inline-block;
    font-size: 14px;
    font-weight: bold;
    line-height: 1.2;
    color: #555;
    white-space: nowrap
}
.table-list-milestones .stat+.stat {
    margin-left: 15px
}
.table-list-milestones .stat-label {
    font-weight: normal;
    color: #767676
}
.milestone-title {
    width: 500px
}
.milestone-title-link {
    margin-top: 0;
    margin-bottom: 5px;
    font-size: 24px;
    font-weight: normal;
    line-height: 1.2
}
.milestone-title-link a {
    color: #333
}
.milestone-title-link a:hover {
    color: #4078c0
}
.milestone-progress {
    width: 420px
}
.milestone-progress .progress-bar {
    margin-top: 7px;
    margin-bottom: 12px
}
.milestone-meta {
    font-size: 14px
}
.milestone-meta-item {
    display: inline-block;
    margin-right: 10px
}
.milestone-meta-item .octicon {
    width: 16px;
    text-align: center
}
.milestone-description-html {
    display: none
}
.milestone-description {
    margin-top: 5px
}
.milestone-description .expand-more {
    color: #4078c0;
    cursor: pointer
}
.milestone-description .expand-more:hover {
    text-decoration: underline
}
.milestone-description.open .milestone-description-plaintext {
    display: none
}
.milestone-description.open .milestone-description-html {
    display: block
}
.milestone-actions {
    margin-top: 8px;
    font-size: 13px
}
.milestone-action {
    display: inline-block;
    margin-right: 10px
}
.milestone-calender-container {
    margin-left: 30px
}
.task-progress {
    color: #767676;
    text-decoration: none
}
.task-progress .octicon {
    margin-right: 5px;
    color: #ccc;
    vertical-align: middle
}
.task-progress .progress-bar {
    display: inline-block;
    width: 120px;
    height: 5px;
    vertical-align: 2px;
    background-color: #eee
}
.task-progress .progress-bar .progress {
    background-color: #ccc
}
.task-progress-counts {
    display: inline-block;
    margin-right: 6px;
    margin-left: -2px;
    font-size: 12px
}
a.task-progress:hover {
    color: #4078c0
}
a.task-progress:hover .octicon {
    color: inherit
}
a.task-progress:hover .progress-bar .progress {
    background-color: #4078c0
}
.issues-listing {
    position: relative
}
.issues-listing .octocat-search {
    position: absolute;
    right: 0;
    height: 250px;
    margin: -132px -4px;
    -webkit-transform: scaleX(-1);
    -ms-transform: scaleX(-1);
    transform: scaleX(-1)
}
.issues-listing .table-list-issues .issue-title {
    width: 740px;
    padding-top: 12px
}
.issues-listing .table-list-issues .labels {
    display: inline-block;
    margin-bottom: 2px;
    vertical-align: 1px
}
.issues-listing .table-list-issues .table-list-cell-avatar {
    padding-top: 13px
}
.issues-listing .table-list-issues .table-list-cell-avatar a {
    display: inline-block
}
.issues-listing .table-list-issues .table-list-cell-avatar img {
    border-radius: 2px
}
.issues-listing .table-list-issues .navigation-focus {
    background-color: #f5f5f5
}
.issue-title-link {
    padding-right: 3px;
    margin-bottom: 2px;
    font-size: 15px;
    font-weight: bold;
    line-height: 1.2;
    color: #333
}
.issue-title-link:hover {
    color: #4078c0;
    text-decoration: none
}
.issue-title-link:hover .num {
    color: inherit
}
.issue-nwo-link {
    color: #767676
}
.issue-pr-status {
    display: inline-block;
    margin-right: 3px;
    vertical-align: -1px
}
.issue-meta {
    margin-top: 1px;
    margin-bottom: 2px;
    font-weight: normal;
    color: #767676
}
.issue-meta-section {
    margin-right: 10px
}
.issue-milestone {
    max-width: 240px
}
.issue-comments {
    width: 60px;
    padding-top: 13px;
    font-weight: bold;
    text-align: right;
    white-space: nowrap
}
.issue-comments .octicon {
    margin-right: 1px
}
.issue-comments-no-comment {
    color: #eee
}
.milestone-link {
    color: #767676
}
.milestone-link .octicon {
    font-size: 14px;
    color: #ccc
}
.milestone-link:hover {
    color: #4078c0;
    text-decoration: none
}
.milestone-link:hover .octicon {
    color: inherit
}
.issue-review-status {
    float: right;
    margin-top: 1px
}
.new-issue-form {
    padding-top: 20px;
    margin-top: 20px;
    border-top: 1px solid #ddd
}
.new-issue-form:before {
    display: table;
    content: ""
}
.new-issue-form:after {
    display: table;
    clear: both;
    content: ""
}
.new-issue-form .discussion-timeline:before {
    display: none
}
.new-pr-form {
    margin-bottom: 20px
}
.new-pr-form:before {
    display: table;
    content: ""
}
.new-pr-form:after {
    display: table;
    clear: both;
    content: ""
}
.new-pr-form .discussion-timeline:before {
    display: none
}
.new-pr-form .discussion-sidebar {
    position: static
}
.new-pr-form .form-actions {
    padding-left: 0;
    padding-right: 0;
    text-align: left
}
.new-pr-form .form-actions .btn {
    float: none;
    margin-left: 0
}
.label-select-menu .color {
    display: inline-block;
    width: 14px;
    height: 14px;
    margin-right: 2px;
    margin-bottom: 2px;
    vertical-align: middle;
    border-radius: 3px
}
.label-select-menu .selected .select-menu-item-icon {
    color: inherit !important
}
.label-select-menu .selected:active {
    background-color: transparent !important
}
.label-select-menu .select-menu-item.navigation-focus {
    color: inherit;
    background-color: #f4f4f4
}
.label-select-menu .select-menu-item.navigation-focus .select-menu-item-icon {
    color: transparent
}
.label-select-menu .select-menu-item .octicon-x {
    display: none;
    float: right;
    margin: 1px 10px 0 0;
    opacity: 0.6
}
.label-select-menu .select-menu-item.selected .octicon-x {
    display: block;
    color: inherit
}
.label-select-menu>form {
    position: relative
}
.closed-banner {
    height: 7px;
    margin: 15px 0 15px 60px;
    overflow: hidden;
    background: url(https://github.com/images/modules/comments/closed_pattern.gif);
    border-radius: 3px
}
.subnav .btn+.issues-search {
    padding-right: 10px;
    border-right: 1px solid #eee
}
.merge-branch-heading {
    margin: 0
}
.merge-branch-description {
    margin-top: 0;
    margin-right: 160px;
    margin-bottom: -5px;
    line-height: 1.6em;
    color: #767676
}
.merge-branch-description .zeroclipboard-link .octicon {
    top: 2px
}
.alt-merge-options {
    display: inline-block;
    margin-top: 0;
    margin-bottom: 0;
    margin-left: 4px;
    vertical-align: middle
}
.merged .merge-branch-description .commit-ref .css-truncate-target {
    max-width: 180px
}
.merge-branch-prh-output {
    margin-top: 10px
}
.cancel-request-form {
    display: inline
}
.merge-branch-form {
    padding-left: 64px;
    display: none
}
.merge-branch.open .merge-branch-form {
    display: block
}
.merge-branch.open .merge-message {
    display: none
}
.merge-branch-manually {
    display: none;
    padding-top: 15px;
    margin-top: 14px;
    background-color: transparent;
    border-top: 1px solid #ddd
}
.merge-branch-manually h3,
.merge-branch-manually p {
    margin: 0
}
.merge-branch-manually h3 {
    margin-bottom: 10px
}
.merge-branch-manually .intro {
    padding-bottom: 10px;
    margin-top: 0
}
.merge-branch-manually .step {
    margin: 15px 0 5px
}
.merge-branch-manually .url-box {
    padding: 0;
    margin-left: 0;
    border: 0
}
.merge-branch-manually .clone-urls {
    width: 100%
}
.merge-branch-manually .copyable-terminal {
    background-color: #f2f2f2
}
.open>.merge-branch-manually {
    display: block
}
#network .network-tree {
    vertical-align: middle
}
#network .gravatar {
    margin-right: 4px;
    border-radius: 3px;
    vertical-align: middle
}
#network .octicon {
    margin-left: 2px;
    vertical-align: middle;
    width: 16px;
    display: inline-block;
    text-align: center
}
#network .current-repository {
    background-color: #fff6a9
}
#network .network-graph-container {
    position: relative;
    margin-bottom: 20px;
    border: 1px solid #ddd;
    border-radius: 3px;
    overflow: hidden;
    line-height: 0
}
#network .network-graph-container .large-loading-area {
    position: absolute;
    top: 0;
    right: 0;
    left: 0
}
.page-new-repo .octicon-repo {
    color: #bbb
}
.page-new-repo .octicon-lock {
    color: #e9dba5
}
.page-new-repo ul.repo-templates {
    margin: 10px 0
}
.page-new-repo ul.repo-templates>li {
    list-style-type: none;
    display: inline-block;
    margin: 0 10px 0 0
}
.page-new-repo ul.repo-templates .select-menu {
    float: left
}
.page-new-repo .team-select {
    display: none
}
.page-new-repo .form-checkbox .mega-octicon {
    font-size: 24px;
    float: left;
    margin-right: 5px
}
.page-new-repo .license-info {
    float: left;
    margin-top: 5px;
    margin-left: 10px;
    color: #ccc
}
.new-repo-container {
    width: 700px;
    margin: 0 auto;
    padding-top: 20px
}
.new-repo-container h2 {
    font-size: 22px;
    font-weight: normal;
    color: #666;
    border-bottom: 1px solid #ddd;
    padding-bottom: 5px;
    margin-bottom: 0.5em
}
.owner-reponame dl.form {
    margin-top: 5px;
    margin-bottom: 0
}
.owner-reponame dl.form.warn dd.warning,
.owner-reponame dl.form.warn dd.error,
.owner-reponame dl.form.errored dd.warning,
.owner-reponame dl.form.errored dd.error {
    display: inline-block;
    position: absolute;
    z-index: 10;
    margin: 2px 0 0;
    padding: 5px 8px;
    font-size: 13px;
    font-weight: normal;
    border-radius: 3px
}
.owner-reponame dl.form.warn dd.warning:after,
.owner-reponame dl.form.warn dd.warning:before,
.owner-reponame dl.form.warn dd.error:after,
.owner-reponame dl.form.warn dd.error:before,
.owner-reponame dl.form.errored dd.warning:after,
.owner-reponame dl.form.errored dd.warning:before,
.owner-reponame dl.form.errored dd.error:after,
.owner-reponame dl.form.errored dd.error:before {
    bottom: 100%;
    z-index: 15;
    left: 15px;
    border: solid transparent;
    content: " ";
    height: 0;
    width: 0;
    position: absolute;
    pointer-events: none
}
.owner-reponame dl.form.warn dd.warning:after,
.owner-reponame dl.form.warn dd.error:after,
.owner-reponame dl.form.errored dd.warning:after,
.owner-reponame dl.form.errored dd.error:after {
    border-width: 5px;
    margin-left: -5px
}
.owner-reponame dl.form.warn dd.warning:before,
.owner-reponame dl.form.warn dd.error:before,
.owner-reponame dl.form.errored dd.warning:before,
.owner-reponame dl.form.errored dd.error:before {
    border-width: 6px;
    margin-left: -6px
}
.owner-reponame dl.form.warn dd.warning {
    color: #4e401e;
    background-color: #ffe5a7;
    border: 1px solid #e7ce94
}
.owner-reponame dl.form.warn dd.warning:after {
    border-color: rgba(255, 229, 167, 0);
    border-bottom-color: #ffe5a7
}
.owner-reponame dl.form.warn dd.warning:before {
    border-color: rgba(205, 186, 131, 0);
    border-bottom-color: #cdb683
}
.owner-reponame dl.form.errored dd.error {
    color: #fff;
    background-color: #bf1515;
    border-color: #911;
    font-size: 13px
}
.owner-reponame dl.form.errored dd.error:after {
    border-color: rgba(191, 21, 21, 0);
    border-bottom-color: #bf1515
}
.owner-reponame dl.form.errored dd.error:before {
    border-color: rgba(153, 17, 17, 0);
    border-bottom-color: #911
}
.owner-reponame .slash {
    float: left;
    font-size: 21px;
    color: #666;
    padding-top: 32px;
    margin: 0 8px
}
.owner-reponame .icon-preview {
    display: none;
    font-size: 32px;
    position: absolute;
    text-align: right;
    width: 100px;
    top: 23px;
    left: -115px
}
.owner-reponame .icon-preview.icon-preview-public {
    top: 25px
}
.reponame-suggestion {
    color: #34631a;
    cursor: pointer
}
.upgrade-upsell {
    padding-left: 33px
}
.cc-upgrade {
    padding-left: 20px
}
.featured-license {
    font-weight: bold
}
.license-container {
    border-left: 1px solid #ccc;
    padding-left: 15px
}
.notification-routing .notification-email .edit-link {
    margin-right: 10px;
    font-weight: bold
}
.notification-routing .notification-email .btn-sm {
    float: none;
    margin: -2px 0 0
}
.notification-routing .notification-email .edit-form {
    display: none
}
.notification-routing .notification-email.open .edit-form {
    display: block
}
.notification-routing .notification-email.open .email-display {
    display: none
}
.notifications .list-group-item {
    padding-top: 8px;
    padding-bottom: 8px;
    padding-left: 35px;
    border-width: 1px 0 0
}
.notifications .list-group-item:first-child {
    border: 0
}
.notifications .list-group-item-name {
    display: block;
    max-width: 400px;
    font-size: 14px;
    line-height: 1.5em
}
.notifications .list-group-item-name a {
    display: block;
    max-width: 460px
}
.notifications .notifications-more {
    padding: 0
}
.notifications .notifications-more>a {
    display: block;
    padding: 10px 15px;
    font-weight: bold;
    color: #4078c0;
    text-align: center
}
.notifications .notifications-more>a:hover {
    text-decoration: underline
}
.notifications .read .type-icon {
    color: #767676
}
.notifications .read .list-group-item-name>a {
    color: #767676
}
.notifications .read .notification-actions {
    color: #767676
}
.notifications .read .avatar-stack {
    opacity: 0.5
}
.notifications .read .undo {
    display: block
}
.notifications .read .delete {
    visibility: hidden
}
.notifications .read.navigation-focus {
    background-color: #f5f9fc
}
.notifications .muted .unmute {
    display: block
}
.notifications .muted .mute {
    display: none
}
.notifications .unmute {
    display: none
}
.type-icon-state-none {
    color: #767676
}
.type-icon-state-open {
    color: #6cc644
}
.type-icon-state-closed {
    color: #bd2c00
}
.type-icon-state-merged {
    color: #6e5494
}
.notifications-list {
    float: left;
    width: 100%
}
.notifications-list .notifications-repo-link {
    max-width: 500px
}
.notifications-list .boxed-group .text-success {
    position: absolute;
    color: #6cc644;
    right: 3px;
    width: 210px;
    margin-top: 4px;
    text-align: right;
    opacity: 0;
    visibility: hidden;
    -webkit-transition: opacity 0.35s ease-in-out, -webkit-transform 0.35s ease-in-out;
    transition: opacity 0.35s ease-in-out, transform 0.35s ease-in-out;
    -webkit-transform: translateX(10px);
    -ms-transform: translateX(10px);
    transform: translateX(10px)
}
.notifications-list .mark-all-as-read {
    margin-right: 0;
    margin-top: 0;
    padding: 2px 6px 5px 10px;
    line-height: 20px;
    color: #767676
}
.notifications-list .mark-all-as-read-confirmed .text-success {
    visibility: visible;
    -webkit-transform: translateX(0);
    -ms-transform: translateX(0);
    transform: translateX(0);
    opacity: 1
}
.notifications-list .mark-all-as-read-confirmed .mark-all-as-read {
    visibility: hidden
}
.notifications-list .confirmation {
    color: #666;
    text-align: center;
    padding: 0;
    -webkit-transition: all 0.4s ease-in-out;
    transition: all 0.4s ease-in-out;
    max-height: 0;
    opacity: 0;
    overflow: hidden
}
.notifications-list .confirmation+.list-group-item {
    margin-top: -1px;
    border-top-color: #d5d5d5
}
.notifications-list .confirmation.mark-all-as-read-confirmed {
    padding: 10px 0;
    max-height: 300px;
    opacity: 1
}
.notification-actions {
    position: absolute;
    right: 10px;
    top: 8px;
    list-style: none
}
.notification-actions li {
    float: right;
    margin-left: 10px;
    font-size: 16px;
    line-height: 20px
}
.notification-actions .age {
    width: 120px;
    font-size: 12px;
    color: #767676
}
.notification-actions .undo {
    display: none;
    position: absolute;
    top: 0
}
.notification-actions .btn-link {
    padding-left: 5px;
    padding-right: 5px;
    color: #767676;
    line-height: inherit
}
.notification-actions .btn-link:hover {
    color: #4078c0;
    text-decoration: none
}
.repo-subscription-container {
    margin: 0 auto;
    width: 600px
}
.repo-subscription-container .spinner {
    float: right
}
.repo-subscription-container h2 {
    font-size: 22px;
    margin-bottom: -10px;
    font-weight: normal
}
.repo-subscription-container .intro {
    font-size: 14px;
    color: #666
}
.repo-subscription-label {
    display: inline-block
}
.subscriptions-content .sorted-by {
    color: #999;
    font-size: 12px;
    margin: 10px;
    float: right
}
.subscriptions-content .repo-icon {
    vertical-align: middle;
    color: #666;
    margin-right: 5px
}
.subscriptions-content .repo-list form {
    display: inline
}
.subscriptions-content .repo-list .only-loading {
    display: none
}
.subscriptions-content .repo-list .loading .only-loading {
    display: inline-block
}
.subscriptions-content .repo-list .only-unsubed {
    display: none
}
.subscriptions-content .repo-list .unsubscribed .only-unsubed {
    display: inline
}
.subscriptions-content .repo-list .unsubscribed .only-subed {
    display: none
}
.subscriptions-content .repo-list .only-unignored {
    display: none
}
.subscriptions-content .repo-list .unsubscribed .only-unignored {
    display: inline
}
.subscriptions-content .repo-list .unsubscribed .only-ignored {
    display: none
}
.thread-subscription-status {
    margin: 40px 0 20px;
    padding: 10px;
    color: #767676;
    background-color: #fff;
    border: 1px solid #eee;
    border-radius: 3px
}
.thread-subscription-status .mega-octicon {
    vertical-align: middle;
    margin-right: 10px;
    margin-left: 4px;
    color: #ccc
}
.thread-subscription-status .btn-sm>.octicon {
    margin-right: 1px
}
.thread-subscription-status .reason {
    display: inline-block;
    margin: 0 10px;
    vertical-align: middle
}
.thread-subscription-status .thread-subscribe-form {
    display: inline-block;
    vertical-align: middle
}
.subscription .loading {
    opacity: 0.5
}
.oauth-connection-illustration {
    position: relative;
    float: right;
    width: 200px;
    margin-top: 10px;
    margin-right: 10px;
    margin-left: 40px
}
.oauth-connection-illustration .oauth-image {
    float: left;
    padding: 5px;
    background-color: #fff;
    background-clip: padding-box;
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 6px
}
.oauth-connection-illustration .oauth-image img {
    display: block;
    width: 75px;
    height: 75px;
    border-radius: 3px
}
.oauth-connection-illustration .oauth-image.oauth-image-user {
    margin-top: 20px;
    margin-left: -20px
}
.setup-wrapper .oauth-permissions {
    margin-bottom: 25px;
    border: 1px solid #ddd;
    border-radius: 3px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05)
}
.setup-wrapper .oauth-permissions-details {
    background-color: #fff
}
.setup-wrapper.oauth-restriction-wrapper {
    padding-top: 0
}
.ellipsis-button {
    display: inline-block;
    height: 12px;
    padding: 0 5px;
    margin-left: 2px;
    font-size: 12px;
    font-weight: bold;
    line-height: 6px;
    color: #555;
    text-decoration: none;
    vertical-align: middle;
    background-color: #ddd;
    border-radius: 1px
}
.ellipsis-button:hover {
    text-decoration: none;
    background-color: #ccc
}
.ellipsis-button:before {
    content: "\2026"
}
.oauth-permissions-details {
    position: relative;
    padding: 15px;
    margin: 0;
    list-style: none;
    border-bottom: 1px solid #f2f2f2
}
.oauth-permissions-details:first-child {
    border-radius: 3px 3px 0 0
}
.oauth-permissions-details:last-child {
    border: 0;
    border-radius: 0 0 3px 3px
}
.oauth-permissions-details.oauth-public-data-only {
    border-radius: 3px
}
.oauth-permissions-details .markdown-body {
    font-size: 13px
}
.oauth-permissions-details .content {
    display: none;
    margin-left: 45px
}
.oauth-permissions-details .content .form-checkbox {
    margin-left: 0
}
.oauth-permissions-details .content .form-checkbox:last-child {
    margin-bottom: 0
}
.oauth-permissions-details .mega-octicon {
    float: left;
    width: 32px;
    margin-top: 1px;
    margin-left: 0;
    color: #767676;
    text-align: center
}
.oauth-permissions-details .permission-help {
    font-size: 13px
}
.oauth-permissions-details .permission-help ul {
    padding-left: 20px;
    margin: 1em 0
}
.oauth-permissions-details .permission-summary {
    margin-left: 45px
}
.oauth-permissions-details .permission-summary .access-details {
    position: relative;
    color: #767676
}
.oauth-permissions-details .permission-summary em.highlight {
    position: relative;
    padding: 2px 3px;
    margin-right: -2px;
    margin-left: -3px;
    font-style: normal;
    color: #4c4a42;
    background: #fff9ea;
    border-radius: 3px
}
.oauth-permissions-details .permission-title {
    display: block;
    color: #000
}
.oauth-permissions-details a.btn-sm {
    float: right;
    margin-top: 4px
}
.oauth-permissions-details.open a.btn-sm {
    background-color: #dcdcdc;
    background-image: none;
    border-color: #b5b5b5;
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.15)
}
.oauth-permissions-details.open .content {
    display: block
}
.oauth-permissions-details.default:not(.delete) .no-access,
.oauth-permissions-details.default:not(.delete) .default-access,
.oauth-permissions-details.none .no-access,
.oauth-permissions-details.none .default-access {
    display: inline
}
.oauth-permissions-details.default:not(.delete) .access-details,
.oauth-permissions-details.default:not(.delete) .permission-title,
.oauth-permissions-details.none .access-details,
.oauth-permissions-details.none .permission-title {
    color: #999
}
.oauth-permissions-details.default:not(.delete) .mega-octicon,
.oauth-permissions-details.none .mega-octicon {
    color: #ccc
}
.oauth-permissions-details.default .default-access {
    display: inline
}
.oauth-permissions-details.full .full-access {
    display: inline
}
.oauth-details-toggle {
    position: absolute;
    top: 0;
    right: 0;
    padding: 20px 10px
}
.oauth-details-toggle .mega-octicon {
    font-size: 22px
}
.oauth-details-toggle .octicon-chevron-up {
    display: none
}
.open .oauth-details-toggle .octicon-chevron-down {
    display: none
}
.open .oauth-details-toggle .octicon-chevron-up {
    display: block
}
.oauth-user-permissions .full-access,
.oauth-user-permissions .limited-access,
.oauth-user-permissions .limited-access-none,
.oauth-user-permissions .limited-access-followers,
.oauth-user-permissions .limited-access-emails,
.oauth-user-permissions .no-access {
    display: none
}
.oauth-user-permissions.limited .limited-access-none {
    display: inline
}
.oauth-user-permissions.limited.limited-email .limited-access,
.oauth-user-permissions.limited.limited-email .limited-access-none {
    display: none
}
.oauth-user-permissions.limited.limited-email .limited-access-emails {
    display: inline
}
.oauth-user-permissions.limited.limited-email.limited-follow .limited-access {
    display: inline
}
.oauth-user-permissions.limited.limited-email.limited-follow .limited-access-none,
.oauth-user-permissions.limited.limited-email.limited-follow .limited-access-emails,
.oauth-user-permissions.limited.limited-email.limited-follow .limited-access-followers {
    display: none
}
.oauth-user-permissions.limited.limited-follow .limited-access,
.oauth-user-permissions.limited.limited-follow .limited-access-none {
    display: none
}
.oauth-user-permissions.limited.limited-follow .limited-access-followers {
    display: inline
}
.oauth-repo-permissions .default-access,
.oauth-repo-permissions .public-access,
.oauth-repo-permissions .full-access {
    display: none
}
.oauth-repo-permissions.public .public-access {
    display: inline
}
.oauth-delete-repo-permissions .octicon-alert {
    color: #bd2c00
}
.oauth-repo-status-permissions .no-access,
.oauth-repo-status-permissions .full-access,
.oauth-repo-deployment-permissions .no-access,
.oauth-repo-deployment-permissions .full-access {
    display: none
}
.oauth-notifications-permissions .no-access,
.oauth-notifications-permissions .read-access,
.oauth-notifications-permissions .via-public-access,
.oauth-notifications-permissions .via-full-access {
    display: none
}
.oauth-notifications-permissions.read .read-access {
    display: inline
}
.oauth-notifications-permissions.via-public .via-public-access {
    display: inline
}
.oauth-notifications-permissions.via-public .octicon {
    display: none
}
.oauth-notifications-permissions.via-full .via-full-access {
    display: inline
}
.oauth-gist-permissions .no-access,
.oauth-gist-permissions .full-access {
    display: none
}
.oauth-granular-permissions .no-access,
.oauth-granular-permissions .read-access,
.oauth-granular-permissions .write-access,
.oauth-granular-permissions .full-access {
    display: none
}
.oauth-granular-permissions.none .no-access {
    display: inline
}
.oauth-granular-permissions.read .read-access {
    display: inline
}
.oauth-granular-permissions.write .write-access {
    display: inline
}
.oauth-granular-permissions.full .full-access {
    display: inline
}
.oauth-secondary .setup-info-module {
    margin-top: 0
}
.oauth-secondary .setup-info-module .no-description {
    color: #767676
}
.oauth-secondary .setup-info-module .features-list {
    padding-bottom: 0
}
.oauth-no-description {
    color: #767676
}
.oauth-org-access-details a:hover {
    text-decoration: none
}
.oauth-org-access-details .boxed-group-list>li {
    line-height: 24px
}
.oauth-org-access-details .boxed-group-list>li .loading-indicator {
    display: none;
    margin: 4px
}
.oauth-org-access-details .boxed-group-list>li.on {
    background: #fff
}
.oauth-org-access-details .boxed-group-list>li.on:hover {
    background: #ffe
}
.oauth-org-access-details .boxed-group-list>li.on .authorized-tools {
    display: block
}
.oauth-org-access-details .boxed-group-list>li.on .unauthorized-tools {
    display: none
}
.oauth-org-access-details .boxed-group-list>li.on strong {
    color: #333
}
.oauth-org-access-details .boxed-group-list>li.on .octicon-check {
    display: inline
}
.oauth-org-access-details .boxed-group-list>li.on .octicon-x {
    display: none
}
.oauth-org-access-details .boxed-group-list>li.loading .unauthorized-tools,
.oauth-org-access-details .boxed-group-list>li.loading .authorized-tools {
    display: none
}
.oauth-org-access-details .boxed-group-list>li.loading .loading-indicator {
    display: block
}
.oauth-org-access-details .boxed-group-list>li .authorized-tools {
    display: none
}
.oauth-org-access-details .boxed-group-list>li .unauthorized-tools {
    display: block
}
.oauth-org-access-details .boxed-group-list>li .btn {
    padding: 0 10px;
    margin-left: 15px;
    line-height: 24px
}
.oauth-org-access-details .octicon {
    color: #979797
}
.oauth-org-access-details .octicon-check {
    display: none;
    color: #6cc644
}
.oauth-org-access-details .octicon-x {
    display: inline
}
.oauth-org-access-details .octicon-x.org-access-denied {
    color: #bd2c00
}
.deleted-permission {
    color: #bd2c00
}
.added-permission {
    color: #6cc644
}
.permission-title {
    margin-top: 0
}
.oauth-application-whitelist .request-info {
    display: block;
    margin-left: 25px;
    color: #9b9b9b
}
.oauth-application-whitelist .request-info strong {
    color: #333
}
.oauth-application-whitelist .request-info .application-description {
    display: none
}
.oauth-application-whitelist .request-info.open .application-description {
    display: block
}
.oauth-application-whitelist .avatar {
    margin-top: 0
}
.oauth-application-whitelist .requestor {
    font-weight: bold
}
.oauth-application-whitelist .octicon-alert {
    color: #c9510c
}
.oauth-application-whitelist .octicon-check,
.oauth-application-whitelist .approved-request {
    color: #6cc644
}
.oauth-application-whitelist .denied-request {
    color: #bd2c00
}
.oauth-application-whitelist .request-indicator {
    margin-left: 10px
}
.oauth-application-whitelist .edit-link {
    color: #999
}
.oauth-application-whitelist .edit-link:hover {
    color: #4078c0
}
.oauth-application-whitelist .boxed-group-list {
    margin-top: 1em
}
.oauth-application-whitelist .boxed-group-list li {
    padding: 10px
}
.oauth-application-info {
    min-height: 70px;
    padding-top: 10px;
    margin-bottom: 30px
}
.boxed-group-inner .oauth-application-info {
    margin-bottom: 10px
}
.oauth-application-info .application-title,
.oauth-application-info .application-description,
.oauth-application-info .application-meta-info {
    margin-left: 70px
}
.oauth-application-info .application-title {
    font-size: 14px;
    font-weight: bold;
    color: #333
}
.oauth-application-info .application-description {
    margin-top: 3px;
    margin-bottom: 0
}
.oauth-application-info .app-info {
    display: inline-block;
    margin-right: 10px;
    color: #999
}
.oauth-application-info .app-info .octicon {
    margin-right: 5px
}
.oauth-application-info .meta-link {
    color: #999
}
.oauth-application-info .meta-link:hover {
    color: #4078c0
}
.oauth-application-info .application-meta-info {
    margin-top: 3px;
    font-size: 12px
}
.oauth-application-info .app-denied,
.oauth-application-info .app-approved {
    margin-left: 10px;
    font-weight: normal;
    white-space: nowrap
}
.oauth-application-info .app-approved,
.oauth-application-info .octicon-check {
    color: #6cc644
}
.oauth-application-info .app-denied,
.oauth-application-info .octicon-x {
    color: #c9510c
}
.restrict-oauth-access-button {
    margin-right: 20px
}
.restrict-oauth-access-info {
    margin-bottom: 40px;
    font-size: 15px
}
.restrict-oauth-access-list {
    padding-left: 25px
}
.restrict-oauth-access-list li {
    margin-bottom: 10px
}
.restrict-oauth-access-list li:last-child {
    margin-bottom: 0
}
.app-transfer-actions form {
    display: inline
}
.application-authorizations:target {
    background-color: #ffe
}
.application-authorizations .oauth-logo-cell {
    width: 20px
}
.application-authorizations .oauth-app-access-name {
    font-size: 15px;
    font-weight: bold;
    line-height: 1.2;
    color: #333
}
.application-authorizations .oauth-app-access-name:hover {
    color: #4078c0;
    text-decoration: none
}
.application-authorizations .oauth-app-list-meta {
    margin-top: 1px;
    margin-bottom: 2px;
    font-weight: normal;
    color: #999
}
.application-authorizations .oauth-info-cell {
    width: 599px;
    padding-right: 0;
    padding-left: 0
}
.application-authorizations .oauth-app-owner {
    color: #999
}
.application-authorizations .oauth-app-owner:hover {
    color: #4078c0;
    text-decoration: none
}
.application-authorizations .oauth-view-revoke-cell {
    width: 90px;
    text-align: right
}
.developer-app-item .developer-app-avatar-cell {
    width: 60px
}
.developer-app-item .developer-app-name {
    font-size: 15px;
    font-weight: bold;
    line-height: 1.2;
    color: #333
}
.developer-app-item .developer-app-name:hover {
    color: #4078c0;
    text-decoration: none
}
.developer-app-item .developer-app-info-cell {
    padding-left: 0
}
.developer-app-item .developer-app-list-meta {
    margin-top: 3px;
    margin-bottom: 2px;
    font-weight: normal;
    color: #999
}
.org-transfer-requests {
    margin: 10px 0 20px
}
.org-oauth-applications-header {
    margin-top: 0
}
.org-header {
    padding-top: 20px;
    padding-bottom: 20px;
    margin-bottom: 20px;
    color: #666;
    border-bottom: 1px solid #eee;
    background-color: #fcfcfc
}
.org-header .edit-org {
    position: relative;
    top: -6px;
    display: inline-block;
    padding: 3px 5px;
    font-size: 14px;
    color: #aaa;
    border: 1px solid #e5e5e5;
    border-radius: 3px
}
.org-header .edit-org:hover {
    color: #4078c0;
    background-color: #fff
}
.org-header .edit-org .octicon {
    font-size: 14px
}
.org-header-wrapper .avatar {
    display: block;
    width: 100px
}
.org-header-wrapper .flex-table-item-primary {
    padding-left: 20px;
    white-space: normal
}
.org-name {
    margin-top: 0;
    margin-bottom: 5px;
    color: #333;
    font-size: 36px;
    font-weight: normal
}
.org-description {
    margin-top: 0;
    margin-bottom: 8px;
    font-size: 16px;
    line-height: 1.25
}
.org-header-meta {
    font-size: 12px;
    line-height: 1.5;
    list-style: none
}
.org-header-meta .meta-item {
    display: inline-block;
    padding-right: 18px;
    overflow: hidden;
    max-width: 100%;
    text-overflow: ellipsis;
    white-space: nowrap
}
.org-header-meta .meta-item .meta-link {
    color: #666
}
.org-header-meta .octicon {
    position: relative;
    top: 1px;
    margin-right: 2px;
    color: #ccc
}
.org-header-meta.has-email.has-blog .meta-item,
.org-header-meta.has-email.has-location .meta-item,
.org-header-meta.has-blog.has-email .org-header-meta.has-blog.has-location .meta-item,
.org-header-meta.has-location.has-blog .meta-item,
.org-header-meta.has-location.has-email .meta-item {
    max-width: 278px
}
.org-header-meta.has-email.has-blog.has-location .meta-item {
    max-width: 186px
}
.org-link {
    color: #333
}
.org-link:hover {
    color: #4078c0;
    text-decoration: none
}
.org-main {
    float: left;
    width: 640px
}
.org-sidebar {
    float: right;
    width: 280px
}
.org-sidebar .member-badge {
    display: block;
    padding-top: 0;
    padding-bottom: 0;
    border-top: 0
}
.pagehead-orgs .pagehead-tabs {
    margin-top: 10px;
    margin-left: -15px
}
.simple-box {
    padding: 15px;
    margin-bottom: 20px;
    border: 1px solid #ddd;
    background-color: #fff;
    border-radius: 3px
}
.simple-box-title {
    margin: -15px -15px 0;
    font-size: 18px;
    padding: 15px;
    border-bottom: 1px solid #eee
}
.simple-box-footer {
    margin: 10px -15px -15px;
    padding: 15px;
    border-top: 1px solid #eee;
    background-color: #fcfcfc;
    border-radius: 0 0 3px 3px
}
.orgs-help-shelf {
    margin-top: -20px;
    margin-bottom: 20px;
    padding-top: 20px;
    padding-bottom: 20px;
    border-bottom: 1px solid #eee
}
.orgs-help-shelf .orgs-help-title {
    font-size: 30px;
    font-weight: normal
}
.orgs-help-shelf-content {
    width: 800px;
    margin: 50px auto;
    text-align: center
}
.orgs-help-shelf-content .orgs-help-lead,
.orgs-help-shelf-content .orgs-help-description {
    font-size: 18px
}
.orgs-help-shelf-content .orgs-help-lead {
    padding-left: 45px;
    padding-right: 45px
}
.orgs-help-shelf-content .orgs-help-divider {
    width: 150px;
    margin: 40px auto;
    display: block;
    border-top: 1px solid #ddd;
    content: ""
}
.orgs-help-lead {
    margin-bottom: 30px
}
.orgs-help-items {
    margin-bottom: 40px
}
.orgs-help-item-octicon {
    width: 70px;
    height: 70px;
    margin: 0 auto 15px;
    text-align: center;
    border: solid 1px #e5e5e5;
    border-radius: 50px;
    background-color: #fff
}
.orgs-help-item-octicon .mega-octicon {
    color: #4078c0;
    font-size: 26px;
    line-height: 68px
}
.orgs-help-item-title {
    margin-bottom: 10px;
    font-weight: normal
}
.orgs-help-item-content {
    margin-top: 0;
    color: #666;
    font-size: 14px
}
.orgs-help-dismiss {
    margin-top: 5px;
    margin-right: 10px;
    float: right;
    font-size: 12px;
    color: #767676
}
.orgs-help-dismiss:hover {
    color: #4078c0;
    text-decoration: none
}
.orgs-help-dismiss .octicon {
    position: relative;
    top: 1px
}
#revoke_all_repo_access,
#revoke_active_repo_access {
    display: none
}
.orgs-help-title {
    margin-top: 0;
    margin-bottom: 0
}
.orgs-help-description {
    font-size: 14px
}
.orgs-help-lead,
.orgs-help-description {
    margin-top: 10px;
    color: #666
}
.orgs-help-button {
    margin-right: 10px
}
.org-module-title {
    margin: -15px -15px 0;
    font-size: 18px;
    border-bottom: 1px solid #eee
}
.org-module-link {
    display: block;
    padding: 15px;
    color: #333
}
.org-module-link:hover,
.org-module-link:hover .org-stats {
    text-decoration: none;
    color: #4078c0
}
.org-stats {
    margin-top: 3px;
    float: right;
    font-size: 14px;
    color: #767676
}
.org-members-title {
    margin-bottom: 0;
    border-bottom: 0
}
.member-avatar-group {
    margin: -1px
}
.member-avatar-group:before {
    display: table;
    content: ""
}
.member-avatar-group:after {
    display: table;
    clear: both;
    content: ""
}
.member-avatar {
    float: left;
    margin: 1px
}
.member-row {
    display: block;
    padding-bottom: 15px;
    margin-bottom: 15px;
    margin-top: 15px;
    font-size: 14px;
    color: #333;
    border-bottom: 1px solid #eee
}
.member-row:before {
    display: table;
    content: ""
}
.member-row:after {
    display: table;
    clear: both;
    content: ""
}
.member-row:hover {
    color: #4078c0;
    text-decoration: none
}
.member-row:last-child {
    margin-bottom: 0;
    padding-bottom: 0;
    border-bottom: 0
}
.member-row .avatar {
    float: left;
    margin-right: 10px
}
.member-row .member-name {
    display: block
}
.member-fullname {
    color: #767676
}
.org-no-members {
    text-align: center;
    color: #767676;
    margin-top: 20px;
    margin-bottom: 10px
}
.org .no-results {
    padding: 10px;
    color: #767676
}
.org-toolbar.disabled {
    pointer-events: none
}
.org-toolbar .subnav-search {
    width: 320px;
    margin-left: 0
}
.org-toolbar .non-admin-search .subnav-divider-right {
    padding-right: 0;
    border-width: 0
}
.org-toolbar .subnav-search-context+.subnav-search {
    margin-left: -1px
}
.org-toolbar input.subnav-search-input {
    width: 100%
}
.member-list-select-all-label {
    font-weight: normal
}
.member-list-select-all-label .some-selected {
    display: none
}
.member-list-select-all-label.has-selected-members .some-selected {
    display: inline
}
.member-list-select-all-label.has-selected-members .none-selected {
    display: none
}
.pending-invitations-link {
    padding-left: 15px;
    padding-right: 15px
}
.member-toolbar-actions {
    margin-top: 9px;
    margin-right: 9px
}
.member-action {
    margin-right: 5px
}
.member-role-select {
    display: inline
}
.member-role-select .select-menu-modal {
    left: -190px;
    width: 310px
}
.member-role-menu .select-menu-item-text {
    padding-right: 8px
}
.legacy-contributor-note {
    padding: 10px;
    color: #767676;
    font-size: 11px;
    background-color: #f8f8f8;
    border-bottom: 1px solid #eee
}
.legacy-contributor-note-content {
    margin-top: 0;
    margin-bottom: 0
}
.auto-search-group {
    position: relative
}
.auto-search-group .auto-search-input {
    padding-left: 30px
}
.auto-search-group .spinner,
.auto-search-group>.octicon {
    position: absolute;
    left: 10px;
    width: 16px;
    height: 16px;
    z-index: 5
}
.auto-search-group .spinner {
    top: 9px;
    background-color: #fff
}
.auto-search-group>.octicon {
    top: 10px;
    font-size: 14px;
    color: #bbb;
    text-align: center
}
.org-list .list-item {
    position: relative;
    padding-top: 15px;
    padding-bottom: 15px;
    border-bottom: 1px solid #eee
}
.org-list .list-item:before {
    display: table;
    content: ""
}
.org-list .list-item:after {
    display: table;
    clear: both;
    content: ""
}
.org-list .cancel-link {
    color: #767676
}
.org-repos .blankslate,
.org-team-main .blankslate {
    margin-top: 15px
}
.org-repos-mini {
    padding: 0;
    margin: 0
}
.org-repos-mini .org-repo-mini-item:first-child .org-repo-mini-cell {
    border-top: 0
}
.org-repos-mini .org-repo-icon {
    vertical-align: middle
}
.org-repos-mini .org-repo-name {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 14px;
    word-wrap: break-word
}
.org-repos-mini .org-repo-name .octicon-repo {
    color: #767676
}
.org-repos-mini .org-repo-name .octicon-lock {
    color: #e9dba5
}
.org-repos-mini .org-repo-name .repo-prefix {
    font-weight: normal;
    text-transform: lowercase
}
.org-repos-mini .org-repo-name .repo-slash {
    display: inline-block;
    margin-left: -4px;
    margin-right: -4px
}
.org-repos-mini .org-repo-forked {
    max-width: 270px;
    margin-top: 0;
    margin-bottom: 0;
    display: inline-block;
    font-weight: normal;
    font-size: 12px
}
.org-repo-mini-cell {
    padding-top: 15px;
    padding-bottom: 15px;
    vertical-align: middle
}
.org-repo-meta {
    width: 165px
}
.org-repo-meta .access-level {
    cursor: default
}
.org-repo-access-level {
    text-align: center
}
.org-repo-manage {
    width: 270px
}
.org-repo-higher-access {
    display: none;
    margin-left: 16px;
    margin-top: 2px;
    font-size: 11px
}
.org-higher-access-member .manage-access {
    font-size: 12px;
    position: relative;
    top: 2px
}
.with-higher-access .org-repo-higher-access {
    display: block
}
.with-higher-access .table-list-cell-checkbox {
    vertical-align: top
}
.permission-level-cell .select-menu-button {
    width: 100px
}
.permission-level-cell .select-menu-button:after {
    position: absolute;
    top: 10px;
    right: 10px
}
.permission-level-cell .spinner,
.permission-level-cell .permission-success-icon {
    position: absolute;
    margin-left: 15px;
    display: inline-block;
    opacity: 0;
    -webkit-transition: opacity 0.2s ease-in-out;
    transition: opacity 0.2s ease-in-out
}
.permission-level-cell .permission-success-icon {
    margin-top: 4px;
    color: #6cc644
}
.permission-level-cell .is-updating .spinner,
.permission-level-cell .was-successful .permission-success-icon {
    opacity: 1
}
.org-repo-permission-select .select-menu-modal .description {
    padding-right: 20px
}
.org-repo-permission-select .select-menu-option-title {
    margin-top: 0;
    margin-bottom: 0
}
.org-repo-permission-select .navigation-focus .select-menu-option-title {
    color: #fff
}
.add-member-wrapper {
    position: relative;
    width: 680px;
    margin: 40px auto
}
.add-member-wrapper .owners-team-info {
    position: relative;
    padding-top: 10px;
    padding-left: 42px;
    color: #767676;
    border-top: 1px solid #e5e5e5
}
.add-member-wrapper .owners-team-info .octicon-info {
    position: absolute;
    left: 8px;
    color: #767676;
    font-size: 18px
}
.add-member-wrapper .available-seats {
    color: #767676
}
.add-member-wrapper .buy-more-link {
    margin-right: 5px
}
.add-member-wrapper .send-invitation-button {
    float: none
}
.invitation-role-group {
    padding-top: 25px;
    padding-bottom: 30px
}
.invitation-role-group:before {
    display: table;
    content: ""
}
.invitation-role-group:after {
    display: table;
    clear: both;
    content: ""
}
.invitation-role {
    position: relative;
    padding: 15px 10px;
    display: block;
    font-weight: normal;
    text-align: center;
    cursor: pointer
}
.invitation-role:first-child {
    margin-left: 0
}
.invitation-role .invitation-role-item-border {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    border: 1px solid #eee;
    border-radius: 3px;
    -webkit-transition: border 0.15s ease-in-out;
    transition: border 0.15s ease-in-out
}
.invitation-role input {
    display: block;
    margin-right: auto;
    margin-left: auto
}
.invitation-role input:checked ~ .invitation-role-item-border {
    border: 1px solid #3b99fc;
    box-shadow: 0 0 5px rgba(59, 153, 252, 0.4)
}
.invitation-role-title {
    margin-top: 5px;
    margin-bottom: 5px
}
.invitation-role-description {
    margin-top: 0;
    color: #767676
}
.add-member-title {
    margin-bottom: 0;
    font-size: 30px;
    font-weight: normal
}
.add-member-lead {
    margin-top: 5px;
    margin-bottom: 0;
    padding-bottom: 15px;
    border-bottom: 1px solid #eee
}
.add-member-lead.no-border {
    padding-bottom: 0;
    border-bottom: 0
}
.add-member-team-list {
    margin-bottom: 15px;
    list-style: none
}
.add-member-team-list .team {
    display: block;
    padding: 15px 0;
    font-weight: normal;
    cursor: pointer
}
.add-member-team-list .team:first-child {
    border-top: 1px solid #f2f2f2
}
.add-member-team-list .team .btn-sm {
    float: right
}
.add-member-team-list .team-info {
    max-width: 80%;
    color: #000;
    text-decoration: none
}
.add-member-team-list .team-info:hover {
    color: #4078c0
}
.add-member-team-list .team-name {
    font-size: 14px
}
.add-member-team-list .team-meta {
    color: #767676;
    margin-top: 2px;
    margin-bottom: 2px
}
.add-member-team-list .team-description {
    margin-top: 2px;
    margin-bottom: 2px;
    color: #333
}
.add-member-team-list .team-toggler .turn-on {
    display: inline-block
}
.add-member-team-list .team-toggler.on .turn-off {
    display: inline-block
}
.add-member-team-list .team-toggler .turn-off {
    display: none
}
.add-member-team-list .team-toggler.on .turn-on {
    display: none
}
.team-list-footer .show-all-link .octicon {
    margin-left: 5px;
    color: #767676
}
.invite-team-member-list .team {
    display: table-row;
    cursor: default
}
.invite-team-member-list .team:first-child .table-list-cell {
    border-top: 0
}
.invite-team-member-list .team .table-list-cell {
    padding-top: 15px;
    padding-bottom: 15px
}
.invite-team-member-list .table-list-cell-checkbox {
    width: 42px
}
.invite-team-member-list .team-info {
    width: 250px;
    vertical-align: middle;
    color: #222
}
.invite-team-member-list .team-description {
    display: block;
    padding-top: 0;
    padding-bottom: 0;
    font-weight: normal
}
.invite-team-member-list .team-meta {
    width: 100px;
    text-align: left;
    vertical-align: middle
}
.invite-team-member-list .team-link {
    text-align: right;
    color: #4078c0
}
.member-list-item .table-list-cell {
    vertical-align: middle;
    padding-top: 10px;
    padding-bottom: 10px
}
.member-list-item .table-list-cell-checkbox {
    width: 30px
}
.member-list-item.adminable .member-info {
    padding-left: 5px
}
.member-list-item .member-link {
    display: block;
    text-decoration: none
}
.member-list-item .member-link:hover .member-username {
    color: #4078c0
}
.member-visibility .octicon {
    font-size: 14px
}
.member-info {
    padding-left: 10px;
    font-size: 14px
}
.member-info .member-list-avatar {
    float: left;
    margin-right: 15px
}
.member-info .member-fullname {
    font-weight: normal;
    color: #767676
}
.member-username {
    margin-top: 4px;
    display: block;
    color: #333
}
.member-username .octicon {
    position: relative;
    top: -2px;
    margin-left: 2px;
    font-size: 12px;
    color: #aaa
}
.member-username.css-truncate-target {
    display: block
}
.member-security .octicon {
    color: #c9510c;
    font-size: 14px
}
.member-meta {
    width: 140px;
    font-size: 11px;
    color: #767676;
    text-align: center
}
.member-meta .access-link {
    color: #767676
}
.member-meta .access-link:hover {
    color: #4078c0;
    text-decoration: none
}
.member-meta .btn-link {
    color: #767676
}
.member-meta .btn-link:hover {
    color: #4078c0;
    text-decoration: none
}
.member-meta .select-menu-modal {
    width: 310px
}
.member-meta .select-menu-modal-holder {
    right: 0;
    text-align: left
}
.member-meta .octicon {
    font-size: 14px
}
.non-member-meta {
    width: 300px
}
.member-follow {
    text-align: right
}
.member-selected-actions {
    display: inline
}
.org-people-blankslate {
    margin-top: -20px;
    border-top-width: 0;
    border-top-left-radius: 0;
    border-top-right-radius: 0
}
.migration-wrapper {
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif
}
.migration-jumbotron {
    height: 70vh;
    min-height: 450px;
    max-height: 650px
}
.migration-jumbotron,
.migration-sub-header {
    position: relative;
    margin-top: -1px;
    background-color: #3f4851;
    background-image: -webkit-linear-gradient(#3f4851 0, #282d33 100%);
    background-image: linear-gradient(#3f4851 0, #282d33 100%)
}
.migration-jumbotron:after,
.migration-sub-header:after {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    content: " ";
    background-image: url(https://github.com/images/modules/orgs/dots-bg.png);
    background-size: 80%;
    background-repeat: repeat;
    opacity: 0.75;
    z-index: 1
}
.migration-jumbotron-content {
    position: relative;
    top: 50%;
    width: 980px;
    margin: 0 auto;
    padding: 100px 60px;
    text-align: center;
    -webkit-transform: translateY(-50%);
    -ms-transform: translateY(-50%);
    transform: translateY(-50%);
    z-index: 2
}
.migration-jumbotron-octicons {
    height: 60px;
    margin-bottom: 20px;
    text-align: center
}
.migration-jumbotron-octicon-item {
    position: relative;
    display: inline-block;
    width: 60px;
    height: 60px;
    margin-right: 10px;
    margin-left: 10px;
    border-radius: 50px
}
.migration-jumbotron-octicon-item:after {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 1;
    content: " ";
    background-image: -webkit-linear-gradient(135deg, #6e5494 30%, #c9510c 100%);
    background-image: linear-gradient(-45deg, #6e5494 30%, #c9510c 100%);
    border-radius: 50px
}
.migration-jumbotron-octicon-item .mega-octicon {
    position: absolute;
    top: 1px;
    left: 1px;
    z-index: 2;
    width: 58px;
    height: 58px;
    color: rgba(255, 255, 255, 0.9);
    font-size: 24px;
    line-height: 60px;
    background-color: #383f47;
    border-radius: 50px
}
.migration-jumbotron-title,
.migration-jumbotron-lead,
.migration-sub-title,
.migration-sub-lead {
    color: #fff;
    text-shadow: 0 1px 1px rgba(0, 0, 0, 0.05)
}
.migration-jumbotron-title,
.migration-section-title {
    margin-bottom: 10px;
    font-weight: 300
}
.migration-jumbotron-title {
    font-size: 40px
}
.migration-section-title {
    margin-top: 0;
    font-size: 30px
}
.migration-jumbotron-lead {
    margin-top: 0;
    font-size: 24px;
    opacity: 0.85
}
.migration-section-lead {
    margin-top: 20px;
    margin-bottom: 20px
}
.migration-jumbotron-btn {
    padding: 12px 18px;
    font-size: 16px;
    color: #6e5494;
    background-color: #fff;
    border-width: 0;
    box-shadow: 0 3px 3px rgba(0, 0, 0, 0.05)
}
.migration-jumbotron-btn:hover {
    color: #6e5494;
    background-color: #eee
}
.migration-jumbotron-soon {
    padding: 12px 18px;
    font-size: 16px;
    color: #ddd;
    border: 1px solid rgba(255, 255, 255, 0.2)
}
.migration-section {
    padding-top: 100px;
    padding-bottom: 100px;
    overflow: hidden;
    border-bottom: 1px solid #ddd
}
.migration-feature-list {
    margin-top: 30px;
    margin-bottom: 20px;
    overflow: hidden;
    color: #767676;
    font-size: 14px;
    list-style: none
}
.migration-feature-list:before {
    content: "";
    width: 100px;
    margin-bottom: 30px;
    display: block;
    border-top: 1px solid #ddd
}
.migration-feature-list .octicon {
    width: 22px;
    margin-left: -3px;
    color: #767676;
    text-align: center
}
.migration-feature-list-item {
    width: 50%;
    margin-bottom: 15px;
    float: left
}
.migration-section-grey {
    background-color: #fcfcfc
}
.migration-illustration-wrapper:before {
    display: table;
    content: ""
}
.migration-illustration-wrapper:after {
    display: table;
    clear: both;
    content: ""
}
.migration-illustration {
    width: 700px;
    margin-top: -30px;
    margin-bottom: -50px;
    border: 1px solid #ddd;
    border-radius: 3px;
    box-shadow: 0 0 15px rgba(0, 0, 0, 0.05)
}
.migration-illustration-left {
    float: right;
    margin-right: 50px
}
.migration-illustration-right {
    float: left;
    margin-left: 50px
}
.migration-section-privileges {
    padding-top: 80px;
    padding-bottom: 80px
}
.migration-footer {
    position: relative;
    z-index: 1;
    padding-top: 60px;
    padding-bottom: 60px;
    margin-bottom: -41px;
    border-bottom: 1px solid #ddd
}
.migration-footer-content {
    width: 800px;
    margin: 0 auto;
    text-align: center
}
.migration-footer-title,
.migration-footer-lead {
    margin-bottom: 0
}
.migration-footer-lead {
    margin-top: 10px
}
.migration-footer-btn,
.migration-footer-soon {
    margin-top: 20px
}
.migration-footer-soon {
    padding: 6px 12px;
    font-size: 14px;
    color: #666;
    border: 1px solid rgba(0, 0, 0, 0.2)
}
.migration-soon-tag {
    display: inline-block;
    line-height: 20px;
    white-space: nowrap;
    vertical-align: middle;
    border-radius: 3px
}
.migration-sub-header {
    margin-bottom: 40px;
    padding-top: 40px;
    padding-bottom: 40px
}
.org-settings-updating {
    padding: 15px;
    margin-top: 0;
    margin-bottom: 30px;
    background-color: #fff;
    border: 1px solid #ddd;
    border-radius: 3px
}
.org-settings-updating .spinner {
    display: inline-block;
    vertical-align: middle;
    margin-top: -2px
}
.org-disabled-settings {
    pointer-events: none;
    opacity: 0.5
}
.migration-sub-header-content {
    width: 68%
}
.migration-sub-title {
    margin-bottom: 0
}
.migration-sub-lead {
    margin-top: 10px;
    margin-bottom: 0
}
.migration-org-avatar {
    margin-top: 6px;
    margin-right: 72px;
    border: 3px solid #fff;
    border-radius: 3px
}
.org-migration-settings-sidebar .migrate-org-roles {
    margin-top: 0;
    margin-bottom: 10px
}
.org-migration-settings-sidebar .preserve-member-privileges-btn {
    display: none
}
.org-migration-settings-sidebar .member-privilege-radios-preserved .preserve-member-privileges-btn {
    display: block
}
.org-migration-settings-sidebar .member-privilege-radios-preserved .save-member-privileges-btn {
    display: none
}
.org-migration-settings-section {
    position: relative;
    margin-right: 60px;
    margin-bottom: 50px;
    padding-bottom: 50px;
    border-bottom: 1px solid #ddd
}
.org-migration-settings-section:last-child {
    margin-bottom: 0;
    padding-bottom: 0;
    border-bottom: 0
}
.org-migration-settings-section .disabled {
    opacity: 0.5;
    pointer-events: none
}
.org-migration-settings-section .spinner {
    display: inline-block;
    margin-bottom: -3px
}
.org-migration-settings-icon {
    position: absolute;
    left: -45px;
    color: #ccc
}
.org-migration-settings-title {
    margin-bottom: 0;
    font-size: 22px;
    font-weight: normal
}
.org-migration-settings-info {
    margin-top: 5px;
    margin-bottom: 30px;
    font-size: 16px;
    color: #767676
}
.migrate-owners-wrapper {
    position: relative;
    min-height: 550px;
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif
}
.migrate-owners-content-about,
.migrate-owners-content-rename {
    position: absolute;
    top: 0;
    left: 50%;
    margin: 30px auto 0;
    -webkit-transition: opacity 0.2s ease-in-out, -webkit-transform 0.3s ease-in-out;
    transition: opacity 0.2s ease-in-out, transform 0.3s ease-in-out;
    -webkit-transform: translate(-50%, 0);
    -ms-transform: translate(-50%, 0);
    transform: translate(-50%, 0)
}
.migrate-owners-content-hidden {
    -webkit-transform: translate(-50%, 150px);
    -ms-transform: translate(-50%, 150px);
    transform: translate(-50%, 150px);
    opacity: 0;
    pointer-events: none;
    z-index: 20
}
.migrate-owners-content-about {
    width: 700px;
    text-align: center
}
.migrate-owners-title {
    font-size: 35px;
    font-weight: normal
}
.migrate-owners-lead {
    margin-top: 0;
    margin-bottom: 20px
}
.migrate-owners-content-rename {
    width: 520px
}
.rename-owners-error span {
    display: inline-block;
    padding: 5px;
    margin-bottom: 10px;
    font-size: 11px;
    font-weight: bold;
    color: #494620;
    background: #f7ea57;
    border: 1px solid #c0b536;
    border-top-color: #fff;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.rename-owners-spinner {
    position: absolute;
    top: 30px;
    right: 30px
}
.delete-owners-button {
    color: #767676
}
.delete-owners-button:hover {
    color: #bd2c00
}
.rename-owners-team-form .rename-owners-team-input {
    font-size: 22px;
    font-weight: bold
}
.rename-owners-team-form .note {
    margin-top: 5px;
    margin-bottom: 15px;
    color: #767676
}
.legacy-contributors-title {
    margin-top: 30px;
    margin-bottom: 0;
    font-size: 24px;
    font-weight: normal
}
.legacy-contributors-lead {
    margin-top: 10px;
    font-size: 16px;
    line-height: 24px
}
.migration-help-collabs {
    margin-top: 145px
}
.migration-help-robots {
    margin-top: 80px
}
.migration-help-teams {
    margin-top: 50px
}
.migration-help-title {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 14px
}
.migration-help-content {
    margin-top: 5px;
    margin-bottom: 5px;
    color: #767676
}
.migrate-org-roles {
    width: 100%;
    margin-top: -20px;
    display: table;
    border: 1px solid #ddd;
    border-radius: 3px
}
.migrate-org-roles .tooltipped:after {
    width: 150px;
    white-space: normal
}
.no-avatars-roles-matrix .migrate-org-roles {
    margin-top: 5px
}
.migrate-org-roles-item {
    display: table-cell;
    width: 33.33%;
    border-right: 1px solid #ddd
}
.migrate-org-roles-item:last-child {
    border-right: 0
}
.repo-creation-content {
    font-size: 14px;
    color: #767676
}
.repo-setting-check {
    margin-top: 1px;
    margin-right: 8px;
    margin-bottom: 15px;
    float: left
}
.repo-setting-saved,
.repo-setting-spinner {
    float: right
}
.repo-setting-saved {
    color: #6cc644;
    font-weight: bold
}
.migrate-ability-list {
    margin: 15px 0;
    list-style: none
}
.migrate-ability-list-item {
    padding-top: 5px;
    padding-bottom: 5px;
    margin: 0 20px;
    font-size: 14px
}
.migrate-ability-list-item:first-child {
    border-top: 0
}
.migrate-ability-list-item .octicon-check,
.migrate-ability-list-item .octicon-x {
    width: 15px
}
.migrate-ability-list-item .octicon-check {
    color: #6cc644
}
.migrate-ability-list-item .octicon-x {
    color: #aaa
}
.migrate-ability-list-item .octicon-question {
    color: #555;
    font-size: 12px
}
.migrate-org-roles-legacy-item {
    background-color: #f8f8f8
}
.migrate-ability-not-possible {
    color: #767676
}
.default-repository-permission .octicon-x,
.members-can-create-repositories .octicon-x,
.team-privacy .octicon-x {
    display: none
}
.default-repository-permission.migrate-ability-not-possible .octicon-x,
.members-can-create-repositories.migrate-ability-not-possible .octicon-x,
.team-privacy.migrate-ability-not-possible .octicon-x {
    display: inline-block
}
.default-repository-permission.migrate-ability-not-possible .octicon-check,
.members-can-create-repositories.migrate-ability-not-possible .octicon-check,
.team-privacy.migrate-ability-not-possible .octicon-check {
    display: none
}
.migrate-org-roles-header {
    padding: 15px 20px;
    border-bottom: 1px solid #ddd
}
.migrate-org-roles-title {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 18px;
    font-weight: normal
}
.migrate-org-roles-lead {
    margin-top: 4px;
    margin-bottom: 0;
    font-size: 14px;
    color: #767676
}
.migrate-org-badge {
    padding: 3px 5px;
    color: #333;
    font-size: 10px;
    letter-spacing: 1px;
    text-transform: uppercase;
    border: 1px solid #ddd;
    border-radius: 3px
}
.migrate-org-roles-count {
    padding: 10px 20px;
    color: #767676;
    border-top: 1px solid #ddd
}
.migrate-org-avatar-list {
    margin-top: 5px;
    margin-bottom: 10px
}
.migrate-org-avatar-list:before {
    display: table;
    content: ""
}
.migrate-org-avatar-list:after {
    display: table;
    clear: both;
    content: ""
}
.migrate-org-avatar-list .migrate-org-avatar,
.migrate-org-avatar-list .migrate-org-avatar-empty {
    float: left;
    margin-left: 2px
}
.migrate-org-avatar-list .migrate-org-avatar:first-child,
.migrate-org-avatar-list .migrate-org-avatar-empty:first-child {
    margin-left: 0
}
.migrate-org-avatar-list .migrate-org-avatar-empty {
    width: 30px;
    height: 30px;
    border-radius: 3px
}
.migrate-org-avatar-list .migrate-org-more-ellipsis,
.migrate-org-avatar-list .migrate-org-zero {
    font-size: 18px;
    text-align: center;
    line-height: 30px;
    color: #767676
}
.migrate-org-avatar-list .migrate-org-more-ellipsis {
    font-weight: bold;
    line-height: 20px;
    background-color: #f5f5f5
}
.migrate-org-avatar-list .migrate-org-zero {
    color: #767676;
    border: 1px dashed #ddd
}
.migrate-org-avatar-list .migrate-org-more-ellipsis:hover {
    text-decoration: none
}
.migrate-org-avatar-list .tooltipped:after {
    width: auto;
    white-space: nowrap
}
.default-permission-update-in-progress .form {
    opacity: 0.5;
    pointer-events: none
}
.default-permission-update-in-progress .spinner {
    margin-top: -3px;
    margin-left: 5px;
    display: inline-block;
    vertical-align: middle
}
.default-permission-updating {
    margin-top: 10px;
    margin-right: 10px;
    float: right
}
.default-permission-update-text {
    color: #767676
}
.org-settings-teams:before {
    display: table;
    content: ""
}
.org-settings-teams:after {
    display: table;
    clear: both;
    content: ""
}
.org-settings-team-item {
    width: 50%;
    float: left;
    padding: 20px 40px 25px;
    text-align: center
}
.org-settings-team-item:first-child {
    border-right: 1px solid #ddd
}
.org-settings-team-count {
    font-size: 30px;
    color: #000
}
.org-settings-team-type {
    margin-top: 0;
    margin-bottom: 10px;
    font-size: 14px;
    font-weight: normal;
    color: #000
}
.org-settings-team-description {
    margin-top: 0
}
.migrate-org-create-repos-wrapper .repo-setting-saved {
    display: none
}
.migrate-org-create-repos-wrapper.loading .spinner {
    display: inline-block
}
.migrate-org-create-repos-wrapper.success .repo-setting-saved {
    display: inline
}
.migrate-org-create-repos-wrapper .disabled {
    opacity: 0.5;
    pointer-events: none
}
.migrate-org-create-repos-wrapper .note {
    margin-top: 0;
    margin-bottom: 0
}
.legacy-contributor-list {
    margin-bottom: 15px
}
.legacy-contributor-actions {
    width: 395px;
    text-align: right
}
.legacy-contributor-button {
    display: inline
}
.legacy-contributor-button:last-child {
    margin-left: 5px
}
.legacy-contributor-cell,
.legacy-contributor-cell-contents {
    -webkit-transition: all 0.25s ease-in-out;
    transition: all 0.25s ease-in-out
}
.legacy-contributor-cell-contents {
    max-height: 80px
}
.hide-legacy-contributor .legacy-contributor-cell {
    padding-top: 0;
    padding-bottom: 0;
    opacity: 0
}
.hide-legacy-contributor .legacy-contributor-cell-contents {
    max-height: 0;
    overflow: hidden
}
.load-more-contributors {
    font-size: 14px;
    font-weight: bold
}
.migrate-back-step {
    margin-top: 9px;
    float: left
}
.team-grid {
    position: relative;
    margin-left: -10px;
    margin-right: -10px
}
.team-grid:before {
    display: table;
    content: ""
}
.team-grid:after {
    display: table;
    clear: both;
    content: ""
}
.team-grid .team {
    position: relative;
    float: left;
    width: 480px;
    padding: 15px;
    margin-bottom: 20px;
    margin-left: 10px;
    margin-right: 10px;
    border: 1px solid #eee;
    border-radius: 3px
}
.team-grid .team-name {
    display: block;
    margin-top: -2px;
    color: #333;
    font-size: 18px;
    font-weight: bold;
    text-decoration: none
}
.team-grid .team-name .css-truncate-target {
    max-width: 315px
}
.team-grid .team-name:focus,
.team-grid .team-name:hover {
    color: #4078c0
}
.team-grid .team-name:focus {
    outline: none
}
.team-grid .team-description {
    max-width: 80%;
    margin-top: 5px;
    font-size: 14px;
    color: #767676;
    text-overflow: ellipsis;
    white-space: nowrap
}
.team-grid .team-description .label-private {
    text-transform: uppercase
}
.team-grid .team-label-ldap {
    float: right
}
.team-grid .team-members {
    width: 478px;
    padding: 10px 15px;
    margin: 0 -15px -15px;
    border-top: 1px solid #eee;
    border-radius: 0 0 3px 3px;
    background-color: #f8f8f8
}
.team-grid .team-members .btn-sm {
    margin-top: 2px;
    margin-bottom: 2px
}
.team-grid .team-member {
    display: inline-block;
    vertical-align: top;
    width: 30px;
    height: 30px
}
.team-grid .team-member:hover {
    text-decoration: none
}
.team-grid .blankslate {
    margin-left: 10px;
    margin-right: 10px
}
.team-grid .team-actions-form {
    float: right
}
.team-label-ldap {
    display: inline-block;
    padding: 0 9px;
    line-height: 25px;
    border: 1px solid #eaeaea;
    border-radius: 3px;
    box-shadow: none;
    color: #767676;
    font-size: 11px;
    text-transform: uppercase;
    cursor: default
}
.team-label-ldap.header-label-ldap {
    padding: 3px 5px
}
.team-member-ellipsis {
    display: inline-block;
    vertical-align: top;
    width: 30px;
    height: 30px;
    line-height: 24px;
    color: #767676;
    font-weight: bold;
    text-align: center;
    background-color: #ddd;
    border-radius: 3px
}
.team-member-ellipsis:hover {
    color: #333;
    text-decoration: none
}
.typeahead-result {
    position: relative;
    display: block;
    min-width: 100%;
    padding: 10px;
    margin-top: 0;
    color: #333;
    cursor: pointer
}
.typeahead-result:before {
    display: table;
    content: ""
}
.typeahead-result:after {
    display: table;
    clear: both;
    content: ""
}
.typeahead-result:first-child {
    border-top: 0
}
.typeahead-result:focus,
.typeahead-result:hover,
.typeahead-result.navigation-focus {
    text-decoration: none
}
.typeahead-result:hover,
.typeahead-result.navigation-focus {
    color: #fff;
    background-color: #4078c0
}
.typeahead-result:hover .octicon-plus,
.typeahead-result.navigation-focus .octicon-plus {
    color: #fff
}
.member-suggestion {
    padding-left: 44px
}
.member-suggestion .avatar {
    float: left;
    margin-left: -34px;
    margin-right: 10px
}
.member-suggestion .member-suggestion-info {
    width: 75%;
    overflow: hidden;
    margin-top: 2px;
    margin-bottom: 0;
    white-space: nowrap;
    text-overflow: ellipsis
}
.member-suggestion .member-name {
    color: #767676;
    font-size: 12px
}
.member-suggestion .octicon-plus,
.member-suggestion .octicon-check {
    position: absolute;
    top: 50%;
    right: 15px;
    margin-top: -8px;
    color: #ddd
}
.member-suggestion .already-member-note,
.member-suggestion .non-member-note,
.member-suggestion .non-member-action {
    margin-top: 0;
    margin-bottom: 0;
    color: #767676;
    font-size: 11px
}
.member-suggestion .non-member-action {
    display: none
}
.member-suggestion:hover .member-name,
.member-suggestion:hover .non-member-note,
.member-suggestion:hover .already-member-note,
.member-suggestion:hover .non-member-action,
.member-suggestion.navigation-focus .member-name,
.member-suggestion.navigation-focus .non-member-note,
.member-suggestion.navigation-focus .already-member-note,
.member-suggestion.navigation-focus .non-member-action {
    color: #fff
}
.member-suggestion:hover .non-member-note,
.member-suggestion.navigation-focus .non-member-note {
    display: none
}
.member-suggestion:hover .non-member-action,
.member-suggestion.navigation-focus .non-member-action {
    display: block
}
.member-suggestion:hover .octicon-plus,
.member-suggestion:hover .octicon-check,
.member-suggestion.navigation-focus .octicon-plus,
.member-suggestion.navigation-focus .octicon-check {
    color: #fff
}
.member-suggestion.not-a-member .member-info,
.member-suggestion.disabled .member-info {
    margin-top: -2px
}
.member-suggestion.disabled {
    opacity: 0.5
}
.team-suggestion {
    padding-left: 32px
}
.team-suggestion .octicon {
    float: left;
    margin-left: -22px;
    margin-top: 2px
}
.team-suggestion .team-suggestion-info {
    margin: 2px 0 0
}
.team-suggestion .team-suggestion-info .css-truncate-target {
    max-width: none
}
.team-suggestion .team-size,
.team-suggestion .team-description {
    color: #767676;
    font-size: 12px
}
.team-suggestion.navigation-focus .team-size,
.team-suggestion.navigation-focus .team-description {
    color: #fff
}
.repo-access-add-team .team-name {
    font-size: 13px
}
.repo-access-add-team .team-description {
    display: block
}
.repo-access-add-team .team-size,
.repo-access-add-team .team-description {
    color: #767676;
    font-size: 12px
}
.repo-access-add-team.navigation-focus .team-size,
.repo-access-add-team.navigation-focus .team-description {
    color: #fff
}
.menu-item .org-avatar,
.menu-item .org-octicon-credit-card {
    position: absolute
}
.menu-item .org-octicon-credit-card {
    right: 0
}
.org-settings-link {
    display: block;
    padding: 0 30px;
    word-wrap: break-word
}
.team-info-card {
    position: relative;
    margin-bottom: 20px
}
.team-info-card .team-label-ldap {
    font-size: 13px;
    line-height: 32px
}
.team-info-card .team-description {
    margin-top: 10px;
    color: #666;
    font-size: 14px;
    line-height: 20px;
    word-break: break-word
}
.team-info-card .team-description .link {
    color: #767676;
    cursor: pointer
}
.team-info-card .team-description .link:hover {
    text-decoration: underline
}
.team-info-card .description-toggler .turn-on {
    display: inline-block
}
.team-info-card .description-toggler.on .turn-off {
    display: inline-block
}
.team-info-card .description-toggler .turn-off {
    display: none
}
.team-info-card .description-toggler.on .turn-on {
    display: none
}
.team-title {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 22px;
    line-height: 26px
}
.team-stats {
    margin-right: -15px;
    margin-bottom: -10px;
    margin-left: -15px;
    padding-right: 15px;
    padding-left: 15px;
    border-top: 1px solid #eee
}
.stats-group {
    display: table;
    table-layout: fixed;
    width: 100%
}
.stats-group-stat {
    display: table-cell;
    padding-top: 10px;
    padding-bottom: 10px;
    padding-left: 15px;
    font-size: 12px;
    color: #767676;
    text-transform: uppercase
}
.stats-group-stat:first-child {
    padding-left: 0;
    border-right: 1px solid #eee
}
.stats-group-stat:hover,
.stats-group-stat:hover .stat-number {
    color: #4078c0;
    text-decoration: none
}
.stats-group-stat.no-link:hover {
    color: #767676;
    text-decoration: none
}
.stats-group-stat.no-link:hover .stat-number {
    color: #333
}
.stat-number {
    display: block;
    color: #333;
    font-size: 16px
}
.team-description-form {
    width: 100%;
    margin-top: 10px;
    margin-bottom: 20px
}
.team-description-field {
    width: 100%;
    height: 100px;
    margin-bottom: 10px;
    font-size: 14px
}
.team-actions .octicon {
    margin-right: 0
}
.team-actions-form {
    display: inline-block
}
.org-team-sidebar {
    float: left;
    width: 280px
}
.org-team-sidebar .team-note {
    color: #767676;
    font-size: 13px;
    text-align: center
}
.org-team-sidebar .team-note .note-emphasis {
    color: #333
}
.org-team-main {
    float: right;
    width: 660px
}
.permission-title {
    margin-top: 0;
    margin-bottom: 0
}
.owners-member-title {
    margin-top: 10px;
    margin-bottom: 0;
    font-size: 18px;
    color: #767676;
    font-weight: normal
}
.owners-notice {
    background-color: #f0f8ff
}
.owners-notice-title {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 18px
}
.owners-notice-link {
    margin-top: 10px
}
.owners-team-repo-note {
    margin-top: 12px;
    margin-bottom: 0
}
.owners-team-repo-note .octicon {
    font-size: 14px
}
.team-member-list {
    list-style: none
}
.team-member-list .table-list-cell {
    padding-top: 15px;
    padding-bottom: 15px
}
.team-member-list .team-member-content {
    margin-left: 50px
}
.team-member-list .team-member-username {
    font-weight: bold;
    margin: 0;
    font-size: 14px;
    line-height: 20px
}
.team-member-list .team-member-description {
    color: #767676;
    margin: 0;
    font-size: 14px;
    line-height: 20px
}
.team-member-list .label-admin,
.team-member-list .label-generic {
    cursor: default
}
.team-member-list .manage-team-member {
    float: right
}
.team-member-list .manage-team-member .select-menu-modal {
    width: 225px;
    left: -176px
}
.team-member-list .manage-team-member .select-menu-item.disabled {
    color: #bbb;
    cursor: not-allowed
}
.team-member-list .manage-team-member .select-menu-item .btn-link {
    width: 100%;
    margin-left: 0;
    color: #767676
}
.team-member-list .manage-team-member .select-menu-item .btn-danger {
    color: #bd2c00
}
.team-member-list .manage-team-member .navigation-focus.disabled {
    color: #bbb;
    background-color: #fff
}
.team-member-list .manage-team-member .navigation-focus .btn-link {
    color: #fff;
    text-decoration: none
}
.team-member-list .manage-team-member .navigation-focus .btn-danger {
    background: #bd2c00
}
.team-member-list-avatar {
    float: left;
    margin-right: 10px
}
.loading .legacy-contributor-button .btn-sm {
    opacity: 0.5;
    pointer-events: none
}
.org-team-form {
    width: 440px;
    margin: 0 auto
}
.org-team-form .disabled {
    opacity: 0.5
}
.org-validate-group {
    position: relative
}
.org-validate-group .octicon,
.org-validate-group .spinner {
    position: absolute;
    top: 9px;
    right: 10px
}
.org-validate-group .octicon-check {
    color: #6cc644
}
.org-validate-group .octicon-alert {
    color: #bd2c00
}
.team-members {
    margin-bottom: 20px
}
.confirm-removal-container .private-fork-count {
    margin-top: 0;
    color: #767676;
    font-weight: normal;
    font-size: 12px
}
.confirm-removal-container .deleting-private-forks-warning {
    position: relative;
    padding-left: 26px
}
.confirm-removal-container .deleting-private-forks-warning .octicon {
    position: absolute;
    top: 2px;
    left: 0;
    color: #bd2c00
}
.confirm-removal-list-container {
    border: 1px solid #eaeaea;
    border-radius: 3px;
    margin-bottom: 15px
}
.facebox .confirm-removal-list {
    margin-left: 0;
    margin-bottom: 0;
    padding-left: 0;
    max-height: 182px;
    overflow: auto
}
.confirm-removal-list-item {
    font-size: 14px;
    font-weight: bold;
    margin: 0;
    padding: 10px;
    border-top: 1px solid #eaeaea
}
.confirm-removal-list-item:first-child {
    border-top: 0
}
.confirm-removal-list-item.cutoff-member-summary {
    font-weight: normal
}
.confirm-removal-team .octicon,
.confirm-removal-repo .octicon {
    margin-right: 3px;
    color: #767676
}
#convert-to-outside-collaborator {
    display: none
}
.org-blankslate {
    display: none
}
.org-section.is-empty .org-blankslate {
    display: block
}
.manage-user-info {
    margin-left: -15px;
    margin-right: -15px;
    padding-right: 15px;
    padding-bottom: 10px;
    padding-left: 15px;
    border-bottom: 1px solid #eee
}
.manage-user-info:before {
    display: table;
    content: ""
}
.manage-user-info:after {
    display: table;
    clear: both;
    content: ""
}
.manage-user-info .member-username {
    margin-top: 0
}
.manage-user-info .member-username,
.manage-user-info .member-fullname {
    display: block;
    overflow-x: hidden;
    white-space: nowrap;
    text-overflow: ellipsis
}
.manage-user-info .avatar {
    margin-top: 2px;
    margin-right: 10px
}
.manage-user-role {
    position: relative;
    padding-top: 15px;
    padding-bottom: 5px
}
.manage-user-role .select-menu-item-text .description {
    font-size: 12px;
    line-height: 16px
}
.manage-user-role .non-member-info {
    color: #767676
}
.manage-member-meta {
    list-style: none
}
.manage-member-meta-item {
    margin-top: 12px;
    color: #767676
}
.manage-member-meta-item:first-child {
    margin-top: 0
}
.manage-member-meta-item .btn-link {
    color: #767676
}
.manage-member-meta-item>.octicon {
    width: 14px;
    text-align: center;
    margin-right: 5px;
    color: #767676
}
.manage-member-meta-item>.octicon-alert {
    color: #c9510c
}
.member-two-factor-disabled {
    color: #bd2c00
}
.manage-member-button {
    margin-bottom: 10px
}
.org-person-repo-header {
    margin-top: 0
}
.org-person-repo-search {
    margin-top: 5px;
    margin-right: 5px
}
.org-user-notice-title {
    margin-top: 0;
    margin-bottom: 0
}
.org-user-notice-content {
    margin-top: 10px;
    margin-bottom: 10px;
    font-size: 14px
}
.org-user-notice-content strong {
    color: #333
}
.org-user-notice-content:last-child {
    margin-bottom: 0
}
.org-user-notice-content .octicon {
    color: #767676
}
.org-user-notice-icon {
    margin: 10px 10px 20px;
    float: right;
    font-size: 45px;
    color: #ccc
}
.org-migration-list {
    margin-left: 20px;
    margin-bottom: 20px;
    font-size: 14px
}
.org-migration-list-item {
    margin-bottom: 5px
}
.org-migration-actions:before {
    display: table;
    content: ""
}
.org-migration-actions:after {
    display: table;
    clear: both;
    content: ""
}
.org-migration-actions form {
    margin-right: 10px;
    float: left
}
.manage-repo-access-header {
    margin-top: 30px;
    margin-bottom: 30px
}
.manage-repo-access-header:before {
    display: table;
    content: ""
}
.manage-repo-access-header:after {
    display: table;
    clear: both;
    content: ""
}
.manage-repo-access-header .btn {
    margin-top: 8px
}
.manage-repo-access-header .tooltipped:after {
    width: 250px;
    white-space: normal
}
.manage-repo-access-heading {
    margin-top: -2px;
    margin-bottom: 0;
    font-weight: normal;
    font-size: 24px
}
.manage-repo-access-lead {
    margin-top: 3px;
    margin-bottom: 0;
    font-size: 16px;
    color: #767676
}
.manage-repo-access-group {
    background-color: #fff;
    border: 1px solid #ddd;
    border-radius: 3px
}
.manage-repo-access-title {
    margin-top: 0;
    margin-bottom: 0;
    padding: 12px 15px;
    font-size: 14px;
    border-bottom: 1px solid #ddd;
    border-radius: 3px 3px 0 0;
    background-color: #f8f8f8
}
.manage-repo-access-wrapper {
    position: relative;
    padding-left: 25px
}
.manage-repo-access-wrapper:before {
    position: absolute;
    top: 15px;
    bottom: 15px;
    left: 20px;
    z-index: 1;
    width: 2px;
    display: block;
    content: "";
    background-color: #eee
}
.manage-repo-access-icon {
    position: relative;
    margin-top: -3px;
    margin-left: -25px;
    padding-top: 2px;
    padding-bottom: 2px;
    z-index: 2;
    float: left;
    background: #fff
}
.manage-repo-access-icon .octicon {
    color: #ccc;
    font-size: 14px
}
.manage-repo-access-list {
    list-style: none
}
.manage-repo-access-list-item {
    padding: 15px
}
.manage-repo-access-list-item:last-child {
    border-bottom: 0;
    border-radius: 0 0 3px 3px
}
.manage-repo-access-teams-group {
    margin-top: -20px;
    list-style: none;
    border: 1px solid #ddd;
    border-radius: 3px
}
.manage-repo-access-team-item {
    border-top: 1px solid #eee
}
.manage-repo-access-team-item:first-child {
    border-top: 0
}
.manage-repo-access-description {
    margin-top: 3px;
    margin-bottom: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    word-wrap: break-word;
    white-space: nowrap
}
.manage-repo-access-not-active {
    color: #333;
    background-color: #fafafa
}
.manage-repo-access-not-active .manage-repo-access-icon {
    background: #f9f9f9
}
.manage-access-remove-footer {
    padding: 15px;
    border-top: 1px solid #ddd
}
.manage-access-remove-footer .tooltipped:after {
    width: 250px;
    white-space: normal
}
.manage-access-none {
    margin: 20px 50px;
    text-align: center
}
.ldap-group-dn {
    display: block;
    color: #aaa;
    font-weight: normal
}
.ldap-import-groups-container .blankslate {
    display: none
}
.ldap-import-groups-container.is-empty .blankslate {
    display: block
}
.ldap-import-groups-container.is-empty .ldap-memberships-list {
    display: none
}
.ldap-import-groups-container .team-name-exists {
    display: none
}
.ldap-import-groups-container .is-exists .ldap-mention-as {
    color: #bd2c00
}
.ldap-import-groups-container .is-exists .team-name-exists {
    position: absolute;
    z-index: 1;
    display: inline-block;
    padding: 5px;
    font-size: 11px;
    color: #494620;
    background: #f7ea57;
    border: 1px solid #c0b536;
    border-top-color: #fff;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.ldap-memberships-list {
    margin-bottom: 30px
}
.ldap-memberships-list .table-list-cell {
    font-size: 13px;
    padding-bottom: 10px;
    padding-top: 10px;
    vertical-align: middle
}
.ldap-memberships-list .table-list-cell:last-child {
    width: 92px
}
.ldap-memberships-list .team-name-exists {
    bottom: -19px;
    left: 10px
}
.ldap-memberships-list .ldap-list-team-name {
    width: 380px
}
.ldap-memberships-list .ldap-group-dn {
    font-size: 11px
}
.ldap-memberships-list .ldap-mention-as {
    width: 260px
}
.ldap-memberships-list .edit {
    position: absolute;
    padding: 10px;
    margin-left: -33px;
    color: #4078c0;
    cursor: pointer
}
.ldap-memberships-list .edit-fields {
    display: none
}
.ldap-memberships-list .is-editing .edit-hide {
    display: none
}
.ldap-memberships-list .is-editing .edit-fields {
    display: block
}
.ldap-memberships-list .is-editing .spinner {
    margin-left: 15px;
    vertical-align: middle
}
.ldap-memberships-list .is-removing {
    opacity: 0.25
}
.ldap-memberships-list .is-removing .edit {
    opacity: 0.5
}
.team-name-field {
    height: 33px
}
.ldap-import-form-actions {
    margin-top: 30px
}
.is-importing .team-ldap-group-adder-button .spinner {
    display: inline;
    float: left
}
.team-ldap-group-adder {
    position: relative;
    float: left
}
.team-ldap-group-adder .team-name-exists {
    bottom: -27px
}
.team-ldap-group-adder .subnav-search-input {
    border-radius: 4px 0 0 4px
}
.team-ldap-group-adder-button {
    border-radius: 0 4px 4px 0;
    margin-left: -1px;
    width: 90px
}
.team-ldap-group-adder-button .loading-indicator {
    display: none
}
#pending-invitations {
    display: none
}
.pending-team-invitations-link {
    margin-top: 20px;
    padding-top: 15px;
    padding-bottom: 15px;
    display: block;
    border-top: 1px solid #eee
}
.invited .team-member-list {
    margin: -10px 0 0
}
.invited .team-member-list .list-item {
    padding: 10px 0;
    border-bottom: 1px solid #eee
}
.invited .team-member-list .list-item:before {
    display: table;
    content: ""
}
.invited .team-member-list .list-item:after {
    display: table;
    clear: both;
    content: ""
}
.invited .team-member-list .list-item:last-of-type {
    border: 0
}
.invited .team-member-list .list-item .edit-invitation,
.invited .team-member-list .list-item .cancel-invitation {
    margin-top: 6px;
    float: right
}
.invited-banner {
    margin-top: 10px;
    padding: 10px;
    border: 1px solid #eaeaea;
    border-radius: 4px;
    background-color: #fff
}
.invited-banner:before {
    display: table;
    content: ""
}
.invited-banner:after {
    display: table;
    clear: both;
    content: ""
}
.invited-banner .btn-sm {
    float: right;
    margin-left: 5px;
    margin-top: -3px
}
.invited-banner p {
    color: #333;
    margin: 0;
    font-size: 15px
}
.invited-banner .inviter-link {
    color: #333;
    font-weight: bold
}
.invitation-container {
    width: 600px;
    margin: 40px auto;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.invitation-container h3 {
    font-size: 16px;
    font-weight: normal
}
.invitation-disclosure {
    position: relative;
    padding: 10px 0 10px 24px;
    list-style: none;
    color: #767676;
    text-align: center
}
.invitation-disclosure .octicon {
    color: #767676;
    text-align: center;
    display: inline-block;
    margin-right: 5px
}
.invitation-header {
    position: relative;
    text-align: center
}
.invitation-header .avatar {
    margin-bottom: 20px
}
.invitation-header .invitation-title {
    font-size: 18px;
    font-weight: normal;
    line-height: 16px;
    margin: 0
}
.invitation-header .inviter {
    font-size: 13px;
    color: #767676;
    margin: 5px 0 10px
}
.invitation-footer {
    margin: 40px 0 20px
}
.invitation-footer form {
    display: inline-block;
    margin-right: 10px
}
.org-user-block-input {
    width: 275px
}
.outline-box-group {
    border-radius: 3px
}
.outline-box {
    padding: 20px;
    border: solid 1px #d8d8d8
}
.outline-box:first-child {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.outline-box:last-child {
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px
}
.outline-box+.outline-box {
    border-top: 0
}
.outline-box-highlighted {
    background-color: #f7fafd;
    border-color: #c9d6e3
}
.owner-select-grid {
    margin-left: -8px
}
.owner-select-grid:before {
    display: table;
    content: ""
}
.owner-select-grid:after {
    display: table;
    clear: both;
    content: ""
}
.owner-select-target {
    float: left;
    padding: 10px;
    margin: 0 10px 20px;
    text-align: center;
    background-color: #f2f2f2;
    border-radius: 3px;
    border: 0;
    font-weight: bold;
    cursor: pointer
}
.owner-select-target:hover,
.owner-select-target:focus {
    color: #fff;
    background-color: #4078c0
}
.owner-select-target:active {
    color: #fff;
    background-color: #33609a
}
.owner-select-target .css-truncate-target {
    max-width: 90px
}
.owner-select-target.disabled {
    cursor: not-allowed;
    color: #999
}
.owner-select-target.disabled .user-mention {
    color: #999
}
.owner-select-target.disabled .owner-select-avatar {
    opacity: 0.3
}
.owner-select-avatar {
    display: block;
    margin-bottom: 9px
}
.page-notice {
    margin: 15px auto;
    width: 400px;
    padding: 20px;
    color: #333;
    font-size: 14px;
    background: #fffeeb;
    border: 1px solid #ddd;
    border-radius: 3px
}
.page-notice h2 {
    margin: 0;
    font-size: 16px;
    color: #000
}
.page-notice p:last-child {
    margin-bottom: 0
}
#editor-body-buffer {
    display: none
}
#pages-composer {
    padding-bottom: 5px;
    margin-bottom: 20px;
    border-bottom: 1px solid #eee
}
#pages-composer label {
    display: inline-block;
    margin-bottom: 10px;
    font-size: 16px
}
#pages-composer input {
    margin-bottom: 15px
}
#pages-composer p {
    margin-top: -10px;
    margin-bottom: 10px;
    color: #767676
}
#pages-composer #gollum-editor-function-bar {
    margin-top: 0
}
#pages-composer #gollum-editor {
    margin: 0;
    padding: 0;
    border: 0
}
#pages-composer #gollum-editor-body {
    margin-top: 10px
}
.gollum-readme {
    display: inline-block;
    margin-left: 10px
}
#gollum-editor-function-bar #undo-load-readme {
    display: none
}
.theme-picker {
    margin-bottom: -1px;
    background-color: #fff;
    background-clip: padding-box;
    border-bottom: 1px solid rgba(0, 0, 0, 0.1);
    box-shadow: 0 5px 10px rgba(0, 0, 0, 0.1)
}
.theme-picker>.container {
    position: relative;
    overflow: hidden;
    text-align: center
}
.theme-picker-thumbs {
    border-bottom: 1px solid #eee
}
.theme-picker-footer {
    position: relative;
    padding-bottom: 15px
}
.theme-toggle {
    width: 32px;
    height: 32px;
    color: #ccc;
    padding: 0;
    background: none;
    border: 0
}
.theme-toggle:hover {
    color: #0084c8;
    text-decoration: none
}
.theme-toggle.disabled,
.theme-toggle.disabled:hover {
    color: #ccc;
    opacity: 0.3;
    cursor: not-allowed
}
.theme-toggle-full-left,
.theme-toggle-full-right {
    position: absolute;
    top: 50px;
    width: 32px;
    height: 32px;
    overflow: hidden
}
.theme-toggle-full-left {
    left: 0
}
.theme-toggle-full-right {
    right: 0
}
.theme-selector {
    height: 102px;
    margin: 15px 46px;
    white-space: nowrap;
    overflow: hidden
}
.theme-selector-thumbnail {
    display: inline-block;
    padding: 2px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.theme-selector-thumbnail+.theme-selector-thumbnail {
    margin-left: 15px
}
.theme-selector-thumbnail:hover {
    text-decoration: none;
    background-color: #f5f5f5
}
.theme-selector-thumbnail.selected {
    padding: 3px;
    background-color: #4078c0;
    border: 0
}
.theme-selector-thumbnail.selected .theme-selector-img {
    border: 1px solid #fff
}
.theme-selector-img {
    display: block;
    width: 126px;
    height: 96px;
    border-radius: 1px
}
.theme-selector-name {
    display: none
}
.theme-picker-spinner {
    position: absolute;
    top: 16px;
    left: 50%;
    margin-left: -16px;
    opacity: 0;
    -webkit-transition: 0.2s, opacity ease-in-out;
    transition: 0.2s, opacity ease-in-out;
    background-color: #fff
}
.theme-picker-spinner.visible {
    opacity: 1
}
.theme-picker-spinner.visible ~ .theme-picker-controls .theme-name {
    opacity: 0
}
.theme-selector-actions {
    padding-top: 15px;
    text-align: right
}
.theme-selector-actions:before {
    display: table;
    content: ""
}
.theme-selector-actions:after {
    display: table;
    clear: both;
    content: ""
}
.theme-selector-actions .page-edit,
.theme-selector-actions .page-publish {
    display: inline-block;
    margin-left: 5px
}
.theme-picker-view-toggle {
    float: left
}
.theme-picker-view-toggle .for-hiding {
    display: none
}
.theme-picker-view-toggle.open .for-hiding {
    display: inline
}
.theme-picker-view-toggle.open .for-showing {
    display: none
}
.theme-picker-controls {
    position: absolute;
    top: 15px;
    left: 50%;
    display: none;
    width: 220px;
    margin-left: -110px;
    line-height: 34px;
    text-align: center
}
.theme-picker-controls .theme-toggle {
    vertical-align: middle
}
.theme-name {
    display: inline-block;
    margin-left: 10px;
    margin-right: 10px;
    font-size: 20px;
    vertical-align: middle;
    line-height: 1
}
#page-preview {
    position: relative;
    z-index: -100;
    display: block;
    width: 100%;
    height: 100%;
    padding: 0;
    background-color: #fff;
    border: 0
}
.feed-icon a {
    display: block;
    width: 18px;
    height: 18px;
    background: #f37538;
    color: #fff;
    border-radius: 3px;
    padding: 1px;
    text-align: center
}
body.page-profile .select-menu-modal {
    width: 130px
}
body.page-profile .select-menu-modal-holder {
    right: 0
}
body.page-profile .tab-content {
    position: relative
}
body.page-profile .feed-icon {
    position: absolute;
    right: 0;
    z-index: 2
}
body.page-profile .user-actions {
    display: inline-block;
    position: relative;
    margin-left: 10px
}
.profilecols .orgs h3 {
    margin: 0 0 5px;
    font-size: 12px
}
.profilecols .orgs h3 a {
    font-weight: normal;
    margin-left: 5px
}
.profilecols .repo-search {
    display: inline
}
.profilecols .filter-bar {
    position: relative;
    padding: 0 0 15px;
    background-color: #fff;
    border-bottom: 1px solid #ddd
}
.profilecols .filter-bar .new-repo {
    float: right;
    margin-left: 15px
}
.profilecols .filter-bar .filter_input {
    width: 260px
}
.profilecols .filter-bar .repo_filterer {
    float: right;
    margin-top: 8px
}
.profilecols .filter-bar li {
    position: relative;
    float: right;
    list-style: none;
    margin-left: 10px;
    font-size: 14px
}
.profilecols .filter-bar li a {
    display: inline-block
}
.profilecols .filter-bar li .filter-selected {
    color: #000;
    font-weight: bold
}
.profilecols .blankslate {
    margin-top: 30px
}
.vcard-avatar {
    position: relative;
    display: block
}
.vcard-avatar .avatar {
    border-radius: 6px
}
.vcard-names {
    margin-top: 5px;
    line-height: 1
}
.vcard-fullname {
    display: block;
    overflow: hidden;
    width: 100%;
    font-size: 26px;
    line-height: 30px;
    text-overflow: ellipsis
}
.vcard-username {
    display: block;
    overflow: hidden;
    width: 100%;
    font-size: 20px;
    font-style: normal;
    font-weight: 300;
    line-height: 24px;
    color: #666;
    text-overflow: ellipsis
}
.vcard-details {
    list-style: none;
    padding-top: 15px;
    padding-bottom: 15px;
    border-top: 1px solid #eee
}
.vcard-detail {
    width: 100%;
    padding: 2px 0 2px 24px;
    overflow-x: hidden;
    white-space: nowrap;
    font-size: 14px;
    text-overflow: ellipsis
}
.vcard-detail .octicon {
    float: left;
    width: 16px;
    text-align: center;
    margin-left: -24px;
    color: #ccc
}
.vcard .staff-badge {
    position: relative;
    top: -1px;
    padding: 2px 5px;
    font-size: 10px;
    font-weight: bold;
    color: #fff;
    text-transform: uppercase;
    background-color: #4078c0;
    border-radius: 3px
}
.member-badge {
    display: block;
    padding: 10px 0 8px 24px;
    font-size: 14px;
    color: #4078c0;
    border-top: 1px solid #eee
}
.member-badge .octicon {
    float: left;
    width: 16px;
    margin-left: -24px;
    color: #ccc;
    text-align: center
}
.member-badge+.member-badge {
    padding-top: 0;
    margin-top: -3px;
    border-top: 0
}
.vcard-stats {
    margin-bottom: 10px;
    padding-top: 15px;
    padding-bottom: 15px;
    text-align: center;
    border-top: 1px solid #eee;
    border-bottom: 1px solid #eee
}
.vcard-stats:before {
    display: table;
    content: ""
}
.vcard-stats:after {
    display: table;
    clear: both;
    content: ""
}
.vcard-stat {
    float: left;
    width: 33.333%;
    font-size: 11px;
    text-transform: capitalize
}
.vcard-stat-count {
    display: block;
    font-size: 28px;
    font-weight: bold;
    line-height: 1
}
.vcard-stat:hover {
    text-decoration: none
}
.vcard-stat:hover .text-muted {
    color: inherit
}
.new-user-avatar-cta {
    padding: 8px 16px;
    margin-bottom: 20px;
    background-color: #f1f6fb;
    color: #244f79;
    border: solid 1px #d0e5f8;
    border-radius: 3px;
    font-size: 14px
}
.new-user-avatar-cta .btn-sm {
    float: right;
    margin-left: 30px
}
.btn-block-user {
    color: inherit
}
.btn-block-user:hover {
    text-decoration: none
}
.steps {
    width: 100%;
    display: table;
    margin: 30px auto 0;
    padding: 0;
    overflow: hidden;
    list-style: none;
    border: 1px solid #ddd;
    border-radius: 3px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05)
}
.steps li {
    display: table-cell;
    width: 33.3%;
    padding: 10px 15px;
    color: #ccc;
    cursor: default;
    border-left: 1px solid #ddd;
    background-color: #fafafa
}
.steps li.current {
    background-color: #fff;
    color: #333
}
.steps li.current .mega-octicon {
    color: #4078c0
}
.steps li .mega-octicon {
    float: left;
    margin-right: 15px;
    padding-bottom: 5px
}
.steps li .step {
    display: block
}
.steps li:first-child {
    border-left: 0
}
.steps .complete {
    color: #767676
}
.steps .complete .mega-octicon {
    color: #6cc644
}
.prose-diff .anchor {
    display: none
}
.prose-diff .show-rich-diff {
    cursor: pointer;
    color: #4183c4;
    text-decoration: none
}
.prose-diff .show-rich-diff:hover {
    text-decoration: underline
}
.prose-diff.collapsed .rich-diff-level-zero.expandable {
    cursor: pointer
}
.prose-diff.collapsed .rich-diff-level-zero.expandable .vicinity {
    display: block
}
.prose-diff.collapsed .rich-diff-level-zero.expandable .unchanged:not(.vicinity) {
    display: none
}
.prose-diff.collapsed .rich-diff-level-zero.expandable:first-child::before {
    margin-top: 1em
}
.prose-diff.collapsed .rich-diff-level-zero.expandable:before {
    font-family: "octicons";
    content: "\f039";
    color: #d3d3d3;
    display: block;
    text-align: center;
    font-size: 24px;
    letter-spacing: 2px;
    line-height: 0;
    margin-top: 1em;
    margin-bottom: 1em;
    padding: 0
}
.prose-diff.collapsed .rich-diff-level-zero.expandable:hover:before {
    color: #000
}
.prose-diff.collapsed .rich-diff-level-zero.expandable:only-child:before {
    content: "Sorry, no visible changes to display.";
    color: #d3d3d3;
    font-size: 18px
}
.prose-diff.collapsed .rich-diff-level-zero.expandable:only-child:hover:before {
    color: #000
}
.prose-diff.collapsed .rich-diff-level-zero.expandable>.removed,
.prose-diff.collapsed .rich-diff-level-zero.expandable>del {
    text-decoration: none;
    display: none
}
.prose-diff .markdown-body {
    padding: 30px;
    padding-left: 15px
}
.prose-diff .markdown-body>ins {
    box-shadow: inset 4px 0 0 #7fcb5c
}
.prose-diff .markdown-body>del {
    box-shadow: inset 4px 0 0 #c94114;
    text-decoration: none
}
.prose-diff .markdown-body>ins,
.prose-diff .markdown-body>del {
    display: block;
    border-radius: 0
}
.prose-diff .markdown-body>ins>.rich-diff-level-zero,
.prose-diff .markdown-body>ins>.rich-diff-level-one,
.prose-diff .markdown-body>del>.rich-diff-level-zero,
.prose-diff .markdown-body>del>.rich-diff-level-one {
    margin-left: 15px
}
.prose-diff .markdown-body>ins:first-child *,
.prose-diff .markdown-body>del:first-child * {
    margin-top: 0
}
.prose-diff .rich-diff-level-zero.added {
    box-shadow: inset 4px 0 0 #7fcb5c
}
.prose-diff .rich-diff-level-zero.removed {
    box-shadow: inset 4px 0 0 #c94114
}
.prose-diff .rich-diff-level-zero.changed {
    box-shadow: inset 4px 0 0 #ffc045
}
.prose-diff .rich-diff-level-zero.unchanged,
.prose-diff .rich-diff-level-zero.vicinity {
    margin-left: 15px
}
.prose-diff .rich-diff-level-zero.added,
.prose-diff .rich-diff-level-zero.removed,
.prose-diff .rich-diff-level-zero.changed {
    display: block;
    border-radius: 0
}
.prose-diff .rich-diff-level-zero.added>.rich-diff-level-one,
.prose-diff .rich-diff-level-zero.removed>.rich-diff-level-one,
.prose-diff .rich-diff-level-zero.changed>.rich-diff-level-one {
    margin-left: 15px
}
.prose-diff .rich-diff-level-zero.added:first-child *,
.prose-diff .rich-diff-level-zero.removed:first-child *,
.prose-diff .rich-diff-level-zero.changed:first-child * {
    margin-top: 0
}
.prose-diff:not(.changed)>:not(.github-user-ins):not(.github-user-del)>.removed,
.prose-diff:not(.changed)>:not(.github-user-ins):not(.github-user-del)>del {
    text-decoration: none
}
.prose-diff .changed del,
.prose-diff .changed del pre,
.prose-diff .changed del code,
.prose-diff .changed del>div,
.prose-diff .changed .removed,
.prose-diff .changed .removed pre,
.prose-diff .changed .removed code,
.prose-diff .changed .removed>div {
    text-decoration: line-through;
    color: #a33;
    background: #ffeaea
}
.prose-diff .changed ins,
.prose-diff .changed ins code,
.prose-diff .changed ins pre,
.prose-diff .changed .added {
    background: #eaffea;
    border-bottom: 1px solid MediumSeaGreen
}
.prose-diff>.markdown-body .github-user-ins {
    text-decoration: underline
}
.prose-diff>.markdown-body .github-user-del {
    text-decoration: line-through
}
.prose-diff>.markdown-body li ul.added {
    background: #eaffea
}
.prose-diff>.markdown-body li ul.removed {
    color: #a33;
    background: #ffeaea
}
.prose-diff>.markdown-body li ul.removed:not(.github-user-ins) {
    text-decoration: line-through
}
.prose-diff>.markdown-body li.added.moved-up:before {
    font-family: "octicons";
    content: "\f03d ";
    color: #d3d3d3
}
.prose-diff>.markdown-body li.added.moved-down:before {
    font-family: "octicons";
    content: "\f03f ";
    color: #d3d3d3
}
.prose-diff>.markdown-body li.added.moved {
    background: #ffffea
}
.prose-diff>.markdown-body li.removed.moved {
    display: none
}
.prose-diff>.markdown-body pre {
    padding: 10px 20px
}
.prose-diff>.markdown-body th.changed,
.prose-diff>.markdown-body td.changed {
    border-left-color: #ddd;
    background: #ffffea
}
.prose-diff>.markdown-body:not(li.moved).removed {
    color: #a33;
    text-decoration: line-through;
    background: #ffeaea
}
.prose-diff>.markdown-body:not(.github-user-ins):not(li.moved).removed {
    text-decoration: line-through
}
.prose-diff>.markdown-body:not(li.moved).added,
.prose-diff>.markdown-body li:not(.moved).added {
    background: #eaffea
}
.prose-diff>.markdown-body:not(.github-user-del):not(li.moved).added li:not(.moved):not(.github-user-del).added {
    text-decoration: none
}
.prose-diff>.markdown-body li:not(.moved).removed {
    color: #a33;
    background: #ffeaea
}
.prose-diff>.markdown-body li:not(.moved):not(.github-user-ins).removed {
    text-decoration: line-through
}
.prose-diff>.markdown-body .added,
.prose-diff>.markdown-body ins+.added,
.prose-diff>.markdown-body ins {
    border-bottom: 0;
    border-top: 0
}
.prose-diff>.markdown-body .added:not(.github-user-del):not(.github-user-ins),
.prose-diff>.markdown-body ins+.added:not(.github-user-del):not(.github-user-ins),
.prose-diff>.markdown-body ins:not(.github-user-del):not(.github-user-ins) {
    text-decoration: none
}
.prose-diff>.markdown-body img.added,
.prose-diff>.markdown-body img.removed {
    border-width: 1px;
    border-style: solid
}
.prose-diff>.markdown-body ins pre:not(.github-user-del):not(.github-user-ins),
.prose-diff>.markdown-body ins code:not(.github-user-del):not(.github-user-ins),
.prose-diff>.markdown-body ins>div:not(.github-user-del):not(.github-user-ins) {
    text-decoration: none
}
.prose-diff>.markdown-body ul>ins,
.prose-diff>.markdown-body ul>del {
    display: block;
    padding: 0
}
.prose-diff>.markdown-body .added>li,
.prose-diff>.markdown-body .removed>li {
    margin-top: 0;
    margin-bottom: 0
}
span.changed_tag,
em.changed_tag,
strong.changed_tag,
b.changed_tag,
i.changed_tag,
code.changed_tag {
    border-bottom: 1px dotted #808080;
    border-radius: 0
}
a.added_href,
a.changed_href,
span.removed_href {
    border-bottom: 1px dotted #808080;
    border-radius: 0
}
.diff-view .file-type-prose .rich-diff {
    display: none
}
.diff-view .file-type-prose.display-rich-diff .rich-diff {
    display: block
}
.diff-view .file-type-prose.display-rich-diff .file-diff {
    display: none
}
.pull-request-tab-content {
    display: none
}
.pull-request-tab-content.is-visible {
    display: block
}
.discussion-timeline p.explain {
    margin: 0;
    font-size: 12px
}
.pull-request-ref-restore {
    display: none
}
.pull-request-ref-restore .animated-ellipsis-container {
    line-height: 16px
}
.pull-request-ref-restore-text {
    display: block
}
.pull-discussion-timeline.is-pull-restorable .pull-request-ref-restore.last {
    display: block
}
.signed-out-comment {
    margin-top: 15px;
    margin-left: 64px;
    padding: 15px;
    background-color: #fff9ea;
    border: solid 1px #dfd8c2;
    border-radius: 3px
}
.signed-out-comment .btn {
    vertical-align: baseline;
    margin-right: 3px
}
.inline-comment-form .signed-out-comment {
    margin: 5px;
    padding: 0;
    background-color: transparent;
    border: 0
}
.stale-files-tab {
    display: none;
    margin-bottom: 10px
}
.files-bucket {
    margin-bottom: 15px
}
.files-bucket.is-stale .stale-files-tab {
    display: block
}
.pull-request-link {
    float: left;
    margin-right: 5px;
    font-size: 13px;
    font-weight: bold;
    padding: 2px 8px;
    line-height: 20px;
    border: 1px solid rgba(65, 131, 196, 0.5);
    border-radius: 3px
}
.pull-request-link:hover {
    background: #4078c0;
    border-color: #4078c0;
    color: #fff;
    text-decoration: none
}
.tabnav-callout {
    position: absolute;
    top: 4px;
    left: 100%;
    display: inline-block;
    padding: 6px 10px;
    margin-left: 3px;
    font-weight: bold;
    line-height: 1;
    white-space: nowrap;
    vertical-align: middle;
    border-radius: 3px;
    box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.1);
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
.tabnav-callout:before {
    position: absolute;
    top: 50%;
    right: 100%;
    margin-top: -7px;
    display: inline-block;
    content: "";
    border: 7px solid transparent
}
.tabnav-callout .octicon {
    vertical-align: text-top
}
.callout-review {
    color: #696143;
    background-color: #fceb9b
}
.callout-review:before {
    border-right-color: #fceb9b
}
.callout-success {
    color: #376a20;
    background-color: #d8f0cd
}
.callout-success:before {
    border-right-color: #d8f0cd
}
.header-with-actions {
    position: relative
}
.header-with-actions h3 {
    margin-top: 5px
}
.header-with-actions .select-menu {
    float: right;
    margin-top: -5px
}
.header-with-actions .select-menu-modal-holder {
    right: 0
}
.header-with-actions .select-menu-modal {
    width: 140px
}
.pulse-blankslate {
    margin-top: 20px
}
.diffstat-summary {
    font-size: 16px;
    vertical-align: middle;
    border-radius: 3px;
    color: #767676;
    line-height: 1.8;
    text-align: left;
    padding: 0 20px 0 0
}
.diffstat-summary a {
    color: #555
}
.diffstat-summary strong {
    color: #333
}
.pulse-graph {
    border-bottom: 1px solid #eee;
    float: left;
    width: 50%;
    padding: 15px 15px 0
}
.pulse-graph:first-child {
    border-right: 1px solid #eee
}
.authors-and-code .insertions {
    color: #6cc644
}
.authors-and-code .deletions {
    color: #bd2c00
}
.authors-and-code .section {
    height: 150px;
    display: table-cell;
    width: 459px
}
.pulse-authors-graph {
    position: relative;
    height: 150px
}
.pulse-authors-graph>svg {
    width: 100%
}
.pulse-authors-graph .dots {
    position: absolute;
    top: 40px;
    left: 0;
    right: 0;
    margin: 0 auto;
    width: 64px;
    height: 64px
}
.pulse-authors-graph .bar rect {
    fill: #c9510c;
    fill-opacity: 0.7
}
.pulse-authors-graph .bar rect:hover {
    fill-opacity: 1
}
.summary-stats {
    display: table;
    width: 100%;
    table-layout: fixed
}
.summary-stats li {
    display: table-cell;
    margin: 0;
    text-align: center;
    color: #767676;
    border-left: 1px solid #eee
}
.summary-stats li a {
    display: block;
    text-decoration: none;
    color: #767676;
    padding-bottom: 10px
}
.summary-stats li a:hover {
    background: #fafafa
}
.summary-stats li .octicon-git-pull-request {
    color: #6e5494
}
.summary-stats li .octicon-git-branch {
    color: #6cc644
}
.summary-stats li .octicon-issue-closed {
    color: #bd2c00
}
.summary-stats li .octicon-issue-opened {
    color: #6cc644
}
.summary-stats li:first-child {
    border-left: 0;
    border-bottom-left-radius: 3px
}
.summary-stats li .num {
    display: block;
    padding-top: 10px;
    font-size: 16px;
    font-weight: bold;
    color: #000
}
.pulse-sections {
    clear: both;
    margin-top: 20px
}
.pulse-section {
    clear: both;
    padding: 0;
    font-size: 14px;
    color: #666
}
.pulse-section p {
    margin-top: 20px
}
.radio-group:before {
    display: table;
    content: ""
}
.radio-group:after {
    display: table;
    clear: both;
    content: ""
}
.radio-label {
    padding: 8px 10px 8px 35px;
    border: 1px solid #d9d9d9;
    margin-left: -1px;
    color: #333;
    float: left;
    cursor: pointer
}
:checked+.radio-label {
    z-index: 1;
    position: relative;
    border-color: #4078c0
}
.radio-label .octicon {
    padding-right: 5px
}
.radio-label:first-of-type {
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px;
    margin-left: 0
}
.radio-label:last-of-type {
    border-top-right-radius: 3px;
    border-bottom-right-radius: 3px;
    padding-right: 16px
}
.radio-input {
    float: left;
    margin: 11px -35px 0 12px;
    z-index: 3
}
#readme.contributing>div {
    max-height: 250px;
    overflow: auto
}
#readme .markdown-body,
#readme .plain {
    background-color: #fff;
    border: 1px solid #ddd;
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px;
    padding: 30px;
    word-wrap: break-word
}
#readme .plain pre {
    font-size: 15px;
    white-space: pre-wrap
}
.file #readme .markdown-body {
    border: 0;
    padding: 30px;
    border-radius: 0
}
.file #readme table[data-table-type="yaml-metadata"] {
    line-height: 1;
    font-size: 12px
}
.file #readme table[data-table-type="yaml-metadata"] table {
    margin: 0
}
.user-recommendations-header {
    width: 550px;
    margin-top: 40px
}
.recommendations-intro-wrapper {
    height: 96px
}
.recommendations-outro {
    padding: 0 100px;
    color: #666;
    border: solid 1px #eee;
    border-radius: 3px
}
.recommendations-complete {
    display: none;
    height: 104px;
    padding: 14px 20px;
    margin: 30px 0 45px;
    font-size: 16px;
    border: solid 1px #eee;
    border-radius: 3px
}
.user-recommendations-form {
    margin: 30px 0
}
.user-interests-label {
    display: block;
    margin-bottom: 10px;
    font-size: 17px;
    font-weight: bold
}
.user-interests-input[type="text"] {
    width: 100%;
    min-height: 40px;
    font-size: 16px
}
.user-interests-examples-wrapper {
    height: 21px;
    margin-top: 13px
}
.user-interests-examples {
    margin: 0;
    color: #767676
}
.user-interests-examples a {
    color: #444
}
.user-interests-list-wrapper {
    height: 41px
}
.skip-button-wrapper {
    display: block;
    text-align: center
}
.button-skip {
    height: 30px;
    padding: 0 12px;
    margin: 0 auto;
    font-size: 13px;
    font-weight: normal;
    line-height: 30px;
    color: #666;
    background-color: transparent;
    background-image: none;
    border: 1px solid rgba(0, 0, 0, 0.1);
    border-radius: 3px
}
.button-skip:hover {
    background-color: #f5f5f5;
    background-image: none;
    border-color: rgba(0, 0, 0, 0.15)
}
.user-interests-list {
    width: 100%;
    text-align: left;
    list-style: none
}
.user-interests-item {
    position: relative;
    float: left;
    height: 41px;
    min-height: 41px;
    padding: 5px 8px 5px 10px;
    margin: 8px;
    font-size: 16px;
    line-height: 30px;
    background-color: #f5f5f5;
    border-left: solid 10px #4078c0;
    border-radius: 3px
}
.user-interests-item.hidden {
    visibility: hidden
}
.user-interests-item.loading {
    border-left: solid 10px #ddd
}
.user-interests-item .spinner {
    display: inline-block;
    width: 16px;
    height: 16px;
    position: relative;
    top: 1px
}
.user-interests-item .octicon-x {
    color: #aaa
}
.user-interests-item .octicon-x:hover {
    color: #4078c0;
    text-decoration: none
}
.remove-user-interest-form {
    display: inline-block
}
.user-interests-item-remove {
    border: 0;
    background-color: transparent;
    outline: none
}
.recommendations-wrapper {
    display: table;
    width: 100%;
    margin-top: 30px;
    border-collapse: collapse;
    border-top: solid 1px #eee
}
.recommendations-wrapper h2 {
    margin-top: 30px;
    font-size: 18px
}
.recommendations-wrapper.disabled {
    color: #ccc
}
.recommendations-wrapper.no_users .recommendations-people {
    display: none
}
.recommendations-wrapper.only_repos_users .recommendations-guides,
.recommendations-wrapper.only_repos_users .recommendations-showcases {
    display: none
}
.recommendations-wrapper.only_repos .recommendations-people,
.recommendations-wrapper.only_repos .recommendations-guides,
.recommendations-wrapper.only_repos .recommendations-showcases {
    display: none
}
.recommendations-left {
    display: table-cell;
    width: 50%;
    padding-top: 10px;
    padding-right: 30px
}
.recommendations-right {
    display: table-cell;
    width: 50%;
    padding-top: 10px;
    padding-left: 30px
}
.recommended-repos {
    min-height: 500px;
    margin-top: 20px;
    list-style: none
}
.recommended-repo-item {
    position: relative;
    padding-right: 80px;
    padding-left: 50px;
    margin-bottom: 30px
}
.recommended-repo-item .starring-container {
    position: absolute;
    top: 0;
    right: 0;
    float: right
}
.recommended-repo-item .author-avatar {
    float: left;
    margin-top: 5px;
    margin-left: -50px;
    border-radius: 3px
}
.recommended-repo-item .repo-meta {
    margin-top: 5px;
    font-size: 13px;
    color: #767676
}
.recommended-repo-item .meta-info {
    margin-right: 10px
}
.recommended-repo-item .repo-title {
    margin-bottom: 4px;
    font-size: 23px
}
.recommended-repo-item .repo-title .repo-author {
    font-weight: normal
}
.recommended-repo-item .repo-title .separator {
    margin: 0 3px;
    font-weight: normal;
    color: #666
}
.recommended-repo-item .repo-description {
    margin-bottom: 4px;
    font-size: 15px;
    line-height: 1.4
}
.recommended-repo-item.placeholder .author-avatar {
    width: 40px;
    height: 40px;
    border: dashed 2px #ccc
}
.recommended-repo-item.placeholder .repo-info {
    position: relative;
    top: 4px;
    display: block;
    height: 100px;
    border: dashed 2px #ccc;
    border-radius: 3px
}
.language-circle {
    position: relative;
    top: -2px;
    display: inline-block;
    width: 12px;
    height: 12px;
    vertical-align: middle;
    border-radius: 50%
}
.recommended-guides {
    margin-top: 22px;
    list-style: none
}
.recommended-guide-item {
    margin-bottom: 10px;
    font-size: 23px
}
.recommended-guide-item .mega-octicon {
    position: relative;
    top: -2px;
    color: #aaa;
    vertical-align: middle
}
.recommended-guide-item.placeholder {
    height: 48px;
    border: dashed 2px #ccc;
    border-radius: 3px
}
.recommended-people-wrapper {
    display: table;
    width: 100%;
    border-collapse: collapse
}
.recommended-people-left,
.recommended-people-right {
    display: table-cell;
    width: 48%
}
.recommended-people-left {
    padding-right: 2%
}
.recommended-people-right {
    padding-left: 2%
}
.recommended-people {
    list-style: none
}
.recommended-person-item {
    padding-left: 50px;
    margin-bottom: 20px;
    font-size: 18px
}
.recommended-person-item .user-following-container {
    margin-top: 5px
}
.recommended-person-item .avatar {
    position: relative;
    top: 5px;
    float: left;
    margin-left: -50px
}
.recommended-person-item .person-meta {
    margin-top: 4px;
    font-size: 13px;
    color: #767676
}
.recommended-person-item .meta-info {
    margin-right: 10px
}
.recommended-person-item.placeholder .avatar {
    width: 40px;
    height: 40px;
    border: dashed 2px #ccc
}
.recommended-person-item.placeholder .person-placeholder {
    position: relative;
    top: 5px;
    display: inline-block;
    width: 140px;
    height: 40px;
    border: dashed 2px #ccc;
    border-radius: 3px
}
.recommended-showcase-link {
    display: table;
    float: left;
    width: 48%;
    margin-bottom: 4%;
    color: #fff
}
.recommended-showcase-link:nth-child(2n+1) {
    margin-left: 4%
}
.recommended-showcase {
    display: table-cell;
    height: 100px;
    font-size: 16px;
    text-align: center;
    vertical-align: middle;
    border-radius: 3px
}
.releases-tag-list {
    width: 100%;
    margin-bottom: 20px;
    border-top: 1px solid #eee
}
.releases-tag-list tr {
    border-bottom: 1px solid #eee
}
.releases-tag-list td {
    padding: 12px 0;
    vertical-align: top
}
.releases-tag-list td.date {
    padding-right: 10px;
    white-space: nowrap
}
.releases-tag-list td.date a {
    color: #767676
}
.releases-tag-list td.main {
    padding-right: 10px
}
.releases-tag-list td.ancillary {
    white-space: nowrap;
    text-align: right
}
.releases-tag-list h4 {
    margin: 0;
    font-size: 14px
}
.releases-tag-list p {
    margin: 0;
    color: #767676;
    font-size: 13px
}
.releases-tag-list p a {
    color: #666;
    font-weight: bold
}
.tag-info h3 {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 14px;
    line-height: 20px
}
.tag-info h3 a {
    color: #666
}
.tag-info h3 a .tag-name {
    color: #000
}
.tag-references {
    margin: 0;
    list-style-type: none;
    font-size: 13px
}
.tag-references>li {
    margin-right: 10px;
    display: inline-block
}
.tag-references>li.commit {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px;
    line-height: 20px
}
.tag-references>li a {
    color: #767676;
    text-decoration: none
}
.tag-references>li a:hover {
    color: #4078c0
}
.release-downloads-header {
    margin-top: 30px
}
.release-downloads {
    margin-top: 10px;
    font-size: 14px;
    border-top: 1px solid #eee
}
.release-downloads li {
    display: block;
    padding-top: 8px;
    padding-bottom: 8px;
    border-bottom: 1px solid #eee
}
.release-downloads .octicon {
    margin-top: 2px;
    margin-right: 5px
}
.release-timeline {
    position: relative;
    border-top: 1px solid #eee
}
.release-timeline-tags {
    list-style-type: none
}
.release-timeline-tags>li {
    display: block
}
.release-timeline-tags>li:before {
    display: table;
    content: ""
}
.release-timeline-tags>li:after {
    display: table;
    clear: both;
    content: ""
}
.release-timeline-tags .date,
.release-timeline-tags .main {
    position: relative;
    float: left;
    padding: 20px
}
.release-timeline-tags .main {
    width: 80%;
    border-left: 2px solid #eee
}
.release-timeline-tags .date {
    width: 20%;
    line-height: 40px;
    text-align: right;
    color: #767676;
    padding-left: 0
}
.release-timeline-tags .date:after {
    box-sizing: border-box;
    content: " ";
    display: block;
    position: absolute;
    top: 50%;
    right: -7px;
    z-index: 10;
    width: 12px;
    height: 12px;
    margin-top: -6px;
    background-color: #eee;
    border: 2px solid #fff;
    border-radius: 6px
}
.release-timeline-tags .octicon-tag {
    padding-left: 5px;
    color: #ccc
}
.release-timeline-tags .expander {
    position: relative;
    display: none
}
.release-timeline-tags .expander .date {
    padding-right: 35px;
    line-height: 20px
}
.release-timeline-tags .expander .date:after {
    display: none
}
.release-timeline-tags .expander .main {
    padding-left: 35px;
    line-height: 20px
}
.release-timeline-tags.is-collapsed .expander {
    display: block
}
.release-timeline-tags.is-collapsed>.collapsable {
    display: none
}
.release-timeline-tags .expander-dots {
    position: absolute;
    top: 18px;
    left: -22px;
    width: 44px;
    text-align: center;
    background-color: #eee;
    border: 2px solid #fff;
    border-radius: 4px;
    z-index: 10;
    cursor: pointer
}
.release-timeline-tags .expander-dots .expander-dot {
    display: inline-block;
    margin-top: -2px;
    width: 4px;
    height: 4px;
    vertical-align: middle;
    border-radius: 2px;
    background-color: #767676
}
.release-timeline-tags .expander-text {
    font-weight: bold;
    color: #666;
    cursor: pointer
}
.release-timeline-tags .expander-text:hover {
    color: #4078c0
}
.release-timeline-tags .expander-text:hover .expander-dots {
    background-color: #4078c0
}
.release-timeline-tags .expander-text:hover .expander-dots .expander-dot {
    background-color: #fff
}
.release:before {
    display: table;
    content: ""
}
.release:after {
    display: table;
    clear: both;
    content: ""
}
.release .tag-references {
    margin-top: 8px
}
.release .tag-references>li {
    display: block;
    margin: 0 0 5px
}
.release-meta {
    float: left;
    width: 20%;
    padding: 40px 20px;
    text-align: right;
    vertical-align: top
}
.release-body {
    float: left;
    width: 80%;
    padding: 40px 20px;
    border-left: 2px solid #eee
}
.release-body .commit-desc pre {
    white-space: pre-line
}
.release-title {
    margin: 0 60px 0 0
}
.release-edit {
    float: right
}
.release-authorship {
    margin-top: 5px;
    margin-bottom: 20px;
    font-size: 14px;
    color: #767676
}
.release-authorship a {
    font-weight: bold;
    color: #666
}
.release-label {
    display: inline-block;
    margin-top: 1px;
    margin-bottom: 10px;
    padding: 5px 10px;
    font-size: 14px;
    font-weight: bold;
    color: #fff;
    background-color: #000;
    border-radius: 3px
}
.release-label.latest {
    background-color: #6cc644
}
.release-label.draft {
    background-color: #bd2c00
}
.release-label.prerelease {
    background-color: #c9510c
}
.release-label a {
    color: #fff
}
.new-release .sidebar h3 {
    margin: 40px 0 -10px;
    font-size: 14px
}
.new-release .sidebar h3:first-child {
    margin-top: 15px
}
.new-release .default,
.new-release .saved,
.new-release .saving,
.new-release .error {
    display: none
}
.new-release .error {
    color: #bd2c00
}
.new-release .is-default .default,
.new-release .is-saving .saving,
.new-release .is-saved .saved,
.new-release .is-failed .error {
    display: inline-block
}
.new-release .saving img {
    vertical-align: top
}
.drop-target .mega-octicon {
    vertical-align: middle;
    color: #e5e5e5
}
.drop-target p {
    padding: 16px 0;
    height: 65px;
    font-size: 14px;
    text-align: center;
    border-color: #ddd;
    border-style: dashed
}
.drop-target .octospinner {
    vertical-align: middle
}
.uploaded-files {
    background: #fff;
    border-top-right-radius: 3px;
    border-top-left-radius: 3px
}
.uploaded-files.not-populated+.drop-target p {
    border-top-right-radius: 3px;
    border-top-left-radius: 3px;
    border-top: dashed 1px #ccc
}
.uploaded-files.is-populated {
    border: 1px solid #ddd;
    border-bottom-color: #e5e5e5
}
.uploaded-files.is-populated+.drop-target p {
    border-top-right-radius: 0;
    border-top-left-radius: 0;
    border-top: 0
}
.uploaded-files>li {
    list-style-type: none;
    margin: 0;
    padding: 8px 10px;
    border-top: 1px solid #eee;
    line-height: 22px
}
.uploaded-files>li.template {
    display: none
}
.uploaded-files>li .delete-pending {
    display: none
}
.uploaded-files>li.delete {
    background: #f9f9f9;
    color: #767676
}
.uploaded-files>li.delete:nth-child(2) {
    border-top-right-radius: 3px;
    border-top-left-radius: 3px
}
.uploaded-files>li.delete .delete-pending {
    display: block
}
.uploaded-files>li.delete .live {
    display: none
}
.uploaded-files>li.delete .filename {
    color: #bd2c00
}
.uploaded-files>li:nth-child(2) {
    border-top: 0
}
.uploaded-files .filename {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 11px
}
.uploaded-files .filesize {
    font-size: 12px;
    color: #767676
}
.uploaded-files input[type=text] {
    width: 490px;
    margin-right: 6px;
    padding: 2px 4px;
    border-radius: 2px
}
.uploaded-files .remove {
    float: right;
    margin-top: 2px;
    color: #767676
}
.uploaded-files .remove:hover {
    color: #bd2c00
}
.uploaded-files .undo {
    float: right
}
.upload-progress {
    background: #fff;
    border: 0;
    border-radius: 30px;
    box-shadow: 0 1px 1px #fff, inset 0 1px 1px rgba(255, 255, 255, 0.5);
    height: 3px;
    margin-top: 3px;
    position: relative
}
.upload-progress .upload-meter {
    background-image: -webkit-linear-gradient(#8dd2f7, #58b8f4);
    background-image: linear-gradient(#8dd2f7, #58b8f4);
    border-radius: 30px;
    height: 100%;
    position: absolute;
    top: 0
}
.release-body-form .previewable-comment-form .comment-form-head.tabnav {
    padding: 0;
    background-color: transparent
}
.release-body-form .previewable-comment-form .write-content,
.release-body-form .previewable-comment-form .preview-content {
    padding: 0 0 10px;
    margin: 0
}
.release-tag-form .for-loading,
.release-tag-form .for-empty,
.release-tag-form .for-valid,
.release-tag-form .for-invalid,
.release-tag-form .for-duplicate,
.release-tag-form .for-pending {
    display: none
}
.release-tag-form.is-loading .for-loading {
    display: block
}
.release-tag-form.is-empty .for-empty {
    display: block
}
.release-tag-form.is-valid .for-valid {
    display: block
}
.release-tag-form.is-invalid .for-invalid {
    display: block
}
.release-tag-form.is-duplicate .for-duplicate {
    display: block
}
.release-tag-form.is-pending .for-pending {
    display: block
}
.release-target-wrapper {
    display: inline-block
}
.release-target-wrapper.hidden {
    display: none
}
.releases-target-menu {
    display: inline-block;
    margin-left: 5px
}
.releases-target-menu .btn-sm {
    line-height: 32px
}
.releases-target-menu .select-menu-button:before {
    top: 14px
}
.release-show {
    border-top: 1px solid #eee
}
.release-show .release-edit {
    display: none
}
.render-container {
    background: #ddd;
    text-align: center;
    padding: 30px;
    line-height: 0
}
.render-container .render-viewer {
    border: 0;
    display: none;
    width: 100%;
    height: 100%
}
.render-container .octospinner {
    display: none
}
.render-container .render-viewer-error,
.render-container .render-viewer-fatal,
.render-container .render-viewer-invalid {
    display: none
}
.render-container.is-render-automatic .octospinner {
    display: inline-block
}
.render-container.is-render-requested .octospinner {
    display: inline-block
}
.render-container.is-render-requested.is-render-failed .render-viewer-error {
    display: inline-block
}
.render-container.is-render-requested.is-render-failed .render-viewer,
.render-container.is-render-requested.is-render-failed .render-viewer-fatal,
.render-container.is-render-requested.is-render-failed .render-viewer-invalid,
.render-container.is-render-requested.is-render-failed .octospinner {
    display: none
}
.render-container.is-render-requested.is-render-failed-fatal .render-viewer-fatal {
    display: inline-block
}
.render-container.is-render-requested.is-render-failed-fatal .render-viewer,
.render-container.is-render-requested.is-render-failed-fatal .render-viewer-error,
.render-container.is-render-requested.is-render-failed-fatal .render-viewer-invalid .octospinner {
    display: none
}
.render-container.is-render-requested.is-render-failed-invalid .render-viewer-invalid {
    display: inline-block
}
.render-container.is-render-requested.is-render-failed-invalid .render-viewer,
.render-container.is-render-requested.is-render-failed-invalid .render-viewer-error,
.render-container.is-render-requested.is-render-failed-invalid .render-viewer-fatal,
.render-container.is-render-requested.is-render-failed-invalid .octospinner {
    display: none
}
.render-container.is-render-ready.is-render-requested:not(.is-render-failed) {
    background: none;
    height: 500px;
    padding: 0
}
.render-container.is-render-ready.is-render-requested:not(.is-render-failed) .render-viewer {
    display: block
}
.render-container.is-render-ready.is-render-requested:not(.is-render-failed) .render-viewer-error,
.render-container.is-render-ready.is-render-requested:not(.is-render-failed) .render-viewer-fatal,
.render-container.is-render-ready.is-render-requested:not(.is-render-failed) .octospinner {
    display: none
}
.render-notice {
    padding: 20px 15px;
    font-size: 14px;
    color: #4c4a42;
    background-color: #fff9ea;
    border-color: #dfd8c2
}
.repohead.experiment-repo-nav {
    padding-bottom: 0;
    background-color: #fafafa
}
.repohead.mirror h1,
.repohead.fork h1 {
    margin-top: -5px;
    margin-bottom: 15px;
    height: auto
}
.repohead h1 {
    position: relative;
    float: left;
    padding-left: 34px;
    color: #666
}
.repohead h1.private .mega-octicon,
.repohead h1.private .octicon {
    color: #e9dba5
}
.repohead h1 .octicon-lock,
.repohead h1 .octicon-repo,
.repohead h1 .octicon-mirror,
.repohead h1 .octicon-repo-forked,
.repohead h1 .octicon-gist,
.repohead h1 .octicon-gist-secret {
    position: absolute;
    left: 0;
    top: 12px;
    line-height: 32px;
    margin-top: -13px;
    color: #bbb
}
.repohead .octicon-mirror {
    left: -3px
}
.repohead .octicon-lock {
    top: 10px
}
.repohead span.fork-flag,
.repohead span.mirror-flag {
    display: block;
    font-size: 11px;
    line-height: 10px;
    white-space: nowrap
}
.reponav {
    position: relative;
    top: 1px;
    margin-top: 17px
}
.reponav:before {
    display: table;
    content: ""
}
.reponav:after {
    display: table;
    clear: both;
    content: ""
}
.reponav-dropdown {
    position: relative;
    float: left
}
.reponav-dropdown.active .dropdown-menu-content {
    display: block
}
.reponav-dropdown.active .dropdown-menu {
    left: 15px
}
.reponav-item {
    float: left;
    padding: 7px 15px 8px;
    color: #666;
    white-space: nowrap;
    border: solid transparent;
    border-width: 3px 1px 1px;
    border-radius: 3px 3px 0 0
}
.reponav-item .octicon {
    color: #bbb
}
.reponav-item .counter {
    color: #555
}
.reponav-item:hover {
    color: #333;
    text-decoration: none
}
.reponav-item.selected {
    font-weight: bold;
    color: #333;
    background-color: #fff;
    border-color: #d26911 #e5e5e5 transparent
}
.reponav-item.selected>.octicon {
    color: inherit
}
.mini-repo-list {
    list-style: none
}
.mini-repo-list>li:first-child .mini-repo-list-item {
    border-top: 0
}
.mini-repo-list>li:last-child .mini-repo-list-item {
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.mini-repo-list .no-repo {
    padding: 15px;
    color: #767676;
    text-align: center
}
.mini-repo-list .repo-name {
    font-weight: bold
}
.mini-repo-list-item {
    position: relative;
    display: block;
    padding: 6px 64px 6px 30px;
    font-size: 14px;
    border-top: 1px solid #e5e5e5
}
.mini-repo-list-item:hover {
    text-decoration: none
}
.mini-repo-list-item:hover .repo,
.mini-repo-list-item:hover .owner {
    text-decoration: underline
}
.mini-repo-list-item .repo-icon {
    float: left;
    margin-top: 2px;
    margin-left: -20px;
    color: #666
}
.mini-repo-list-item .repo-and-owner {
    max-width: 220px
}
.mini-repo-list-item .owner {
    max-width: 110px
}
.mini-repo-list-item .repo {
    font-weight: bold
}
.mini-repo-list-item .stars {
    position: absolute;
    top: 0;
    right: 10px;
    margin-top: 6px;
    font-size: 12px;
    color: #888
}
.mini-repo-list-item .repo-description {
    display: block;
    max-width: 100%;
    font-size: 12px;
    color: #767676;
    line-height: 21px
}
.popular-repos .mini-repo-list-item .stars {
    margin-top: 16px
}
.popular-repos .no-description .mini-repo-list-item {
    padding-top: 17px;
    padding-bottom: 16px
}
.private .mini-repo-list-item {
    background-color: #fff9ea
}
.private .mini-repo-list-item .repo-icon {
    color: #a1882b
}
.filter-bar {
    padding: 10px;
    background-color: #fafafa;
    border-bottom: 1px solid #e5e5e5
}
.filter-bar:before {
    display: table;
    content: ""
}
.filter-bar:after {
    display: table;
    clear: both;
    content: ""
}
.filter-bar .filter-input {
    width: 100%;
    min-height: 26px;
    padding: 3px 10px;
    font-size: 11px;
    border-radius: 12px
}
.user-repos .filter-bar {
    text-align: center
}
.filter-repos {
    padding-bottom: 0
}
.repo-filterer {
    display: inline-block;
    margin-top: 6px;
    list-style: none
}
.repo-filterer li {
    display: inline-block
}
.repo-filterer .repo-filter {
    display: inline-block;
    padding: 5px 5px 6px;
    margin-right: 5px;
    font-size: 11px;
    color: #767676;
    border-bottom: 2px solid transparent
}
.repo-filterer .repo-filter:hover {
    text-decoration: none;
    border-bottom-color: #e5e5e5
}
.repo-filterer .repo-filter.filter-selected {
    color: #333;
    text-decoration: none;
    border-bottom-color: #d26911;
    outline: none
}
.more-repos {
    text-align: center;
    box-shadow: inset 0 1px 0 #e5e5e5
}
.more-repos img {
    margin: 11px auto
}
.more-repos-link {
    display: block;
    padding: 10px;
    color: #7aa1d3
}
.more-repos-link:hover {
    color: #4078c0;
    text-decoration: none
}
.more-repos-link.is-loading {
    text-indent: -9999px;
    cursor: default;
    background-image: url(https://github.com/images/spinners/octocat-spinner-16px.gif);
    background-repeat: no-repeat;
    background-position: center center
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    .more-repos-link.is-loading {
        background-image: url(https://github.com/images/spinners/octocat-spinner-32.gif);
        background-size: 16px 16px
    }
}
.empty-repo {
    font-size: 14px
}
.empty-repo .url-box {
    display: block;
    width: 100%;
    height: auto;
    padding: 0;
    margin: 0;
    border: 0
}
.empty-repo .clone-urls {
    width: 100%
}
.empty-repo .or-text {
    margin-right: 5px;
    margin-left: 5px
}
.empty-repo-setup-option .copyable-terminal-content {
    font-size: 14px
}
.empty-repo-setup-option h3 {
    margin-top: 0
}
.empty-repo-setup-option p:last-child {
    margin-bottom: 0
}
.give-access-setup-option {
    margin-bottom: 20px
}
.timeout {
    width: auto;
    height: 300px;
    padding: 0;
    margin: 20px 0;
    background-color: transparent;
    border: 0
}
.timeout h3 {
    padding-top: 100px;
    color: #767676
}
.repo-container {
    min-height: 345px
}
.repo-nav .counter {
    display: none
}
.repo-nav .full-word {
    display: none
}
.with-full-navigation .repo-nav .counter {
    display: block
}
.with-full-navigation .repo-nav .full-word {
    display: inline-block;
    width: 128px;
    vertical-align: top
}
.only-with-full-nav {
    display: none
}
.with-full-navigation .only-with-full-nav {
    display: block
}
.only-with-full-nav .btn-block {
    margin-bottom: 5px
}
.repository-with-sidebar:before {
    display: table;
    content: ""
}
.repository-with-sidebar:after {
    display: table;
    clear: both;
    content: ""
}
.repository-with-sidebar .repository-sidebar {
    float: right;
    width: 38px
}
.repository-with-sidebar .repository-sidebar .sidebar-button {
    width: 100%;
    margin: 0 0 10px;
    text-align: center
}
.repository-with-sidebar .repository-sidebar h3 {
    margin-bottom: 5px;
    font-size: 11px;
    font-weight: normal;
    color: #767676
}
.repository-with-sidebar .repository-sidebar .clone-url {
    display: none;
    margin-top: -5px
}
.repository-with-sidebar .repository-sidebar .clone-url.open {
    display: block
}
.repository-with-sidebar .repository-sidebar .clone-options {
    margin: 8px 0 15px;
    font-size: 11px;
    color: #666
}
.repository-with-sidebar .repository-sidebar .clone-options .octicon-question {
    position: relative;
    bottom: 1px;
    font-size: 11px;
    color: #000;
    cursor: pointer
}
.repository-with-sidebar .repository-content {
    float: left;
    width: 920px
}
.repository-with-sidebar.with-full-navigation .repository-content {
    width: 790px
}
.repository-with-sidebar.with-full-navigation .repository-sidebar {
    width: 170px
}
.repository-with-sidebar.with-full-navigation .sunken-menu-group .tooltipped:before,
.repository-with-sidebar.with-full-navigation .sunken-menu-group .tooltipped:after {
    display: none
}
.overall-summary {
    position: relative;
    margin-bottom: 10px;
    border: 1px solid #ddd;
    border-radius: 3px
}
.overall-summary-bottomless {
    margin-bottom: 0;
    border-bottom: 0;
    border-radius: 3px 3px 0 0
}
.numbers-summary {
    display: table;
    width: 100%;
    table-layout: fixed
}
.numbers-summary li {
    display: table-cell;
    padding: 0;
    margin: 0;
    text-align: center;
    white-space: nowrap
}
.numbers-summary a,
.numbers-summary .nolink {
    display: block;
    padding: 10px 0;
    color: #767676;
    text-decoration: none
}
.numbers-summary .octicon {
    opacity: 0.5
}
.numbers-summary a:hover {
    color: #4078c0
}
.numbers-summary a:hover .num {
    color: inherit
}
.repo-private-label,
.gist-secret-label {
    display: inline-block;
    padding: 4px 5px 3px;
    font-size: 11px;
    font-weight: 300;
    line-height: 11px;
    color: #a1882b;
    text-transform: uppercase;
    vertical-align: middle;
    background-color: #ffefc6;
    border-radius: 3px
}
.repository-meta {
    margin: 0 0 13px;
    font-size: 16px;
    color: #666
}
.repository-meta:before {
    display: table;
    content: ""
}
.repository-meta:after {
    display: table;
    clear: both;
    content: ""
}
.repository-meta .repository-description {
    display: inline;
    word-wrap: break-word
}
.repository-meta .repository-website {
    display: inline-block
}
.repository-meta .repo-description-field {
    width: 380px
}
.repository-meta .repo-website-field {
    width: 270px
}
.repository-meta .edit-repository-meta {
    display: none;
    font-size: 13px;
    margin-bottom: 5px
}
.repository-meta .edit-repository-meta .field {
    display: inline-block;
    margin-right: 5px
}
.repository-meta .edit-repository-meta label {
    display: block;
    margin-bottom: 6px;
    font-weight: bold;
    color: #333
}
.repository-meta.open .repository-meta-content,
.repository-meta.open .repository-description,
.repository-meta.open .repository-website,
.repository-meta.open .edit-link {
    display: none
}
.repository-meta.open .edit-repository-meta {
    display: block
}
.experiment-repo-nav .repo-description-field {
    width: 540px
}
.experiment-repo-nav .repo-website-field {
    width: 300px
}
.experiment-repo-nav .new-issue-form {
    border-top: 0;
    padding-top: 0
}
.experiment-repo-nav.repohead h1 {
    padding-left: 18px;
    font-size: 18px;
    line-height: 26px
}
.experiment-repo-nav.repohead h1 .octicon {
    line-height: 26px;
    margin-top: 0;
    top: 0
}
.experiment-repo-nav.repohead .repo-mirror {
    padding-left: 20px
}
.file-navigation:before {
    display: table;
    content: ""
}
.file-navigation:after {
    display: table;
    clear: both;
    content: ""
}
.file-navigation .select-menu,
.file-navigation .btn-group,
.file-navigation .breadcrumb {
    margin-bottom: 10px
}
.file-navigation.in-mid-page {
    margin-top: 10px
}
.file-navigation .select-menu {
    margin-right: 10px
}
.file-navigation .breadcrumb {
    float: left;
    margin-top: 0;
    margin-bottom: 5px
}
.file-navigation .breadcrumb .octicon-btn.disabled {
    color: #bbb;
    cursor: default
}
.file-navigation .breadcrumb .octicon-btn.disabled:hover {
    color: #bbb
}
.file-navigation .btn-group {
    margin-left: 10px
}
.file-navigation .compare-button {
    margin-right: 5px
}
.file-navigation-new {
    margin-bottom: 10px
}
.file-navigation-new.in-mid-page {
    margin-top: 15px
}
.file-navigation-new .select-menu,
.file-navigation-new .btn-group,
.file-navigation-new .breadcrumb {
    margin-bottom: 0
}
.file-navigation-new+.breadcrumb {
    margin-bottom: 10px
}
.file-navigation-new .btn-block+.btn-block {
    margin-top: 10px
}
.iconbutton .octicon {
    margin-right: 0
}
.file-wrap {
    margin-bottom: 10px;
    border: 1px solid #ddd;
    border-top: 0;
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px
}
.file-wrap .include-fragment-error {
    display: none
}
.file-wrap.is-error .include-fragment-error {
    display: table-row
}
table.files {
    width: 100%;
    background: #f8f8f8;
    border-radius: 2px
}
table.files td {
    padding: 6px 3px;
    line-height: 20px;
    border-top: 1px solid #eee
}
table.files td.icon {
    width: 17px;
    padding-right: 2px;
    padding-left: 10px;
    color: #767676
}
table.files td.icon .octicon-file-directory {
    color: #80a6cd
}
table.files td.icon .spinner {
    position: relative;
    top: 3px;
    display: none;
    margin-top: -3px;
    margin-left: -2px
}
table.files td .simplified-path {
    color: #888
}
table.files td .css-truncate {
    max-width: 100%
}
table.files td.content {
    max-width: 180px
}
table.files td.message {
    max-width: 442px;
    padding-left: 10px;
    overflow: hidden;
    color: #888
}
table.files td.message .emoji {
    vertical-align: top
}
table.files td.message a {
    color: #888
}
table.files td.message a:hover {
    color: #4078c0
}
table.files td.age {
    max-width: 140px;
    padding-right: 10px;
    color: #888;
    text-align: right;
    white-space: nowrap
}
table.files tr.is-loading td.icon .octicon {
    display: none
}
table.files tr.is-loading td.icon .spinner {
    display: inline-block
}
table.files tr.up-tree {
    border-bottom: 1px solid #eee
}
table.files tr.up-tree a {
    padding: 3px 6px;
    margin-left: -3px;
    font-weight: bold;
    border-radius: 2px
}
table.files tr.up-tree a:hover {
    background-color: #eee
}
table.files tbody tr:first-child td {
    border-top: 0
}
.branch-infobar {
    padding: 8px;
    color: #767676;
    background-color: #fafafa;
    border: solid #ddd;
    border-width: 1px 1px 0;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.branch-infobar:before {
    display: table;
    content: ""
}
.branch-infobar:after {
    display: table;
    clear: both;
    content: ""
}
.branch-infobar .muted-link {
    display: inline-block;
    margin-left: 10px
}
.branch-infobar .muted-link .octicon {
    color: #bbb
}
.branch-infobar .muted-link:hover .octicon {
    color: inherit
}
.branch-infobar+.commit-tease {
    border-radius: 0
}
.fork-select-fragment {
    text-align: center
}
.spinner-forking {
    display: block;
    margin: 20px auto 40px
}
.prereceive-feedback {
    padding: 15px;
    margin-bottom: 15px;
    border: 1px solid #ddd;
    border-left: 3px solid #cea61b;
    border-radius: 3px
}
.prereceive-feedback-heading {
    margin-top: 0;
    margin-bottom: 10px;
    color: #cea61b
}
.file-navigation-options {
    display: inline-block;
    vertical-align: middle
}
.file-navigation-options.active .dropdown-menu-content {
    display: block
}
.file-navigation-options .dropdown-menu {
    padding: 15px;
    width: 300px
}
.file-navigation-options .dropdown-divider {
    margin: 15px -15px
}
.file-navigation-options .clone-url {
    display: none
}
.file-navigation-options .clone-url.open {
    display: block
}
.file-navigation-options .clone-url h3 {
    margin-top: 0;
    margin-bottom: 5px;
    font-size: 14px
}
.file-navigation-options .clone-options {
    margin: 8px 0 15px;
    font-size: 11px;
    color: #666
}
.file-navigation-options .clone-options .octicon-question {
    position: relative;
    bottom: 1px;
    font-size: 11px;
    color: #000;
    cursor: pointer
}
.file-navigation-option {
    position: relative;
    display: inline-block;
    margin-left: 3px
}
.url-box {
    width: 100%;
    margin-top: 10px;
    margin-left: -10px;
    padding: 10px 10px 0;
    border-top: 1px solid #ddd;
    height: 26px
}
.url-box p {
    float: left;
    margin: 0 0 0 5px;
    height: 26px;
    line-height: 26px;
    font-size: 11px;
    color: #666
}
.url-box p strong {
    color: #000
}
.clone-urls {
    display: table;
    float: left;
    width: 585px
}
.clone-url-button {
    display: table-cell;
    width: 1%;
    vertical-align: top
}
.clone-url-button:first-child .clone-url-link {
    border-top-left-radius: 3px;
    border-bottom-left-radius: 3px
}
.clone-url-button>.clone-url-link {
    position: relative;
    display: block;
    padding: 0 9px;
    margin-right: -1px;
    font-size: 11px;
    font-weight: bold;
    line-height: 24px;
    color: #333;
    text-decoration: none;
    text-shadow: 0 1px 0 #fff;
    background-image: -webkit-linear-gradient(#fafafa, #eaeaea);
    background-image: linear-gradient(#fafafa, #eaeaea);
    border: 1px solid #ccc;
    white-space: nowrap;
    cursor: pointer
}
.clone-url-button>.clone-url-link:hover,
.clone-url-button>.clone-url-link:active {
    z-index: 3;
    color: #fff;
    text-decoration: none;
    text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.25);
    background-image: -webkit-linear-gradient(#599bcd, #3072b3);
    background-image: linear-gradient(#599bcd, #3072b3);
    border-color: #2a65a0
}
.clone-url-button>.clone-url-link:active {
    background-color: #3072b3;
    background-image: none;
    border-color: #25588c;
    box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.15)
}
.clone-url-button>.clone-url-link:focus {
    outline: 0
}
.clone-url-button+.clone-url-button>.clone-url-link {
    box-shadow: inset 1px 0 0 #fff
}
.clone-url-button+.clone-url-button>.clone-url-link:hover {
    box-shadow: none
}
.clone-url-button+.clone-url-button>.clone-url-link:active {
    box-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.15)
}
.clone-url-button.selected>.clone-url-link,
.clone-url-button.selected>.clone-url-link:hover {
    z-index: 2;
    color: #333;
    text-shadow: 0 1px 0 rgba(255, 255, 255, 0.6);
    border-color: #bbb;
    background-color: #ccc;
    background-image: -webkit-linear-gradient(#ccc, #d5d5d5);
    background-image: linear-gradient(#ccc, #d5d5d5);
    box-shadow: inset 0 2px 3px rgba(0, 0, 0, 0.075)
}
input.url-field {
    position: relative;
    width: 100%;
    min-height: 26px;
    padding: 0 5px;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 12px;
    border-radius: 0
}
input.url-field:focus {
    z-index: 2
}
.url-box-clippy .zeroclipboard-button {
    border-radius: 0 3px 3px 0;
    border-left: 0;
    margin-left: 0 !important
}
.pagehead.repohead .select-menu .select-menu-modal-holder {
    z-index: 25
}
.audit-log-map-container {
    position: relative;
    margin-bottom: 10px
}
.audit-log-map-container .activity {
    position: absolute;
    text-align: center;
    z-index: 99999;
    top: 120px;
    left: 450px;
    display: none
}
.audit-log-map-container .is-graph-loading .activity {
    display: block
}
.audit-search-form {
    position: relative
}
.audit-search-form:before {
    display: table;
    content: ""
}
.audit-search-form:after {
    display: table;
    clear: both;
    content: ""
}
.audit-search-form .subnav-search-input {
    width: 250px
}
.audit-log-map {
    overflow: hidden;
    height: 325px;
    border-radius: 3px;
    box-shadow: inset 1px 1px 0 rgba(0, 0, 0, 0.2);
    background-color: #4078c0;
    margin-bottom: 25px
}
.map-background {
    fill: #4078c0;
    cursor: -webkit-grab;
    cursor: grab;
    pointer-events: all
}
.land {
    fill: none;
    stroke: #256aae;
    stroke-width: 2;
    shape-rendering: crispedges
}
.country {
    cursor: pointer;
    fill: #d7c7ad;
    shape-rendering: crispedges
}
.country.hk {
    stroke: #a5967e
}
.country:hover {
    fill: #c8b28e
}
.country.active {
    fill: #f6e5ca
}
.borders {
    fill: none;
    stroke: #a5967e;
    shape-rendering: crispedges
}
.graticule {
    pointer-events: none;
    fill: none;
    stroke: #fff;
    stroke-opacity: 0.2;
    shape-rendering: crispedges
}
.graticule:nth-child(2n) {
    stroke-dasharray: 2, 2
}
.security-map-legend circle {
    stroke: #fff;
    stroke-width: 1.5;
    fill-opacity: 0
}
.security-map-legend text {
    fill: #fff;
    font-size: 10px;
    text-anchor: end
}
.security-map-legend .link {
    stroke-width: 1.5;
    stroke: #fff
}
.audit-point {
    pointer-events: none;
    fill: #bd2c00;
    fill-opacity: 0.8;
    stroke: #bd2c00
}
.country-info {
    opacity: 0;
    position: absolute;
    top: 10px;
    right: 10px;
    padding: 10px;
    pointer-events: none;
    background: rgba(255, 255, 255, 0.9);
    border-radius: 2px
}
.audit-log-search .member-info {
    width: 300px
}
.audit-log-search .member-info .member-avatar {
    float: left;
    margin-right: 15px
}
.audit-log-search .member-info .member-link {
    display: block
}
.audit-log-search .member-info .member-list-avatar {
    margin-right: 0
}
.audit-log-search .member-info .ghost {
    color: #767676;
    display: inline-block
}
.audit-log-search .audit-action-info {
    float: left;
    max-width: 220px
}
.audit-log-search .audit-action-info .member-username {
    font-size: 15px
}
.audit-log-search .audit-action-info a {
    color: #4078c0
}
.audit-log-search .blankslate {
    border-top-left-radius: 0;
    border-top-right-radius: 0
}
.audit-results-header {
    padding: 10px;
    border: 1px solid #dcdcdc;
    border-bottom: 0;
    background-color: #f7f7f7;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.audit-results-header:before {
    display: table;
    content: ""
}
.audit-results-header:after {
    display: table;
    clear: both;
    content: ""
}
.audit-results-header-title {
    margin-top: 10px;
    margin-bottom: 0;
    font-size: 14px;
    font-weight: normal;
    color: #333
}
.audit-search-clear {
    padding: 10px;
    border: 1px solid #e5e5e5;
    border-bottom: 0
}
.audit-search-clear .issues-reset-query {
    margin-bottom: 0
}
.audit-action-info {
    margin-top: 3px;
    font-weight: normal;
    font-size: 12px;
    color: #767676
}
.audit-action-info .context {
    color: #333
}
.audit-type {
    width: 200px;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap
}
.audit-type .octicon {
    margin-right: 3px;
    font-size: 14px;
    font-weight: normal
}
.audit-type .repo {
    color: #c9510c
}
.audit-type .team {
    color: #6cc644
}
.audit-type .user {
    color: #6e5494
}
.audit-type .oauth_access {
    color: #bd2c00
}
.audit-type .hook {
    color: #e1bf4e
}
.export-actions {
    display: inline-block;
    margin-left: 10px
}
.export-actions a {
    color: #999;
    margin-top: -3px
}
.export-actions a:hover {
    color: #4078c0;
    text-decoration: none
}
.export-actions .select-menu-button:after {
    position: absolute;
    right: 15px;
    top: 50%;
    margin-top: -2px
}
.export-actions .select-menu-modal {
    width: 111px
}
.export-actions .select-menu-item-text {
    padding: 8px 0;
    text-align: center
}
.export-phrase {
    margin-top: 5px
}
.export-phrase pre {
    border-left: 1px solid #eee;
    color: #666;
    font-size: 11px;
    padding-left: 10px;
    white-space: pre-wrap;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.audit-log-export-button {
    -webkit-transition: 0.25s width ease-in-out;
    transition: 0.25s width ease-in-out;
    width: 110px;
    height: 34px
}
.audit-log-export-button .loader {
    display: none;
    position: absolute;
    left: 11px;
    top: 50%;
    margin-top: -9px
}
.audit-log-export-button .octicon {
    position: absolute;
    left: 11px;
    top: 50%;
    margin-top: -9px
}
.audit-log-export-button .audit-log-export-status {
    position: absolute;
    left: 35px;
    top: 7px
}
.audit-log-export-button.disabled {
    width: 125px
}
.audit-log-export-button.disabled:after {
    display: none
}
.audit-log-export-button.disabled .octicon {
    display: none
}
.audit-log-export-button.disabled .loader {
    display: block
}
.full-export .audit-log-export-button {
    width: 137px
}
.full-export.export-actions .select-menu-modal {
    width: 137px
}
.context-loader-container .large-format-loader {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    padding-top: 190px;
    background: rgba(255, 255, 255, 0.8);
    z-index: 9999;
    text-align: center;
    color: #767676
}
.profile-picture {
    margin: 10px 0 0
}
.profile-picture p {
    float: left;
    margin-top: 7px
}
.profile-picture img {
    float: left;
    margin: 0 10px 0 0;
    border-radius: 3px
}
.app-owner {
    margin: 10px 0 -10px
}
.edit-profile-avatar .drag-and-drop {
    padding: 0;
    color: #666;
    border-width: 0
}
.edit-profile-avatar input {
    cursor: pointer
}
.edit-profile-avatar.is-bad-file {
    border: 0
}
.edit-profile-avatar .manual-file-chooser {
    position: absolute;
    top: 0;
    left: 0;
    width: 146px;
    height: 34px;
    margin-left: 0;
    padding: 0;
    cursor: pointer
}
.button-change-profile-picture {
    overflow: hidden
}
.croppable-avatar {
    display: none
}
.profile-picture-cropper {
    max-width: 400px;
    text-align: center;
    margin: 0 auto 15px
}
.profile-picture-cropper>img {
    max-width: 100%
}
.profile-picture-cropper .jcrop-holder {
    display: inline-block
}
.profile-picture-spinner {
    display: inline-block;
    background-image: url(https://github.com/images/spinners/octocat-spinner-128.gif);
    background-repeat: no-repeat;
    background-position: center;
    background-size: 64px 64px
}
.profile-picture-spinner.hidden {
    display: none
}
.avatar-upload {
    float: left;
    width: 340px;
    margin-left: 20px
}
.avatar-upload .flash {
    width: 100%;
    padding: 30px 15px;
    border: dashed 1px #bd2c00;
    box-shadow: none
}
.avatar-upload .upload-state {
    display: none;
    padding: 10px 0
}
.avatar-upload .upload-state p {
    margin: 0;
    font-size: 12px;
    color: #767676
}
.avatar-upload .avatar-upload .octicon {
    display: inline-block
}
.is-default .avatar-upload .default {
    display: block
}
.is-uploading .avatar-upload .loading {
    display: block;
    padding: 0
}
.is-uploading .avatar-upload .loading img {
    vertical-align: top
}
.is-uploading .avatar-upload .button-change-profile-picture {
    display: none
}
.is-bad-file .avatar-upload .bad-file {
    display: block;
    margin: 0
}
.is-too-big .avatar-upload .too-big {
    display: block;
    margin: 0
}
.is-bad-dimensions .avatar-upload .bad-dimensions {
    display: block;
    margin: 0
}
.is-failed .avatar-upload .failed-request {
    display: block;
    margin: 0
}
.is-empty .avatar-upload .file-empty {
    display: block;
    margin: 0
}
.is-bad-browser .avatar-upload .bad-browser {
    display: block;
    margin: 0
}
dl.new-email-form {
    padding: 10px 10px 0;
    margin: 0 -10px 10px;
    border-top: 1px solid #e5e5e5
}
span.label.default {
    margin-left: 4px;
    padding: 4px 6px;
    background-color: #6cc644;
    color: #fff;
    border-radius: 3px
}
span.label.visibility {
    margin-left: 4px;
    padding: 4px 6px;
    background-color: #999;
    color: #fff;
    border-radius: 3px
}
span.label.bouncing {
    margin-left: 4px;
    padding: 4px 6px;
    background-color: #daa520;
    color: #fff;
    border-radius: 3px
}
.email-actions {
    float: right
}
.email-actions>span {
    float: left
}
.email-actions form {
    display: inline
}
.email-actions span.label {
    font-size: 13px;
    color: #767676;
    padding: 0 10px
}
.email-actions .octicon-alert {
    color: #ca5633
}
.boxed-group .fork-flag {
    margin-left: 16px;
    font-size: 12px;
    color: #767676
}
li.ssh-key {
    position: relative;
    line-height: 18px;
    padding: 15px
}
li.ssh-key .btn {
    float: right;
    margin-top: 4px
}
li.ssh-key .team-label-ldap {
    float: right
}
li.ssh-key .ssh-key-state-indicator {
    float: left;
    width: 8px;
    height: 8px;
    margin-top: 10px;
    border-radius: 5px
}
li.ssh-key .ssh-key-state-indicator.recent {
    background-color: #6cc644;
    box-shadow: 0 0 10px rgba(108, 198, 68, 0.5)
}
li.ssh-key .ssh-key-state-indicator.not-recent {
    box-shadow: 0 1px 0 #fff;
    background-image: -webkit-linear-gradient(#aaa, #ccc);
    background-image: linear-gradient(#aaa, #ccc)
}
li.ssh-key .ssh-key-icon {
    float: left;
    width: 32px;
    margin-top: 1px;
    margin-left: 15px;
    text-align: center
}
.ssh-key-details {
    position: relative;
    margin-left: 70px;
    margin-right: 150px
}
.ssh-key-details .ssh-key-title {
    display: block;
    max-width: 360px
}
.ssh-key-fingerprint {
    display: block;
    font-weight: normal;
    color: #767676
}
#notification-center .overview {
    padding: 0 10px 10px;
    border-bottom: 1px solid #ddd
}
.oauth-stats dl.keys {
    float: right;
    margin: 0;
    text-align: right
}
.oauth-stats dl.keys dt {
    color: #767676;
    font-weight: bold
}
.oauth-stats dl.keys dd {
    color: #333;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace
}
.user-count {
    float: left;
    margin: 12px 0 0;
    font-size: 36px;
    color: #767676;
    font-weight: 300
}
.access-token-group .boxed-group-inner .help {
    margin-top: 0
}
.access-token .zeroclipboard-link {
    display: inline-block
}
.access-token.new-token {
    background-color: rgba(108, 198, 68, 0.1)
}
.access-token.new-token .octicon-check {
    color: #6cc644
}
.access-token .token-description {
    max-width: 450px
}
.access-token .token {
    font-size: 14px
}
.token-scope {
    display: inline-block;
    width: 220px;
    margin: 0;
    padding: 5px 0 5px 20px;
    color: #333
}
.callback-urls dl dd input[type=text] {
    width: 100%
}
.callback-urls.has-many .callback-url-action-cell {
    display: table-cell
}
.callback-description {
    margin-top: 20px
}
.callback-description .octicon {
    padding-left: 0
}
.callback-url .label {
    display: none;
    width: 64px;
    text-align: center
}
.callback-url.is-default-callback .label {
    display: inline-block
}
.callback-url.is-default-callback .btn {
    display: none
}
.callback-url-wrap {
    display: table;
    width: 100%
}
.callback-url-field-cell {
    display: table-cell
}
.callback-url-action-cell {
    display: none;
    width: 70px;
    text-align: right
}
.boxed-group.application-show-group .logo-upload {
    float: right;
    width: 142px;
    background-color: #eee;
    position: relative
}
.boxed-group.application-show-group .logo-upload a.delete {
    position: absolute;
    right: 0;
    padding: 5px;
    display: none
}
.boxed-group.application-show-group .logo-upload a.delete:hover {
    color: #bd2c00
}
.boxed-group.application-show-group .logo-box {
    border: 1px solid #ccc;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px;
    height: 140px
}
.boxed-group.application-show-group .logo-box img {
    height: 140px;
    width: 140px;
    border-radius: 2px 2px 0 0;
    display: none
}
.boxed-group.application-show-group .logo-placeholder {
    height: 140px;
    width: 140px;
    color: #767676;
    text-align: center;
    text-shadow: 0 1px 0 #fff
}
.boxed-group.application-show-group .logo-placeholder span {
    margin: 45px 0 0
}
.boxed-group.application-show-group .logo-placeholder p {
    margin: 0;
    font-size: 16px
}
.boxed-group.application-show-group .has-uploaded-logo .logo-placeholder,
.boxed-group.application-show-group .has-uploaded-logo .or {
    display: none
}
.boxed-group.application-show-group .has-uploaded-logo:hover a.delete {
    display: block
}
.boxed-group.application-show-group .has-uploaded-logo .logo-box img {
    display: block
}
.boxed-group.application-show-group dl.form>dd input[type="text"].wide {
    width: 460px
}
.boxed-group.application-show-group dl.form>dd input[type="textarea"].short {
    height: 50px;
    min-height: 50px
}
.application-show-group .errored .note {
    display: none
}
.application-show-group .drag-and-drop {
    padding: 8px 5px 7px;
    text-align: center
}
.application-show-group .drag-and-drop img {
    vertical-align: bottom;
    margin-bottom: 1px
}
.application-show-group .drag-and-drop span {
    padding: 0
}
.application-show-group .dragover .logo-box {
    box-shadow: #c9ff00 0 0 3px
}
.application-show-group .is-uploading .loading {
    display: block
}
.application-show-group .is-uploading .default {
    display: none
}
.application-show-group .is-failed .failed-request {
    display: block
}
.application-show-group .is-failed .default {
    display: none
}
.application-show-group .is-bad-file .bad-file {
    display: block
}
.application-show-group .is-bad-file .default {
    display: none
}
.application-show-group .is-too-big .file-too-big {
    display: block
}
.application-show-group .is-too-big .default {
    display: none
}
.application-show-group .is-default .default {
    display: block
}
.security-history .security-history-timestamp {
    float: right;
    color: #767676
}
table.security-history-detail {
    width: 100%;
    font-size: 12px
}
table.security-history-detail td {
    max-width: 200px;
    word-wrap: break-word
}
.org-two-factor .btn {
    float: right;
    margin: 10px 0 0 20px
}
.org-two-factor .flash-global {
    margin-top: 0
}
.two-factor-disabled .flash-global {
    display: block
}
.settings-email .email-actions .settings-remove-email {
    float: right;
    margin-left: 5px;
    padding-right: 7px;
    padding-left: 7px;
    line-height: 24px;
    color: #bd2c00
}
.settings-email .email-actions .settings-disabled-remove-email {
    display: none
}
.settings-email:only-child .email-actions .settings-remove-email {
    display: none
}
.settings-email:only-child .email-actions .settings-disabled-remove-email {
    display: block;
    cursor: default;
    color: #999
}
.settings-email .octicon-info {
    padding-left: 5px
}
.settings-email .public.label {
    display: inline
}
.settings-email .private.label {
    display: none
}
.settings-email.private .public.label {
    display: none
}
.settings-email.private .private.label {
    display: inline
}
.settings-email .css-truncate-target {
    max-width: 300px
}
.email-preference-exceptions {
    font-size: 12px
}
.email-preference-exceptions h5 {
    margin: 15px 0 5px;
    color: #666
}
.email-preference-exceptions .exception-list {
    list-style: none
}
.email-preference-exceptions .exception {
    max-width: 400px;
    padding: 5px;
    padding-left: 0;
    border-top: 1px solid #eee
}
.email-preference-exceptions .exception:last-child {
    border-bottom: 1px solid #eee
}
.email-preference-exceptions .exception-action {
    float: right
}
.email-preference-exceptions.opt-in-list {
    display: none
}
.transactional-only .email-preference-exceptions.opt-in-list {
    display: block
}
.transactional-only .email-preference-exceptions.opt-out-list {
    display: none
}
.email-preferences .status-indicator-opt-out,
.email-preferences.transactional-only .status-indicator-opt-in {
    display: none
}
.email-preferences.transactional-only .status-indicator-opt-out {
    display: inline-block
}
.two-factor-intro {
    width: 675px;
    margin: 40px auto 0
}
.two-factor-intro .two-factor-graphic {
    margin: 20px 0
}
.two-factor-intro .two-factor-explain {
    margin: 0 0 40px;
    padding: 0;
    font-size: 13px;
    list-style: none
}
.two-factor-intro .two-factor-explain li {
    float: left;
    margin: 0;
    padding: 0
}
.two-factor-intro .two-factor-explain .step-one {
    width: 185px;
    margin-right: 36px
}
.two-factor-intro .two-factor-explain .step-two {
    width: 230px;
    margin-right: 42px
}
.two-factor-intro .two-factor-explain .step-three {
    width: 180px
}
.two-factor-graphic {
    background-image: url(https://github.com/images/modules/settings/2fa_guide.png);
    background-repeat: no-repeat;
    width: 675px;
    height: 135px
}
.two-factor-recovery-codes {
    height: 240px;
    margin-top: 15px;
    padding-left: 60px
}
.two-factor-recovery-code {
    display: inline-block;
    width: 49%;
    line-height: 1.1
}
.two-factor-recovery-code::before {
    content: "\25a1";
    font-size: 26px;
    margin-right: 10px;
    color: #eaeaea;
    position: relative;
    top: 1px
}
@media only screen and (-webkit-min-device-pixel-ratio: 2),
only screen and (min-device-pixel-ratio: 2),
only screen and (min-resolution: 2dppx) {
    .two-factor-graphic {
        background-image: url(https://github.com/images/modules/settings/2fa_guide@2x.png);
        background-size: 675px 135px
    }
}
.markdown-body .sms-or-app {
    width: 100%;
    margin: 0;
    padding: 40px 0 0;
    border-top: 1px solid #ddd
}
.markdown-body .sms-or-app:before {
    display: table;
    content: ""
}
.markdown-body .sms-or-app:after {
    display: table;
    clear: both;
    content: ""
}
.markdown-body .sms-or-app li {
    float: left;
    width: 325px;
    padding: 0;
    list-style: none
}
.markdown-body .sms-or-app li:first-child {
    margin-right: 25px
}
.markdown-body .sms-or-app li .btn {
    display: block;
    text-align: center;
    margin: 10px 0;
    padding-top: 12px;
    padding-bottom: 12px;
    font-size: 15px;
    height: 100%;
    width: 100%
}
.markdown-body .sms-or-app small {
    font-size: 80%
}
.markdown-body .app-only {
    padding: 20px 0 0
}
.markdown-body .app-only li {
    float: none;
    width: auto
}
.markdown-body .app-only li .btn {
    display: inline-block;
    width: auto;
    padding-left: 20px;
    padding-right: 20px
}
.two-factor-setup-container {
    width: 600px;
    margin: 0 auto
}
.two-factor-setup-container .form label {
    font-style: normal
}
.two-factor-setup-container .form dd {
    margin: 0;
    padding: 0
}
.two-factor-setup-container .octicon-alert {
    color: #bd2c00
}
.two-factor-setup-container .error-icon {
    position: relative;
    left: 5px;
    top: 2px;
    color: #bd2c00
}
.two-factor-setup-container .sent-message {
    position: relative;
    left: 5px;
    top: 2px;
    color: #6cc644
}
.two-factor-setup-container .select-menu {
    float: left
}
.two-factor-setup-container .select-menu .btn-sm {
    padding-top: 4px;
    padding-bottom: 4px;
    margin-right: 5px
}
.two-factor-setup-container .select-menu .btn-sm input[type="radio"],
.two-factor-setup-container .select-menu .btn-sm .country {
    display: none
}
.two-factor-setup-container .select-menu .select-menu-button:before {
    top: 14px
}
.container.two-factor-toggle {
    width: 700px
}
.two-factor-step-container {
    margin: 0 0 20px;
    font-size: 86.6667%
}
.two-factor-step-container .sms-form .form {
    margin-left: 15px
}
.two-factor-step-container .form label {
    color: #767676
}
.two-factor-step-container:last-of-type {
    margin-bottom: 0;
    padding-bottom: 0;
    border-bottom: 0
}
.two-factor-step-container h4 {
    margin: 0;
    font-size: 13px
}
.two-factor-step-container p:last-child {
    margin-bottom: 0
}
.two-factor-toggle {
    margin-top: 40px
}
.two-factor-toggle .two-factor-status {
    padding: 20px 0;
    margin: 0 0 20px;
    border-bottom: 1px solid #eaeaea;
    color: #767676
}
.two-factor-toggle .two-factor-status p {
    margin: 0
}
.two-factor-toggle .two-factor-status .btn {
    position: relative;
    top: -3px;
    float: right
}
.two-factor-toggle .two-factor-on {
    margin-right: 5px;
    padding: 3px 5px;
    border-radius: 2px;
    background-color: #6cc644;
    color: #fff;
    text-shadow: 0 1px 1px rgba(0, 0, 0, 0.1)
}
.two-factor-settings-group {
    position: relative;
    margin: 0 0 20px;
    padding: 0 0 20px 220px;
    border-bottom: 1px solid #ddd
}
.two-factor-settings-group>h3 {
    position: absolute;
    top: -15px;
    left: 0;
    width: 200px;
    font-size: 14px
}
.two-factor-settings-group>h3 .octicon {
    position: absolute;
    left: -24px;
    color: #bd2c00
}
.two-factor-settings-group li {
    list-style: none;
    line-height: 1.5
}
.u2f-registrations {
    padding-left: 24px
}
.u2f-registration {
    border-bottom: 1px solid #f8f8f8;
    margin-bottom: 8px;
    padding-bottom: 8px;
    position: relative
}
.u2f-registration.is-sending .u2f-registration-delete {
    display: none
}
.u2f-registration.is-sending .spinner {
    position: relative;
    top: 3px
}
.u2f-registration-nickname {
    font-weight: bold
}
.u2f-registration-icon {
    color: #dac268;
    position: absolute;
    left: -24px
}
.new-u2f-registration {
    position: relative
}
.new-u2f-registration .add-u2f-registration-form {
    display: none;
    margin-bottom: 10px
}
.new-u2f-registration.is-active .add-u2f-registration-link {
    display: none
}
.new-u2f-registration.is-active .add-u2f-registration-form {
    display: block
}
.new-u2f-registration .u2f-request-interaction,
.new-u2f-registration .u2f-request-error {
    display: none
}
.new-u2f-registration.is-sending .u2f-request-interaction {
    display: block
}
.new-u2f-registration.is-showing-error .u2f-request-error {
    display: block
}
.new-u2f-registration .u2f-error-icon {
    font-size: 64px
}
.u2f-box .u2f-sorry {
    display: block
}
.u2f-box .new-u2f-registration {
    display: none
}
.u2f-box.available .u2f-sorry {
    display: none
}
.u2f-box.available .new-u2f-registration {
    display: block
}
.github-access-banner {
    position: relative;
    margin: 0 0 20px;
    padding: 10px 20px 10px 70px;
    border: 1px solid #ddd;
    border-radius: 3px;
    font-size: 14px
}
.github-access-banner .mega-octicon {
    position: absolute;
    left: 20px;
    top: 20px;
    color: #bd2c00
}
.error-icon,
.spinner,
.sent-message,
.sms-error-message,
#text-code {
    display: none
}
.is-sending .spinner {
    display: inline-block
}
.is-sent .sent-message {
    display: inline-block
}
.is-not-sent .sms-error-message {
    display: block
}
.is-not-sent .error-icon {
    display: inline-block
}
.two-factor-secret {
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    font-size: 13px
}
.markdown-body .qr-code-table,
.qr-code-table {
    width: auto;
    float: right;
    margin: 0 0 0 40px;
    border: 1px solid #ddd
}
.markdown-body .qr-code-table tr,
.qr-code-table tr {
    background: transparent;
    border: 0
}
.markdown-body .qr-code-table th,
.markdown-body .qr-code-table td,
.qr-code-table th,
.qr-code-table td {
    border: 0;
    padding: 0
}
.markdown-body .qr-code-table td,
.qr-code-table td {
    width: 3px;
    height: 3px
}
.markdown-body .qr-code-table .black,
.qr-code-table .black {
    background: #000
}
.markdown-body .qr-code-table .white,
.qr-code-table .white {
    background: #fff
}
.markdown-body .two-factor-actions {
    clear: both;
    padding: 20px 0 0;
    margin: 20px 0 0;
    border-top: 1px solid #eaeaea;
    font-size: 13px
}
.markdown-body .two-factor-actions ul {
    width: 600px;
    margin: 0;
    padding: 0
}
.markdown-body .two-factor-actions li {
    list-style: none;
    display: inline-block;
    margin-right: 10px
}
.two-factor-banner {
    position: relative;
    padding-left: 60px;
    margin: 40px auto;
    width: 700px;
    background: #fff;
    border: 1px solid #ddd;
    color: #444
}
.two-factor-banner:hover {
    border-color: #ddd
}
.two-factor-banner .mega-octicon {
    position: absolute;
    top: 15px;
    left: 15px;
    color: #bd2c00
}
.two-factor-banner h2 {
    margin-top: 0;
    line-height: 32px
}
.two-factor-banner p {
    margin-top: 0
}
.two-factor-mini-banner {
    display: block;
    width: 100%;
    margin: 0 0 20px;
    padding: 15px 15px 15px 42px;
    background: #fff
}
.two-factor-mini-banner .btn-sm {
    position: relative;
    top: -4px;
    float: right
}
.two-factor-mini-banner p {
    margin-bottom: 0;
    line-height: 1.5
}
.two-factor-mini-banner .octicon {
    position: absolute;
    top: 15px;
    left: 15px;
    color: #bd2c00
}
.orgs-settings {
    margin-bottom: 15px
}
.confirmation-phrase {
    font-weight: normal;
    font-style: italic
}
.do-not-copy-me {
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none
}
li.session-device {
    position: relative;
    line-height: 18px;
    padding: 15px;
    background-color: #fafafa;
    color: #767676
}
li.session-device .btn {
    float: right;
    margin-top: 4px
}
li.session-device .session-state-indicator {
    float: left;
    width: 8px;
    height: 8px;
    margin-top: 10px;
    border-radius: 5px
}
li.session-device .session-state-indicator.recent {
    background-color: #6cc644;
    box-shadow: 0 0 10px rgba(108, 198, 68, 0.5)
}
li.session-device .session-state-indicator.not-recent {
    box-shadow: 0 1px 0 #fff;
    background-image: -webkit-linear-gradient(#aaa, #ccc);
    background-image: linear-gradient(#aaa, #ccc)
}
li.session-device .session-icon {
    float: left;
    width: 32px;
    margin-top: 1px;
    margin-left: 15px;
    text-align: center;
    color: #bbb
}
li.session-device .sessions-more-info {
    position: relative;
    display: none;
    margin-top: 10px
}
li.session-device.session-current {
    background-color: #fff
}
li.session-device.session-current .session-last-accessed {
    color: #767676
}
li.session-device.session-current .sessions-more-info {
    color: #767676
}
li.session-device.session-current .sessions-more-info:after {
    border-top-color: #fff
}
li.session-device.session-current .mega-octicon {
    color: #767676
}
.session-details {
    position: relative;
    width: 350px;
    margin-left: 70px
}
.session-details:hover .octicon {
    color: #4078c0;
    cursor: pointer
}
.session-details.open .sessions-more-info {
    display: block
}
.session-title {
    display: block
}
.collaborators .collab-list {
    border-bottom-width: 0
}
.collaborators .collab-list-item:first-child .collab-list-cell {
    border-top-width: 0
}
.collaborators .collab-list-cell {
    padding-top: 15px;
    padding-bottom: 15px;
    vertical-align: middle
}
.collaborators .collab-meta {
    width: 140px
}
.collaborators .collab-permission {
    text-align: center
}
.collaborators .collab-remove {
    padding-right: 20px;
    text-align: right
}
.collaborators .collab-remove .remove-link {
    color: #767676
}
.collaborators .collab-remove .remove-link:hover {
    color: #bd2c00
}
.collaborators .collab-team-link {
    width: 300px
}
.collaborators .collab-team-link:hover {
    text-decoration: none
}
.collaborators .collab-team-link .avatar {
    float: left;
    margin-top: 1px;
    margin-right: 10px
}
.collaborators .collab-team-link.disabled {
    pointer-events: none
}
.collaborators .collab-info {
    color: #666
}
.collaborators .collab-info .description {
    padding-right: 50px;
    margin-top: 3px;
    margin-bottom: 3px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap
}
.collaborators .collab-info .collab-name {
    display: block;
    font-size: 14px
}
.access-sub-heading {
    float: right;
    font-weight: normal;
    line-height: 1.4;
    color: #767676
}
.access-form-wrapper {
    padding: 10px;
    background-color: #fcfcfc;
    border-top: 1px solid #ddd;
    border-radius: 0 0 3px 3px
}
.access-flash {
    margin-bottom: 10px;
    margin-left: 10px;
    margin-right: 10px;
    padding: 8px
}
.repo-access-group .blankslate {
    display: none
}
.repo-access-group.is-empty .blankslate {
    display: block
}
.repo-access-group.no-form .add-team-form {
    display: none
}
.repo-access-group .select-menu-item.has-access {
    display: none
}
.oauth-pending-deletion-list-item {
    background-color: #fafafa;
    box-shadow: inset 0 0 8px #eee
}
.oauth-pending-deletion-list-item:hover {
    background-color: #fafafa
}
.oauth-pending-deletion-list-item .oauth-pending-deletion {
    display: inline
}
.oauth-pending-deletion-list-item .active {
    display: none
}
.oauth-pending-deletion {
    display: none;
    width: 100%
}
.boxed-group-list .access-level {
    color: #767676
}
.boxed-group-list .access-level.css-truncate-target {
    max-width: 500px
}
.icon-shield {
    display: inline-block;
    width: 14px;
    height: 16px;
    background: url(https://github.com/images/icons/shield-16.png) no-repeat center center 14px 16px;
    opacity: 0.35
}
.icon-shield.mega {
    width: 28px;
    height: 32px;
    background-image: url(https://github.com/images/icons/shield-32.png);
    background-size: 28px 32px
}
.settings-next {
    font-size: 14px;
    line-height: 1.5
}
.settings-next .form-checkbox label {
    font-size: 14px
}
.settings-next .form-checkbox .note {
    font-size: 13px
}
.settings-next .form-checkbox input[type=radio],
.settings-next .form-checkbox input[type=checkbox] {
    margin-top: 4px
}
.form-hr {
    margin-top: 20px;
    margin-bottom: 20px;
    border-bottom-color: #e5e5e5
}
.listgroup {
    border: 1px solid #e5e5e5;
    border-radius: 3px
}
.listgroup-item {
    padding: 10px;
    font-size: 13px;
    color: #767676;
    line-height: 26px
}
.listgroup-item:before {
    display: table;
    content: ""
}
.listgroup-item:after {
    display: table;
    clear: both;
    content: ""
}
.listgroup-item+.listgroup-item {
    border-top: 1px solid #e5e5e5
}
.listgroup-header {
    border-top: 0;
    border-bottom: 1px solid #e5e5e5
}
.listgroup-overflow {
    max-height: 240px;
    overflow-y: auto;
    background-color: #f5f5f5
}
.listgroup-sm .listgroup-item {
    padding-top: 5px;
    padding-bottom: 5px
}
.protected-branches {
    margin-top: 15px;
    margin-bottom: 15px
}
.protected-branch-options {
    opacity: 0.5
}
.protected-branch-options.active {
    opacity: 1
}
.automated-check-options {
    margin-top: 10px
}
.automated-check-options .listgroup-item label {
    font-size: inherit
}
.automated-check-options .listgroup-item input[type="checkbox"] {
    float: none;
    margin-top: -2px;
    margin-left: 0;
    margin-right: 5px
}
.automated-check-options .label {
    margin-top: 4px
}
.automated-check-options .check-details {
    display: none
}
.logged_out.signup .header-logged-out .container,
.logged_out.signup .site-footer {
    width: 750px
}
.logged_out.signup .site-footer {
    margin-right: auto;
    margin-left: auto
}
.logged_out.signup .site-footer .octicon-mark-github {
    top: 30px
}
.logged_out.signup .header-actions .primary,
.logged_out.signup .site-footer-links,
.logged_out.signup .site-search {
    display: none
}
.setup-wrapper {
    width: 750px;
    padding-top: 30px;
    margin: 0 auto;
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif
}
.setup-header {
    overflow: hidden;
    padding-bottom: 20px;
    margin: 0 auto 30px;
    text-align: left;
    text-shadow: 0 1px 0 white;
    border-bottom: 1px solid #ddd
}
.setup-header h1 {
    margin-top: 0;
    margin-bottom: 0;
    font-size: 45px;
    font-weight: normal;
    letter-spacing: -1px
}
.setup-header h1 .mega-octicon {
    color: #bbb
}
.setup-header .lead {
    margin-top: 2px;
    margin-bottom: 0;
    font-size: 21px
}
.setup-header .lead a {
    color: #767676
}
.setup-header .lead a:hover {
    color: #4078c0;
    text-decoration: none
}
.setup-org {
    padding-bottom: 0;
    border-bottom: 0
}
.setup-main {
    float: left;
    width: 450px
}
.setup-main.without-secondary {
    margin-left: 150px
}
.setup-secondary {
    float: right;
    width: 250px
}
.setup-secondary .info {
    padding-top: 0;
    padding-bottom: 0;
    margin-top: -10px;
    color: #767676;
    font-size: 12px;
    line-height: 18px;
    text-align: center
}
.setup-info-module {
    margin-bottom: 30px;
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 3px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.075)
}
.setup-info-module h2 {
    padding: 15px;
    margin-top: 0;
    margin-bottom: 15px;
    overflow: hidden;
    font-size: 16px;
    border-bottom: 1px solid #ddd
}
.setup-info-module h2 .price {
    float: right;
    font-weight: bold;
    color: #767676
}
.setup-info-module h3 {
    padding: 0 15px;
    margin: 0 0 -7px;
    font-size: 14px
}
.setup-info-module p {
    padding: 0 15px;
    margin: 15px 0
}
.setup-info-module .setup-section-title {
    margin-bottom: 10px
}
.setup-info-module .setup-info-note {
    background: #f9f9f9;
    padding: 1px 0;
    margin: 0;
    border-top: 1px solid #e0e0e0
}
.features-list {
    padding: 0 15px 15px;
    margin: 0;
    font-size: 14px;
    list-style: none
}
.features-list li {
    margin-top: 10px
}
.features-list li:first-child {
    margin-top: 0
}
.features-list .list-divider {
    margin: 15px -15px;
    border-top: 1px solid #eee
}
.features-list .octicon-check {
    margin-right: 5px;
    color: #60b044
}
.features-list .octicon-question {
    color: #555;
    font-size: 12px
}
.features-list .tooltipped:after {
    width: 250px;
    white-space: normal
}
.features-list.features-list-org {
    padding-bottom: 0
}
.setup-form-container .setup-form-title {
    margin-top: 0;
    font-size: 16px
}
.setup-form-container .secure {
    float: right;
    margin-top: 2px;
    color: #60b044;
    font-size: 11px;
    text-transform: uppercase
}
.setup-form-container hr {
    margin-top: 25px;
    margin-bottom: 25px
}
.setup-form-container .form-actions {
    padding-top: 0;
    padding-bottom: 0;
    text-align: left
}
.team-member-container {
    margin-bottom: 20px
}
.team-member-container .team-member-username {
    line-height: 1.2
}
.setup-form {
    padding-bottom: 15px
}
.setup-form .form dd input[type="text"],
.setup-form .form dd input[type="password"],
.setup-form .form dd input[type="email"] {
    width: 100%
}
.setup-form .form dd input[type="text"].short,
.setup-form .form dd input[type="password"].short,
.setup-form .form dd input[type="email"].short {
    width: 250px
}
.setup-form dd {
    position: relative
}
.setup-form dd .octicon {
    position: absolute;
    right: 25px;
    top: 8px
}
.setup-form .octicon-alert:before {
    color: #bd2c00
}
.setup-form .octicon-check:before {
    color: #6cc644
}
.setup-form .text-muted {
    margin-top: 5px
}
.setup-form .tos-info,
.setup-form .setup-organization-next {
    margin: 15px 0;
    border-top: 1px solid #eee;
    border-bottom: 1px solid #eee
}
.setup-form .tos-info {
    padding: 15px 0
}
.setup-form .setup-organization-next {
    padding-top: 15px;
    padding-bottom: 15px
}
.setup-form .setup-plans {
    margin-bottom: 25px;
    border: solid 1px #ccc;
    border-collapse: separate;
    border-radius: 3px;
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.075);
    overflow: hidden
}
.setup-form .setup-plans tr.selected {
    background-color: #f0f7fd
}
.setup-form .setup-plans th,
.setup-form .setup-plans td {
    vertical-align: middle;
    border-bottom: 1px solid #ccc
}
.setup-form .setup-plans .name {
    font-weight: bold
}
.setup-form .setup-plans .choose_plan input[type=radio] {
    display: none
}
.setup-creditcard-form .cc-extras {
    margin-bottom: 15px
}
.setup-creditcard-form .expiration-form {
    width: 120px
}
.setup-creditcard-form .expiration-form dd {
    line-height: 32px
}
.setup-creditcard-form .expiration-form,
.setup-creditcard-form .cvv-form,
.setup-creditcard-form .country-form,
.setup-creditcard-form .state-form {
    word-wrap: normal;
    float: left;
    margin: 0
}
.setup-creditcard-form .form dd input.input-cvv {
    width: 130px
}
.setup-creditcard-form .form select.select-country {
    width: 162px;
    margin-right: 5px
}
.setup-creditcard-form .form select.select-state {
    width: 84px
}
.setup-creditcard-form.is-vat-country .vat-field {
    display: block
}
.setup-creditcard-form.is-international .form select.select-country {
    width: 250px
}
.setup-creditcard-form.is-international .state-form {
    display: none
}
.setup-creditcard-form dd .octicon-credit-card {
    position: inherit
}
.setup-creditcard-form .enter-new-card {
    display: none
}
.setup-creditcard-form.has-credit-card .enter-new-card {
    display: inline-block
}
.setup-creditcard-form.has-credit-card .card-select-number-field,
.setup-creditcard-form.has-credit-card .cancel-enter-new-card,
.setup-creditcard-form.has-credit-card .cards-select {
    display: none
}
.setup-creditcard-form .vat-field {
    display: none
}
.setup-creditcard-form .vat-field.prefilled {
    display: block
}
.setup-creditcard-form .help-text {
    font-size: 80%;
    font-weight: normal;
    color: #767676
}
.collection-head {
    height: 225px;
    margin-top: -20px;
    margin-bottom: 20px;
    background: #555 url(https://github.com/images/modules/home/octicons-bg.png) center repeat;
    box-shadow: inset 0 10px 20px rgba(0, 0, 0, 0.1);
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    color: #fff
}
.collection-head .collection-info {
    margin: 0
}
.collection-head .collection-info .meta-info {
    margin-right: 15px
}
.collection-head .collection-info .avatar {
    background-color: rgba(255, 255, 255, 0.7);
    border: 1px solid rgba(255, 255, 255, 0.7)
}
.collection-head .container {
    position: relative
}
.collection-head .draft-tag {
    position: absolute;
    top: 0;
    left: 0
}
.collection-title {
    display: table-cell;
    height: 225px;
    vertical-align: middle
}
.collection-header {
    margin-top: 0;
    font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
    font-size: 45px;
    font-weight: normal
}
.collection-description {
    position: relative;
    font-size: 16px
}
.collection-page .collection-info {
    margin-top: 10px;
    margin-bottom: 20px;
    font-size: 13px;
    color: #767676
}
.collection-page .column.main {
    margin-right: 260px !important
}
.collection-page .column.sidebar {
    width: 240px
}
.collection-page .other-content {
    padding: 20px 0 20px 20px;
    border-left: 1px solid #f1f1f1
}
.collection-page .other-content .subnav-search {
    margin-left: 0
}
.collection-page .other-content input.subnav-search-input {
    width: 100%
}
.collection-page .other-content-title {
    margin-top: 40px
}
.collection-page .other-content-title:first-child {
    margin-top: 0
}
.side-collection-list {
    margin: 0;
    list-style-type: none
}
.side-collection-link {
    display: table;
    width: 100%;
    height: 100px;
    color: #fff
}
.side-collection-item-title {
    font-size: 16px;
    font-weight: 100
}
.side-collection-image {
    background: #555 url(https://github.com/images/modules/home/octicons-bg.png) center repeat;
    box-shadow: inset 0 10px 20px rgba(0, 0, 0, 0.1);
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    color: #fff;
    display: table-cell;
    width: 100%;
    height: 100%;
    margin-bottom: 5px;
    text-align: center;
    vertical-align: middle;
    border-radius: 3px
}
.side-collection-list-item {
    margin-bottom: 20px
}
.collection-tools {
    list-style-type: none;
    margin-bottom: 10px;
    font-size: 15px
}
.collection-tools .edit-link {
    color: #333
}
.collection-tools .edit-link:hover {
    color: #4078c0;
    cursor: pointer
}
.collection-tools .octicon {
    margin-right: 5px
}
.collection-tools .select-menu-button {
    position: relative;
    display: inline-block;
    color: #333
}
.collection-tools .select-menu-button:hover {
    color: #4078c0;
    cursor: pointer
}
.collection-tool {
    margin-left: 20px
}
.collection-search-results em {
    padding: 0.1em;
    background-color: #faffa6
}
.collection-search-result {
    margin-bottom: 40px;
    list-style-type: none
}
.collection-search-result-title {
    margin-top: 0
}
.collection-search-page .search-results-info {
    line-height: 33px;
    float: right;
    margin-left: 10px;
    font-size: 15px
}
.draft-tag {
    padding: 5px 10px;
    font-weight: bold;
    color: #eee;
    background-color: #404040
}
.collection-card {
    position: relative;
    float: left;
    width: 313px;
    margin-right: 20px;
    margin-bottom: 20px;
    list-style-type: none;
    background: #f7f7f7;
    border: 1px solid #ddd;
    border-radius: 3px
}
.collection-card .draft-tag {
    position: absolute;
    top: -1px;
    left: 10px
}
.collection-card:nth-child(3n+3) {
    margin-right: 0
}
.collection-card-title {
    padding: 0 15px;
    margin: 10px 0;
    display: table-cell;
    width: 100%;
    height: 100%;
    font-size: 19px;
    font-weight: bold;
    text-align: center;
    vertical-align: middle
}
.collection-card-body {
    padding: 0 15px;
    margin: 10px 0;
    height: 6em;
    margin-top: 0;
    overflow: hidden;
    font-size: 15px;
    line-height: 1.5em
}
.collection-card-image {
    position: relative;
    display: table;
    width: 313px;
    height: 120px;
    margin: -1px;
    margin-bottom: 15px;
    background: #555 url(https://github.com/images/modules/home/octicons-bg.png) center repeat;
    box-shadow: inset 0 10px 20px rgba(0, 0, 0, 0.1);
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    color: #fff;
    border-top-right-radius: 3px;
    border-top-left-radius: 3px
}
.collection-card-meta {
    padding: 0 15px;
    margin-top: 5px;
    margin-bottom: 15px;
    color: #767676
}
.collection-card-meta .meta-info {
    margin-right: 10px
}
.collection-card-meta .last-updated {
    float: right;
    margin-right: 0
}
.collection-listing-search {
    margin-bottom: 20px
}
.collection-listing-search .subnav-search {
    margin-left: 0;
    margin-right: 25%
}
.collection-feed-icon {
    float: right;
    margin-top: 28px
}
.featured-grid {
    position: relative;
    list-style: none;
    margin-top: -10px
}
.featured-grid-outer {
    position: relative;
    height: 100%
}
.featured-grid-more-info {
    padding: 20px
}
.featured-showcase-meta {
    position: absolute;
    bottom: 15px;
    left: 20px
}
.featured-showcase-meta .meta-info {
    margin-right: 10px
}
.featured-grid-link {
    display: table;
    width: 100%;
    height: 100%;
    border-radius: 5px;
    background: #555 url(https://github.com/images/modules/home/octicons-bg.png) center repeat;
    box-shadow: inset 0 10px 20px rgba(0, 0, 0, 0.1);
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
    color: #fff
}
.featured-grid-inner {
    display: table-cell;
    padding: 10px 20px;
    font-size: 15px;
    text-align: center;
    vertical-align: middle
}
.grid-item {
    position: relative;
    display: block;
    float: left;
    width: 25%;
    height: 122.5px;
    padding: 10px
}
.grid-item[data-grid-item="0"] {
    position: absolute;
    width: 50%;
    height: 245px
}
.grid-item[data-grid-item="1"],
.grid-item[data-grid-item="2"] {
    margin-right: 25%;
    margin-left: 50%
}
.grid-item[data-grid-item="3"] {
    height: 245px
}
.grid-item[data-grid-item="4"] {
    width: 50%;
    height: 245px
}
.grid-item[data-grid-item="7"] {
    position: absolute;
    top: 0;
    right: 0;
    height: 245px
}
.showcase-featured .see-more {
    text-align: center
}
.showcase-featured .in-yo-face .featured-grid-outer {
    overflow: hidden
}
.showcase-featured .in-yo-face .showcase-info {
    position: absolute;
    right: 0;
    bottom: -20%;
    left: 0;
    padding: 10px 20px;
    font-size: 13px;
    text-align: left;
    background: rgba(0, 0, 0, 0.6);
    border-bottom-right-radius: 3px;
    border-bottom-left-radius: 3px;
    opacity: 0;
    visibility: hidden;
    -webkit-transition: all 0.2s ease-in-out;
    transition: all 0.2s ease-in-out
}
.showcase-featured .in-yo-face .showcase-name {
    font-size: 25px;
    color: #fff
}
.showcase-featured .in-yo-face .meta-info {
    margin-right: 10px
}
.showcase-featured .in-yo-face:hover .showcase-info {
    bottom: 0;
    opacity: 1;
    visibility: visible
}
.showcase-featured .mo-info .featured-grid-outer {
    background: #f7f7f7;
    border: 1px solid #ddd;
    border-radius: 3px
}
.showcase-featured .mo-info .featured-grid-link {
    width: 225px;
    height: 102px;
    margin-top: -1px;
    margin-right: -1px;
    margin-left: -1px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0
}
.showcase-featured .normal-intensity .showcase-info {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    padding-top: 43px;
    font-size: 13px;
    text-align: center;
    vertical-align: middle;
    background: rgba(0, 0, 0, 0.6);
    border-radius: 3px;
    opacity: 0;
    -webkit-transition: opacity 0.3s ease-in-out;
    transition: opacity 0.3s ease-in-out
}
.showcase-featured .normal-intensity .octicon {
    display: inline
}
.showcase-featured .normal-intensity .meta-info {
    margin-right: 10px
}
.showcase-featured .normal-intensity .text {
    display: none
}
.showcase-featured .normal-intensity .name-text {
    display: block
}
.showcase-featured .normal-intensity a:hover {
    text-decoration: none
}
.showcase-featured .normal-intensity:hover .showcase-info {
    opacity: 1
}
.showcase-featured .normal-intensity:hover .name-text {
    color: transparent;
    text-shadow: 0 0 5px rgba(255, 255, 255, 0.8)
}
.showcase-featured .normal-intensity .showcase-name,
.showcase-featured .mo-info .showcase-name {
    font-size: 16px;
    font-weight: 100
}
.signup-plans-actions {
    margin: -10px 0 20px
}
.btn.plans-signup-button {
    padding: 12px 50px;
    font-size: 16px
}
.signup-plans-or {
    margin: 0 5px 0 8px
}
.signup-plans-collabocat {
    float: right;
    width: 300px;
    margin-left: 30px
}
.signup-plans {
    width: 100%;
    margin: 20px 0 40px;
    border-collapse: separate
}
.signup-plans th,
.signup-plans td {
    text-align: left;
    padding: 9px;
    font-size: 14px;
    border: solid #e5e5e5
}
.signup-plans th {
    padding: 14px 9px;
    font-size: 20px;
    border-width: 1px 0 0 1px
}
.signup-plans th small {
    display: block;
    font-size: 14px;
    color: #7a7a7a
}
.signup-plans thead .empty {
    border-width: 0
}
.signup-plans thead .plan-free {
    border-radius: 5px 0 0
}
.signup-plans thead th:last-child {
    border-right-width: 1px;
    border-radius: 0 5px 0 0
}
.signup-plans tbody tr:first-child td:first-child {
    border-top-width: 1px;
    border-radius: 5px 0 0
}
.signup-plans td {
    border-width: 1px 0 0 1px
}
.signup-plans td:last-child {
    border-right-width: 1px
}
.signup-plans tr:last-child td {
    border-bottom-width: 1px
}
.signup-plans tr:last-child td:first-child {
    border-radius: 0 0 0 5px
}
.signup-plans tr:last-child td:last-child {
    border-radius: 0 0 5px
}
.signup-plans tr:nth-child(odd) td {
    background-color: #f5f5f5
}
.signup-plans .row-label {
    width: 18%;
    font-weight: bold
}
.signup-plans-personal th {
    color: #4078c0
}
.signup-plans-personal tr:nth-child(odd) td {
    background-color: #edf2f9
}
.signup-plans-orgs th {
    color: #6cc644
}
.signup-plans-orgs tr:nth-child(odd) td {
    background-color: #f6fcf4
}
.signup-plans-toggle-currency {
    float: right
}
.signup-plans-currency-notice {
    margin: 10px auto 30px;
    width: 800px;
    text-align: center
}
.simple-stacked-bar {
    display: table;
    height: 10px;
    width: 100%;
    background-color: #eee
}
.bar-section {
    display: table-cell
}
.bar-section[style='width:0.0%'] {
    display: none
}
.bar-section-positive {
    background-color: #6cc644
}
.bar-section-negative {
    background-color: #bd2c00
}
.bar-section-alt {
    background-color: #6e5494
}
.stars-browser .sort-bar .filter_input {
    width: 400px
}
.stars-browser .repo-list {
    margin-top: -20px
}
.facebox .sudo {
    padding: 0
}
.facebox .sudo .auth-form-header {
    border-width: 0 0 1px
}
.facebox .sudo .auth-form-header .mini-icon {
    display: none
}
.facebox .sudo .auth-form-body {
    border-width: 0
}
.facebox .sudo+.facebox-close {
    padding: 5px;
    color: #fff
}
.sudo-prompt,
.sudo-error {
    display: none
}
.survey-flow .header {
    background: none;
    border-bottom-color: rgba(255, 255, 255, 0.2)
}
.survey-flow .header-logo-invertocat {
    position: relative;
    left: 50%;
    margin-left: -14px;
    color: #fff
}
.survey-flow .site-search,
.survey-flow .header-nav,
.survey-flow .flash-global,
.survey-flow .site-footer {
    display: none
}
.survey-background {
    position: fixed;
    top: 0;
    left: 0;
    z-index: -1;
    width: 100%;
    height: 100%;
    background-image: -webkit-linear-gradient(#1dadee, #86d1ee);
    background-image: linear-gradient(#1dadee, #86d1ee)
}
.survey-content {
    position: relative
}
.survey-screen {
    position: absolute;
    top: 0;
    width: 100%;
    z-index: 1;
    opacity: 0
}
.survey-screen.out {
    z-index: 0;
    -webkit-animation: fadeOutUp 0.8s ease-in-out 0s;
    animation: fadeOutUp 0.8s ease-in-out 0s
}
.survey-screen.in {
    opacity: 1;
    -webkit-animation: fadeInUpShort 0.8s ease-in-out 0s;
    animation: fadeInUpShort 0.8s ease-in-out 0s
}
.survey-intro,
.survey-outro {
    margin-top: 170px
}
.survey-intro .survey-title,
.survey-intro .lead,
.survey-outro .survey-title,
.survey-outro .lead {
    color: #fff;
    text-shadow: 0 2px 3px rgba(75, 138, 190, 0.8)
}
.survey-intro .survey-title,
.survey-outro .survey-title {
    margin-bottom: 0;
    font-size: 50px;
    font-weight: 400;
    letter-spacing: -1px
}
.survey-intro .lead,
.survey-outro .lead {
    margin: 0 auto 30px;
    max-width: 600px;
    font-weight: normal
}
.survey-intro .btn,
.survey-outro .btn {
    padding: 10px 15px;
    font-size: 16px
}
.survey-questions {
    max-width: 600px;
    margin: 90px auto 0;
    position: relative
}
.survey-question {
    width: 100%;
    background-color: #fafafa;
    border: 1px solid #859acf;
    border-radius: 6px;
    box-shadow: 0 1px 2px 1px rgba(0, 0, 0, 0.05);
    position: absolute;
    opacity: 0;
    display: none
}
.survey-question.active {
    -webkit-animation: fadeInUpShort 0.3s ease-in-out 0s;
    animation: fadeInUpShort 0.3s ease-in-out 0s;
    opacity: 1;
    display: block
}
.survey-question.complete {
    -webkit-animation: fadeOutUp 0.3s ease-in-out 0s;
    animation: fadeOutUp 0.3s ease-in-out 0s;
    z-index: 0;
    display: block
}
.survey-question-1 {
    z-index: 9
}
.survey-question-2 {
    z-index: 8
}
.survey-question-3 {
    z-index: 7
}
.survey-question-4 {
    z-index: 6
}
.survey-question-5 {
    z-index: 5
}
.survey-question-6 {
    z-index: 4
}
.survey-question-7 {
    z-index: 3
}
.survey-question-8 {
    z-index: 2
}
.survey-question-9 {
    z-index: 1
}
.survey-question-10 {
    z-index: 0
}
.skip-survey-link {
    position: absolute;
    top: 4px;
    right: 7px;
    padding: 5px;
    color: #666
}
.survey-question-index {
    margin-top: 20px;
    font-size: 16px;
    color: #999
}
.survey-question-title {
    margin-top: -15px;
    padding: 0 10px;
    font-size: 20px
}
.survey-choices {
    text-align: left;
    padding: 0 100px;
    display: inline-block
}
.survey-choices .survey-choice {
    display: block;
    margin: 5px 0
}
.survey-choices .survey-choice-label {
    font-weight: normal;
    position: relative;
    padding-left: 20px;
    display: inline-block
}
.survey-choices .survey-choice-radio {
    position: absolute;
    top: 1px;
    left: 0
}
.survey-choice-other-field {
    display: inline-block
}
.survey-action {
    margin-top: 20px;
    padding: 15px 100px 8px;
    width: 100%;
    background: #f6f6f6;
    border-radius: 0 0 6px 6px;
    border-top: 1px solid #eaeaea
}
.survey-action .btn {
    margin-bottom: 7px
}
.survey-spinner {
    position: absolute;
    z-index: 1;
    top: 0;
    left: 0;
    display: none;
    width: 100%;
    height: 100%;
    line-height: 225px;
    background-color: rgba(255, 255, 255, 0.9);
    border-radius: 4px
}
.survey-clouds {
    -webkit-animation: fadeInUp 1s ease 0s;
    animation: fadeInUp 1s ease 0s;
    -webkit-transition: all 0.8s ease-in-out;
    transition: all 0.8s ease-in-out
}
.quiz-active .survey-clouds {
    -webkit-animation: fadeOutUp 0.8s ease-in-out 0s;
    animation: fadeOutUp 0.8s ease-in-out 0s;
    opacity: 0
}
.survey-cloud {
    position: absolute;
    z-index: 0;
    -webkit-animation: float 6s ease 0s infinite;
    animation: float 6s ease 0s infinite
}
.survey-cloud,
.cloud-center {
    -webkit-transition-timing-function: cubic-bezier(0.01, 0.38, 1, 0.73);
    transition-timing-function: cubic-bezier(0.01, 0.38, 1, 0.73)
}
.cloud-right {
    right: 20px;
    top: 400px;
    height: 175px
}
.cloud-center {
    right: 260px;
    top: 490px;
    height: 140px;
    -webkit-animation: floatInverted 5s ease 0s infinite;
    animation: floatInverted 5s ease 0s infinite
}
.cloud-left {
    top: 360px;
    left: 0;
    height: 225px
}
@-webkit-keyframes fadeOutUp {
    0% {
        opacity: 1
    }
    100% {
        opacity: 0;
        -webkit-transform: translate3d(0, -100%, 0);
        transform: translate3d(0, -100%, 0)
    }
}
@keyframes fadeOutUp {
    0% {
        opacity: 1
    }
    100% {
        opacity: 0;
        -webkit-transform: translate3d(0, -100%, 0);
        transform: translate3d(0, -100%, 0)
    }
}
@-webkit-keyframes fadeInUp {
    0% {
        opacity: 0;
        -webkit-transform: translate3d(0, 1000px, 0);
        transform: translate3d(0, 1000px, 0)
    }
    100% {
        opacity: 1;
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@keyframes fadeInUp {
    0% {
        opacity: 0;
        -webkit-transform: translate3d(0, 1000px, 0);
        transform: translate3d(0, 1000px, 0)
    }
    100% {
        opacity: 1;
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@-webkit-keyframes fadeInUpShort {
    0% {
        opacity: 0;
        -webkit-transform: translate3d(0, 200px, 0);
        transform: translate3d(0, 200px, 0)
    }
    50% {
        opacity: 0
    }
    100% {
        opacity: 1;
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@keyframes fadeInUpShort {
    0% {
        opacity: 0;
        -webkit-transform: translate3d(0, 200px, 0);
        transform: translate3d(0, 200px, 0)
    }
    50% {
        opacity: 0
    }
    100% {
        opacity: 1;
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@-webkit-keyframes float {
    0% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
    65% {
        -webkit-transform: translate3d(0, 10px, 0);
        transform: translate3d(0, 10px, 0)
    }
    100% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@keyframes float {
    0% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
    65% {
        -webkit-transform: translate3d(0, 10px, 0);
        transform: translate3d(0, 10px, 0)
    }
    100% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@-webkit-keyframes floatInverted {
    0% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
    65% {
        -webkit-transform: translate3d(0, -5px, 0);
        transform: translate3d(0, -5px, 0)
    }
    100% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
@keyframes floatInverted {
    0% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
    65% {
        -webkit-transform: translate3d(0, -5px, 0);
        transform: translate3d(0, -5px, 0)
    }
    100% {
        -webkit-transform: translate3d(0, 0, 0);
        transform: translate3d(0, 0, 0)
    }
}
.tree-browser {
    width: 100%;
    margin: 0;
    border-bottom: 1px solid #cacaca;
    border-left: 0;
    border-right: 0
}
.tree-browser td {
    background: #f8f8f8;
    border-bottom: 1px solid #eee;
    padding: 7px 3px;
    color: #484848;
    vertical-align: middle;
    white-space: nowrap
}
.tree-browser td.icon {
    width: 17px;
    padding-right: 2px;
    padding-left: 10px
}
.tree-browser td:first-child {
    border-left: 1px solid #cacaca
}
.tree-browser td:last-child {
    border-right: 1px solid #cacaca
}
.tree-browser td a.message {
    color: #484848
}
.tree-browser td span.ref {
    color: #aaa
}
.tree-browser img {
    vertical-align: text-bottom
}
.tree-browser tbody tr:last-child td {
    border-bottom: 0
}
.tree-browser .history {
    float: right;
    padding-right: 5px
}
.tree-browser .octicon-chevron-right {
    color: transparent
}
.tree-browser tr.navigation-focus td {
    background-color: #fff
}
.tree-browser tr.navigation-focus td .octicon-chevron-right {
    color: #4078c0
}
.tree-browser .octicon-file-directory {
    color: #80a6cd
}
.tree-browser .octicon-file-submodule {
    color: #3cbf5e
}
.tree-browser .octicon-file-text {
    color: #767676
}
.tree-browser .content {
    max-width: 220px
}
.tree-browser .message {
    max-width: 420px
}
.tree-browser .css-truncate-target {
    max-width: 100%
}
.tree-browser-result-template {
    display: none
}
.tree-browser-result .css-truncate-target {
    max-width: 870px
}
.tree-browser-result mark {
    background-color: transparent;
    color: #4078c0;
    font-weight: bold
}
input.tree-finder-input,
input.tree-finder-input:focus {
    position: relative;
    top: 1px;
    border: 0;
    padding: 0;
    outline: none;
    font-size: 100%;
    box-shadow: none;
    min-height: 0;
    height: 22px;
    line-height: 1px;
    vertical-align: top;
    -webkit-appearance: none;
    margin-left: 5px
}
.tree-finder .no-results {
    display: none
}
.tree-finder .no-results th {
    text-align: center
}
.tree-finder tr td.icon {
    cursor: pointer
}
.tree-finder .tree-browser {
    border-top: 1px solid #cacaca
}
.tree-finder .filterable-empty+.no-results {
    display: block
}
#user-content-toc {
    overflow: visible
}
#user-content-toc tr {
    border-top: 0
}
#user-content-toc td {
    padding: 0 20px;
    background-color: #f7f7f7;
    border: 0;
    border-radius: 3px
}
#user-content-toc ul {
    padding-left: 0;
    font-weight: bold;
    list-style: none
}
#user-content-toc ul li {
    padding-left: 0.2em
}
#user-content-toc ul ul {
    font-weight: normal
}
#user-content-toc ul ul li:before {
    float: left;
    margin-top: -0.2em;
    margin-right: 0.2em;
    font-size: 1.2em;
    line-height: 1;
    color: #aaa;
    content: "\231e"
}
#user-content-toc ul ul ul {
    padding-left: 0.9em
}
#user-content-toctitle h2 {
    margin-top: 1em;
    margin-bottom: 0.5em;
    font-size: 1.25em;
    border-bottom: 0
}
.user-list em {
    font-weight: bold;
    background-color: rgba(255, 255, 140, 0.5);
    padding: 3px;
    border-radius: 3px;
    font-style: normal
}
.user-list .avatar {
    position: absolute;
    top: 0;
    left: 0
}
.user-list-info {
    padding: 0;
    min-height: 48px;
    font-weight: normal;
    font-size: 18px;
    line-height: 20px
}
.user-list-meta {
    font-size: 11px;
    margin: 8px 0 0;
    list-style-type: none;
    overflow: hidden;
    color: #999
}
.user-list-meta>li {
    float: left;
    margin-right: 10px
}
.user-list-meta a {
    color: #333
}
.user-list-item {
    border-bottom: 1px solid #f1f1f1;
    padding: 0 0 20px 58px;
    margin: 0 0 20px;
    position: relative
}
.follow-list {
    list-style-type: none
}
.follow-list .follow-list-container {
    margin-left: 90px
}
.follow-list .follow-list-item {
    float: left;
    width: 305px;
    padding-bottom: 20px;
    margin-bottom: 20px;
    margin-right: 20px;
    border-bottom: 1px solid #eee
}
.repository-with-sidebar .follow-list .follow-list-item {
    width: 285px
}
.follow-list .follow-list-name {
    font-weight: normal;
    margin: 0 0 5px
}
.follow-list .follow-list-name a {
    color: inherit
}
.follow-list .follow-list-info {
    margin-top: 0;
    margin-bottom: 0.6em;
    font-size: 12px;
    color: #767676
}
.follow-list .css-truncate.css-truncate-target {
    max-width: 195px
}
.repository-with-sidebar .follow-list .css-truncate.css-truncate-target {
    max-width: 170px
}
.follow-list .gravatar {
    float: left;
    display: block;
    width: 75px;
    height: 75px
}
#wiki-body {
    margin-top: 20px
}
#wiki-body .markdown-body {
    padding: 0 30px;
    margin: 0 -30px
}
#wiki-rightbar {
    float: right;
    width: 230px
}
#wiki-rightbar .markdown-body {
    font-size: 13px
}
#wiki-rightbar .markdown-body .anchor {
    display: none
}
#wiki-rightbar .markdown-body h1 {
    padding-bottom: 5px;
    font-size: 1.6em;
    line-height: 1.2;
    border-color: #eee
}
#wiki-rightbar .markdown-body h2 {
    padding-bottom: 5px;
    font-size: 1.4em;
    line-height: 1.2;
    border-color: #eee
}
#wiki-rightbar .markdown-body h3,
#wiki-rightbar .markdown-body h4,
#wiki-rightbar .markdown-body h5,
#wiki-rightbar .markdown-body h6 {
    font-size: 1.2em;
    line-height: 1.2;
    border-color: #eee
}
#wiki-rightbar .boxed-group>h3 {
    cursor: pointer
}
#wiki-rightbar .boxed-group .caret-collapsed {
    display: none
}
#wiki-rightbar .boxed-group.collapsed .caret-expanded {
    display: none
}
#wiki-rightbar .boxed-group.collapsed .caret-collapsed {
    display: inline
}
#wiki-rightbar .boxed-group.collapsed>h3 {
    border-bottom: 1px solid #d8d8d8;
    border-radius: 3px
}
#wiki-rightbar .boxed-group.collapsed .boxed-group-inner {
    display: none
}
#wiki-rightbar p:last-child,
#wiki-rightbar ul:last-child,
#wiki-rightbar ol:last-child {
    margin-bottom: 0
}
.wiki-pages {
    padding: 0;
    margin: 0;
    list-style-type: none
}
.wiki-page-link {
    display: block;
    padding: 6px 10px;
    word-wrap: break-word
}
.has-rightbar #wiki-body,
.has-rightbar #wiki-footer {
    margin-right: 280px
}
#wiki-footer {
    margin: 20px 0 50px;
    clear: both
}
#wiki-footer .markdown-body {
    font-size: 13px
}
.wiki-wrapper .blankslate.wiki {
    padding: 115px 0
}
.wiki-wrapper .blankslate.wiki p.has-fixed-width {
    text-align: center
}
.wiki-wrapper .gh-header .divider {
    padding: 0 3px 0 2px
}
.wiki-wrapper .gh-header-meta {
    padding-bottom: 15px;
    margin-top: 6px
}
.wiki-wrapper a.history {
    color: inherit
}
.wiki-wrapper a.history:hover {
    color: #555
}
.wiki-wrapper.edit h1 {
    font-weight: normal;
    color: inherit
}
.wiki-wrapper.edit h1 strong {
    color: #000
}
.wiki-wrapper .wiki-empty-box {
    display: block;
    padding: 10px 0;
    margin: 20px 0;
    color: #767676;
    text-align: center;
    border: 1px dashed #ddd;
    border-radius: 3px;
    -webkit-transition: all 0.1s ease-in-out;
    transition: all 0.1s ease-in-out
}
.wiki-wrapper .wiki-empty-box .octicon-plus {
    margin-right: 4px;
    opacity: 0.4
}
.wiki-wrapper .wiki-empty-box:hover {
    color: #767676;
    text-decoration: none;
    border-color: #ccc
}
.wiki-wrapper .wiki-auxiliary-content {
    background: #f6f6f6;
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.06)
}
.wiki-wrapper .wiki-auxiliary-content.markdown-body.wiki-writable>*:nth-child(2) {
    margin-top: 0 !important
}
.wiki-wrapper .wiki-auxiliary-content .wiki-edit-link {
    position: relative;
    z-index: 2;
    float: right;
    color: #767676;
    opacity: 0.2;
    -webkit-transition: opacity 0.2s ease-in-out;
    transition: opacity 0.2s ease-in-out
}
.wiki-wrapper .wiki-auxiliary-content .wiki-edit-link:hover {
    text-decoration: none;
    opacity: 1
}
.wiki-wrapper .wiki-auxiliary-content-no-bg {
    background: #fff
}
.wiki-wrapper .wiki-custom-sidebar {
    padding: 10px;
    margin-bottom: 20px;
    border: solid 1px #e2e2e2;
    border-radius: 3px
}
.wiki-wrapper .wiki-custom-sidebar>:nth-child(2) {
    margin-top: 0
}
.wiki-wrapper .wiki-custom-sidebar .octicon-pencil {
    position: relative;
    z-index: 10;
    float: right;
    margin-left: 15px;
    color: #767676
}
.wiki-wrapper .wiki-custom-sidebar .octicon-pencil:hover {
    color: #333;
    text-decoration: none
}
.wiki-wrapper #wiki-footer {
    margin: 30px 30px 0;
    clear: none
}
.wiki-wrapper #wiki-footer .markdown-body {
    padding: 10px 15px
}
.wiki-wrapper #wiki-footer .wiki-empty-box {
    margin: 0 -30px
}
.wiki-wrapper #wiki-footer .wiki-edit-link {
    right: -5px
}
.wiki-wrapper #wiki-footer .wiki-auxiliary-content {
    border-top-left-radius: 3px;
    border-top-right-radius: 3px
}
.wiki-wrapper.compare .gh-header {
    margin-bottom: 20px
}
.wiki-wrapper .wiki-history {
    margin-top: 20px
}
.wiki-wrapper .wiki-history .checkbox {
    width: 30px;
    text-align: center
}
.wiki-wrapper .wiki-history .author {
    width: 200px
}
.wiki-wrapper .wiki-history .author img {
    display: block;
    float: left;
    margin-right: 6px
}
.wiki-wrapper .wiki-history .date {
    color: #bbb;
    white-space: nowrap
}
.wiki-wrapper .wiki-history .commit {
    max-width: 450px;
    overflow: hidden;
    text-overflow: ellipsis
}
.wiki-wrapper .wiki-history .commit-meta {
    width: 160px;
    padding-right: 10px;
    text-align: right;
    white-space: nowrap
}
.wiki-wrapper .wiki-history .commit-meta code {
    display: inline-block;
    font-family: Consolas, "Liberation Mono", Menlo, Courier, monospace;
    line-height: 16px;
    vertical-align: top
}
.wiki-wrapper .wiki-history .commit-id {
    color: #bbb
}
.wiki-wrapper .wiki-history .commit-id:hover {
    color: #4078c0
}
.wiki-wrapper #wiki-rightbar .sidebar-button {
    margin-top: 10px
}
.wiki-wrapper #wiki-content {
    clear: both
}
.wiki-wrapper #wiki-content .markdown-body {
    word-break: break-word
}
.wiki-wrapper #wiki-content #gollum-editor-title-field {
    margin: 0 0 14px
}
.wiki-wrapper #wiki-content .file-wrap {
    margin-top: 20px;
    border-top: 1px solid #ddd;
    border-radius: 3px
}
.wiki-wrapper #wiki-content .file-wrap .files {
    border-radius: 3px
}
.wiki-pages-box .wiki-more-pages {
    display: none
}
.wiki-pages-box.wiki-show-more .wiki-more-pages,
.wiki-pages-box .filterable-active .wiki-more-pages {
    display: block
}
.wiki-pages-box.wiki-show-more .wiki-more-pages-link,
.wiki-pages-box .filterable-active .wiki-more-pages-link {
    display: none
}
.wiki-pages-box .wiki-more-pages-link {
    box-shadow: inset 0 1px 0 #e5e5e5
}
.wiki-pages-box .wiki-more-pages-link a {
    display: block;
    padding: 3px;
    color: #7aa1d3;
    text-align: center
}
.wiki-pages-box .wiki-more-pages-link a:hover {
    color: #4078c0;
    text-decoration: none
}
.wiki-wrapper.history #gollum-footer ul.actions li {
    margin: 0 0.6em 0 0
}
.wiki-wrapper.results #results {
    border-bottom: 1px solid #ccc;
    margin-bottom: 2em;
    padding-bottom: 2em
}
.wiki-wrapper #results ul {
    margin: 2em 0 0;
    padding: 0
}
.wiki-wrapper #results li {
    font-size: 1.2em;
    line-height: 1.6em;
    list-style-position: outside;
    padding: 0.2em 0
}
.wiki-wrapper #results .count {
    color: #767676
}
.wiki-wrapper .results #no-results {
    font-size: 1.2em;
    line-height: 1.6em;
    margin-top: 2em
}
.wiki-actions {
    display: block;
    list-style-type: none;
    overflow: hidden;
    padding: 0
}
.results .wiki-actions li {
    margin: 0 1em 0 0
}
.compare .wiki-actions {
    margin-bottom: 1.4em
}
.compare .wiki-actions li {
    margin-left: 0;
    margin-right: 0.6em
}
.wiki-wrapper .file .data .line_numbers {
    width: 1%;
    font-size: 12px
}
.zeroclipboard-link {
    color: #4078c0;
    cursor: pointer;
    background: none;
    border: 0;
    padding: 0;
    margin: 0
}
.zeroclipboard-link .octicon {
    display: block
}
`)
