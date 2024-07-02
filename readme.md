# crypto-rate-bot

## Описание

Телеграм бот помогает отслеживать курсы криптовалют и отправлять оповещения об изменении их цены.

## Команды бота

- `/current` - показать текущие цены всех криптовалют с Binance. Для фильтрации валют можно добавить название в формате BTC/BNB/etc. Например `/current BNB BTC`.
- `/price` - показать изменение цены криптовалюты за определенный период. Например `/price BNB BTC 5m`.
- `/price_background` - запустить фоновую задачу, которая будет оповещать об изменении цены за последние 5 минут на 3 процента роста.

---

## Лицензия
Этот проект лицензируется по лицензии MIT. Подробности смотрите в файле LICENSE.


# crypto-rate-bot

## Description

Telegram bot helps track cryptocurrency rates and send alerts about their price changes.

## Bot Commands

- `/current` - show current prices of all cryptocurrencies from Binance. To filter currencies, you can add the name in the format BTC/BNB/etc. For example, `/current BNB BTC`.
- `/price` - show the price change of a cryptocurrency for a specific period. For example, `/price BNB BTC 5m`.
- `/price_background` - start a background task that will notify about price changes of 3 percent growth over the last 5 minutes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
