// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedGroovyScript) DeepCopyInto(out *AppliedGroovyScript) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedGroovyScript.
func (in *AppliedGroovyScript) DeepCopy() *AppliedGroovyScript {
	if in == nil {
		return nil
	}
	out := new(AppliedGroovyScript)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapRef) DeepCopyInto(out *ConfigMapRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapRef.
func (in *ConfigMapRef) DeepCopy() *ConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(ConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationAsCode) DeepCopyInto(out *ConfigurationAsCode) {
	*out = *in
	in.Customization.DeepCopyInto(&out.Customization)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationAsCode.
func (in *ConfigurationAsCode) DeepCopy() *ConfigurationAsCode {
	if in == nil {
		return nil
	}
	out := new(ConfigurationAsCode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Container) DeepCopyInto(out *Container) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Lifecycle != nil {
		in, out := &in.Lifecycle, &out.Lifecycle
		*out = new(v1.Lifecycle)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Container.
func (in *Container) DeepCopy() *Container {
	if in == nil {
		return nil
	}
	out := new(Container)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Customization) DeepCopyInto(out *Customization) {
	*out = *in
	out.Secret = in.Secret
	if in.Configurations != nil {
		in, out := &in.Configurations, &out.Configurations
		*out = make([]ConfigMapRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Customization.
func (in *Customization) DeepCopy() *Customization {
	if in == nil {
		return nil
	}
	out := new(Customization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroovyScripts) DeepCopyInto(out *GroovyScripts) {
	*out = *in
	in.Customization.DeepCopyInto(&out.Customization)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroovyScripts.
func (in *GroovyScripts) DeepCopy() *GroovyScripts {
	if in == nil {
		return nil
	}
	out := new(GroovyScripts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Handler) DeepCopyInto(out *Handler) {
	*out = *in
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(v1.ExecAction)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Handler.
func (in *Handler) DeepCopy() *Handler {
	if in == nil {
		return nil
	}
	out := new(Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jenkins) DeepCopyInto(out *Jenkins) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jenkins.
func (in *Jenkins) DeepCopy() *Jenkins {
	if in == nil {
		return nil
	}
	out := new(Jenkins)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jenkins) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsAPISettings) DeepCopyInto(out *JenkinsAPISettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsAPISettings.
func (in *JenkinsAPISettings) DeepCopy() *JenkinsAPISettings {
	if in == nil {
		return nil
	}
	out := new(JenkinsAPISettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsImage) DeepCopyInto(out *JenkinsImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsImage.
func (in *JenkinsImage) DeepCopy() *JenkinsImage {
	if in == nil {
		return nil
	}
	out := new(JenkinsImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsImageList) DeepCopyInto(out *JenkinsImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]JenkinsImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsImageList.
func (in *JenkinsImageList) DeepCopy() *JenkinsImageList {
	if in == nil {
		return nil
	}
	out := new(JenkinsImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsImageSpec) DeepCopyInto(out *JenkinsImageSpec) {
	*out = *in
	out.BaseImage = in.BaseImage
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]JenkinsPlugin, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsImageSpec.
func (in *JenkinsImageSpec) DeepCopy() *JenkinsImageSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsImageStatus) DeepCopyInto(out *JenkinsImageStatus) {
	*out = *in
	if in.InstalledPlugins != nil {
		in, out := &in.InstalledPlugins, &out.InstalledPlugins
		*out = make([]JenkinsPlugin, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsImageStatus.
func (in *JenkinsImageStatus) DeepCopy() *JenkinsImageStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsList) DeepCopyInto(out *JenkinsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jenkins, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsList.
func (in *JenkinsList) DeepCopy() *JenkinsList {
	if in == nil {
		return nil
	}
	out := new(JenkinsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JenkinsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsMaster) DeepCopyInto(out *JenkinsMaster) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AnnotationsDeprecated != nil {
		in, out := &in.AnnotationsDeprecated, &out.AnnotationsDeprecated
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BasePlugins != nil {
		in, out := &in.BasePlugins, &out.BasePlugins
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsMaster.
func (in *JenkinsMaster) DeepCopy() *JenkinsMaster {
	if in == nil {
		return nil
	}
	out := new(JenkinsMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsPlugin) DeepCopyInto(out *JenkinsPlugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsPlugin.
func (in *JenkinsPlugin) DeepCopy() *JenkinsPlugin {
	if in == nil {
		return nil
	}
	out := new(JenkinsPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSpec) DeepCopyInto(out *JenkinsSpec) {
	*out = *in
	in.Master.DeepCopyInto(&out.Master)
	if in.SeedJobs != nil {
		in, out := &in.SeedJobs, &out.SeedJobs
		*out = make([]SeedJob, len(*in))
		copy(*out, *in)
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]Notification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Service.DeepCopyInto(&out.Service)
	in.SlaveService.DeepCopyInto(&out.SlaveService)
	in.Backup.DeepCopyInto(&out.Backup)
	in.Restore.DeepCopyInto(&out.Restore)
	in.GroovyScripts.DeepCopyInto(&out.GroovyScripts)
	in.ConfigurationAsCode.DeepCopyInto(&out.ConfigurationAsCode)
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]rbacv1.RoleRef, len(*in))
		copy(*out, *in)
	}
	in.ServiceAccount.DeepCopyInto(&out.ServiceAccount)
	out.JenkinsAPISettings = in.JenkinsAPISettings
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSpec.
func (in *JenkinsSpec) DeepCopy() *JenkinsSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsStatus) DeepCopyInto(out *JenkinsStatus) {
	*out = *in
	if in.ProvisionStartTime != nil {
		in, out := &in.ProvisionStartTime, &out.ProvisionStartTime
		*out = (*in).DeepCopy()
	}
	if in.BaseConfigurationCompletedTime != nil {
		in, out := &in.BaseConfigurationCompletedTime, &out.BaseConfigurationCompletedTime
		*out = (*in).DeepCopy()
	}
	if in.UserConfigurationCompletedTime != nil {
		in, out := &in.UserConfigurationCompletedTime, &out.UserConfigurationCompletedTime
		*out = (*in).DeepCopy()
	}
	if in.CreatedSeedJobs != nil {
		in, out := &in.CreatedSeedJobs, &out.CreatedSeedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AppliedGroovyScripts != nil {
		in, out := &in.AppliedGroovyScripts, &out.AppliedGroovyScripts
		*out = make([]AppliedGroovyScript, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsStatus.
func (in *JenkinsStatus) DeepCopy() *JenkinsStatus {
	if in == nil {
		return nil
	}
	out := new(JenkinsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Mailgun) DeepCopyInto(out *Mailgun) {
	*out = *in
	out.APIKeySecretKeySelector = in.APIKeySecretKeySelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Mailgun.
func (in *Mailgun) DeepCopy() *Mailgun {
	if in == nil {
		return nil
	}
	out := new(Mailgun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MicrosoftTeams) DeepCopyInto(out *MicrosoftTeams) {
	*out = *in
	out.WebHookURLSecretKeySelector = in.WebHookURLSecretKeySelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MicrosoftTeams.
func (in *MicrosoftTeams) DeepCopy() *MicrosoftTeams {
	if in == nil {
		return nil
	}
	out := new(MicrosoftTeams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Notification) DeepCopyInto(out *Notification) {
	*out = *in
	if in.Slack != nil {
		in, out := &in.Slack, &out.Slack
		*out = new(Slack)
		**out = **in
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = new(MicrosoftTeams)
		**out = **in
	}
	if in.Mailgun != nil {
		in, out := &in.Mailgun, &out.Mailgun
		*out = new(Mailgun)
		**out = **in
	}
	if in.SMTP != nil {
		in, out := &in.SMTP, &out.SMTP
		*out = new(SMTP)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Notification.
func (in *Notification) DeepCopy() *Notification {
	if in == nil {
		return nil
	}
	out := new(Notification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restore) DeepCopyInto(out *Restore) {
	*out = *in
	in.Action.DeepCopyInto(&out.Action)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restore.
func (in *Restore) DeepCopy() *Restore {
	if in == nil {
		return nil
	}
	out := new(Restore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMTP) DeepCopyInto(out *SMTP) {
	*out = *in
	out.UsernameSecretKeySelector = in.UsernameSecretKeySelector
	out.PasswordSecretKeySelector = in.PasswordSecretKeySelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMTP.
func (in *SMTP) DeepCopy() *SMTP {
	if in == nil {
		return nil
	}
	out := new(SMTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedJob) DeepCopyInto(out *SeedJob) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedJob.
func (in *SeedJob) DeepCopy() *SeedJob {
	if in == nil {
		return nil
	}
	out := new(SeedJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Slack) DeepCopyInto(out *Slack) {
	*out = *in
	out.WebHookURLSecretKeySelector = in.WebHookURLSecretKeySelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slack.
func (in *Slack) DeepCopy() *Slack {
	if in == nil {
		return nil
	}
	out := new(Slack)
	in.DeepCopyInto(out)
	return out
}
