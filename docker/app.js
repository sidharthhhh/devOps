const express = require("express");
const app = express();


app.get("/",(req,res)=>{
    res.send("hellllllo");
})


app.listen(4000,()=>{
    console.log("server is running in port 4000");
})