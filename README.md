# tf-azurerm-module_primitive-monitor_private_link_scoped_service

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This module is used to associate an Azure Monitor resource with a Private Link Scope

## Usage

This repository has an example configuration in:

- [examples/with_log_analytics](examples/with_log_analytics)

## Module Development

Use this repository as a standard Launch Terraform primitive module.

- Keep examples and tests aligned with code changes because they are part of the public contract.
- Preserve generated files and automation patterns from the shared skeleton unless a module-specific exception is required.
- Prefer make targets and pre-commit hooks over ad hoc commands to match CI behavior.

## Pre-Requisites

The following commands should be available on your system:

- asdf or mise
- make
- python3 (for pre-commit)

Install pinned tool versions and bootstrap dependencies from the repository root:

```sh
make configure
```

## Pre-Commit Hooks

This repository uses [.pre-commit-config.yaml](.pre-commit-config.yaml) to run Terraform, Go, and repository hygiene checks.

Install local hooks:

```sh
pre-commit install --hook-type commit-msg
```

Run all hooks manually:

```sh
pre-commit run --all-files
```

## Local Validation

Run the same validations used in CI:

```sh
make lint
make check
```

If a hook or generated file changes content (for example terraform-docs), commit the updates and rerun the checks.

## Review And Merge Process

- Open a pull request with a clear summary of functional and test-impacting changes.
- Resolve all review comments and ensure CI is green before merge.
- Keep commits focused and use conventional commit messages when possible.

## Automatic Updates

This repository receives periodic updates from the shared launch-terraform-skeleton baseline via Copier automation. Keep skeleton-managed files aligned with upstream expectations so automated updates continue to merge cleanly.
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.5 |
| <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) | ~> 3.67 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [azurerm_monitor_private_link_scoped_service.service](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/monitor_private_link_scoped_service) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_monitor_private_link_scope_id"></a> [monitor\_private\_link\_scope\_id](#input\_monitor\_private\_link\_scope\_id) | Resource group name | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Name of the private link scoped service | `string` | n/a | yes |
| <a name="input_resource_id"></a> [resource\_id](#input\_resource\_id) | ID of the resource to associate with the Private Link Scope | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_id"></a> [id](#output\_id) | resource ID of the scoped service |
<!-- END_TF_DOCS -->
