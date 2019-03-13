const app = require('express')()
const bodyParser = require('body-parser')
const { Nuxt, Builder } = require('nuxt')
const axios = require('axios')

app.use(bodyParser.json())

const apiUrl = process.env.TODO_API_URL
if (!apiUrl) {
  console.log('TODO_API_URL is undefined')
  process.exit(1)
}
console.log(`TODO_API_URL=${apiUrl}`)

app.get('/api/todo', (req, res) => {
  const status = req.query.status
  axios.get(`${apiUrl}/todo?status=${status}`)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.post('/api/todo', (req, res) => {
  const param = {
    title: req.body.title,
    content: req.body.content
  }

  axios.post(`${apiUrl}/todo`, param)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

let config = require('./nuxt.config.js')
config.dev = !(process.env.NODE_ENV === 'production')
const nuxt = new Nuxt(config)
if (config.dev) {
  const builder = new Builder(nuxt)
  builder.build()
}

app.use(nuxt.render)
app.listen(3000)
