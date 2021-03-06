import os
import time
import logging
from behave import given, when, then
from datetime import date
from pyshould import should
from kubernetes import config, client
from smoke.features.steps.openshift import Openshift
from smoke.features.steps.project import Project


config.load_kube_config()
v1 = client.CoreV1Api()
podStatus = {}
operatorPods = {}
operator_image = "quay.io/redhat-developer/openshift-jenkins-operator:latest"
def get_filename_datetime():
    # Use current date to get a text file name.
    return "" + str(date.today()) + ".txt"


# Get full path for writing.
file_name = get_filename_datetime()
path = "./smoke/logs-" + file_name
crd_path = "./deploy/crds/jenkins_all_crd.yaml"
template_path = "./smoke/jenkins-operator-template.yaml"
cr_path = "./deploy/crds/openshift_jenkins_v1alpha2_jenkins_cr.yaml"
logging.basicConfig(filename=path, format='%(asctime)s: %(levelname)s: %(message)s', datefmt='%m/%d/%Y %I:%M:%S %p')
logger = logging.getLogger()
logger.setLevel(logging.INFO)

current_project = ''
paramfile = 'smoke/templates.params'
oc = Openshift()
# STEP
@given(u'Project "{project_name}" is used')
def given_project_is_used(context, project_name):
    project = Project(project_name)
    global current_project
    current_project = project_name
    context.current_project = current_project
    context.oc = oc
    if not project.is_present():
        print("Project is not present, creating project: {}...".format(project_name))
        project.create() | should.be_truthy.desc("Project {} is created".format(project_name))
    print("Project {} is created!!!".format(project_name))
    logger.info("Project {} is created!!!".format(project_name))
    context.project = project


# STEP
@given(u'Project [{project_env}] is used')
def given_namespace_from_env_is_used(context, project_env):
    env = os.getenv(project_env)
    assert env is not None, f"{project_env} environment variable needs to be set"
    print(f"{project_env} = {env}")
    given_project_is_used(context, env)




@given(u'We have a openshift cluster')
def loginCluster(context):
    print("Using [{}]".format(current_project))
    

@when(u'the resources are created using the crd')
def createResources(context):
    logger.info("using crd to create the required resources")
    res = oc.oc_create_from_yaml(crd_path)
    
    

@then(u'We create template from yaml')
def createTemplate(context):
    logger.info('Creating the template')
    res = oc.oc_create_from_yaml(template_path)
    logger.info(res)
    time.sleep(20)
    if not 'jenkins-operator-template' in res:
        raise AssertionError


@then(u'Apply template with oc new-app')
def createOperator(context):
    template_name = 'jenkins-operator-template'
    logger.info('Applying the template')
    if os.path.exists(paramfile):
        print("Removing param file")
        os.remove(paramfile)
    if not os.path.exists(paramfile):
        print("Creating the param file for the template to use")
        logger.info("Creating the param file for the template to use")
        f = open(paramfile, "w")
        f.write("JENKINS_OPERATOR_NAMESPACE = %s\n"%current_project)
        f.write("JENKINS_OPERATOR_IMAGE = %s\n"%operator_image)
        f.close()

    oc.new_app_with_params(template_name,paramfile)
    time.sleep(10)
    


@then(u'Check for pod creation and state')
def checkOperatorPod(context):
    time.sleep(30)
    context.v1 = v1

    logger.info('Searching resources')
    pods = v1.list_namespaced_pod(current_project)
    for i in pods.items:
        logger.info("Getting pod list")
        logger.info(i.status.pod_ip)
        logger.info(i.metadata.name)
        logger.info(i.status.phase)
        podStatus[i.metadata.name] = i.status.phase
        logger.info('---> Validating...')
        if not i.metadata.name in oc.search_pod_in_namespace(i.metadata.name,current_project):
            raise AssertionError

    logger.info('waiting to get pod status')
    time.sleep(60)
    for pod in podStatus.keys():
        status = podStatus[pod]
        if 'Running' in status:
            logger.info(pod)
            logger.info(podStatus[pod])
        else:
            raise AssertionError