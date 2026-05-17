# Run Go Unit Tests on HyperExecute with TestMu AI (Formerly LambdaTest)

<p align="center">
  <a href="https://www.testmuai.com/"><img src="https://img.shields.io/badge/MADE%20BY%20TestMu%20AI-000000.svg?style=for-the-badge&labelColor=000" alt="Made by TestMu AI"></a>
  <a href="https://pkg.go.dev/testing"><img src="https://img.shields.io/github/go-mod/go-version/LambdaTest/golang-hyperexecute-unit-test-sample-repo.svg?style=for-the-badge&labelColor=000000" alt="Go version"></a>
  <a href="https://community.testmuai.com/"><img src="https://img.shields.io/badge/Join%20the%20community-blueviolet.svg?style=for-the-badge&labelColor=000000" alt="Community"></a>
</p>

## Getting Started

[TestMu AI](https://www.testmuai.com/) (Formerly LambdaTest) is the world's first full-stack AI Agentic Quality Engineering platform that empowers teams to test intelligently, smarter, and ship faster. Built for scale, it offers a full-stack testing cloud with 10K+ real devices and 3,000+ browsers. With AI-native test management, MCP servers, and agent-based automation, TestMu AI supports Selenium, Appium, Playwright, and all major frameworks.

With TestMu AI (Formerly LambdaTest), you can run Go unit tests on HyperExecute across different platforms with auto-split and matrix execution modes.

- [Sign up on TestMu AI](https://www.testmuai.com/register/) (Formerly LambdaTest).
- Follow the [TestMu AI documentation](https://www.testmuai.com/support/docs/) (Formerly LambdaTest) for the full setup walkthrough.

### Prerequisites

Before using HyperExecute, download the HyperExecute CLI binary corresponding to your host OS. Along with it, export the environment variables `LT_USERNAME` and `LT_ACCESS_KEY` available from your TestMu AI (Formerly LambdaTest) profile page.

**Download HyperExecute CLI**

Download the HyperExecute CLI binary for your platform (recommended: place it in the project's parent directory):

* Mac: https://downloads.lambdatest.com/hyperexecute/darwin/hyperexecute
* Linux: https://downloads.lambdatest.com/hyperexecute/linux/hyperexecute
* Windows: https://downloads.lambdatest.com/hyperexecute/windows/hyperexecute.exe

### Setup

Set your credentials as environment variables.

**macOS / Linux:**

```bash
export LT_USERNAME="YOUR_USERNAME"
export LT_ACCESS_KEY="YOUR_ACCESS_KEY"
```

**Windows:**

```bash
set LT_USERNAME="YOUR_USERNAME"
set LT_ACCESS_KEY="YOUR_ACCESS_KEY"
```

## Auto-Split Execution with Go

Auto-split execution lets you run tests at predefined concurrency and distribute tests over the available infrastructure. Concurrency can be achieved at different levels — file, module, test suite, test, scenario, etc.

For more information about auto-split execution, check out the [TestMu AI documentation](https://www.testmuai.com/support/docs/).

### Core

The auto-split YAML file (`.hyperexecute_autosplit.yaml`) in the repo contains the following configuration:

```yaml
globalTimeout: 90
testSuiteTimeout: 90
testSuiteStep: 90
```

Global timeout, testSuite timeout, and testSuite step timeout are set to 90 minutes.

The `runson` key determines the platform on which the tests are executed. Here it is set to Windows.

```yaml
runson: win
```

Auto-split is set to true in the YAML file.

```yaml
 autosplit: true
```

`retryOnFailure` is set to true, instructing HyperExecute to retry failed commands. The retry operation runs until the number of retries in `maxRetries` are exhausted or the command results in a pass. Concurrency (number of parallel sessions) is set to 2.

```yaml
retryOnFailure: true
runson: win
maxRetries: 2
```

### Pre Steps and Dependency Caching

To leverage dependency caching in HyperExecute, the integrity of `go.exe` is checked using the checksum functionality.

```yaml
cacheKey: '{{ checksum "go.exe" }}'
```

Set the array of files and directories to be cached. In the example, all packages are cached in the `C:\Program Files\Go\bin` directory.

```yaml
cacheDirectories:
- C:\Program Files\Go\bin
```

The `testDiscovery` directive contains the command that gives details of the mode of execution and the command used for test execution. Here, the list of test file scenarios is fetched and passed to `testRunnerCommand`.

```yaml
testDiscovery:
  type: raw
  mode: dynamic
  command: |
    printf 'golang-hyperexecute-unit-test-sample-repo/Tests/armstrong/\ngolang-hyperexecute-unit-test-sample-repo/Tests/array/\ngolang-hyperexecute-unit-test-sample-repo/Tests/case1/\ngolang-hyperexecute-unit-test-sample-repo/Tests/greet/\ngolang-hyperexecute-unit-test-sample-repo/Tests/Hello/\ngolang-hyperexecute-unit-test-sample-repo/Tests/integer/\ngolang-hyperexecute-unit-test-sample-repo/Tests/multiples/\ngolang-hyperexecute-unit-test-sample-repo/Tests/oops/\ngolang-hyperexecute-unit-test-sample-repo/Tests/palindrome/\ngolang-hyperexecute-unit-test-sample-repo/Tests/repeat/\n'
```

Running the above command on the terminal gives a list of test scenario lines located in the project folder:

```
golang-hyperexecute-unit-test-sample-repo/Tests/armstrong/
golang-hyperexecute-unit-test-sample-repo/Tests/array/
golang-hyperexecute-unit-test-sample-repo/Tests/case1/
golang-hyperexecute-unit-test-sample-repo/Tests/greet/
golang-hyperexecute-unit-test-sample-repo/Tests/Hello/
golang-hyperexecute-unit-test-sample-repo/Tests/integer/
golang-hyperexecute-unit-test-sample-repo/Tests/multiples/
golang-hyperexecute-unit-test-sample-repo/Tests/oops/
golang-hyperexecute-unit-test-sample-repo/Tests/palindrome/
golang-hyperexecute-unit-test-sample-repo/Tests/repeat/
```

The `testRunnerCommand` contains the command used for triggering the test. The output from `testDiscovery` acts as input to the test runner.

```yaml
testRunnerCommand: go test $test -coverpkg=$test -coverprofile=coverage/profile.txt
```

### Artifacts Management

The `mergeArtifacts` directive (default: `false`) is set to `true` for merging artifacts generated under each task.

The `uploadArtefacts` directive instructs HyperExecute to upload artifacts (files, reports, etc.) generated after task completion. In the example, `path` contains a regex for parsing the directory that contains test reports.

```yaml
mergeArtifacts: true

uploadArtefacts:
- name: Reports
  path:
  - coverage/**
```

HyperExecute also lets you download the artifacts to your local machine. Click on the **Artifacts** button corresponding to the associated TestID.

### Run tests

```bash
./hyperexecute --config --verbose .hyperexecute_autosplit.yaml
```

View results on your TestMu AI dashboard.

## Matrix Execution with Go

Matrix-based test execution runs the same tests across different test or input combinations. The Matrix directive in HyperExecute YAML file is a `key:value` pair where value is an array of strings.

For more information about matrix multiplexing, check out the [TestMu AI documentation](https://www.testmuai.com/support/docs/).

### Core

The matrix YAML file (`.hyperexecute_matrix.yaml`) in the repo contains the following configuration:

```yaml
globalTimeout: 90
testSuiteTimeout: 90
testSuiteStep: 90
```

The target platform is set to Windows. Set the `runson` key to `win` if the tests have to be executed on the Windows platform.

```yaml
runson: win
```

Test files in the `Tests` folder contain the test scenarios to run on the HyperExecute grid. Tests run in parallel based on the specified input combinations.

```yaml
matrix:
  os: [win]
  files: ["golang-hyperexecute-unit-test-sample-repo/Tests/armstrong", "golang-hyperexecute-unit-test-sample-repo/Tests/array", "golang-hyperexecute-unit-test-sample-repo/Tests/case1", "golang-hyperexecute-unit-test-sample-repo/Tests/greet", "golang-hyperexecute-unit-test-sample-repo/Tests/Hello", "golang-hyperexecute-unit-test-sample-repo/Tests/integer", "golang-hyperexecute-unit-test-sample-repo/Tests/multiples", "golang-hyperexecute-unit-test-sample-repo/Tests/oops", "golang-hyperexecute-unit-test-sample-repo/Tests/palindrome", "golang-hyperexecute-unit-test-sample-repo/Tests/repeat"]
```

The `testSuites` object contains a list of commands. The `go test` command is used to run tests in `.go` files. Tags are mentioned as an array to the `files` key that is a part of the matrix.

```yaml
testSuites:
  - go test $files -cover -coverpkg=$files -coverprofile=coverage/profile.txt
```

### Pre Steps and Dependency Caching

Dependency caching is enabled in the YAML file to ensure package dependencies are not re-downloaded in subsequent runs.

```yaml
cacheKey: '{{ checksum "go.exe" }}'
```

Set the array of files and directories to be cached.

```yaml
cacheDirectories:
    - C:\Program Files\Go\bin
```

Steps that must run before test execution are listed in the `pre` run step.

```yaml
pre:
  - curl -O https://dl.google.com/go/go1.20.3.windows-amd64.msi
```

### Artifacts Management

```yaml
mergeArtifacts: true

uploadArtefacts:
- name: Reports
  path:
  - coverage/**
```

### Run tests

```bash
./hyperexecute --config --verbose .hyperexecute_matrix.yaml
```

## Run Go Tests on Windows and Linux

Run the following command on the terminal to trigger tests on Windows:

```bash
./hyperexecute.exe --config --verbose .hyperexecute_autosplit.yaml
```

Run the following command on the terminal to trigger tests on Linux:

```bash
./hyperexecute --config --verbose .hyperexecute_autosplit.yaml
```

## Secrets Management

To use secret keys in the YAML file, set them by clicking on the **Secrets** button on the dashboard. Create an environment variable that uses the secret key:

```yaml
env:
  AccessKey: ${{.secrets.AccessKey}}
```

## Navigation in Automation Dashboard

HyperExecute lets you navigate between **Test Logs** in the Automation Dashboard and **HyperExecute Logs**. You get relevant test details like video, network log, commands, exceptions, and more in the dashboard.

## Contributions

Contributions are welcome. Open an issue to discuss your idea before submitting a pull request. When reporting bugs, include your Go version, OS, and HyperExecute CLI version.

## TestMu AI (Formerly LambdaTest) Community

Connect with testers and developers in the [TestMu AI Community](https://community.testmuai.com/). Ask questions, share what you are building, and discuss best practices in test automation and DevOps.

## TestMu AI (Formerly LambdaTest) Certifications

Earn free [TestMu AI Certifications](https://www.testmuai.com/certifications/) for testers, developers, and QA engineers. Validate your skills in Selenium, Cypress, Playwright, Appium, Espresso and more. Industry-recognized, shareable on LinkedIn, and built by practitioners, not marketers.

## Learning Resources by TestMu AI (Formerly LambdaTest)

Learn modern testing through tutorials, guides, videos, and weekly updates:

* [TestMu AI Blog](https://www.testmuai.com/blog/)
* [TestMu AI Learning Hub](https://www.testmuai.com/learning-hub/)
* [TestMu AI on YouTube](https://www.youtube.com/@TestMuAI)
* [TestMu AI Newsletter](https://www.testmuai.com/newsletter/)

## LambdaTest is Now TestMu AI

On **January 12, 2026**, [LambdaTest evolved to TestMu AI](https://www.testmuai.com/lambdatest-is-now-testmuai/), the world's first fully autonomous **Agentic AI Quality Engineering Platform**.

Same team. Same infrastructure. Same customer accounts. All existing LambdaTest logins, scripts, capabilities, and integrations continue to work without change.

👉 Find the new home for [LambdaTest](https://www.testmuai.com).

### How LambdaTest Evolved into TestMu AI

In 2017, we launched LambdaTest with a simple mission: make testing fast, reliable, and accessible. As LambdaTest grew, we expanded into Test Intelligence, Visual Regression Testing, Accessibility Testing, API Testing, and Performance Testing, covering the full depth of the testing lifecycle.

As software development entered the AI era, testing had to evolve, too. We rebuilt the architecture to be AI-native from the ground up, with autonomous agents that **plan, author, execute, analyze, and optimize tests** while keeping humans in the loop. The platform integrates with your repos, CI, IDEs, and terminals, continuously learning from every code change and development signal.

That evolution earned a new name: **TestMu AI**, built for an AI-first future of quality engineering. TestMu is not a new name for us. It is the name of our annual community conference, which has brought together 100,000+ quality engineers to discuss how AI would reshape testing, long before that became an industry norm.

What started as a high-performance cloud testing platform has transformed into an AI-native, multi-agent system powering a connected, end-to-end quality layer. That evolution defined a new identity: LambdaTest evolved into TestMu AI, built for an AI-first future of quality engineering.

## Support

Got a question? Email [support@testmuai.com](mailto:support@testmuai.com) or chat with us 24x7 from our chat portal.
