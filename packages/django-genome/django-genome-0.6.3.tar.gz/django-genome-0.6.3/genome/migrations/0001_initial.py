# -*- coding: utf-8 -*-
# Generated by Django 1.11.8 on 2018-01-09 13:07
from __future__ import unicode_literals

from django.db import migrations, models
import django.db.models.deletion
import django.utils.timezone
import model_utils.fields


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Chromosome',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('label', models.CharField(max_length=2)),
                ('length', models.PositiveIntegerField()),
                ('active', models.BooleanField(default=True)),
            ],
            options={
                'verbose_name': 'Chromosome',
                'verbose_name_plural': 'Chromosomes',
            },
        ),
        migrations.CreateModel(
            name='CytoBand',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('label', models.CharField(max_length=20)),
                ('start', models.PositiveIntegerField()),
                ('end', models.PositiveIntegerField()),
                ('stain', models.CharField(max_length=10)),
                ('active', models.BooleanField(default=True)),
                ('chromosome', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='cytobands', to='genome.Chromosome')),
            ],
            options={
                'verbose_name': 'Cytogenetic Band',
                'verbose_name_plural': 'Cytogenetic Bands',
            },
        ),
        migrations.CreateModel(
            name='Exon',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('number', models.PositiveSmallIntegerField()),
                ('active', models.BooleanField(default=True)),
                ('start', models.PositiveIntegerField()),
                ('end', models.PositiveIntegerField()),
                ('cds', models.BooleanField(default=False)),
            ],
            options={
                'verbose_name': 'Exon',
                'verbose_name_plural': 'Exons',
            },
        ),
        migrations.CreateModel(
            name='Gene',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('symbol', models.CharField(db_index=True, max_length=255, unique=True)),
                ('name', models.TextField(blank=True)),
                ('hgnc_id', models.PositiveIntegerField(blank=True, null=True)),
                ('status', models.PositiveSmallIntegerField(choices=[(1, 'approved'), (2, 'entry_withdrawn'), (3, 'symbol_withdrawn'), (4, 'ucsc_gene')])),
                ('active', models.BooleanField(default=True)),
                ('previous_name', models.TextField(blank=True)),
                ('locus_type', models.CharField(blank=True, max_length=100)),
                ('locus_group', models.CharField(blank=True, max_length=100)),
                ('ensembl', models.TextField(blank=True)),
                ('refseq', models.TextField(blank=True)),
                ('not_curated_ensembl', models.TextField(blank=True, verbose_name='Ensembl ID (supplied by Ensembl)')),
                ('not_curated_refseq', models.TextField(blank=True, verbose_name='RefSeq (supplied by NCBI)')),
                ('not_curated_ucsc', models.TextField(blank=True, verbose_name='UCSC ID (supplied by UCSC)')),
                ('not_curated_omim', models.TextField(blank=True, verbose_name='OMIM ID (supplied by OMIM)')),
                ('not_curated_uniprot', models.TextField(blank=True, verbose_name='UniProt ID (supplied by UniProt)')),
                ('not_curated_mouse_genome_database', models.TextField(blank=True, verbose_name='Mouse Genome Database ID (supplied by MGI)')),
                ('not_curated_rat_genome_database', models.TextField(blank=True, verbose_name='Rat Genome Database ID (supplied by RGD)')),
                ('chromosome', models.ForeignKey(blank=True, null=True, on_delete=django.db.models.deletion.SET_NULL, related_name='genes', to='genome.Chromosome')),
            ],
            options={
                'verbose_name': 'Gene',
                'verbose_name_plural': 'Genes',
            },
        ),
        migrations.CreateModel(
            name='GeneSynonym',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('label', models.CharField(db_index=True, max_length=255, unique=True)),
                ('active', models.BooleanField(default=True)),
            ],
            options={
                'verbose_name': 'Gene Synonym',
                'verbose_name_plural': 'Gene Synonyms',
            },
        ),
        migrations.CreateModel(
            name='Genome',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('label', models.CharField(max_length=50, unique=True)),
                ('description', models.CharField(blank=True, max_length=255)),
                ('slug', models.SlugField(unique=True)),
                ('active', models.BooleanField(default=True)),
                ('description_url', models.URLField(blank=True)),
            ],
            options={
                'verbose_name': 'Genome Build',
                'verbose_name_plural': 'Genome Builds',
            },
        ),
        migrations.CreateModel(
            name='Transcript',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('created', model_utils.fields.AutoCreatedField(default=django.utils.timezone.now, editable=False, verbose_name='created')),
                ('modified', model_utils.fields.AutoLastModifiedField(default=django.utils.timezone.now, editable=False, verbose_name='modified')),
                ('label', models.CharField(db_index=True, max_length=255, unique=True)),
                ('active', models.BooleanField(default=True)),
                ('strand', models.PositiveSmallIntegerField(choices=[(1, '+'), (2, '-')])),
                ('transcription_start', models.PositiveIntegerField()),
                ('transcription_end', models.PositiveIntegerField()),
                ('cds_start', models.PositiveIntegerField()),
                ('cds_end', models.PositiveIntegerField()),
                ('gene', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='transcripts', to='genome.Gene')),
            ],
            options={
                'verbose_name': 'Transcript',
                'verbose_name_plural': 'Transcripts',
            },
        ),
        migrations.AddField(
            model_name='gene',
            name='synonyms',
            field=models.ManyToManyField(blank=True, related_name='genes', to='genome.GeneSynonym'),
        ),
        migrations.AddField(
            model_name='exon',
            name='transcript',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='exons', to='genome.Transcript'),
        ),
        migrations.AddField(
            model_name='chromosome',
            name='genome',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='chromosomes', to='genome.Genome'),
        ),
        migrations.AlterUniqueTogether(
            name='chromosome',
            unique_together=set([('genome', 'label')]),
        ),
    ]
