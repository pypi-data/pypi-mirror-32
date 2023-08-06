# EMR Launcher

Launches EMR clusters using config files for consistent run-time behavior when setting up a cluster.

## To Run

Starting a new cluster:
```
PYTHONPATH=/path/to/emr_launcher python emr_launcher/launcher.py /path/to/config/<my_config>.json
```

Adding steps to an existing cluster
```
PYTHONPATH=/path/to/emr_launcher python emr_launcher/launcher.py config/<my_config>.json --job-id <job_id_of_existing_cluster>
```

## Creating configs

Create a JSON file in the configs directory. Fill the config based on the parameters defined here: http://boto3.readthedocs.io/en/latest/reference/services/emr.html#EMR.Client.run_job_flow

*or*

build off the example config

## Using templating in configs

In any JSON config file function defined in the `template_functions` module can be used inside the config using jinja2 style templating. If you require a new function, add it to the `template_functions` module and it will be available to use in any config.


