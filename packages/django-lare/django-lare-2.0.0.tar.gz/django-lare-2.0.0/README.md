# django-lare [![Build Status](https://secure.travis-ci.org/lare-team/django-lare.png)](http://travis-ci.org/lare-team/django-lare)

## How to install django-lare?

There are just two steps needed to install django-lare:

1. Install django-lare to your virtual env:

	```bash
	pip install django-lare==1.0.1
	```

2. Configure your django installation with the following lines:

	```python
    # django lare
    INSTALLED_APPS += ('django_lare', )

    MIDDLEWARE_CLASSES += (
        'django_lare.middlewares.LareMiddleware',
    )
    if django.VERSION <= (1, 9):
        from django.template.base import add_to_builtins
        add_to_builtins('django_lare.templatetags.lare_extends',)
        TEMPLATE_CONTEXT_PROCESSORS = ('django_lare.context_processors.lare_information',)
    else:
        # for your templates:

        TEMPLATES = [{
            ...
            'OPTIONS': {
                ...
                'context_processors': [
                    ...
                    'django_lare.context_processors.lare_information',
                ],
                'builtins': [
                    ...
                    'django_lare.templatetags.lare_extends',
                ],
            },
        }]

    DEFAULT_LARE_TEMPLATE = "django_lare/lare.html"
	```

3. Set your namespaces:

    In your views:

	```python
    class YourView(DefaultLareViewMixin, View):
        lare_current_namespace = "YourSite.YourView"
        ...
	```

	In your templates:

	```html
	...
    <body data-lare-namespace="{% block lare_namespace %}{{ lare.current_namespace }}{% endblock lare_namespace %}">
        ...
    </body>
	```

## What do you need for django-lare?

1. Django >= 1.4
2. [lare.js](https://github.com/lare-team/lare.js)

## Projects using django-lare

1. [socialfunders.org](https://socialfunders.org/)
2. [iekadou.com](http://www.iekadou.com/)

If you are using django-lare, please contact us, and tell me in which projects you are using it. Thank you!

Happy speeding up your django project!

For further information read [django-lare on iekadou.com](http://www.iekadou.com/programming/django-lare)
