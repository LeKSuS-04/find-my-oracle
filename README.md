# Find my oracle

<hr>

## ⚠️ LIABILITY DISCLAMER ⚠️

Provided software is intended for educational purposes ONLY and cannot be used in
any kind of illegal, malicious or destructive activity. For more details check out
LICENSE file.

<hr>

## Information
This is simple cloud scanner for oracle cloud. It's intened for searching maching
in public ips, when no other methods for identifing it's ip address are available.
You need to edit `checker` function in checker/checker.go file to include custom
logic for identifing your instance.

### Build from source
```bash
git clone https://github.com/LeKSuS-04/find-my-oracle/
cd find-my-oracle
# ...
# Edit checker/checker.go file
# ...
go build
```

### Usage
```
Usage of find-my-oracle:
  -no-cache
        don't use local sqlite database to cache requests
  -region string
        region to search server in
  -threads int
        amount of threads (default 20)
  -timeout int
        timeout in milliseconds to use in checker function (default 10000)
```
