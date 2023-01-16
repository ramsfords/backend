import { Hono } from "hono";
const index = new Hono();

index.get("*", async (c) => {
    let authTokenKey = c.req.headers.get("Authorization")
    //@ts-ignore
    let auth_tokenss =  await AUTHTOKENS.get(authTokenKey)
    if(auth_tokenss == null || auth_tokenss == undefined || auth_tokenss == ""){
        return c.text("authorization not found", 404)
    }
    let cleanUrl = c.req.url.replace("/firstshipper", "")
    cleanUrl = c.req.url.replace("/menuloom", "")
    let newUrl = new URL(cleanUrl)
    newUrl.host = "api.firstshipper.com"
    newUrl.hostname = "api.firstshipper.com"
    //@ts-ignore
    let newRes = await fetch(newUrl.toString(), c.req)
    let resdata = await newRes.json()
    let status  =  newRes.status
    return c.json({
        body: resdata,
        status: status
    });
});

export default index;