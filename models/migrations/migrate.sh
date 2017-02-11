#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mysql -u root < $DIR/reset.sql
mysql -u root do_it_dev < $DIR/01_users_actions_goals_create.sql
mysql -u root do_it_dev < $DIR/seeds.sql