# kvdb
Simple In-memory KV Database implementation(SET, GET, BEGIN, END, COMMIT, ROLLBACK)

 Commands to Use:
 SET KEY VALUE // Sets Key to value. If part of Txn, till COMMIT
               // is called, its in mem. If ROLLBACK, then deleted.
 GET KEY       // Gets value first from Pending Txns and then Global DB.
 BEGIN         // Starts a transaction of multi keys.
 ROLLBACK      // Takes back all members of the transacted items.
 COMMIT        // Adds key-value pairs from transactions to DB.
 END           // Dont commit the transactions.
 EXIT          // Exit from the program.
