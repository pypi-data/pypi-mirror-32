import requests
import json
from urllib.request import urlopen

from django.conf import settings
from django.core.files.base import ContentFile

from core.models import User

class ElggBackend:

    def authenticate(self, request, username=None, password=None):
        if not settings.ELGG_URL:
            return None

        elgg_url = settings.ELGG_URL

        # Check if user exists (case-insensitive)
        try:
            user = User.objects.get(email__iexact=username)
            if user.check_password(password):
                return user
        except User.DoesNotExist:
            # Verify username/password combination
            valid_user_request = requests.post(elgg_url + "/services/api/rest/json/", data={'method': 'pleio.verifyuser', 'user': username, 'password': password})
            valid_user_json = json.loads(valid_user_request.text)
            valid_user_result = valid_user_json["result"] if 'result' in valid_user_json else []
            valid_user = valid_user_result["valid"] if 'valid' in valid_user_result else False
            name = valid_user_result["name"] if 'name' in valid_user_result else username
            admin = valid_user_result["admin"] if 'admin' in valid_user_result else False

            # If valid, create new user with Elgg attributes
            if valid_user is True:
                user = User.objects.create_user(
                    name=name,
                    email=username,
                    password=password,
                    accepted_terms=True,
                    receives_newsletter=True
                )
                user.is_active = True
                user.is_admin = admin
                user.save()
                return user
            else:
                return None

    def get_user(self, user_id):
        try:
            return User.objects.get(pk=user_id)
        except User.DoesNotExist:
            return None
