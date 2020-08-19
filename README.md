# Github Profiler

Github Profiler is a command-line tool for collecting user data from Github API.

It requires an authorization token in order to access Github API.
Token can be the least privileged scope. It will only query for public accessible entities.

## Usage

`-user` should be string

`-token` should be string

```
docker run --rm yigitsadic/gogithubprofiler -user=<username> -token=<auth token>
```

## Output

JSON output is like below:

```json
{
  "userName": "Yiğit Sadıç",
  "name": "yigitsadic",
  "profilePicture": "https://avatars3.githubusercontent.com/u/727840?u=1cd118b339885ab5d46aeec8bb40d6ba31652203\u0026v=4",
  "totalPoint": 582,
  "stars": 3,
  "followers": 44,
  "repos": 15,
  "languages": null
}
```
