# lazyjira

A Terminal UI for Jira.

The motivation for the project came from using [jira-cli](https://github.com/ankitpokhrel/jira-cli) and wishing it had
a TUI similar to these awesome projects:

- [lazygit](https://github.com/jesseduffield/lazygit)
- [lazydocker](https://github.com/jesseduffield/lazydocker)
- [k9s](https://github.com/derailed/k9s)

Using `lazyjira` you can:

- Do nothing at the moment

## Local Development

### Setup

- Install and activate [hermit](https://cashapp.github.io/hermit/)
- Install [pre-commit](https://pre-commit.com/) hooks:
  ```shell
  pre-commit install
  pre-commit install -t commit-msg
  ```
- Install [Docker](https://www.docker.com/)
- Spin up a Jira instance by running `docker compose up [-d]` and following the instructions once its up
  - This involves setting up a trial license from Atlassian
- Create a sample project in Jira, this should contain a bunch of pre-populated tickets

## Contributing

### Git

- Commits must follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) spec

> [!TIP]
> Use the pre-installed `comet` CLI to make writing commit messages easier.
