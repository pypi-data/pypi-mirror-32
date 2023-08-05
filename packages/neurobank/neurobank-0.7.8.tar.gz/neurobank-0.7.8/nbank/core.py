# -*- coding: utf-8 -*-
# -*- mode: python -*-
"""functions for managing a data archive

Copyright (C) 2013 Dan Meliza <dan@meliza.org>
Created Mon Nov 25 08:52:28 2013
"""
from __future__ import absolute_import
from __future__ import print_function
from __future__ import division
from __future__ import unicode_literals

import os
import logging

from nbank import util

log = logging.getLogger('nbank')   # root logger
env_registry = "NBANK_REGISTRY"


def default_registry():
    """Return the registry URL associated with the default registry environment variable """
    return os.environ.get(env_registry, None)


def deposit(archive_path, files, dtype=None, hash=False, auto_id=False, auth=None, **metadata):
    """Main entry point to deposit resources into an archive

    Yields the short IDs for each deposited item in files

    Here's how a variety of error conditions are handled:

    - unable to contact registry: ConnectionError
    - attempt to add unallowed directory: skip the directory
    - failed to register resource for any reason: HTTPError, usually 400 error code
    - unable to match archive path to archive in registry: RuntimeError
    - failed to add the file (usually b/c the identifier is taken): RuntimeError

    The last two errors indicate a major problem where the archive has perhaps
    been moved, or the archive int he registry is not pointing to the right
    location, or data has been stored in the archive without contacting the
    registry. These are currently hairy enough problems that the user is going
    to have to fix them herself for now.

    """
    import uuid
    from nbank.archive import get_config, store_resource
    from nbank.registry import add_resource, find_archive_by_path, full_url
    archive_path = os.path.abspath(archive_path)
    cfg = get_config(archive_path)
    if cfg is None:
        raise ValueError("%s is not a valid archive" % archive_path)
    log.info("archive: %s", archive_path)
    registry_url = cfg["registry"]
    log.info("   registry: %s", registry_url)
    auto_id = cfg['policy']['auto_identifiers'] or auto_id
    auto_id_type = cfg['policy'].get('auto_id_type', None)
    allow_dirs = cfg['policy']['allow_directories']

    # check that archive exists for this path
    archive = find_archive_by_path(registry_url, archive_path)
    log.info("   archive name: %s", archive)
    if archive is None:
        raise RuntimeError("archive '%s' not in registry. did it move?" % archive_path)

    for src in files:
        log.info("processing '%s':", src)
        if not os.path.exists(src):
           log.info("   does not exist; skipping")
           continue
        if not allow_dirs and os.path.isdir(src):
            log.info("   is a directory; skipping")
            continue
        if auto_id:
            if auto_id_type == "uuid":
                id = str(uuid.uuid4())
            else:
                id = None
        else:
            id = util.id_from_fname(src)
        if hash or cfg['policy']['require_hash']:
            sha1 = util.hash(src)
            log.info("   sha1: %s", sha1)
        else:
            sha1 = None
        result = add_resource(registry_url, id, dtype, archive, sha1, auth, **metadata)
        id = full_url(registry_url, result["name"])
        log.info("   registered as %s", id)
        tgt = store_resource(cfg, src, id=result["name"])
        log.info("   deposited in %s", tgt)
        yield {"source": src, "id": result["name"]}


def find(id, registry_url=None, local_only=False):
    """Generates a sequence of paths or URLs where id can be located

    If local_only is True, only files that can be found on the local filesystem
    are yielded.

    """
    from nbank.registry import parse_resource_id, get_locations
    from nbank.archive import find_resource
    if registry_url is None:
        registry_url = default_registry()
    base, sid = parse_resource_id(registry_url, id)
    if base is None:
        raise ValueError("short identifier supplied without a registry to resolve it")
    for loc in get_locations(base, sid):
        path = get_path_or_url(loc)
        if local_only:
            yield find_resource(path)
        else:
            yield path

def get(id, registry_url=None, local_only=False):
    """Returns the first path or URL where id can be found, or None if no match.

    If local_only is True, only files that can be found on the local filesystem
    are considered.

    """
    try:
        return next(find(id, registry_url, local_only))
    except StopIteration:
        pass


def describe(id, registry_url=None):
    """Returns the database record for id, or None if no match can be found """
    from nbank.registry import get_resource, parse_resource_id
    if registry_url is None:
        registry_url = default_registry()
    base, sid = parse_resource_id(registry_url, id)
    return get_resource(base, sid)


def get_path_or_url(location):
    """Return the path or URL associated with location

    location is a dict with 'scheme', 'path', and 'resource_name' (like what's
    yielded by registry.get_locations). Note that for local (neurobank)
    locations, the resource may have an extension; this is not included.

    """
    from nbank.archive import resource_path
    try:
        from urllib.parse import urlunparse
    except ImportError:
        from urlparse import urlunparse
    if location["scheme"] == "neurobank":
        return resource_path(location["root"], location["resource_name"])
    else:
        return urlunparse((location["scheme"], location["root"],
                           location["resource_name"], '', '', ''))



# Variables:
# End:
