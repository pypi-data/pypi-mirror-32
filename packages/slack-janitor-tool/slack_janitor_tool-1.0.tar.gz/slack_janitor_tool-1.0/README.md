# Slack Janitor: Your stubborn Slack cleaner

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity)

If you are relying on the Slack free tier you might have noticed that after a while you cannot upload images or files anymore.
Slack Janitor helps removing files and documents stacking up in your Slack space.

## Installation

Installation is not required as you can just clone that repository and run the following command directly:

```
python -m slack_janitor.main
```

But it can be more convenient to install the janitor on your system:

```
python setup.py install
```

## Usage

The Slack Janitor requires a Slack Token that can be retrieve [here](https://api.slack.com/custom-integrations/legacy-tokens). You'll need to pass it the janitor either by using the `-t` flag or filling the `SLACK_TOKEN` variable in your session.
To use the `slack_janitor` just call him from a terminal:

```
slack_janitor
```

the `-h` will provide you with additional parameters.



[![forthebadge made-with-python](http://ForTheBadge.com/images/badges/made-with-python.svg)](https://www.python.org/)
