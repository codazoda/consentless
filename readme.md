# Consentless

Track website analytics while preserving visitor privacy. No cookie banners or consent required.

See [consentless.joeldare.com](https://consentless.joeldare.com) for documentation and rationale.


## Instructions

- Host consentless on a VPS or Cloud Provider
- Add the script tag to each website you want to track


## Downloading

If you redirect standard output to a file it will produce a CSV and you can use SCP to grab it from the server.

```
scp root@143.110.236.139:/some/path/consentless.log ~/consentless.log
```


## Query with SQLite

I use SQLite to query the file directly.

```
sqlite ':memory:' -cmd ".mode csv" -cmd ".import consentless.log c"
```

Then run a query like.

```
select count(*) from c;
```
