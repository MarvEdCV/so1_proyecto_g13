const redis = require('redis');
var PROTO_PATH = './proto/config.proto';
const crypto = require('crypto');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var usactar_proto = grpc.loadPackageDefinition(packageDefinition).usactar;

const HOST_REDIS = 'azureCache-Redis.redis.cache.windows.net'
const KEY_REDIS = 'M63eFLchNx4pcX11OR6qJYsfjvX5wuumXAzCaMHhark='
const client = redis.createClient({
    //REDIS FOR TLS
    url:`rediss://${HOST_REDIS}:6380`,
    password: KEY_REDIS
});
async function AddPrediction(call,callback){
    let id = crypto.randomUUID();
    console.log(id)
    await client.set(id,JSON.stringify({
        team1: call.request.team1,
        team2: call.request.team2,
        score: call.request.score,
        phase: call.request.phase
    }))
    callback(null,{message: `Caso insertado en la base de datos${id}`})
}

async function main(){
    /*await client.connect();

    console.log(`Ping test`);
    console.log(`Redis cache response:${await client.ping()}`)

    // Simple get and put of integral data types into the cache
    console.log("\nCache command: GET Message");
    console.log("Cache response : " + await client.get("Message"));

    console.log("\nCache command: SET Message");
    console.log("Cache response : " + await client.set("Message",
        "Hola esto es una prueba"));

    // Demonstrate "SET Message" executed as expected...
    console.log("\nCache command: GET Message");
    console.log("Cache response : " + await client.get("Message"));

    // Get the client list, useful to see if connection list is growing...
    console.log("\nCache command: CLIENT LIST");
    console.log("Cache response : " + await client.sendCommand(["CLIENT", "LIST"]));
     */
    await client.connect();
    var server=new grpc.Server();
    server.addService(usactar_proto.GrpcConnection.service,{
        AddPrediction:AddPrediction
    });
    server.bindAsync('0.0.0.0:50051',grpc.ServerCredentials.createInsecure(),() =>{
    server.start();
    console.log('gRCP server on port 50051')
})
    console.log("\nEsperando trafico");

}
main();
