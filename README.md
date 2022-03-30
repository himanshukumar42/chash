# chash
chash is basically a command line tool built with golang by using cobra framework

Some of the Commands are 
  1. chash - it will basically create a sqlite database in a local repository.
  2. random - This is a chash command that will print random jokes in the terminal.
  3. init - This command will basically creates a new chash table in the sqlite3 database.
  4. note - This chash command will have the notes functionallity. This will do the following things.
          1. chash note new - 
                This command will create a new note.
                It will ask a prompt for entering the notes title and description
                It will prompt to ask to select the category.
          2. chash note list -
                This command will list all the notes present in the database.
