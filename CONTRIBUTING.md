# How to contribute to a Solace Project

We'd love for you to contribute and welcome your help. Here are some guidelines to follow:

- [Issues and Bugs](#issue)
- [Submitting a fix](#submitting)
- [Feature Requests](#features)
- [Questions](#questions)

## <a name="issue"></a> Did you find a issue?

- **Ensure the bug was not already reported** by searching on GitHub under [Issues](https://github.com/SolaceProducts/terraform-provider-solacebroker/issues).

- If you're unable to find an open issue addressing the problem, [open a new one](https://github.com/SolaceProducts/terraform-provider-solacebroker/issues). Be sure to include a **title and clear description**, as much relevant information as possible, and a **code sample** or an **executable test case** demonstrating the expected behavior that is not occurring.

## <a name="submitting"></a> Did you write a patch that fixes a bug?

Open a new GitHub pull request with the patch following the steps outlined below. Ensure the PR description clearly describes the problem and solution. Include the relevant issue number if applicable.

Before you submit your pull request consider the following guidelines:

- Search [GitHub](/https://github.com/SolaceProducts/terraform-provider-solacebroker/pulls) for an open or closed Pull Request
  that relates to your submission. You don't want to duplicate effort.

### Submitting a Pull Request

Please follow these steps for all pull requests. These steps are derived from the [GitHub flow](https://help.github.com/articles/github-flow/).

#### Step 1: Fork

Fork the project and clone your fork
locally.

```sh
prompt:~$ git clone https://github.com/SolaceProducts/terraform-provider-solacebroker
```

#### Step 2: Branch

Make your changes on a new git branch in your fork of the repository.

```sh
prompt:~$ git checkout -b my-fix-branch master
```

#### Step 3: Commit

Commit your changes using a descriptive commit message.

```sh
prompt:~$ git commit -a -m "Your Commit Message"
```

Note: the optional commit `-a` command line option will automatically "add" and "rm" edited files.

#### Step 4a: Rebase (if you have not yet pushed your branch to origin, else goto step 4b.)

Assuming you have not yet pushed your branch to origin, use `git rebase` (not `git merge`) to synchronize your work with the main
repository.

If you have not set the upstream, do so as follows:

```sh
prompt:~$ git remote add upstream https://github.com/SolaceProducts/terraform-provider-solacebroker
```

then:

```sh
prompt:~$ git fetch upstream
prompt:~$ git rebase upstream/master
```

#### Step 4b: Merge (if you have already pushed our branch to origin)

Assuming you have already pushed your branch to origin, use `git merge` (not `git rebase`) to synchronize your work with the main
repository.

First ensure there are not any changes to master that you need to pick up, then merge in your changes.
You may need to resolve any conflicts on either of the merge steps.

```sh
prompt:~$ git merge master
prompt:~$ git checkout master
prompt:~$ git merge my-fix-branch
```


#### Step 5: Push

Push your branch to your fork in GitHub:

```sh
prompt:~$ git push origin my-fix-branch
```

#### Step 6: Pull Request

In GitHub, send a pull request to `solace-samples-semp:master`.

When fixing an existing issue, use the [commit message keywords](https://help.github.com/articles/closing-issues-via-commit-messages/) to close the associated GitHub issue.

- If we suggest changes then:
  - Make the required updates.
  - Commit these changes to your branch (ex: my-fix-branch)

That's it! Thank you for your contribution!

## <a name="releasing"></a> Releasing (maintainers)

The software provider ([terraform-provider-solacebroker](https://github.com/SolaceProducts/terraform-provider-solacebroker)) and the appliance provider ([terraform-provider-solacebrokerappliance](https://github.com/SolaceProducts/terraform-provider-solacebrokerappliance)) are released together, at the same version and from the same broker build. Do each step in both repos.

### 1. Update the SEMP API spec (when moving to a new broker version)

1. Check the generator first. On `main`, run `make generate-code`: it builds the latest `main` of [broker-terraform-code-generator](https://github.com/SolaceDev/broker-terraform-code-generator). Expect no change in `internal/broker/generated/`. If it changes, the generator has changed: open that as its own PR before the spec update.
2. On a branch off `main`, replace the spec in `ci/swagger_spec/` with the one from the broker version's branch (for example `10.26.8`) of the SolaceDev fork: `semp-v2-swagger-config-extended.<build>.vm.json` from [SolaceDev/terraform-provider-solacebroker](https://github.com/SolaceDev/terraform-provider-solacebroker) for the software provider, `<build>.lm.json` from [SolaceDev/terraform-provider-solacebrokerappliance](https://github.com/SolaceDev/terraform-provider-solacebrokerappliance) for the appliance. Delete the old spec: `make generate-code` needs exactly one.
3. Run `make generate-code` and `make generate-docs`. Commit the spec, `internal/broker/generated/` and `docs/`.
4. In the PR, list the attributes added, removed and newly deprecated; they go in the release notes.

### 2. Update `third-party-licenses.txt`

It lists the modules built into the provider: the output of `go list -deps .`, without the standard library and this module. If `go.mod` changed since the last release, update the versions, add or remove modules, and copy any new copyright lines from each module's `LICENSE` in `$(go env GOMODCACHE)`. The file is the same in both repos.

### 3. Release

- Choose the version: `X.Y.Z`, no leading `v`. A minor bump when attributes are added, otherwise a patch bump. Both providers get the same version.
- Run the **Release** workflow from the Actions tab on `main`, with **dry-run** ticked first: it tests and builds without tagging or publishing. Then run it again with dry-run unticked.
- The release job waits for approval in the `release` environment. Ask another member of `dax-developers`: you can't approve your own run.
- Never push a `v*` tag by hand. It releases nothing, and the Release run for that version is then refused; `v*` tags can't be deleted without an admin turning off the `release-tags` ruleset. If a release fails after creating its tag, use **Re-run failed jobs** on that run.
- The GitHub release is published with an empty description. Add the release notes straight away, following the previous release's: the line `Release vX.Y.Z of the ... Provider. Generated from the SEMP API spec of broker version <build>.`, then **Features**, **Fixes**, **Security** and **Misc** as needed. The notes are public: no internal links or ticket keys.

### 4. Check the registry

The Terraform Registry picks up the release by itself. Run the `verify-registry-release.yml` workflow from the Actions tab with the version and **public_release** ticked: it installs the release from registry.terraform.io and runs a short apply, destroy and `generate`.

## <a name="features"></a> **Do you have an ideas for a new feature or a change to an existing one?**

- Open a GitHub [enhancement request issue](https://github.com/SolaceProducts/terraform-provider-solacebroker/issues) and describe the new functionality.

##  <a name="questions"></a> Do you have questions about the source code?

- Ask any question about the code or how to use Solace messaging in the [Solace community](http://dev.solace.com/community/).
