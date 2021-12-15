# Validators <!-- omit in toc -->

- [Description](#description)
- [Current validators](#current-validators)
- [Roadmap](#roadmap)

## Description

This document provides more details on what our validators do, and all implemented validation.

Validators are meant to cross-validate an addon metadata with it's bundles. The goal is to perform advanced static checks and capture addon misconfigurations before they are deployed to production.

## Current validators

View `AllValidators` here: `https://github.com/mt-sre/addon-metadata-operator/blob/master/pkg/validate/validate.go`

| Name                  | Description                                                                  | Error Codes (TODO) |
| --------------------- | ---------------------------------------------------------------------------- | ------------------ |
| `001_default_channel` | Ensure defaultChannel is present in list of channels.                        |                    |
| `002_label_format`    | Ensure `label` follows the format `api.openshift.com/addon-<operator-id>`    |                    |
| `003_csv_names`       | Validate CSV.name and CSV.replaces against the operatorName and semver check |                    |
| `007_csv_names`       | Validate CSV.installModes against the metadata.installMode.                  |                    |

## Roadmap

- [ ] Add error codes in a github wiki, like `AMXXXX` similar to other linters/validators framework:
  - Hadolint (expand Pages on the right): https://github.com/hadolint/hadolint/wiki
  - Shellcheck: https://gist.github.com/nicerobot/53cee11ee0abbdc997661e65b348f375
