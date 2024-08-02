import { createServer } from 'node:http'

const server = createServer((req, res) => 
{
    res.writeHead(200, {'Content-Type': 'text/plain'});
    res.end('Hello World!\n');
});

server.listen(3000, '192.168.18.8', () => 
{
    console.log('Listening on 192.168.18.8:3000');
});