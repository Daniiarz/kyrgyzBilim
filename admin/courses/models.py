from django.db import models


class Course(models.Model):
    name = models.CharField(max_length=255)
    description = models.CharField(max_length=255)

    class Meta:
        db_table = 'courses'
        managed = False
    
    def __str__(self):
        return self.name


class Section(models.Model):
    course = models.ForeignKey(Course, on_delete=models.CASCADE, related_name='sections')
    name = models.CharField(max_length=255)
    icon = models.ImageField(upload_to='section/', blank=True, null=True)
    type = models.CharField(max_length=255, blank=True, null=True)

    def save(self, *args, **kwargs):
        self.type = self.name
        super(Section, self).save(*args, **kwargs)

    class Meta:
        db_table = 'sections'
        managed = False

    def __str__(self):
        return f"{self.name} - {self.course.name}"


class Topic(models.Model):
    section = models.ForeignKey(Section, on_delete=models.CASCADE, related_name='topics')
    name = models.CharField(max_length=255)
    translated_name = models.CharField(max_length=255, blank=True, null=True)
    icon = models.FileField(upload_to='section/', blank=True, null=True)
    type = models.CharField(max_length=255, blank=True, null=True)

    class Meta:
        db_table = 'topics'
        managed = False

    def __str__(self):
        return self.name


class SubTopic(models.Model):
    topic = models.ForeignKey(Topic, on_delete=models.CASCADE, related_name='subtopics')
    text = models.TextField()
    translated_text = models.TextField(blank=True, null=True)
    audio = models.FileField(upload_to='audio/', blank=True, null=True)
    image = models.FileField(upload_to='subtopic/', blank=True, null=True)

    order = models.PositiveIntegerField(default=0, blank=False, null=False)

    class Meta:
        db_table = 'sub_topics'
        managed = False
        ordering = ['order']

    def __str__(self):
        return f'{self.text} - {self.topic.name}'
