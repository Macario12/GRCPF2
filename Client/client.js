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

const client = new GamesService.GameService('0.0.0.0:50051',
grpc.credentials.createInsecure());


module.exports = client;

