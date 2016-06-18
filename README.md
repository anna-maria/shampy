# shampy
playing with go

## Setup

This application is configured to run on Heroku. As such, there are some steps
to follow to get it running in your local.

1. Install Heroku Toolbelt:

    Via Homebrew: 

    ```sh
    $ brew install heroku
    ```

    Or Via [Heroku](https://toolbelt.heroku.com/)

2. Create your .env file

    ```sh
    $ cp env.example .env
    ```

3. Update any values in `.env` as needed

4. Run the application locally

    ```sh
    $ heroku local
    ```

## Deploying

Assuming you have the proper heroku app setup + git remote

```sh
$ git push heroku [BRANCH]:master
```

