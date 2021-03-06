package e2e

import (
	goctx "context"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/jenkinsci/kubernetes-operator/internal/try"
	"github.com/jenkinsci/kubernetes-operator/pkg/apis/jenkins/v1alpha2"
	"github.com/jenkinsci/kubernetes-operator/pkg/configuration/base/resources"

	framework "github.com/operator-framework/operator-sdk/pkg/test"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	retryInterval               = time.Second * 5
	timeout                     = time.Second * 60
	namespaceTerminationTimeout = time.Second * 120
	retries                     = 60
)

// checkConditionFunc is used to check if a condition for the jenkins CR is set
type checkConditionFunc func(*v1alpha2.Jenkins, error) bool

func waitForJenkinsBaseConfigurationToComplete(t *testing.T, jenkins *v1alpha2.Jenkins) {
	t.Logf("Waiting for Jenkins base configuration to complete: Will retry %+v times every %+v", retries, retryInterval)
	var completedTime *metav1.Time
	isJenkinsBaseConfigurationCompletedTimeSet := func(jenkins *v1alpha2.Jenkins, err error) bool {
		completedTime = jenkins.Status.BaseConfigurationCompletedTime
		t.Logf("Current Jenkins status BaseConfigurationCompletedTime: '%+v', error '%s' in namespace %s", completedTime, err, jenkins.Namespace)
		jenkinsPhase := jenkins.Status.Phase
		if len(jenkinsPhase) > 0 {
			t.Logf("Jenkins instance %s in Phase: %s", jenkins.Name, jenkinsPhase)
		}
		return err == nil && completedTime != nil
	}
	_, err := WaitForCondition(retryInterval, retries, jenkins, isJenkinsBaseConfigurationCompletedTimeSet)
	if err != nil {
		t.Errorf("Waiting for BaseConfiguration to complete failed with : %+v", jenkins)
		t.Fatal(err)
	}
	t.Log("Jenkins pod is running")
	// update jenkins CR because Operator sets default values
	namespacedName := types.NamespacedName{Namespace: jenkins.Namespace, Name: jenkins.Name}
	err = framework.Global.Client.Get(goctx.TODO(), namespacedName, jenkins)

	assert.NoError(t, err)

	t.Log("Sleep for 10 seconds after BaseConfiguration completes")
	time.Sleep(10 * time.Second)
}

func waitForJenkinsPodRecreation(t *testing.T, jenkins *v1alpha2.Jenkins) {
	t.Logf("Waiting for Jenkins Master Pod to be ready: Will every %+v until %v", retryInterval, 30*retryInterval)
	IsJenkinsMasterPodRecreated := func() (bool, error) {
		lo := metav1.ListOptions{
			LabelSelector: labels.SelectorFromSet(resources.BuildResourceLabels(jenkins)).String(),
		}
		podList, err := framework.Global.KubeClient.CoreV1().Pods(jenkins.ObjectMeta.Namespace).List(lo)
		if err != nil {
			t.Logf("Error while listing Jenkins Pods List : %+v \n Error: %+v", podList, err)
			return false, err
		}
		pods := podList.Items
		if len(pods) != 1 {
			if len(pods) == 2 {
				t.Logf("Two pods are present for the Deployment selector")
				for i, p := range pods {
					t.Logf("Pod no. %d is in the %s Status.Phase", i, p.Status.Phase)
				}
				return false, nil
			}
			return false, nil
		} else if len(pods) == 1 {
			jenkinsPod := pods[0]
			t.Logf("Detected single Jenkins Pod %s for given selector with DeletionTimestamp %s", jenkinsPod.Name, jenkinsPod.DeletionTimestamp)
			if jenkinsPod.Status.Phase == "Running" && jenkinsPod.DeletionTimestamp == nil {
				t.Logf("Jenkins Pod %s  is in Running state now", jenkinsPod.Name)
				return true, nil
			}
			return false, nil
		}
		return false, err
	}
	err := wait.Poll(retryInterval, 30*retryInterval, IsJenkinsMasterPodRecreated)
	if err != nil {
		t.Errorf("Waiting for Jenkins Master Pod to recreate failed with jenkins: %+v", jenkins)
		t.Fatal(err)
	}
	t.Log("Jenkins pod has been recreated")
}

// WaitForCondition retries until the specified condition check becomes true for the jenkins CR
func WaitForCondition(retryInterval time.Duration, retries int, jenkins *v1alpha2.Jenkins, checkCondition checkConditionFunc) (*v1alpha2.Jenkins, error) {
	jenkinsStatus := &v1alpha2.Jenkins{}
	err := wait.Poll(retryInterval, time.Duration(retries)*retryInterval, func() (bool, error) {
		namespacedName := types.NamespacedName{Namespace: jenkins.Namespace, Name: jenkins.Name}
		err := framework.Global.Client.Get(goctx.TODO(), namespacedName, jenkinsStatus)
		return checkCondition(jenkinsStatus, err), nil
	})
	if err != nil {
		return nil, err
	}
	return jenkinsStatus, nil
}

func waitUntilNamespaceDestroyed(namespace string) error {
	isNamespaceDestroyed := func() (bool, error) {
		var namespaceList v1.NamespaceList
		err := framework.Global.Client.List(context.TODO(), &namespaceList, &client.ListOptions{})
		if err != nil {
			return true, err
		}
		exists := false
		for _, namespaceItem := range namespaceList.Items {
			if namespaceItem.Name == namespace {
				exists = true
				break
			}
		}
		return !exists, nil
	}
	err := try.Until(isNamespaceDestroyed, time.Second, namespaceTerminationTimeout)
	return err
}
