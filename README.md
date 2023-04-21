# git-pinned

A simple go scraper that gets all of your pinned repos on GitHub via an API endpoint.

## 📑 About

git-pinned is a simple go scraper that is running on a gin server. It is designed to be used as an API endpoint to get all of your pinned repos on GitHub. It is currently being used on my [portfolio](https://xilaluna.com/projects). I initially made it for this reason, but I figured it would be useful for others as well. Feel free to fork or clone this repo. If you have any suggestions or issues, please open an issue or pull request.

## 📚 Stack

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [Colly](https://go-colly.org/)
- [Railway](https://railway.app/)

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
If you use this project, please just give the project a star. Thanks!
