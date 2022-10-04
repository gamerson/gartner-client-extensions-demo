const cookieParser = require('cookie-parser');
const express = require('express');
const path = require('path');
const winston = require('winston');


const logger = winston.createLogger({
  level: 'info',
  format: winston.format.json(),
  transports: [new winston.transports.Console()],
});

const app = express();

app.use(express.json({ limit: 64 * 1000 }));
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.post('/coupons/issued', function(req, res) {
  try {
    logger.info("HERE IS THE OBJECT JSON: " + JSON.stringify(req.body));
    res.status(201).send('OK');
  } catch (exc) {
    logger.info(exc);
    res.status(400).send('Invalid JSON');
  }
});

module.exports = app;
