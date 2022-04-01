var PROTO_PATH = '../protos/Game.proto';
const grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");


const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const GamesService = grpc.loadPackageDefinition(packageDefinition).Game;
/*
const client = new NewsService(
  "127.0.0.1:50051",
  grpc.credentials.createInsecure()
);
/*
client.HelloWord({}, function(err, response) {
  console.log('Greeting:', response);
});*/
/*
client.AddGame({Game_id:1,Players:23}, function(err, response) {
  console.log(err);
  console.log('Response:', response);
});*/
//module.exports = client;*/

function main() {
  var client = new GamesService.GameService('localhost:50051',
                                       grpc.credentials.createInsecure());

  client.AddGame({game_id:3,players:23}, function(err, response) {
    console.log('Greeting:', response.message);
  });
}

main();