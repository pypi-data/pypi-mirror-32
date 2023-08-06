# testcube-client

[![TestCube PyPI](https://img.shields.io/pypi/v/testcube-client.svg)](https://pypi.org/project/testcube-client/)
[![Build](https://img.shields.io/travis/tobyqin/testcube-client.svg)](https://travis-ci.org/tobyqin/testcube-client)
[![Updates](https://pyup.io/repos/github/tobyqin/testcube-client/shield.svg)](https://pyup.io/repos/github/tobyqin/testcube-client/)

A Python client for testcube.

- Free software: MIT license
- TestCube Project: <https://github.com/tobyqin/testcube>
- TestCube Client Project: <https://github.com/tobyqin/testcube-client>

## Get Started

You should have python 2.7 or 3.x installed on your machine. Then follow steps here.

### Installation

Install testcube-client via pip is the most easy way.

```shell
pip install testcube-client -U
```

### Register to Server

You must have a [TestCube](https://github.com/tobyqin/testcube) server deployed at somewhere, then run `--register` command.:

```shell
testcube-client --register http://testcube-server:8000
```

### Upload Run Results

Once you registered to a server, you will have 2 ways to upload test results, first choice is calling `--start-run` at the beginning and calling `--finish--run` when finished.:

```shell
# call --start-run before run started
testcube-client --start-run -name "nightly run for testcube"  --team Core --product TestCube

# call --finish-run after run finished
 testcube-client --finish-run --xunit-files "**/results/*.xml"
```

In this way, TestCube will record the **exact** `start_time` and `end_time` for the run. You should provide `team` and `product` name when start the run.

Another choice is using `--run` command to upload test results like a batch.:

```shell
# put this command at the end of a run
testcube-client --run -n "smoke tests for testcube" -t XPower -p TestCube -v v1.0 -x "**/smoke*.xml"
```

With this choice, TestCube will use current time as `end_time` for the run, and guess `start_time`according to run duration. `team` and `product` name are also required.

### Upload Result Files

Sometimes your test will generate result files, for example, screenshots, log files. These files are useful when analyze test failure, so we can upload them to TestCube as well:

```shell
# call --result-files to upload files
testcube-client --result-files "**/*.png"

# use -i as a shortcut
testcube-client -i "**/*.jpg"
```

**Note:**

In order to link a file to a test result, the `file name` must **contains** `test case name` so that TestCube can map them together.

For example:

```shell
test case name => "test_student_login"
result file name => "test_student_login_failed.png"
```

### Command-line Options

The optional arguments:

```shell
-h, --help            show this help message and exit
-r REGISTER, --register REGISTER
                      Register to the TestCube server, e.g.
                      http://server:8000
-run, --run           Upload run info at one time, require team,product,name
                      and xunit files.
-start, --start-run   Start a run, require team, product and a name.
-finish, --finish-run
                      Finish a run, require xunit files.
-x XUNIT_FILES, --xunit-files XUNIT_FILES
                      Specify the xunit xml results, e.g "**/result*.xml"
-i RESULT_FILES, --result-files RESULT_FILES
                      Specify the result files, e.g "**/output/**/*.png"
-n NAME, --name NAME  Specify the run name.
-t TEAM, --team TEAM  Specify the team name.
-p PRODUCT, --product PRODUCT
                      Specify the product name.
-v PRODUCT_VERSION, --product-version PRODUCT_VERSION
                      Specify the product version. [Optional]
-f, --force           Force the action, support --register command.
-vb, --verbose        Show verbose log info.
```