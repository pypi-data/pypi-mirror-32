# pylint: disable=no-member,redefined-outer-name,unused-argument
from collections import OrderedDict

from django.conf.urls import include, url

from rest_framework import routers, permissions

from .base import Response
from ..utils import import_class


class _AbstractRouter(routers.DefaultRouter):

    def __init__(self, *args, **kwargs):
        self.custom_urls = list()
        self.permission_classes = kwargs.pop("perms", None)
        super(_AbstractRouter, self).__init__(*args, **kwargs)

    def get_default_base_name(self, viewset):
        base_name = getattr(viewset, 'base_name', None)
        if base_name is not None:
            return base_name  # nocv
        queryset = getattr(viewset, 'queryset', None)
        model = getattr(viewset, 'model', None)
        if queryset is None:
            assert model is not None, \
                '`base_name` argument not specified, and could ' \
                'not automatically determine the name from the viewset, as ' \
                'it does not have a `.queryset` or `.model` attribute.'
            return model._meta.object_name.lower()
        # can't be tested because this initialization takes place before any
        # test code can be run
        return super(_AbstractRouter,
                     self).get_default_base_name(viewset)  # nocv

    def register_view(self, prefix, view, name=None):
        name = name or view().get_view_name().lower()
        self.custom_urls.append((prefix, view, name))

    def _unreg(self, prefix, objects_list):
        del self._urls
        index = 0
        for reg_prefix, _, _ in objects_list:
            if reg_prefix == prefix:
                del objects_list[index]
                break
            index += 1  # nocv
        return objects_list

    def unregister_view(self, prefix):
        self.custom_urls = self._unreg(prefix, self.custom_urls)  # nocv

    def unregister(self, prefix):
        self.registry = self._unreg(prefix, self.registry)

    def generate(self, views_list):
        for prefix, options in views_list.items():
            args = [prefix, import_class(options['view']), options.get('name', None)]
            if options.get('type', 'viewset') == 'viewset':
                self.register(*args)
            elif options.get('type', 'viewset') == 'view':
                self.register_view(*args)
            else:  # nocv
                raise Exception('Unknown type of view')


class APIRouter(_AbstractRouter):
    root_view_name = 'api-v1'

    def get_api_root_view(self, *args, **kwargs):
        api_root_dict = OrderedDict()
        list_name = self.routes[0].name
        for prefix, _, basename in self.registry:
            api_root_dict[prefix] = list_name.format(basename=basename)

        class API(self.APIRootView):
            root_view_name = self.root_view_name
            if self.permission_classes:
                permission_classes = self.permission_classes
            custom_urls = self.custom_urls

            def get_view_name(self): return self.root_view_name

            def get(self, request, *args, **kwargs):
                registers = super(API, self).get(request, *args, **kwargs)
                routers_list = OrderedDict()
                for prefix, _, name in self.custom_urls:  # nocv
                    fpath = request.get_full_path().split("?")
                    path = request.build_absolute_uri(fpath[0]) + prefix
                    if len(fpath) > 1:
                        path += "?{}".format(fpath[1])
                    routers_list[name] = path
                routers_list.update(registers.data)
                return Response(routers_list, 200).resp

        return API.as_view(api_root_dict=api_root_dict)

    def get_urls(self):
        urls = super(APIRouter, self).get_urls()
        for prefix, view, _ in self.custom_urls:  # nocv
            urls.append(url("^{}/$".format(prefix), view.as_view()))
        return urls


class MainRouter(_AbstractRouter):
    routers = []

    def get_api_root_view(self, *args, **kwargs):
        api_root_dict = OrderedDict()
        list_name = self.routes[0].name
        for prefix, _, basename in self.registry:
            api_root_dict[prefix] = list_name.format(basename=basename)  # nocv

        class API(self.APIRootView):
            if self.permission_classes:
                permission_classes = self.permission_classes
            routers = self.routers
            custom_urls = self.custom_urls

            def get_view_name(self): return "API"

            def get(self, request, *args, **kwargs):
                registers = super(API, self).get(request, *args, **kwargs)
                routers_list = OrderedDict()
                for prefix, _, name in self.routers + self.custom_urls:
                    fpath = request.get_full_path().split("?")
                    path = request.build_absolute_uri(fpath[0]) + prefix
                    if len(fpath) > 1:
                        path += "?{}".format(fpath[1])  # nocv
                    routers_list[name] = path
                routers_list.update(registers.data)
                return Response(routers_list, 200).resp

        return API.as_view(api_root_dict=api_root_dict)

    def register_router(self, prefix, router, name=None):
        name = name or router.root_view_name
        self.routers.append((prefix, router, name))

    def unregister_router(self, prefix):
        self.routers = self._unreg(prefix, self.routers)  # nocv

    def get_urls(self):
        urls = super(MainRouter, self).get_urls()
        for prefix, router, _ in self.routers:
            urls.append(url(prefix, include(router.urls)))
        for prefix, view, _ in self.custom_urls:
            # can't be tested because this initialization takes place before
            # any test code can be run
            urls.append(url(prefix, view.as_view()))  # nocv
        return urls

    def generate_routers(self, api):
        for version, views_list in api.items():
            router = APIRouter(perms=(permissions.IsAuthenticated,))
            router.root_view_name = version
            router.generate(views_list)
            self.register_router(version+'/', router)
