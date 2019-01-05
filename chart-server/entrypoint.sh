#!/usr/bin/env bash

case "$1" in
  chart-server)
    exec ./chart-server
    ;;
  *)
    # The command is something like bash, not an airflow subcommand. Just run it in the right environment.
    exec "$@"
    ;;
esac
