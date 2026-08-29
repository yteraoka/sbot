# sbot

[English](README.md) | 日本語

`sbot` は SwitchBot API を操作するためのコマンドラインツール (CLI) です。

## 概要

このツールを使うと、SwitchBot デバイスをコマンドラインから管理・制御できます。

## インストール

### Homebrew (macOS)

`sbot` は [yteraoka/cask](https://github.com/yteraoka/homebrew-cask) tap の cask として配布しています。

```bash
brew install --cask yteraoka/cask/sbot
```

先に tap を追加しておけば、短い名前でインストールできます。

```bash
brew tap yteraoka/cask
brew install --cask sbot
```

アップグレード・アンインストールは次の通りです。

```bash
brew upgrade --cask sbot
brew uninstall --cask sbot
```

### go install

Go がインストールされていれば、`go install` で導入できます。

```bash
go install github.com/yteraoka/sbot@latest
```

### 配布バイナリ

Linux・macOS・Windows (amd64 / arm64) 向けのアーカイブを各
[リリース](https://github.com/yteraoka/sbot/releases) に添付しています。お使いのプラットフォーム用の
ものをダウンロードして展開し、`sbot` バイナリを `PATH` の通ったディレクトリに置いてください。

### ソースからビルド

ローカル開発では GoReleaser でビルドできます。

```bash
goreleaser build --clean --snapshot
```

`dist` ディレクトリに配布用のバイナリが作成されます。

## 設定

`sbot` を使うには、SwitchBot API の認証情報を環境変数で渡す必要があります。

*   `SWITCHBOT_TOKEN`: SwitchBot API のトークン。
*   `SWITCHBOT_CLIENT_SECRET`: SwitchBot API のクライアントシークレット。

これらは SwitchBot アプリから取得できます。

## 使い方

### 汎用デバイスコマンド

#### `list`

登録されているデバイスを一覧表示します。

```bash
sbot list
```

#### `describe [DEVICE_ID]`

指定したデバイスの詳細を JSON 形式で表示します。

```bash
sbot describe <device-id>
```

#### `on [DEVICE_NAME_OR_ID]`

デバイスをオンにします。デバイスは名前または ID で指定できます。

```bash
sbot on "Bedroom Bot"
```

#### `off [DEVICE_NAME_OR_ID]`

デバイスをオフにします。デバイスは名前または ID で指定できます。

```bash
sbot off "Bedroom Bot"
```

#### `run-customize [DEVICE_NAME_OR_ID] [BUTTON_NAME]`

赤外線リモコンのカスタムボタンを実行します。

```bash
sbot run-customize "Living Room TV" "Menu"
```

### カラー電球コマンド

#### `bulb set brightness [DEVICE_NAME_OR_ID] [LEVEL]`

カラー電球の明るさを設定します。`LEVEL` は 1 から 100 の整数です。

```bash
sbot bulb set brightness "My Bulb" 75
```

#### `bulb set colortemperature [DEVICE_NAME_OR_ID] [KELVIN]`

カラー電球の色温度を設定します。`KELVIN` は 2700 から 6500 の整数です。

```bash
sbot bulb set colortemperature "My Bulb" 4000
```

### テレビコマンド

#### `tv set-channel [DEVICE_NAME_OR_ID] [CHANNEL]`

テレビのチャンネルを設定します。

```bash
sbot tv set-channel "Living Room TV" 5
```

#### `tv volume-up [DEVICE_NAME_OR_ID]`

テレビの音量を上げます。

```bash
sbot tv volume-up "Living Room TV"
```

#### `tv volume-down [DEVICE_NAME_OR_ID]`

テレビの音量を下げます。

```bash
sbot tv volume-down "Living Room TV"
```

#### `tv channel-up [DEVICE_NAME_OR_ID]`

次のチャンネルに切り替えます。

```bash
sbot tv channel-up "Living Room TV"
```

#### `tv channel-down [DEVICE_NAME_OR_ID]`

前のチャンネルに切り替えます。

```bash
sbot tv channel-down "Living Room TV"
```

### エアコンコマンド

#### `ac on [DEVICE_NAME_OR_ID] --temperature [TEMP] --mode [MODE] --fan-speed [SPEED]`

指定した設定でエアコンをオンにします。温度は必須です。

*   `--temperature`, `-t`: 温度 (摂氏)。
*   `--mode`, `-m`: `auto`, `cool`, `dry`, `fan`, `heat` (デフォルト: `auto`)。
*   `--fan-speed`, `-f`: `auto`, `low`, `medium`, `high` (デフォルト: `auto`)。

```bash
sbot ac on "Bedroom AC" -t 25 --mode cool --fan-speed medium
```

#### `ac off [DEVICE_NAME_OR_ID]`

エアコンをオフにします。

```bash
sbot ac off "Bedroom AC"
```

### シーンコマンド

#### `scene list`

登録されているシーンを一覧表示します。

```bash
sbot scene list
```

#### `scene run [SCENE_NAME_OR_ID]`

指定したシーンを実行します。

```bash
sbot scene run "Movie Time"
```

### Webhook コマンド

#### `webhook create [URL]`

Webhook を作成します。

```bash
sbot webhook create "https://example.com/webhook"
```

#### `webhook update [URL]`

Webhook を更新します。

```bash
sbot webhook update "https://new.example.com/webhook"
```

#### `webhook delete`

Webhook を削除します。

```bash
sbot webhook delete
```

## シェル補完

`sbot` は Bash と Zsh 向けの補完スクリプトを生成できます。

### `completion [bash|zsh]`

補完スクリプトを生成します。インストール方法はコマンドが出力する手順に従ってください。

**Bash の例:**

```bash
source <(sbot completion bash)
```

**Zsh の例:**

```bash
sbot completion zsh > "${fpath[1]}/_sbot"
```

## リリース

リリースは [tagpr](https://github.com/Songmu/tagpr) と [GoReleaser](https://goreleaser.com/) で自動化しています。

1. 通常どおりプルリクエストを `main` にマージします。
2. tagpr がバージョンを更新し `CHANGELOG.md` を書き換えるリリース PR を作成 (または更新) します。
   デフォルトではパッチバージョンの更新を提案します。マイナー / メジャーリリースにしたい場合は、
   その PR に `minor` または `major` ラベルを付けてください。
3. リリース PR をマージすると、tagpr が対応する `vX.Y.Z` タグを push します。
4. タグの push により GoReleaser のワークフローが起動し、バイナリのビルド、GitHub Release の作成、
   [Homebrew cask](https://github.com/yteraoka/homebrew-cask) の更新が行われます。

手動でタグを作成する必要はありません。

### 必要なリポジトリ設定

デフォルトの `GITHUB_TOKEN` で push したタグは他のワークフローを起動しないため、どちらのワークフローも
GitHub App として認証しています。

| 名前 | 種別 | 使用するワークフロー |
| --- | --- | --- |
| `TAGPR_APP_ID` | Variable | `tagpr.yml` |
| `TAGPR_APP_PRIVATE_KEY` | Secret | `tagpr.yml` |
| `HOMEBREW_APP_ID` | Variable | `release.yml` |
| `HOMEBREW_APP_PRIVATE_KEY` | Secret | `release.yml` |

tagpr 用の App にはこのリポジトリへの `Contents: write` と `Pull requests: write` が必要です。
リリース用の App にはこのリポジトリと `homebrew-cask` への `Contents: write` が必要です。

さらに Settings -> Actions で *Allow GitHub Actions to create and approve pull requests* を
有効にしておく必要があります。

## ライセンス

このプロジェクトは LICENSE ファイルの条件に従ってライセンスされています。
