# (c) Copyright IBM Corp. 2010, 2017. All Rights Reserved.

"""Pytest Fixture for Testing Resilient Circuits Components"""
from __future__ import print_function

import os
import sys
import time
import calendar
import json
import traceback
import logging
import pytest
from pytest_resilient_circuits.circuits_fixtures import manager, watcher
import resilient
from resilient import SimpleHTTPException
import resilient_circuits.app

DATATABLE_TYPE_ID = 8


class ConfiguredAppliance:
    """ configure resilient org with specs from the test module """
    def __init__(self, request):
        # TODO: replace/supplement with the 'resilient.circuits.customize' entry-point
        # TODO: Add support for phases, tasks, incident types.
        # TODO: Add support for optional action and custom fields (currently all set as required)

        host = os.environ.get("TEST_RESILIENT_APPLIANCE", request.config.option.resilient_host)
        url = "https://%s:443" % host
        org = os.environ.get("TEST_RESILIENT_ORG", request.config.option.resilient_org)
        user = os.environ.get("TEST_RESILIENT_USER", request.config.option.resilient_email)
        password = os.environ.get("TEST_RESILIENT_PASSWORD", request.config.option.resilient_password)
        assert all((host, org, user, password))

        # Connect to Resilient
        self.client = resilient.SimpleClient(org_name=org, base_url=url, verify=False)
        session = self.client.connect(user, password)

        # Retrieve constants from appliance
        constants = self._get_constants()["actions_framework"]
        # Invert the dictionaries so they are {"name": id,...}
        action_object_types = {value: int(key) for key, value in constants["action_object_types"].items()}
        action_types = {value: int(key) for key, value in constants["action_types"].items()}
        destination_types = {value: int(key) for key, value in constants["destination_types"].items()}

        # Delete all existing configuration items
        self._clear_org()

        # Create all required configuration items for this class of tests
        action_fields = getattr(request.cls, "action_fields", None)
        if action_fields:
            for field_name, info in action_fields.items():
                success = self._create_action_field(field_name,
                                                    info[0], info[1], info[2])
                assert success
        custom_fields = getattr(request.cls, "custom_fields", None)
        if custom_fields:
            for field_name, info in custom_fields.items():
                success = self._create_custom_field(field_name,
                                                    info[0], info[1], info[2])
                assert success
        destinations = getattr(request.cls, "destinations", None)
        if destinations:
            for queue_name in destinations:
                success = self._create_destination(queue_name,
                                                   destination_types["Queue"])
                assert success

        destinations = self.client.get("/message_destinations")["entities"]
        destinations = {dest["programmatic_name"]: int(dest["id"]) for dest in destinations}
        manual_actions = getattr(request.cls, "manual_actions", None)
        if manual_actions:
            for action_name, info in manual_actions.items():
                success = self._create_manual_action(action_name,
                                                     destinations[info[0]],
                                                     action_object_types[info[1]],
                                                     action_types["Manual"],
                                                     info[2])
                assert success
        automatic_actions = getattr(request.cls, "automatic_actions", None)
        if automatic_actions:
            for action_name, info in automatic_actions.items():
                success = self._create_automatic_action(action_name,
                                                        destinations[info[0]],
                                                        action_object_types[info[1]],
                                                        action_types["Automatic"],
                                                        info[2])
                assert success

        data_tables = getattr(request.cls, "data_tables", None)
        print("create data tables: %s" % data_tables)
        if data_tables:
            for table_name, columns in data_tables.items():
                success = self._create_data_table(table_name, columns)
                assert success

    # end __init__

    def _clear_org(self):
        """ Delete all existing destinations, actions, fields, etc from org """

        actions = self.client.get("/actions")["entities"]
        if actions:
            for action in actions:
                self.client.delete("/actions/%s" % action['id'])

        destinations = self.client.get("/message_destinations")["entities"]
        if destinations:
            for destination in destinations:
                print("/message_destinations/%s" % destination['id'])
                self.client.delete("/message_destinations/%s" % destination['id'])

        fields = self.client.get("/types/actioninvocation/fields")
        if fields:
            for field in fields:
                self.client.delete("/types/actioninvocation/fields/%s" % field['id'])

        fields = self.client.get("/types/incident/fields")
        if fields:
            fields = [field for field in fields if field['prefix'] == "properties"]
            for field in fields:
                print("URL IS /types/incident/fields/%s" % field['id'])
                self.client.delete("/types/incident/fields/%s" % field['id'])

        types = self.client.get("/types")
        data_tables = [res_type["type_name"] for res_type in types.values() if res_type['type_id'] == DATATABLE_TYPE_ID]
        for dt_name in data_tables:
            print("Delete /types/%s" % dt_name)
            self.client.delete("/types/%s" % dt_name)
    # end _clear_org

    def _get_constants(self):
        """ get data from /rest/const """
        url = "{0}/rest/const".format(self.client.base_url)
        response = self.client._execute_request(self.client.session.get,
                                                url,
                                                proxies=self.client.proxies,
                                                cookies=self.client.cookies,
                                                headers=self.client._SimpleClient__make_headers())
        if response.status_code != 200:
            raise SimpleHTTPException(response)
        return json.loads(response.text)
    # end _get_constants

    def _create_action_field(self, programmatic_name, field_type, display_name, values):
        """ Create action field in resilient """
        endpoint = "/types/actioninvocation/fields"
        action_field = {"text": display_name,
                        "required": "always",
                        "blank_option": False,
                        "input_type": field_type,
                        "name": programmatic_name}
        if values:
            action_field["values"] = [{"label": value} for value in values]

        try:
            field_def = self.client.post(endpoint, action_field)
            if not field_def:
                print("Failed to create action field %s" % programmatic_name)
                return False
        except Exception as e:
            print("Failed to create action field %s" % programmatic_name)
            traceback.print_exc()
            return False
        return True
    # end _create_action_field

    def _create_custom_field(self, programmatic_name, field_type, display_name, values):
        """ Create custom field in resilient """
        endpoint = "/types/incident/fields"
        custom_field = {"text": display_name,
                        "input_type": field_type,
                        "name": programmatic_name}
        if values:
            custom_field["values"] = [{"label": value} for value in values]

        response = self.client.post(endpoint, custom_field)
        return True if response else False
    # end _create_custom_field

    def _create_destination(self, name, destination_type):
        """ Create destination queue """
        user_id = self.client.user_id
        endpoint = "/message_destinations"
        destination_obj = {"name": name,
                           "expect_ack": True,
                           "destination_type": destination_type,
                           "programmatic_name": name,
                           "users": [user_id]}
        try:
            destination = self.client.post(endpoint, destination_obj)
            if not destination:
                print("Failed to create destination %s" % name)
                return False
        except Exception as e:
            print("Failed to create destination %s" % name)
            traceback.print_exc()
            return False
        return True
    # end _create_destination

    def _create_manual_action(self, action_name, destination_id, object_type_id, action_type_id, fields):
        """ Create and configure a manual action """
        endpoint = "/actions"
        action = {"name": action_name,
                  "type": action_type_id,
                  "object_type": object_type_id,
                  "message_destinations": [destination_id],
                  "view_items": [{"field_type": "actioninvocation",
                                  "element": "field",
                                  "content": fieldname} for fieldname in fields]
                  }
        try:
            action_obj = self.client.post(endpoint, action)
            if not action_obj:
                print("Failed to create action %s" % action_name)
                return False
        except Exception as e:
            print("Failed to create action %s" % action_name)
            traceback.print_exc()
            return False

        return True
    # end _create_manual_action

    def _create_automatic_action(self, action_name, destination_id, object_type_id, action_type_id, conditions):
        """ Create and configure automatic action """
        endpoint = "/actions"
        action = {"name": action_name,
                  "type": action_type_id,
                  "object_type": object_type_id,
                  "message_destinations": [destination_id]
                  }
        if conditions:
            action["conditions"] = conditions
        try:
            action_obj = self.client.post(endpoint, action)
            if not action_obj:
                print("Failed to create action %s" % action_name)
                return False
        except Exception as e:
            print("Failed to create action %s" % action_name)
            traceback.print_exc()
            return False

        return True
    # end _create_automatic_action

    def _create_data_table(self, table_name, columns):
        """" Create a data table """
        endpoint = "/types"
        table_columns = {}
        print("create table %s" % table_name)
        table = {"type_name": table_name,
                 "display_name": table_name,
                 "type_id": DATATABLE_TYPE_ID,
                 "parent_types": ["incident", ]}
        try:
            dt_obj = self.client.post(endpoint, table)
            if not dt_obj:
                print("Failed to create data table %s" % table_name)
                return False
        except Exception as e:
            print("Failed to create data table %s" % table_name)
            traceback.print_exc()
            return False

        endpoint = "/types/%s/fields" % table_name
        for col_name, (col_type, col_values) in columns.items():
            if not col_values:
                values = []
            else:
                values = [{"label": value} for value in col_values]
            field = {"name": col_name,
                     "text": col_name,
                     "input_type": col_type,
                     "values": values}
            try:
                field_obj = self.client.post(endpoint, field)
                if not field_obj:
                    print("Failed to create data table field %s" % col_name)
                    return False
            except Exception as e:
                print("Failed to create data table field %s" % col_name)
                traceback.print_exc()
                return False

        return True
    # end _create_data_table

