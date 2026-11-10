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

The software provider ([terraform-provider-solacebroker](https://github.com/SolaceProducts/terraform-provider-solacebroker)) and the appliance provider ([terraform-provider-solacebrokerappliance](https://github.com/SolaceProducts/terraform-provider-solacebrokerappliance)) are released together, at the same version and from the same event broker build. Do each step in both repositories.

### 1. Update the SEMP API Spec (When Moving to a New Event Broker Version)

The **Promote spec** workflow in [broker-terraform-code-generator](https://github.com/SolaceDev/broker-terraform-code-generator), an internal repository, opens the spec PR in both repositories. Don't copy specs or regenerate by hand.

1. In the generator, check that the build has a `spec/<build>` branch and that its generator gate passed.
2. Run **Promote spec** from the generator's Actions tab on `main`, with the build (for example, `10.26.9.13589`). It regenerates both providers and pushes `spec/<build>` to each. This repository's **Spec PR** workflow then opens the PR "Spec <build>" with the spec in `ci/swagger_spec/`, `internal/broker/generated/`, `docs/`, and `ci/generator-stamp.json`. A repository already on that build gets no PR. Running it again for the same build updates the open PR.
3. The PR's checks run on the `spec/<build>` push. Review the diff, and add a comment listing the attributes added, removed, and newly deprecated; they go in the release notes.
4. Merge it. Don't edit the branch: to change it, fix the generator or the spec and run **Promote spec** again.

### 2. Update `third-party-licenses.txt`

It lists the modules built into the provider: the output of `go list -deps .`, without the standard library and this module. If `go.mod` changed since the last release, update the versions, add or remove modules, and copy any new copyright lines from each module's `LICENSE` in `$(go env GOMODCACHE)`. The file is the same in both repos.

### 3. Release

- Choose the version: `X.Y.Z`, no leading `v`. A minor bump when attributes are added, otherwise a patch bump. If attributes are removed or a default changes, stop and agree the version with the team before releasing: a major bump also changes every resource's schema version. Both providers get the same version.
- Run the **Release** workflow from the Actions tab on `main`, with **dry-run** ticked first: it tests and builds without tagging or publishing. Then run it again with dry-run unticked.
- The release job waits for approval in the `release` environment. Ask another member of `dax-developers`: you can't approve your own run.
- Never push a `v*` tag by hand. It releases nothing, and the Release run for that version is then refused; `v*` tags can't be deleted without an admin turning off the `release-tags` ruleset. If a release fails after creating its tag, use **Re-run failed jobs** on that run.
- The GitHub release is published with an empty description. Add the release notes straight away, following the previous release's: the line `Release vX.Y.Z of the ... Provider. Generated from the SEMP API spec of broker version <build>.`, then **Features**, **Fixes**, **Security** and **Misc** as needed. The notes are public: no internal links or ticket keys.

### 4. Check the registry

The Terraform Registry picks up the release by itself. Run the `verify-registry-release.yml` workflow from the Actions tab with `release_version` set to `X.Y.Z` (not its default) and **public_release** ticked: it installs the release from registry.terraform.io and runs a short apply, destroy and `generate`.

## <a name="features"></a> **Do you have an ideas for a new feature or a change to an existing one?**

- Open a GitHub [enhancement request issue](https://github.com/SolaceProducts/terraform-provider-solacebroker/issues) and describe the new functionality.

##  <a name="questions"></a> Do you have questions about the source code?

- Ask any question about the code or how to use Solace messaging in the [Solace community](http://dev.solace.com/community/).
