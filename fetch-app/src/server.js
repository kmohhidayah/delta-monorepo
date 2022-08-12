require('dotenv').config();
const express = require('express')
const fetch = require("node-fetch")
const jwt = require("jsonwebtoken")

const app = express()
const router = express.Router()
const port = 3000;

const decodeToken = (token) => {
  token = token.replace("Bearer", "")
  const decode = jwt.verify(token, "secret", (_, decode) => decode)
  return decode
}

const isAuth = async (req, _, next) => {
  if (req.headers.authorization && req.headers.authorization.split(" ")[0] === "Bearer") {
    const token = req.headers.authorization.split(" ")[1];
    const payload = decodeToken(token)
    
    req.context = {
      auth: payload
    }

    return next()
  }
}

router.use(isAuth)

router.get("/", async (req, res) => {
  console.log(req.context.auth)

  if (req.context.auth.roles !== "user") {
    throw new Error("Unauthorized")
  }

  const request = await fetch(process.env.STEIN_EFISHERY);
  const response = await request.json();
  

  const data = response.filter(res => res.size !== null).map(result => ({
    size: result.size,
    area_provinsi: result.area_provinsi
  })) || []
  
  res.send(data)
})

app.use("/", router)

app.listen(port, () => {
  console.log(`App running in ${port}`)
})
