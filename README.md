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

# Credits

- [HardHat Tutorial](https://hardhat.org/tutorial/testing-contracts)
- [yuichiroaoki.medium.com](https://yuichiroaoki.medium.com/testing-erc20-smart-contracts-in-typescript-hardhat-9ad20eb40502)

![PogChamp](https://blog.cdn.own3d.tv/resize=fit:crop,height:400,width:600/pKwIyI8RyGtPW35ZFg2m)
