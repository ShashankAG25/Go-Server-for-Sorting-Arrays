Go server listening on port 8000 with two endpoints (/process-single and /process-concurrent)<br/>
The /process-single endpoint, sort each sub-array sequentially and the /process-concurrent endpoint, sort each sub-array concurrently <br/>
<br/>
Test run :<br/>
// sorting each sub-array sequentially<br/>
// POST request with the JSON input<br/>
$ curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[9, 8, 7], [6, 5, 4], [3, 2, 1]]}' https://go-for-sorting-arrays.onrender.com/process_single
<br/>
//Response<br/>
{"sorted_arrays":[[7,8,9],[4,5,6],[1,2,3]],"time_ns":125262}<br/>
<br/>
// sorting each sub-array concurrently<br/>
// POST request with the JSON input<br/>
$ curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[9, 8, 7], [6, 5, 4], [3, 2, 1]]}' https://go-for-sorting-arrays.onrender.com/process_concurrent<br/>
<br/>
//Response<br/> 
{"sorted_arrays":[[7,8,9],[4,5,6],[1,2,3]],"time_ns":7820}
<br/>

 
