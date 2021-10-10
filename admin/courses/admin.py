from adminsortable2.admin import SortableAdminMixin
from django.contrib import admin

from . import models

# Register your models here.

@admin.register(models.SubTopic)
class SubtopicAdmin(SortableAdminMixin, admin.ModelAdmin):
    list_display = ('text', 'topic_name', 'audio', 'image')
    search_fields = ('text', )
    list_filter = ('topic',)
    list_editable = ('audio', 'image')

    def topic_name(self, obj):
        return obj.topic.name


admin.site.register(models.Course)
admin.site.register(models.Section)
admin.site.register(models.Topic)
