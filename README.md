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
  - [x] [AdaptiveCard](https://adaptivecards.io/explorer/AdaptiveCard.html)
- [x] Card Elements
  - [x] [TextBlock](https://adaptivecards.io/explorer/TextBlock.html)
  - [x] [Image](https://adaptivecards.io/explorer/Image.html)
  - [x] [Media](https://adaptivecards.io/explorer/Media.html)
  - [x] [MediaSource](https://adaptivecards.io/explorer/MediaSource.html)
  - [x] [RichTextBlock](https://adaptivecards.io/explorer/RichTextBlock.html)
  - [x] [TextRun](https://adaptivecards.io/explorer/TextRun.html)
- [ ] Containers
  - [ ] [ActionSet](https://adaptivecards.io/explorer/ActionSet.html)
  - [ ] [Container](https://adaptivecards.io/explorer/Container.html)
  - [ ] [ColumnSet](https://adaptivecards.io/explorer/ColumnSet.html)
  - [ ] [Column](https://adaptivecards.io/explorer/Column.html)
  - [ ] [FactSet](https://adaptivecards.io/explorer/FactSet.html)
  - [ ] [Fact](https://adaptivecards.io/explorer/Fact.html)
  - [ ] [ImageSet](https://adaptivecards.io/explorer/ImageSet.html)
  - [ ] [Table](https://adaptivecards.io/explorer/Table.html)
  - [ ] [TableCell](https://adaptivecards.io/explorer/TableCell.html)
- [ ] Actions
  - [x] [Action.OpenUrl](https://adaptivecards.io/explorer/Action.OpenUrl.html)
  - [ ] [Action.Submit](https://adaptivecards.io/explorer/Action.Submit.html)
  - [ ] [Action.ShowCard](https://adaptivecards.io/explorer/Action.ShowCard.html)
  - [ ] [Action.ToggleVisibility](https://adaptivecards.io/explorer/Action.ToggleVisibility.html)
  - [ ] [TargetElement](https://adaptivecards.io/explorer/TargetElement.html)
  - [ ] [Action.Execute](https://adaptivecards.io/explorer/Action.Execute.html)
- [x] Inputs
  - [x] [Input.Text](https://adaptivecards.io/explorer/Input.Text.html)
  - [x] [Input.Number](https://adaptivecards.io/explorer/Input.Number.html)
  - [x] [Input.Date](https://adaptivecards.io/explorer/Input.Date.html)
  - [x] [Input.Time](https://adaptivecards.io/explorer/Input.Time.html)
  - [x] [Input.Toggle](https://adaptivecards.io/explorer/Input.Toggle.html)
  - [x] [Input.ChoiceSet](https://adaptivecards.io/explorer/Input.ChoiceSet.html)
  - [x] [Input.Choice](https://adaptivecards.io/explorer/Input.Choice.html)
- [x] Types
  - [x] [BackgroundImage](https://adaptivecards.io/explorer/BackgroundImage.html)
  - [ ] [Refresh](https://adaptivecards.io/explorer/Refresh.html)
  - [ ] [Authentication](https://adaptivecards.io/explorer/Authentication.html)
  - [ ] [TokenExchangeResource](https://adaptivecards.io/explorer/TokenExchangeResource.html)
  - [ ] [AuthCardButton](https://adaptivecards.io/explorer/AuthCardButton.html)