# Tool builder: `gcr.io/cloud-builders/mvn`

This Container Builder build step runs Maven.

You might also consider using an [official `maven` image](https://hub.docker.com/_/maven/) and specifying the `mvn` entrypoint:

 ```yaml
 steps:
 - name: maven:3.5.4
   entrypoint: mvn
   args: ['build']

## Building this builder

To build this builder, run the following command in this directory.

    $ gcloud builds submit . --config=cloudbuild.yaml
