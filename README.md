# PogToken

[PogChamp Token](https://goerli.etherscan.io/address/0xd727259569e631ad36d12b9c18cd754b69a8c4de)

Исходный код контракта: [POGToken.sol](./contracts/POGToken.sol)

## Взаимодействие с контрактом

Реализовано на Go в директории [client](./client/).

### Доступные команды

Необходимые env:

- PRIVATE_KEY &mdash; приватный ключ тестовой сети
- INFURA_URL &mdash; секретный URL infura с API токеном
- CONTRACT_ADDRESS &mdash; адрес контракта в сети для взаимодействия

В директории [client](./client/) можно выполнить следующие команды:

- `go run . -- add <address> <username>` &mdash; добавить пользователя
- `go run . -- remove <address>` &mdash; удалить существующего пользователя
- `go run . -- logs add` &mdash; посмотреть лог добавлений
- `go run . -- logs remove` &mdash; посмотреть лог удалений

### Демонстрация взаимодействия с контрактом

```bash
[$] go run . -- add 0xD727259569e631AD36d12b9C18cd754b69A8C4de blockchain
viewer has been added: 0xe22ad90a418314fb8bdd14066f3cdcc0952a8fff23247da6400d317d8d3deddc

[$] go run . -- logs add
[0xD727259569e631AD36d12b9C18cd754b69A8C4de] New viewer added:
 - address:  0xB87e1B6A462660E26Ec8913513aB86b084A4F46c
 - username:  lodthe

[0xD727259569e631AD36d12b9C18cd754b69A8C4de] New viewer added:
 - address:  0xD727259569e631AD36d12b9C18cd754b69A8C4de
 - username:  lodtheeee

[0xD727259569e631AD36d12b9C18cd754b69A8C4de] New viewer added:
 - address:  0xD727259569e631AD36d12b9C18cd754b69A8C4de
 - username:  blockchain

[$] go run . -- remove 0xD727259569e631AD36d12b9C18cd754b69A8C4de
viewer has been removed: 0xfd5d2205514f44fecc73e1b6cd6e5052ac5c32f99582f06a4a8096386f34cb11

[$] go run . -- logs remove
[0xD727259569e631AD36d12b9C18cd754b69A8C4de] A viewer has been removed:
 - address:  0xD727259569e631AD36d12b9C18cd754b69A8C4de
 - username:  blockchain
```

# Ответы на вопросы

## Какие есть расширения для обозревателя, которые умеют взаимодействовать с Ethereum?

- Кошельки: MetaMask, Trust Wallet, Opera, ZenGo Wallet, CoinBase Wallet.

- Обозреватели: EtherScan, BlockScout.

- Удобный оступ к блокчейну через API: Infura.

## Что делает MetaMask, чтобы web-приложение могло взаимодействовать с блокчейном через JavaScript?

MetaMask встраивает javascript библиотеку web3js, к которой обращается веб-приложение. Также MetaMask хранит в local storage браузера аутентификационные данные, чтобы взаимодействовать с блокчейном через json-rpc.

## Каким образом обрабатываются пользовательские данные в web-приложении, чтобы они были приняты блокчейном?

Чтобы пользовательские данные могли быть приняты блокчейном, они должны быть закодированы в виде байтового кода, который соответствует стандарту ABI (Application Binary Interface). Стандарт ABI определяет, как функции и переменные в смарт-контракте должны быть представлены в виде байтового кода.

Логика кодирования в ABI как правило находится в библиотеке, которая может получать ABI контракта напрямую из сети, и в соответствии с данным ABI "запаковывать данные".

Отправка транзакций / прочие базовые функции реализуются через готовый общий интерфейс.

## Какие есть способы взаимодействия с блокчейном, кроме RPC-сервера и в чём их особенность?

Со стороны клиента есть следующие варианты:

- Веб-приложения (которые часто базируются на js-библиотеках)
- Command line утилиты
- SDK для разных языков программирования

Однако внутри как правило все равно лежит общение либо через RPC, либо через нативный протокол.
Различные варианты отличаются по удобности и своей эффективности. Если необходимо разрабатывать вебприложение, подойдёт аналог web3js. Если необходимо выполнять логику где-то на бэкенде, лучшим вариантом будет SDK.

## Какие есть ещё способы получения данных из блокчейна, кроме отправка call методов и событий?

Есть следующие источники информации:

- Чтение состояния контракта открыто, и нем может содержаться полезная информация (например, балансы ERC20)

- Получение баланса кошелька

- Получение блоков блокчейна и содержащихся в них транзакциях

Получить их как правило можно либо через нативный протокол блокчейна, либо через RPC (например, json-rpc в Ethereum).

# Credits

- [HardHat Tutorial](https://hardhat.org/tutorial/testing-contracts)
- [yuichiroaoki.medium.com](https://yuichiroaoki.medium.com/testing-erc20-smart-contracts-in-typescript-hardhat-9ad20eb40502)

![PogChamp](https://blog.cdn.own3d.tv/resize=fit:crop,height:400,width:600/pKwIyI8RyGtPW35ZFg2m)
