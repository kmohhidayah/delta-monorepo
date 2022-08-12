require('dotenv').config();
const express = require('express')
const fetch = require("node-fetch")
const jwt = require("jsonwebtoken")
const moment = require('moment')

const app = express()
const router = express.Router()
const port = 3000;

const decodeToken = (token) => {
  token = token.replace("Bearer", "")
  const decode = jwt.verify(token, "secret", (_, decode) => decode)
  return decode
}

const isAuth = (req, _, next) => {
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

  if (req.context.auth.roles !== "admin") {
    throw new Error("Unauthorized")
  }

  const request = await fetch(steinUrl);
  const response = await request.json();
  
  const data = response.filter(res => res.size !== null).map(result => {
    let tglParsed = result.tgl_parsed
    
    if (!result.tgl_parsed) tglParsed = new Date();
    
    tglParsed = moment(new Date(result.tgl_parsed)).format(
      "YYYY-MM-DD HH:mm:ss"
    )

    const year = moment(new Date(result.tgl_parsed)).format("YYYY")
    const month = moment(new Date(result.tgl_parsed)).format("MM")
    const week = moment(new Date(result.tgl_parsed)).format("w")
    
    return {
      tgl_parse: tglParsed,
      year,
      month,
      week,
      province: result.area_provinsi,
      size_aggregate: {
        min: result.size < 10 ? result.size : 9,
        median: result.size < 50 ? result.size : 40,
        max: result.size > 50 ? result.size : 50
      }
    }
   }) || []
  
  res.send(data)
})

app.use("/", router)

app.listen(port, () => {
  console.log(`App running in ${port}`)
})
