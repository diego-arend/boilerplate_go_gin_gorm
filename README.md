# Descrição

# boilerplate_go_gin_gorm

Boilerplate WebApi em Golang com frameworks GIN para tratamento de requests e GORM para comunicação com banco de dados e midleware de autenticação.

# Disclaimer

Este boilerplate está emc onstante desenvolvimento.
Portanto é altamente recomendável que sente-se confortávelmente em sua cadeira, pegue um refresco e leia este documento. Ele fornecerá informações importantes para que você seja mais ágil e não perca minutos importantes do seu descanço Debugando informações conhecidas :).

# Kit básico de funcionamento

Ë ncessário ter o GoLang instalado em sua máquina. Para maiores informações acessar:
`https://go.dev`

# Comandos

- Após clonar o repositório digite o comando `go mod tidy` para instalar as depências necessárias.
- Para iniciar o projeto digite `go run main.go`

# Detalhes de arquitetura.

- Este boilerplate utiliza a framework GIN-GONIC para tratamento de resquest. Mais informações em `https://gin-gonic.com/`

- Para comunicação com o banco de dados é utilizado o framework GORM. Mais informações em `https://gorm.io/`

- As migrates já são executadas automaticamente toda vez que o projeto é incializado com o comando `go run main.go` através da função <pre>
<code>func RunMigrations(db \*gorm.DB) {</code>
<code> db.AutoMigrate(models.Car{})</code>
<code> db.AutoMigrate(models.User{})</code>
<code>}</code>
</pre> 
contida no arquivo `./database/migrations.go`. No arquivo "./database/database.go" a função `StartDB` contém o comando <pre><code>migrations.RunMigrations(db)</code></pre> que executa a migração sempre que a conexão com o banco for inicializada.

- No arquivo `./routes/router.go`estão contidas as rotas apontando para os devidos controllers.

# Fontes de pesquisa e créditos
- Série sobre GoLang do canal `https://www.youtube.com/@2yuri`
- Série sobre GoLang do canal `https://www.youtube.com/@GustavoNoronhaDev`


