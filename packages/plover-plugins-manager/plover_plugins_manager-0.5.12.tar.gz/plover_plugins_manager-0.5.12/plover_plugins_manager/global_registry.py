
from collections import defaultdict
from datetime import datetime
from xmlrpc.client import ServerProxy
import json
import os

from pkg_resources import parse_version, safe_name
try:
    import pip._internal as pip_internal
except ImportError:
    import pip as pip_internal

from plover.oslayer.config import CONFIG_DIR

from plover_plugins_manager.plugin_metadata import PluginMetadata


CACHE_FILE = os.path.join(CONFIG_DIR, '.cache', 'plugins.json')
CACHE_VERSION = 3


def load_cache():
    if not os.path.exists(CACHE_FILE):
        return {}
    with open(CACHE_FILE, 'r') as fp:
        return json.load(fp)

def save_cache(**kwargs):
    dirname = os.path.dirname(CACHE_FILE)
    if not os.path.exists(dirname):
        os.makedirs(dirname)
    with open(CACHE_FILE, 'w') as fp:
        json.dump(kwargs, fp, indent=2, sort_keys=True)


def list_plugins():
    session = pip_internal.download.PipSession()
    index_url = pip_internal.models.PyPI.pypi_url
    # We use pip's session/transport to avoid SSL errors on Windows/macOS...
    transport = pip_internal.download.PipXmlrpcTransport(index_url, session)
    pypi = ServerProxy(index_url, transport)
    cache = load_cache()
    if cache.get('version') == CACHE_VERSION and \
       (cache.get('timestamp', 0.0) + 600.0) >= datetime.utcnow().timestamp():
        plugins = {
            name: [PluginMetadata(*[
                v.get(k, '')
                for k in PluginMetadata._fields
            ]) for v in versions]
            for name, versions in cache.get('plugins', {}).items()
        }
        return plugins
    plugins = defaultdict(list)
    for match in pypi.search({'keywords': 'plover_plugin'}):
        name, version = match['name'], match['version']
        metadata_dict = pypi.release_data(name, version)
        # Can happen if a package has been deleted.
        if not metadata_dict:
            continue
        plugin_metadata = PluginMetadata(*[
            metadata_dict.get(k, '')
            for k in PluginMetadata._fields
        ])
        assert name == plugin_metadata.name
        assert version == plugin_metadata.version
        plugins[safe_name(name)].append(plugin_metadata)
    plugins = {
        name: list(sorted(versions))
        for name, versions in plugins.items()
    }
    save_cache(version=CACHE_VERSION,
               timestamp=datetime.utcnow().timestamp(),
               plugins={
                   name: [v.to_dict() for v in versions]
                   for name, versions in plugins.items()
               })
    return plugins
