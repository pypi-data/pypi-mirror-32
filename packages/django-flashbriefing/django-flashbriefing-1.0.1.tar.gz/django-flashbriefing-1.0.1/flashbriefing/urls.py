from django.conf.urls import url

from .views import FeedView

app_name = 'flashbriefing'
urlpatterns = [
    url(r'^feeds/(?P<uuid>\w{32}).(?P<format>\w+)$',
        FeedView.as_view(), name='feed'),
]
