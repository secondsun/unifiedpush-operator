package unifiedpushserver

import (
	"encoding/base64"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	pushv1alpha1 "github.com/aerogear/unifiedpush-operator/pkg/apis/push/v1alpha1"
	enmassev1beta "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newQueue(cr *pushv1alpha1.UnifiedPushServer, address string) *enmassev1beta.Address {
	name := fmt.Sprintf("ups.%s", strings.ToLower(address))
	return &enmassev1beta.Address{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": cr.Name,
			},
		},
		Spec: enmassev1beta.AddressSpec{
			Address: address,
			Type:    "queue",
			Plan:    "brokered-queue",
		},
	}
}

func newTopic(cr *pushv1alpha1.UnifiedPushServer, address string) *enmassev1beta.Address {

	name := fmt.Sprintf("ups.%s", strings.ToLower(strings.Replace(address, "topic/", "", 1))) //a topic has a prefix.
	return &enmassev1beta.Address{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: cr.Namespace,
			Labels: map[string]string{
				"app": cr.Name,
			},
		},
		Spec: enmassev1beta.AddressSpec{
			Address: address,
			Type:    "topic",
			Plan:    "brokered-topic",
		},
	}
}

func newMessagingUser(cr *pushv1alpha1.UnifiedPushServer) *unstructured.Unstructured {
	password := "password"
	encoded := base64.StdEncoding.EncodeToString([]byte(password))

	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "user.enmasse.io/v1beta1",
			"kind":       "MessagingUser",
			"metadata": map[string]interface{}{
				"name":      "ups.upsuser",
				"namespace": cr.Namespace,
				"labels": map[string]interface{}{
					"app":     cr.Name,
					"service": fmt.Sprintf("%s-%s", cr.Name, "ups.upsuser"),
				},
			},
			"spec": map[string]interface{}{
				"username": "upsuser",
				"authentication": map[string]interface{}{
					"type":     "password",
					"password": encoded,
				},
				"authorization": []map[string]interface{}{
					map[string]interface{}{
						"addresses": []string{
							"*",
						},
						"operations": []string{
							"send",
							"recv",
						},
					},
				},
			},
		},
	}
}

func newAddressSpace(cr *pushv1alpha1.UnifiedPushServer) *enmassev1beta.AddressSpace {
	return &enmassev1beta.AddressSpace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ups",
			Namespace: cr.Namespace,
			Labels:    labels(cr, "ups"),
		},
		Spec: enmassev1beta.AddressSpaceSpec{
			Type: "brokered",
			Plan: "brokered-single-broker",
		},
	}
}
