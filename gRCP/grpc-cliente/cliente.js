var PROTO_PATH = './proto/config.proto';

var parseArgs = require('minimist');
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

var argv = parseArgs(process.argv.slice(2), {
    string: 'target'
});
var target;

//VARIABLES API
const express = require('express');
var cors = require('cors');
const app = express();

app.use(express.json());
app.use(cors());


if (argv.target) {
    target = argv.target;
} else {
    target = '20.81.86.208:8083';
    //target = '0.0.0.0:50051';
}
var client = new usactar_proto.GrpcConnection(target,grpc.credentials.createInsecure());

app.post('/client-grcp',function (req,res){
    console.log(req.body);

    client.AddPrediction(req.body,function (err,response){
        res.status(200).json({mensaje: response.message})
    })
})

app.listen(3000,()=>{
    console.log('Servidor cliente en el puerto',3000);
})

