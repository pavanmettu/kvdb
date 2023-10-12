# kvdb
Simple In-memory KV Database implementation(SET, GET, BEGIN, END, COMMIT, ROLLBACK)<br />

 Commands to Use:<br />
 SET KEY VALUE // Sets Key to value. If part of Txn, till COMMIT
               // is called, its in mem. If ROLLBACK, then deleted.<br />
               
 GET KEY       // Gets value first from Pending Txns and then Global DB.<br />
 
 BEGIN         // Starts a transaction of multi keys.<br />
 
 ROLLBACK      // Takes back all members of the transacted items.<br />
 
 COMMIT        // Adds key-value pairs from transactions to DB.<br />
 
 END           // Dont commit the transactions.<br />
 
 EXIT          // Exit from the program.<br />
