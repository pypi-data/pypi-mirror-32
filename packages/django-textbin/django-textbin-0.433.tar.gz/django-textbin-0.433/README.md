# Textbin

Python приложение для Django представляющее собой web-сервис для размещения 
текстов и обмена ими.
* Длина текста не ограничена.
* Возможно прикрепление ссылок на изображения и видео
* Приложение предоставляет WEB и REST API интерфейсы.


## Reqiurements

* Python 2.X or 3.X
* Django 1.9.1+
* Django REST framework 3.3.2+
* Django-bootstrap3 6.2.2+
* Django-recaptcha2 0.1.7+

## Installation

1) Установить приложение и его зависимости одним из способов:
```bash
$ pip install django-textbin
$ sudo pip install -e git+https://bitbucket.org/zokalo/django-textbin.git#egg=django-textbin
$ sudo pip install -e git+https://github.com/zokalo/django-textbin.git#egg=django-textbin
$ sudo python setup.py install
```
Eсли у Вас не установлен setuptools, то setup.py uses использует distutils 
который не поддерживает установку зависимостей, установите зависимости из 
requirements.txt самостоятельно:
```bash
pip install -r requirements.txt
```

2) Textbin может использоваться как самостоятельный проект или как приложение.

**Для использования в качестве самостоятельного проекта**
используйте встроенный settings.py и передайте для него настройки в переменных 
окружения:

* TEXTBIN_DEBUG
* TEXTBIN_SECRET_KEY
* TEXTBIN_DATABASE_ENGINE
* TEXTBIN_DATABASE_NAME
* TEXTBIN_DATABASE_USER
* TEXTBIN_DATABASE_PASSWORD
* TEXTBIN_DATABASE_HOST
* TEXTBIN_RECAPTCHA_PUBLIC_KEY
* TEXTBIN_RECAPTCHA_PRIVATE_KEY
* TEXTBIN_STATIC_ROOT
* TEXTBIN_STATIC_URL
* TEXTBIN_ALLOWED_HOSTS
* TEXTBIN_DJANGO_LOGGING

**Для использования в качестве приложения к проекту:**
Добавить в settings.py приложения в INSTALLED_APPS
    'textbin',
    'rest_framework',
    'bootstrap3',
    'snowpenguin.django.recaptcha2',

Установить параметры для приложения и его компонентов по аналогии с settings.py 
приложения.
Установить в urls.py проекта URL приложения 
Кроме того, должны быть подключены staticfiles
```python
from django.contrib.staticfiles.urls import staticfiles_urlpatterns

urlpatterns += [url(r'^textbin/', include('textbin.urls')),]
urlpatterns += staticfiles_urlpatterns()
```

4) Добавить модели приложения в базу данных проекта
```bash
python manage.py migrate
```

## History
* **v0.432 (2018-06-08)**
    * Mark text_detail page with canonical (required when redirected)
* **v0.431 (2016-02-20)**
    * Добавлены настройки логгирования
* **v0.420 (2016-02-16)**
    * Доработка административного интерфейса: фильтры, поиск, отображение даты публикации
    * Кнопка `submit` формы создания нового сообщения перенесена вправо
* **v0.411 (2016-02-12)**
    * Добавлена возможность явного указания типа контента по ссылке при создании записи через API (int-идентификатором или строкой с именем типа)
    * Добавлено требование авторизации для использования API
* **v0.356 (2016-02-09)**
    * Добавлено распознавание ссылок на YouTube
    * Добавлено указание параметров textbin в settings.py:
        - длина id присваемых сообщениям
        - максимальное количество прикрепляемых к сообщению ссылок
        - разрешить прикрепление ссылок, контент которых не распознается
    * Изменён внешний вид форм: верстка, тема, шрифты, галерея изображений
    * Добавлен тест api
    * Добавлена локализация:
        - русский
        - английский
    * Добавлено распознавание прикрепляемых изображений по спецификатору формата, завершающему ссылку
* **v0.344 (2016-02-05)**
    * Удалён повторный импорт статики bootstrap`а
    * Исправлен метод определения значения DEBUG из переменной окружения TEXTBIN_DEBUG
    * Добавлен  импорт ALLOWED_HOSTS из переменной окружения TEXTBIN_ALLOWED_HOSTS
* **v0.33 (2016-02-04)**
    * Импорт настроек в settings.py из переменных окружения
    * Реорганизация urls.py
* **v0.22 (2016-02-02)**
    * Для оформления форм применён фреймворк Django-bootstrap3
    * Добавлена капча django-recaptcha2
* **v0.21 (2016-02-01)**
    * изменено оформление (bootstrap)
    * предпросмотр изображений с кнопками перелистывания 
    * запрещены пустые сообщения: должен содержаться текст или прикрепленные ссылки
    * добавлена внутренняя страница "404" при запросе несуществующего поста
    * введена проверка типа контента, на который указывает прикрепляемый url
    * для основного функционала добавлены unitest`ы
* **v0.2 (2016-01-27)**
    * Переработаны модели данных
    * Переработаны представления
    * Добавлена возможность прикрепления видео
    * REST API реализован Django REST Framework
* **v0.1 (2016-01-13)**
    * Реализован базовый функционал:
        - опубликование текста
        - возможность указания автора, заголовка
        - возможность добавления url-изображений
        - web интерфейс
        - REST API для создания новых и получения существующих записей 
          (реализован Tastypie)
        - django-admin интерфейс

## License

GPL v3.0 License. See the LICENSE file.