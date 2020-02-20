package configs

import (
	"reflect"
	"testing"

	"github.com/nginxinc/kubernetes-ingress/pkg/apis/configuration/v1alpha1"
)

func TestParseGlobalConfiguration(t *testing.T) {
	globalConfiguration := &v1alpha1.GlobalConfiguration{
		Spec: v1alpha1.GlobalConfigurationSpec{
			Listeners: []v1alpha1.Listener{
				{
					Name:     "tcp-listener",
					Port:     53,
					Protocol: "TCP",
				},
				{
					Name:     "udp-listener",
					Port:     53,
					Protocol: "UDP",
				},
			},
		},
	}

	expected := &GlobalConfigParams{
		Listeners: map[string]Listener{
			"tcp-listener": {
				Port:     53,
				Protocol: "TCP",
			},
			"udp-listener": {
				Port:     53,
				Protocol: "UDP",
			},
		},
	}

	result := ParseGlobalConfiguration(globalConfiguration)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseGlobalConfiguration() returned \n%+v but expected \n%+v", result, expected)
	}
}
