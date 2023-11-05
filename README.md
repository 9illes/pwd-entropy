# Password strengh

Simple project for testing Giu.

/!\ strength is inacurrate

## Usage

Open the interface

```sh
go run . gui
```

Use the command line

```sh
go run . cli
```

## Tests

```sh
go test ./...
```

## Refs

* [Password strength](https://en.wikipedia.org/wiki/Password_strength) on wikipedia
* More info on [ANSSI](https://www.ssi.gouv.fr/administration/precautions-elementaires/calculer-la-force-dun-mot-de-passe/) website

## Libs

* GUI framework [Giu](https://pkg.go.dev/github.com/AllenDang/giu) based on Dear ImGu
* CLI library [Cobra](https://github.com/spf13/cobra)