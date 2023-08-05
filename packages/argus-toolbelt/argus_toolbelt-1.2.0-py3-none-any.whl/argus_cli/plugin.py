import imp, inspect
from os.path import splitext, basename, exists
from glob import glob

from argus_cli import arguments
from argus_cli.helpers.collections import ImmutableDeepDict
from argus_cli.helpers.log import log
from argus_cli.helpers.formatting import to_caterpillar_case, to_camel_case, to_snake_case, is_snake_case, is_camel_case

#: Contains all plugins and their functions
_plugins = ImmutableDeepDict()

#: The API that plugins will interact with
api = None


def _get_or_set_plugin(name: tuple) -> dict:
    """Get's a plugin. If it doesn't exist it registers it"""
    new_plugin = {}

    plugin = _plugins.setdefault(name, new_plugin)
    if plugin is new_plugin:
        arguments.register_plugin(name)

    return plugin


def _add_command(func: callable, plugin: dict, command_name: str) -> dict:
    """Adds a command to a plugin

    :param func: The function to add
    :param plugin: The plugin to add the function to
    :param command_name: The name of the function
    :raises NameError: If a command already has the same name
    """
    import inspect
    # Only raise this error if the plugin is a different function
    # This is because if the function is imported multiple times, e.g as module.function,
    # and as from module import function, they will be registered with different IDs,
    # and will attempt to register twice with their decorator. Thus, we have to compare
    # the function.__code__ to only raise the warning if the functions are actually two
    # different functions.
    if command_name in plugin and func.__code__ != plugin[command_name].__code__:
        raise NameError(
            "Command '%s' already registered by %s (#%s). Cannot re-assign to %s (#%s)" %
            (command_name,
             inspect.getsourcefile(plugin[command_name]),
             str(id(plugin[command_name])),
             inspect.getsourcefile(func),
             str(id(func.__module__)))
        )

    plugin[command_name] = func

    return plugin


def register_command(alias: str = None, extending: tuple = None) -> callable:
    """Decorator used to register commands to a plugin

    :param alias: If the user wants a custom name on the plugin
    :param extending: A existing plugin to extend
    """
    def decorate(func):
        if extending:
            plugin_name = (extending,) if isinstance(extending, str) else extending
        else:
            plugin_name = (func.__module__,)
        plugin_name = tuple(map(to_caterpillar_case, plugin_name))
        command_name = to_caterpillar_case(alias or func.__name__)

        plugin = _get_or_set_plugin(plugin_name)
        _add_command(func, plugin, command_name)
        arguments.register_command(plugin_name, command_name)

        return func

    return decorate


def register_command_metadata(plugin_name: tuple, command_name: str) -> None:
    """Registers the commands metadata

    Wraps the function with the same name in arguments.
    This is because this module contains the actual function to parse.

    :param plugin_name: The name of the plugin
    :param command_name: The name of the function
    """
    arguments.register_command_metadata(plugin_name, command_name, _plugins[plugin_name][command_name])


def get_plugin_modules(locations: list) -> list:
    """Loads plugins from default plugin location and user defined plugin location,
    and attempts to load them as python modules. 
    
    Directories can be loaded, provided they are python packages, i.e containing an __init__.py file. 
    If a plugin is defined as a Python package with an __init__.py file, this file must export
    all functions decorated with `@register_command`, since these will be registered on import, and
    only the __init__.py will be initially imported.
     
    NOTE: Renames plugins to common unix command naming scheme

    :param list locations: Folder with plunspgins
    :rtype: list
    :returns: A list of python files with paths
    """
    modules = []

    for path in locations:
        log.debug("Loading plugins from %s..." % path)

        if not exists(path):
            log.warning("Plugin directory does not exist: %s" % path)
            continue

        # Load plugins that dont start with __ (__pycache__, __init__, __main__, etc)
        # and force the paths to the filenames
        for plugin in map(basename, glob("%s/[!__]*" % path)):
            log.debug("Extracting plugin metadata from: %s" % plugin)

            # Get the file without file extension
            module_name, file_ending = splitext(plugin)

            if file_ending and not file_ending.startswith(".py"):
                continue

            try:
                # Explicitly show the return types, to avoid confusion:
                file_reference, path_to_file, file_information = imp.find_module("%s" % module_name, [path])
            except ImportError:
                log.critical("Could not load module: %s (%s)" % (module_name, path))
                continue

            log.debug("Loaded plugin %s" % to_caterpillar_case(module_name))
            modules.append({
                "name": to_caterpillar_case(module_name),
                "info": (file_reference, path_to_file, file_information)
            })

    return modules


def load_plugin_module(plugin: dict) -> bool:
    """Loads a plugin

    :param dict plugin: A dict with the module name and info
    :returns: True if module was successfully loaded
    :rtype: bool
    """
    log.info("Loading plugin: %s" % plugin["name"])

    try:
        imp.load_module(plugin["name"].replace("-", "_"), *plugin["info"])
    except Exception:
        log.exception("Error while loading plugin module %s:" % plugin["name"])
        return False
    finally:
        if plugin["info"][0]:
            plugin["info"][0].close()

    return True


def run_command(plugin_name: tuple, command: str, arguments: dict = None) -> None:
    """Runs the specified command

    :param plugin_name: The plugin(module) to source the command from
    :param command: The command to run
    :param arguments: Arguments to pass to the function
    :returns: Anything the command returns
    """

    command = _plugins[plugin_name][command]

    if arguments:    
        # Arguments in caterpillar-case must be turned into snake_case again,
        # if method parameters are snake_cased; the API methods are camelCased though,
        # and cater-pillar cased CLI arguments must turn into camelCase instead:
        signature = inspect.signature(command)
        argument_count = len(signature.parameters.keys())

        # Check how many of the functions arguments are considered to be in camelCase
        camel_case_probability = sum([
            100/argument_count
            if is_camel_case(x)
            else 0
            for x in signature.parameters.keys()
        ])

        # Compare this to how many of the functions arguments are considered to be snake_case
        snake_case_probability = sum([
            100/argument_count
            if is_snake_case(x)
            else 0
            for x in signature.parameters.keys()
        ])

        # ... then transform the caterpillar-case arguments from command line to the appropriate case
        # This is done because the auto-generated API classes have parameters corresponding to the 
        # camelCased API, while we otherwise prefer plugins to have snake_cased parameters.
        # To allow API classes to still have parameter names that reflect the actual API, while
        # encouraging plugins to use snake_cased parameters, we have to cast the command-line caterpillar-case
        # to either snake_case or camelCase depending on the function
        formatter = to_snake_case if snake_case_probability > camel_case_probability else to_camel_case

        formatted_arguments = {formatter(key): value for key, value in arguments.items() if value is not None}
        log.debug("Running command \"%s %s\" with arguments: %s" % (plugin_name, command, formatted_arguments))
        return command(**formatted_arguments)
    else:
        return command()
