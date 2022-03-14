module.exports = (req, res) => {
  const { name = "World" } = req.query;
  res.send(`Hello ${name}!`);
};
