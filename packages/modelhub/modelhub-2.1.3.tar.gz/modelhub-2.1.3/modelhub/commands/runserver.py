# -*- coding: utf-8 -*-
from __future__ import print_function
import click
import os
from .base import BaseCommand, argument, option, types, register  # noqa
from modelhub.core.models import Model
from modelhub.core.utils import cached_property
from collections import defaultdict

# TODO: move to config
REDIS_URL_ENV_KEY = "RD_REDIS_URL"
REDIS_PREFIX_ENV_KEY = "RD_REDIS_PREFIX"
REDIS_TIMEOUT = 10


@register("runserver")
class Command(BaseCommand):
    arguments = [
        argument("models_name", nargs=-1),
        option("-b", "--bin_path", type=click.Path(exists=True)),
        option("-r", "--register_model", is_flag=True)
    ]

    def __init__(self, *args, **kwargs):
        super(Command, self).__init__(*args, **kwargs)
        self.redis_url = os.getenv(REDIS_URL_ENV_KEY)
        self.redis_prefix = os.getenv(REDIS_PREFIX_ENV_KEY, "")
        # bin_path = "/usr/bin/tensorflow_model_server"

    @cached_property
    def bin_path(self):
        import subprocess
        return subprocess.check_output(["which", "tensorflow_model_server"]).strip().decode()

    def run(self, models_name, bin_path=None, register_model=False):
        """run tensorflow_model_server with models"""
        if not models_name:
            self.echo("no model name provided")
            return

        self.is_register_model = register_model

        if bin_path:
            self.bin_path = bin_path

        model_name_version_list = [Model.split_name_version(model_name) for model_name in models_name]
        model_versions_map = defaultdict(list)
        for model_name, version in model_name_version_list:
            model_versions_map[model_name].append(version)

        model_versions_list = [(Model.get_local(model_name), list(filter(None, versions))) for model_name, versions in model_versions_map.items()]

        from tempfile import NamedTemporaryFile
        with NamedTemporaryFile("r+") as f:
            f.write(self.generate_model_config(model_versions_list))
            f.flush()
            self.echo("server bin_path:\t%s" % self.bin_path)
            self.echo("model_config_file:\t%s" % f.name)
            self._run_server(
                models=model_versions_list,
                model_config_file=f.name,
            )

    def generate_model_config(self, model_versions_list):
        from tensorflow_serving.config.model_server_config_pb2 import ModelServerConfig, ModelConfigList, ModelConfig
        from tensorflow_serving.sources.storage_path.file_system_storage_path_source_pb2 import FileSystemStoragePathSourceConfig

        def build_model_config(model, versions):
            if not versions:
                versions_list = [model.latest_local_version]
            else:
                versions_list = []
                for version in versions:
                    version = model.get_version(version)
                    assert version.local_exists, "model {model.name}@{version} is not exist in local, try ‘modelhub checkout {model.name}@{version}"
                    versions_list.append(version)

            assert all(version.manifest.is_saved_model for version in versions_list), "%s is not a TensorFlow Model, cannot serving" % model.name
            # import ipdb; ipdb.set_trace()
            return ModelConfig(
                name=model.name,
                base_path=model.local_path,
                model_platform="tensorflow",
                model_version_policy=None if not versions else FileSystemStoragePathSourceConfig.ServableVersionPolicy(specific=FileSystemStoragePathSourceConfig.ServableVersionPolicy.Specific(versions=versions))
            )

        res = str(ModelServerConfig(
            model_config_list=ModelConfigList(
                config=[build_model_config(model, versions) for model, versions in model_versions_list]
            )
        ))
        # print(res)
        return res

    auto_discovery = None

    def _register_model(self, models, port):
        if not self.auto_discovery:
            from modelhub.tf_serving import AutoDiscovery, parse_redis_url, get_local_ip
            self.hostport = "%s:%d" % (get_local_ip(), port)
            if not self.redis_url:
                raise ValueError("No REDIS url configured, should set env '%s'" % REDIS_URL_ENV_KEY)
            ad_args = parse_redis_url(self.redis_url)
            ad_args["redis_prefix"] = self.redis_prefix
            self.auto_discovery = AutoDiscovery(**ad_args)
        for model, versions in models:
            print("register", model.name, versions)
            versions = versions or [model.latest_local_version.manifest.seq]
            for v in versions:
                self.auto_discovery.register(
                    model.name,
                    v,
                    self.hostport
                )

    def _run_server(self, models, model_config_file, port=8500, **kwargs):
        from subprocess import Popen, PIPE, TimeoutExpired
        import threading
        popen = Popen(
            [
                self.bin_path,
                "--port=%s" % port,
                "--model_config_file=%s" % model_config_file
            ] + [
                "--%s=%s" % (key, value) for key, value in kwargs.items()
            ],
            stdin=PIPE,
            stderr=PIPE if self.is_register_model else None,
            universal_newlines=True
        )
        try:
            if not self.is_register_model:
                return popen.wait()
            # initialized = False
            while True:
                line = popen.stderr.readline()
                print(line, end="")
                if "Running ModelServer at" in line:
                    # initialized = True
                    self._register_model(models, port)
                    break

            def print_stderr():
                while True:
                    print("start read")
                    line = popen.stderr.readline().strip()
                    if not line:
                        break
                    print(line)
            t = threading.Thread(target=print_stderr, name='print_stderr')
            t.start()

            while True:
                try:
                    popen.wait(timeout=REDIS_TIMEOUT)
                    # print("out", stderr)
                except TimeoutExpired:
                    self._register_model(models, port)
                else:
                    # print(popen.stderr.read().decode())
                    break
        finally:
            print("Killing")
            popen.kill()
