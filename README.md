# About

This is a project containing information about nyla's treatment.

# Technologies

- Storage:
    - MySql
- CLI Tool:
    - golang

# Schema
    - Medication table
        - id (pk)
        - name
        - type
        - time_period_hr
    - TreatmentTime table
        - id (pk)
        - recorded_time_taken (YYYY-MM-DD HH:MI:SS)
        - medicine (fk)
# Time format
    - We will use `datetime`, (i.e. YYYY-MM-DD HH:MI:SS). 
# Time Utilities
    - Info we want 
        - day (MM-DD)
        - time taken (HH:MM)
        - how long ago last taken, i.e. current_time - last_time_taken
        - next time to take, i.e. current_time + time_period (need to convert time_period to integer)
        - Day of the week (Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
    - make use of mysql function `extract()` to get a portion of datetime value
    - make use of mysql function `date_add()`
    - make use of mysql function `datediff()`

# How to insert into table with fk:
    - https://stackoverflow.com/a/15523692

# Next feature:
    - Use google calendar api to send calendar invite for reminder
        - This will be via a command
# Other operational concerns:
    - Dump the database to back up the contents of the db
