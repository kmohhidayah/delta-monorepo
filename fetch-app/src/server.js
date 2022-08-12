require("dotenv").config();
const express = require("express");
const fetch = require("node-fetch");
const jwt = require("jsonwebtoken");
const moment = require("moment");
const _ = require("lodash");

const app = express();
const router = express.Router();
const port = 3000;

const decodeToken = (token) => {
  token = token.replace("Bearer", "");
  const decode = jwt.verify(token, process.env.JWT_SIGNATURE, (_, decode) => decode);
  return decode;
};

const isAuth = (req, _, next) => {
  if (
    req.headers.authorization &&
    req.headers.authorization.split(" ")[0] === "Bearer"
  ) {
    const token = req.headers.authorization.split(" ")[1];
    const payload = decodeToken(token);

    req.context = {
      auth: payload,
    };

    return next();
  }
};

router.use(isAuth);

router.get("/api/v1/products/aggregation", async (req, res) => {
  if (req.context.auth.role !== "admin") {
    throw new Error("Unauthorized");
  }

  const request = await fetch(process.env.STEIN_EFISHERY);
  const response = await request.json();

  let sizes = [];

  const data = response.filter((res) => res.size !== null).map((result) => {
    const year = moment(new Date(result.tgl_parsed) || new Date()).format(
      "YYYY",
    );
    const month = moment(new Date(result.tgl_parsed) || new Date()).format(
      "MM",
    );
    const week = moment(new Date(result.tgl_parsed) || new Date()).format("w");

    sizes.push(parseInt(result.size, 10));
  
    return {
      year,

      month,
      week,
      price: parseInt(result.price, 10),
      province: result.area_provinsi,
      size: parseInt(result.size, 10),
    };
  }) || [];

res.send({

    aggregate_by_province: _.keyBy(data, (res) => res.province),
    aggregate_by_week: _.keyBy(data, (res) => parseInt(res.week, 10)),
    aggregate_by_price: _.keyBy(data, (res) => res.price),
    average: {
      min: _.min(sizes),
      max: _.max(sizes)
    }
  })
});

app.use("/", router);

app.listen(port, () => {
  console.log(`App running in ${port}`);
});
