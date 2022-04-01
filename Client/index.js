const client = require("./client");

client.getAllNews({}, (error, news) => {
  if (!error) throw error;
  console.table(news.m);
});