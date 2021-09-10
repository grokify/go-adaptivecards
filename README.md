# Go SDK for Authoring AdaptiveCards

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

Golang implementation of [AdaptiveCards](https://adaptivecards.io/).

 [build-status-svg]: https://github.com/grokify/go-adaptivecards/workflows/go%20build/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/go-adaptivecards/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-adaptivecards
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-adaptivecards
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-adaptivecards
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-adaptivecards
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-adaptivecards/blob/master/LICENSE

## Schema

The following have been implemented.

- [x] Cards
  - [x] AdaptiveCard
- [x] Card Elements
  - [x] [TextBlock](https://adaptivecards.io/explorer/TextBlock.html)
  - [x] Image
  - [x] Media
  - [x] MediaSource
  - [x] RichTextBlock
  - [x] TextRun
- [ ] Containers
  - [ ] ActionSet
  - [ ] Container
  - [ ] ColumnSet
  - [ ] Column
  - [ ] FactSet
  - [ ] Fact
  - [ ] ImageSet
- [ ] Actions
  - [x] [Action.OpenUrl](https://adaptivecards.io/explorer/Action.OpenUrl.html)
  - [ ] Action.Submit
  - [ ] Action.ShowCard
  - [ ] Action.ToggleVisibility
  - [ ] TargetElement
  - [ ] Action.Execute
- [x] Inputs
  - [x] Input.Text
  - [x] Input.Numbers
  - [x] Input.Date
  - [x] Input.Time
  - [x] Input.Toggle
  - [x] Input.ChoiceSet
  - [x] Input.Choice
- [x] Types
  - [x] BackgroundImage
  - [ ] Refresh
  - [ ] Authentication
  - [ ] TokenExchangeResource
  - [ ] AuthCardButton