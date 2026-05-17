# Run Golang HyperExecute Unit Tests on TestMu AI (Formerly LambdaTest)

<p align="center">
  <a href="https://www.testmuai.com/"><img src="https://img.shields.io/badge/MADE%20BY%20TestMu%20AI-000000.svg?style=for-the-badge&labelColor=000" alt="Made by TestMu AI"></a>
  <a href="https://community.testmuai.com/"><img src="https://img.shields.io/badge/Join%20the%20community-blueviolet.svg?style=for-the-badge&labelColor=000000" alt="Community"></a>
</p>


## Getting Started

[TestMu AI](https://www.testmuai.com/) (Formerly LambdaTest) is the world's first full-stack AI Agentic Quality Engineering platform that empowers teams to test intelligently, smarter, and ship faster. Built for scale, it offers a full-stack testing cloud with 10K+ real devices and 3,000+ browsers. With AI-native test management, MCP servers, and agent-based automation, TestMu AI supports Selenium, Appium, Playwright, and all major frameworks.

With TestMu AI (Formerly LambdaTest), you can run Golang HyperExecute unit tests across real browsers and operating systems.

- [Sign up on TestMu AI](https://www.testmuai.com/register/) (Formerly LambdaTest).
- Follow the [TestMu AI Documentation](https://www.testmuai.com/support/docs/) for the full setup walkthrough.


# Pre-requisites

Before using HyperExecute, you have to download HyperExecute CLI corresponding to the host OS. Along with it, you also need to export the environment variables *LT_USERNAME* and *LT_ACCESS_KEY* that are available in your TestMu AI (Formerly LambdaTest) profile page.

## Download HyperExecute CLI

HyperExecute CLI is the CLI for interacting and running the tests on the HyperExecute Grid. The CLI provides a host of other useful Tests that accelerate test execution. In order to trigger tests using the CLI, you need to download the HyperExecute CLI binary corresponding to the platform (or OS) from where the tests are triggered:

Also, it is recommended to download the binary in the project's parent directory. Shown below is the location from where you can download the HyperExecute CLI binary:

* Mac: https://downloads.lambdatest.com/hyperexecute/darwin/hyperexecute
* Linux: https://downloads.lambdatest.com/hyperexecute/linux/hyperexecute
* Windows: https://downloads.lambdatest.com/hyperexecute/windows/hyperexecute.exe

## Configure Environment Variables

Before the tests are run, please set the environment variables LT_USERNAME & LT_ACCESS_KEY from the terminal. The account details are available on your TestMu AI (Formerly LambdaTest) profile page.

For macOS:

```bash
export LT_USERNAME=LT_USERNAME
export LT_ACCESS_KEY=LT_ACCESS_KEY
```

For Linux:

```bash
export LT_USERNAME=LT_USERNAME
export LT_ACCESS_KEY=LT_ACCESS_KEY
```

For Windows:

```bash
set LT_USERNAME=LT_USERNAME
set LT_ACCESS_KEY=LT_ACCESS_KEY
```

## Auto-Split Execution with Golang

Auto-split execution mechanism lets you run tests at predefined concurrency and distribute the tests over the available infrastructure. Concurrency can be achieved at different levels - file, module, test suite, test, scenario, etc.

For more information about auto-split execution, check out the [TestMu AI Documentation](https://www.testmuai.com/support/docs/)

### Core

Auto-split YAML file (*.hyperexecute_autosplit.yaml*) in the repo contains the following configuration:

```yaml
globalTimeout: 90
testSuiteTimeout: 90
testSuiteStep: 90
```

Global timeout, testSuite timeout, and testSuite timeout are set to 90 minutes.

The *runson* key determines the platform (or operating system) on which the tests are executed. Here we have set the target OS as Windows.

```yaml
runson: win
```

Auto-split is set to true in the YAML file.

```yaml
 autosplit: true
```

*retryOnFailure* is set to true, instructing HyperExecute to retry failed command(s). The retry operation is carried out till the number of retries mentioned in *maxRetries* are exhausted or the command execution results in a *Pass*. In addition, the concurrency (i.e. number of parallel sessions) is set to 2.

```yaml
retryOnFailure: true
runson: win
maxRetries: 2
```

## Pre Steps and Dependency Caching

To leverage the advantage offered by *Dependency Caching* in HyperExecute, the integrity of *go.exe* is checked using the checksum functionality.

```yaml
cacheKey: '{{ checksum "go.exe" }}'
```

Set the array of files & directories to be cached. In the example, all the packages will be cached in the *C:\Program Files\Go\bin* directory.


```yaml
cacheDirectories:
- C:\Program Files\Go\bin
```

The *testDiscovery* directive contains the command that gives details of the mode of execution, along with detailing the command that is used for test execution. Here, we are fetching the list of Test file scenario that would be further executed using the *value* passed in the *testRunnerCommand*

```yaml
testDiscovery:
  type: raw
  mode: dynamic
  command: |
    printf 'golang-hyperexecute-unit-test-sample-repo/Tests/armstrong/\ngolang-hyperexecute-unit-test-sample-repo/Tests/array/\ngolang-hyperexecute-unit-test-sample-repo/Tests/case1/\ngolang-hyperexecute-unit-test-sample-repo/Tests/greet/\ngolang-hyperexecute-unit-test-sample-repo/Tests/Hello/\ngolang-hyperexecute-unit-test-sample-repo/Tests/integer/\ngolang-hyperexecute-unit-test-sample-repo/Tests/multiples/\ngolang-hyperexecute-unit-test-sample-repo/Tests/oops/\ngolang-hyperexecute-unit-test-sample-repo/Tests/palindrome/\ngolang-hyperexecute-unit-test-sample-repo/Tests/repeat/\n'

```

Running the above command on the terminal will give a list of Test Scenario lines that are located in the Project folder:

Test Discovery Output:
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

The *testRunnerCommand* contains the command that is used for triggering the test. The output fetched from the *testDiscoverer* command acts as an input to the *testRunner* command.

```yaml
testRunnerCommand: go test $test -coverpkg=$test -coverprofile=coverage/profile.txt
```


### Artifacts Management

The *mergeArtifacts* directive (which is by default *false*) is set to *true* for merging the artifacts and combing artifacts generated under each task.

The *uploadArtefacts* directive informs HyperExecute to upload artifacts [files, reports, etc.] generated after task completion.  In the example, *path* consists of a regex for parsing the directory (i.e. *reports* that contains the test reports).

```yaml
mergeArtifacts: true

uploadArtefacts:
- name: Reports
  path:
  - coverage/**

```
HyperExecute also facilitates the provision to download the artifacts on your local machine. To download the artifacts, click on *Artifacts* button corresponding to the associated TestID.

### Test Execution

The CLI option *--config* is used for providing the custom HyperExecute YAML file (i.e. *.hyperexecute_autosplit.yaml*). Run the following command on the terminal to trigger the tests in go files on the HyperExecute grid. The *--download-artifacts* option is used to inform HyperExecute to download the artifacts for the job.

```bash
./hyperexecute --config --verbose .hyperexecute_autosplit.yaml
```

# Matrix Execution with Golang

Matrix-based test execution is used for running the same tests across different test (or input) combinations. The Matrix directive in HyperExecute YAML file is a *key:value* pair where value is an array of strings.

Also, the *key:value* pairs are opaque strings for HyperExecute. For more information about matrix multiplexing, check out the [TestMu AI Documentation](https://www.testmuai.com/support/docs/)

### Core

In the current example, matrix YAML file (*.hyperexecute_matrix.yaml*) in the repo contains the following configuration:

```yaml
globalTimeout: 90
testSuiteTimeout: 90
testSuiteStep: 90
```

Global timeout, testSuite timeout, and testSuite timeout are set to 90 minutes.

The target platform is set to Windows. Please set the *[runson]* key to *[win]* if the tests have to be executed on the macOS platform.

```yaml
runson: win
```

Test files in the 'Tests' folder contain the Test Scenario to run on the HyperExecute grid. In the example, the Test files run in parallel on the basis of scenario by using the specified input combinations.

```yaml
matrix:
  os: [win]
  files: ["golang-hyperexecute-unit-test-sample-repo/Tests/armstrong", "golang-hyperexecute-unit-test-sample-repo/Tests/array", "golang-hyperexecute-unit-test-sample-repo/Tests/case1", "golang-hyperexecute-unit-test-sample-repo/Tests/greet", "golang-hyperexecute-unit-test-sample-repo/Tests/Hello", "golang-hyperexecute-unit-test-sample-repo/Tests/integer", "golang-hyperexecute-unit-test-sample-repo/Tests/multiples", "golang-hyperexecute-unit-test-sample-repo/Tests/oops", "golang-hyperexecute-unit-test-sample-repo/Tests/palindrome", "golang-hyperexecute-unit-test-sample-repo/Tests/repeat"]

```

The *testSuites* object contains a list of commands (that can be presented in an array). In the current YAML file, commands for executing the tests are put in an array (with a '-' preceding each item). The go test command is used to run tests in *.go* files. The tags are mentioned as an array to the *files* key that is a part of the matrix.

```yaml
testSuites:
  - go test $files -cover -coverpkg=$files -coverprofile=coverage/profile.txt
```

### Pre Steps and Dependency Caching

Dependency caching is enabled in the YAML file to ensure that the package dependencies are not downloaded in subsequent runs. The first step is to set the Key used to cache directories.

```yaml
cacheKey: '{{ checksum "go.exe" }}'
```

Set the array of files & directories to be cached. In the example, all the packages will be cached in the *C:\Program Files\Go\bin* directory.

```yaml
cacheDirectories:
    - C:\Program Files\Go\bin
```

Steps (or commands) that must run before the test execution are listed in the *pre* run step.

```yaml
pre:
  - curl -O https://dl.google.com/go/go1.20.3.windows-amd64.msi
```

### Artifacts Management

The *mergeArtifacts* directive (which is by default *false*) is set to *true* for merging the artifacts and combing artifacts generated under each task.

The *uploadArtefacts* directive informs HyperExecute to upload artifacts [files, reports, etc.] generated after task completion. In the example, *path* consists of a regex for parsing the directory (i.e. *reports* that contains the test reports).

```yaml
mergeArtifacts: true

uploadArtefacts:
- name: Reports
  path:
  - coverage/**
```

HyperExecute also facilitates the provision to download the artifacts on your local machine. To download the artifacts, click on Artifacts button corresponding to the associated TestID.


## Test Execution

The CLI option *--config* is used for providing the custom HyperExecute YAML file (i.e. *.hyperexecute_matrix.yaml*). Run the following command on the terminal to trigger the tests in Test file Scenario on the HyperExecute grid.

```bash
./hyperexecute --config --verbose .hyperexecute_matrix.yaml
```

## Run Golang tests on Windows and Linux platforms

The CLI option *--config* is used for providing the custom HyperExecute YAML file (i.e. *.hyperexecute_autosplit.yaml* for Windows and *.hyperexecute_autosplit.yaml* for Linux).

Run the following command on the terminal to trigger tests on Windows platform:

```bash
./hyperexecute.exe --config --verbose .hyperexecute_autosplit.yaml
```

Run the following command on the terminal to trigger tests on Linux platform:

```bash
./hyperexecute --config --verbose .hyperexecute_autosplit.yaml
```

## Secrets Management

In case you want to use any secret keys in the YAML file, the same can be set by clicking on the *Secrets* button the dashboard.


All you need to do is create an environment variable that uses the secret key:

```yaml
env:
  AccessKey: ${{.secrets.AccessKey}}
```

## Navigation in Automation Dashboard

HyperExecute lets you navigate from/to *Test Logs* in Automation Dashboard from/to *HyperExecute Logs*. You also get relevant get relevant test details like video, network log, commands, Exceptions & more in the Dashboard. Effortlessly navigate from the automation dashboard to HyperExecute logs (and vice-versa) to get more details of the test execution.

## LambdaTest is Now TestMu AI

On **January 12, 2026**, [LambdaTest evolved to TestMu AI](https://www.testmuai.com/lambdatest-is-now-testmuai/), the world's first fully autonomous **Agentic AI Quality Engineering Platform**.

Same team. Same infrastructure. Same customer accounts. All existing LambdaTest logins, scripts, capabilities, and integrations continue to work without change.

ð Find the new home for [LambdaTest](https://www.testmuai.com).

### How LambdaTest Evolved into TestMu AI

In 2017, we launched LambdaTest with a simple mission: make testing fast, reliable, and accessible. As LambdaTest grew, we expanded into Test Intelligence, Visual Regression Testing, Accessibility Testing, API Testing, and Performance Testing, covering the full depth of the testing lifecycle.

As software development entered the AI era, testing had to evolve, too. We rebuilt the architecture to be AI-native from the ground up, with autonomous agents that **plan, author, execute, analyze, and optimize tests** while keeping humans in the loop. The platform integrates with your repos, CI, IDEs, and terminals, continuously learning from every code change and development signal.

That evolution earned a new name: **TestMu AI**, built for an AI-first future of quality engineering. TestMu is not a new name for us. It is the name of our annual community conference, which has brought together 100,000+ quality engineers to discuss how AI would reshape testing, long before that became an industry norm.

What started as a high-performance cloud testing platform has transformed into an AI-native, multi-agent system powering a connected, end-to-end quality layer. That evolution defined a new identity: LambdaTest evolved into TestMu AI, built for an AI-first future of quality engineering.

## Support

Got a question? Email [support@testmuai.com](mailto:support@testmuai.com) or chat with us 24x7 from our chat portal.
