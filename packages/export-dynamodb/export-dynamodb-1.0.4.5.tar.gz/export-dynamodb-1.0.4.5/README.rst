Export DynamoDb Cli


.. code-block:: bash

    # Install export-dynamodb
    $ pip install export-dynamodb

    # List of all export-dynamodb cli options.
    $ export-dynamodb --help

    # Export table and write to TABLE_NAME.csv file
    $ export-dynamodb -t TABLE_NAME -f csv

    # Export table and write to output.csv file
    $ export-dynamodb -t TABLE_NAME -f csv -o output.csv
