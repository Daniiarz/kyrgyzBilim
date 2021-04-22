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
    icon = models.ImageField(upload_to='section/')

    class Meta:
        db_table = 'sections'
        managed = False

    def __str__(self):
        return self.name


class Topic(models.Model):
    section = models.ForeignKey(Section, on_delete=models.CASCADE, related_name='topics')
    name = models.CharField(max_length=255)
    translated_name = models.CharField(max_length=255)
    icon = models.FileField(upload_to='section/')
    type = models.CharField(max_length=255)

    class Meta:
        db_table = 'topics'
        managed = False

    def __str__(self):
        return self.name


class SubTopic(models.Model):
    topic = models.ForeignKey(Topic, on_delete=models.CASCADE, related_name='subtopics')
    text = models.TextField()
    translated_text = models.TextField()
    audio = models.FileField('audio/')
    image = models.FileField('subtopic/')

    class Meta:
        db_table = 'sub_topics'
        managed = False

    def __str__(self):
        return self.text
