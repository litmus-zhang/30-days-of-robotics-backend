package services

// The Peer Review Algorithm
//- Get the submissions in the usertask table if length > 2, unique users submitted, grade is not 0
//- Get the userid and grader_id from the submissions
//- Create a hashmap where the userID is the key and the graderID is the value,
//- initialize the hashmap keys to the userID
//- Go through the hashmap, set the values to the current value if the key != value, after setting, pop the values.
//- Run this algorithm daily at 00:00
