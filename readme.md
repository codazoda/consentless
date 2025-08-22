# Counter

Counter is a minimalist application for tracking website statistics. It aims to be GDPR compliant without requiring consent. It does so by tracking only impressions on URLs and not IP address or other user identifiable information. Time is tracked only to minute precision to further reduce the ability to pinpoint users with supplemental data.

It's designed to track impressions instead of users.

## Standard Output

All CSV data, from impressions, goes to stdout. This allows you to redirect output to a file or attach stdout to a service like pub/sub.

## Warning

If you pass PII in the query string parameter, or as any other part of a URL, then it will be logged. If someone else passes PII in a link to one of your URLs then you might log that.
