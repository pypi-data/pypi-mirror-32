"""cli entry point for pipeline runner.

Parse command line arguments in, invoke pipelinerunner.
"""
import argparse
import pypyrcli.pipelinerunner


def get_args():
    """Parse arguments passed in from shell."""
    parser = argparse.ArgumentParser(
        allow_abbrev=True,
        description='pypyr pipeline runner')
    parser.add_argument('--name', dest='pipeline_name', required=True,
                        help='Name of pipeline to run. It should exist in the '
                        '/pipelines directory.')
    parser.add_argument('--context', dest='pipeline_context',
                        help='String for context values. Parsed by '
                        'pipeline''s context_parser function.')
    parser.add_argument('--loglevel', dest='log_level', type=int, default=10,
                        help='Integer log level. Defaults to 10 (Debug). '
                        '10=DEBUG 20=INFO 30=WARNING 40=ERROR 50=CRITICAL')
    parser.add_argument('--version', action='version',
                        help='Echo version number.',
                        version=f'%(prog)s {pypyrcli.__version__}')

    return parser.parse_args()


def main():
    args = get_args()
    return pypyrcli.pipelinerunner.main(
        pipeline_name=args.pipeline_name,
        pipeline_context_input=args.pipeline_context,
        log_level=args.log_level)