# end ConfiguredAppliance


class ResilientCircuits:
    def __init__(self, tmpdir_factory, manager, watcher, request):

        # Reset the rest client b/c it is a global variable
        resilient_circuits.rest_helper.reset_resilient_client()
        resilient_config_data = """
[resilient]
logfile = app.log
loglevel = INFO
cafile = false
stomp_port = 65001
no_prompt_password = True
port = 443
test_actions = True
"""
        print("CURRENT WORKING DIR:  Addr: ", os.getcwd(), id(self))

        resilient_mock = getattr(request.module, "resilient_mock", None)
        self.config_file = tmpdir_factory.mktemp('data').join("%dapp.config" % id(self))
        print("TEST config={}".format(self.config_file.strpath))

        self.logs = tmpdir_factory.mktemp("logs")
        print("TEST logdir={}".format(self.logs))

        config_data = getattr(request.module, "config_data", "")
        self.config_file.write(resilient_config_data)

        host = os.environ.get("TEST_RESILIENT_APPLIANCE")
        if host:
            print("TEST host={} (from environment)".format(host))
        else:
            host = request.config.option.resilient_host
            print("TEST host={} (from configuration data)".format(host))
        self.config_file.write("host = %s\n" % host, mode='a')

        self.org = os.environ.get("TEST_RESILIENT_ORG")
        if self.org:
            print("TEST org={} (from environment)".format(self.org))
        else:
            self.org = request.config.option.resilient_org
            print("TEST org={} (from configuration data)".format(self.org))
        self.config_file.write("org = %s\n" % self.org, mode='a')

        self.user = os.environ.get("TEST_RESILIENT_USER")
        if self.user:
            print("TEST user={} (from environment)".format(self.user))
        else:
            self.user = request.config.option.resilient_email
            print("TEST user={} (from configuration data)".format(self.user))
        self.config_file.write("email = %s\n" % self.user, mode='a')

        password = os.environ.get("TEST_RESILIENT_PASSWORD", request.config.option.resilient_password)
        self.config_file.write("password = %s\n" % password, mode='a')
        self.config_file.write("log_http_responses = %s\n" % self.logs, mode='a')
        self.config_file.write("logdir = %s\n" % self.logs, mode='a')
        self.config_file.write("test_port = 0\n", mode='a')
        if resilient_mock:
            if isinstance(resilient_mock, type):
                self.config_file.write("resilient_mock = %s.%s\n" % (resilient_mock.__module__, resilient_mock.__name__), mode='a')
            else:
                self.config_file.write("resilient_mock = %s\n" % resilient_mock, mode='a')

        self.config_file.write(config_data, mode='a')
        os.environ["APP_CONFIG_FILE"] = self.config_file.strpath
        # Set this manually b/c it is only read on import in app.py
        resilient_circuits.app.APP_CONFIG_FILE = self.config_file.strpath

        self.manager = manager
        self.watcher = watcher

        # Remove the pytest commandline arguments so they don't break ArgParse in resilient
        sys.argv = sys.argv[0:1]

        self.app = resilient_circuits.app.App().register(manager)
        assert watcher.wait("registered")
        pytest.wait_for(manager, "_running", True)
        assert watcher.wait("load_all_success", timeout=10)

    # end __init__

    def finalizer(self):
        """ Unregister and reset CWD back to where it was originally """
        print("Finalizer Running")
        self.flush_logs()
        self.manager.stop()
        pytest.wait_for(self.manager, "_running", False)

    def flush_logs(self):
        handlers = logging.getLogger().handlers
        for handler in handlers:
            if isinstance(handler, logging.FileHandler) or isinstance(handler, logging.handlers.RotatingFileHandler):
                handler.flush()

