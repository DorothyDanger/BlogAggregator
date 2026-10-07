# BlogAggregator
This is a guided learning project done on boot.dev

You will need Postgres and Go installed to run this program.

Install gator with "go install ..."

Create a .gatorconfig.json file in your home directory with the structure:

{
    "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
    "current_user_name": "username"
}

Change your directory to /sql/schema and run the command:
goose postgres <connection_string> up
With the connection string being that used in the config file.

Gator should be up and running after this.

The commands available to you are in the form "gator <command>":

    # login <username>
        Login to your account if it exists in the database

    # register <name>
        Register's a new user

    # reset
        This will reset the users table

    # users
        This will print a list of users to the console

    # agg <time>
        This will start aggregating posts from the oldest feeds

    # addfeed <name> <url>
        Connects a feed to the current user given a provided feed name and url

    # feeds
        Lists all the feeds in the database to the console

    # follow <url>
        Takes a url and creats a new record of the user following the feed

    # following
        Lists all of the feeds you are following

    # unfollow
        Unfollows a feed

    # browse <limit>
        Outputs the posts you are following up to the limit with a default of 2 if none is specified.