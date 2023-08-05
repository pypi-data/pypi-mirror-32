<?php

define('CLI_SCRIPT', true);
require(dirname(__FILE__).'/config.php');
require_once($CFG->libdir.'/clilib.php');
require("$CFG->dirroot/version.php");

cli_separator();
cli_heading('Resetting all version numbers');

$manager = core_plugin_manager::instance();

$plugininfo = $manager->get_plugins();
foreach ($plugininfo as $type => $plugins) {
    foreach ($plugins as $name => $plugin) {
        if ($plugin->get_status() !== core_plugin_manager::PLUGIN_STATUS_DOWNGRADE) {
            continue;
        }

        $frankenstyle = sprintf("%s_%s", $type, $name);

        mtrace("Updating {$frankenstyle} from {$plugin->versiondb} to {$plugin->versiondisk}");
        $DB->set_field('config_plugins', 'value', $plugin->versiondisk, array('name' => 'version', 'plugin' => $frankenstyle));
    }
}

// Check that the main version hasn't changed.
if ((float) $CFG->version !== $version) {
    set_config('version', $version);
    mtrace("Updated main version from {$CFG->version} to {$version}");
}

// Purge all caches.
$cache = cache::make('core', 'plugin_manager');
$cache->purge();

$cache = cache::make('core', 'config');
$cache->purge();
