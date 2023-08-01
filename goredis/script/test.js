import http from 'k6/http'

export default function(){
    // without docker container use
    // run by k6 run ./script/test.js
    http.get("http://localhost:8000/hello")


    // with docker container 
    //http.get("http://host.docker.internal:8000/hello")
}