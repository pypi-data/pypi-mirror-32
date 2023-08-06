# apigett

#### Doc: [pt-br](#ptbr); [eng-usa](#engusa).

<h2 id="ptbr">Documentação em português:</h2>

Esta aplicação pega os resultados de uma API, baseada em JSON, e os salva como ".json", ".csv" ou ".txt".

### Requerimentos:

- pip v. >= 9.0.1;
- python v. >= 3.5.2;
- requests v. >= 2.9.1

### Ambiente:

- Linux

### Níveis:

- apigett (principal)
  - apigett
    - __init__.py
  - LICENSE
  - README.md
  - MANIFEST.in
  - setup.py

### Instalação:

Na linha de comando:

  *pip install apigett*

### Comandos:

- **-add**: Adiciona uma API;
  - Ex: *apigett -add "apiName" "apiEndPoint".*
- **-doc**: Mostra a documentação;
  - Ex: *apigett -doc*
- **-license**: Mostra a licença utilizada;
  - Ex: *apigett -license.*
- **-list**: Mostra uma lista de APIs registradas;
  - Ex: *apigett -list.*
- **-query**: Faz uma consulta em uma API registrada;
  - Ex: *apigett -query "apiName" "apiQueryVar" "queryArg" "apiResultKey".*
- **-remove**: Remove uma API registrada;
  - Ex: *apigett -remove "apiName"*
- **-version**: Mostra a atual versão.
  - Ex: *apigett -version.*

Todos os valores de variáveis devem ser passados, obrigratoriamente, com aspas.

### Variáveis:

- **apiName**: O nome da API, não necessáriamente o verdadeiro;
- **apiEndPoint**: A URL base da API;
- **apiQueryVar**: A variável GET que a API utiliza em sua URL base;
- **apiResultKey**: A chave, no JSON retornado pela API, onde se encontram os resultados desejados;
- **queryArg**: O que se deseja buscar na consulta.

### Exemplo:

No exemplo a seguir, será utilizada a [API xeno-canto 2.0](https://www.xeno-canto.org/article/153) (Jonathon Jongsma, 2013), onde é possível obter dados de ocorrências de espécie do mundo inteiro. Essa API utiliza a URL base (*endpoint*) "`http://www.xeno-canto.org/api/2/recordings`" e o parametro "query" para realizar as consultas via método GET, portanto, a variável "apiEndPoint" terá o valor da URL base, e a variável "apiQueryVar" terá o valor do parametro "query". Quando o nome de uma espécie é inserido na consulta, a mesma retorna  um objeto JSON com as seguintes chaves: "numRecordings", "numSpecies", "page", "numPages" e "recordings". A chave "recordings" se mostra mais interessante, nela estão todos os registros da espécie consultada, portanto, a consulta será realizda utilizando essa chave como valor da variável "apiResultKey". Como a API que está sendo utilizada faz uma consulta pelo nome científico de uma espécie, devemos então, obviamente, colocar esse nome como valor da variável "queryArg".

Para prosseguir com a consulta, antes, deve-se adicionar a API no apigett. Para isso, utiliza-se o comando "-add".

*apigett* **-add** "apiName" "apiEndPoint"

Com os dados da API:

*apigett* **-add** "xeno-canto 2.0" "`https://www.xeno-canto.org/article/153`"

Com a API adicionada no apigett, pode-se seguir com a consulta. Para isso, utiliza-se o comando "-query":

*apigett* **-query** "apiName" "apiQueryVar" "queryArg" "apiResultKey"

Com os dados da API (adicionada):

*apigett* **-query** "xeno-canto 2.0" "query" "chamaeza meruloides" "recordings"

Após a consulta, se algum valor for retornado, o apigett irá mostrar-lo na tela, e perguntará se é desejado salva-lo em um arquivo, solicitando como resposta, somente, "y" (sim), ou qualquer outro valor para não. Caso a resposta seja "y", o apigett pedirá que o usuário escolha uma das três opções de formato (JSON, CSV e TXT), sendo que cada opção estará relacionada à um número, solicitando que a resposta seja, somente, um desses números. Após, será solicitado o local onde se deseja salvar esse arquivo, onde deve-se colocar como resposta, o caminho (path) inteiro, por exemplo, /home/user/Documents.

<h2 id="engusa">English documentation:</h2>

This application takes the results of an API, based on JSON, and saves them as ".json", ".csv" or ".txt".

### Requiriments:

- pip v. >= 9.0.1;
- python v. >= 3.5.2;
- requests v. >= 2.9.1

### Environment:

- Linux

### Scaffolding:

- apigett (root)
  - apigett
    - __init__.py
  - LICENSE
  - README.md
  - MANIFEST.in
  - setup.py

### Installation:

On the command line:

  *pip install apigett*

### Commands:

- **-add**: Adds an API/ Add an API;
  - Ex: *apigett -add "apiName" apiEndPoint.*
- **-doc**: Shows the documentation;
  - Ex: *apigett -doc*
- **-license**: Shows the used license;
  - Ex: *apigett -license.*
- **-list**: Shows a list of registered APIs;
  - Ex: *apigett -list.*
- **-query**: Do a query in a registered API;
  - Ex: *apigett -query "apiName" "apiQueryVar" "queryArg" "apiResultKey".*
- **-remove**: Removes a registered API;
  - Ex: *apigett -remove "apiName"
- **-version**: Show the current version.
  - Ex: *apigett -version.*

All variable values ​​must be passed with double or single quotation marks.

### Variables:

- **apiName**: The name of the API, not necessarily the true one;
- **apiEndPoint**: The API's URL base;
- **apiQueryVar**: The GET variable in the API's URL base;
- **apiResultKey**: The key, in the JSON returned by the API, where the desired results are found;
- **queryArg**: What to look for in the query.

### Example:

In the following example, the [xeno-corner API 2.0] (https://www.xeno-canto.org/article/153) (Jonathon Jongsma, 2013) will be used, where it is possible to obtain data of species occurrences in the world all. This API uses the base URL (* endpoint *) "`http://www.xeno-canto.org/api/2/recordings`" and the "query" parameter to perform queries via the GET method, so the variable "apiEndPoint "will have the value of the base URL, and the variable" apiQueryVar "will have the value of the" query "parameter. When the name of a species is inserted into the query, it returns a JSON object with the following keys: "numRecordings", "numSpecies", "page", "numPages" and "recordings". The key "recordings" is more interesting, in it are all records of the species consulted, so the query will be performed using this key as the value of the variable "apiResultKey". Since the API being used makes a query by the scientific name of a species, then we must obviously put that name as the value of the "queryArg" variable.

To proceed with the query, first, add the API in apigett. To do this, use the command "-add".

*apigett* **-add** "apiName" "apiEndPoint"

With API data:

*apigett* **-add** "xeno-canto 2.0" "`https://www.xeno-canto.org/article/153`"

With the API added in apigett, you can proceed with the query. For this, the command "-query" is used:

*apigett* **-query** "apiName" "apiQueryVar" "queryArg" "apiResultKey"

With API data (added):

*apigett* **-query** "xeno-canto 2.0" "query" "chamaeza meruloides" "recordings"

After the query, if any value is returned, apigett will show it on the screen, and will ask if it is desired to save it in a file, requesting only "y" (yes), or any other value to recuse. If the answer is "y", apigett will ask the user to choose one of the three format options (JSON, CSV and TXT), each option being related to a number, requesting that the answer be only one of those numbers . After that, you will be prompted for the location where you want to save this file, where you must put the entire path, for example, /home/user/Documents.


