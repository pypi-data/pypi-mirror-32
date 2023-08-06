"""Script entry point."""

import argparse

from toolrack.script import ErrorExitMessage

from prometheus_aioexporter.script import PrometheusExporterScript
from prometheus_aioexporter.metric import InvalidMetricType

from .config import load_config, ConfigError
from .loop import QueryLoop


class QueryExporterScript(PrometheusExporterScript):
    """Periodically run database queries and export results to Prometheus."""

    name = 'query-exporter'

    description = __doc__

    def configure_argument_parser(self, parser):
        parser.add_argument(
            'config', type=argparse.FileType('r'),
            help='configuration file')

    def configure(self, args):
        config = self._load_config(args.config)
        self.create_metrics(config.metrics)
        self.query_loop = QueryLoop(
            config, self.registry, self.logger, self.loop)

    async def on_application_startup(self, application):
        application['exporter'].set_metric_update_handler(
            self._update_handler)
        await self.query_loop.start()

    async def on_application_shutdown(self, application):
        await self.query_loop.stop()

    async def _update_handler(self, metrics):
        """Run queries with no specified interval on each request."""
        await self.query_loop.run_aperiodic_queries()

    def _load_config(self, config_file):
        """Load the application configuration."""
        try:
            config = load_config(config_file)
        except (InvalidMetricType, ConfigError) as error:
            raise ErrorExitMessage(str(error))
        finally:
            config_file.close()
        return config


script = QueryExporterScript()
