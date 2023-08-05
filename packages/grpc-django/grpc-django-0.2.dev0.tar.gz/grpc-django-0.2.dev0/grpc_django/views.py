import json
import traceback

from django.contrib.auth.models import AnonymousUser
from django.db.models import QuerySet

from grpc_django.models import ContextUser
from grpc_django.protobuf_to_dict import dict_to_protobuf
from grpc_django.settings import settings
from grpc_django.exceptions import InvalidArgument, NotAuthenticated, ExceptionHandler


class GenericGrpcView(object):
    queryset = None
    serializer_class = None
    # Protocol Buffer class which needs to be returned
    response_proto = None
    # Whether the RPC call requires authentication
    requires_authentication = False
    # Lookup field identifier in request proto
    lookup_kwarg = "id"
    # If you want to use object lookups other than pk, set 'lookup_field'.
    # For more complex lookup requirements override `get_object()`.
    lookup_field = "pk"

    def __init__(self, request, context):
        assert self.response_proto, "Missing response_proto declaration"
        self.request_user = self.get_user(context)
        self.request = request
        self.context = context

    @staticmethod
    def get_user(context):
        user_json = json.loads(dict(context.invocation_metadata()).get(settings.auth_user_metakey, "{}"))
        if user_json is None or len(user_json) == 0:
            return AnonymousUser()
        else:
            return ContextUser(**user_json)

    def get_queryset(self):
        assert self.queryset is not None, (
            "{}' should either include a `queryset` attribute, "
            "or override the `get_queryset()` method.".format(self.__class__.__name__)
        )
        queryset = self.queryset
        if isinstance(queryset, QuerySet):
            # Ensure queryset is re-evaluated on each request.
            queryset = queryset.all()
        return queryset

    def check_object_permissions(self, user, obj):
        """
        Override this function to check if the request should be permitted for a given object.
        Raise an appropriate exception if the request is not permitted.
        """
        pass

    def perform_authentication(self, user):
        """
        Perform authentication on the incoming request.
        Raise an appropriate exception if the request is not authenticated.
        :param user:
        :return:
        """
        if self.requires_authentication and isinstance(user, AnonymousUser):
            raise NotAuthenticated
        pass

    def get_object(self):
        if not hasattr(self.request, self.lookup_kwarg):
            raise InvalidArgument("Missing argument {}".format(self.lookup_kwarg))
        queryset = self.get_queryset()
        obj = queryset.get(**{self.lookup_field: getattr(self.request, self.lookup_kwarg)})
        self.check_object_permissions(self.request_user, obj)
        return obj


class RetrieveGRPCView(GenericGrpcView):
    def retrieve(self, request):
        """
        Override this function to implement retrieval
        :param request:
        :return: dictionary of object
        """
        instance = self.get_object()
        serializer = self.serializer_class(instance)
        return serializer.data

    def __call__(self):
        try:
            self.perform_authentication(self.request_user)
            result = self.retrieve(self.request)
            return dict_to_protobuf(self.response_proto, values=result, ignore_none=True)
        except Exception as ex:
            self.context = ExceptionHandler(self.context).__call__(ex, traceback.format_exc())
            return self.response_proto()
