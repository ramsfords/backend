import { Hono } from 'hono'
import { cors } from 'hono/cors'
import index from './routes/index'
const app = new Hono()
let auth_tokens:KVNamespace
app.use("*", cors(
	{
		origin: 'http://menuloom.com',
		allowHeaders: ['X-Custom-Header', 'Upgrade-Insecure-Requests', 'Content-Type', 'Authorization', 'Access-Control-Allow-Origin', 'Access-Control-Allow-Headers', 'Access-Control-Allow-Methods', 'Access-Control-Allow-Credentials', 'Access-Control-Max-Age'],
		allowMethods: ['POST', 'GET', 'OPTIONS', "PUT", "DELETE"],
		exposeHeaders: ['Content-Length', 'X-Kuma-Revision', "X-Custom-Header", "Upgrade-Insecure-Requests", "Content-Type", "Authorization", "Accept","Allow-Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Max-Age"],
		maxAge: 600,
		credentials: true,
	  }
))
app.route("/", index);
app.get("/name", async (c, env) =>{
	const value  = auth_tokens.get("token")
	console.log("value is ", value)
	try {
		let url = new URL("http://127.0.0.1:8090/api/ping")
		// url.hostname = "api.firstshipper.com"
		// url.host = "api.firstshipper.com"
		let pathParams = c.req.query("name") 
		console.log("path params is ", pathParams)
		 //@ts-ignore
		let newRes = await fetch(url.toString(), c.req)
		let resdata = await newRes.json()
		console.log("respone data was ", resdata)
		return c.json(resdata)
	} catch (error) {
		console.log("error is ", error)
		return new Response("error")
	}
})
app.fire()
