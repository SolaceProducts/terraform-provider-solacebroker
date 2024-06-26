# Terraform Provider for Solace PubSub+ Software Event Broker
[![Actions Status](https://github.com/SolaceProducts/terraform-provider-solacebroker/actions/workflows/core-pipeline-main-branch-only.yml/badge.svg?branch=main)](https://github.com/SolaceProducts/terraform-provider-solacebroker/actions?query=workflow:%22Test+Provider%22+branch:main)
[![Go Report Card](https://goreportcard.com/badge/github.com/solaceproducts/terraform-provider-solacebroker)](https://goreportcard.com/report/github.com/solaceproducts/terraform-provider-solacebroker)


This provider, maintained by Solace, is a plugin for Terraform that enables you to configure PubSub+ Software Event Brokers.

The provider is available from the [Terraform Providers Registry](https://registry.terraform.io/providers/solaceproducts/solacebroker/latest).

The minimum required PubSub+ Software Event Broker version is 10.4.

## Limitations

> This version of the PubSub+ Terraform provider is generally available for production services contained within a Message VPN.  Use of any resource not contained within a Message VPN is not supported in production by Solace.

## Quick Start

1. Ensure you have admin access to a Solace PubSub+ Software Event Broker. Options include [local deployment of a containerized version](https://docs.solace.com/Software-Broker/SW-Broker-Set-Up/Containers/Set-Up-Container-Image.htm) or use of a free broker from [PubSub+ Cloud](https://docs.solace.com/Cloud/cloud-lp.htm).
2. Install the [Terraform CLI](https://www.terraform.io/downloads)
3. Create the [`examples/sampleconfig.tf`](examples/sampleconfig.tf) sample file in a new directory, adjust the `url`, and the management credential parameters `username` and `password` to your broker's setup.
4. From this directory run `terraform plan`, then `terraform apply` (if prompted, `terraform init` may also be required).
5. Open the `url` link in your browser to access the broker's web management UI. The credentials are the same as used in the Terraform config. Observe the new objects created: a new Message VPN and a messaging queue under that Message VPN.
6. Run `terraform destroy` from your command line and observe that the Message VPN is deleted.
   
Note that the provider also offers the unique ability to generate a config file from an already configured broker. For more information, refer to the [full documentation](https://registry.terraform.io/providers/SolaceProducts/solacebroker/latest/docs).

## Documentation

Full documentation is available on the [Terraform Providers Registry website](https://registry.terraform.io/providers/solaceproducts/solacebroker/latest/docs).

We recommended that you familiarize yourself with Solace technology and broker management. For more information, refer to the [Resources section](#resources).

## Development

### Requirements

* [Terraform](https://www.terraform.io/downloads) (>= 1.0)
* [Go](https://go.dev/doc/install) (1.22)
* [Make](https://www.gnu.org/software/make/)

### Building

1. `git clone` this repository and `cd` into its directory.
2. `make install` will trigger the Golang build of the provider in your [`${GOBIN}`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) (defaults to `${GOPATH}/bin` or `${HOME}/go/bin` if `${GOPATH}` is not set). Repeat
this every time you make changes to the provider locally.

The provided `makefile` defines additional commands that are useful during development.

### Using a Development Build

Create or update your `${HOME}/.terraformrc` (Unix) / `%APPDATA%\terraform.rc` (Windows) configuration with a `provider_installation` section that contains the following `dev_overrides`:

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/solaceproducts/solacebroker" = "${GOBIN}" //< replace `${GOBIN}` with the actual path on your system
  }

  direct {}
}
```

Note that it is also possible to use a dedicated Terraform configuration file and invoke `terraform` while setting
the environment variable `TF_CLI_CONFIG_FILE=my_terraform_config_file`.

Once the `dev_overrides` are in place, any local execution of `terraform plan` and `terraform apply` will
use the version of the provider found in the given `${GOBIN}` directory, instead of the one indicated in your terraform configuration.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the Apache License, Version 2.0. - See the [LICENSE](LICENSE) file for details.

## Resources

For more information about Solace technology in general please visit these resources:

- Understanding [Solace technology](https://docs.solace.com/Get-Started/Solace-PubSub-Platform.htm)
- The Solace Developer Portal website at: [solace.dev](https://solace.dev/)
- Ask the [Solace community](https://solace.community/).