# end ResilientCircuits


@pytest.fixture(scope="class")
def circuits_app(tmpdir_factory, manager, watcher, request):
    circuits_app_fixture = ResilientCircuits(tmpdir_factory, manager, watcher, request)
    yield circuits_app_fixture
    circuits_app_fixture.finalizer()


@pytest.fixture(scope="class")
def configure_resilient(request):
    """ Create necessary destinations, actions, fields, etc in Resilient """
    yield ConfiguredAppliance(request)
# end configure_resilient


@pytest.fixture(scope="class")
def new_incident(circuits_app):
    """ Create but don't post a minimally populated incident for testing """
    client = circuits_app.app.action_component.rest_client()
    incident = {}
    fields = client.get('/types/incident/fields')
    for field_def in fields:
        if field_def.get('required', '') != 'always':
            continue
        fieldname = field_def['name']
        field_type = field_def['input_type']

        if field_def['values']:
            # Just use whatever the first valid value is
            value = field_def['values'][0]["value"]
        elif field_type == "boolean":
            value = True
        elif field_type in ("text", "textarea"):
            value = "test"
        elif field_type == "datetimepicker":
            value = int(calendar.timegm(time.gmtime())) * 1000
        elif field_type == "number":
            value = 1

        if field_def['prefix']:
            incident[field_def['prefix']][fieldname] = value
        else:
            incident[fieldname] = value
    yield incident
# end new_incident
