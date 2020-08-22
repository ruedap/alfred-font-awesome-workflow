# Font Awesome Workflow for Alfred

[![Latest Version](http://img.shields.io/github/release/ruedap/alfred-font-awesome-workflow.svg?style=flat-square)](https://github.com/ruedap/alfred-font-awesome-workflow/releases)
[![License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://ruedap.mit-license.org/2015)

You can incrementally search for [Font Awesome icons](https://fontawesome.com/changelog/latest) and paste it to front most app.

#### :pencil: for coding HTML/CSS

![screencast.gif](https://github.com/ruedap/alfred-font-awesome-workflow/raw/master/screenshots/screencast.gif)

#### :art: for designing in Photoshop, Illustrator, Sketch, Keynote, etc

![screencast-illustrator.gif](https://github.com/ruedap/alfred-font-awesome-workflow/raw/master/screenshots/screencast-illustrator.gif)

:gem: See also: [Font Awesome Workflow with Sketch 3 — QuickCast.](http://quick.as/46rbfrqr)

## Installation

Download **[Font-Awesome.alfredworkflow](https://github.com/ruedap/alfred-font-awesome-workflow/releases)** and import to [Alfred](http://www.alfredapp.com/) (requires Powerpack).

The current supported version is **Font Awesome 5.14.0 Free**. Previous versions are available [here](https://github.com/ruedap/alfred-font-awesome-workflow/releases).

## Usage

**Keyword** `fa`: Search through [Font Awesome icons](https://fontawesome.com/icons?d=gallery&m=free).

- `Enter`: Paste class name (**for coding HTML/CSS**. e.g. `fa-arrow-circle-o-right`)
- `Ctrl + Enter`: Paste character name and ligatures shoud convert the letter into the correct symbol. (**for designing in Photoshop, Illustrator, etc**)
- `Shift + Enter`: Paste character code (e.g. `f18e`)
- `Command + Enter`: Open in browser (e.g. <http://fontawesome.io/icon/arrow-circle-o-right>)

![Workflow](https://github.com/ruedap/alfred-font-awesome-workflow/raw/master/screenshots/workflow.png)

## Options

**Disable pasting**: Turn off "Automatically paste to front most app" in Workflow's Preferences.

![Disable pasting](https://github.com/ruedap/alfred-font-awesome-workflow/raw/master/screenshots/option-disable-pasting.png)

## Changelog

See details changes for each version in the [release notes](https://github.com/ruedap/alfred-font-awesome-workflow/releases).

## Links

- [Syntax.fm Podcast Episode 11 - Our favourite Productivity Hacks](https://syntax.fm/show/011/our-favourite-productivity-hacks)
- [Automating Front-end Workflow (Addy Osmani)](https://speakerdeck.com/addyosmani/automating-front-end-workflow?slide=151)
- [Hidden Productivity Secrets With Alfred | Smashing Coding](http://coding.smashingmagazine.com/2013/10/25/hidden-productivity-secrets-with-alfred/)
- [Can I import FontAwesome icons in Sketch 3? - Quora](http://www.quora.com/Can-I-import-FontAwesome-icons-in-Sketch-3)
  - [Quickly search and insert icons to Sketch using Alfed - YouTube](https://www.youtube.com/watch?v=nEFW_NmC-TA)
- [Font Awesome Workflow with Sketch 3 — QuickCast.](http://quick.as/46rbfrqr)
- [Sketch 3 - Insert Font Awesone icons w/ Alfred workflow — QuickCast.](http://quick.as/dvxup47)

## Development

```sh
$ npm i
$ npm t
$ npm run build
$ npm run open
```

## License

Released under the [MIT license](http://ruedap.mit-license.org/2015).

## Author

<a href="https://github.com/ruedap"><img src="https://avatars.githubusercontent.com/u/289671?v=3&s=300" alt="ruedap" title="ruedap" width="100" height="100"></a>
