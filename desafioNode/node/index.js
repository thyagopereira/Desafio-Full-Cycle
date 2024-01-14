const express = require('express');
const { faker } = require(`@faker-js/faker`);

const app = express()
const port = 3000
const config = {
    host: 'db',
    user: 'root',
    password: 'fullcycleDB',
    database:'nodedb'
};

const mysql = require('mysql')
const connection = mysql.createConnection(config)

app.get('/', (req,res) => {
    const nome = faker.internet.userName();
    connection.query(`INSERT INTO people(nome) values('${nome}');`)
    connection.query(`SELECT * FROM people;`, (error, result, fields) => {
        res.send(`
        <h1>Full cycle Rocks!</h1>
        <li>
            ${!!result.length ? result.map(elem => `<li>${elem.nome}<li>`).join('') : ""}
        </li>
        `)
    })
})

app.listen(port, ()=> {
    console.log('Rodando na porta ' + port)
})