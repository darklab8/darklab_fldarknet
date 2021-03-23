# Generated by Django 3.1.7 on 2021-03-23 16:01

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('ship', '0005_auto_20210323_2259'),
    ]

    operations = [
        migrations.AlterField(
            model_name='ship',
            name='camera_horizontal_turn_angle',
            field=models.IntegerField(blank=True, null=True, verbose_name='hor ang'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='camera_turn_look_ahead_slerp_amount',
            field=models.FloatField(blank=True, null=True, verbose_name='look ahead'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='camera_vertical_turn_down_angle',
            field=models.IntegerField(blank=True, null=True, verbose_name='turn down'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='camera_vertical_turn_up_angle',
            field=models.IntegerField(blank=True, null=True, verbose_name='turn up'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='explosion_resistance',
            field=models.FloatField(blank=True, null=True, verbose_name='exp res'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='nanobot_limit',
            field=models.IntegerField(blank=True, null=True, verbose_name='nanobots'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='shield_battery_limit',
            field=models.IntegerField(blank=True, null=True, verbose_name='batteries'),
        ),
        migrations.AlterField(
            model_name='ship',
            name='strafe_power_usage',
            field=models.IntegerField(blank=True, null=True, verbose_name='strafe usage'),
        ),
    ]
