require("dotenv").config();
const express = require("express");
const fetch = require("node-fetch");
const jwt = require("jsonwebtoken");
const moment = require("moment");
const _ = require("lodash");

const app = express();
const router = express.Router();
const port = process.env.PORT || 3000;

const decodeToken = (token) => {
  try {
    token = token?.replace("Bearer", "");
    const decode = jwt.verify(token, process.env.JWT_SIGNATURE || "SECRET");
    return decode;
  } catch (err) {
    return err;
  }
};

const isAuth = (req, res, next) => {
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
  } else res.send({ message: "Unauthorized" });
};

router.use(isAuth);

router.get("/products/aggregate", async (req, res) => {
  try {
    if (req?.context?.auth?.role !== "admin") {
      return res.send({ message: "Unauthorized", payload: req?.context?.auth });
    }

    const request = await fetch(process.env.STEIN_EFISHERY);
    const response = await request.json();

    let sizes = [];

    const data = response
      .filter((res) => res.size !== null)
      .map((result) => {
        const tgl_parsed = new Date(result?.tgl_parsed);
        const currentDate = new Date();

        const year = moment(tgl_parsed || currentDate).format("YYYY");
        const month = moment(tgl_parsed || currentDate).format("MM");
        const week = moment(tgl_parsed || currentDate).format("w");

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
        max: _.max(sizes),
      },
    });
  } catch (err) {
    res.send({ message: err });
  }
});

router.get("/products", async (req, res) => {
  try {
    if (
      req?.context?.auth?.role == "admin" ||
      req?.context?.auth?.role == "user"
    ) {
      const request = await fetch(process.env.STEIN_EFISHERY);
      const response = await request.json();
      const data = response.filter((res) => res.uuid !== null);
      res.send(data);
    } else res.send({ message: "Unauthorized", payload: req?.context?.auth });
  } catch (err) {
    res.send({ message: err });
  }
});

app.use(router);

app.listen(port, () => {
  console.log(`App running in ${port}`);
});
