const grpc = require("@grpc/grpc-js");
const PROTO_PATH = "./news.proto";
var protoLoader = require("@grpc/proto-loader");

const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};
var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);
const newsProto = grpc.loadPackageDefinition(packageDefinition);

let news = [
  { id: "1", title: "Note 1", body: "Content 1", postImage: "Post image 1" },
  { id: "2", title: "Note 2", body: "Content 2", postImage: "Post image 2" },
];
function HelloWord(_, callback) {
    callback(null, {mensaje: "holaa"});
}

function GetAllNews(_, callback) {
    callback(null, news);
}



function main() {
    var server = new grpc.Server();
    server.addService(newsProto.NewsService.service, /*{HelloWord: HelloWord},*/{GetAllNews:GetAllNews});
    server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
      server.start();
    });
  }
  
  main ();