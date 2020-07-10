# Notes

## Ideas

- [ ] use OAuth2 instead of API key
- [x] add swagger docs
  - [ ] build automatically using something like `@go generate`
- [x] embed gql files as resources using go-bindata
- [ ] database required for automation?
  - [ ] put in sqlite3 database?
  - [ ] gorm
  - [ ] alternatively use an in-memory [KV database](https://github.com/etcd-io/bbolt)
- [ ] import types
  - [ ] yaml (preferred)
  - [ ] xlsx (maybe) use [this](https://github.com/tealeg/xlsx)
- [ ] fix goreleaser build
  - [ ] prevent go mod from changing on tags
  - [ ] currently can only push to brew repo using `GITHUB_TOKEN=... gorelease --rm-dist`

## Populating Data

- [x] Locations
- [x] Departments
- [ ] Users
- [ ] Objectives
- [ ] One-on-one templates
- [ ] Recognition badges
- [ ] Competencies
  - [ ] Competency levels
- [ ] Role templates
  - [ ] Responsibilities

## Automating Data
