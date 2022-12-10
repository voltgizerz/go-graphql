<h1 align="center">GraphQL GO using GQLGEN 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-0.0.1-blue.svg?cacheSeconds=2592000" />
  <a href="https://twitter.com/felix\_fernand0" target="_blank">
    <img alt="Twitter: felix_fernand0" src="https://img.shields.io/twitter/follow/felix_fernand0.svg?style=social" />
  </a>
</p>

> Example go graphql using library gqlgen

### 🏠 [Homepage](https://graphql-poke.herokuapp.com/)

### ✨ [Playground](https://graphql-poke.herokuapp.com/)

## Installation

```sh
create .env file
DATA_SOURCE="your-database-source"
PORT=8080
GOLANG_ENV="test"
```

## Test

```sh
make test
```

## Usage

```sh
go run server.go
```

## Linter

```sh
make check
```

## Fix Linter

```sh
make format
```

## GraphQL Features

💡 **Request Header**

```
{
  "Authorization":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzA2NzI0MjQsImlzcyI6Impob24iLCJVc2VySUQiOjgxLCJVc2VybmFtZSI6ImZlbGl4IiwiTmFtZSI6IkZlbGl4IiwiSXNBZG1pbiI6dHJ1ZSwiRXhwaXJlc0F0IjoxNjcwNjcyNDI0fQ.noKdYN3fhr7iwQmlDy1xYn1FqLsZ4gY-ORkWf3oeP_I"
}
```

💡 **Query**
* Login get token
```
query login {
  login(input: {username: "felix", password: "felix"}) {
    success
    message
    token
  }
}
```
* Show all data pokemon (limit, offset, sort)
```
query pokemons {
  pokemons {
    id
    height
    name
  }
}
```
* Show data pokemon by ID
```
query pokemon {
  pokemon(pokemonID: 1) {
    name
    types {
      id
    }
    height
    weight
  }
}
```

💡 **Mutation**

* Update Pokemon
```
mutation update{
  updatePokemon(input:{id: 1}){
    success
    pokemon{id}
  }
}
```
* Delete Pokemon
```
mutation delete {
  deletePokemon(input:{id:11}){
    success
  }
}
```

## Author

👤 **Felix Fernando**

* Website: voltgizerz.github.io
* Twitter: [@felix_fernand0](https://twitter.com/felix_fernand0)
* Github: [@voltgizerz](https://github.com/voltgizerz)
* LinkedIn: [@felix-fernando-wijaya](https://linkedin.com/in/felix-fernando-wijaya)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/voltgizerz/go-graphql/issues). 

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_