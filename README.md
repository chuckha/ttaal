# Two Truths and a Lie

## Internal ORM

The interesting part of this codebase is the internal ORM.

## CLI

### Migrator

The migrator will look at structs and generate queries for the configured backend to migrate (create, update) database tables.

## Models

Note: this section is likely out of date

Poll has many options
Options have many votes
Vote has one user
Vote has one Option

users
| id | created |
* Keep track of people who have signed in

polls
| id | user_id | ended | preamble | created | updated |
* id is required as it is a reference
* user_id indicates owner
* created so we know when it was created (for auto-expire after 30 days)
* updated so we can display if the introduction was edited or not
* preamble to be used to make an introduction

statements
| id | poll_id | is a lie | statement | created | updated |
* id is required as it is a reference
* poll_id is required to be the referencing of the owning poll
* is a lie tells us which statement is a lie
* statement is the text of
* created and updated together let us know if it has been edited

votes
| id | statement_id | user_id | created | deleted |
* id is a vote id, referenced by comments
* statement_id links a vote for a statement to be the lie
* user_id links the owner of the vote
* created tells us when the vote was cast, updates if statement_id updates
* deleted allows us to mark this vote as deleted so a user can change votes

comments
| id | poll_id | user_id | comment | created | updated | deleted |
* id is used as a primary key
* poll_id links a comment and a poll
* user_id tells us who voted
* comment is the actual comment
* created and updated allow us to show if a comment was edited
* deleted allows us to mark the comment as deleted

app enforces constraints
* Polls must have exactly 3 options
* One user can only own one poll (maybe)
* One statement must be a lie
* One vote per user per poll
* Cannot change votes on a poll that has ended
* A statement MUST be associated with a poll.


## Operation

What if a model must change?
Need a way to modify tables. Could a cli that prints the migration based on the input table?

`./ttaal migrate <model> --config config.yaml`

And then it connects to the db, looks to see if the new model (model ./ttaal was compiled with) is the same as the existing model.
If it's not it will print the sql commands necessary to migrate

## Bonus features

## Leaderboard
 - Shows an ordered list of users that have the most number of correct guesses

## slack integration
 - ping a slack channel when a new poll is created

