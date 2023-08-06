from django.db import models
from django.contrib import admin
from django.utils.text import slugify
from django.shortcuts import get_object_or_404

from .dash_wrapper import get_stateless_by_name

import json

class DashApp(models.Model):
    '''
    An instance of this model represents a dash application and its internal state
    '''
    app_name = models.CharField(max_length=100, blank=False, null=False, unique=False)
    instance_name = models.CharField(max_length=100, unique=True, blank=True, null=False)
    slug = models.SlugField(max_length=110, unique=True, blank=True)
    base_state = models.TextField(null=False, default="{}") # If mandating postgresql then this could be a JSONField
    creation = models.DateTimeField(auto_now_add=True)
    update = models.DateTimeField(auto_now=True)
    save_on_change = models.BooleanField(null=False,default=False)

    def __str__(self):
        return self.instance_name

    def save(self, *args, **kwargs):
        if not self.instance_name:
            existing_count = DashApp.objects.filter(app_name=self.app_name).count()
            self.instance_name = "%s-%i" %(self.app_name, existing_count+1)
        if not self.slug or len(self.slug) < 2:
            self.slug = slugify(self.instance_name)
        super(DashApp, self).save(*args,**kwargs)

    def _stateless_dash_app(self):
        dd = getattr(self,'_stateless_dash_app_instance',None)
        if not dd:
            dd = get_stateless_by_name(self.app_name)
            setattr(self,'_stateless_dash_app_instance',dd)
        return dd

    def handle_current_state(self):
        '''
        Check to see if the current hydrated state and the saved state are different.

        If they are, then persist the current state in the database by saving the model instance.
        '''
        if getattr(self,'_current_state_hydrated_changed',False) and self.save_on_change:
            new_base_state = json.dumps(getattr(self,'_current_state_hydrated',{}))
            if new_base_state != self.base_state:
                self.base_state = new_base_state
                self.save()

    def have_current_state_entry(self, wid, key):
        cscoll = self.current_state()
        cs = cscoll.get(wid,{})
        return key in cs

    def update_current_state(self, wid, key, value):
        '''
        Update current state with a (possibly new) value associated with key

        If the key does not represent an existing entry, then ignore it
        '''
        cscoll = self.current_state()
        cs = cscoll.get(wid,{})
        if key in cs:
            current_value = cs.get(key,None)
            if current_value != value:
                cs[key] = value
                setattr(self,'_current_state_hydrated_changed',True)

    def current_state(self):
        '''
        Return the current internal state of the model instance.

        This is not necessarily the same as the persisted state stored in the self.base_state variable.
        '''
        cs = getattr(self,'_current_state_hydrated',None)
        if not cs:
            cs = json.loads(self.base_state)
            setattr(self,'_current_state_hydrated',cs)
            setattr(self,'_current_state_hydrated_changed',False)
        return cs

    def as_dash_instance(self):
        dd = self._stateless_dash_app()
        base = self.current_state()
        return dd.form_dash_instance(replacements=base,
                                     specific_identifier=self.slug)

    def _get_base_state(self):
        '''
        Get the base state of the object, as defined by the app.layout code, as a python dict
        '''
        base_app_inst = self._stateless_dash_app().as_dash_instance()

        # Get base layout response, from a base object
        base_resp = base_app_inst.locate_endpoint_function('dash-layout')()

        base_obj = json.loads(base_resp.data.decode('utf-8'))

        # Walk the base layout and find all values; insert into base state map
        obj = {}
        base_app_inst.walk_tree_and_extract(base_obj, obj)
        return obj

    def populate_values(self):
        '''
        Add values from the underlying dash layout configuration
        '''
        obj = self._get_base_state()
        self.base_state = json.dumps(obj)

    @staticmethod
    def locate_item(id, stateless=False):
        if stateless:
            da = get_stateless_by_name(id)
        else:
            da = get_object_or_404(DashApp,slug=id)

        app = da.as_dash_instance()
        return da, app

class DashAppAdmin(admin.ModelAdmin):
    list_display = ['instance_name','app_name','slug','creation','update','save_on_change',]
    list_filter = ['creation','update','save_on_change','app_name',]

    def _populate_values(self, request, queryset):
        for da in queryset:
            da.populate_values()
            da.save()
    _populate_values.short_description = "Populate app"

    actions = ['_populate_values',]
