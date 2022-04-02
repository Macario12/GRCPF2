const express = require('express')
const bodParser = require('body-parser')
let cors = require('cors')


const client = require("./client");

const app = express()


app.use(bodParser.json({limit:'50mb', extended:true}))
app.use(bodParser.urlencoded({limit:'50mb', extended:true}))
app.use(cors())

app.post('/Games',(req,res)=>{

  var respuesta;
    client.AddGame(req.body, function(err, response) {
      //console.log(err);
      console.log('Response:', response);
      res.send(response)
    });
    
})

app.listen('3000', ()=>{
    console.log("Servidor en puerto 3000")
})