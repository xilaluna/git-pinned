# git-pinned

A simple go scraper that gets all of your pinned repos on GitHub via an API endpoint.

## 🧑‍💻 How to Use

The most basic route defaults to my account and gets my pinned repos. Feel free to change this in the `main.go` file.

```
api/
```

The next api route allows you to enter your username after the slash and fetch the pinned repos for that specific user.

```
api/:username
```

In both cases, the response will be a JSON object with the following structure:

```JSON
[
    {
        "title": "repo-name",
        "description": "repo-description",
        "link": "repo-link",
        "language": "repo-language",
        "stars": "repo-stars",
        "forks": "repo-forks"
    },
]
```

## 🪪 License

Licensed under the [MIT License](./LICENSE).
If you use this project, please just give a shout out and link it back to this repo. Thanks!
