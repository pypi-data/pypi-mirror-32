
django-sssoon
=============

Django-sssoon is a simple Django app to add beautiful coming soon webpage to your django website. This template is
based on on Bootstrap 3 and designed by [Creative Tim](https://www.creative-tim.com/).

![Screenshot](./docs/images/screencapture.png "Screenshot")

Detailed documentation is in the "docs" directory.

Quick start
-----------
1. django-sssoon is available on the Python Package Index (PyPI), so it can be installed with standard Python tools like `pip` or `easy_install`:

```python
pip install django-sssoon
```

2. Add "sssoon" to your INSTALLED_APPS setting like this:

```python
INSTALLED_APPS = [
    ...
    'sssoon',
]
```

2. Include the sssoon URLconf in your project urls.py like this to make your index page coming sssoon:

```python
url(r'^', include('sssoon.urls', namespace="sssoon")),
```

3. Collect static files

```python
python manage.py collectstatic
```

4. Start the development server and visit http://127.0.0.1:8000/
