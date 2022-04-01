const grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
const PROTO_PATH ="../protos/Game.proto";

const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);

const GameService = grpc.loadPackageDefinition(packageDefinition).GameService;

const client = new GameService(
  "127.0.0.1:50051",
  grpc.credentials.createInsecure()
);
/*
client.HelloWord({}, function(err, response) {
  console.log('Greeting:', response);
});*/

client.AddGame({game_id:1,players:32}, function(err, response) {
  console.log('Greeting:', response);
});
//module.exports = client;