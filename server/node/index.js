import express from "express";
const app = express();
const port = 3000;
const hostname = "0.0.0.0"

app.get("/helloworld/node", (_, res) => {
	res.send("Hello World!!")
});

app.listen(port, hostname, () => {
	console.log(`server running at http://${hostname}:${port}/`);
});