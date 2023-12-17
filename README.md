Go server listening on port 8000 with two endpoints (/process-single and /process-concurrent)
The /process-single endpoint, sort each sub-array sequentially and the /process-concurrent endpoint, sort each sub-array concurrently 

Test run :
// sorting each sub-array sequentially
// POST request with the JSON input
$ curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[9, 8, 7], [6, 5, 4], [3, 2, 1]]}' http://localhost:8000/process_single

//Response
{"sorted_arrays":[[7,8,9],[4,5,6],[1,2,3]],"time_ns":55326}

// sorting each sub-array concurrently
// POST request with the JSON input
$ curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[9, 8, 7], [6, 5, 4], [3, 2, 1]]}' http://localhost:8000/process_concurrent

//Response
{"sorted_arrays":[[7,8,9],[4,5,6],[1,2,3]],"time_ns":8613}

 
