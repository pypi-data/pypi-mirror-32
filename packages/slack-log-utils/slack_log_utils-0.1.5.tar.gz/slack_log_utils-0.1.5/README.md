# slack_log_utils

[![PyPI](https://img.shields.io/pypi/v/slack_log_utils.svg)](https://pypi.org/project/slack-log-utils/)
[![PyPI](https://img.shields.io/pypi/l/slack_log_utils.svg)](https://pypi.org/project/slack-log-utils/)
[![PyPI](https://img.shields.io/pypi/pyversions/slack_log_utils.svg)](https://pypi.org/project/slack-log-utils/)
[![Build Status](https://img.shields.io/travis/KeltonKarboviak/slack_log_utils.svg?logo=travis)](https://travis-ci.org/KeltonKarboviak/slack_log_utils)
[![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/KeltonKarboviak)

A Python logging handler & formatter for Slack integration.

### How To Install
To install slack-log-utils, simply:

```bash
pip install slack-log-utils
```

or using [pipenv](https://docs.pipenv.org/):

```bash
pipenv install slack-log-utils
```


### Getting Started
Get an Incoming Webhook URL on [Slack](https://my.slack.com/services/new/incoming-webhook/).

Instantiate manually:
```python
import logging
from slack_log_utils import SlackWebhookFormatter, SlackWebhookHandler

formatter = SlackWebhookFormatter()
handler = SlackWebhookHandler('https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX')
handler.setFormatter(formatter)
handler.setLevel(logging.WARNING)

logger = logging.getLogger('sample_log')
logger.addHandler(handler)
```

Instantiate using the `logging.config.dictConfig` method:
```python
import logging.config
from slack_log_utils import SlackWebhookFormatter, SlackWebhookHandler

logging_config = {
    'version': 1,
    'formatters': {
        'slack': {
            '()': SlackWebhookFormatter,
        },
    },
    'handlers': {
        'slack': {
            '()': SlackWebhookHandler,
            'formatter': 'slack',
            'level': logging.WARNING,
            'url': 'https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX',
        },
    },
    'loggers': {
        'sample_log': {
            'level': logging.DEBUG,
            'handlers': ['slack'],        
        },
    },
}

logging.config.dictConfig(logging_config)
logger = logging.getLogger('sample_log')
```

Read more about logging configuration in the [Python docs](https://docs.python.org/3/library/logging.config.html).

You can then use it just like any logger:
```python
logger.info('This will NOT post to Slack')  # Since handler's level is set to WARNING
logger.error('This will post to Slack')

logger.error('[%s]: There are %s errors!', 'CRITICAL', 8)
```

Below is the logic contained within `SlackWebhookFormatter` on how a logging.LogRecord is formatted: 
```python
import json 

formatted_msg = super(SlackWebhookFormatter, self).format(record)

color = '#36a64f' \ 
    if record.levelno < logging.WARNING \
    else '#ff0000'

attachment = {
    'author_name': record.name,
    'title': getattr(record, 'title', ''),
    'text': formatted_msg,
    'fallback': formatted_msg,
    'color': color,
}

return json.dumps({
    'attachments': [attachment]
})
```

First, it uses the `logging.Formatter`'s base `format()` method in order to get the %-style format string.
Then, the color is decided based on the record's level.  If it's less than `logging.WARNING`, it will be Green, else Red.
The `author_name` is set from `record.name` which is the name of the logger used to log the call ('sample_log' in this case).
The `title` is optionally set by using the key in the `extra` dict argument for the call to the logger:

```python
logger.error('Sample Message Attachment!', extra={'title': 'Sample Title'})
```

An example of what this will look like can be seen in the [Slack Message Builder](https://api.slack.com/docs/messages/builder?msg=%7B%22attachments%22%3A%5B%7B%22author_name%22%3A%22sample_log%22%2C%22title%22%3A%22Sample%20Title%22%2C%22text%22%3A%22Sample%20Message%20Attachment!%22%2C%22fallback%22%3A%22Sample%20Message%20Attachment!%22%2C%22color%22%3A%22%23ff0000%22%7D%5D%7D).

Please see the [Python LogRecord docs](https://docs.python.org/3/library/logging.html#logrecord-attributes) for the exact meaning of `record.levelno`, `record.name`, and other attrtibutes.

For further details on formatting messages, refer to the following Slack API pages:
* [Message Formatting](https://api.slack.com/docs/message-formatting)
* [Attachments](https://api.slack.com/docs/message-attachments)

Currently, all log messages will be sent as an attachment to Slack and it is only possible to send a single attachment at a time.
Future releases may make this more configurable.
