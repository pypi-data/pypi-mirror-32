
# Promium

[![build_status](https://gitlab.uaprom/uaprom/promium/badges/master/build.svg)](https://gitlab.uaprom/uaprom/promium/pipelines)
[![python_versions](https://img.shields.io/badge/python-2.7%2C%203.5%2C%203.6-blue.svg)](https://gitlab.uaprom/uaprom/promium/)

watch [documentation](http://doc.promium.gitlab.uaprom/)



# promium / promium & Selenium Grid & Vagga (container) in Ubuntu 16.04
----------------------------------------------------------------------------------
----------------------------------------------------------------------------------

# promium 
## Requirements
* Python3 (3.5.2)
* Selenium (3.8.0)
* PyTest (3.2.3)

### Install requirements
pip3 install -r requirements.txt

## Drivers for browsers
You can downloan drivers here http://docs.seleniumhq.org/download/

You must put drivers in folder and write path in environment variable PATH

As an alternative, you can put driver in framework folder and write path in module test_case.py:

webdriver.Chrome(executable_path="{путь к chromedriver}")

For example:
* webdriver.Chrome(executable_path="./chromedriver", chrome_options=get_chrome_options(device))

## Run tests
#### Command for run tests
* py.test tests_path

For example:
* py.test tests

#### Run tests from module. You can add path to module

For example:
* py.test tests/test_catalog_main_page.py

#### Run certain test. You can use option "-k" and test name

For example:
* py.test tests/test_catalog_main_page.py -k test_main_page_links

#### Run parametrize test with certain parameter. You can add [] with needful parameter after test name

For example:
* py.test tests/test_catalog_main_page.py -k test_main_page_links[footer_customers]

#### Run tests in multiple threads. You can use option "-n" with number of threads

For example:
* py.test tests -n3 (3 - number of threads).

#### Run tests in certain browser. You can use option with browsers name

For example:
* py.test tests --chrome

You can select one from browsers:
* --chrome
* --firefox
* --opera
* --edge
* --ie

If browser parameter does not use, then framework use Chrome by default

#### Run tests in Selenium Grid. You must run Selenium Grid, write host and port in config.py and use option "--grid"

For example:
* py.test tests --grid

#### Run tests and highlight used elements. You can use option "--highlight"

For example:
* py.test tests --highlight

#### Run last failed tests. You can use option "--last-failed"

For example:
* py.test tests --last-failed

#### Run tests with frontend errors in browsers. You can use option "--check-console"

For example:
* py.test tests --check-console

### Run tests with many options. You can use more one options

For example:
* py.test tests --opera --grid -n2 --last-failed --check-console

Run last failed tests in two threads in Opera browser in Selenium Grid and check frontend errors. The order of the options can be arbitrary.

## Framework

Package name / Module name  | Package contents / Module contents
----------------------|----------------------
lib                   | Libraries
screenshots           | Screenshots folder
selenovo              | Framework Selenovo
selenovo/elements     | Elements package
selenovo/files        | Files for tests
selenovo/base.py      | Base module
selenovo/common.py    | Common functions
selenovo/exception.py | Exceptions module
selenovo/waits.py     | Waits module
state                 | Methods for work DB using SQLSoup
tests                 | Tests
assertions.py         | Assert methods
config.py             | Configuration file (host, port, etc.)
conftest.py           | Pytest
device_config.py      | Configuration mobile device (Android, iOS)
helpers.py            | Auxiliary methods and functions
logger.py             | Log functions
pytest.ini            | Configuration file pytest
requirements.txt      | Requirements file
services.py           | Services
test_case.py          | Driver initialization, description of browser options
urls.py               | Urls methods

## Pattern Page

The pages are described in the package *"pages"*. The page class is inherited from the base Page class.
The page describes the url of the page, blocks and elements that it contains. If the block is contained only on this page, then is described above in the module.
If a block is used on several pages, then it is described in a separate module that is in the *"blocks"* package and is defined in the "blocks_mixin.py" (The page is additionally inherited from the "BlocksMixin" block).

```
from selenium.webdriver.common.by import By

from promium import Page, Block, Element, Link
from urls import collect_url

from .blocks.blocks_mixin import BlocksMixin


class ProductBlock(Block):
    all_characteristics_link = Link(
        By.CSS_SELECTOR,
        '[data-qaid="all_characteristics_link"]'
    )
    buy_button = Element(
        By.CSS_SELECTOR,
        '[data-qaid="buy_button"]'
    )


class MainPage(Page, BlocksMixin):

    url = collect_url('')

    products_block = ProductBlock.as_list(
        By.CSS_SELECTOR,
        '[data-qaid="searches_block"]'
    )
    header_block = HeaderBlock(
        By.CSS_SELECTOR,
        '[data-qaid="header_block"]'
    )

```

## Example tests and pages

Examples of tests are in */tests*

Examples of page object and blocks are in */tests/pages* 

### test_catalog_header.py

Checks the desktop version of the site using Selenium Webdriver

### test_mobile_main_page.py

Check the mobile version of the site using Selenium Webdriver

### test_catalog_main_page.py

The test that does NOT use Selenium Webdriver, but uses the library of requests

----------------------------------------------------------------------------------
----------------------------------------------------------------------------------
----------------------------------------------------------------------------------

# promium & Selenium Grid & Vagga (container) in Ubuntu 16.04
## Install Vagga

http://vagga.readthedocs.io/en/latest/installation.html#ubuntu

Create configuration file .vagga.yaml in "/home/user_name":
```
external-volumes:
  bin: /usr/bin
```

## Install requirements, browsers and drivers

Nothing to install is not necessary, everything is already in the container :)

## View available Vagga commands

At the root of the project, run the "vagga" command

* vagga

```
Available commands:
    faq                 View FAQ
    pep8                Check code style
    run-tests           Run selenium tests
    selenium-hub        Run Selenium Grid Hub on the local machine
    selenium-node       Run Selenium Grid Node
    versions            View which selenium and browsers versions is using
```

## Command "vagga lint"
### To run the lint test, run the command

* vagga lint

## Command "vagga run-tests"
### Command for run tests

* vagga run-tests -- tests_path

For example:
* vagga run-tests -- tests

### Run tests with options

* vagga run-tests -- tests/test_catalog_main_page.py -k test_main_page_links --firefox --grid -n2 --last-failed --highlight

Tests are started in analogy with py.test. All options are listed after "vagga run-tests -- "

You can select one from browsers:
* --chrome
* --firefox
* --opera

## Commands "vagga selenium-hub" and "selenium-node"
### Run Selenium Grid

Run two commands in parallel in the terminal in the following order:
* vagga selenium-hub

After running in the console, we get this:
```
13:30:13.934 INFO - Nodes should register to http://0.0.0.0:13400/grid/register/
13:30:13.934 INFO - Selenium Grid hub is up and running
```

* vagga selenium-node

After running in the console, we get this:
```
13:30:14.351 INFO - Selenium Grid node is up and ready to register to the hub
13:30:14.358 INFO - Starting auto registration thread. Will try to register every 5000 ms.
13:30:14.358 INFO - Registering the node to the hub: http://localhost:13400/grid/register
13:30:14.406 INFO - The node is registered to the hub and ready to use
```

Now you can run tests in the Selenium Grid using the "--grid" option

http://0.0.0.0:13400/grid/console

## Command "vagga versions"

* vagga versions

```
Google Chrome 63.0.3239.132 
ChromeDriver 2.35.528139 (47ead77cb35ad2a9a83248b292151462a66cd881)
Mozilla Firefox 57.0.4
geckodriver 0.19.1
Opera 50.0.2762.58
OperaDriver 2.32 (cfa164127aab5f93e5e47d9dcf8407380eb42c50)
Selenium server version: 3.8.1, revision: 6e95a6684b
```

# FAQ

* **If an error occurred during the test run:**

```
Traceback (most recent call last):
  File "/usr/local/lib/python3.5/dist-packages/_pytest/config.py", line 342, in _getconftestmodules
    return self._path2confmods[path]
KeyError: local("/work/tests")

During handling of the above exception, another exception occurred:
Traceback (most recent call last):
  File "/usr/local/lib/python3.5/dist-packages/_pytest/config.py", line 373, in _importconftest
    return self._conftestpath2mod[conftestpath]
KeyError: local("/work/conftest.py")

During handling of the above exception, another exception occurred:
Traceback (most recent call last):
  File "/usr/local/lib/python3.5/dist-packages/_pytest/config.py", line 379, in _importconftest
    mod = conftestpath.pyimport()
  File "/usr/local/lib/python3.5/dist-packages/py/_path/local.py", line 686, in pyimport
    raise self.ImportMismatchError(modname, modfile, self)
py._path.local.LocalPath.ImportMismatchError: ("work.conftest", "/home/user/workspace/promium3/conftest.py", local("/work/conftest.py"))
ERROR: could not load /work/conftest.py
```
It occurs if you run the tests first through the py.test command, and then through the vagga run-tests or in the reverse order.
You can correct the error by running the following command at the root of the project:

```
sudo find -name '*.pyc' -delete
```

* **To remove the collected vagga containers (for example, they are obsolete or take up too much disk space), you need to run the following command in the project root:**

```
vagga _clean --everything
```

* **How can I find out what's in the container? Using the following command, you can look through the terminal what is there:**

```
vagga _run promium bash
```

----------------------------------------------------------------------------------
----------------------------------------------------------------------------------
----------------------------------------------------------------------------------

# promium
## Зависимости
* Python3 (3.5.2)
* Selenium (3.6.0)
* PyTest (3.2.3)

### Установка зависимостей
pip3 install -r requirements.txt

## Драйвера браузеров
Драйвера браузеров можно скачать здесь http://docs.seleniumhq.org/download/

Скаченные драйвера положить в один каталог и указать к нему путь в PATH

Как альтернатива можно положить драйвер в корень фреймворка и указать к нему путь в модуле test_case.py при создании объекта:

webdriver.Chrome(executable_path="{путь к chromedriver}")

Пример:
* webdriver.Chrome(executable_path="./chromedriver", chrome_options=get_chrome_options(device))

## Запуск тестов
#### Для запуска тестов необходимо выполнить минимальную команду
* py.test <путь к тестам>

Пример:
* py.test tests

#### Для запуска тестов с определенного модуля нужно указать этот модуль в пути

Пример:
* py.test tests/test_catalog_main_page.py

#### Для запуска определенного теста нужно указать его в пути после параметра "-k"

Пример:
* py.test tests/test_catalog_main_page.py -k test_main_page_links

#### Если тест параметризирован, то при запуске также в [] можно указать параметр с которым нужно запустить тест

Пример:
* py.test tests/test_catalog_main_page.py -k test_main_page_links[footer_customers]

#### Для запусков тестов в несколько потоков нужно использовать параметр "-n" с количеством потоков

Пример:
* py.test tests -n3 (где 3 - количество потоков).

#### Для выбора браузера необходимо указать его в параметре

Пример:
* py.test tests --chrome

Доступны следующие браузеры:
* --chrome
* --firefox
* --opera
* --edge
* --ie

Если параметр браузера не узазать то по-умолчанию берется Сhrome

#### Чтобы запустить тесты в Selenium Grid нужно запустить грид, указать хост и порт в config.py и указать параметр "--grid" при запуске тестов

Пример:
* py.test tests --grid

#### Чтобы видеть с каким элементом в данный моммент работает драйвер нужно указать параметр "--highlight"

Пример:
* py.test tests --highlight

#### Чтобы запустить только упавшие тесты нужно указать параметр "--last-failed"

Пример:
* py.test tests --last-failed

#### Для проверки консольных ошибок в браузере нужно указать параметр "--check-console"

Пример:
* py.test tests --check-console

### В итоге можно запускать тесты с различными комбинациями

Пример:
* py.test tests --firefox --grid -n2 --last-failed --highlight

Запустятся только ранее упавшие тесты в два потока в браузере Opera в гриде. При прогоне тестов активные элементы будут подсвечиваться. Порядок параметров может быть произвольным.

## Структура фреймворка

Название пакета / модуля  | Содержание пакета / модуля
----------------------|----------------------
lib                   | Вспомогательные библиотеки
screenshots           | Место сохранения скриншотов упавших тестов
selenovo              | Фреймворк selenovo
selenovo/elements     | Пакет с элементами
selenovo/files        | Место хранения файлов необходимых для тестов
selenovo/base.py      | Базовый модуль фреймворка
selenovo/common.py    | Общие функции фреймворка
selenovo/exception.py | Исключения фреймворка
selenovo/waits.py     | Функции ожидания фреймворка
state                 | Методы для работы с DB используя SQLSoup
tests                 | Тесты
assertions.py         | Методы assert
config.py             | Файл конфигурации (хост, порт и т.д.)
conftest.py           | Функции pytest
device_config.py      | Конфигурация мобильных устройств (Android, iOS)
helpers.py            | Вспомогательные методы и функции
logger.py             | Функции логирования
pytest.ini            | Файл конфигурации pytest
requirements.txt      | Файл с необходимыми зависимостями
services.py           | Подключение сервисов
test_case.py          | Инициализация драйвера, описания опций браузеров
urls.py               | Методы работы с урлами

## Структура Page

Описание пейджей производится в пакете *"pages"*. Класс пейджы наследуется от базового класса Page.
В пейдже описывается url пейджы, блоки и элементы которые содержит. Если блок содержиться только на этой пейдже то он описывается выше в модуле.
Если блок используется на нескольких пейджах то он описывается в отдельном модуле который находится в пакете *"blocks"* и определяется в модуле "blocks_mixin.py" (пейджа дополнительно наследуется от блока "BlocksMixin").

```
from selenium.webdriver.common.by import By

from promium import Page, Block, Element, Link
from urls import collect_url

from .blocks.blocks_mixin import BlocksMixin


class ProductBlock(Block):
    all_characteristics_link = Link(
        By.CSS_SELECTOR,
        '[data-qaid="all_characteristics_link"]'
    )
    buy_button = Element(
        By.CSS_SELECTOR,
        '[data-qaid="buy_button"]'
    )


class MainPage(Page, BlocksMixin):

    url = collect_url('')

    products_block = ProductBlock.as_list(
        By.CSS_SELECTOR,
        '[data-qaid="searches_block"]'
    )
    header_block = HeaderBlock(
        By.CSS_SELECTOR,
        '[data-qaid="header_block"]'
    )

```

## Примеры тестов с пейджами

Примеры тестов находятся в */tests*

Примеры page object и блоков находятся в */tests/pages* 

### test_catalog_header.py

Тест проверяющий десктопную версию сайта с помошью Selenium Webdriver

### test_mobile_main_page.py

Тест проверяющий мобильную версию сайта с помошью Selenium Webdriver

### test_catalog_main_page.py

Тест который НЕ использует Selenium Webdriver, а использующий библиотеку requests

----------------------------------------------------------------------------------
----------------------------------------------------------------------------------
----------------------------------------------------------------------------------

# promium & Selenium Grid & Vagga (container) in Ubuntu 16.04
## Установка Vagga

http://vagga.readthedocs.io/en/latest/installation.html#ubuntu

Для корректной работы с тестами необходимо в домашней директории /home/user_name создать файл конфигурации .vagga.yaml с таким содержимым:
```
external-volumes:
  bin: /usr/bin
```

## Установка зависимостей, браузеров и драйверов

Ничего устанавливать не нужно, все уже в контейнере :)

## Просмотр доступных комманд Vagga

В корне проекта выполнить команду vagga

* vagga

```
Available commands:
    faq                 View FAQ
    lint                Check code style
    run-tests           Run selenium tests
    selenium-hub        Run Selenium Grid Hub on the local machine
    selenium-node       Run Selenium Grid Node
    versions            View which selenium and browsers versions is using
```

## Команда "vagga lint"
### Для запуска проверки lint необходимо выполнить команду

* vagga lint

## Команда "vagga run-tests"
### Для запуска тестов необходимо выполнить минимальную команду

* vagga run-tests -- <путь к тестам>

Пример:
* vagga run-tests -- tests

### Запуск тестов с параметрами

* vagga run-tests -- tests/test_catalog_main_page.py -k test_main_page_links --firefox --grid -n2 --last-failed --highlight

Тесты запускаются по аналогии с py.test. Все опции указываются после "vagga run-tests -- "

Доступны следующие браузеры:
* --chrome
* --firefox
* --opera

## Команда "vagga selenium-hub" и "selenium-node"
### Запуск Selenium Grid

Запустить две команды параллельно в терминале в следующем порядке:
* vagga selenium-hub

После запуска в консоли получаем такое:
```
13:30:13.934 INFO - Nodes should register to http://0.0.0.0:13400/grid/register/
13:30:13.934 INFO - Selenium Grid hub is up and running
```

* vagga selenium-node

После запуска в консоли получаем такое:
```
13:30:14.351 INFO - Selenium Grid node is up and ready to register to the hub
13:30:14.358 INFO - Starting auto registration thread. Will try to register every 5000 ms.
13:30:14.358 INFO - Registering the node to the hub: http://localhost:13400/grid/register
13:30:14.406 INFO - The node is registered to the hub and ready to use
```

Теперь можно запускать тесты в гриде используя параметр --grid

http://0.0.0.0:13400/grid/console

## Команда "vagga versions"

* vagga versions

```
Google Chrome 63.0.3239.132 
ChromeDriver 2.35.528139 (47ead77cb35ad2a9a83248b292151462a66cd881)
Mozilla Firefox 57.0.4
geckodriver 0.19.1
Opera 50.0.2762.58
OperaDriver 2.32 (cfa164127aab5f93e5e47d9dcf8407380eb42c50)
Selenium server version: 3.8.1, revision: 6e95a6684b
```
# FAQ

* **Если при запуске тестов возникли ошибка вида:**

```
Traceback (most recent call last):
  File "/usr/local/lib/python3.5/dist-packages/_pytest/config.py", line 342, in _getconftestmodules
    return self._path2confmods[path]
KeyError: local("/work/tests")

During handling of the above exception, another exception occurred:
Traceback (most recent call last):
  File "/usr/local/lib/python3.5/dist-packages/_pytest/config.py", line 373, in _importconftest
    return self._conftestpath2mod[conftestpath]
KeyError: local("/work/conftest.py")

During handling of the above exception, another exception occurred:
Traceback (most recent call last):
  File "/usr/local/lib/python3.5/dist-packages/_pytest/config.py", line 379, in _importconftest
    mod = conftestpath.pyimport()
  File "/usr/local/lib/python3.5/dist-packages/py/_path/local.py", line 686, in pyimport
    raise self.ImportMismatchError(modname, modfile, self)
py._path.local.LocalPath.ImportMismatchError: ("work.conftest", "/home/user/workspace/promium3/conftest.py", local("/work/conftest.py"))
ERROR: could not load /work/conftest.py
```
Возникает если запускать тесты сначала через команду py.test, а потом через vagga run-tests или в обратном порядке.
Исправить ошибку можно выполнив в терменале в корне проекта команду:

```
sudo find -name '*.pyc' -delete
```

* **Чтобы удалить собранные контейнеры vagga (например они устарели или занимают уже много места на диске) нужно выполнить в корне проекта следующую команду:**

```
vagga _clean --everything
```

* **Как узнать что находиться в контейнере? Используя слудующую команду можно через терминал посмотреть что же там:**

```
vagga _run promium bash
```
