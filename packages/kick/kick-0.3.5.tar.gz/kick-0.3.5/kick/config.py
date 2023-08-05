import os
import re
import shutil
import inspect
import pathlib

import toml
import addict
from first import first

CONFIGDIR = pathlib.Path.home() / ".config"
ENV_VAR_PATTERN = re.compile("@([A-Z_]+)")


def get_caller_path():
    stack = inspect.stack()
    try:
        frame = first(
            stack, key=lambda frame: "kick." in (frame.code_context[0] or [""])
        )
        path = pathlib.Path(frame.filename).parent
    except:
        path = None
    finally:
        del stack

    return path


def get_local_config_path(variant):
    local_dir = get_caller_path() or pathlib.Path(".")
    path = local_dir / "config" / "{}.toml".format(variant)
    if not path.exists():
        path = local_dir / "config" / "config.toml"
    return path


def env_var_or_key(match):
    return os.getenv(match.group(1), match.group(0))


def replace_env_variables(config):
    if isinstance(config, str):
        return ENV_VAR_PATTERN.sub(env_var_or_key, config)
    if isinstance(config, list):
        return [replace_env_variables(v) for v in config]
    if isinstance(config, dict):
        config = addict.Dict(config)
        for key, value in config.items():
            if isinstance(value, str):
                config[key] = ENV_VAR_PATTERN.sub(env_var_or_key, value)
            elif isinstance(value, list):
                config[key] = [replace_env_variables(v) for v in value]
            elif isinstance(value, dict):
                config[key] = replace_env_variables(value)
    return config


def Config(name, path=None, variant="config"):
    config_path = CONFIGDIR / name / "{}.toml".format(variant)
    if not config_path.exists():
        config_path.parent.mkdir(parents=True, exist_ok=True)
        local_config_path = path or get_local_config_path(variant)
        shutil.copy(str(local_config_path), str(config_path))
        print("Created config: {}".format(config_path))

    config = addict.Dict(toml.loads(config_path.read_text()))
    config = replace_env_variables(config)
    return config


def update_config(name, path=None, variant="config"):
    config_path = CONFIGDIR / name / "{}.toml".format(variant)
    config = Config(name, path, variant)
    local_config = addict.Dict(
        toml.loads((path or get_local_config_path(variant)).read_text())
    )
    local_config.update(config)
    config_path.write_text(toml.dumps(local_config))
    print("Updated config: {}".format(config_path))
