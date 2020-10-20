import express from "express";
import url from "url";
import bodyParser from "body-parser"

function fullUrl(req) {
	return url.format({
		protocol: req.protocol,
		host: req.get('host'),
		pathname: req.originalUrl
	});
};

const app = express();
app.use(bodyParser.json());

const hostname = "0.0.0.0";
const port = 3000;


app.get("/helloworld/node", (_, res) => {
	res.send("Hello World!!")
});

app.get("*", (req, res) => {
	res.redirect(`/helloworld/404.html?url=${fullUrl(req)}`);
});

app.listen(port, hostname, () => {
	console.log(`server running at http://${hostname}:${port}/`);
});